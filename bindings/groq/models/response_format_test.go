package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func TestMarshalResponseFormat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		responseFormat models.ResponseFormat

		expect string
	}{
		{
			name: "Text",

			responseFormat: models.ResponseFormatText,

			expect: `{"type":"text"}`,
		},
		{
			name: "JSON",

			responseFormat: models.ResponseFormatJSON,

			expect: `{"type":"json_object"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			mrsh, err := json.Marshal(testCase.responseFormat)
			require.NoError(t, err)

			require.JSONEq(t, testCase.expect, string(mrsh))
		})
	}
}

func TestUnmarshalResponseFormat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		responseFormat string

		expect models.ResponseFormat
	}{
		{
			name: "Text",

			responseFormat: `{"type":"text"}`,

			expect: models.ResponseFormatText,
		},
		{
			name: "JSON",

			responseFormat: `{"type":"json_object"}`,

			expect: models.ResponseFormatJSON,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var responseFormat models.ResponseFormat

			require.NoError(t, json.Unmarshal([]byte(testCase.responseFormat), &responseFormat))

			require.Equal(t, testCase.expect, responseFormat)
		})
	}
}
