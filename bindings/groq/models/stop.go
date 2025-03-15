package models

import "encoding/json"

type Stop [4]string

func (stop Stop) MarshalJSON() ([]byte, error) {
	switch {
	case stop[3] != "":
		return json.Marshal(stop[:4])
	case stop[2] != "":
		return json.Marshal(stop[:3])
	case stop[1] != "":
		return json.Marshal(stop[:2])
	case stop[0] != "":
		return json.Marshal(stop[0])
	default:
		return []byte("null"), nil
	}
}

func (stop *Stop) UnmarshalJSON(data []byte) error {
	var outSlice [4]string
	err := json.Unmarshal(data, &outSlice)

	if err == nil {
		*stop = outSlice

		return nil
	}

	var outString string
	err = json.Unmarshal(data, &outString)

	if err == nil {
		*stop = [4]string{outString}

		return nil
	}

	return err
}
