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

// checks if the SandboxTesterV2UpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTesterV2UpdateRequest{}

// SandboxTesterV2UpdateRequest struct for SandboxTesterV2UpdateRequest
type SandboxTesterV2UpdateRequest struct {
	Data SandboxTesterV2UpdateRequestData `json:"data"`
}

type _SandboxTesterV2UpdateRequest SandboxTesterV2UpdateRequest

// NewSandboxTesterV2UpdateRequest instantiates a new SandboxTesterV2UpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTesterV2UpdateRequest(data SandboxTesterV2UpdateRequestData) *SandboxTesterV2UpdateRequest {
	this := SandboxTesterV2UpdateRequest{}
	this.Data = data
	return &this
}

// NewSandboxTesterV2UpdateRequestWithDefaults instantiates a new SandboxTesterV2UpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTesterV2UpdateRequestWithDefaults() *SandboxTesterV2UpdateRequest {
	this := SandboxTesterV2UpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SandboxTesterV2UpdateRequest) GetData() SandboxTesterV2UpdateRequestData {
	if o == nil {
		var ret SandboxTesterV2UpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2UpdateRequest) GetDataOk() (*SandboxTesterV2UpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SandboxTesterV2UpdateRequest) SetData(v SandboxTesterV2UpdateRequestData) {
	o.Data = v
}

func (o SandboxTesterV2UpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTesterV2UpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SandboxTesterV2UpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSandboxTesterV2UpdateRequest := _SandboxTesterV2UpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxTesterV2UpdateRequest)

	if err != nil {
		return err
	}

	*o = SandboxTesterV2UpdateRequest(varSandboxTesterV2UpdateRequest)

	return err
}

type NullableSandboxTesterV2UpdateRequest struct {
	value *SandboxTesterV2UpdateRequest
	isSet bool
}

func (v NullableSandboxTesterV2UpdateRequest) Get() *SandboxTesterV2UpdateRequest {
	return v.value
}

func (v *NullableSandboxTesterV2UpdateRequest) Set(val *SandboxTesterV2UpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTesterV2UpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTesterV2UpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTesterV2UpdateRequest(val *SandboxTesterV2UpdateRequest) *NullableSandboxTesterV2UpdateRequest {
	return &NullableSandboxTesterV2UpdateRequest{value: val, isSet: true}
}

func (v NullableSandboxTesterV2UpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTesterV2UpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

