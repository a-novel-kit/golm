package groq

import (
	"net/http"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/bindings/groq/models"
)

const DefaultEndpoint = "https://api.groq.com/openai/v1"

const ChatCompletionRoute = "/chat/completions"

type ChatBinding = golm.ChatBinding[
	models.ChatCompletionRequest,
	*http.Response,
	models.ChatCompletionChunkResponse,
]

type Chat = golm.Chat[models.ChatCompletionRequest, *http.Response]

var _ ChatBinding = (*Binding)(nil)

type Binding struct {
	apiKey   string
	endpoint string
	model    models.Model
}

// New returns a new groq binding. You may use a single binding instance for all your chats.
//
// Requires a Groq Cloud API key: https://console.groq.com/keys
//
// You must also specify a model compatible with chat completion api: https://console.groq.com/docs/models
func New(apiKey string, model models.Model) *Binding {
	return &Binding{
		apiKey:   apiKey,
		endpoint: DefaultEndpoint,
		model:    model,
	}
}
