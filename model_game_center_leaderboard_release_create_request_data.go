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

// checks if the GameCenterLeaderboardReleaseCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardReleaseCreateRequestData{}

// GameCenterLeaderboardReleaseCreateRequestData struct for GameCenterLeaderboardReleaseCreateRequestData
type GameCenterLeaderboardReleaseCreateRequestData struct {
	Type string `json:"type"`
	Relationships GameCenterLeaderboardReleaseCreateRequestDataRelationships `json:"relationships"`
}

type _GameCenterLeaderboardReleaseCreateRequestData GameCenterLeaderboardReleaseCreateRequestData

// NewGameCenterLeaderboardReleaseCreateRequestData instantiates a new GameCenterLeaderboardReleaseCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardReleaseCreateRequestData(type_ string, relationships GameCenterLeaderboardReleaseCreateRequestDataRelationships) *GameCenterLeaderboardReleaseCreateRequestData {
	this := GameCenterLeaderboardReleaseCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewGameCenterLeaderboardReleaseCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardReleaseCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardReleaseCreateRequestDataWithDefaults() *GameCenterLeaderboardReleaseCreateRequestData {
	this := GameCenterLeaderboardReleaseCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardReleaseCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleaseCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardReleaseCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterLeaderboardReleaseCreateRequestData) GetRelationships() GameCenterLeaderboardReleaseCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterLeaderboardReleaseCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardReleaseCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardReleaseCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterLeaderboardReleaseCreateRequestData) SetRelationships(v GameCenterLeaderboardReleaseCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterLeaderboardReleaseCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardReleaseCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *GameCenterLeaderboardReleaseCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"relationships",
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

	varGameCenterLeaderboardReleaseCreateRequestData := _GameCenterLeaderboardReleaseCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardReleaseCreateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardReleaseCreateRequestData(varGameCenterLeaderboardReleaseCreateRequestData)

	return err
}

type NullableGameCenterLeaderboardReleaseCreateRequestData struct {
	value *GameCenterLeaderboardReleaseCreateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardReleaseCreateRequestData) Get() *GameCenterLeaderboardReleaseCreateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardReleaseCreateRequestData) Set(val *GameCenterLeaderboardReleaseCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardReleaseCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardReleaseCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardReleaseCreateRequestData(val *GameCenterLeaderboardReleaseCreateRequestData) *NullableGameCenterLeaderboardReleaseCreateRequestData {
	return &NullableGameCenterLeaderboardReleaseCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardReleaseCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardReleaseCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


