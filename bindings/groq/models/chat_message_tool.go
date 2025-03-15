package models

import (
	"encoding/json"
	"fmt"
	"strings"
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
	serializedContent, err := json.Marshal(message.Content)
	if err != nil {
		return nil, fmt.Errorf("ToolCallMessage.MarshalJSON: %w", err)
	}

	messageData := []string{`"role":"` + string(MessageRoleTool) + `"`, `"content":` + string(serializedContent)}

	if message.ToolCallID != "" {
		messageData = append(messageData, `"tool_call_id":"`+message.ToolCallID+`"`)
	}

	return []byte("{" + strings.Join(messageData, ",") + "}"), nil
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
