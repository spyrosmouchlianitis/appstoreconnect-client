/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the AppClipAppStoreReviewDetailRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailRelationships{}

// AppClipAppStoreReviewDetailRelationships struct for AppClipAppStoreReviewDetailRelationships
type AppClipAppStoreReviewDetailRelationships struct {
	AppClipDefaultExperience *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience `json:"appClipDefaultExperience,omitempty"`
}

// NewAppClipAppStoreReviewDetailRelationships instantiates a new AppClipAppStoreReviewDetailRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailRelationships() *AppClipAppStoreReviewDetailRelationships {
	this := AppClipAppStoreReviewDetailRelationships{}
	return &this
}

// NewAppClipAppStoreReviewDetailRelationshipsWithDefaults instantiates a new AppClipAppStoreReviewDetailRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailRelationshipsWithDefaults() *AppClipAppStoreReviewDetailRelationships {
	this := AppClipAppStoreReviewDetailRelationships{}
	return &this
}

// GetAppClipDefaultExperience returns the AppClipDefaultExperience field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailRelationships) GetAppClipDefaultExperience() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience
		return ret
	}
	return *o.AppClipDefaultExperience
}

// GetAppClipDefaultExperienceOk returns a tuple with the AppClipDefaultExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailRelationships) GetAppClipDefaultExperienceOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience, bool) {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		return nil, false
	}
	return o.AppClipDefaultExperience, true
}

// HasAppClipDefaultExperience returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailRelationships) HasAppClipDefaultExperience() bool {
	if o != nil && !IsNil(o.AppClipDefaultExperience) {
		return true
	}

	return false
}

// SetAppClipDefaultExperience gets a reference to the given AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience and assigns it to the AppClipDefaultExperience field.
func (o *AppClipAppStoreReviewDetailRelationships) SetAppClipDefaultExperience(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) {
	o.AppClipDefaultExperience = &v
}

func (o AppClipAppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppClipDefaultExperience) {
		toSerialize["appClipDefaultExperience"] = o.AppClipDefaultExperience
	}
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailRelationships struct {
	value *AppClipAppStoreReviewDetailRelationships
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailRelationships) Get() *AppClipAppStoreReviewDetailRelationships {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailRelationships) Set(val *AppClipAppStoreReviewDetailRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailRelationships(val *AppClipAppStoreReviewDetailRelationships) *NullableAppClipAppStoreReviewDetailRelationships {
	return &NullableAppClipAppStoreReviewDetailRelationships{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


