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

// checks if the GameCenterAchievementImageCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementImageCreateRequestDataRelationships{}

// GameCenterAchievementImageCreateRequestDataRelationships struct for GameCenterAchievementImageCreateRequestDataRelationships
type GameCenterAchievementImageCreateRequestDataRelationships struct {
	GameCenterAchievementLocalization GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization `json:"gameCenterAchievementLocalization"`
}

type _GameCenterAchievementImageCreateRequestDataRelationships GameCenterAchievementImageCreateRequestDataRelationships

// NewGameCenterAchievementImageCreateRequestDataRelationships instantiates a new GameCenterAchievementImageCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementImageCreateRequestDataRelationships(gameCenterAchievementLocalization GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization) *GameCenterAchievementImageCreateRequestDataRelationships {
	this := GameCenterAchievementImageCreateRequestDataRelationships{}
	this.GameCenterAchievementLocalization = gameCenterAchievementLocalization
	return &this
}

// NewGameCenterAchievementImageCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterAchievementImageCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementImageCreateRequestDataRelationshipsWithDefaults() *GameCenterAchievementImageCreateRequestDataRelationships {
	this := GameCenterAchievementImageCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterAchievementLocalization returns the GameCenterAchievementLocalization field value
func (o *GameCenterAchievementImageCreateRequestDataRelationships) GetGameCenterAchievementLocalization() GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization {
	if o == nil {
		var ret GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization
		return ret
	}

	return o.GameCenterAchievementLocalization
}

// GetGameCenterAchievementLocalizationOk returns a tuple with the GameCenterAchievementLocalization field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageCreateRequestDataRelationships) GetGameCenterAchievementLocalizationOk() (*GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterAchievementLocalization, true
}

// SetGameCenterAchievementLocalization sets field value
func (o *GameCenterAchievementImageCreateRequestDataRelationships) SetGameCenterAchievementLocalization(v GameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization) {
	o.GameCenterAchievementLocalization = v
}

func (o GameCenterAchievementImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementImageCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameCenterAchievementLocalization"] = o.GameCenterAchievementLocalization
	return toSerialize, nil
}

func (o *GameCenterAchievementImageCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gameCenterAchievementLocalization",
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

	varGameCenterAchievementImageCreateRequestDataRelationships := _GameCenterAchievementImageCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementImageCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementImageCreateRequestDataRelationships(varGameCenterAchievementImageCreateRequestDataRelationships)

	return err
}

type NullableGameCenterAchievementImageCreateRequestDataRelationships struct {
	value *GameCenterAchievementImageCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterAchievementImageCreateRequestDataRelationships) Get() *GameCenterAchievementImageCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementImageCreateRequestDataRelationships) Set(val *GameCenterAchievementImageCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementImageCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementImageCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementImageCreateRequestDataRelationships(val *GameCenterAchievementImageCreateRequestDataRelationships) *NullableGameCenterAchievementImageCreateRequestDataRelationships {
	return &NullableGameCenterAchievementImageCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementImageCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


