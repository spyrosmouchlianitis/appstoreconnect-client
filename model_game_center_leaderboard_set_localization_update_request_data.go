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

// checks if the GameCenterLeaderboardSetLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationUpdateRequestData{}

// GameCenterLeaderboardSetLocalizationUpdateRequestData struct for GameCenterLeaderboardSetLocalizationUpdateRequestData
type GameCenterLeaderboardSetLocalizationUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BundleIdUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _GameCenterLeaderboardSetLocalizationUpdateRequestData GameCenterLeaderboardSetLocalizationUpdateRequestData

// NewGameCenterLeaderboardSetLocalizationUpdateRequestData instantiates a new GameCenterLeaderboardSetLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationUpdateRequestData(type_ string, id string) *GameCenterLeaderboardSetLocalizationUpdateRequestData {
	this := GameCenterLeaderboardSetLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardSetLocalizationUpdateRequestDataWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationUpdateRequestDataWithDefaults() *GameCenterLeaderboardSetLocalizationUpdateRequestData {
	this := GameCenterLeaderboardSetLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetAttributes() BundleIdUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BundleIdUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) GetAttributesOk() (*BundleIdUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BundleIdUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) SetAttributes(v BundleIdUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterLeaderboardSetLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetLocalizationUpdateRequestData := _GameCenterLeaderboardSetLocalizationUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationUpdateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationUpdateRequestData(varGameCenterLeaderboardSetLocalizationUpdateRequestData)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationUpdateRequestData struct {
	value *GameCenterLeaderboardSetLocalizationUpdateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) Get() *GameCenterLeaderboardSetLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) Set(val *GameCenterLeaderboardSetLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationUpdateRequestData(val *GameCenterLeaderboardSetLocalizationUpdateRequestData) *NullableGameCenterLeaderboardSetLocalizationUpdateRequestData {
	return &NullableGameCenterLeaderboardSetLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


