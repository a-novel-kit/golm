package groq

import (
	"net/http"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/bindings/groq/models"
)

const DefaultEndpoint = "https://api.groq.com/openai/v1"

const ChatCompletionRoute = "/chat/completions"

var _ golm.ChatBinding[
	models.ChatCompletionRequest,
	*http.Response,
	models.ChatCompletionChunkResponse,
] = (*Binding)(nil)

type Binding struct {
	apiKey   string
	endpoint string
	model    models.Model
}

// New returns a new groq binding. You may use a single binding instance for all your chats.
func New(apiKey string, model models.Model) *Binding {
	return &Binding{
		apiKey:   apiKey,
		endpoint: DefaultEndpoint,
		model:    model,
	}
}
