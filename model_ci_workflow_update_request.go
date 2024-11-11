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

// checks if the CiWorkflowUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowUpdateRequest{}

// CiWorkflowUpdateRequest struct for CiWorkflowUpdateRequest
type CiWorkflowUpdateRequest struct {
	Data CiWorkflowUpdateRequestData `json:"data"`
}

type _CiWorkflowUpdateRequest CiWorkflowUpdateRequest

// NewCiWorkflowUpdateRequest instantiates a new CiWorkflowUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowUpdateRequest(data CiWorkflowUpdateRequestData) *CiWorkflowUpdateRequest {
	this := CiWorkflowUpdateRequest{}
	this.Data = data
	return &this
}

// NewCiWorkflowUpdateRequestWithDefaults instantiates a new CiWorkflowUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowUpdateRequestWithDefaults() *CiWorkflowUpdateRequest {
	this := CiWorkflowUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *CiWorkflowUpdateRequest) GetData() CiWorkflowUpdateRequestData {
	if o == nil {
		var ret CiWorkflowUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowUpdateRequest) GetDataOk() (*CiWorkflowUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiWorkflowUpdateRequest) SetData(v CiWorkflowUpdateRequestData) {
	o.Data = v
}

func (o CiWorkflowUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CiWorkflowUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCiWorkflowUpdateRequest := _CiWorkflowUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiWorkflowUpdateRequest)

	if err != nil {
		return err
	}

	*o = CiWorkflowUpdateRequest(varCiWorkflowUpdateRequest)

	return err
}

type NullableCiWorkflowUpdateRequest struct {
	value *CiWorkflowUpdateRequest
	isSet bool
}

func (v NullableCiWorkflowUpdateRequest) Get() *CiWorkflowUpdateRequest {
	return v.value
}

func (v *NullableCiWorkflowUpdateRequest) Set(val *CiWorkflowUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowUpdateRequest(val *CiWorkflowUpdateRequest) *NullableCiWorkflowUpdateRequest {
	return &NullableCiWorkflowUpdateRequest{value: val, isSet: true}
}

func (v NullableCiWorkflowUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


