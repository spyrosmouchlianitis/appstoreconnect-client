/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the GameCenterLeaderboardSetLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationAttributes{}

// GameCenterLeaderboardSetLocalizationAttributes struct for GameCenterLeaderboardSetLocalizationAttributes
type GameCenterLeaderboardSetLocalizationAttributes struct {
	Locale *string `json:"locale,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGameCenterLeaderboardSetLocalizationAttributes instantiates a new GameCenterLeaderboardSetLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationAttributes() *GameCenterLeaderboardSetLocalizationAttributes {
	this := GameCenterLeaderboardSetLocalizationAttributes{}
	return &this
}

// NewGameCenterLeaderboardSetLocalizationAttributesWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationAttributesWithDefaults() *GameCenterLeaderboardSetLocalizationAttributes {
	this := GameCenterLeaderboardSetLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GameCenterLeaderboardSetLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GameCenterLeaderboardSetLocalizationAttributes) SetName(v string) {
	o.Name = &v
}

func (o GameCenterLeaderboardSetLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetLocalizationAttributes struct {
	value *GameCenterLeaderboardSetLocalizationAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationAttributes) Get() *GameCenterLeaderboardSetLocalizationAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationAttributes) Set(val *GameCenterLeaderboardSetLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationAttributes(val *GameCenterLeaderboardSetLocalizationAttributes) *NullableGameCenterLeaderboardSetLocalizationAttributes {
	return &NullableGameCenterLeaderboardSetLocalizationAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

