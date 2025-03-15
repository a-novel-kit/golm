package groq_test

import (
	"os"
	"testing"

	"github.com/a-novel-kit/golm/bindings/groq"
	"github.com/a-novel-kit/golm/bindings/groq/models"
	"github.com/a-novel-kit/golm/testutils"
)

func TestGroqBindings(t *testing.T) {
	t.Parallel()

	binding := groq.New(os.Getenv("GROQ_TOKEN"), models.ModelLlama3370BVersatile)

	testutils.BindingImplementation(t, binding)
}
