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

// GameCenterLeaderboardLocalizationsResponseIncludedInner - struct for GameCenterLeaderboardLocalizationsResponseIncludedInner
type GameCenterLeaderboardLocalizationsResponseIncludedInner struct {
	GameCenterLeaderboard *GameCenterLeaderboard
	GameCenterLeaderboardImage *GameCenterLeaderboardImage
}

// GameCenterLeaderboardAsGameCenterLeaderboardLocalizationsResponseIncludedInner is a convenience function that returns GameCenterLeaderboard wrapped in GameCenterLeaderboardLocalizationsResponseIncludedInner
func GameCenterLeaderboardAsGameCenterLeaderboardLocalizationsResponseIncludedInner(v *GameCenterLeaderboard) GameCenterLeaderboardLocalizationsResponseIncludedInner {
	return GameCenterLeaderboardLocalizationsResponseIncludedInner{
		GameCenterLeaderboard: v,
	}
}

// GameCenterLeaderboardImageAsGameCenterLeaderboardLocalizationsResponseIncludedInner is a convenience function that returns GameCenterLeaderboardImage wrapped in GameCenterLeaderboardLocalizationsResponseIncludedInner
func GameCenterLeaderboardImageAsGameCenterLeaderboardLocalizationsResponseIncludedInner(v *GameCenterLeaderboardImage) GameCenterLeaderboardLocalizationsResponseIncludedInner {
	return GameCenterLeaderboardLocalizationsResponseIncludedInner{
		GameCenterLeaderboardImage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterLeaderboardLocalizationsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GameCenterLeaderboard
	err = newStrictDecoder(data).Decode(&dst.GameCenterLeaderboard)
	if err == nil {
		jsonGameCenterLeaderboard, _ := json.Marshal(dst.GameCenterLeaderboard)
		if string(jsonGameCenterLeaderboard) == "{}" { // empty struct
			dst.GameCenterLeaderboard = nil
		} else {
			if err = validator.Validate(dst.GameCenterLeaderboard); err != nil {
				dst.GameCenterLeaderboard = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterLeaderboard = nil
	}

	// try to unmarshal data into GameCenterLeaderboardImage
	err = newStrictDecoder(data).Decode(&dst.GameCenterLeaderboardImage)
	if err == nil {
		jsonGameCenterLeaderboardImage, _ := json.Marshal(dst.GameCenterLeaderboardImage)
		if string(jsonGameCenterLeaderboardImage) == "{}" { // empty struct
			dst.GameCenterLeaderboardImage = nil
		} else {
			if err = validator.Validate(dst.GameCenterLeaderboardImage); err != nil {
				dst.GameCenterLeaderboardImage = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterLeaderboardImage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GameCenterLeaderboard = nil
		dst.GameCenterLeaderboardImage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterLeaderboardLocalizationsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterLeaderboardLocalizationsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterLeaderboardLocalizationsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.GameCenterLeaderboard != nil {
		return json.Marshal(&src.GameCenterLeaderboard)
	}

	if src.GameCenterLeaderboardImage != nil {
		return json.Marshal(&src.GameCenterLeaderboardImage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterLeaderboardLocalizationsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GameCenterLeaderboard != nil {
		return obj.GameCenterLeaderboard
	}

	if obj.GameCenterLeaderboardImage != nil {
		return obj.GameCenterLeaderboardImage
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterLeaderboardLocalizationsResponseIncludedInner struct {
	value *GameCenterLeaderboardLocalizationsResponseIncludedInner
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) Get() *GameCenterLeaderboardLocalizationsResponseIncludedInner {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) Set(val *GameCenterLeaderboardLocalizationsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationsResponseIncludedInner(val *GameCenterLeaderboardLocalizationsResponseIncludedInner) *NullableGameCenterLeaderboardLocalizationsResponseIncludedInner {
	return &NullableGameCenterLeaderboardLocalizationsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


