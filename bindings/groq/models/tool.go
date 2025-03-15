package models

type Tool struct {
	// The type of the tool. Currently, only ToolTypeFunction is supported.
	Type ToolType `json:"type,omitempty"`
}

type ToolFunction struct {
	// A description of what the function does, used by the model to choose when and how to call the function.
	Description string `json:"description,omitempty"`
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	// length of 64.
	Name string `json:"name"`
	// The parameters the functions accepts, described as a JSON Schema object. See the docs on tool use for examples,
	// and the JSON Schema reference for documentation about the format.
	//
	// Omitting parameters defines a function with an empty parameter list.
	//
	// https://console.api.com/docs/tool-use
	// https://json-schema.org/understanding-json-schema/reference
	Parameters map[string]any `json:"parameters,omitempty"`
}
