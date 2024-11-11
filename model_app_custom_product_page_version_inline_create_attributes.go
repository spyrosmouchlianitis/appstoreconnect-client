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

// checks if the AppCustomProductPageVersionInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionInlineCreateAttributes{}

// AppCustomProductPageVersionInlineCreateAttributes struct for AppCustomProductPageVersionInlineCreateAttributes
type AppCustomProductPageVersionInlineCreateAttributes struct {
	DeepLink *string `json:"deepLink,omitempty"`
}

// NewAppCustomProductPageVersionInlineCreateAttributes instantiates a new AppCustomProductPageVersionInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionInlineCreateAttributes() *AppCustomProductPageVersionInlineCreateAttributes {
	this := AppCustomProductPageVersionInlineCreateAttributes{}
	return &this
}

// NewAppCustomProductPageVersionInlineCreateAttributesWithDefaults instantiates a new AppCustomProductPageVersionInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionInlineCreateAttributesWithDefaults() *AppCustomProductPageVersionInlineCreateAttributes {
	this := AppCustomProductPageVersionInlineCreateAttributes{}
	return &this
}

// GetDeepLink returns the DeepLink field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionInlineCreateAttributes) GetDeepLink() string {
	if o == nil || IsNil(o.DeepLink) {
		var ret string
		return ret
	}
	return *o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionInlineCreateAttributes) GetDeepLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DeepLink) {
		return nil, false
	}
	return o.DeepLink, true
}

// HasDeepLink returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionInlineCreateAttributes) HasDeepLink() bool {
	if o != nil && !IsNil(o.DeepLink) {
		return true
	}

	return false
}

// SetDeepLink gets a reference to the given string and assigns it to the DeepLink field.
func (o *AppCustomProductPageVersionInlineCreateAttributes) SetDeepLink(v string) {
	o.DeepLink = &v
}

func (o AppCustomProductPageVersionInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeepLink) {
		toSerialize["deepLink"] = o.DeepLink
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionInlineCreateAttributes struct {
	value *AppCustomProductPageVersionInlineCreateAttributes
	isSet bool
}

func (v NullableAppCustomProductPageVersionInlineCreateAttributes) Get() *AppCustomProductPageVersionInlineCreateAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageVersionInlineCreateAttributes) Set(val *AppCustomProductPageVersionInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionInlineCreateAttributes(val *AppCustomProductPageVersionInlineCreateAttributes) *NullableAppCustomProductPageVersionInlineCreateAttributes {
	return &NullableAppCustomProductPageVersionInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


