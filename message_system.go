package golm

import (
	"fmt"
	"strings"
	"text/template"
)

type SystemMessage struct {
	Content string
}

func (message SystemMessage) String() string {
	return "[System]:\n" + message.Content
}

func NewSystemMessage(content string) *SystemMessage {
	return &SystemMessage{Content: content}
}

func NewSystemMessageF(format string, args ...interface{}) *SystemMessage {
	return &SystemMessage{Content: fmt.Sprintf(format, args...)}
}

func NewSystemMessageT(t *template.Template, tName string, data any) (*SystemMessage, error) {
	builder := new(strings.Builder)

	if err := t.ExecuteTemplate(builder, tName, data); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return NewSystemMessage(builder.String()), nil
}
