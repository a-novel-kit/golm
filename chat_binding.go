package golm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/samber/lo"

	"github.com/a-novel-kit/golm/utils"
)

// ChatBinding represents the minimum set of methods that a binding should implement, for it to be turned into a
// Chat instance.
type ChatBinding[RawRequest, RawResponse, StreamResponse any] interface {
	RawQuery(ctx context.Context, request RawRequest, history ChatHistory) (response RawResponse, err error)
	Completion(
		ctx context.Context, message UserMessage, options CompletionParams, history ChatHistory,
	) (response *AssistantMessage, err error)
	CompletionStream(
		ctx context.Context, message UserMessage, options CompletionParams, history ChatHistory,
	) (response <-chan StreamResponse, wait utils.StreamWaitFn)
	StreamResponseToMessage(response StreamResponse) (message string)
}

var _ Chat[any, any] = (*ChatWithBinding[any, any, any])(nil)

type ChatWithBinding[RawRequest, RawResponse, StreamResponse any] struct {
	binding ChatBinding[RawRequest, RawResponse, StreamResponse]
	ChatHistory
}

func (chat *ChatWithBinding[RawRequest, RawResponse, StreamResponse]) RawQuery(
	ctx context.Context, request RawRequest,
) (RawResponse, error) {
	return chat.binding.RawQuery(ctx, request, chat.ChatHistory.GetHistory())
}

func (chat *ChatWithBinding[RawRequest, RawResponse, StreamResponse]) Completion(
	ctx context.Context, message UserMessage, options CompletionParams,
) (*AssistantMessage, error) {
	resp, err := chat.binding.Completion(ctx, message, options, chat.ChatHistory.GetHistory())

	if err == nil {
		chat.PushHistory(message, lo.FromPtr(resp))
	}

	return resp, err
}

func (chat *ChatWithBinding[RawRequest, RawResponse, StreamResponse]) CompletionJSON(
	ctx context.Context, message UserMessage, options CompletionParams, dest any,
) error {
	resp, err := chat.Completion(ctx, message, options)
	if err != nil {
		return fmt.Errorf("completion: %w", err)
	}

	err = json.Unmarshal([]byte(resp.Content), dest)
	if err != nil {
		return errors.Join(
			fmt.Errorf("unmarshal response: %w", err),
			errors.New(resp.Content),
		)
	}

	return nil
}

func (chat *ChatWithBinding[RawRequest, RawResponse, StreamResponse]) CompletionStream(
	ctx context.Context, message UserMessage, options CompletionParams,
) (<-chan string, utils.StreamWaitFn) {
	var assistantMessage AssistantMessage

	out, waitFn := chat.binding.CompletionStream(ctx, message, options, chat.ChatHistory.GetHistory())

	// As we need to catch the returned channel to keep history up-to-date, we need to wrap the returned channel
	// into a new channel.
	userOut, cpWaitFn := utils.NewStreamer[string](ctx, func(ctx context.Context, inC chan<- string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case response, ok := <-out:
			// Stream closed.
			if !ok {
				return utils.ErrStreamerClosed
			}

			decodedMessage := chat.binding.StreamResponseToMessage(response)

			// Add the content to our internal history, then forward it to the user output.
			assistantMessage.Content += decodedMessage
			inC <- decodedMessage
		}

		return nil
	})

	return userOut, func() error {
		upstreamErr := waitFn()
		cpErr := cpWaitFn()

		if err := errors.Join(upstreamErr, cpErr); err != nil {
			return err
		}

		chat.PushHistory(message, assistantMessage)

		return nil
	}
}
