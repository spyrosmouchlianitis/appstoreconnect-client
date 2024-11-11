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

// CapabilityType the model 'CapabilityType'
type CapabilityType string

// List of CapabilityType
const (
	ICLOUD CapabilityType = "ICLOUD"
	IN_APP_PURCHASE CapabilityType = "IN_APP_PURCHASE"
	GAME_CENTER CapabilityType = "GAME_CENTER"
	PUSH_NOTIFICATIONS CapabilityType = "PUSH_NOTIFICATIONS"
	WALLET CapabilityType = "WALLET"
	INTER_APP_AUDIO CapabilityType = "INTER_APP_AUDIO"
	MAPS CapabilityType = "MAPS"
	ASSOCIATED_DOMAINS CapabilityType = "ASSOCIATED_DOMAINS"
	PERSONAL_VPN CapabilityType = "PERSONAL_VPN"
	APP_GROUPS CapabilityType = "APP_GROUPS"
	HEALTHKIT CapabilityType = "HEALTHKIT"
	HOMEKIT CapabilityType = "HOMEKIT"
	WIRELESS_ACCESSORY_CONFIGURATION CapabilityType = "WIRELESS_ACCESSORY_CONFIGURATION"
	APPLE_PAY CapabilityType = "APPLE_PAY"
	DATA_PROTECTION CapabilityType = "DATA_PROTECTION"
	SIRIKIT CapabilityType = "SIRIKIT"
	NETWORK_EXTENSIONS CapabilityType = "NETWORK_EXTENSIONS"
	MULTIPATH CapabilityType = "MULTIPATH"
	HOT_SPOT CapabilityType = "HOT_SPOT"
	NFC_TAG_READING CapabilityType = "NFC_TAG_READING"
	CLASSKIT CapabilityType = "CLASSKIT"
	AUTOFILL_CREDENTIAL_PROVIDER CapabilityType = "AUTOFILL_CREDENTIAL_PROVIDER"
	ACCESS_WIFI_INFORMATION CapabilityType = "ACCESS_WIFI_INFORMATION"
	NETWORK_CUSTOM_PROTOCOL CapabilityType = "NETWORK_CUSTOM_PROTOCOL"
	COREMEDIA_HLS_LOW_LATENCY CapabilityType = "COREMEDIA_HLS_LOW_LATENCY"
	SYSTEM_EXTENSION_INSTALL CapabilityType = "SYSTEM_EXTENSION_INSTALL"
	USER_MANAGEMENT CapabilityType = "USER_MANAGEMENT"
	APPLE_ID_AUTH CapabilityType = "APPLE_ID_AUTH"
)

// All allowed values of CapabilityType enum
var AllowedCapabilityTypeEnumValues = []CapabilityType{
	"ICLOUD",
	"IN_APP_PURCHASE",
	"GAME_CENTER",
	"PUSH_NOTIFICATIONS",
	"WALLET",
	"INTER_APP_AUDIO",
	"MAPS",
	"ASSOCIATED_DOMAINS",
	"PERSONAL_VPN",
	"APP_GROUPS",
	"HEALTHKIT",
	"HOMEKIT",
	"WIRELESS_ACCESSORY_CONFIGURATION",
	"APPLE_PAY",
	"DATA_PROTECTION",
	"SIRIKIT",
	"NETWORK_EXTENSIONS",
	"MULTIPATH",
	"HOT_SPOT",
	"NFC_TAG_READING",
	"CLASSKIT",
	"AUTOFILL_CREDENTIAL_PROVIDER",
	"ACCESS_WIFI_INFORMATION",
	"NETWORK_CUSTOM_PROTOCOL",
	"COREMEDIA_HLS_LOW_LATENCY",
	"SYSTEM_EXTENSION_INSTALL",
	"USER_MANAGEMENT",
	"APPLE_ID_AUTH",
}

func (v *CapabilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CapabilityType(value)
	for _, existing := range AllowedCapabilityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CapabilityType", value)
}

// NewCapabilityTypeFromValue returns a pointer to a valid CapabilityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCapabilityTypeFromValue(v string) (*CapabilityType, error) {
	ev := CapabilityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CapabilityType: valid values are %v", v, AllowedCapabilityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CapabilityType) IsValid() bool {
	for _, existing := range AllowedCapabilityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CapabilityType value
func (v CapabilityType) Ptr() *CapabilityType {
	return &v
}

type NullableCapabilityType struct {
	value *CapabilityType
	isSet bool
}

func (v NullableCapabilityType) Get() *CapabilityType {
	return v.value
}

func (v *NullableCapabilityType) Set(val *CapabilityType) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityType) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityType(val *CapabilityType) *NullableCapabilityType {
	return &NullableCapabilityType{value: val, isSet: true}
}

func (v NullableCapabilityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

