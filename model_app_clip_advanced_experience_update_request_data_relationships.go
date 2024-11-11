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

// checks if the AppClipAdvancedExperienceUpdateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceUpdateRequestDataRelationships{}

// AppClipAdvancedExperienceUpdateRequestDataRelationships struct for AppClipAdvancedExperienceUpdateRequestDataRelationships
type AppClipAdvancedExperienceUpdateRequestDataRelationships struct {
	AppClip *AppClipAdvancedExperienceRelationshipsAppClip `json:"appClip,omitempty"`
	HeaderImage *AppClipAdvancedExperienceRelationshipsHeaderImage `json:"headerImage,omitempty"`
	Localizations *AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations `json:"localizations,omitempty"`
}

// NewAppClipAdvancedExperienceUpdateRequestDataRelationships instantiates a new AppClipAdvancedExperienceUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceUpdateRequestDataRelationships() *AppClipAdvancedExperienceUpdateRequestDataRelationships {
	this := AppClipAdvancedExperienceUpdateRequestDataRelationships{}
	return &this
}

// NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsWithDefaults instantiates a new AppClipAdvancedExperienceUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsWithDefaults() *AppClipAdvancedExperienceUpdateRequestDataRelationships {
	this := AppClipAdvancedExperienceUpdateRequestDataRelationships{}
	return &this
}

// GetAppClip returns the AppClip field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetAppClip() AppClipAdvancedExperienceRelationshipsAppClip {
	if o == nil || IsNil(o.AppClip) {
		var ret AppClipAdvancedExperienceRelationshipsAppClip
		return ret
	}
	return *o.AppClip
}

// GetAppClipOk returns a tuple with the AppClip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetAppClipOk() (*AppClipAdvancedExperienceRelationshipsAppClip, bool) {
	if o == nil || IsNil(o.AppClip) {
		return nil, false
	}
	return o.AppClip, true
}

// HasAppClip returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) HasAppClip() bool {
	if o != nil && !IsNil(o.AppClip) {
		return true
	}

	return false
}

// SetAppClip gets a reference to the given AppClipAdvancedExperienceRelationshipsAppClip and assigns it to the AppClip field.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) SetAppClip(v AppClipAdvancedExperienceRelationshipsAppClip) {
	o.AppClip = &v
}

// GetHeaderImage returns the HeaderImage field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetHeaderImage() AppClipAdvancedExperienceRelationshipsHeaderImage {
	if o == nil || IsNil(o.HeaderImage) {
		var ret AppClipAdvancedExperienceRelationshipsHeaderImage
		return ret
	}
	return *o.HeaderImage
}

// GetHeaderImageOk returns a tuple with the HeaderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetHeaderImageOk() (*AppClipAdvancedExperienceRelationshipsHeaderImage, bool) {
	if o == nil || IsNil(o.HeaderImage) {
		return nil, false
	}
	return o.HeaderImage, true
}

// HasHeaderImage returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) HasHeaderImage() bool {
	if o != nil && !IsNil(o.HeaderImage) {
		return true
	}

	return false
}

// SetHeaderImage gets a reference to the given AppClipAdvancedExperienceRelationshipsHeaderImage and assigns it to the HeaderImage field.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) SetHeaderImage(v AppClipAdvancedExperienceRelationshipsHeaderImage) {
	o.HeaderImage = &v
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetLocalizations() AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations {
	if o == nil || IsNil(o.Localizations) {
		var ret AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations
		return ret
	}
	return *o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) GetLocalizationsOk() (*AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations, bool) {
	if o == nil || IsNil(o.Localizations) {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) HasLocalizations() bool {
	if o != nil && !IsNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations and assigns it to the Localizations field.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationships) SetLocalizations(v AppClipAdvancedExperienceUpdateRequestDataRelationshipsLocalizations) {
	o.Localizations = &v
}

func (o AppClipAdvancedExperienceUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceUpdateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppClip) {
		toSerialize["appClip"] = o.AppClip
	}
	if !IsNil(o.HeaderImage) {
		toSerialize["headerImage"] = o.HeaderImage
	}
	if !IsNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceUpdateRequestDataRelationships struct {
	value *AppClipAdvancedExperienceUpdateRequestDataRelationships
	isSet bool
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) Get() *AppClipAdvancedExperienceUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) Set(val *AppClipAdvancedExperienceUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceUpdateRequestDataRelationships(val *AppClipAdvancedExperienceUpdateRequestDataRelationships) *NullableAppClipAdvancedExperienceUpdateRequestDataRelationships {
	return &NullableAppClipAdvancedExperienceUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


