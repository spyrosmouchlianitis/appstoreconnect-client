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

// checks if the GameCenterAchievementImageRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementImageRelationships{}

// GameCenterAchievementImageRelationships struct for GameCenterAchievementImageRelationships
type GameCenterAchievementImageRelationships struct {
	GameCenterAchievementLocalization *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization `json:"gameCenterAchievementLocalization,omitempty"`
}

// NewGameCenterAchievementImageRelationships instantiates a new GameCenterAchievementImageRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementImageRelationships() *GameCenterAchievementImageRelationships {
	this := GameCenterAchievementImageRelationships{}
	return &this
}

// NewGameCenterAchievementImageRelationshipsWithDefaults instantiates a new GameCenterAchievementImageRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementImageRelationshipsWithDefaults() *GameCenterAchievementImageRelationships {
	this := GameCenterAchievementImageRelationships{}
	return &this
}

// GetGameCenterAchievementLocalization returns the GameCenterAchievementLocalization field value if set, zero value otherwise.
func (o *GameCenterAchievementImageRelationships) GetGameCenterAchievementLocalization() GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization {
	if o == nil || IsNil(o.GameCenterAchievementLocalization) {
		var ret GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization
		return ret
	}
	return *o.GameCenterAchievementLocalization
}

// GetGameCenterAchievementLocalizationOk returns a tuple with the GameCenterAchievementLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageRelationships) GetGameCenterAchievementLocalizationOk() (*GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization, bool) {
	if o == nil || IsNil(o.GameCenterAchievementLocalization) {
		return nil, false
	}
	return o.GameCenterAchievementLocalization, true
}

// HasGameCenterAchievementLocalization returns a boolean if a field has been set.
func (o *GameCenterAchievementImageRelationships) HasGameCenterAchievementLocalization() bool {
	if o != nil && !IsNil(o.GameCenterAchievementLocalization) {
		return true
	}

	return false
}

// SetGameCenterAchievementLocalization gets a reference to the given GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization and assigns it to the GameCenterAchievementLocalization field.
func (o *GameCenterAchievementImageRelationships) SetGameCenterAchievementLocalization(v GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) {
	o.GameCenterAchievementLocalization = &v
}

func (o GameCenterAchievementImageRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementImageRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterAchievementLocalization) {
		toSerialize["gameCenterAchievementLocalization"] = o.GameCenterAchievementLocalization
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementImageRelationships struct {
	value *GameCenterAchievementImageRelationships
	isSet bool
}

func (v NullableGameCenterAchievementImageRelationships) Get() *GameCenterAchievementImageRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementImageRelationships) Set(val *GameCenterAchievementImageRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementImageRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementImageRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementImageRelationships(val *GameCenterAchievementImageRelationships) *NullableGameCenterAchievementImageRelationships {
	return &NullableGameCenterAchievementImageRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementImageRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementImageRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


