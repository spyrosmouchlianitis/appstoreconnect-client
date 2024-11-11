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

// checks if the BuildBundleRelationshipsBetaAppClipInvocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBundleRelationshipsBetaAppClipInvocations{}

// BuildBundleRelationshipsBetaAppClipInvocations struct for BuildBundleRelationshipsBetaAppClipInvocations
type BuildBundleRelationshipsBetaAppClipInvocations struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData `json:"data,omitempty"`
}

// NewBuildBundleRelationshipsBetaAppClipInvocations instantiates a new BuildBundleRelationshipsBetaAppClipInvocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBundleRelationshipsBetaAppClipInvocations() *BuildBundleRelationshipsBetaAppClipInvocations {
	this := BuildBundleRelationshipsBetaAppClipInvocations{}
	return &this
}

// NewBuildBundleRelationshipsBetaAppClipInvocationsWithDefaults instantiates a new BuildBundleRelationshipsBetaAppClipInvocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBundleRelationshipsBetaAppClipInvocationsWithDefaults() *BuildBundleRelationshipsBetaAppClipInvocations {
	this := BuildBundleRelationshipsBetaAppClipInvocations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetData() []BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData {
	if o == nil || IsNil(o.Data) {
		var ret []BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) GetDataOk() ([]BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData and assigns it to the Data field.
func (o *BuildBundleRelationshipsBetaAppClipInvocations) SetData(v []BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocationData) {
	o.Data = v
}

func (o BuildBundleRelationshipsBetaAppClipInvocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBundleRelationshipsBetaAppClipInvocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuildBundleRelationshipsBetaAppClipInvocations struct {
	value *BuildBundleRelationshipsBetaAppClipInvocations
	isSet bool
}

func (v NullableBuildBundleRelationshipsBetaAppClipInvocations) Get() *BuildBundleRelationshipsBetaAppClipInvocations {
	return v.value
}

func (v *NullableBuildBundleRelationshipsBetaAppClipInvocations) Set(val *BuildBundleRelationshipsBetaAppClipInvocations) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBundleRelationshipsBetaAppClipInvocations) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBundleRelationshipsBetaAppClipInvocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBundleRelationshipsBetaAppClipInvocations(val *BuildBundleRelationshipsBetaAppClipInvocations) *NullableBuildBundleRelationshipsBetaAppClipInvocations {
	return &NullableBuildBundleRelationshipsBetaAppClipInvocations{value: val, isSet: true}
}

func (v NullableBuildBundleRelationshipsBetaAppClipInvocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBundleRelationshipsBetaAppClipInvocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


