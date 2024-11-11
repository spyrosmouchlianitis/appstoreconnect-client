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

// checks if the AppClipAppStoreReviewDetailCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailCreateRequestDataRelationships{}

// AppClipAppStoreReviewDetailCreateRequestDataRelationships struct for AppClipAppStoreReviewDetailCreateRequestDataRelationships
type AppClipAppStoreReviewDetailCreateRequestDataRelationships struct {
	AppClipDefaultExperience AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience `json:"appClipDefaultExperience"`
}

type _AppClipAppStoreReviewDetailCreateRequestDataRelationships AppClipAppStoreReviewDetailCreateRequestDataRelationships

// NewAppClipAppStoreReviewDetailCreateRequestDataRelationships instantiates a new AppClipAppStoreReviewDetailCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailCreateRequestDataRelationships(appClipDefaultExperience AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) *AppClipAppStoreReviewDetailCreateRequestDataRelationships {
	this := AppClipAppStoreReviewDetailCreateRequestDataRelationships{}
	this.AppClipDefaultExperience = appClipDefaultExperience
	return &this
}

// NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsWithDefaults instantiates a new AppClipAppStoreReviewDetailCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailCreateRequestDataRelationshipsWithDefaults() *AppClipAppStoreReviewDetailCreateRequestDataRelationships {
	this := AppClipAppStoreReviewDetailCreateRequestDataRelationships{}
	return &this
}

// GetAppClipDefaultExperience returns the AppClipDefaultExperience field value
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationships) GetAppClipDefaultExperience() AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience {
	if o == nil {
		var ret AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience
		return ret
	}

	return o.AppClipDefaultExperience
}

// GetAppClipDefaultExperienceOk returns a tuple with the AppClipDefaultExperience field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationships) GetAppClipDefaultExperienceOk() (*AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppClipDefaultExperience, true
}

// SetAppClipDefaultExperience sets field value
func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationships) SetAppClipDefaultExperience(v AppClipAppStoreReviewDetailCreateRequestDataRelationshipsAppClipDefaultExperience) {
	o.AppClipDefaultExperience = v
}

func (o AppClipAppStoreReviewDetailCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appClipDefaultExperience"] = o.AppClipDefaultExperience
	return toSerialize, nil
}

func (o *AppClipAppStoreReviewDetailCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appClipDefaultExperience",
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

	varAppClipAppStoreReviewDetailCreateRequestDataRelationships := _AppClipAppStoreReviewDetailCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipAppStoreReviewDetailCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = AppClipAppStoreReviewDetailCreateRequestDataRelationships(varAppClipAppStoreReviewDetailCreateRequestDataRelationships)

	return err
}

type NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships struct {
	value *AppClipAppStoreReviewDetailCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) Get() *AppClipAppStoreReviewDetailCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) Set(val *AppClipAppStoreReviewDetailCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailCreateRequestDataRelationships(val *AppClipAppStoreReviewDetailCreateRequestDataRelationships) *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships {
	return &NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

