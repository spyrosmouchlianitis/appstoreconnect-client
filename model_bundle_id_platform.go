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

// BundleIdPlatform the model 'BundleIdPlatform'
type BundleIdPlatform string

// List of BundleIdPlatform
const (
	IOS BundleIdPlatform = "IOS"
	MAC_OS BundleIdPlatform = "MAC_OS"
)

// All allowed values of BundleIdPlatform enum
var AllowedBundleIdPlatformEnumValues = []BundleIdPlatform{
	"IOS",
	"MAC_OS",
}

func (v *BundleIdPlatform) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BundleIdPlatform(value)
	for _, existing := range AllowedBundleIdPlatformEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BundleIdPlatform", value)
}

// NewBundleIdPlatformFromValue returns a pointer to a valid BundleIdPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBundleIdPlatformFromValue(v string) (*BundleIdPlatform, error) {
	ev := BundleIdPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BundleIdPlatform: valid values are %v", v, AllowedBundleIdPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BundleIdPlatform) IsValid() bool {
	for _, existing := range AllowedBundleIdPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BundleIdPlatform value
func (v BundleIdPlatform) Ptr() *BundleIdPlatform {
	return &v
}

type NullableBundleIdPlatform struct {
	value *BundleIdPlatform
	isSet bool
}

func (v NullableBundleIdPlatform) Get() *BundleIdPlatform {
	return v.value
}

func (v *NullableBundleIdPlatform) Set(val *BundleIdPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdPlatform(val *BundleIdPlatform) *NullableBundleIdPlatform {
	return &NullableBundleIdPlatform{value: val, isSet: true}
}

func (v NullableBundleIdPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

