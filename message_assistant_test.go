package golm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
)

func TestNewAssistantMessage(t *testing.T) {
	t.Parallel()

	assistantMessage := golm.NewAssistantMessage("content")

	require.Equal(t, "[Assistant]:\ncontent", assistantMessage.String())
}

func TestNewAssistantMessageF(t *testing.T) {
	t.Parallel()

	assistantMessage := golm.NewAssistantMessageF("content %s", "formatted")

	require.Equal(t, "[Assistant]:\ncontent formatted", assistantMessage.String())
}

func TestNewAssistantMessageT(t *testing.T) {
	t.Parallel()

	assistantMessage, err := golm.NewAssistantMessageT(TestTemplate, "test", TestTemplateData{Content: "foobar"})

	require.NoError(t, err)
	require.Equal(t, "[Assistant]:\nfoobar", assistantMessage.String())
}
