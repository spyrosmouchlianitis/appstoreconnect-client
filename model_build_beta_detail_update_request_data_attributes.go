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

// checks if the BuildBetaDetailUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaDetailUpdateRequestDataAttributes{}

// BuildBetaDetailUpdateRequestDataAttributes struct for BuildBetaDetailUpdateRequestDataAttributes
type BuildBetaDetailUpdateRequestDataAttributes struct {
	AutoNotifyEnabled *bool `json:"autoNotifyEnabled,omitempty"`
}

// NewBuildBetaDetailUpdateRequestDataAttributes instantiates a new BuildBetaDetailUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailUpdateRequestDataAttributes() *BuildBetaDetailUpdateRequestDataAttributes {
	this := BuildBetaDetailUpdateRequestDataAttributes{}
	return &this
}

// NewBuildBetaDetailUpdateRequestDataAttributesWithDefaults instantiates a new BuildBetaDetailUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailUpdateRequestDataAttributesWithDefaults() *BuildBetaDetailUpdateRequestDataAttributes {
	this := BuildBetaDetailUpdateRequestDataAttributes{}
	return &this
}

// GetAutoNotifyEnabled returns the AutoNotifyEnabled field value if set, zero value otherwise.
func (o *BuildBetaDetailUpdateRequestDataAttributes) GetAutoNotifyEnabled() bool {
	if o == nil || IsNil(o.AutoNotifyEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoNotifyEnabled
}

// GetAutoNotifyEnabledOk returns a tuple with the AutoNotifyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailUpdateRequestDataAttributes) GetAutoNotifyEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoNotifyEnabled) {
		return nil, false
	}
	return o.AutoNotifyEnabled, true
}

// HasAutoNotifyEnabled returns a boolean if a field has been set.
func (o *BuildBetaDetailUpdateRequestDataAttributes) HasAutoNotifyEnabled() bool {
	if o != nil && !IsNil(o.AutoNotifyEnabled) {
		return true
	}

	return false
}

// SetAutoNotifyEnabled gets a reference to the given bool and assigns it to the AutoNotifyEnabled field.
func (o *BuildBetaDetailUpdateRequestDataAttributes) SetAutoNotifyEnabled(v bool) {
	o.AutoNotifyEnabled = &v
}

func (o BuildBetaDetailUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaDetailUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoNotifyEnabled) {
		toSerialize["autoNotifyEnabled"] = o.AutoNotifyEnabled
	}
	return toSerialize, nil
}

type NullableBuildBetaDetailUpdateRequestDataAttributes struct {
	value *BuildBetaDetailUpdateRequestDataAttributes
	isSet bool
}

func (v NullableBuildBetaDetailUpdateRequestDataAttributes) Get() *BuildBetaDetailUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableBuildBetaDetailUpdateRequestDataAttributes) Set(val *BuildBetaDetailUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailUpdateRequestDataAttributes(val *BuildBetaDetailUpdateRequestDataAttributes) *NullableBuildBetaDetailUpdateRequestDataAttributes {
	return &NullableBuildBetaDetailUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBuildBetaDetailUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


