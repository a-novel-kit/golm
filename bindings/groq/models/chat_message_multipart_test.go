package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func TestMarshalMessageParts(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		content models.MultipartMessage

		expect string
	}{
		{
			name: "Simple",

			content: models.NewMultipartStaticMessage("Hello world!"),

			expect: `"Hello world!"`,
		},
		{
			name: "List",

			content: models.NewMultipartMessage(
				models.MultipartMessageTextContent{
					Text: "Message 1",
					Type: "text",
				},
				models.MultipartMessageImageContent{
					ImageURL: models.MessageImageURL{
						Detail: "auto",
						URL:    "https://example.com/image.jpg",
					},
					Type: "image",
				},
				models.MultipartMessageTextContent{
					Text: "Message 2",
					Type: "text",
				},
			),

			expect: `[
				{"text":"Message 1","type":"text"},
				{"image_url":{"detail":"auto","url":"https://example.com/image.jpg"},"type":"image"},
				{"text":"Message 2","type":"text"}
			]`,
		},
		{
			name: "ListMinimal",

			content: models.NewMultipartMessage(
				models.MultipartMessageTextContent{
					Text: "Message 1",
				},
				models.MultipartMessageImageContent{
					ImageURL: models.MessageImageURL{
						URL: "https://example.com/image.jpg",
					},
				},
				models.MultipartMessageTextContent{
					Text: "Message 2",
				},
			),

			expect: `[
				{"text":"Message 1"},
				{"image_url":{"url":"https://example.com/image.jpg"}},
				{"text":"Message 2"}
			]`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			mrsh, err := json.Marshal(testCase.content)
			require.NoError(t, err)

			require.JSONEq(t, testCase.expect, string(mrsh))
		})
	}
}

func TestUnmarshalMessageParts(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		content string

		expect models.MultipartMessage
	}{
		{
			name: "Simple",

			content: `"Hello world!"`,

			expect: models.MultipartMessageStatic("Hello world!"),
		},
		{
			name: "List",

			content: `[
				{"text":"Message 1","type":"text"},
				{"image_url":{"detail":"auto","url":"https://example.com/image.jpg"},"type":"image"},
				{"text":"Message 2","type":"text"}
			]`,

			expect: models.MultipartMessageList{
				models.MultipartMessageTextContent{
					Text: "Message 1",
					Type: "text",
				},
				models.MultipartMessageImageContent{
					ImageURL: models.MessageImageURL{
						Detail: "auto",
						URL:    "https://example.com/image.jpg",
					},
					Type: "image",
				},
				models.MultipartMessageTextContent{
					Text: "Message 2",
					Type: "text",
				},
			},
		},
		{
			name: "ListMinimal",

			content: `[
				{"text":"Message 1"},
				{"image_url":{"url":"https://example.com/image.jpg"}},
				{"text":"Message 2"}
			]`,

			expect: models.MultipartMessageList{
				models.MultipartMessageTextContent{
					Text: "Message 1",
				},
				models.MultipartMessageImageContent{
					ImageURL: models.MessageImageURL{
						URL: "https://example.com/image.jpg",
					},
				},
				models.MultipartMessageTextContent{
					Text: "Message 2",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var content models.MultipartMessageAny
			err := json.Unmarshal([]byte(testCase.content), &content)
			require.NoError(t, err)

			require.Equal(t, testCase.expect, content.MultipartMessage())
		})
	}
}
