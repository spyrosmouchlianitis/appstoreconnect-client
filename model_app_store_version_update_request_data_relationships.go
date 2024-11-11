/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AppStoreVersionUpdateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionUpdateRequestDataRelationships{}

// AppStoreVersionUpdateRequestDataRelationships struct for AppStoreVersionUpdateRequestDataRelationships
type AppStoreVersionUpdateRequestDataRelationships struct {
	Build *AppStoreVersionCreateRequestDataRelationshipsBuild `json:"build,omitempty"`
	AppClipDefaultExperience *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience `json:"appClipDefaultExperience,omitempty"`
}

// NewAppStoreVersionUpdateRequestDataRelationships instantiates a new AppStoreVersionUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequestDataRelationships() *AppStoreVersionUpdateRequestDataRelationships {
	this := AppStoreVersionUpdateRequestDataRelationships{}
	return &this
}

// NewAppStoreVersionUpdateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestDataRelationshipsWithDefaults() *AppStoreVersionUpdateRequestDataRelationships {
	this := AppStoreVersionUpdateRequestDataRelationships{}
	return &this
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetBuild() AppStoreVersionCreateRequestDataRelationshipsBuild {
	if o == nil || IsNil(o.Build) {
		var ret AppStoreVersionCreateRequestDataRelationshipsBuild
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetBuildOk() (*AppStoreVersionCreateRequestDataRelationshipsBuild, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given AppStoreVersionCreateRequestDataRelationshipsBuild and assigns it to the Build field.
func (o *AppStoreVersionUpdateRequestDataRelationships) SetBuild(v AppStoreVersionCreateRequestDataRelationshipsBuild) {
	o.Build = &v
}

// GetAppClipDefaultExperience returns the AppClipDefaultExperience field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetAppClipDefaultExperience() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience
		return ret
	}
	return *o.AppClipDefaultExperience
}

// GetAppClipDefaultExperienceOk returns a tuple with the AppClipDefaultExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) GetAppClipDefaultExperienceOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience, bool) {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		return nil, false
	}
	return o.AppClipDefaultExperience, true
}

// HasAppClipDefaultExperience returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataRelationships) HasAppClipDefaultExperience() bool {
	if o != nil && !IsNil(o.AppClipDefaultExperience) {
		return true
	}

	return false
}

// SetAppClipDefaultExperience gets a reference to the given AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience and assigns it to the AppClipDefaultExperience field.
func (o *AppStoreVersionUpdateRequestDataRelationships) SetAppClipDefaultExperience(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) {
	o.AppClipDefaultExperience = &v
}

func (o AppStoreVersionUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionUpdateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !IsNil(o.AppClipDefaultExperience) {
		toSerialize["appClipDefaultExperience"] = o.AppClipDefaultExperience
	}
	return toSerialize, nil
}

type NullableAppStoreVersionUpdateRequestDataRelationships struct {
	value *AppStoreVersionUpdateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) Get() *AppStoreVersionUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) Set(val *AppStoreVersionUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequestDataRelationships(val *AppStoreVersionUpdateRequestDataRelationships) *NullableAppStoreVersionUpdateRequestDataRelationships {
	return &NullableAppStoreVersionUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


