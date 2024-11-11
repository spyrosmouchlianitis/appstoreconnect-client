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

// checks if the GameCenterLeaderboardSetLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationRelationships{}

// GameCenterLeaderboardSetLocalizationRelationships struct for GameCenterLeaderboardSetLocalizationRelationships
type GameCenterLeaderboardSetLocalizationRelationships struct {
	GameCenterLeaderboardSet *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet `json:"gameCenterLeaderboardSet,omitempty"`
	GameCenterLeaderboardSetImage *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage `json:"gameCenterLeaderboardSetImage,omitempty"`
}

// NewGameCenterLeaderboardSetLocalizationRelationships instantiates a new GameCenterLeaderboardSetLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationRelationships() *GameCenterLeaderboardSetLocalizationRelationships {
	this := GameCenterLeaderboardSetLocalizationRelationships{}
	return &this
}

// NewGameCenterLeaderboardSetLocalizationRelationshipsWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationRelationshipsWithDefaults() *GameCenterLeaderboardSetLocalizationRelationships {
	this := GameCenterLeaderboardSetLocalizationRelationships{}
	return &this
}

// GetGameCenterLeaderboardSet returns the GameCenterLeaderboardSet field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationRelationships) GetGameCenterLeaderboardSet() GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet {
	if o == nil || IsNil(o.GameCenterLeaderboardSet) {
		var ret GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet
		return ret
	}
	return *o.GameCenterLeaderboardSet
}

// GetGameCenterLeaderboardSetOk returns a tuple with the GameCenterLeaderboardSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationships) GetGameCenterLeaderboardSetOk() (*GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSet) {
		return nil, false
	}
	return o.GameCenterLeaderboardSet, true
}

// HasGameCenterLeaderboardSet returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationships) HasGameCenterLeaderboardSet() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSet) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSet gets a reference to the given GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet and assigns it to the GameCenterLeaderboardSet field.
func (o *GameCenterLeaderboardSetLocalizationRelationships) SetGameCenterLeaderboardSet(v GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet) {
	o.GameCenterLeaderboardSet = &v
}

// GetGameCenterLeaderboardSetImage returns the GameCenterLeaderboardSetImage field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationRelationships) GetGameCenterLeaderboardSetImage() GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage {
	if o == nil || IsNil(o.GameCenterLeaderboardSetImage) {
		var ret GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage
		return ret
	}
	return *o.GameCenterLeaderboardSetImage
}

// GetGameCenterLeaderboardSetImageOk returns a tuple with the GameCenterLeaderboardSetImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationships) GetGameCenterLeaderboardSetImageOk() (*GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSetImage) {
		return nil, false
	}
	return o.GameCenterLeaderboardSetImage, true
}

// HasGameCenterLeaderboardSetImage returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationships) HasGameCenterLeaderboardSetImage() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSetImage) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSetImage gets a reference to the given GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage and assigns it to the GameCenterLeaderboardSetImage field.
func (o *GameCenterLeaderboardSetLocalizationRelationships) SetGameCenterLeaderboardSetImage(v GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) {
	o.GameCenterLeaderboardSetImage = &v
}

func (o GameCenterLeaderboardSetLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterLeaderboardSet) {
		toSerialize["gameCenterLeaderboardSet"] = o.GameCenterLeaderboardSet
	}
	if !IsNil(o.GameCenterLeaderboardSetImage) {
		toSerialize["gameCenterLeaderboardSetImage"] = o.GameCenterLeaderboardSetImage
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetLocalizationRelationships struct {
	value *GameCenterLeaderboardSetLocalizationRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationships) Get() *GameCenterLeaderboardSetLocalizationRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationships) Set(val *GameCenterLeaderboardSetLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationRelationships(val *GameCenterLeaderboardSetLocalizationRelationships) *NullableGameCenterLeaderboardSetLocalizationRelationships {
	return &NullableGameCenterLeaderboardSetLocalizationRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


