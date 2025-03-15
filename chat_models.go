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

type SystemMessage struct {
	Content string
}

func (message SystemMessage) String() string {
	return "[System]:\n" + message.Content
}

func NewSystemMessage(content string) *SystemMessage {
	return &SystemMessage{Content: content}
}

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

type CompletionParams struct {
	Temperature     *float64
	MaxOutputLength int
	JSON            bool
}
