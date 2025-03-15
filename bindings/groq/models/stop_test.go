package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func TestMarshalStop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		data models.Stop

		expect string
	}{
		{
			name: "Empty",

			data: models.Stop{},

			expect: `null`,
		},
		{
			name: "One",

			data: models.Stop{"one"},

			expect: `"one"`,
		},
		{
			name: "Two",

			data: models.Stop{"one", "two"},

			expect: `["one","two"]`,
		},
		{
			name: "Max",

			data: models.Stop{"one", "two", "three", "four"},

			expect: `["one","two","three","four"]`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			mrsh, err := json.Marshal(testCase.data)
			require.NoError(t, err)

			require.JSONEq(t, testCase.expect, string(mrsh))
		})
	}
}

func TestUnmarshalStop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		data string

		expect models.Stop
	}{
		{
			name: "One",

			data: `"one"`,

			expect: models.Stop{"one"},
		},
		{
			name: "OneSlice",

			data: `["one"]`,

			expect: models.Stop{"one"},
		},
		{
			name: "Two",

			data: `["one","two"]`,

			expect: models.Stop{"one", "two"},
		},
		{
			name: "Max",

			data: `["one","two","three","four"]`,

			expect: models.Stop{"one", "two", "three", "four"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var stop models.Stop

			require.NoError(t, json.Unmarshal([]byte(testCase.data), &stop))

			require.Equal(t, testCase.expect, stop)
		})
	}
}
