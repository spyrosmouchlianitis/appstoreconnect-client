/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// GameCenterLeaderboardSetReleasesResponseIncludedInner - struct for GameCenterLeaderboardSetReleasesResponseIncludedInner
type GameCenterLeaderboardSetReleasesResponseIncludedInner struct {
	GameCenterDetail *GameCenterDetail
	GameCenterLeaderboardSet *GameCenterLeaderboardSet
}

// GameCenterDetailAsGameCenterLeaderboardSetReleasesResponseIncludedInner is a convenience function that returns GameCenterDetail wrapped in GameCenterLeaderboardSetReleasesResponseIncludedInner
func GameCenterDetailAsGameCenterLeaderboardSetReleasesResponseIncludedInner(v *GameCenterDetail) GameCenterLeaderboardSetReleasesResponseIncludedInner {
	return GameCenterLeaderboardSetReleasesResponseIncludedInner{
		GameCenterDetail: v,
	}
}

// GameCenterLeaderboardSetAsGameCenterLeaderboardSetReleasesResponseIncludedInner is a convenience function that returns GameCenterLeaderboardSet wrapped in GameCenterLeaderboardSetReleasesResponseIncludedInner
func GameCenterLeaderboardSetAsGameCenterLeaderboardSetReleasesResponseIncludedInner(v *GameCenterLeaderboardSet) GameCenterLeaderboardSetReleasesResponseIncludedInner {
	return GameCenterLeaderboardSetReleasesResponseIncludedInner{
		GameCenterLeaderboardSet: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterLeaderboardSetReleasesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GameCenterDetail
	err = newStrictDecoder(data).Decode(&dst.GameCenterDetail)
	if err == nil {
		jsonGameCenterDetail, _ := json.Marshal(dst.GameCenterDetail)
		if string(jsonGameCenterDetail) == "{}" { // empty struct
			dst.GameCenterDetail = nil
		} else {
			if err = validator.Validate(dst.GameCenterDetail); err != nil {
				dst.GameCenterDetail = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterDetail = nil
	}

	// try to unmarshal data into GameCenterLeaderboardSet
	err = newStrictDecoder(data).Decode(&dst.GameCenterLeaderboardSet)
	if err == nil {
		jsonGameCenterLeaderboardSet, _ := json.Marshal(dst.GameCenterLeaderboardSet)
		if string(jsonGameCenterLeaderboardSet) == "{}" { // empty struct
			dst.GameCenterLeaderboardSet = nil
		} else {
			if err = validator.Validate(dst.GameCenterLeaderboardSet); err != nil {
				dst.GameCenterLeaderboardSet = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterLeaderboardSet = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GameCenterDetail = nil
		dst.GameCenterLeaderboardSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterLeaderboardSetReleasesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterLeaderboardSetReleasesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterLeaderboardSetReleasesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.GameCenterDetail != nil {
		return json.Marshal(&src.GameCenterDetail)
	}

	if src.GameCenterLeaderboardSet != nil {
		return json.Marshal(&src.GameCenterLeaderboardSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterLeaderboardSetReleasesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GameCenterDetail != nil {
		return obj.GameCenterDetail
	}

	if obj.GameCenterLeaderboardSet != nil {
		return obj.GameCenterLeaderboardSet
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterLeaderboardSetReleasesResponseIncludedInner struct {
	value *GameCenterLeaderboardSetReleasesResponseIncludedInner
	isSet bool
}

func (v NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) Get() *GameCenterLeaderboardSetReleasesResponseIncludedInner {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) Set(val *GameCenterLeaderboardSetReleasesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetReleasesResponseIncludedInner(val *GameCenterLeaderboardSetReleasesResponseIncludedInner) *NullableGameCenterLeaderboardSetReleasesResponseIncludedInner {
	return &NullableGameCenterLeaderboardSetReleasesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetReleasesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


