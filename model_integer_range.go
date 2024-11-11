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

// checks if the IntegerRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegerRange{}

// IntegerRange struct for IntegerRange
type IntegerRange struct {
	Minimum *int32 `json:"minimum,omitempty"`
	Maximum *int32 `json:"maximum,omitempty"`
}

// NewIntegerRange instantiates a new IntegerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerRange() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// NewIntegerRangeWithDefaults instantiates a new IntegerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerRangeWithDefaults() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *IntegerRange) GetMinimum() int32 {
	if o == nil || IsNil(o.Minimum) {
		var ret int32
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetMinimumOk() (*int32, bool) {
	if o == nil || IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *IntegerRange) HasMinimum() bool {
	if o != nil && !IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given int32 and assigns it to the Minimum field.
func (o *IntegerRange) SetMinimum(v int32) {
	o.Minimum = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *IntegerRange) GetMaximum() int32 {
	if o == nil || IsNil(o.Maximum) {
		var ret int32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetMaximumOk() (*int32, bool) {
	if o == nil || IsNil(o.Maximum) {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *IntegerRange) HasMaximum() bool {
	if o != nil && !IsNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given int32 and assigns it to the Maximum field.
func (o *IntegerRange) SetMaximum(v int32) {
	o.Maximum = &v
}

func (o IntegerRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegerRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !IsNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	return toSerialize, nil
}

type NullableIntegerRange struct {
	value *IntegerRange
	isSet bool
}

func (v NullableIntegerRange) Get() *IntegerRange {
	return v.value
}

func (v *NullableIntegerRange) Set(val *IntegerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerRange(val *IntegerRange) *NullableIntegerRange {
	return &NullableIntegerRange{value: val, isSet: true}
}

func (v NullableIntegerRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

