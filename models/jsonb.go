package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB is a custom data type for handling JSONB fields in PostgreSQL
type JSONBRelationship []APTFeedRelationshipObject
type JSONBData []APTFeedDataObject
type JSONBTechnic []APTFeedTechnicObject
type JSONBTactic []APTFeedTacticObject
type JSONBMitigation []APTFeedMitigationObject
type JSONBExternalReferences []ExternalReferences
type JSONBKillChainPhases []KillChainPhases

// Value converts the JSONB object to a JSON-encoded byte slice for storage in the database
func (j JSONBRelationship) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan converts a JSON-encoded byte slice from the database to a JSONB object
func (j *JSONBRelationship) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

// Implement Value and Scan methods for each type
func (j JSONBData) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONBTechnic) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBTechnic) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONBTactic) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBTactic) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONBMitigation) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBMitigation) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

func (e *JSONBExternalReferences) Scan(value interface{}) error {
	if value == nil {
		*e = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, e)
}

func (e JSONBExternalReferences) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}

	return json.Marshal(e)
}

func (j JSONBKillChainPhases) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONBKillChainPhases) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}