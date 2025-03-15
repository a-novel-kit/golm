package testutils

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
)

func BindingImplementation[RawRequest, RawResponse, StreamResponse any](
	t *testing.T, binding golm.ChatBinding[RawRequest, RawResponse, StreamResponse],
) {
	t.Helper()

	t.Run("ChatCompletion", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name string

			message golm.UserMessage
			options golm.CompletionParams
			history golm.ChatHistory

			expect *golm.AssistantMessage
		}{
			{
				history: golm.ChatHistory{
					System: &golm.SystemMessage{
						Content: `You are a counter. Just count. Return your answer as a list of numbers separated by
							commas. Example: 1,2,3`,
					},
				},
				options: golm.CompletionParams{
					Temperature: lo.ToPtr(float64(0)),
				},
				message: golm.UserMessage{Content: "Count from 1 to ten."},

				expect: &golm.AssistantMessage{
					Content: "1,2,3,4,5,6,7,8,9,10",
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				chat := golm.NewChat(binding)
				chat.SetHistory(testCase.history)

				response, err := chat.Completion(t.Context(), testCase.message, testCase.options)
				require.NoError(t, err)
				require.Equal(t, testCase.expect, response)
			})
		}
	})

	t.Run("ChatCompletionJSON", func(t *testing.T) {
		t.Parallel()

		whitespaceRegexp := regexp.MustCompile(`\s+`)

		testCases := []struct {
			name string

			message golm.UserMessage
			options golm.CompletionParams
			history golm.ChatHistory

			expect *golm.AssistantMessage
		}{
			{
				history: golm.ChatHistory{
					System: &golm.SystemMessage{
						Content: `You are a counter. Just count. Return your answer as a JSON array of numbers, 
							separated by commas.
							Wrap your array in the key "count". Example: { "count": [1,2,3] }`,
					},
				},
				options: golm.CompletionParams{
					Temperature: lo.ToPtr(float64(0)),
				},
				message: golm.UserMessage{Content: "Count from 1 to ten."},

				expect: &golm.AssistantMessage{
					Content: "{\"count\":[1,2,3,4,5,6,7,8,9,10]}",
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				chat := golm.NewChat(binding)
				chat.SetHistory(testCase.history)

				response, err := chat.Completion(t.Context(), testCase.message, testCase.options)
				require.NoError(t, err)

				// Remove whitespaces and line breaks from the response.
				response.Content = whitespaceRegexp.ReplaceAllString(response.Content, "")
				require.Equal(t, testCase.expect, response)
			})
		}
	})

	t.Run("Stream", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name string

			message golm.UserMessage
			options golm.CompletionParams
			history golm.ChatHistory

			expect *golm.AssistantMessage
		}{
			{
				history: golm.ChatHistory{
					System: &golm.SystemMessage{
						Content: `You are a counter. Just count. Return your answer as a list of numbers separated by
							commas. Example: 1,2,3`,
					},
				},
				options: golm.CompletionParams{
					Temperature: lo.ToPtr(float64(0)),
				},
				message: golm.UserMessage{Content: "Count from 1 to ten."},

				expect: &golm.AssistantMessage{
					Content: "1,2,3,4,5,6,7,8,9,10",
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				t.Parallel()

				chat := golm.NewChat(binding)
				chat.SetHistory(testCase.history)

				timedCtx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
				t.Cleanup(cancel)

				outC, waitFn := chat.CompletionStream(timedCtx, testCase.message, testCase.options)

				response := &golm.AssistantMessage{}

				for s := range outC {
					response.Content += s
				}

				require.NoError(t, waitFn())
				require.Equal(t, testCase.expect, response)
			})
		}
	})
}
