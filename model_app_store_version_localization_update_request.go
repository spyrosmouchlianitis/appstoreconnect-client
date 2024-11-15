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

// checks if the AppStoreVersionLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationUpdateRequest{}

// AppStoreVersionLocalizationUpdateRequest struct for AppStoreVersionLocalizationUpdateRequest
type AppStoreVersionLocalizationUpdateRequest struct {
	Data AppStoreVersionLocalizationUpdateRequestData `json:"data"`
}

type _AppStoreVersionLocalizationUpdateRequest AppStoreVersionLocalizationUpdateRequest

// NewAppStoreVersionLocalizationUpdateRequest instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationUpdateRequest(data AppStoreVersionLocalizationUpdateRequestData) *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionLocalizationUpdateRequestWithDefaults instantiates a new AppStoreVersionLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationUpdateRequestWithDefaults() *AppStoreVersionLocalizationUpdateRequest {
	this := AppStoreVersionLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationUpdateRequest) GetData() AppStoreVersionLocalizationUpdateRequestData {
	if o == nil {
		var ret AppStoreVersionLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationUpdateRequest) GetDataOk() (*AppStoreVersionLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationUpdateRequest) SetData(v AppStoreVersionLocalizationUpdateRequestData) {
	o.Data = v
}

func (o AppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreVersionLocalizationUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionLocalizationUpdateRequest := _AppStoreVersionLocalizationUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionLocalizationUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppStoreVersionLocalizationUpdateRequest(varAppStoreVersionLocalizationUpdateRequest)

	return err
}

type NullableAppStoreVersionLocalizationUpdateRequest struct {
	value *AppStoreVersionLocalizationUpdateRequest
	isSet bool
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) Get() *AppStoreVersionLocalizationUpdateRequest {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Set(val *AppStoreVersionLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationUpdateRequest(val *AppStoreVersionLocalizationUpdateRequest) *NullableAppStoreVersionLocalizationUpdateRequest {
	return &NullableAppStoreVersionLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


