package groq

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/bindings/groq/models"
)

var ErrRawQuery = errors.New("groq.Binding.RawQuery")

func NewErrRawQuery(err error) error {
	return errors.Join(err, ErrRawQuery)
}

func (binding *Binding) processHistory(history golm.ChatHistory) models.Messages {
	messages := make(models.Messages, 0, len(history.History)+1)

	if history.System != nil {
		messages = append(messages, models.SystemMessage{
			Content: history.System.Content,
		})
	}

	for _, message := range history.History {
		switch message.Role() {
		case golm.MessageRoleUser:
			messages = append(messages, models.UserMessage{
				Content: models.NewMultipartStaticMessage(message.GetContent()),
			})
		case golm.MessageRoleAssistant:
			messages = append(messages, models.AssistantMessage{
				Content: message.GetContent(),
			})
		}
	}

	return messages
}

func (binding *Binding) RawQuery(
	ctx context.Context, request models.ChatCompletionRequest, history golm.ChatHistory,
) (*http.Response, error) {
	// If request has not set a custom history, set it.
	if len(request.Messages) == 0 {
		request.Messages = binding.processHistory(history)
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, NewErrRawQuery(fmt.Errorf("marshal body: %w", err))
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, binding.endpoint+ChatCompletionRoute, bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, NewErrRawQuery(fmt.Errorf("create request: %w", err))
	}

	req.Header.Set("Authorization", "Bearer "+binding.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, NewErrRawQuery(fmt.Errorf("do request: %w", err))
	}

	return resp, nil
}
