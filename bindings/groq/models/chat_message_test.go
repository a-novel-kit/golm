package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func TestMarshalChatCompletionMessages(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		message models.Message

		expect string
	}{
		{
			name: "SystemMessage",

			message: models.SystemMessage{
				Name:    "Test Message",
				Content: "Test Content",
			},

			expect: `{
				"role":"system",
				"content":"Test Content",
				"name":"Test Message"
			}`,
		},
		{
			name: "SystemMessageMinimal",

			message: models.SystemMessage{
				Content: "Test Content",
			},

			expect: `{
				"role":"system",
				"content":"Test Content"
			}`,
		},
		{
			name: "UserMessage",

			message: models.UserMessage{
				Name:    "Test Message",
				Content: models.NewMultipartStaticMessage("Test Content"),
			},

			expect: `{
				"role":"user",
				"content":"Test Content",
				"name":"Test Message"
			}`,
		},
		{
			name: "UserMessageMinimal",

			message: models.UserMessage{
				Content: models.NewMultipartStaticMessage("Test Content"),
			},

			expect: `{
				"role":"user",
				"content":"Test Content"
			}`,
		},
		{
			name: "AssistantMessage",

			message: models.AssistantMessage{
				Name:    "Test Message",
				Content: "Test Content",
				ToolCalls: []models.ToolCall{
					{
						ID:   "test-id",
						Type: models.ToolTypeFunction,
						Function: &models.ToolCallFunction{
							Arguments: `{"test":"test"}`,
							Name:      "test-function",
						},
					},
				},
			},

			expect: `{
				"role":"assistant",
				"content":"Test Content",
				"name":"Test Message",
				"tool_calls":[{
					"id":"test-id",
					"type":"function",
					"function":{
						"arguments":"{\"test\":\"test\"}",
						"name":"test-function"
					}
				}]
			}`,
		},
		{
			name: "AssistantMessageMinimalWithToolCalls",

			message: models.AssistantMessage{
				ToolCalls: []models.ToolCall{
					{
						ID: "test-id",
					},
				},
			},

			expect: `{
				"role":"assistant",
				"tool_calls":[{
					"id":"test-id"
				}]
			}`,
		},
		{
			name: "AssistantMessageMinimal",

			message: models.AssistantMessage{
				Content: "Test Content",
			},

			expect: `{
				"role":"assistant",
				"content":"Test Content"
			}`,
		},
		{
			name: "ToolCallMessage",

			message: models.ToolCallMessage{
				Content:    "Test Content",
				ToolCallID: "test-id",
			},

			expect: `{
				"role":"tool",
				"content":"Test Content",
				"tool_call_id":"test-id"
			}`,
		},
		{
			name: "ToolCallMessageMinimal",

			message: models.ToolCallMessage{
				Content: "Test Content",
			},

			expect: `{
				"role":"tool",
				"content":"Test Content"
			}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			mrsh, err := json.Marshal(testCase.message)
			require.NoError(t, err)

			require.JSONEq(t, testCase.expect, string(mrsh))
		})
	}
}

func TestUnmarshalChatCompletionMessages(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		message string

		expect models.Message
	}{
		{
			name: "SystemMessage",

			message: `{
				"role":"system",
				"content":"Test Content",
				"name":"Test Message"
			}`,

			expect: models.SystemMessage{
				Name:    "Test Message",
				Content: "Test Content",
			},
		},
		{
			name: "SystemMessageMinimal",

			message: `{
				"role":"system",
				"content":"Test Content"
			}`,

			expect: models.SystemMessage{
				Content: "Test Content",
			},
		},
		{
			name: "UserMessage",

			message: `{
				"role":"user",
				"content":"Test Content",
				"name":"Test Message"
			}`,

			expect: models.UserMessage{
				Name:    "Test Message",
				Content: models.NewMultipartStaticMessage("Test Content"),
			},
		},
		{
			name: "UserMessageMinimal",

			message: `{
				"role":"user",
				"content":"Test Content"
			}`,

			expect: models.UserMessage{
				Content: models.NewMultipartStaticMessage("Test Content"),
			},
		},
		{
			name: "AssistantMessage",

			message: `{
				"role":"assistant",
				"content":"Test Content",
				"name":"Test Message",
				"tool_calls":[{
					"id":"test-id",
					"type":"function",
					"function":{
						"arguments":"{\"test\":\"test\"}",
						"name":"test-function"
					}
				}]
			}`,

			expect: models.AssistantMessage{
				Name:    "Test Message",
				Content: "Test Content",
				ToolCalls: []models.ToolCall{
					{
						ID:   "test-id",
						Type: models.ToolTypeFunction,
						Function: &models.ToolCallFunction{
							Arguments: `{"test":"test"}`,
							Name:      "test-function",
						},
					},
				},
			},
		},
		{
			name: "AssistantMessageMinimalWithToolCalls",

			message: `{
				"role":"assistant",
				"tool_calls":[{
					"id":"test-id"
				}]
			}`,

			expect: models.AssistantMessage{
				ToolCalls: []models.ToolCall{
					{
						ID: "test-id",
					},
				},
			},
		},
		{
			name: "AssistantMessageMinimal",

			message: `{
				"role":"assistant",
				"content":"Test Content"
			}`,

			expect: models.AssistantMessage{
				Content: "Test Content",
			},
		},
		{
			name: "ToolCallMessage",

			message: `{
				"role":"tool",
				"content":"Test Content",
				"tool_call_id":"test-id"
			}`,

			expect: models.ToolCallMessage{
				Content:    "Test Content",
				ToolCallID: "test-id",
			},
		},
		{
			name: "ToolCallMessageMinimal",

			message: `{
				"role":"tool",
				"content":"Test Content"
			}`,

			expect: models.ToolCallMessage{
				Content: "Test Content",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			message, err := models.UnmarshalMessage([]byte(testCase.message))
			require.NoError(t, err)

			require.Equal(t, testCase.expect, message.Message())
		})
	}
}
