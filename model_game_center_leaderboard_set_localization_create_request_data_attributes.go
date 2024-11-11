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

// checks if the GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes{}

// GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes struct for GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes
type GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
	Name string `json:"name"`
}

type _GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes

// NewGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes(locale string, name string) *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes {
	this := GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	this.Name = name
	return &this
}

// NewGameCenterLeaderboardSetLocalizationCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes {
	this := GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetName returns the Name field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locale",
		"name",
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

	varGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes := _GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes(varGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes struct {
	value *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) Get() *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) Set(val *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes(val *GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes {
	return &NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


