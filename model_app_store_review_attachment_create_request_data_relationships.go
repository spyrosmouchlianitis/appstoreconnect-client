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

// checks if the AppStoreReviewAttachmentCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentCreateRequestDataRelationships{}

// AppStoreReviewAttachmentCreateRequestDataRelationships struct for AppStoreReviewAttachmentCreateRequestDataRelationships
type AppStoreReviewAttachmentCreateRequestDataRelationships struct {
	AppStoreReviewDetail AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail `json:"appStoreReviewDetail"`
}

type _AppStoreReviewAttachmentCreateRequestDataRelationships AppStoreReviewAttachmentCreateRequestDataRelationships

// NewAppStoreReviewAttachmentCreateRequestDataRelationships instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentCreateRequestDataRelationships(appStoreReviewDetail AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) *AppStoreReviewAttachmentCreateRequestDataRelationships {
	this := AppStoreReviewAttachmentCreateRequestDataRelationships{}
	this.AppStoreReviewDetail = appStoreReviewDetail
	return &this
}

// NewAppStoreReviewAttachmentCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreReviewAttachmentCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentCreateRequestDataRelationshipsWithDefaults() *AppStoreReviewAttachmentCreateRequestDataRelationships {
	this := AppStoreReviewAttachmentCreateRequestDataRelationships{}
	return &this
}

// GetAppStoreReviewDetail returns the AppStoreReviewDetail field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) GetAppStoreReviewDetail() AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail {
	if o == nil {
		var ret AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail
		return ret
	}

	return o.AppStoreReviewDetail
}

// GetAppStoreReviewDetailOk returns a tuple with the AppStoreReviewDetail field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) GetAppStoreReviewDetailOk() (*AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppStoreReviewDetail, true
}

// SetAppStoreReviewDetail sets field value
func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) SetAppStoreReviewDetail(v AppStoreReviewAttachmentCreateRequestDataRelationshipsAppStoreReviewDetail) {
	o.AppStoreReviewDetail = v
}

func (o AppStoreReviewAttachmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appStoreReviewDetail"] = o.AppStoreReviewDetail
	return toSerialize, nil
}

func (o *AppStoreReviewAttachmentCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appStoreReviewDetail",
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

	varAppStoreReviewAttachmentCreateRequestDataRelationships := _AppStoreReviewAttachmentCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreReviewAttachmentCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = AppStoreReviewAttachmentCreateRequestDataRelationships(varAppStoreReviewAttachmentCreateRequestDataRelationships)

	return err
}

type NullableAppStoreReviewAttachmentCreateRequestDataRelationships struct {
	value *AppStoreReviewAttachmentCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Get() *AppStoreReviewAttachmentCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Set(val *AppStoreReviewAttachmentCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentCreateRequestDataRelationships(val *AppStoreReviewAttachmentCreateRequestDataRelationships) *NullableAppStoreReviewAttachmentCreateRequestDataRelationships {
	return &NullableAppStoreReviewAttachmentCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


