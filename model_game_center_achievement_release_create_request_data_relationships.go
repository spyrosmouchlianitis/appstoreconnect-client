/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameCenterAchievementReleaseCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementReleaseCreateRequestDataRelationships{}

// GameCenterAchievementReleaseCreateRequestDataRelationships struct for GameCenterAchievementReleaseCreateRequestDataRelationships
type GameCenterAchievementReleaseCreateRequestDataRelationships struct {
	GameCenterDetail GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail `json:"gameCenterDetail"`
	GameCenterAchievement GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement `json:"gameCenterAchievement"`
}

type _GameCenterAchievementReleaseCreateRequestDataRelationships GameCenterAchievementReleaseCreateRequestDataRelationships

// NewGameCenterAchievementReleaseCreateRequestDataRelationships instantiates a new GameCenterAchievementReleaseCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementReleaseCreateRequestDataRelationships(gameCenterDetail GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail, gameCenterAchievement GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement) *GameCenterAchievementReleaseCreateRequestDataRelationships {
	this := GameCenterAchievementReleaseCreateRequestDataRelationships{}
	this.GameCenterDetail = gameCenterDetail
	this.GameCenterAchievement = gameCenterAchievement
	return &this
}

// NewGameCenterAchievementReleaseCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterAchievementReleaseCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleaseCreateRequestDataRelationshipsWithDefaults() *GameCenterAchievementReleaseCreateRequestDataRelationships {
	this := GameCenterAchievementReleaseCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterDetail returns the GameCenterDetail field value
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) GetGameCenterDetail() GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail {
	if o == nil {
		var ret GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail
		return ret
	}

	return o.GameCenterDetail
}

// GetGameCenterDetailOk returns a tuple with the GameCenterDetail field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) GetGameCenterDetailOk() (*GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterDetail, true
}

// SetGameCenterDetail sets field value
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) SetGameCenterDetail(v GameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail) {
	o.GameCenterDetail = v
}

// GetGameCenterAchievement returns the GameCenterAchievement field value
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) GetGameCenterAchievement() GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement {
	if o == nil {
		var ret GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement
		return ret
	}

	return o.GameCenterAchievement
}

// GetGameCenterAchievementOk returns a tuple with the GameCenterAchievement field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) GetGameCenterAchievementOk() (*GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterAchievement, true
}

// SetGameCenterAchievement sets field value
func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) SetGameCenterAchievement(v GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement) {
	o.GameCenterAchievement = v
}

func (o GameCenterAchievementReleaseCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementReleaseCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameCenterDetail"] = o.GameCenterDetail
	toSerialize["gameCenterAchievement"] = o.GameCenterAchievement
	return toSerialize, nil
}

func (o *GameCenterAchievementReleaseCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gameCenterDetail",
		"gameCenterAchievement",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGameCenterAchievementReleaseCreateRequestDataRelationships := _GameCenterAchievementReleaseCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementReleaseCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementReleaseCreateRequestDataRelationships(varGameCenterAchievementReleaseCreateRequestDataRelationships)

	return err
}

type NullableGameCenterAchievementReleaseCreateRequestDataRelationships struct {
	value *GameCenterAchievementReleaseCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterAchievementReleaseCreateRequestDataRelationships) Get() *GameCenterAchievementReleaseCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementReleaseCreateRequestDataRelationships) Set(val *GameCenterAchievementReleaseCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementReleaseCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementReleaseCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementReleaseCreateRequestDataRelationships(val *GameCenterAchievementReleaseCreateRequestDataRelationships) *NullableGameCenterAchievementReleaseCreateRequestDataRelationships {
	return &NullableGameCenterAchievementReleaseCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementReleaseCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementReleaseCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


