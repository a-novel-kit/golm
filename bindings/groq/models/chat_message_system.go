package models

import (
	"encoding/json"
	"fmt"
	"strings"
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
	serializedContent, err := json.Marshal(message.Content)
	if err != nil {
		return nil, fmt.Errorf("SystemMessage.MarshalJSON: %w", err)
	}

	messageData := []string{`"role":"` + string(MessageRoleSystem) + `"`, `"content":` + string(serializedContent)}

	if message.Name != "" {
		messageData = append(messageData, `"name":"`+message.Name+`"`)
	}

	return []byte("{" + strings.Join(messageData, ",") + "}"), nil
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
