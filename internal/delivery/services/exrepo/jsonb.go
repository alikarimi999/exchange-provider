package exrepo

import (
	"database/sql/driver"
	"encoding/json"
)

type jsonb map[string]interface{}

func (j jsonb) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *jsonb) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
