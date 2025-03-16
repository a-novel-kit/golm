package models

import (
	"encoding/json"
	"fmt"
)

type SystemMessage struct {
	// An optional name for the participant. Provides the model information to differentiate between participants of
	// the same role.
	Name string
	// The contents of the system message.
	Content string
}

func (message SystemMessage) Message() Message {
	return message
}

func (message SystemMessage) MarshalJSON() ([]byte, error) {
	out := struct {
		Role    MessageRole `json:"role"`
		Name    string      `json:"name,omitempty"`
		Content string      `json:"content"`
	}{
		Role:    MessageRoleSystem,
		Name:    message.Name,
		Content: message.Content,
	}

	return json.Marshal(out)
}

func (message *SystemMessage) UnmarshalJSON(data []byte) error {
	messageData := struct {
		Name    string      `json:"name"`
		Content string      `json:"content"`
		Role    MessageRole `json:"role"`
	}{}

	if err := json.Unmarshal(data, &messageData); err != nil {
		return fmt.Errorf("SystemMessage.UnmarshalJSON: %w", err)
	}

	if messageData.Role != MessageRoleSystem {
		return fmt.Errorf("SystemMessage.UnmarshalJSON: invalid role %s", messageData.Role)
	}

	message.Name = messageData.Name
	message.Content = messageData.Content

	return nil
}
