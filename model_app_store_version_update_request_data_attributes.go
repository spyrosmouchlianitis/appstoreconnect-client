/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"time"
)

// checks if the AppStoreVersionUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionUpdateRequestDataAttributes{}

// AppStoreVersionUpdateRequestDataAttributes struct for AppStoreVersionUpdateRequestDataAttributes
type AppStoreVersionUpdateRequestDataAttributes struct {
	VersionString *string `json:"versionString,omitempty"`
	Copyright *string `json:"copyright,omitempty"`
	ReviewType *string `json:"reviewType,omitempty"`
	ReleaseType *string `json:"releaseType,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
	Downloadable *bool `json:"downloadable,omitempty"`
}

// NewAppStoreVersionUpdateRequestDataAttributes instantiates a new AppStoreVersionUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionUpdateRequestDataAttributes() *AppStoreVersionUpdateRequestDataAttributes {
	this := AppStoreVersionUpdateRequestDataAttributes{}
	return &this
}

// NewAppStoreVersionUpdateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionUpdateRequestDataAttributesWithDefaults() *AppStoreVersionUpdateRequestDataAttributes {
	this := AppStoreVersionUpdateRequestDataAttributes{}
	return &this
}

// GetVersionString returns the VersionString field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetVersionString() string {
	if o == nil || IsNil(o.VersionString) {
		var ret string
		return ret
	}
	return *o.VersionString
}

// GetVersionStringOk returns a tuple with the VersionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetVersionStringOk() (*string, bool) {
	if o == nil || IsNil(o.VersionString) {
		return nil, false
	}
	return o.VersionString, true
}

// HasVersionString returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasVersionString() bool {
	if o != nil && !IsNil(o.VersionString) {
		return true
	}

	return false
}

// SetVersionString gets a reference to the given string and assigns it to the VersionString field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetVersionString(v string) {
	o.VersionString = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetCopyright(v string) {
	o.Copyright = &v
}

// GetReviewType returns the ReviewType field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetReviewType() string {
	if o == nil || IsNil(o.ReviewType) {
		var ret string
		return ret
	}
	return *o.ReviewType
}

// GetReviewTypeOk returns a tuple with the ReviewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetReviewTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewType) {
		return nil, false
	}
	return o.ReviewType, true
}

// HasReviewType returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasReviewType() bool {
	if o != nil && !IsNil(o.ReviewType) {
		return true
	}

	return false
}

// SetReviewType gets a reference to the given string and assigns it to the ReviewType field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetReviewType(v string) {
	o.ReviewType = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetEarliestReleaseDate returns the EarliestReleaseDate field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetEarliestReleaseDate() time.Time {
	if o == nil || IsNil(o.EarliestReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestReleaseDate
}

// GetEarliestReleaseDateOk returns a tuple with the EarliestReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetEarliestReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestReleaseDate) {
		return nil, false
	}
	return o.EarliestReleaseDate, true
}

// HasEarliestReleaseDate returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasEarliestReleaseDate() bool {
	if o != nil && !IsNil(o.EarliestReleaseDate) {
		return true
	}

	return false
}

// SetEarliestReleaseDate gets a reference to the given time.Time and assigns it to the EarliestReleaseDate field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetEarliestReleaseDate(v time.Time) {
	o.EarliestReleaseDate = &v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetDownloadable() bool {
	if o == nil || IsNil(o.Downloadable) {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) GetDownloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Downloadable) {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *AppStoreVersionUpdateRequestDataAttributes) HasDownloadable() bool {
	if o != nil && !IsNil(o.Downloadable) {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *AppStoreVersionUpdateRequestDataAttributes) SetDownloadable(v bool) {
	o.Downloadable = &v
}

func (o AppStoreVersionUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionString) {
		toSerialize["versionString"] = o.VersionString
	}
	if !IsNil(o.Copyright) {
		toSerialize["copyright"] = o.Copyright
	}
	if !IsNil(o.ReviewType) {
		toSerialize["reviewType"] = o.ReviewType
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.EarliestReleaseDate) {
		toSerialize["earliestReleaseDate"] = o.EarliestReleaseDate
	}
	if !IsNil(o.Downloadable) {
		toSerialize["downloadable"] = o.Downloadable
	}
	return toSerialize, nil
}

type NullableAppStoreVersionUpdateRequestDataAttributes struct {
	value *AppStoreVersionUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionUpdateRequestDataAttributes) Get() *AppStoreVersionUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionUpdateRequestDataAttributes) Set(val *AppStoreVersionUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionUpdateRequestDataAttributes(val *AppStoreVersionUpdateRequestDataAttributes) *NullableAppStoreVersionUpdateRequestDataAttributes {
	return &NullableAppStoreVersionUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


