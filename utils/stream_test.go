package utils_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/utils"
)

func TestStreamer(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	testCases := []struct {
		name string

		callback utils.StreamerCallback[string]

		expect    string
		expectErr error
	}{
		{
			name: "Success",

			callback: func(_ context.Context, stream chan<- string) error {
				stream <- "test"

				return utils.ErrStreamerClosed
			},

			expect: "test",
		},
		{
			name: "Error",

			callback: func(_ context.Context, _ chan<- string) error {
				return errFoo
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			timedCtx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
			t.Cleanup(cancel)

			streamer, waitFn := utils.NewStreamer[string](timedCtx, testCase.callback)

			var resp string

			for s := range streamer {
				resp = s
			}

			require.ErrorIs(t, waitFn(), testCase.expectErr)
			require.Equal(t, testCase.expect, resp)
		})
	}
}
