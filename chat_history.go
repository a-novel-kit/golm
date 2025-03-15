package golm

import "strings"

type ChatHistory struct {
	System  *SystemMessage
	History []MessageWithRole
}

func (chatHistory *ChatHistory) SetHistory(history ChatHistory) {
	*chatHistory = history
}

func (chatHistory *ChatHistory) GetHistory() ChatHistory {
	output := ChatHistory{
		History: make([]MessageWithRole, 0, len(chatHistory.History)),
	}

	if chatHistory.System != nil {
		output.System = &SystemMessage{
			Content: chatHistory.System.Content,
		}
	}

	for _, message := range chatHistory.History {
		switch typ := message.(type) {
		case UserMessage, *UserMessage:
			output.History = append(output.History, UserMessage{
				Content: typ.GetContent(),
			})
		case AssistantMessage, *AssistantMessage:
			output.History = append(output.History, AssistantMessage{
				Content: typ.GetContent(),
			})
		}
	}

	return output
}

func (chatHistory *ChatHistory) PushHistory(messages ...MessageWithRole) {
	chatHistory.History = append(chatHistory.History, messages...)
}

func (chatHistory *ChatHistory) SetSystem(system *SystemMessage) {
	chatHistory.System = system
}

func (chatHistory *ChatHistory) String() string {
	messages := make([]string, 0, len(chatHistory.History)+1)

	if chatHistory.System != nil {
		messages = append(messages, chatHistory.System.String())
	}

	for _, message := range chatHistory.History {
		messages = append(messages, message.String())
	}

	return strings.Join(messages, "\n\n")
}
