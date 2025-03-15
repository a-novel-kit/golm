package models

import "encoding/json"

type ToolChoiceObject struct {
	// The type of the tool. Currently, only ToolTypeFunction is supported.
	Type     ToolType            `json:"type,omitempty"`
	Function *ToolChoiceFunction `json:"function,omitempty"`
}

type ToolChoiceFunction struct {
	// The name of the function to call.
	Name string `json:"name,omitempty"`
}

type ToolChoiceStatic string

const (
	ToolChoiceStaticNone ToolChoiceStatic = "none"
	ToolChoiceStaticAuto ToolChoiceStatic = "auto"
)

type ToolChoice struct {
	static ToolChoiceStatic
	object ToolChoiceObject
}

func (choice ToolChoice) MarshalJSON() ([]byte, error) {
	switch {
	case choice.static != "":
		return json.Marshal(choice.static)
	default:
		return json.Marshal(choice.object)
	}
}

func (choice *ToolChoice) UnmarshalJSON(data []byte) error {
	var outStatic ToolChoiceStatic
	err := json.Unmarshal(data, &outStatic)

	if err == nil {
		choice.static = outStatic

		return nil
	}

	var outObject ToolChoiceObject
	err = json.Unmarshal(data, &outObject)

	if err == nil {
		choice.object = outObject

		return nil
	}

	return err
}

func NewToolChoiceStatic(static ToolChoiceStatic) *ToolChoice {
	return &ToolChoice{static: static}
}

func NewToolChoiceObject(object ToolChoiceObject) *ToolChoice {
	return &ToolChoice{object: object}
}
