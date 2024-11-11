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

// checks if the GameCenterAchievementRelationshipsReleasesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementRelationshipsReleasesDataInner{}

// GameCenterAchievementRelationshipsReleasesDataInner struct for GameCenterAchievementRelationshipsReleasesDataInner
type GameCenterAchievementRelationshipsReleasesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _GameCenterAchievementRelationshipsReleasesDataInner GameCenterAchievementRelationshipsReleasesDataInner

// NewGameCenterAchievementRelationshipsReleasesDataInner instantiates a new GameCenterAchievementRelationshipsReleasesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementRelationshipsReleasesDataInner(type_ string, id string) *GameCenterAchievementRelationshipsReleasesDataInner {
	this := GameCenterAchievementRelationshipsReleasesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAchievementRelationshipsReleasesDataInnerWithDefaults instantiates a new GameCenterAchievementRelationshipsReleasesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementRelationshipsReleasesDataInnerWithDefaults() *GameCenterAchievementRelationshipsReleasesDataInner {
	this := GameCenterAchievementRelationshipsReleasesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementRelationshipsReleasesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelationshipsReleasesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementRelationshipsReleasesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAchievementRelationshipsReleasesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelationshipsReleasesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAchievementRelationshipsReleasesDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterAchievementRelationshipsReleasesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementRelationshipsReleasesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GameCenterAchievementRelationshipsReleasesDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGameCenterAchievementRelationshipsReleasesDataInner := _GameCenterAchievementRelationshipsReleasesDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementRelationshipsReleasesDataInner)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementRelationshipsReleasesDataInner(varGameCenterAchievementRelationshipsReleasesDataInner)

	return err
}

type NullableGameCenterAchievementRelationshipsReleasesDataInner struct {
	value *GameCenterAchievementRelationshipsReleasesDataInner
	isSet bool
}

func (v NullableGameCenterAchievementRelationshipsReleasesDataInner) Get() *GameCenterAchievementRelationshipsReleasesDataInner {
	return v.value
}

func (v *NullableGameCenterAchievementRelationshipsReleasesDataInner) Set(val *GameCenterAchievementRelationshipsReleasesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementRelationshipsReleasesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementRelationshipsReleasesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementRelationshipsReleasesDataInner(val *GameCenterAchievementRelationshipsReleasesDataInner) *NullableGameCenterAchievementRelationshipsReleasesDataInner {
	return &NullableGameCenterAchievementRelationshipsReleasesDataInner{value: val, isSet: true}
}

func (v NullableGameCenterAchievementRelationshipsReleasesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementRelationshipsReleasesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


