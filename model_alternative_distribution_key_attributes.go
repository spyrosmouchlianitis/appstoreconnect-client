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

// checks if the AlternativeDistributionKeyAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionKeyAttributes{}

// AlternativeDistributionKeyAttributes struct for AlternativeDistributionKeyAttributes
type AlternativeDistributionKeyAttributes struct {
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewAlternativeDistributionKeyAttributes instantiates a new AlternativeDistributionKeyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionKeyAttributes() *AlternativeDistributionKeyAttributes {
	this := AlternativeDistributionKeyAttributes{}
	return &this
}

// NewAlternativeDistributionKeyAttributesWithDefaults instantiates a new AlternativeDistributionKeyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionKeyAttributesWithDefaults() *AlternativeDistributionKeyAttributes {
	this := AlternativeDistributionKeyAttributes{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *AlternativeDistributionKeyAttributes) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionKeyAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *AlternativeDistributionKeyAttributes) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *AlternativeDistributionKeyAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o AlternativeDistributionKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionKeyAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableAlternativeDistributionKeyAttributes struct {
	value *AlternativeDistributionKeyAttributes
	isSet bool
}

func (v NullableAlternativeDistributionKeyAttributes) Get() *AlternativeDistributionKeyAttributes {
	return v.value
}

func (v *NullableAlternativeDistributionKeyAttributes) Set(val *AlternativeDistributionKeyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionKeyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionKeyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionKeyAttributes(val *AlternativeDistributionKeyAttributes) *NullableAlternativeDistributionKeyAttributes {
	return &NullableAlternativeDistributionKeyAttributes{value: val, isSet: true}
}

func (v NullableAlternativeDistributionKeyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionKeyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


