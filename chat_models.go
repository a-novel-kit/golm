package golm

import "fmt"

type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
)

type MessageWithRole interface {
	Role() MessageRole
	GetContent() string
	fmt.Stringer
}

type CompletionParams struct {
	Temperature     *float64
	MaxOutputLength int
	User            string
	JSON            bool
}
