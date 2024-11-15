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

// checks if the AppInfoLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationCreateRequest{}

// AppInfoLocalizationCreateRequest struct for AppInfoLocalizationCreateRequest
type AppInfoLocalizationCreateRequest struct {
	Data AppInfoLocalizationCreateRequestData `json:"data"`
}

type _AppInfoLocalizationCreateRequest AppInfoLocalizationCreateRequest

// NewAppInfoLocalizationCreateRequest instantiates a new AppInfoLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationCreateRequest(data AppInfoLocalizationCreateRequestData) *AppInfoLocalizationCreateRequest {
	this := AppInfoLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewAppInfoLocalizationCreateRequestWithDefaults instantiates a new AppInfoLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationCreateRequestWithDefaults() *AppInfoLocalizationCreateRequest {
	this := AppInfoLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppInfoLocalizationCreateRequest) GetData() AppInfoLocalizationCreateRequestData {
	if o == nil {
		var ret AppInfoLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationCreateRequest) GetDataOk() (*AppInfoLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppInfoLocalizationCreateRequest) SetData(v AppInfoLocalizationCreateRequestData) {
	o.Data = v
}

func (o AppInfoLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppInfoLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppInfoLocalizationCreateRequest := _AppInfoLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppInfoLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = AppInfoLocalizationCreateRequest(varAppInfoLocalizationCreateRequest)

	return err
}

type NullableAppInfoLocalizationCreateRequest struct {
	value *AppInfoLocalizationCreateRequest
	isSet bool
}

func (v NullableAppInfoLocalizationCreateRequest) Get() *AppInfoLocalizationCreateRequest {
	return v.value
}

func (v *NullableAppInfoLocalizationCreateRequest) Set(val *AppInfoLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationCreateRequest(val *AppInfoLocalizationCreateRequest) *NullableAppInfoLocalizationCreateRequest {
	return &NullableAppInfoLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


