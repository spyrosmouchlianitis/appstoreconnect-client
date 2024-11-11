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

// checks if the AppClipAdvancedExperienceImageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImageCreateRequest{}

// AppClipAdvancedExperienceImageCreateRequest struct for AppClipAdvancedExperienceImageCreateRequest
type AppClipAdvancedExperienceImageCreateRequest struct {
	Data AppClipAdvancedExperienceImageCreateRequestData `json:"data"`
}

type _AppClipAdvancedExperienceImageCreateRequest AppClipAdvancedExperienceImageCreateRequest

// NewAppClipAdvancedExperienceImageCreateRequest instantiates a new AppClipAdvancedExperienceImageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImageCreateRequest(data AppClipAdvancedExperienceImageCreateRequestData) *AppClipAdvancedExperienceImageCreateRequest {
	this := AppClipAdvancedExperienceImageCreateRequest{}
	this.Data = data
	return &this
}

// NewAppClipAdvancedExperienceImageCreateRequestWithDefaults instantiates a new AppClipAdvancedExperienceImageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageCreateRequestWithDefaults() *AppClipAdvancedExperienceImageCreateRequest {
	this := AppClipAdvancedExperienceImageCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceImageCreateRequest) GetData() AppClipAdvancedExperienceImageCreateRequestData {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageCreateRequest) GetDataOk() (*AppClipAdvancedExperienceImageCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceImageCreateRequest) SetData(v AppClipAdvancedExperienceImageCreateRequestData) {
	o.Data = v
}

func (o AppClipAdvancedExperienceImageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImageCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppClipAdvancedExperienceImageCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipAdvancedExperienceImageCreateRequest := _AppClipAdvancedExperienceImageCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipAdvancedExperienceImageCreateRequest)

	if err != nil {
		return err
	}

	*o = AppClipAdvancedExperienceImageCreateRequest(varAppClipAdvancedExperienceImageCreateRequest)

	return err
}

type NullableAppClipAdvancedExperienceImageCreateRequest struct {
	value *AppClipAdvancedExperienceImageCreateRequest
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImageCreateRequest) Get() *AppClipAdvancedExperienceImageCreateRequest {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequest) Set(val *AppClipAdvancedExperienceImageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImageCreateRequest(val *AppClipAdvancedExperienceImageCreateRequest) *NullableAppClipAdvancedExperienceImageCreateRequest {
	return &NullableAppClipAdvancedExperienceImageCreateRequest{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


