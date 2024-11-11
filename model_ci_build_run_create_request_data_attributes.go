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

// checks if the CiBuildRunCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestDataAttributes{}

// CiBuildRunCreateRequestDataAttributes struct for CiBuildRunCreateRequestDataAttributes
type CiBuildRunCreateRequestDataAttributes struct {
	Clean *bool `json:"clean,omitempty"`
}

// NewCiBuildRunCreateRequestDataAttributes instantiates a new CiBuildRunCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestDataAttributes() *CiBuildRunCreateRequestDataAttributes {
	this := CiBuildRunCreateRequestDataAttributes{}
	return &this
}

// NewCiBuildRunCreateRequestDataAttributesWithDefaults instantiates a new CiBuildRunCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataAttributesWithDefaults() *CiBuildRunCreateRequestDataAttributes {
	this := CiBuildRunCreateRequestDataAttributes{}
	return &this
}

// GetClean returns the Clean field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataAttributes) GetClean() bool {
	if o == nil || IsNil(o.Clean) {
		var ret bool
		return ret
	}
	return *o.Clean
}

// GetCleanOk returns a tuple with the Clean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataAttributes) GetCleanOk() (*bool, bool) {
	if o == nil || IsNil(o.Clean) {
		return nil, false
	}
	return o.Clean, true
}

// HasClean returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataAttributes) HasClean() bool {
	if o != nil && !IsNil(o.Clean) {
		return true
	}

	return false
}

// SetClean gets a reference to the given bool and assigns it to the Clean field.
func (o *CiBuildRunCreateRequestDataAttributes) SetClean(v bool) {
	o.Clean = &v
}

func (o CiBuildRunCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clean) {
		toSerialize["clean"] = o.Clean
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestDataAttributes struct {
	value *CiBuildRunCreateRequestDataAttributes
	isSet bool
}

func (v NullableCiBuildRunCreateRequestDataAttributes) Get() *CiBuildRunCreateRequestDataAttributes {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestDataAttributes) Set(val *CiBuildRunCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestDataAttributes(val *CiBuildRunCreateRequestDataAttributes) *NullableCiBuildRunCreateRequestDataAttributes {
	return &NullableCiBuildRunCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


