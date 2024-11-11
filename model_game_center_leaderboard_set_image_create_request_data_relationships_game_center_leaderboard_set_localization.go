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

// checks if the GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization{}

// GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization struct for GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization
type GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization struct {
	Data GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData `json:"data"`
}

type _GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization

// NewGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization instantiates a new GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization(data GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData) *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization {
	this := GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalizationWithDefaults instantiates a new GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalizationWithDefaults() *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization {
	this := GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) GetData() GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData {
	if o == nil {
		var ret GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) GetDataOk() (*GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) SetData(v GameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization := _GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization(varGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization)

	return err
}

type NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization struct {
	value *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization
	isSet bool
}

func (v NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) Get() *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) Set(val *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization(val *GameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) *NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization {
	return &NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


