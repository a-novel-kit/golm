package models

import (
	"encoding/json"
	"fmt"
)

type UserMessage struct {
	// An optional name for the participant. Provides the model information to differentiate between participants of
	// the same role.
	Name string
	// The contents of the user message.
	Content MultipartMessage
}

func (message UserMessage) Message() Message {
	return message
}

func (message UserMessage) MarshalJSON() ([]byte, error) {
	out := struct {
		Role    MessageRole      `json:"role"`
		Name    string           `json:"name,omitempty"`
		Content MultipartMessage `json:"content"`
	}{
		Role:    MessageRoleUser,
		Name:    message.Name,
		Content: message.Content,
	}

	return json.Marshal(out)
}

func (message *UserMessage) UnmarshalJSON(data []byte) error {
	messageData := struct {
		Name    string              `json:"name"`
		Content MultipartMessageAny `json:"content"`
		Role    MessageRole         `json:"role"`
	}{}

	if err := json.Unmarshal(data, &messageData); err != nil {
		return fmt.Errorf("UserMessage.UnmarshalJSON: %w", err)
	}

	if messageData.Role != MessageRoleUser {
		return fmt.Errorf("UserMessage.UnmarshalJSON: invalid role %s", messageData.Role)
	}

	message.Name = messageData.Name
	message.Content = messageData.Content.MultipartMessage()

	return nil
}
