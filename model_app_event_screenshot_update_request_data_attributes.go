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

// checks if the AppEventScreenshotUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotUpdateRequestDataAttributes{}

// AppEventScreenshotUpdateRequestDataAttributes struct for AppEventScreenshotUpdateRequestDataAttributes
type AppEventScreenshotUpdateRequestDataAttributes struct {
	Uploaded *bool `json:"uploaded,omitempty"`
}

// NewAppEventScreenshotUpdateRequestDataAttributes instantiates a new AppEventScreenshotUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotUpdateRequestDataAttributes() *AppEventScreenshotUpdateRequestDataAttributes {
	this := AppEventScreenshotUpdateRequestDataAttributes{}
	return &this
}

// NewAppEventScreenshotUpdateRequestDataAttributesWithDefaults instantiates a new AppEventScreenshotUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotUpdateRequestDataAttributesWithDefaults() *AppEventScreenshotUpdateRequestDataAttributes {
	this := AppEventScreenshotUpdateRequestDataAttributes{}
	return &this
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise.
func (o *AppEventScreenshotUpdateRequestDataAttributes) GetUploaded() bool {
	if o == nil || IsNil(o.Uploaded) {
		var ret bool
		return ret
	}
	return *o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotUpdateRequestDataAttributes) GetUploadedOk() (*bool, bool) {
	if o == nil || IsNil(o.Uploaded) {
		return nil, false
	}
	return o.Uploaded, true
}

// HasUploaded returns a boolean if a field has been set.
func (o *AppEventScreenshotUpdateRequestDataAttributes) HasUploaded() bool {
	if o != nil && !IsNil(o.Uploaded) {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given bool and assigns it to the Uploaded field.
func (o *AppEventScreenshotUpdateRequestDataAttributes) SetUploaded(v bool) {
	o.Uploaded = &v
}

func (o AppEventScreenshotUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uploaded) {
		toSerialize["uploaded"] = o.Uploaded
	}
	return toSerialize, nil
}

type NullableAppEventScreenshotUpdateRequestDataAttributes struct {
	value *AppEventScreenshotUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventScreenshotUpdateRequestDataAttributes) Get() *AppEventScreenshotUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventScreenshotUpdateRequestDataAttributes) Set(val *AppEventScreenshotUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotUpdateRequestDataAttributes(val *AppEventScreenshotUpdateRequestDataAttributes) *NullableAppEventScreenshotUpdateRequestDataAttributes {
	return &NullableAppEventScreenshotUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventScreenshotUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


