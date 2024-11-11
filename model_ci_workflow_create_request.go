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

// checks if the CiWorkflowCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowCreateRequest{}

// CiWorkflowCreateRequest struct for CiWorkflowCreateRequest
type CiWorkflowCreateRequest struct {
	Data CiWorkflowCreateRequestData `json:"data"`
}

type _CiWorkflowCreateRequest CiWorkflowCreateRequest

// NewCiWorkflowCreateRequest instantiates a new CiWorkflowCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowCreateRequest(data CiWorkflowCreateRequestData) *CiWorkflowCreateRequest {
	this := CiWorkflowCreateRequest{}
	this.Data = data
	return &this
}

// NewCiWorkflowCreateRequestWithDefaults instantiates a new CiWorkflowCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowCreateRequestWithDefaults() *CiWorkflowCreateRequest {
	this := CiWorkflowCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *CiWorkflowCreateRequest) GetData() CiWorkflowCreateRequestData {
	if o == nil {
		var ret CiWorkflowCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequest) GetDataOk() (*CiWorkflowCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiWorkflowCreateRequest) SetData(v CiWorkflowCreateRequestData) {
	o.Data = v
}

func (o CiWorkflowCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CiWorkflowCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCiWorkflowCreateRequest := _CiWorkflowCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiWorkflowCreateRequest)

	if err != nil {
		return err
	}

	*o = CiWorkflowCreateRequest(varCiWorkflowCreateRequest)

	return err
}

type NullableCiWorkflowCreateRequest struct {
	value *CiWorkflowCreateRequest
	isSet bool
}

func (v NullableCiWorkflowCreateRequest) Get() *CiWorkflowCreateRequest {
	return v.value
}

func (v *NullableCiWorkflowCreateRequest) Set(val *CiWorkflowCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowCreateRequest(val *CiWorkflowCreateRequest) *NullableCiWorkflowCreateRequest {
	return &NullableCiWorkflowCreateRequest{value: val, isSet: true}
}

func (v NullableCiWorkflowCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

