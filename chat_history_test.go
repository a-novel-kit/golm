package golm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
)

func TestCopyGetHistory(t *testing.T) {
	t.Parallel()

	history := golm.ChatHistory{
		System: &golm.SystemMessage{Content: "system 1"},
		History: []golm.MessageWithRole{
			golm.UserMessage{Content: "user 1 - 1"},
			golm.AssistantMessage{Content: "assistant 2 - 1"},
		},
	}

	copyHistory := history.GetHistory()

	require.Equal(t, history, copyHistory)

	copyHistory.System.Content = "system 2"
	copyHistory.History = []golm.MessageWithRole{
		golm.UserMessage{Content: "user 1 - 2"},
	}

	require.NotEqual(t, history, copyHistory)
}
