package groq

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/bindings/groq/models"
	"github.com/a-novel-kit/golm/utils"
)

const ChatCompletionStreamPrefix = "data: "

var ErrCompletionStream = errors.New("groq.Binding.ErrCompletionStream")

func NewErrCompletionStream(err error) error {
	return errors.Join(err, ErrCompletion)
}

func (binding *Binding) StreamResponseToMessage(response models.ChatCompletionChunkResponse) string {
	if len(response.Choices) == 0 {
		return ""
	}

	return response.Choices[0].Delta.Content
}

func (binding *Binding) CompletionStream(
	ctx context.Context, message golm.UserMessage, options golm.CompletionParams, history golm.ChatHistory,
) (<-chan models.ChatCompletionChunkResponse, utils.StreamWaitFn) {
	request := binding.buildRequest(models.ChatCompletionRequest{Stream: models.NewStream(true)}, options)

	history.PushHistory(message)

	resp, err := binding.RawQuery(ctx, request, history) //nolint:bodyclose
	if err != nil {
		return nil, func() error { return NewErrCompletionStream(err) }
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)

		body, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			err = errors.Join(err, fmt.Errorf("read body: %w", bodyErr))
		} else {
			err = errors.Join(err, fmt.Errorf("body: %s", string(body)))
		}

		_ = resp.Body.Close()

		return nil, func() error { return NewErrCompletionStream(err) }
	}

	reader := bufio.NewReader(resp.Body)

	out, waitFn := utils.NewStreamer[models.ChatCompletionChunkResponse](
		ctx,
		func(ctx context.Context, inC chan<- models.ChatCompletionChunkResponse) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var line []byte

				// Groq returns streams as pseudo-json strings, with the following format:
				//
				//  data: {...}
				//
				//  data: {...}
				//
				//  data: [DONE]
				//
				// Where the `data: [DONE]` represents the end of the stream, and each other lines
				// data is followed by a json chunk. Because the full line is not valid JSON, we must first read it
				// as a raw string, then depending on its content, extract the actual JSON chunk from it.

				line, err = reader.ReadBytes('\n')
				if err != nil {
					return fmt.Errorf("read line: %w", err)
				}

				// Ignore empty lines.
				if string(line) == "\n" {
					return nil
				}

				// Stream end.
				if string(line) == "data: [DONE]\n" {
					return utils.ErrStreamerClosed
				}

				if !strings.HasPrefix(string(line), ChatCompletionStreamPrefix) {
					return fmt.Errorf("invalid line: '%s'", string(line))
				}

				// Extract the JSON chunk from the line.
				var chunk models.ChatCompletionChunkResponse

				if err = json.Unmarshal(line[len(ChatCompletionStreamPrefix):], &chunk); err != nil {
					return errors.Join(fmt.Errorf("unmarshal chunk: %w", err), errors.New(string(line)))
				}

				inC <- chunk
			}

			return nil
		},
	)

	return out, func() error {
		if err = errors.Join(resp.Body.Close(), waitFn()); err != nil {
			return NewErrCompletionStream(err)
		}

		return nil
	}
}
