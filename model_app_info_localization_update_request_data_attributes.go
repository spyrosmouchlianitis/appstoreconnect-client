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

// checks if the AppInfoLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationUpdateRequestDataAttributes{}

// AppInfoLocalizationUpdateRequestDataAttributes struct for AppInfoLocalizationUpdateRequestDataAttributes
type AppInfoLocalizationUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	PrivacyPolicyUrl *string `json:"privacyPolicyUrl,omitempty"`
	PrivacyChoicesUrl *string `json:"privacyChoicesUrl,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
}

// NewAppInfoLocalizationUpdateRequestDataAttributes instantiates a new AppInfoLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationUpdateRequestDataAttributes() *AppInfoLocalizationUpdateRequestDataAttributes {
	this := AppInfoLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewAppInfoLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new AppInfoLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationUpdateRequestDataAttributesWithDefaults() *AppInfoLocalizationUpdateRequestDataAttributes {
	this := AppInfoLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyUrl() string {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasPrivacyPolicyUrl() bool {
	if o != nil && !IsNil(o.PrivacyPolicyUrl) {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}

// GetPrivacyChoicesUrl returns the PrivacyChoicesUrl field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyChoicesUrl() string {
	if o == nil || IsNil(o.PrivacyChoicesUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyChoicesUrl
}

// GetPrivacyChoicesUrlOk returns a tuple with the PrivacyChoicesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyChoicesUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyChoicesUrl) {
		return nil, false
	}
	return o.PrivacyChoicesUrl, true
}

// HasPrivacyChoicesUrl returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasPrivacyChoicesUrl() bool {
	if o != nil && !IsNil(o.PrivacyChoicesUrl) {
		return true
	}

	return false
}

// SetPrivacyChoicesUrl gets a reference to the given string and assigns it to the PrivacyChoicesUrl field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetPrivacyChoicesUrl(v string) {
	o.PrivacyChoicesUrl = &v
}

// GetPrivacyPolicyText returns the PrivacyPolicyText field value if set, zero value otherwise.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyText() string {
	if o == nil || IsNil(o.PrivacyPolicyText) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyText
}

// GetPrivacyPolicyTextOk returns a tuple with the PrivacyPolicyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) GetPrivacyPolicyTextOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyText) {
		return nil, false
	}
	return o.PrivacyPolicyText, true
}

// HasPrivacyPolicyText returns a boolean if a field has been set.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) HasPrivacyPolicyText() bool {
	if o != nil && !IsNil(o.PrivacyPolicyText) {
		return true
	}

	return false
}

// SetPrivacyPolicyText gets a reference to the given string and assigns it to the PrivacyPolicyText field.
func (o *AppInfoLocalizationUpdateRequestDataAttributes) SetPrivacyPolicyText(v string) {
	o.PrivacyPolicyText = &v
}

func (o AppInfoLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.PrivacyPolicyUrl) {
		toSerialize["privacyPolicyUrl"] = o.PrivacyPolicyUrl
	}
	if !IsNil(o.PrivacyChoicesUrl) {
		toSerialize["privacyChoicesUrl"] = o.PrivacyChoicesUrl
	}
	if !IsNil(o.PrivacyPolicyText) {
		toSerialize["privacyPolicyText"] = o.PrivacyPolicyText
	}
	return toSerialize, nil
}

type NullableAppInfoLocalizationUpdateRequestDataAttributes struct {
	value *AppInfoLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) Get() *AppInfoLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) Set(val *AppInfoLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationUpdateRequestDataAttributes(val *AppInfoLocalizationUpdateRequestDataAttributes) *NullableAppInfoLocalizationUpdateRequestDataAttributes {
	return &NullableAppInfoLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


