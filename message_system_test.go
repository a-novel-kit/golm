package golm_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm"
)

func TestNewSystemMessage(t *testing.T) {
	t.Parallel()

	SystemMessage := golm.NewSystemMessage("content")

	require.Equal(t, "[System]:\ncontent", SystemMessage.String())
}

func TestNewSystemMessageF(t *testing.T) {
	t.Parallel()

	SystemMessage := golm.NewSystemMessageF("content %s", "formatted")

	require.Equal(t, "[System]:\ncontent formatted", SystemMessage.String())
}

func TestNewSystemMessageT(t *testing.T) {
	t.Parallel()

	SystemMessage, err := golm.NewSystemMessageT(TestTemplate, "test", TestTemplateData{Content: "foobar"})

	require.NoError(t, err)
	require.Equal(t, "[System]:\nfoobar", SystemMessage.String())
}
