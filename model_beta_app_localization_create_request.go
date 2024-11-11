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

// checks if the BetaAppLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationCreateRequest{}

// BetaAppLocalizationCreateRequest struct for BetaAppLocalizationCreateRequest
type BetaAppLocalizationCreateRequest struct {
	Data BetaAppLocalizationCreateRequestData `json:"data"`
}

type _BetaAppLocalizationCreateRequest BetaAppLocalizationCreateRequest

// NewBetaAppLocalizationCreateRequest instantiates a new BetaAppLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationCreateRequest(data BetaAppLocalizationCreateRequestData) *BetaAppLocalizationCreateRequest {
	this := BetaAppLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppLocalizationCreateRequestWithDefaults instantiates a new BetaAppLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationCreateRequestWithDefaults() *BetaAppLocalizationCreateRequest {
	this := BetaAppLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppLocalizationCreateRequest) GetData() BetaAppLocalizationCreateRequestData {
	if o == nil {
		var ret BetaAppLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationCreateRequest) GetDataOk() (*BetaAppLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppLocalizationCreateRequest) SetData(v BetaAppLocalizationCreateRequestData) {
	o.Data = v
}

func (o BetaAppLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BetaAppLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBetaAppLocalizationCreateRequest := _BetaAppLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaAppLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = BetaAppLocalizationCreateRequest(varBetaAppLocalizationCreateRequest)

	return err
}

type NullableBetaAppLocalizationCreateRequest struct {
	value *BetaAppLocalizationCreateRequest
	isSet bool
}

func (v NullableBetaAppLocalizationCreateRequest) Get() *BetaAppLocalizationCreateRequest {
	return v.value
}

func (v *NullableBetaAppLocalizationCreateRequest) Set(val *BetaAppLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationCreateRequest(val *BetaAppLocalizationCreateRequest) *NullableBetaAppLocalizationCreateRequest {
	return &NullableBetaAppLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

