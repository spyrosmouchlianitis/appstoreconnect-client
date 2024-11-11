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

// checks if the ScmProviderRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmProviderRelationships{}

// ScmProviderRelationships struct for ScmProviderRelationships
type ScmProviderRelationships struct {
	Repositories *AnalyticsReportInstanceRelationshipsSegments `json:"repositories,omitempty"`
}

// NewScmProviderRelationships instantiates a new ScmProviderRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmProviderRelationships() *ScmProviderRelationships {
	this := ScmProviderRelationships{}
	return &this
}

// NewScmProviderRelationshipsWithDefaults instantiates a new ScmProviderRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmProviderRelationshipsWithDefaults() *ScmProviderRelationships {
	this := ScmProviderRelationships{}
	return &this
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ScmProviderRelationships) GetRepositories() AnalyticsReportInstanceRelationshipsSegments {
	if o == nil || IsNil(o.Repositories) {
		var ret AnalyticsReportInstanceRelationshipsSegments
		return ret
	}
	return *o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProviderRelationships) GetRepositoriesOk() (*AnalyticsReportInstanceRelationshipsSegments, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ScmProviderRelationships) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given AnalyticsReportInstanceRelationshipsSegments and assigns it to the Repositories field.
func (o *ScmProviderRelationships) SetRepositories(v AnalyticsReportInstanceRelationshipsSegments) {
	o.Repositories = &v
}

func (o ScmProviderRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmProviderRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	return toSerialize, nil
}

type NullableScmProviderRelationships struct {
	value *ScmProviderRelationships
	isSet bool
}

func (v NullableScmProviderRelationships) Get() *ScmProviderRelationships {
	return v.value
}

func (v *NullableScmProviderRelationships) Set(val *ScmProviderRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableScmProviderRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableScmProviderRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmProviderRelationships(val *ScmProviderRelationships) *NullableScmProviderRelationships {
	return &NullableScmProviderRelationships{value: val, isSet: true}
}

func (v NullableScmProviderRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmProviderRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

