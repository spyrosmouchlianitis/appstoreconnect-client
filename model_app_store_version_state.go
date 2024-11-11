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

// AppStoreVersionState the model 'AppStoreVersionState'
type AppStoreVersionState string

// List of AppStoreVersionState
const (
	ACCEPTED AppStoreVersionState = "ACCEPTED"
	DEVELOPER_REMOVED_FROM_SALE AppStoreVersionState = "DEVELOPER_REMOVED_FROM_SALE"
	DEVELOPER_REJECTED AppStoreVersionState = "DEVELOPER_REJECTED"
	IN_REVIEW AppStoreVersionState = "IN_REVIEW"
	INVALID_BINARY AppStoreVersionState = "INVALID_BINARY"
	METADATA_REJECTED AppStoreVersionState = "METADATA_REJECTED"
	PENDING_APPLE_RELEASE AppStoreVersionState = "PENDING_APPLE_RELEASE"
	PENDING_CONTRACT AppStoreVersionState = "PENDING_CONTRACT"
	PENDING_DEVELOPER_RELEASE AppStoreVersionState = "PENDING_DEVELOPER_RELEASE"
	PREPARE_FOR_SUBMISSION AppStoreVersionState = "PREPARE_FOR_SUBMISSION"
	PREORDER_READY_FOR_SALE AppStoreVersionState = "PREORDER_READY_FOR_SALE"
	PROCESSING_FOR_APP_STORE AppStoreVersionState = "PROCESSING_FOR_APP_STORE"
	READY_FOR_REVIEW AppStoreVersionState = "READY_FOR_REVIEW"
	READY_FOR_SALE AppStoreVersionState = "READY_FOR_SALE"
	REJECTED AppStoreVersionState = "REJECTED"
	REMOVED_FROM_SALE AppStoreVersionState = "REMOVED_FROM_SALE"
	WAITING_FOR_EXPORT_COMPLIANCE AppStoreVersionState = "WAITING_FOR_EXPORT_COMPLIANCE"
	WAITING_FOR_REVIEW AppStoreVersionState = "WAITING_FOR_REVIEW"
	REPLACED_WITH_NEW_VERSION AppStoreVersionState = "REPLACED_WITH_NEW_VERSION"
	NOT_APPLICABLE AppStoreVersionState = "NOT_APPLICABLE"
)

// All allowed values of AppStoreVersionState enum
var AllowedAppStoreVersionStateEnumValues = []AppStoreVersionState{
	"ACCEPTED",
	"DEVELOPER_REMOVED_FROM_SALE",
	"DEVELOPER_REJECTED",
	"IN_REVIEW",
	"INVALID_BINARY",
	"METADATA_REJECTED",
	"PENDING_APPLE_RELEASE",
	"PENDING_CONTRACT",
	"PENDING_DEVELOPER_RELEASE",
	"PREPARE_FOR_SUBMISSION",
	"PREORDER_READY_FOR_SALE",
	"PROCESSING_FOR_APP_STORE",
	"READY_FOR_REVIEW",
	"READY_FOR_SALE",
	"REJECTED",
	"REMOVED_FROM_SALE",
	"WAITING_FOR_EXPORT_COMPLIANCE",
	"WAITING_FOR_REVIEW",
	"REPLACED_WITH_NEW_VERSION",
	"NOT_APPLICABLE",
}

func (v *AppStoreVersionState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppStoreVersionState(value)
	for _, existing := range AllowedAppStoreVersionStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppStoreVersionState", value)
}

// NewAppStoreVersionStateFromValue returns a pointer to a valid AppStoreVersionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppStoreVersionStateFromValue(v string) (*AppStoreVersionState, error) {
	ev := AppStoreVersionState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppStoreVersionState: valid values are %v", v, AllowedAppStoreVersionStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppStoreVersionState) IsValid() bool {
	for _, existing := range AllowedAppStoreVersionStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppStoreVersionState value
func (v AppStoreVersionState) Ptr() *AppStoreVersionState {
	return &v
}

type NullableAppStoreVersionState struct {
	value *AppStoreVersionState
	isSet bool
}

func (v NullableAppStoreVersionState) Get() *AppStoreVersionState {
	return v.value
}

func (v *NullableAppStoreVersionState) Set(val *AppStoreVersionState) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionState) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionState(val *AppStoreVersionState) *NullableAppStoreVersionState {
	return &NullableAppStoreVersionState{value: val, isSet: true}
}

func (v NullableAppStoreVersionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

