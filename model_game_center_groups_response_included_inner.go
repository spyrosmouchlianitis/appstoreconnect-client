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

// GameCenterGroupsResponseIncludedInner - struct for GameCenterGroupsResponseIncludedInner
type GameCenterGroupsResponseIncludedInner struct {
	GameCenterAchievement *GameCenterAchievement
	GameCenterDetail *GameCenterDetail
	GameCenterLeaderboard *GameCenterLeaderboard
	GameCenterLeaderboardSet *GameCenterLeaderboardSet
}

// GameCenterAchievementAsGameCenterGroupsResponseIncludedInner is a convenience function that returns GameCenterAchievement wrapped in GameCenterGroupsResponseIncludedInner
func GameCenterAchievementAsGameCenterGroupsResponseIncludedInner(v *GameCenterAchievement) GameCenterGroupsResponseIncludedInner {
	return GameCenterGroupsResponseIncludedInner{
		GameCenterAchievement: v,
	}
}

// GameCenterDetailAsGameCenterGroupsResponseIncludedInner is a convenience function that returns GameCenterDetail wrapped in GameCenterGroupsResponseIncludedInner
func GameCenterDetailAsGameCenterGroupsResponseIncludedInner(v *GameCenterDetail) GameCenterGroupsResponseIncludedInner {
	return GameCenterGroupsResponseIncludedInner{
		GameCenterDetail: v,
	}
}

// GameCenterLeaderboardAsGameCenterGroupsResponseIncludedInner is a convenience function that returns GameCenterLeaderboard wrapped in GameCenterGroupsResponseIncludedInner
func GameCenterLeaderboardAsGameCenterGroupsResponseIncludedInner(v *GameCenterLeaderboard) GameCenterGroupsResponseIncludedInner {
	return GameCenterGroupsResponseIncludedInner{
		GameCenterLeaderboard: v,
	}
}

// GameCenterLeaderboardSetAsGameCenterGroupsResponseIncludedInner is a convenience function that returns GameCenterLeaderboardSet wrapped in GameCenterGroupsResponseIncludedInner
func GameCenterLeaderboardSetAsGameCenterGroupsResponseIncludedInner(v *GameCenterLeaderboardSet) GameCenterGroupsResponseIncludedInner {
	return GameCenterGroupsResponseIncludedInner{
		GameCenterLeaderboardSet: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterGroupsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GameCenterAchievement
	err = newStrictDecoder(data).Decode(&dst.GameCenterAchievement)
	if err == nil {
		jsonGameCenterAchievement, _ := json.Marshal(dst.GameCenterAchievement)
		if string(jsonGameCenterAchievement) == "{}" { // empty struct
			dst.GameCenterAchievement = nil
		} else {
			if err = validator.Validate(dst.GameCenterAchievement); err != nil {
				dst.GameCenterAchievement = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameCenterAchievement = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GameCenterAchievement = nil
		dst.GameCenterDetail = nil
		dst.GameCenterLeaderboard = nil
		dst.GameCenterLeaderboardSet = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterGroupsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterGroupsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.GameCenterAchievement != nil {
		return json.Marshal(&src.GameCenterAchievement)
	}

	if src.GameCenterDetail != nil {
		return json.Marshal(&src.GameCenterDetail)
	}

	if src.GameCenterLeaderboard != nil {
		return json.Marshal(&src.GameCenterLeaderboard)
	}

	if src.GameCenterLeaderboardSet != nil {
		return json.Marshal(&src.GameCenterLeaderboardSet)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterGroupsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GameCenterAchievement != nil {
		return obj.GameCenterAchievement
	}

	if obj.GameCenterDetail != nil {
		return obj.GameCenterDetail
	}

	if obj.GameCenterLeaderboard != nil {
		return obj.GameCenterLeaderboard
	}

	if obj.GameCenterLeaderboardSet != nil {
		return obj.GameCenterLeaderboardSet
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterGroupsResponseIncludedInner struct {
	value *GameCenterGroupsResponseIncludedInner
	isSet bool
}

func (v NullableGameCenterGroupsResponseIncludedInner) Get() *GameCenterGroupsResponseIncludedInner {
	return v.value
}

func (v *NullableGameCenterGroupsResponseIncludedInner) Set(val *GameCenterGroupsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupsResponseIncludedInner(val *GameCenterGroupsResponseIncludedInner) *NullableGameCenterGroupsResponseIncludedInner {
	return &NullableGameCenterGroupsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


