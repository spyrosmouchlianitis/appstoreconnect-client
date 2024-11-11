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

// checks if the BuildBundleRelationshipsAppClipDomainCacheStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBundleRelationshipsAppClipDomainCacheStatus{}

// BuildBundleRelationshipsAppClipDomainCacheStatus struct for BuildBundleRelationshipsAppClipDomainCacheStatus
type BuildBundleRelationshipsAppClipDomainCacheStatus struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *BuildBundleRelationshipsAppClipDomainCacheStatusData `json:"data,omitempty"`
}

// NewBuildBundleRelationshipsAppClipDomainCacheStatus instantiates a new BuildBundleRelationshipsAppClipDomainCacheStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBundleRelationshipsAppClipDomainCacheStatus() *BuildBundleRelationshipsAppClipDomainCacheStatus {
	this := BuildBundleRelationshipsAppClipDomainCacheStatus{}
	return &this
}

// NewBuildBundleRelationshipsAppClipDomainCacheStatusWithDefaults instantiates a new BuildBundleRelationshipsAppClipDomainCacheStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBundleRelationshipsAppClipDomainCacheStatusWithDefaults() *BuildBundleRelationshipsAppClipDomainCacheStatus {
	this := BuildBundleRelationshipsAppClipDomainCacheStatus{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) GetData() BuildBundleRelationshipsAppClipDomainCacheStatusData {
	if o == nil || IsNil(o.Data) {
		var ret BuildBundleRelationshipsAppClipDomainCacheStatusData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) GetDataOk() (*BuildBundleRelationshipsAppClipDomainCacheStatusData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuildBundleRelationshipsAppClipDomainCacheStatusData and assigns it to the Data field.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatus) SetData(v BuildBundleRelationshipsAppClipDomainCacheStatusData) {
	o.Data = &v
}

func (o BuildBundleRelationshipsAppClipDomainCacheStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBundleRelationshipsAppClipDomainCacheStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuildBundleRelationshipsAppClipDomainCacheStatus struct {
	value *BuildBundleRelationshipsAppClipDomainCacheStatus
	isSet bool
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatus) Get() *BuildBundleRelationshipsAppClipDomainCacheStatus {
	return v.value
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatus) Set(val *BuildBundleRelationshipsAppClipDomainCacheStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBundleRelationshipsAppClipDomainCacheStatus(val *BuildBundleRelationshipsAppClipDomainCacheStatus) *NullableBuildBundleRelationshipsAppClipDomainCacheStatus {
	return &NullableBuildBundleRelationshipsAppClipDomainCacheStatus{value: val, isSet: true}
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


