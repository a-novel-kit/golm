package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func TestMarshalToolChoice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		data models.ToolChoice

		expect string
	}{
		{
			name: "Static/None",

			data: *models.NewToolChoiceStatic(models.ToolChoiceStaticNone),

			expect: `"none"`,
		},
		{
			name: "Static/Auto",

			data: *models.NewToolChoiceStatic(models.ToolChoiceStaticAuto),

			expect: `"auto"`,
		},
		{
			name: "Object",

			data: *models.NewToolChoiceObject(models.ToolChoiceObject{
				Type: models.ToolTypeFunction,
				Function: &models.ToolChoiceFunction{
					Name: "test",
				},
			}),

			expect: `{"type":"function","function":{"name":"test"}}`,
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

func TestUnmarshalToolChoice(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		data string

		expect models.ToolChoice
	}{
		{
			name: "Empty",

			data: `null`,

			expect: models.ToolChoice{},
		},
		{
			name: "Static/None",

			data: `"none"`,

			expect: *models.NewToolChoiceStatic(models.ToolChoiceStaticNone),
		},
		{
			name: "Static/Auto",

			data: `"auto"`,

			expect: *models.NewToolChoiceStatic(models.ToolChoiceStaticAuto),
		},
		{
			name: "Object",

			data: `{"type":"function","function":{"name":"test"}}`,

			expect: *models.NewToolChoiceObject(models.ToolChoiceObject{
				Type: models.ToolTypeFunction,
				Function: &models.ToolChoiceFunction{
					Name: "test",
				},
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var out models.ToolChoice
			err := json.Unmarshal([]byte(testCase.data), &out)
			require.NoError(t, err)

			require.Equal(t, testCase.expect, out)
		})
	}
}
