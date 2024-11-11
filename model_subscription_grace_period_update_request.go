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

// checks if the SubscriptionGracePeriodUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGracePeriodUpdateRequest{}

// SubscriptionGracePeriodUpdateRequest struct for SubscriptionGracePeriodUpdateRequest
type SubscriptionGracePeriodUpdateRequest struct {
	Data SubscriptionGracePeriodUpdateRequestData `json:"data"`
}

type _SubscriptionGracePeriodUpdateRequest SubscriptionGracePeriodUpdateRequest

// NewSubscriptionGracePeriodUpdateRequest instantiates a new SubscriptionGracePeriodUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGracePeriodUpdateRequest(data SubscriptionGracePeriodUpdateRequestData) *SubscriptionGracePeriodUpdateRequest {
	this := SubscriptionGracePeriodUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionGracePeriodUpdateRequestWithDefaults instantiates a new SubscriptionGracePeriodUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGracePeriodUpdateRequestWithDefaults() *SubscriptionGracePeriodUpdateRequest {
	this := SubscriptionGracePeriodUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGracePeriodUpdateRequest) GetData() SubscriptionGracePeriodUpdateRequestData {
	if o == nil {
		var ret SubscriptionGracePeriodUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriodUpdateRequest) GetDataOk() (*SubscriptionGracePeriodUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGracePeriodUpdateRequest) SetData(v SubscriptionGracePeriodUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionGracePeriodUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGracePeriodUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionGracePeriodUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionGracePeriodUpdateRequest := _SubscriptionGracePeriodUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionGracePeriodUpdateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionGracePeriodUpdateRequest(varSubscriptionGracePeriodUpdateRequest)

	return err
}

type NullableSubscriptionGracePeriodUpdateRequest struct {
	value *SubscriptionGracePeriodUpdateRequest
	isSet bool
}

func (v NullableSubscriptionGracePeriodUpdateRequest) Get() *SubscriptionGracePeriodUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionGracePeriodUpdateRequest) Set(val *SubscriptionGracePeriodUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGracePeriodUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGracePeriodUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGracePeriodUpdateRequest(val *SubscriptionGracePeriodUpdateRequest) *NullableSubscriptionGracePeriodUpdateRequest {
	return &NullableSubscriptionGracePeriodUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionGracePeriodUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGracePeriodUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


