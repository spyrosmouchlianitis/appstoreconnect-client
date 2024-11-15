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

// checks if the SubscriptionGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupCreateRequest{}

// SubscriptionGroupCreateRequest struct for SubscriptionGroupCreateRequest
type SubscriptionGroupCreateRequest struct {
	Data SubscriptionGroupCreateRequestData `json:"data"`
}

type _SubscriptionGroupCreateRequest SubscriptionGroupCreateRequest

// NewSubscriptionGroupCreateRequest instantiates a new SubscriptionGroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupCreateRequest(data SubscriptionGroupCreateRequestData) *SubscriptionGroupCreateRequest {
	this := SubscriptionGroupCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionGroupCreateRequestWithDefaults instantiates a new SubscriptionGroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupCreateRequestWithDefaults() *SubscriptionGroupCreateRequest {
	this := SubscriptionGroupCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupCreateRequest) GetData() SubscriptionGroupCreateRequestData {
	if o == nil {
		var ret SubscriptionGroupCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupCreateRequest) GetDataOk() (*SubscriptionGroupCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupCreateRequest) SetData(v SubscriptionGroupCreateRequestData) {
	o.Data = v
}

func (o SubscriptionGroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionGroupCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionGroupCreateRequest := _SubscriptionGroupCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionGroupCreateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionGroupCreateRequest(varSubscriptionGroupCreateRequest)

	return err
}

type NullableSubscriptionGroupCreateRequest struct {
	value *SubscriptionGroupCreateRequest
	isSet bool
}

func (v NullableSubscriptionGroupCreateRequest) Get() *SubscriptionGroupCreateRequest {
	return v.value
}

func (v *NullableSubscriptionGroupCreateRequest) Set(val *SubscriptionGroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupCreateRequest(val *SubscriptionGroupCreateRequest) *NullableSubscriptionGroupCreateRequest {
	return &NullableSubscriptionGroupCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


