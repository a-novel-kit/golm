package models

import (
	"encoding/json"
	"fmt"
)

// Message is a generic interface implemented by every Groq supported message types.
//
// Each Message implementation is mapped to a MessageRole.
type Message interface {
	Message() Message
}

// MessageRole represents the role of the messages author.
type MessageRole string

const (
	MessageRoleSystem    MessageRole = "system"
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleTool      MessageRole = "tool"
)

// UnmarshalMessage decodes a JSON value into a Message implementation. The returned message type depends on the
// "role" value.
func UnmarshalMessage(data json.RawMessage) (Message, error) {
	var messageData struct {
		Role MessageRole `json:"role"`
	}

	if err := json.Unmarshal(data, &messageData); err != nil {
		return nil, fmt.Errorf("UnmarshalMessage: %w", err)
	}

	switch messageData.Role {
	case MessageRoleSystem:
		var message SystemMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalMessage: %w", err)
		}

		return &message, nil
	case MessageRoleUser:
		var message UserMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalMessage: %w", err)
		}

		return &message, nil
	case MessageRoleAssistant:
		var message AssistantMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalMessage: %w", err)
		}

		return &message, nil
	case MessageRoleTool:
		var message ToolCallMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalMessage: %w", err)
		}

		return &message, nil
	default:
		return nil, fmt.Errorf("UnmarshalMessage: invalid role %s", messageData.Role)
	}
}

// Messages represents a list of Message of unknown type.
type Messages []Message

func NewMessages(messages ...Message) Messages {
	anyMessages := messages

	return anyMessages
}

func (messages *Messages) UnmarshalJSON(data []byte) error {
	rawMessages := make([]json.RawMessage, 0)

	if err := json.Unmarshal(data, &rawMessages); err != nil {
		return fmt.Errorf("decode messages list: %w", err)
	}

	*messages = make(Messages, 0, len(rawMessages))

	for i, rawMessage := range rawMessages {
		message, err := UnmarshalMessage(rawMessage)
		if err != nil {
			return fmt.Errorf("decode message at position %d: %w", i, err)
		}

		(*messages)[i] = message
	}

	return nil
}
