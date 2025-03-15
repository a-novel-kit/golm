package golm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/mocks"
)

func TestContext(t *testing.T) {
	t.Parallel()

	var binding golm.ChatBinding[string, string, string] = mocks.NewMockChatBinding[string, string, string](t)

	ctx := golm.WithContext(t.Context(), binding)

	rawChat := golm.ContextWithRaw[string, string](ctx)
	require.NotNil(t, rawChat)
	require.Equal(t, golm.ChatHistory{History: []golm.MessageWithRole{}}, rawChat.GetHistory())

	testHistory := golm.ChatHistory{
		System:  golm.NewSystemMessage("foo"),
		History: []golm.MessageWithRole{},
	}

	rawChat.SetHistory(testHistory)

	// Refetch chat from context.
	rawChat = golm.ContextWithRaw[string, string](ctx)
	require.NotNil(t, rawChat)
	require.Equal(t, rawChat.GetHistory(), testHistory)

	baseChat := golm.Context(ctx)
	require.NotNil(t, baseChat)
	require.Equal(t, baseChat.GetHistory(), testHistory)
}
