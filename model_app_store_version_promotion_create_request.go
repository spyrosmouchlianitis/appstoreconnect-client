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

// checks if the AppStoreVersionPromotionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPromotionCreateRequest{}

// AppStoreVersionPromotionCreateRequest struct for AppStoreVersionPromotionCreateRequest
type AppStoreVersionPromotionCreateRequest struct {
	Data AppStoreVersionPromotionCreateRequestData `json:"data"`
}

type _AppStoreVersionPromotionCreateRequest AppStoreVersionPromotionCreateRequest

// NewAppStoreVersionPromotionCreateRequest instantiates a new AppStoreVersionPromotionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPromotionCreateRequest(data AppStoreVersionPromotionCreateRequestData) *AppStoreVersionPromotionCreateRequest {
	this := AppStoreVersionPromotionCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionPromotionCreateRequestWithDefaults instantiates a new AppStoreVersionPromotionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPromotionCreateRequestWithDefaults() *AppStoreVersionPromotionCreateRequest {
	this := AppStoreVersionPromotionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionPromotionCreateRequest) GetData() AppStoreVersionPromotionCreateRequestData {
	if o == nil {
		var ret AppStoreVersionPromotionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotionCreateRequest) GetDataOk() (*AppStoreVersionPromotionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionPromotionCreateRequest) SetData(v AppStoreVersionPromotionCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionPromotionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPromotionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreVersionPromotionCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionPromotionCreateRequest := _AppStoreVersionPromotionCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionPromotionCreateRequest)

	if err != nil {
		return err
	}

	*o = AppStoreVersionPromotionCreateRequest(varAppStoreVersionPromotionCreateRequest)

	return err
}

type NullableAppStoreVersionPromotionCreateRequest struct {
	value *AppStoreVersionPromotionCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionPromotionCreateRequest) Get() *AppStoreVersionPromotionCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionPromotionCreateRequest) Set(val *AppStoreVersionPromotionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPromotionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPromotionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPromotionCreateRequest(val *AppStoreVersionPromotionCreateRequest) *NullableAppStoreVersionPromotionCreateRequest {
	return &NullableAppStoreVersionPromotionCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionPromotionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPromotionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

