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

// checks if the GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization{}

// GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization struct for GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization
type GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization struct {
	Data *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData `json:"data,omitempty"`
}

// NewGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization instantiates a new GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization() *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization {
	this := GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization{}
	return &this
}

// NewGameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationWithDefaults instantiates a new GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationWithDefaults() *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization {
	this := GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) GetData() GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) GetDataOk() (*GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData and assigns it to the Data field.
func (o *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) SetData(v GameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData) {
	o.Data = &v
}

func (o GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization struct {
	value *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization
	isSet bool
}

func (v NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) Get() *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization {
	return v.value
}

func (v *NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) Set(val *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization(val *GameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) *NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization {
	return &NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization{value: val, isSet: true}
}

func (v NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementImageRelationshipsGameCenterAchievementLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


