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

// checks if the GameCenterGroupAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupAttributes{}

// GameCenterGroupAttributes struct for GameCenterGroupAttributes
type GameCenterGroupAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
}

// NewGameCenterGroupAttributes instantiates a new GameCenterGroupAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupAttributes() *GameCenterGroupAttributes {
	this := GameCenterGroupAttributes{}
	return &this
}

// NewGameCenterGroupAttributesWithDefaults instantiates a new GameCenterGroupAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupAttributesWithDefaults() *GameCenterGroupAttributes {
	this := GameCenterGroupAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterGroupAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroupAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterGroupAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterGroupAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

func (o GameCenterGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	return toSerialize, nil
}

type NullableGameCenterGroupAttributes struct {
	value *GameCenterGroupAttributes
	isSet bool
}

func (v NullableGameCenterGroupAttributes) Get() *GameCenterGroupAttributes {
	return v.value
}

func (v *NullableGameCenterGroupAttributes) Set(val *GameCenterGroupAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupAttributes(val *GameCenterGroupAttributes) *NullableGameCenterGroupAttributes {
	return &NullableGameCenterGroupAttributes{value: val, isSet: true}
}

func (v NullableGameCenterGroupAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


