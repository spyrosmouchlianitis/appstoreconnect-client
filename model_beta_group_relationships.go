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

// checks if the BetaGroupRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupRelationships{}

// BetaGroupRelationships struct for BetaGroupRelationships
type BetaGroupRelationships struct {
	App *BetaAppLocalizationRelationshipsApp `json:"app,omitempty"`
	Builds *AppRelationshipsBuilds `json:"builds,omitempty"`
	BetaTesters *BetaGroupRelationshipsBetaTesters `json:"betaTesters,omitempty"`
}

// NewBetaGroupRelationships instantiates a new BetaGroupRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupRelationships() *BetaGroupRelationships {
	this := BetaGroupRelationships{}
	return &this
}

// NewBetaGroupRelationshipsWithDefaults instantiates a new BetaGroupRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupRelationshipsWithDefaults() *BetaGroupRelationships {
	this := BetaGroupRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BetaGroupRelationships) GetApp() BetaAppLocalizationRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret BetaAppLocalizationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupRelationships) GetAppOk() (*BetaAppLocalizationRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BetaGroupRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given BetaAppLocalizationRelationshipsApp and assigns it to the App field.
func (o *BetaGroupRelationships) SetApp(v BetaAppLocalizationRelationshipsApp) {
	o.App = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *BetaGroupRelationships) GetBuilds() AppRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret AppRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *BetaGroupRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppRelationshipsBuilds and assigns it to the Builds field.
func (o *BetaGroupRelationships) SetBuilds(v AppRelationshipsBuilds) {
	o.Builds = &v
}

// GetBetaTesters returns the BetaTesters field value if set, zero value otherwise.
func (o *BetaGroupRelationships) GetBetaTesters() BetaGroupRelationshipsBetaTesters {
	if o == nil || IsNil(o.BetaTesters) {
		var ret BetaGroupRelationshipsBetaTesters
		return ret
	}
	return *o.BetaTesters
}

// GetBetaTestersOk returns a tuple with the BetaTesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupRelationships) GetBetaTestersOk() (*BetaGroupRelationshipsBetaTesters, bool) {
	if o == nil || IsNil(o.BetaTesters) {
		return nil, false
	}
	return o.BetaTesters, true
}

// HasBetaTesters returns a boolean if a field has been set.
func (o *BetaGroupRelationships) HasBetaTesters() bool {
	if o != nil && !IsNil(o.BetaTesters) {
		return true
	}

	return false
}

// SetBetaTesters gets a reference to the given BetaGroupRelationshipsBetaTesters and assigns it to the BetaTesters field.
func (o *BetaGroupRelationships) SetBetaTesters(v BetaGroupRelationshipsBetaTesters) {
	o.BetaTesters = &v
}

func (o BetaGroupRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.BetaTesters) {
		toSerialize["betaTesters"] = o.BetaTesters
	}
	return toSerialize, nil
}

type NullableBetaGroupRelationships struct {
	value *BetaGroupRelationships
	isSet bool
}

func (v NullableBetaGroupRelationships) Get() *BetaGroupRelationships {
	return v.value
}

func (v *NullableBetaGroupRelationships) Set(val *BetaGroupRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupRelationships(val *BetaGroupRelationships) *NullableBetaGroupRelationships {
	return &NullableBetaGroupRelationships{value: val, isSet: true}
}

func (v NullableBetaGroupRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


