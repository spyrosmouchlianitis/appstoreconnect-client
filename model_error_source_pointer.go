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

// checks if the ErrorSourcePointer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorSourcePointer{}

// ErrorSourcePointer struct for ErrorSourcePointer
type ErrorSourcePointer struct {
	Pointer string `json:"pointer"`
}

type _ErrorSourcePointer ErrorSourcePointer

// NewErrorSourcePointer instantiates a new ErrorSourcePointer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSourcePointer(pointer string) *ErrorSourcePointer {
	this := ErrorSourcePointer{}
	this.Pointer = pointer
	return &this
}

// NewErrorSourcePointerWithDefaults instantiates a new ErrorSourcePointer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourcePointerWithDefaults() *ErrorSourcePointer {
	this := ErrorSourcePointer{}
	return &this
}

// GetPointer returns the Pointer field value
func (o *ErrorSourcePointer) GetPointer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value
// and a boolean to check if the value has been set.
func (o *ErrorSourcePointer) GetPointerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pointer, true
}

// SetPointer sets field value
func (o *ErrorSourcePointer) SetPointer(v string) {
	o.Pointer = v
}

func (o ErrorSourcePointer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorSourcePointer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pointer"] = o.Pointer
	return toSerialize, nil
}

func (o *ErrorSourcePointer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pointer",
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

	varErrorSourcePointer := _ErrorSourcePointer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorSourcePointer)

	if err != nil {
		return err
	}

	*o = ErrorSourcePointer(varErrorSourcePointer)

	return err
}

type NullableErrorSourcePointer struct {
	value *ErrorSourcePointer
	isSet bool
}

func (v NullableErrorSourcePointer) Get() *ErrorSourcePointer {
	return v.value
}

func (v *NullableErrorSourcePointer) Set(val *ErrorSourcePointer) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSourcePointer) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSourcePointer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSourcePointer(val *ErrorSourcePointer) *NullableErrorSourcePointer {
	return &NullableErrorSourcePointer{value: val, isSet: true}
}

func (v NullableErrorSourcePointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSourcePointer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


