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

// checks if the ErrorSourceParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorSourceParameter{}

// ErrorSourceParameter struct for ErrorSourceParameter
type ErrorSourceParameter struct {
	Parameter string `json:"parameter"`
}

type _ErrorSourceParameter ErrorSourceParameter

// NewErrorSourceParameter instantiates a new ErrorSourceParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSourceParameter(parameter string) *ErrorSourceParameter {
	this := ErrorSourceParameter{}
	this.Parameter = parameter
	return &this
}

// NewErrorSourceParameterWithDefaults instantiates a new ErrorSourceParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourceParameterWithDefaults() *ErrorSourceParameter {
	this := ErrorSourceParameter{}
	return &this
}

// GetParameter returns the Parameter field value
func (o *ErrorSourceParameter) GetParameter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *ErrorSourceParameter) GetParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *ErrorSourceParameter) SetParameter(v string) {
	o.Parameter = v
}

func (o ErrorSourceParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorSourceParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameter"] = o.Parameter
	return toSerialize, nil
}

func (o *ErrorSourceParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameter",
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

	varErrorSourceParameter := _ErrorSourceParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorSourceParameter)

	if err != nil {
		return err
	}

	*o = ErrorSourceParameter(varErrorSourceParameter)

	return err
}

type NullableErrorSourceParameter struct {
	value *ErrorSourceParameter
	isSet bool
}

func (v NullableErrorSourceParameter) Get() *ErrorSourceParameter {
	return v.value
}

func (v *NullableErrorSourceParameter) Set(val *ErrorSourceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSourceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSourceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSourceParameter(val *ErrorSourceParameter) *NullableErrorSourceParameter {
	return &NullableErrorSourceParameter{value: val, isSet: true}
}

func (v NullableErrorSourceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSourceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


