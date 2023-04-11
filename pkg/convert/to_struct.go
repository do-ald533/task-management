package convert

import (
	"encoding/json"
)

func ToStruct(encode any, decode any) error {
	data, err := json.Marshal(encode)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, decode); err != nil {
		return err
	}

	return nil
}
