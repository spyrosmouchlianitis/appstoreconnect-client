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

// checks if the SubscriptionImageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionImageUpdateRequest{}

// SubscriptionImageUpdateRequest struct for SubscriptionImageUpdateRequest
type SubscriptionImageUpdateRequest struct {
	Data SubscriptionImageUpdateRequestData `json:"data"`
}

type _SubscriptionImageUpdateRequest SubscriptionImageUpdateRequest

// NewSubscriptionImageUpdateRequest instantiates a new SubscriptionImageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionImageUpdateRequest(data SubscriptionImageUpdateRequestData) *SubscriptionImageUpdateRequest {
	this := SubscriptionImageUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionImageUpdateRequestWithDefaults instantiates a new SubscriptionImageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionImageUpdateRequestWithDefaults() *SubscriptionImageUpdateRequest {
	this := SubscriptionImageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionImageUpdateRequest) GetData() SubscriptionImageUpdateRequestData {
	if o == nil {
		var ret SubscriptionImageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionImageUpdateRequest) GetDataOk() (*SubscriptionImageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionImageUpdateRequest) SetData(v SubscriptionImageUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionImageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionImageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionImageUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionImageUpdateRequest := _SubscriptionImageUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionImageUpdateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionImageUpdateRequest(varSubscriptionImageUpdateRequest)

	return err
}

type NullableSubscriptionImageUpdateRequest struct {
	value *SubscriptionImageUpdateRequest
	isSet bool
}

func (v NullableSubscriptionImageUpdateRequest) Get() *SubscriptionImageUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionImageUpdateRequest) Set(val *SubscriptionImageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionImageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionImageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionImageUpdateRequest(val *SubscriptionImageUpdateRequest) *NullableSubscriptionImageUpdateRequest {
	return &NullableSubscriptionImageUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionImageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionImageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


