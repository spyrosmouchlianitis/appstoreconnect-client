/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GameCenterAppVersionsResponseIncludedInner - struct for GameCenterAppVersionsResponseIncludedInner
type GameCenterAppVersionsResponseIncludedInner struct {
	AppStoreVersion *AppStoreVersion
	GameCenterAppVersion *GameCenterAppVersion
}

// AppStoreVersionAsGameCenterAppVersionsResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in GameCenterAppVersionsResponseIncludedInner
func AppStoreVersionAsGameCenterAppVersionsResponseIncludedInner(v *AppStoreVersion) GameCenterAppVersionsResponseIncludedInner {
	return GameCenterAppVersionsResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// GameCenterAppVersionAsGameCenterAppVersionsResponseIncludedInner is a convenience function that returns GameCenterAppVersion wrapped in GameCenterAppVersionsResponseIncludedInner
func GameCenterAppVersionAsGameCenterAppVersionsResponseIncludedInner(v *GameCenterAppVersion) GameCenterAppVersionsResponseIncludedInner {
	return GameCenterAppVersionsResponseIncludedInner{
		GameCenterAppVersion: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterAppVersionsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			if err = validator.Validate(dst.AppStoreVersion); err != nil {
				dst.AppStoreVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.AppStoreVersion = nil
	}

	// try to unmarshal data into GameCenterAppVersion
	err = newStrictDecoder(data).Decode(&dst.GameCenterAppVersion)
	if err == nil {
		jsonGameCenterAppVersion, _ := json.Marshal(dst.GameCenterAppVersion)
		if string(jsonGameCenterAppVersion) == "{}" { // empty struct
			dst.GameCenterAppVersion = nil
		} else {
			if err = validator.Validate(dst.GameCenterAppVersion); err != nil {
				dst.GameCenterAppVersion = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterAppVersion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppStoreVersion = nil
		dst.GameCenterAppVersion = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterAppVersionsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterAppVersionsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterAppVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	if src.GameCenterAppVersion != nil {
		return json.Marshal(&src.GameCenterAppVersion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterAppVersionsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	if obj.GameCenterAppVersion != nil {
		return obj.GameCenterAppVersion
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterAppVersionsResponseIncludedInner struct {
	value *GameCenterAppVersionsResponseIncludedInner
	isSet bool
}

func (v NullableGameCenterAppVersionsResponseIncludedInner) Get() *GameCenterAppVersionsResponseIncludedInner {
	return v.value
}

func (v *NullableGameCenterAppVersionsResponseIncludedInner) Set(val *GameCenterAppVersionsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionsResponseIncludedInner(val *GameCenterAppVersionsResponseIncludedInner) *NullableGameCenterAppVersionsResponseIncludedInner {
	return &NullableGameCenterAppVersionsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


