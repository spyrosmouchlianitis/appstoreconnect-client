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

// AppVersionState the model 'AppVersionState'
type AppVersionState string

// List of AppVersionState
const (
	ACCEPTED AppVersionState = "ACCEPTED"
	DEVELOPER_REJECTED AppVersionState = "DEVELOPER_REJECTED"
	IN_REVIEW AppVersionState = "IN_REVIEW"
	INVALID_BINARY AppVersionState = "INVALID_BINARY"
	METADATA_REJECTED AppVersionState = "METADATA_REJECTED"
	PENDING_APPLE_RELEASE AppVersionState = "PENDING_APPLE_RELEASE"
	PENDING_DEVELOPER_RELEASE AppVersionState = "PENDING_DEVELOPER_RELEASE"
	PREPARE_FOR_SUBMISSION AppVersionState = "PREPARE_FOR_SUBMISSION"
	PROCESSING_FOR_DISTRIBUTION AppVersionState = "PROCESSING_FOR_DISTRIBUTION"
	READY_FOR_DISTRIBUTION AppVersionState = "READY_FOR_DISTRIBUTION"
	READY_FOR_REVIEW AppVersionState = "READY_FOR_REVIEW"
	REJECTED AppVersionState = "REJECTED"
	REPLACED_WITH_NEW_VERSION AppVersionState = "REPLACED_WITH_NEW_VERSION"
	WAITING_FOR_EXPORT_COMPLIANCE AppVersionState = "WAITING_FOR_EXPORT_COMPLIANCE"
	WAITING_FOR_REVIEW AppVersionState = "WAITING_FOR_REVIEW"
)

// All allowed values of AppVersionState enum
var AllowedAppVersionStateEnumValues = []AppVersionState{
	"ACCEPTED",
	"DEVELOPER_REJECTED",
	"IN_REVIEW",
	"INVALID_BINARY",
	"METADATA_REJECTED",
	"PENDING_APPLE_RELEASE",
	"PENDING_DEVELOPER_RELEASE",
	"PREPARE_FOR_SUBMISSION",
	"PROCESSING_FOR_DISTRIBUTION",
	"READY_FOR_DISTRIBUTION",
	"READY_FOR_REVIEW",
	"REJECTED",
	"REPLACED_WITH_NEW_VERSION",
	"WAITING_FOR_EXPORT_COMPLIANCE",
	"WAITING_FOR_REVIEW",
}

func (v *AppVersionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppVersionState(value)
	for _, existing := range AllowedAppVersionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppVersionState", value)
}

// NewAppVersionStateFromValue returns a pointer to a valid AppVersionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppVersionStateFromValue(v string) (*AppVersionState, error) {
	ev := AppVersionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppVersionState: valid values are %v", v, AllowedAppVersionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppVersionState) IsValid() bool {
	for _, existing := range AllowedAppVersionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppVersionState value
func (v AppVersionState) Ptr() *AppVersionState {
	return &v
}

type NullableAppVersionState struct {
	value *AppVersionState
	isSet bool
}

func (v NullableAppVersionState) Get() *AppVersionState {
	return v.value
}

func (v *NullableAppVersionState) Set(val *AppVersionState) {
	v.value = val
	v.isSet = true
}

func (v NullableAppVersionState) IsSet() bool {
	return v.isSet
}

func (v *NullableAppVersionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppVersionState(val *AppVersionState) *NullableAppVersionState {
	return &NullableAppVersionState{value: val, isSet: true}
}

func (v NullableAppVersionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppVersionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

