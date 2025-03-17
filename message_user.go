package golm

import (
	"fmt"
	"strings"
	"text/template"
)

type UserMessage struct {
	Content string
}

func (message UserMessage) Role() MessageRole {
	return MessageRoleUser
}

func (message UserMessage) GetContent() string {
	return message.Content
}

func (message UserMessage) String() string {
	return "[User]:\n" + message.Content
}

func NewUserMessage(content string) UserMessage {
	return UserMessage{Content: content}
}

func NewUserMessageF(format string, args ...interface{}) UserMessage {
	return UserMessage{Content: fmt.Sprintf(format, args...)}
}

func NewUserMessageT(t *template.Template, tName string, data any) (UserMessage, error) {
	builder := new(strings.Builder)

	if err := t.ExecuteTemplate(builder, tName, data); err != nil {
		return UserMessage{}, fmt.Errorf("execute template: %w", err)
	}

	return NewUserMessage(builder.String()), nil
}
