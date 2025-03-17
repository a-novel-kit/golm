package golm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
)

func TestNewUserMessage(t *testing.T) {
	t.Parallel()

	UserMessage := golm.NewUserMessage("content")

	require.Equal(t, "[User]:\ncontent", UserMessage.String())
}

func TestNewUserMessageF(t *testing.T) {
	t.Parallel()

	UserMessage := golm.NewUserMessageF("content %s", "formatted")

	require.Equal(t, "[User]:\ncontent formatted", UserMessage.String())
}

func TestNewUserMessageT(t *testing.T) {
	t.Parallel()

	UserMessage, err := golm.NewUserMessageT(TestTemplate, "test", TestTemplateData{Content: "foobar"})

	require.NoError(t, err)
	require.Equal(t, "[User]:\nfoobar", UserMessage.String())
}
