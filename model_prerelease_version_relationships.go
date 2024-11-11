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

// checks if the PrereleaseVersionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrereleaseVersionRelationships{}

// PrereleaseVersionRelationships struct for PrereleaseVersionRelationships
type PrereleaseVersionRelationships struct {
	Builds *AppRelationshipsBuilds `json:"builds,omitempty"`
	App *BetaAppLocalizationRelationshipsApp `json:"app,omitempty"`
}

// NewPrereleaseVersionRelationships instantiates a new PrereleaseVersionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrereleaseVersionRelationships() *PrereleaseVersionRelationships {
	this := PrereleaseVersionRelationships{}
	return &this
}

// NewPrereleaseVersionRelationshipsWithDefaults instantiates a new PrereleaseVersionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrereleaseVersionRelationshipsWithDefaults() *PrereleaseVersionRelationships {
	this := PrereleaseVersionRelationships{}
	return &this
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *PrereleaseVersionRelationships) GetBuilds() AppRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret AppRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *PrereleaseVersionRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppRelationshipsBuilds and assigns it to the Builds field.
func (o *PrereleaseVersionRelationships) SetBuilds(v AppRelationshipsBuilds) {
	o.Builds = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *PrereleaseVersionRelationships) GetApp() BetaAppLocalizationRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret BetaAppLocalizationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrereleaseVersionRelationships) GetAppOk() (*BetaAppLocalizationRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *PrereleaseVersionRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given BetaAppLocalizationRelationshipsApp and assigns it to the App field.
func (o *PrereleaseVersionRelationships) SetApp(v BetaAppLocalizationRelationshipsApp) {
	o.App = &v
}

func (o PrereleaseVersionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrereleaseVersionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullablePrereleaseVersionRelationships struct {
	value *PrereleaseVersionRelationships
	isSet bool
}

func (v NullablePrereleaseVersionRelationships) Get() *PrereleaseVersionRelationships {
	return v.value
}

func (v *NullablePrereleaseVersionRelationships) Set(val *PrereleaseVersionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePrereleaseVersionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePrereleaseVersionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrereleaseVersionRelationships(val *PrereleaseVersionRelationships) *NullablePrereleaseVersionRelationships {
	return &NullablePrereleaseVersionRelationships{value: val, isSet: true}
}

func (v NullablePrereleaseVersionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrereleaseVersionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


