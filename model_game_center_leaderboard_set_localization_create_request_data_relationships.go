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

// checks if the GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships{}

// GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships struct for GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships
type GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships struct {
	GameCenterLeaderboardSet GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet `json:"gameCenterLeaderboardSet"`
}

type _GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships

// NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships(gameCenterLeaderboardSet GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet) *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships {
	this := GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships{}
	this.GameCenterLeaderboardSet = gameCenterLeaderboardSet
	return &this
}

// NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsWithDefaults() *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships {
	this := GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterLeaderboardSet returns the GameCenterLeaderboardSet field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) GetGameCenterLeaderboardSet() GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet {
	if o == nil {
		var ret GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet
		return ret
	}

	return o.GameCenterLeaderboardSet
}

// GetGameCenterLeaderboardSetOk returns a tuple with the GameCenterLeaderboardSet field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) GetGameCenterLeaderboardSetOk() (*GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterLeaderboardSet, true
}

// SetGameCenterLeaderboardSet sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) SetGameCenterLeaderboardSet(v GameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet) {
	o.GameCenterLeaderboardSet = v
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameCenterLeaderboardSet"] = o.GameCenterLeaderboardSet
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gameCenterLeaderboardSet",
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

	varGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships := _GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships(varGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships struct {
	value *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) Get() *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) Set(val *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships(val *GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships {
	return &NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


