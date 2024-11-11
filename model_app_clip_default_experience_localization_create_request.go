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

// checks if the AppClipDefaultExperienceLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationCreateRequest{}

// AppClipDefaultExperienceLocalizationCreateRequest struct for AppClipDefaultExperienceLocalizationCreateRequest
type AppClipDefaultExperienceLocalizationCreateRequest struct {
	Data AppClipDefaultExperienceLocalizationCreateRequestData `json:"data"`
}

type _AppClipDefaultExperienceLocalizationCreateRequest AppClipDefaultExperienceLocalizationCreateRequest

// NewAppClipDefaultExperienceLocalizationCreateRequest instantiates a new AppClipDefaultExperienceLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationCreateRequest(data AppClipDefaultExperienceLocalizationCreateRequestData) *AppClipDefaultExperienceLocalizationCreateRequest {
	this := AppClipDefaultExperienceLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewAppClipDefaultExperienceLocalizationCreateRequestWithDefaults instantiates a new AppClipDefaultExperienceLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationCreateRequestWithDefaults() *AppClipDefaultExperienceLocalizationCreateRequest {
	this := AppClipDefaultExperienceLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDefaultExperienceLocalizationCreateRequest) GetData() AppClipDefaultExperienceLocalizationCreateRequestData {
	if o == nil {
		var ret AppClipDefaultExperienceLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationCreateRequest) GetDataOk() (*AppClipDefaultExperienceLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipDefaultExperienceLocalizationCreateRequest) SetData(v AppClipDefaultExperienceLocalizationCreateRequestData) {
	o.Data = v
}

func (o AppClipDefaultExperienceLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppClipDefaultExperienceLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipDefaultExperienceLocalizationCreateRequest := _AppClipDefaultExperienceLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipDefaultExperienceLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = AppClipDefaultExperienceLocalizationCreateRequest(varAppClipDefaultExperienceLocalizationCreateRequest)

	return err
}

type NullableAppClipDefaultExperienceLocalizationCreateRequest struct {
	value *AppClipDefaultExperienceLocalizationCreateRequest
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequest) Get() *AppClipDefaultExperienceLocalizationCreateRequest {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequest) Set(val *AppClipDefaultExperienceLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationCreateRequest(val *AppClipDefaultExperienceLocalizationCreateRequest) *NullableAppClipDefaultExperienceLocalizationCreateRequest {
	return &NullableAppClipDefaultExperienceLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


