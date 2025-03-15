package golm

import (
	"context"

	"github.com/a-novel-kit/golm/utils"
)

// ChatRaw is an interface used to access the raw implementation of a LLM binding. It lets you tinkle with all the
// options offered by the binding, with the downside of letting you in charge of parsing the response and managing
// messages history.
type ChatRaw[RawRequest, RawResponse any] interface {
	RawQuery(ctx context.Context, request RawRequest) (response RawResponse, err error)
}

type ChatBase interface {
	// SetHistory replaces the current history with the provided one.
	SetHistory(history ChatHistory)
	// GetHistory returns the current history.
	//
	// The returned history is a copy of the one used internally, so you can modify it without affecting the interna
	// state.
	GetHistory() ChatHistory
	// PushHistory adds the provided messages to the history.
	PushHistory(messages ...MessageWithRole)
	// SetSystem replaces the current system message with the provided one.
	SetSystem(system *SystemMessage)

	// Completion returns the LLM response to the provided message, based on user history.
	//
	// The LLM response is automatically added to history.
	Completion(
		ctx context.Context, message UserMessage, options CompletionParams,
	) (response *AssistantMessage, err error)
	// CompletionJSON ets the LLM response to the provided message, and attempts to parse it into the desired
	// destination.
	//
	// You might need to instruct your model to return JSON depending on the selected binding.
	//
	// The LLM response is automatically added to history.
	CompletionJSON(ctx context.Context, message UserMessage, options CompletionParams, dest any) (err error)
	// CompletionStream returns a channel that will be filled with the LLM response to the provided message, as
	// data bits are received from the LLM. The channel must be automatically closed by the implementation, so you
	// just have to worry about reading from it.
	//
	// The LLM response is automatically added to history. If the stream is closed early, then the already read bits
	// might form an incomplete response in the history.
	CompletionStream(
		ctx context.Context, message UserMessage, options CompletionParams,
	) (response <-chan string, wait utils.StreamWaitFn)
}

type Chat[RawRequest, RawResponse any] interface {
	ChatBase
	ChatRaw[RawRequest, RawResponse]
}

// NewChat returns a new Chat instance based on the provided binding.
//
// Since a chat instance maintains its own history, it is recommended to create a new chat instance for each
// new process.
func NewChat[Req, Res, Stream any](binding ChatBinding[Req, Res, Stream]) Chat[Req, Res] {
	return &ChatWithBinding[Req, Res, Stream]{binding: binding}
}
