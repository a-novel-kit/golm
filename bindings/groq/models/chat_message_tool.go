package models

import (
	"encoding/json"
	"fmt"
)

type ToolCallMessage struct {
	// The contents of the tool message.
	Content string
	// Tool call that this message is responding to.
	ToolCallID string
}

func (message ToolCallMessage) Message() Message {
	return message
}

func (message ToolCallMessage) MarshalJSON() ([]byte, error) {
	out := struct {
		Role       MessageRole `json:"role"`
		Content    string      `json:"content"`
		ToolCallID string      `json:"tool_call_id,omitempty"`
	}{
		Role:       MessageRoleTool,
		Content:    message.Content,
		ToolCallID: message.ToolCallID,
	}

	return json.Marshal(out)
}

func (message *ToolCallMessage) UnmarshalJSON(data []byte) error {
	messageData := struct {
		Content    string      `json:"content"`
		ToolCallID string      `json:"tool_call_id"`
		Role       MessageRole `json:"role"`
	}{}

	if err := json.Unmarshal(data, &messageData); err != nil {
		return fmt.Errorf("ToolCallMessage.UnmarshalJSON: %w", err)
	}

	if messageData.Role != MessageRoleTool {
		return fmt.Errorf("ToolCallMessage.UnmarshalJSON: invalid role %s", messageData.Role)
	}

	message.Content = messageData.Content
	message.ToolCallID = messageData.ToolCallID

	return nil
}
