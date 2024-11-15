/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the GameCenterAppVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionAttributes{}

// GameCenterAppVersionAttributes struct for GameCenterAppVersionAttributes
type GameCenterAppVersionAttributes struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGameCenterAppVersionAttributes instantiates a new GameCenterAppVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionAttributes() *GameCenterAppVersionAttributes {
	this := GameCenterAppVersionAttributes{}
	return &this
}

// NewGameCenterAppVersionAttributesWithDefaults instantiates a new GameCenterAppVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionAttributesWithDefaults() *GameCenterAppVersionAttributes {
	this := GameCenterAppVersionAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GameCenterAppVersionAttributes) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GameCenterAppVersionAttributes) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GameCenterAppVersionAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GameCenterAppVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGameCenterAppVersionAttributes struct {
	value *GameCenterAppVersionAttributes
	isSet bool
}

func (v NullableGameCenterAppVersionAttributes) Get() *GameCenterAppVersionAttributes {
	return v.value
}

func (v *NullableGameCenterAppVersionAttributes) Set(val *GameCenterAppVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionAttributes(val *GameCenterAppVersionAttributes) *NullableGameCenterAppVersionAttributes {
	return &NullableGameCenterAppVersionAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


