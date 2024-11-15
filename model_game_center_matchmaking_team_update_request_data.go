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

// checks if the GameCenterMatchmakingTeamUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamUpdateRequestData{}

// GameCenterMatchmakingTeamUpdateRequestData struct for GameCenterMatchmakingTeamUpdateRequestData
type GameCenterMatchmakingTeamUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterMatchmakingRuleSetUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _GameCenterMatchmakingTeamUpdateRequestData GameCenterMatchmakingTeamUpdateRequestData

// NewGameCenterMatchmakingTeamUpdateRequestData instantiates a new GameCenterMatchmakingTeamUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamUpdateRequestData(type_ string, id string) *GameCenterMatchmakingTeamUpdateRequestData {
	this := GameCenterMatchmakingTeamUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterMatchmakingTeamUpdateRequestDataWithDefaults instantiates a new GameCenterMatchmakingTeamUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamUpdateRequestDataWithDefaults() *GameCenterMatchmakingTeamUpdateRequestData {
	this := GameCenterMatchmakingTeamUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingTeamUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterMatchmakingTeamUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetAttributes() GameCenterMatchmakingRuleSetUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterMatchmakingRuleSetUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamUpdateRequestData) GetAttributesOk() (*GameCenterMatchmakingRuleSetUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTeamUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterMatchmakingRuleSetUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterMatchmakingTeamUpdateRequestData) SetAttributes(v GameCenterMatchmakingRuleSetUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterMatchmakingTeamUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTeamUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingTeamUpdateRequestData := _GameCenterMatchmakingTeamUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTeamUpdateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTeamUpdateRequestData(varGameCenterMatchmakingTeamUpdateRequestData)

	return err
}

type NullableGameCenterMatchmakingTeamUpdateRequestData struct {
	value *GameCenterMatchmakingTeamUpdateRequestData
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamUpdateRequestData) Get() *GameCenterMatchmakingTeamUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequestData) Set(val *GameCenterMatchmakingTeamUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamUpdateRequestData(val *GameCenterMatchmakingTeamUpdateRequestData) *NullableGameCenterMatchmakingTeamUpdateRequestData {
	return &NullableGameCenterMatchmakingTeamUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


