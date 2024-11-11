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

// checks if the AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail{}

// AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail struct for AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail
type AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail struct {
	Data AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData `json:"data"`
}

type _AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail

// NewAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail(data AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData) *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail{}
	this.Data = data
	return &this
}

// NewAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetailWithDefaults instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetailWithDefaults() *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) GetData() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData {
	if o == nil {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) GetDataOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) SetData(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData) {
	o.Data = v
}

func (o AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail := _AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail)

	if err != nil {
		return err
	}

	*o = AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail(varAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail)

	return err
}

type NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail struct {
	value *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail
	isSet bool
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) Get() *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) Set(val *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail(val *AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) *NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	return &NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


