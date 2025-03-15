package golm_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/mocks"
	"github.com/a-novel-kit/golm/utils"
)

func TestChatBindingRawQuery(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type rawQueryData struct {
		resp string
		err  error
	}

	testCases := []struct {
		name string

		request string
		history golm.ChatHistory

		data rawQueryData

		expect    string
		expectErr error
	}{
		{
			name: "Success",

			request: "request",
			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			data: rawQueryData{
				resp: "response",
			},

			expect: "response",
		},
		{
			name: "Error",

			request: "request",
			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			data: rawQueryData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			binding := mocks.NewMockChatBinding[string, string, string](t)

			chat := golm.NewChat[string, string, string](binding)
			chat.SetHistory(testCase.history)

			binding.EXPECT().
				RawQuery(t.Context(), testCase.request, testCase.history).
				Return(testCase.data.resp, testCase.data.err)

			resp, err := chat.RawQuery(t.Context(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, resp)

			binding.AssertExpectations(t)
		})
	}
}

func TestChatBindingCompletion(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type completionData struct {
		resp *golm.AssistantMessage
		err  error
	}

	testCases := []struct {
		name string

		message golm.UserMessage
		options golm.CompletionParams
		history golm.ChatHistory

		completionData completionData

		expect    *golm.AssistantMessage
		expectErr error
	}{
		{
			name: "Success",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				resp: &golm.AssistantMessage{Content: "response"},
			},

			expect: &golm.AssistantMessage{Content: "response"},
		},
		{
			name: "Error",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			binding := mocks.NewMockChatBinding[string, string, string](t)

			chat := golm.NewChat[string, string, string](binding)
			chat.SetHistory(testCase.history)

			binding.EXPECT().
				Completion(t.Context(), testCase.message, testCase.options, testCase.history).
				Return(testCase.completionData.resp, testCase.completionData.err)

			resp, err := chat.Completion(t.Context(), testCase.message, testCase.options)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, resp)

			if testCase.completionData.err == nil {
				testCase.history.PushHistory(testCase.message, lo.FromPtr(testCase.completionData.resp))
			}

			require.Equal(t, testCase.history, chat.GetHistory())

			binding.AssertExpectations(t)
		})
	}
}

func TestChatBindingCompletionJSON(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type completionData struct {
		resp *golm.AssistantMessage
		err  error
	}

	testCases := []struct {
		name string

		message golm.UserMessage
		options golm.CompletionParams
		history golm.ChatHistory
		dest    any

		completionData completionData

		expect    any
		expectErr error
	}{
		{
			name: "Success",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				resp: &golm.AssistantMessage{Content: `{"foo":"bar"}`},
			},

			dest: &struct {
				Foo string `json:"foo"`
			}{},
			expect: &struct {
				Foo string `json:"foo"`
			}{Foo: "bar"},
		},
		{
			name: "Error",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			binding := mocks.NewMockChatBinding[string, string, string](t)

			chat := golm.NewChat[string, string, string](binding)
			chat.SetHistory(testCase.history)

			binding.EXPECT().
				Completion(t.Context(), testCase.message, testCase.options, testCase.history).
				Return(testCase.completionData.resp, testCase.completionData.err)

			err := chat.CompletionJSON(t.Context(), testCase.message, testCase.options, testCase.dest)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, testCase.dest)

			if testCase.completionData.err == nil {
				testCase.history.PushHistory(testCase.message, lo.FromPtr(testCase.completionData.resp))
			}

			require.Equal(t, testCase.history, chat.GetHistory())

			binding.AssertExpectations(t)
		})
	}
}

func TestChatBindingCompletionStream(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type completionData struct {
		resp func(ctx context.Context) (<-chan string, utils.StreamWaitFn)
	}

	testCases := []struct {
		name string

		message golm.UserMessage
		options golm.CompletionParams
		history golm.ChatHistory

		completionData completionData

		expect    string
		expectErr error
	}{
		{
			name: "Success",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				resp: func(ctx context.Context) (<-chan string, utils.StreamWaitFn) {
					bits := []string{"r", "e", "s", "p", "o", "n", "s", "e"}
					i := 0

					return utils.NewStreamer[string](ctx, func(_ context.Context, inC chan<- string) error {
						if i < len(bits) {
							inC <- bits[i]
							i++

							return nil
						}

						return utils.ErrStreamerClosed
					})
				},
			},

			expect: "response",
		},
		{
			name: "Error",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				resp: func(ctx context.Context) (<-chan string, utils.StreamWaitFn) {
					return utils.NewStreamer[string](ctx, func(_ context.Context, _ chan<- string) error {
						return errFoo
					})
				},
			},

			expectErr: errFoo,
		},
		{
			name: "Timeout",

			message: golm.UserMessage{Content: "message"},

			options: golm.CompletionParams{
				MaxOutputLength: 128,
			},

			history: golm.ChatHistory{
				System:  &golm.SystemMessage{Content: "system"},
				History: []golm.MessageWithRole{},
			},

			completionData: completionData{
				resp: func(ctx context.Context) (<-chan string, utils.StreamWaitFn) {
					return utils.NewStreamer[string](ctx, func(_ context.Context, _ chan<- string) error {
						time.Sleep(time.Second)

						return utils.ErrStreamerClosed
					})
				},
			},

			expectErr: context.DeadlineExceeded,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			binding := mocks.NewMockChatBinding[string, string, string](t)

			chat := golm.NewChat[string, string, string](binding)
			chat.SetHistory(testCase.history)

			timedCtx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			t.Cleanup(cancel)

			binding.EXPECT().
				CompletionStream(timedCtx, testCase.message, testCase.options, testCase.history).
				Return(testCase.completionData.resp(timedCtx))

			binding.EXPECT().
				StreamResponseToMessage(mock.AnythingOfType("string")).
				RunAndReturn(func(s string) string { return s }).Maybe()

			outC, waitFn := chat.CompletionStream(timedCtx, testCase.message, testCase.options)

			var resp string

			for s := range outC {
				resp += s
			}

			require.ErrorIs(t, waitFn(), testCase.expectErr)
			require.Equal(t, testCase.expect, resp)

			if testCase.expectErr == nil {
				testCase.history.PushHistory(testCase.message, golm.AssistantMessage{Content: resp})
			}

			require.Equal(t, testCase.history, chat.GetHistory())

			binding.AssertExpectations(t)
		})
	}
}
