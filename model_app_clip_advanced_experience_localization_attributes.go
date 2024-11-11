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

// checks if the AppClipAdvancedExperienceLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceLocalizationAttributes{}

// AppClipAdvancedExperienceLocalizationAttributes struct for AppClipAdvancedExperienceLocalizationAttributes
type AppClipAdvancedExperienceLocalizationAttributes struct {
	Language *AppClipAdvancedExperienceLanguage `json:"language,omitempty"`
	Title *string `json:"title,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
}

// NewAppClipAdvancedExperienceLocalizationAttributes instantiates a new AppClipAdvancedExperienceLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceLocalizationAttributes() *AppClipAdvancedExperienceLocalizationAttributes {
	this := AppClipAdvancedExperienceLocalizationAttributes{}
	return &this
}

// NewAppClipAdvancedExperienceLocalizationAttributesWithDefaults instantiates a new AppClipAdvancedExperienceLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceLocalizationAttributesWithDefaults() *AppClipAdvancedExperienceLocalizationAttributes {
	this := AppClipAdvancedExperienceLocalizationAttributes{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetLanguage() AppClipAdvancedExperienceLanguage {
	if o == nil || IsNil(o.Language) {
		var ret AppClipAdvancedExperienceLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetLanguageOk() (*AppClipAdvancedExperienceLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given AppClipAdvancedExperienceLanguage and assigns it to the Language field.
func (o *AppClipAdvancedExperienceLocalizationAttributes) SetLanguage(v AppClipAdvancedExperienceLanguage) {
	o.Language = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AppClipAdvancedExperienceLocalizationAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceLocalizationAttributes) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *AppClipAdvancedExperienceLocalizationAttributes) SetSubtitle(v string) {
	o.Subtitle = &v
}

func (o AppClipAdvancedExperienceLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceLocalizationAttributes struct {
	value *AppClipAdvancedExperienceLocalizationAttributes
	isSet bool
}

func (v NullableAppClipAdvancedExperienceLocalizationAttributes) Get() *AppClipAdvancedExperienceLocalizationAttributes {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceLocalizationAttributes) Set(val *AppClipAdvancedExperienceLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceLocalizationAttributes(val *AppClipAdvancedExperienceLocalizationAttributes) *NullableAppClipAdvancedExperienceLocalizationAttributes {
	return &NullableAppClipAdvancedExperienceLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


