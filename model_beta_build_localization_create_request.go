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

// checks if the BetaBuildLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationCreateRequest{}

// BetaBuildLocalizationCreateRequest struct for BetaBuildLocalizationCreateRequest
type BetaBuildLocalizationCreateRequest struct {
	Data BetaBuildLocalizationCreateRequestData `json:"data"`
}

type _BetaBuildLocalizationCreateRequest BetaBuildLocalizationCreateRequest

// NewBetaBuildLocalizationCreateRequest instantiates a new BetaBuildLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationCreateRequest(data BetaBuildLocalizationCreateRequestData) *BetaBuildLocalizationCreateRequest {
	this := BetaBuildLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaBuildLocalizationCreateRequestWithDefaults instantiates a new BetaBuildLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationCreateRequestWithDefaults() *BetaBuildLocalizationCreateRequest {
	this := BetaBuildLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaBuildLocalizationCreateRequest) GetData() BetaBuildLocalizationCreateRequestData {
	if o == nil {
		var ret BetaBuildLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequest) GetDataOk() (*BetaBuildLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaBuildLocalizationCreateRequest) SetData(v BetaBuildLocalizationCreateRequestData) {
	o.Data = v
}

func (o BetaBuildLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BetaBuildLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varBetaBuildLocalizationCreateRequest := _BetaBuildLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaBuildLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = BetaBuildLocalizationCreateRequest(varBetaBuildLocalizationCreateRequest)

	return err
}

type NullableBetaBuildLocalizationCreateRequest struct {
	value *BetaBuildLocalizationCreateRequest
	isSet bool
}

func (v NullableBetaBuildLocalizationCreateRequest) Get() *BetaBuildLocalizationCreateRequest {
	return v.value
}

func (v *NullableBetaBuildLocalizationCreateRequest) Set(val *BetaBuildLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationCreateRequest(val *BetaBuildLocalizationCreateRequest) *NullableBetaBuildLocalizationCreateRequest {
	return &NullableBetaBuildLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


