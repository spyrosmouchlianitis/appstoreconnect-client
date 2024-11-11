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

// GameCenterLeaderboardSetsResponseIncludedInner - struct for GameCenterLeaderboardSetsResponseIncludedInner
type GameCenterLeaderboardSetsResponseIncludedInner struct {
	GameCenterDetail *GameCenterDetail
	GameCenterGroup *GameCenterGroup
	GameCenterLeaderboard *GameCenterLeaderboard
	GameCenterLeaderboardSet *GameCenterLeaderboardSet
	GameCenterLeaderboardSetLocalization *GameCenterLeaderboardSetLocalization
	GameCenterLeaderboardSetRelease *GameCenterLeaderboardSetRelease
}

// GameCenterDetailAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterDetail wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterDetailAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterDetail) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterDetail: v,
	}
}

// GameCenterGroupAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterGroup wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterGroupAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterGroup) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterGroup: v,
	}
}

// GameCenterLeaderboardAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterLeaderboard wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterLeaderboardAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterLeaderboard) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterLeaderboard: v,
	}
}

// GameCenterLeaderboardSetAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterLeaderboardSet wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterLeaderboardSetAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterLeaderboardSet) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterLeaderboardSet: v,
	}
}

// GameCenterLeaderboardSetLocalizationAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterLeaderboardSetLocalization wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterLeaderboardSetLocalizationAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterLeaderboardSetLocalization) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterLeaderboardSetLocalization: v,
	}
}

// GameCenterLeaderboardSetReleaseAsGameCenterLeaderboardSetsResponseIncludedInner is a convenience function that returns GameCenterLeaderboardSetRelease wrapped in GameCenterLeaderboardSetsResponseIncludedInner
func GameCenterLeaderboardSetReleaseAsGameCenterLeaderboardSetsResponseIncludedInner(v *GameCenterLeaderboardSetRelease) GameCenterLeaderboardSetsResponseIncludedInner {
	return GameCenterLeaderboardSetsResponseIncludedInner{
		GameCenterLeaderboardSetRelease: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterLeaderboardSetsResponseIncludedInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into GameCenterGroup
	err = newStrictDecoder(data).Decode(&dst.GameCenterGroup)
	if err == nil {
		jsonGameCenterGroup, _ := json.Marshal(dst.GameCenterGroup)
		if string(jsonGameCenterGroup) == "{}" { // empty struct
			dst.GameCenterGroup = nil
		} else {
			if err = validator.Validate(dst.GameCenterGroup); err != nil {
				dst.GameCenterGroup = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterGroup = nil
	}

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

	// try to unmarshal data into GameCenterLeaderboardSetLocalization
	err = newStrictDecoder(data).Decode(&dst.GameCenterLeaderboardSetLocalization)
	if err == nil {
		jsonGameCenterLeaderboardSetLocalization, _ := json.Marshal(dst.GameCenterLeaderboardSetLocalization)
		if string(jsonGameCenterLeaderboardSetLocalization) == "{}" { // empty struct
			dst.GameCenterLeaderboardSetLocalization = nil
		} else {
			if err = validator.Validate(dst.GameCenterLeaderboardSetLocalization); err != nil {
				dst.GameCenterLeaderboardSetLocalization = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterLeaderboardSetLocalization = nil
	}

	// try to unmarshal data into GameCenterLeaderboardSetRelease
	err = newStrictDecoder(data).Decode(&dst.GameCenterLeaderboardSetRelease)
	if err == nil {
		jsonGameCenterLeaderboardSetRelease, _ := json.Marshal(dst.GameCenterLeaderboardSetRelease)
		if string(jsonGameCenterLeaderboardSetRelease) == "{}" { // empty struct
			dst.GameCenterLeaderboardSetRelease = nil
		} else {
			if err = validator.Validate(dst.GameCenterLeaderboardSetRelease); err != nil {
				dst.GameCenterLeaderboardSetRelease = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterLeaderboardSetRelease = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GameCenterDetail = nil
		dst.GameCenterGroup = nil
		dst.GameCenterLeaderboard = nil
		dst.GameCenterLeaderboardSet = nil
		dst.GameCenterLeaderboardSetLocalization = nil
		dst.GameCenterLeaderboardSetRelease = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterLeaderboardSetsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterLeaderboardSetsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterLeaderboardSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.GameCenterDetail != nil {
		return json.Marshal(&src.GameCenterDetail)
	}

	if src.GameCenterGroup != nil {
		return json.Marshal(&src.GameCenterGroup)
	}

	if src.GameCenterLeaderboard != nil {
		return json.Marshal(&src.GameCenterLeaderboard)
	}

	if src.GameCenterLeaderboardSet != nil {
		return json.Marshal(&src.GameCenterLeaderboardSet)
	}

	if src.GameCenterLeaderboardSetLocalization != nil {
		return json.Marshal(&src.GameCenterLeaderboardSetLocalization)
	}

	if src.GameCenterLeaderboardSetRelease != nil {
		return json.Marshal(&src.GameCenterLeaderboardSetRelease)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterLeaderboardSetsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GameCenterDetail != nil {
		return obj.GameCenterDetail
	}

	if obj.GameCenterGroup != nil {
		return obj.GameCenterGroup
	}

	if obj.GameCenterLeaderboard != nil {
		return obj.GameCenterLeaderboard
	}

	if obj.GameCenterLeaderboardSet != nil {
		return obj.GameCenterLeaderboardSet
	}

	if obj.GameCenterLeaderboardSetLocalization != nil {
		return obj.GameCenterLeaderboardSetLocalization
	}

	if obj.GameCenterLeaderboardSetRelease != nil {
		return obj.GameCenterLeaderboardSetRelease
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterLeaderboardSetsResponseIncludedInner struct {
	value *GameCenterLeaderboardSetsResponseIncludedInner
	isSet bool
}

func (v NullableGameCenterLeaderboardSetsResponseIncludedInner) Get() *GameCenterLeaderboardSetsResponseIncludedInner {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetsResponseIncludedInner) Set(val *GameCenterLeaderboardSetsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetsResponseIncludedInner(val *GameCenterLeaderboardSetsResponseIncludedInner) *NullableGameCenterLeaderboardSetsResponseIncludedInner {
	return &NullableGameCenterLeaderboardSetsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

