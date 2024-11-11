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

// checks if the AppStoreVersionLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationCreateRequest{}

// AppStoreVersionLocalizationCreateRequest struct for AppStoreVersionLocalizationCreateRequest
type AppStoreVersionLocalizationCreateRequest struct {
	Data AppStoreVersionLocalizationCreateRequestData `json:"data"`
}

type _AppStoreVersionLocalizationCreateRequest AppStoreVersionLocalizationCreateRequest

// NewAppStoreVersionLocalizationCreateRequest instantiates a new AppStoreVersionLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationCreateRequest(data AppStoreVersionLocalizationCreateRequestData) *AppStoreVersionLocalizationCreateRequest {
	this := AppStoreVersionLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionLocalizationCreateRequestWithDefaults instantiates a new AppStoreVersionLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationCreateRequestWithDefaults() *AppStoreVersionLocalizationCreateRequest {
	this := AppStoreVersionLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationCreateRequest) GetData() AppStoreVersionLocalizationCreateRequestData {
	if o == nil {
		var ret AppStoreVersionLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequest) GetDataOk() (*AppStoreVersionLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationCreateRequest) SetData(v AppStoreVersionLocalizationCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreVersionLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionLocalizationCreateRequest := _AppStoreVersionLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = AppStoreVersionLocalizationCreateRequest(varAppStoreVersionLocalizationCreateRequest)

	return err
}

type NullableAppStoreVersionLocalizationCreateRequest struct {
	value *AppStoreVersionLocalizationCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionLocalizationCreateRequest) Get() *AppStoreVersionLocalizationCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationCreateRequest) Set(val *AppStoreVersionLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationCreateRequest(val *AppStoreVersionLocalizationCreateRequest) *NullableAppStoreVersionLocalizationCreateRequest {
	return &NullableAppStoreVersionLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


