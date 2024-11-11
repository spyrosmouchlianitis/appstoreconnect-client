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

// checks if the GameCenterLeaderboardImageCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardImageCreateRequestDataRelationships{}

// GameCenterLeaderboardImageCreateRequestDataRelationships struct for GameCenterLeaderboardImageCreateRequestDataRelationships
type GameCenterLeaderboardImageCreateRequestDataRelationships struct {
	GameCenterLeaderboardLocalization GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization `json:"gameCenterLeaderboardLocalization"`
}

type _GameCenterLeaderboardImageCreateRequestDataRelationships GameCenterLeaderboardImageCreateRequestDataRelationships

// NewGameCenterLeaderboardImageCreateRequestDataRelationships instantiates a new GameCenterLeaderboardImageCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardImageCreateRequestDataRelationships(gameCenterLeaderboardLocalization GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization) *GameCenterLeaderboardImageCreateRequestDataRelationships {
	this := GameCenterLeaderboardImageCreateRequestDataRelationships{}
	this.GameCenterLeaderboardLocalization = gameCenterLeaderboardLocalization
	return &this
}

// NewGameCenterLeaderboardImageCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterLeaderboardImageCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardImageCreateRequestDataRelationshipsWithDefaults() *GameCenterLeaderboardImageCreateRequestDataRelationships {
	this := GameCenterLeaderboardImageCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterLeaderboardLocalization returns the GameCenterLeaderboardLocalization field value
func (o *GameCenterLeaderboardImageCreateRequestDataRelationships) GetGameCenterLeaderboardLocalization() GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization {
	if o == nil {
		var ret GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization
		return ret
	}

	return o.GameCenterLeaderboardLocalization
}

// GetGameCenterLeaderboardLocalizationOk returns a tuple with the GameCenterLeaderboardLocalization field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardImageCreateRequestDataRelationships) GetGameCenterLeaderboardLocalizationOk() (*GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameCenterLeaderboardLocalization, true
}

// SetGameCenterLeaderboardLocalization sets field value
func (o *GameCenterLeaderboardImageCreateRequestDataRelationships) SetGameCenterLeaderboardLocalization(v GameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization) {
	o.GameCenterLeaderboardLocalization = v
}

func (o GameCenterLeaderboardImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardImageCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gameCenterLeaderboardLocalization"] = o.GameCenterLeaderboardLocalization
	return toSerialize, nil
}

func (o *GameCenterLeaderboardImageCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gameCenterLeaderboardLocalization",
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

	varGameCenterLeaderboardImageCreateRequestDataRelationships := _GameCenterLeaderboardImageCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardImageCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardImageCreateRequestDataRelationships(varGameCenterLeaderboardImageCreateRequestDataRelationships)

	return err
}

type NullableGameCenterLeaderboardImageCreateRequestDataRelationships struct {
	value *GameCenterLeaderboardImageCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardImageCreateRequestDataRelationships) Get() *GameCenterLeaderboardImageCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardImageCreateRequestDataRelationships) Set(val *GameCenterLeaderboardImageCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardImageCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardImageCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardImageCreateRequestDataRelationships(val *GameCenterLeaderboardImageCreateRequestDataRelationships) *NullableGameCenterLeaderboardImageCreateRequestDataRelationships {
	return &NullableGameCenterLeaderboardImageCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardImageCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


