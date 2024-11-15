/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the GameCenterAchievementReleaseRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementReleaseRelationships{}

// GameCenterAchievementReleaseRelationships struct for GameCenterAchievementReleaseRelationships
type GameCenterAchievementReleaseRelationships struct {
	GameCenterDetail *GameCenterAchievementReleaseRelationshipsGameCenterDetail `json:"gameCenterDetail,omitempty"`
	GameCenterAchievement *GameCenterAchievementReleaseRelationshipsGameCenterAchievement `json:"gameCenterAchievement,omitempty"`
}

// NewGameCenterAchievementReleaseRelationships instantiates a new GameCenterAchievementReleaseRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementReleaseRelationships() *GameCenterAchievementReleaseRelationships {
	this := GameCenterAchievementReleaseRelationships{}
	return &this
}

// NewGameCenterAchievementReleaseRelationshipsWithDefaults instantiates a new GameCenterAchievementReleaseRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleaseRelationshipsWithDefaults() *GameCenterAchievementReleaseRelationships {
	this := GameCenterAchievementReleaseRelationships{}
	return &this
}

// GetGameCenterDetail returns the GameCenterDetail field value if set, zero value otherwise.
func (o *GameCenterAchievementReleaseRelationships) GetGameCenterDetail() GameCenterAchievementReleaseRelationshipsGameCenterDetail {
	if o == nil || IsNil(o.GameCenterDetail) {
		var ret GameCenterAchievementReleaseRelationshipsGameCenterDetail
		return ret
	}
	return *o.GameCenterDetail
}

// GetGameCenterDetailOk returns a tuple with the GameCenterDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseRelationships) GetGameCenterDetailOk() (*GameCenterAchievementReleaseRelationshipsGameCenterDetail, bool) {
	if o == nil || IsNil(o.GameCenterDetail) {
		return nil, false
	}
	return o.GameCenterDetail, true
}

// HasGameCenterDetail returns a boolean if a field has been set.
func (o *GameCenterAchievementReleaseRelationships) HasGameCenterDetail() bool {
	if o != nil && !IsNil(o.GameCenterDetail) {
		return true
	}

	return false
}

// SetGameCenterDetail gets a reference to the given GameCenterAchievementReleaseRelationshipsGameCenterDetail and assigns it to the GameCenterDetail field.
func (o *GameCenterAchievementReleaseRelationships) SetGameCenterDetail(v GameCenterAchievementReleaseRelationshipsGameCenterDetail) {
	o.GameCenterDetail = &v
}

// GetGameCenterAchievement returns the GameCenterAchievement field value if set, zero value otherwise.
func (o *GameCenterAchievementReleaseRelationships) GetGameCenterAchievement() GameCenterAchievementReleaseRelationshipsGameCenterAchievement {
	if o == nil || IsNil(o.GameCenterAchievement) {
		var ret GameCenterAchievementReleaseRelationshipsGameCenterAchievement
		return ret
	}
	return *o.GameCenterAchievement
}

// GetGameCenterAchievementOk returns a tuple with the GameCenterAchievement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseRelationships) GetGameCenterAchievementOk() (*GameCenterAchievementReleaseRelationshipsGameCenterAchievement, bool) {
	if o == nil || IsNil(o.GameCenterAchievement) {
		return nil, false
	}
	return o.GameCenterAchievement, true
}

// HasGameCenterAchievement returns a boolean if a field has been set.
func (o *GameCenterAchievementReleaseRelationships) HasGameCenterAchievement() bool {
	if o != nil && !IsNil(o.GameCenterAchievement) {
		return true
	}

	return false
}

// SetGameCenterAchievement gets a reference to the given GameCenterAchievementReleaseRelationshipsGameCenterAchievement and assigns it to the GameCenterAchievement field.
func (o *GameCenterAchievementReleaseRelationships) SetGameCenterAchievement(v GameCenterAchievementReleaseRelationshipsGameCenterAchievement) {
	o.GameCenterAchievement = &v
}

func (o GameCenterAchievementReleaseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementReleaseRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterDetail) {
		toSerialize["gameCenterDetail"] = o.GameCenterDetail
	}
	if !IsNil(o.GameCenterAchievement) {
		toSerialize["gameCenterAchievement"] = o.GameCenterAchievement
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementReleaseRelationships struct {
	value *GameCenterAchievementReleaseRelationships
	isSet bool
}

func (v NullableGameCenterAchievementReleaseRelationships) Get() *GameCenterAchievementReleaseRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementReleaseRelationships) Set(val *GameCenterAchievementReleaseRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementReleaseRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementReleaseRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementReleaseRelationships(val *GameCenterAchievementReleaseRelationships) *NullableGameCenterAchievementReleaseRelationships {
	return &NullableGameCenterAchievementReleaseRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementReleaseRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementReleaseRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


