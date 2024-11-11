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

// checks if the GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}

// GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct for GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard
type GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct {
	Data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
}

type _GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard

// NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard instantiates a new GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardWithDefaults instantiates a new GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardWithDefaults() *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) GetData() GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) GetDataOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) SetData(v GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

func (o GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard := _GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(varGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard)

	return err
}

type NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct {
	value *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Get() *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Set(val *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(val *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	return &NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


