/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameCenterAchievementLocalizationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationCreateRequestDataRelationships{}

// GameCenterAchievementLocalizationCreateRequestDataRelationships struct for GameCenterAchievementLocalizationCreateRequestDataRelationships
type GameCenterAchievementLocalizationCreateRequestDataRelationships struct {
	GameCenterAchievement GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement `json:"gameCenterAchievement"`
}

type _GameCenterAchievementLocalizationCreateRequestDataRelationships GameCenterAchievementLocalizationCreateRequestDataRelationships

// NewGameCenterAchievementLocalizationCreateRequestDataRelationships instantiates a new GameCenterAchievementLocalizationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationCreateRequestDataRelationships(gameCenterAchievement GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement) *GameCenterAchievementLocalizationCreateRequestDataRelationships {
	this := GameCenterAchievementLocalizationCreateRequestDataRelationships{}
	this.GameCenterAchievement = gameCenterAchievement
	return &this
}

// NewGameCenterAchievementLocalizationCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterAchievementLocalizationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationCreateRequestDataRelationshipsWithDefaults() *GameCenterAchievementLocalizationCreateRequestDataRelationships {
	this := GameCenterAchievementLocalizationCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterAchievement returns the GameCenterAchievement field value
func (o *GameCenterAchievementLocalizationCreateRequestDataRelationships) GetGameCenterAchievement() GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement {
	if o == nil {
		var ret GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement
		return ret
	}

	return o.GameCenterAchievement
}

// GetGameCenterAchievementOk returns a tuple with the GameCenterAchievement field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestDataRelationships) GetGameCenterAchievementOk() (*GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterAchievement, true
}

// SetGameCenterAchievement sets field value
func (o *GameCenterAchievementLocalizationCreateRequestDataRelationships) SetGameCenterAchievement(v GameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement) {
	o.GameCenterAchievement = v
}

func (o GameCenterAchievementLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameCenterAchievement"] = o.GameCenterAchievement
	return toSerialize, nil
}

func (o *GameCenterAchievementLocalizationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGameCenterAchievementLocalizationCreateRequestDataRelationships := _GameCenterAchievementLocalizationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementLocalizationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementLocalizationCreateRequestDataRelationships(varGameCenterAchievementLocalizationCreateRequestDataRelationships)

	return err
}

type NullableGameCenterAchievementLocalizationCreateRequestDataRelationships struct {
	value *GameCenterAchievementLocalizationCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) Get() *GameCenterAchievementLocalizationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) Set(val *GameCenterAchievementLocalizationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationCreateRequestDataRelationships(val *GameCenterAchievementLocalizationCreateRequestDataRelationships) *NullableGameCenterAchievementLocalizationCreateRequestDataRelationships {
	return &NullableGameCenterAchievementLocalizationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

