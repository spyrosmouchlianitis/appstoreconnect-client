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

// checks if the BetaAppLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationRelationships{}

// BetaAppLocalizationRelationships struct for BetaAppLocalizationRelationships
type BetaAppLocalizationRelationships struct {
	App *BetaAppLocalizationRelationshipsApp `json:"app,omitempty"`
}

// NewBetaAppLocalizationRelationships instantiates a new BetaAppLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationRelationships() *BetaAppLocalizationRelationships {
	this := BetaAppLocalizationRelationships{}
	return &this
}

// NewBetaAppLocalizationRelationshipsWithDefaults instantiates a new BetaAppLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationRelationshipsWithDefaults() *BetaAppLocalizationRelationships {
	this := BetaAppLocalizationRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BetaAppLocalizationRelationships) GetApp() BetaAppLocalizationRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret BetaAppLocalizationRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationRelationships) GetAppOk() (*BetaAppLocalizationRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BetaAppLocalizationRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given BetaAppLocalizationRelationshipsApp and assigns it to the App field.
func (o *BetaAppLocalizationRelationships) SetApp(v BetaAppLocalizationRelationshipsApp) {
	o.App = &v
}

func (o BetaAppLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullableBetaAppLocalizationRelationships struct {
	value *BetaAppLocalizationRelationships
	isSet bool
}

func (v NullableBetaAppLocalizationRelationships) Get() *BetaAppLocalizationRelationships {
	return v.value
}

func (v *NullableBetaAppLocalizationRelationships) Set(val *BetaAppLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationRelationships(val *BetaAppLocalizationRelationships) *NullableBetaAppLocalizationRelationships {
	return &NullableBetaAppLocalizationRelationships{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


