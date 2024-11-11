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

// checks if the BetaAppClipInvocationLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationCreateRequest{}

// BetaAppClipInvocationLocalizationCreateRequest struct for BetaAppClipInvocationLocalizationCreateRequest
type BetaAppClipInvocationLocalizationCreateRequest struct {
	Data BetaAppClipInvocationLocalizationCreateRequestData `json:"data"`
}

type _BetaAppClipInvocationLocalizationCreateRequest BetaAppClipInvocationLocalizationCreateRequest

// NewBetaAppClipInvocationLocalizationCreateRequest instantiates a new BetaAppClipInvocationLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationCreateRequest(data BetaAppClipInvocationLocalizationCreateRequestData) *BetaAppClipInvocationLocalizationCreateRequest {
	this := BetaAppClipInvocationLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppClipInvocationLocalizationCreateRequestWithDefaults instantiates a new BetaAppClipInvocationLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationCreateRequestWithDefaults() *BetaAppClipInvocationLocalizationCreateRequest {
	this := BetaAppClipInvocationLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationLocalizationCreateRequest) GetData() BetaAppClipInvocationLocalizationCreateRequestData {
	if o == nil {
		var ret BetaAppClipInvocationLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationCreateRequest) GetDataOk() (*BetaAppClipInvocationLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationLocalizationCreateRequest) SetData(v BetaAppClipInvocationLocalizationCreateRequestData) {
	o.Data = v
}

func (o BetaAppClipInvocationLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BetaAppClipInvocationLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBetaAppClipInvocationLocalizationCreateRequest := _BetaAppClipInvocationLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppClipInvocationLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = BetaAppClipInvocationLocalizationCreateRequest(varBetaAppClipInvocationLocalizationCreateRequest)

	return err
}

type NullableBetaAppClipInvocationLocalizationCreateRequest struct {
	value *BetaAppClipInvocationLocalizationCreateRequest
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationCreateRequest) Get() *BetaAppClipInvocationLocalizationCreateRequest {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationCreateRequest) Set(val *BetaAppClipInvocationLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationCreateRequest(val *BetaAppClipInvocationLocalizationCreateRequest) *NullableBetaAppClipInvocationLocalizationCreateRequest {
	return &NullableBetaAppClipInvocationLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


