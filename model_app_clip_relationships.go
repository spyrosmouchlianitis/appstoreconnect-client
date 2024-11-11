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

// checks if the AppClipRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipRelationships{}

// AppClipRelationships struct for AppClipRelationships
type AppClipRelationships struct {
	App *AlternativeDistributionKeyCreateRequestDataRelationshipsApp `json:"app,omitempty"`
	AppClipDefaultExperiences *AppClipRelationshipsAppClipDefaultExperiences `json:"appClipDefaultExperiences,omitempty"`
	AppClipAdvancedExperiences *AnalyticsReportInstanceRelationshipsSegments `json:"appClipAdvancedExperiences,omitempty"`
}

// NewAppClipRelationships instantiates a new AppClipRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipRelationships() *AppClipRelationships {
	this := AppClipRelationships{}
	return &this
}

// NewAppClipRelationshipsWithDefaults instantiates a new AppClipRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipRelationshipsWithDefaults() *AppClipRelationships {
	this := AppClipRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppClipRelationships) GetApp() AlternativeDistributionKeyCreateRequestDataRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AlternativeDistributionKeyCreateRequestDataRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipRelationships) GetAppOk() (*AlternativeDistributionKeyCreateRequestDataRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppClipRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AlternativeDistributionKeyCreateRequestDataRelationshipsApp and assigns it to the App field.
func (o *AppClipRelationships) SetApp(v AlternativeDistributionKeyCreateRequestDataRelationshipsApp) {
	o.App = &v
}

// GetAppClipDefaultExperiences returns the AppClipDefaultExperiences field value if set, zero value otherwise.
func (o *AppClipRelationships) GetAppClipDefaultExperiences() AppClipRelationshipsAppClipDefaultExperiences {
	if o == nil || IsNil(o.AppClipDefaultExperiences) {
		var ret AppClipRelationshipsAppClipDefaultExperiences
		return ret
	}
	return *o.AppClipDefaultExperiences
}

// GetAppClipDefaultExperiencesOk returns a tuple with the AppClipDefaultExperiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipRelationships) GetAppClipDefaultExperiencesOk() (*AppClipRelationshipsAppClipDefaultExperiences, bool) {
	if o == nil || IsNil(o.AppClipDefaultExperiences) {
		return nil, false
	}
	return o.AppClipDefaultExperiences, true
}

// HasAppClipDefaultExperiences returns a boolean if a field has been set.
func (o *AppClipRelationships) HasAppClipDefaultExperiences() bool {
	if o != nil && !IsNil(o.AppClipDefaultExperiences) {
		return true
	}

	return false
}

// SetAppClipDefaultExperiences gets a reference to the given AppClipRelationshipsAppClipDefaultExperiences and assigns it to the AppClipDefaultExperiences field.
func (o *AppClipRelationships) SetAppClipDefaultExperiences(v AppClipRelationshipsAppClipDefaultExperiences) {
	o.AppClipDefaultExperiences = &v
}

// GetAppClipAdvancedExperiences returns the AppClipAdvancedExperiences field value if set, zero value otherwise.
func (o *AppClipRelationships) GetAppClipAdvancedExperiences() AnalyticsReportInstanceRelationshipsSegments {
	if o == nil || IsNil(o.AppClipAdvancedExperiences) {
		var ret AnalyticsReportInstanceRelationshipsSegments
		return ret
	}
	return *o.AppClipAdvancedExperiences
}

// GetAppClipAdvancedExperiencesOk returns a tuple with the AppClipAdvancedExperiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipRelationships) GetAppClipAdvancedExperiencesOk() (*AnalyticsReportInstanceRelationshipsSegments, bool) {
	if o == nil || IsNil(o.AppClipAdvancedExperiences) {
		return nil, false
	}
	return o.AppClipAdvancedExperiences, true
}

// HasAppClipAdvancedExperiences returns a boolean if a field has been set.
func (o *AppClipRelationships) HasAppClipAdvancedExperiences() bool {
	if o != nil && !IsNil(o.AppClipAdvancedExperiences) {
		return true
	}

	return false
}

// SetAppClipAdvancedExperiences gets a reference to the given AnalyticsReportInstanceRelationshipsSegments and assigns it to the AppClipAdvancedExperiences field.
func (o *AppClipRelationships) SetAppClipAdvancedExperiences(v AnalyticsReportInstanceRelationshipsSegments) {
	o.AppClipAdvancedExperiences = &v
}

func (o AppClipRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.AppClipDefaultExperiences) {
		toSerialize["appClipDefaultExperiences"] = o.AppClipDefaultExperiences
	}
	if !IsNil(o.AppClipAdvancedExperiences) {
		toSerialize["appClipAdvancedExperiences"] = o.AppClipAdvancedExperiences
	}
	return toSerialize, nil
}

type NullableAppClipRelationships struct {
	value *AppClipRelationships
	isSet bool
}

func (v NullableAppClipRelationships) Get() *AppClipRelationships {
	return v.value
}

func (v *NullableAppClipRelationships) Set(val *AppClipRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipRelationships(val *AppClipRelationships) *NullableAppClipRelationships {
	return &NullableAppClipRelationships{value: val, isSet: true}
}

func (v NullableAppClipRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


