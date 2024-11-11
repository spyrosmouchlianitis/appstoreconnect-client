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

// checks if the AlternativeDistributionPackageRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeDistributionPackageRelationships{}

// AlternativeDistributionPackageRelationships struct for AlternativeDistributionPackageRelationships
type AlternativeDistributionPackageRelationships struct {
	Versions *AlternativeDistributionPackageRelationshipsVersions `json:"versions,omitempty"`
}

// NewAlternativeDistributionPackageRelationships instantiates a new AlternativeDistributionPackageRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeDistributionPackageRelationships() *AlternativeDistributionPackageRelationships {
	this := AlternativeDistributionPackageRelationships{}
	return &this
}

// NewAlternativeDistributionPackageRelationshipsWithDefaults instantiates a new AlternativeDistributionPackageRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeDistributionPackageRelationshipsWithDefaults() *AlternativeDistributionPackageRelationships {
	this := AlternativeDistributionPackageRelationships{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AlternativeDistributionPackageRelationships) GetVersions() AlternativeDistributionPackageRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AlternativeDistributionPackageRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeDistributionPackageRelationships) GetVersionsOk() (*AlternativeDistributionPackageRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AlternativeDistributionPackageRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AlternativeDistributionPackageRelationshipsVersions and assigns it to the Versions field.
func (o *AlternativeDistributionPackageRelationships) SetVersions(v AlternativeDistributionPackageRelationshipsVersions) {
	o.Versions = &v
}

func (o AlternativeDistributionPackageRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeDistributionPackageRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableAlternativeDistributionPackageRelationships struct {
	value *AlternativeDistributionPackageRelationships
	isSet bool
}

func (v NullableAlternativeDistributionPackageRelationships) Get() *AlternativeDistributionPackageRelationships {
	return v.value
}

func (v *NullableAlternativeDistributionPackageRelationships) Set(val *AlternativeDistributionPackageRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeDistributionPackageRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeDistributionPackageRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeDistributionPackageRelationships(val *AlternativeDistributionPackageRelationships) *NullableAlternativeDistributionPackageRelationships {
	return &NullableAlternativeDistributionPackageRelationships{value: val, isSet: true}
}

func (v NullableAlternativeDistributionPackageRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeDistributionPackageRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


