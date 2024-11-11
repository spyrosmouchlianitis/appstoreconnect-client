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

// checks if the AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes{}

// AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes struct for AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes
type AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	AppIconName *string `json:"appIconName,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes instantiates a new AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes() *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentUpdateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentUpdateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetAppIconName returns the AppIconName field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) GetAppIconName() string {
	if o == nil || IsNil(o.AppIconName) {
		var ret string
		return ret
	}
	return *o.AppIconName
}

// GetAppIconNameOk returns a tuple with the AppIconName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) GetAppIconNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppIconName) {
		return nil, false
	}
	return o.AppIconName, true
}

// HasAppIconName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) HasAppIconName() bool {
	if o != nil && !IsNil(o.AppIconName) {
		return true
	}

	return false
}

// SetAppIconName gets a reference to the given string and assigns it to the AppIconName field.
func (o *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) SetAppIconName(v string) {
	o.AppIconName = &v
}

func (o AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AppIconName) {
		toSerialize["appIconName"] = o.AppIconName
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes struct {
	value *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) Get() *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) Set(val *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes(val *AppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) *NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes {
	return &NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


