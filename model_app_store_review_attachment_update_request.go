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

// checks if the AppStoreReviewAttachmentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentUpdateRequest{}

// AppStoreReviewAttachmentUpdateRequest struct for AppStoreReviewAttachmentUpdateRequest
type AppStoreReviewAttachmentUpdateRequest struct {
	Data AppStoreReviewAttachmentUpdateRequestData `json:"data"`
}

type _AppStoreReviewAttachmentUpdateRequest AppStoreReviewAttachmentUpdateRequest

// NewAppStoreReviewAttachmentUpdateRequest instantiates a new AppStoreReviewAttachmentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentUpdateRequest(data AppStoreReviewAttachmentUpdateRequestData) *AppStoreReviewAttachmentUpdateRequest {
	this := AppStoreReviewAttachmentUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreReviewAttachmentUpdateRequestWithDefaults instantiates a new AppStoreReviewAttachmentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentUpdateRequestWithDefaults() *AppStoreReviewAttachmentUpdateRequest {
	this := AppStoreReviewAttachmentUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewAttachmentUpdateRequest) GetData() AppStoreReviewAttachmentUpdateRequestData {
	if o == nil {
		var ret AppStoreReviewAttachmentUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentUpdateRequest) GetDataOk() (*AppStoreReviewAttachmentUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewAttachmentUpdateRequest) SetData(v AppStoreReviewAttachmentUpdateRequestData) {
	o.Data = v
}

func (o AppStoreReviewAttachmentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreReviewAttachmentUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreReviewAttachmentUpdateRequest := _AppStoreReviewAttachmentUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreReviewAttachmentUpdateRequest)

	if err != nil {
		return err
	}

	*o = AppStoreReviewAttachmentUpdateRequest(varAppStoreReviewAttachmentUpdateRequest)

	return err
}

type NullableAppStoreReviewAttachmentUpdateRequest struct {
	value *AppStoreReviewAttachmentUpdateRequest
	isSet bool
}

func (v NullableAppStoreReviewAttachmentUpdateRequest) Get() *AppStoreReviewAttachmentUpdateRequest {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentUpdateRequest) Set(val *AppStoreReviewAttachmentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentUpdateRequest(val *AppStoreReviewAttachmentUpdateRequest) *NullableAppStoreReviewAttachmentUpdateRequest {
	return &NullableAppStoreReviewAttachmentUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


