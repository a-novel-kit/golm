package models

import (
	"encoding/json"
	"fmt"
)

type ResponseFormat string

func (responseFormat ResponseFormat) MarshalJSON() ([]byte, error) {
	return []byte(`{"type":"` + string(responseFormat) + `"}`), nil
}

func (responseFormat *ResponseFormat) UnmarshalJSON(data []byte) error {
	var responseFormatData struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &responseFormatData); err != nil {
		return fmt.Errorf("ResponseFormat.UnmarshalJSON: %w", err)
	}

	*responseFormat = ResponseFormat(responseFormatData.Type)

	return nil
}

const (
	ResponseFormatText ResponseFormat = "text"
	ResponseFormatJSON ResponseFormat = "json_object"
)
