/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"fmt"
)

// CiActionType the model 'CiActionType'
type CiActionType string

// List of CiActionType
const (
	BUILD CiActionType = "BUILD"
	ANALYZE CiActionType = "ANALYZE"
	TEST CiActionType = "TEST"
	ARCHIVE CiActionType = "ARCHIVE"
)

// All allowed values of CiActionType enum
var AllowedCiActionTypeEnumValues = []CiActionType{
	"BUILD",
	"ANALYZE",
	"TEST",
	"ARCHIVE",
}

func (v *CiActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CiActionType(value)
	for _, existing := range AllowedCiActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CiActionType", value)
}

// NewCiActionTypeFromValue returns a pointer to a valid CiActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCiActionTypeFromValue(v string) (*CiActionType, error) {
	ev := CiActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CiActionType: valid values are %v", v, AllowedCiActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CiActionType) IsValid() bool {
	for _, existing := range AllowedCiActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CiActionType value
func (v CiActionType) Ptr() *CiActionType {
	return &v
}

type NullableCiActionType struct {
	value *CiActionType
	isSet bool
}

func (v NullableCiActionType) Get() *CiActionType {
	return v.value
}

func (v *NullableCiActionType) Set(val *CiActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableCiActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableCiActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiActionType(val *CiActionType) *NullableCiActionType {
	return &NullableCiActionType{value: val, isSet: true}
}

func (v NullableCiActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

