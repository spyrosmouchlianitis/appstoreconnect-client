/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"fmt"
)

// AppClipAction the model 'AppClipAction'
type AppClipAction string

// List of AppClipAction
const (
	OPEN AppClipAction = "OPEN"
	VIEW AppClipAction = "VIEW"
	PLAY AppClipAction = "PLAY"
)

// All allowed values of AppClipAction enum
var AllowedAppClipActionEnumValues = []AppClipAction{
	"OPEN",
	"VIEW",
	"PLAY",
}

func (v *AppClipAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppClipAction(value)
	for _, existing := range AllowedAppClipActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppClipAction", value)
}

// NewAppClipActionFromValue returns a pointer to a valid AppClipAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppClipActionFromValue(v string) (*AppClipAction, error) {
	ev := AppClipAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppClipAction: valid values are %v", v, AllowedAppClipActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppClipAction) IsValid() bool {
	for _, existing := range AllowedAppClipActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppClipAction value
func (v AppClipAction) Ptr() *AppClipAction {
	return &v
}

type NullableAppClipAction struct {
	value *AppClipAction
	isSet bool
}

func (v NullableAppClipAction) Get() *AppClipAction {
	return v.value
}

func (v *NullableAppClipAction) Set(val *AppClipAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAction(val *AppClipAction) *NullableAppClipAction {
	return &NullableAppClipAction{value: val, isSet: true}
}

func (v NullableAppClipAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

