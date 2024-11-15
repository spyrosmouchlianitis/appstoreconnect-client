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

// checks if the AppCustomProductPageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageUpdateRequest{}

// AppCustomProductPageUpdateRequest struct for AppCustomProductPageUpdateRequest
type AppCustomProductPageUpdateRequest struct {
	Data AppCustomProductPageUpdateRequestData `json:"data"`
}

type _AppCustomProductPageUpdateRequest AppCustomProductPageUpdateRequest

// NewAppCustomProductPageUpdateRequest instantiates a new AppCustomProductPageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageUpdateRequest(data AppCustomProductPageUpdateRequestData) *AppCustomProductPageUpdateRequest {
	this := AppCustomProductPageUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppCustomProductPageUpdateRequestWithDefaults instantiates a new AppCustomProductPageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageUpdateRequestWithDefaults() *AppCustomProductPageUpdateRequest {
	this := AppCustomProductPageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageUpdateRequest) GetData() AppCustomProductPageUpdateRequestData {
	if o == nil {
		var ret AppCustomProductPageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageUpdateRequest) GetDataOk() (*AppCustomProductPageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageUpdateRequest) SetData(v AppCustomProductPageUpdateRequestData) {
	o.Data = v
}

func (o AppCustomProductPageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppCustomProductPageUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppCustomProductPageUpdateRequest := _AppCustomProductPageUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppCustomProductPageUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppCustomProductPageUpdateRequest(varAppCustomProductPageUpdateRequest)

	return err
}

type NullableAppCustomProductPageUpdateRequest struct {
	value *AppCustomProductPageUpdateRequest
	isSet bool
}

func (v NullableAppCustomProductPageUpdateRequest) Get() *AppCustomProductPageUpdateRequest {
	return v.value
}

func (v *NullableAppCustomProductPageUpdateRequest) Set(val *AppCustomProductPageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageUpdateRequest(val *AppCustomProductPageUpdateRequest) *NullableAppCustomProductPageUpdateRequest {
	return &NullableAppCustomProductPageUpdateRequest{value: val, isSet: true}
}

func (v NullableAppCustomProductPageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


