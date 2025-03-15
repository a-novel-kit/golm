package models

import (
	"encoding/json"
	"fmt"
	"strings"
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
	serializedContent, err := json.Marshal(message.Content)
	if err != nil {
		return nil, fmt.Errorf("UserMessage.MarshalJSON: %w", err)
	}

	messageData := []string{`"role":"` + string(MessageRoleUser) + `"`, `"content":` + string(serializedContent)}

	if message.Name != "" {
		messageData = append(messageData, `"name":"`+message.Name+`"`)
	}

	return []byte("{" + strings.Join(messageData, ",") + "}"), nil
}

func (message *UserMessage) UnmarshalJSON(data []byte) error {
	messageData := struct {
		Name    string                   `json:"name"`
		Content MultipartOrStaticMessage `json:"content"`
		Role    MessageRole              `json:"role"`
	}{}

	if err := json.Unmarshal(data, &messageData); err != nil {
		return fmt.Errorf("UserMessage.UnmarshalJSON: %w", err)
	}

	if messageData.Role != MessageRoleUser {
		return fmt.Errorf("UserMessage.UnmarshalJSON: invalid role %s", messageData.Role)
	}

	message.Name = messageData.Name
	message.Content = messageData.Content

	return nil
}
