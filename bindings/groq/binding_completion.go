package groq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/a-novel-kit/golm"
	"github.com/a-novel-kit/golm/bindings/groq/models"
)

var ErrNoCompletionData = errors.New("no completion data found")

var ErrCompletion = errors.New("groq.Binding.ErrCompletion")

func NewErrCompletion(err error) error {
	return errors.Join(err, ErrCompletion)
}

func (binding *Binding) buildRequest(
	base models.ChatCompletionRequest, options golm.CompletionParams,
) models.ChatCompletionRequest {
	request := base

	request.Model = binding.model
	request.Temperature = (*models.Temperature)(options.Temperature)

	if options.MaxOutputLength > 0 {
		request.MaxCompletionTokens = (*models.MaxCompletionTokens)(&options.MaxOutputLength)
	}

	if options.JSON {
		request.ResponseFormat = models.ResponseFormatJSON
	}

	return request
}

func (binding *Binding) Completion(
	ctx context.Context, message golm.UserMessage, options golm.CompletionParams, history golm.ChatHistory,
) (*golm.AssistantMessage, error) {
	request := binding.buildRequest(models.ChatCompletionRequest{}, options)

	history.PushHistory(message)

	resp, err := binding.RawQuery(ctx, request, history)
	if err != nil {
		return nil, NewErrCompletion(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)

		body, bodyErr := io.ReadAll(resp.Body)
		if bodyErr != nil {
			err = errors.Join(err, fmt.Errorf("read body: %w", bodyErr))
		} else {
			err = errors.Join(err, fmt.Errorf("body: %s", string(body)))
		}

		return nil, NewErrCompletion(err)
	}

	var response models.ChatCompletionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, NewErrCompletion(fmt.Errorf("decode response: %w", err))
	}

	if len(response.Choices) == 0 {
		return nil, NewErrCompletion(ErrNoCompletionData)
	}

	return &golm.AssistantMessage{
		Content: response.Choices[0].Message.Content,
	}, nil
}
