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

// checks if the AppStoreVersionExperimentTreatmentLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationAttributes{}

// AppStoreVersionExperimentTreatmentLocalizationAttributes struct for AppStoreVersionExperimentTreatmentLocalizationAttributes
type AppStoreVersionExperimentTreatmentLocalizationAttributes struct {
	Locale *string `json:"locale,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentLocalizationAttributes instantiates a new AppStoreVersionExperimentTreatmentLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationAttributes() *AppStoreVersionExperimentTreatmentLocalizationAttributes {
	this := AppStoreVersionExperimentTreatmentLocalizationAttributes{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationAttributesWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationAttributesWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationAttributes {
	this := AppStoreVersionExperimentTreatmentLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppStoreVersionExperimentTreatmentLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

func (o AppStoreVersionExperimentTreatmentLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentLocalizationAttributes struct {
	value *AppStoreVersionExperimentTreatmentLocalizationAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) Get() *AppStoreVersionExperimentTreatmentLocalizationAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) Set(val *AppStoreVersionExperimentTreatmentLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationAttributes(val *AppStoreVersionExperimentTreatmentLocalizationAttributes) *NullableAppStoreVersionExperimentTreatmentLocalizationAttributes {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


