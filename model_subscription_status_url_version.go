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

// SubscriptionStatusUrlVersion the model 'SubscriptionStatusUrlVersion'
type SubscriptionStatusUrlVersion string

// List of SubscriptionStatusUrlVersion
const (
	V1 SubscriptionStatusUrlVersion = "V1"
	V2 SubscriptionStatusUrlVersion = "V2"
)

// All allowed values of SubscriptionStatusUrlVersion enum
var AllowedSubscriptionStatusUrlVersionEnumValues = []SubscriptionStatusUrlVersion{
	"V1",
	"V2",
}

func (v *SubscriptionStatusUrlVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionStatusUrlVersion(value)
	for _, existing := range AllowedSubscriptionStatusUrlVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionStatusUrlVersion", value)
}

// NewSubscriptionStatusUrlVersionFromValue returns a pointer to a valid SubscriptionStatusUrlVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionStatusUrlVersionFromValue(v string) (*SubscriptionStatusUrlVersion, error) {
	ev := SubscriptionStatusUrlVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionStatusUrlVersion: valid values are %v", v, AllowedSubscriptionStatusUrlVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionStatusUrlVersion) IsValid() bool {
	for _, existing := range AllowedSubscriptionStatusUrlVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionStatusUrlVersion value
func (v SubscriptionStatusUrlVersion) Ptr() *SubscriptionStatusUrlVersion {
	return &v
}

type NullableSubscriptionStatusUrlVersion struct {
	value *SubscriptionStatusUrlVersion
	isSet bool
}

func (v NullableSubscriptionStatusUrlVersion) Get() *SubscriptionStatusUrlVersion {
	return v.value
}

func (v *NullableSubscriptionStatusUrlVersion) Set(val *SubscriptionStatusUrlVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionStatusUrlVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionStatusUrlVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionStatusUrlVersion(val *SubscriptionStatusUrlVersion) *NullableSubscriptionStatusUrlVersion {
	return &NullableSubscriptionStatusUrlVersion{value: val, isSet: true}
}

func (v NullableSubscriptionStatusUrlVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionStatusUrlVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

