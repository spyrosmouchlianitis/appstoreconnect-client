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

// checks if the AppClipDefaultExperienceLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalizationUpdateRequest{}

// AppClipDefaultExperienceLocalizationUpdateRequest struct for AppClipDefaultExperienceLocalizationUpdateRequest
type AppClipDefaultExperienceLocalizationUpdateRequest struct {
	Data AppClipDefaultExperienceLocalizationUpdateRequestData `json:"data"`
}

type _AppClipDefaultExperienceLocalizationUpdateRequest AppClipDefaultExperienceLocalizationUpdateRequest

// NewAppClipDefaultExperienceLocalizationUpdateRequest instantiates a new AppClipDefaultExperienceLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalizationUpdateRequest(data AppClipDefaultExperienceLocalizationUpdateRequestData) *AppClipDefaultExperienceLocalizationUpdateRequest {
	this := AppClipDefaultExperienceLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppClipDefaultExperienceLocalizationUpdateRequestWithDefaults instantiates a new AppClipDefaultExperienceLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationUpdateRequestWithDefaults() *AppClipDefaultExperienceLocalizationUpdateRequest {
	this := AppClipDefaultExperienceLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDefaultExperienceLocalizationUpdateRequest) GetData() AppClipDefaultExperienceLocalizationUpdateRequestData {
	if o == nil {
		var ret AppClipDefaultExperienceLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalizationUpdateRequest) GetDataOk() (*AppClipDefaultExperienceLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipDefaultExperienceLocalizationUpdateRequest) SetData(v AppClipDefaultExperienceLocalizationUpdateRequestData) {
	o.Data = v
}

func (o AppClipDefaultExperienceLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppClipDefaultExperienceLocalizationUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipDefaultExperienceLocalizationUpdateRequest := _AppClipDefaultExperienceLocalizationUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipDefaultExperienceLocalizationUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppClipDefaultExperienceLocalizationUpdateRequest(varAppClipDefaultExperienceLocalizationUpdateRequest)

	return err
}

type NullableAppClipDefaultExperienceLocalizationUpdateRequest struct {
	value *AppClipDefaultExperienceLocalizationUpdateRequest
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequest) Get() *AppClipDefaultExperienceLocalizationUpdateRequest {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequest) Set(val *AppClipDefaultExperienceLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalizationUpdateRequest(val *AppClipDefaultExperienceLocalizationUpdateRequest) *NullableAppClipDefaultExperienceLocalizationUpdateRequest {
	return &NullableAppClipDefaultExperienceLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


