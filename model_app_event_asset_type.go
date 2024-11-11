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

// AppEventAssetType the model 'AppEventAssetType'
type AppEventAssetType string

// List of AppEventAssetType
const (
	CARD AppEventAssetType = "EVENT_CARD"
	DETAILS_PAGE AppEventAssetType = "EVENT_DETAILS_PAGE"
)

// All allowed values of AppEventAssetType enum
var AllowedAppEventAssetTypeEnumValues = []AppEventAssetType{
	"EVENT_CARD",
	"EVENT_DETAILS_PAGE",
}

func (v *AppEventAssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppEventAssetType(value)
	for _, existing := range AllowedAppEventAssetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppEventAssetType", value)
}

// NewAppEventAssetTypeFromValue returns a pointer to a valid AppEventAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppEventAssetTypeFromValue(v string) (*AppEventAssetType, error) {
	ev := AppEventAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppEventAssetType: valid values are %v", v, AllowedAppEventAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppEventAssetType) IsValid() bool {
	for _, existing := range AllowedAppEventAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppEventAssetType value
func (v AppEventAssetType) Ptr() *AppEventAssetType {
	return &v
}

type NullableAppEventAssetType struct {
	value *AppEventAssetType
	isSet bool
}

func (v NullableAppEventAssetType) Get() *AppEventAssetType {
	return v.value
}

func (v *NullableAppEventAssetType) Set(val *AppEventAssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventAssetType(val *AppEventAssetType) *NullableAppEventAssetType {
	return &NullableAppEventAssetType{value: val, isSet: true}
}

func (v NullableAppEventAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

