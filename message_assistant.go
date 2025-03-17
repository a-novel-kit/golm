package golm

import (
	"fmt"
	"strings"
	"text/template"
)

type AssistantMessage struct {
	Content string
}

func (message AssistantMessage) Role() MessageRole {
	return MessageRoleAssistant
}

func (message AssistantMessage) GetContent() string {
	return message.Content
}

func (message AssistantMessage) String() string {
	return "[Assistant]:\n" + message.Content
}

func NewAssistantMessage(content string) AssistantMessage {
	return AssistantMessage{Content: content}
}

func NewAssistantMessageF(format string, args ...interface{}) AssistantMessage {
	return AssistantMessage{Content: fmt.Sprintf(format, args...)}
}

func NewAssistantMessageT(t *template.Template, tName string, data any) (AssistantMessage, error) {
	builder := new(strings.Builder)

	if err := t.ExecuteTemplate(builder, tName, data); err != nil {
		return AssistantMessage{}, fmt.Errorf("execute template: %w", err)
	}

	return NewAssistantMessage(builder.String()), nil
}
