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

// checks if the AlternativeDistributionKeyCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionKeyCreateRequest{}

// AlternativeDistributionKeyCreateRequest struct for AlternativeDistributionKeyCreateRequest
type AlternativeDistributionKeyCreateRequest struct {
	Data AlternativeDistributionKeyCreateRequestData `json:"data"`
}

type _AlternativeDistributionKeyCreateRequest AlternativeDistributionKeyCreateRequest

// NewAlternativeDistributionKeyCreateRequest instantiates a new AlternativeDistributionKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionKeyCreateRequest(data AlternativeDistributionKeyCreateRequestData) *AlternativeDistributionKeyCreateRequest {
	this := AlternativeDistributionKeyCreateRequest{}
	this.Data = data
	return &this
}

// NewAlternativeDistributionKeyCreateRequestWithDefaults instantiates a new AlternativeDistributionKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionKeyCreateRequestWithDefaults() *AlternativeDistributionKeyCreateRequest {
	this := AlternativeDistributionKeyCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AlternativeDistributionKeyCreateRequest) GetData() AlternativeDistributionKeyCreateRequestData {
	if o == nil {
		var ret AlternativeDistributionKeyCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionKeyCreateRequest) GetDataOk() (*AlternativeDistributionKeyCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AlternativeDistributionKeyCreateRequest) SetData(v AlternativeDistributionKeyCreateRequestData) {
	o.Data = v
}

func (o AlternativeDistributionKeyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionKeyCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AlternativeDistributionKeyCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAlternativeDistributionKeyCreateRequest := _AlternativeDistributionKeyCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlternativeDistributionKeyCreateRequest)

	if err != nil {
		return err
	}

	*o = AlternativeDistributionKeyCreateRequest(varAlternativeDistributionKeyCreateRequest)

	return err
}

type NullableAlternativeDistributionKeyCreateRequest struct {
	value *AlternativeDistributionKeyCreateRequest
	isSet bool
}

func (v NullableAlternativeDistributionKeyCreateRequest) Get() *AlternativeDistributionKeyCreateRequest {
	return v.value
}

func (v *NullableAlternativeDistributionKeyCreateRequest) Set(val *AlternativeDistributionKeyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionKeyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionKeyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionKeyCreateRequest(val *AlternativeDistributionKeyCreateRequest) *NullableAlternativeDistributionKeyCreateRequest {
	return &NullableAlternativeDistributionKeyCreateRequest{value: val, isSet: true}
}

func (v NullableAlternativeDistributionKeyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionKeyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


