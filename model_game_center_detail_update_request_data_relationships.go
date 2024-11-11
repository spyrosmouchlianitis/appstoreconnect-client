/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the GameCenterDetailUpdateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailUpdateRequestDataRelationships{}

// GameCenterDetailUpdateRequestDataRelationships struct for GameCenterDetailUpdateRequestDataRelationships
type GameCenterDetailUpdateRequestDataRelationships struct {
	GameCenterGroup *GameCenterAchievementRelationshipsGameCenterGroup `json:"gameCenterGroup,omitempty"`
	DefaultLeaderboard *GameCenterDetailRelationshipsDefaultLeaderboard `json:"defaultLeaderboard,omitempty"`
	DefaultGroupLeaderboard *GameCenterDetailRelationshipsDefaultLeaderboard `json:"defaultGroupLeaderboard,omitempty"`
}

// NewGameCenterDetailUpdateRequestDataRelationships instantiates a new GameCenterDetailUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailUpdateRequestDataRelationships() *GameCenterDetailUpdateRequestDataRelationships {
	this := GameCenterDetailUpdateRequestDataRelationships{}
	return &this
}

// NewGameCenterDetailUpdateRequestDataRelationshipsWithDefaults instantiates a new GameCenterDetailUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailUpdateRequestDataRelationshipsWithDefaults() *GameCenterDetailUpdateRequestDataRelationships {
	this := GameCenterDetailUpdateRequestDataRelationships{}
	return &this
}

// GetGameCenterGroup returns the GameCenterGroup field value if set, zero value otherwise.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetGameCenterGroup() GameCenterAchievementRelationshipsGameCenterGroup {
	if o == nil || IsNil(o.GameCenterGroup) {
		var ret GameCenterAchievementRelationshipsGameCenterGroup
		return ret
	}
	return *o.GameCenterGroup
}

// GetGameCenterGroupOk returns a tuple with the GameCenterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetGameCenterGroupOk() (*GameCenterAchievementRelationshipsGameCenterGroup, bool) {
	if o == nil || IsNil(o.GameCenterGroup) {
		return nil, false
	}
	return o.GameCenterGroup, true
}

// HasGameCenterGroup returns a boolean if a field has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) HasGameCenterGroup() bool {
	if o != nil && !IsNil(o.GameCenterGroup) {
		return true
	}

	return false
}

// SetGameCenterGroup gets a reference to the given GameCenterAchievementRelationshipsGameCenterGroup and assigns it to the GameCenterGroup field.
func (o *GameCenterDetailUpdateRequestDataRelationships) SetGameCenterGroup(v GameCenterAchievementRelationshipsGameCenterGroup) {
	o.GameCenterGroup = &v
}

// GetDefaultLeaderboard returns the DefaultLeaderboard field value if set, zero value otherwise.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetDefaultLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard {
	if o == nil || IsNil(o.DefaultLeaderboard) {
		var ret GameCenterDetailRelationshipsDefaultLeaderboard
		return ret
	}
	return *o.DefaultLeaderboard
}

// GetDefaultLeaderboardOk returns a tuple with the DefaultLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetDefaultLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool) {
	if o == nil || IsNil(o.DefaultLeaderboard) {
		return nil, false
	}
	return o.DefaultLeaderboard, true
}

// HasDefaultLeaderboard returns a boolean if a field has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) HasDefaultLeaderboard() bool {
	if o != nil && !IsNil(o.DefaultLeaderboard) {
		return true
	}

	return false
}

// SetDefaultLeaderboard gets a reference to the given GameCenterDetailRelationshipsDefaultLeaderboard and assigns it to the DefaultLeaderboard field.
func (o *GameCenterDetailUpdateRequestDataRelationships) SetDefaultLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard) {
	o.DefaultLeaderboard = &v
}

// GetDefaultGroupLeaderboard returns the DefaultGroupLeaderboard field value if set, zero value otherwise.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetDefaultGroupLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard {
	if o == nil || IsNil(o.DefaultGroupLeaderboard) {
		var ret GameCenterDetailRelationshipsDefaultLeaderboard
		return ret
	}
	return *o.DefaultGroupLeaderboard
}

// GetDefaultGroupLeaderboardOk returns a tuple with the DefaultGroupLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) GetDefaultGroupLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool) {
	if o == nil || IsNil(o.DefaultGroupLeaderboard) {
		return nil, false
	}
	return o.DefaultGroupLeaderboard, true
}

// HasDefaultGroupLeaderboard returns a boolean if a field has been set.
func (o *GameCenterDetailUpdateRequestDataRelationships) HasDefaultGroupLeaderboard() bool {
	if o != nil && !IsNil(o.DefaultGroupLeaderboard) {
		return true
	}

	return false
}

// SetDefaultGroupLeaderboard gets a reference to the given GameCenterDetailRelationshipsDefaultLeaderboard and assigns it to the DefaultGroupLeaderboard field.
func (o *GameCenterDetailUpdateRequestDataRelationships) SetDefaultGroupLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard) {
	o.DefaultGroupLeaderboard = &v
}

func (o GameCenterDetailUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailUpdateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterGroup) {
		toSerialize["gameCenterGroup"] = o.GameCenterGroup
	}
	if !IsNil(o.DefaultLeaderboard) {
		toSerialize["defaultLeaderboard"] = o.DefaultLeaderboard
	}
	if !IsNil(o.DefaultGroupLeaderboard) {
		toSerialize["defaultGroupLeaderboard"] = o.DefaultGroupLeaderboard
	}
	return toSerialize, nil
}

type NullableGameCenterDetailUpdateRequestDataRelationships struct {
	value *GameCenterDetailUpdateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterDetailUpdateRequestDataRelationships) Get() *GameCenterDetailUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterDetailUpdateRequestDataRelationships) Set(val *GameCenterDetailUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailUpdateRequestDataRelationships(val *GameCenterDetailUpdateRequestDataRelationships) *NullableGameCenterDetailUpdateRequestDataRelationships {
	return &NullableGameCenterDetailUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterDetailUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

