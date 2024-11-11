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

// checks if the AppPreviewCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewCreateRequest{}

// AppPreviewCreateRequest struct for AppPreviewCreateRequest
type AppPreviewCreateRequest struct {
	Data AppPreviewCreateRequestData `json:"data"`
}

type _AppPreviewCreateRequest AppPreviewCreateRequest

// NewAppPreviewCreateRequest instantiates a new AppPreviewCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewCreateRequest(data AppPreviewCreateRequestData) *AppPreviewCreateRequest {
	this := AppPreviewCreateRequest{}
	this.Data = data
	return &this
}

// NewAppPreviewCreateRequestWithDefaults instantiates a new AppPreviewCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewCreateRequestWithDefaults() *AppPreviewCreateRequest {
	this := AppPreviewCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewCreateRequest) GetData() AppPreviewCreateRequestData {
	if o == nil {
		var ret AppPreviewCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewCreateRequest) GetDataOk() (*AppPreviewCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewCreateRequest) SetData(v AppPreviewCreateRequestData) {
	o.Data = v
}

func (o AppPreviewCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppPreviewCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppPreviewCreateRequest := _AppPreviewCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppPreviewCreateRequest)

	if err != nil {
		return err
	}

	*o = AppPreviewCreateRequest(varAppPreviewCreateRequest)

	return err
}

type NullableAppPreviewCreateRequest struct {
	value *AppPreviewCreateRequest
	isSet bool
}

func (v NullableAppPreviewCreateRequest) Get() *AppPreviewCreateRequest {
	return v.value
}

func (v *NullableAppPreviewCreateRequest) Set(val *AppPreviewCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewCreateRequest(val *AppPreviewCreateRequest) *NullableAppPreviewCreateRequest {
	return &NullableAppPreviewCreateRequest{value: val, isSet: true}
}

func (v NullableAppPreviewCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


