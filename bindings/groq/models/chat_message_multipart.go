package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrUnknownMessageContent = errors.New("unknown message content")

// =====================================================================================================================
// MULTIPART MESSAGE PART
// Represents the single element of a MultipartMessage in list format.
// =====================================================================================================================

type MultipartMessagePart interface {
	MultipartMessagePart() MultipartMessagePart
}

type MultipartMessageTextContent struct {
	Text string `json:"text"`
	Type string `json:"type,omitempty"`
}

func (message MultipartMessageTextContent) MultipartMessagePart() MultipartMessagePart {
	return message
}

type MultipartMessageImageContent struct {
	ImageURL MessageImageURL `json:"image_url"`
	Type     string          `json:"type,omitempty"`
}

func (message MultipartMessageImageContent) MultipartMessagePart() MultipartMessagePart {
	return message
}

type MessageImageURL struct {
	// Specifies the detail level of the image. Defaults to `auto`.
	Detail string `json:"detail,omitempty"`
	// Either a URL of the image or the base64 encoded image data.
	URL string `json:"url"`
}

// =====================================================================================================================
// MULTIPART MESSAGE
// Represents a message that can be either a single string or a list of multipart message parts.
// =====================================================================================================================

type MultipartMessage interface {
	MultipartMessage() MultipartMessage
}

type MultipartMessageList []MultipartMessagePart

func (message MultipartMessageList) MultipartMessage() MultipartMessage {
	return message
}

func (message MultipartMessageList) unmarshalListElement(src json.RawMessage) (MultipartMessagePart, error) {
	var out map[string]json.RawMessage

	if err := json.Unmarshal(src, &out); err != nil {
		return nil, fmt.Errorf("unmarshal base map: %w", err)
	}

	switch {
	case out["text"] != nil:
		var text MultipartMessageTextContent
		if err := json.Unmarshal(src, &text); err != nil {
			return nil, fmt.Errorf("unmarshal text content: %w", err)
		}

		return text, nil
	case out["image_url"] != nil:
		var image MultipartMessageImageContent
		if err := json.Unmarshal(src, &image); err != nil {
			return nil, fmt.Errorf("unmarshal image content: %w", err)
		}

		return image, nil
	default:
		return nil, fmt.Errorf("unmarshal list element: %w", ErrUnknownMessageContent)
	}
}

func (message *MultipartMessageList) UnmarshalJSON(data []byte) error {
	var listOut []json.RawMessage

	if err := json.Unmarshal(data, &listOut); err != nil {
		return fmt.Errorf("json.Unmarshal MultipartMessage: %w", err)
	}

	*message = make([]MultipartMessagePart, len(listOut))

	for i, raw := range listOut {
		element, err := message.unmarshalListElement(raw)
		if err != nil {
			return fmt.Errorf("unmarshal list element %d: %w", i, err)
		}

		(*message)[i] = element
	}

	return nil
}

type MultipartMessageStatic string

func (message MultipartMessageStatic) MultipartMessage() MultipartMessage {
	return message
}

// =====================================================================================================================
// MULTIPART MESSAGE WRAPPER
// =====================================================================================================================

// MultipartMessageAny ios a generic wrapper over MultipartMessageList and MultipartMessageStatic. It implements
// custom JSON marshal methods top handle both in a single interface.
type MultipartMessageAny struct {
	messageContent MultipartMessage
}

func (message MultipartMessageAny) MultipartMessage() MultipartMessage {
	return message.messageContent
}

func (message MultipartMessageAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(message.messageContent)
}

func (message *MultipartMessageAny) UnmarshalJSON(data []byte) error {
	// Try to decode it as array first.
	messageContent := make(MultipartMessageList, 0)

	if err := json.Unmarshal(data, &messageContent); err == nil {
		message.messageContent = messageContent

		return nil
	}

	// If it's not an array, try to decode it as a string.
	var staticContent MultipartMessageStatic

	if err := json.Unmarshal(data, &staticContent); err != nil {
		return fmt.Errorf("MultipartMessageAny.UnmarshalJSON: %w", err)
	}

	message.messageContent = staticContent

	return nil
}

func NewMultipartMessage(parts ...MultipartMessagePart) MultipartMessageList {
	return parts
}

func NewMultipartStaticMessage(content string) MultipartMessageStatic {
	return MultipartMessageStatic(content)
}
