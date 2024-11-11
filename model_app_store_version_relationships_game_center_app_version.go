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

// checks if the AppStoreVersionRelationshipsGameCenterAppVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsGameCenterAppVersion{}

// AppStoreVersionRelationshipsGameCenterAppVersion struct for AppStoreVersionRelationshipsGameCenterAppVersion
type AppStoreVersionRelationshipsGameCenterAppVersion struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *AppStoreVersionRelationshipsGameCenterAppVersionData `json:"data,omitempty"`
}

// NewAppStoreVersionRelationshipsGameCenterAppVersion instantiates a new AppStoreVersionRelationshipsGameCenterAppVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsGameCenterAppVersion() *AppStoreVersionRelationshipsGameCenterAppVersion {
	this := AppStoreVersionRelationshipsGameCenterAppVersion{}
	return &this
}

// NewAppStoreVersionRelationshipsGameCenterAppVersionWithDefaults instantiates a new AppStoreVersionRelationshipsGameCenterAppVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsGameCenterAppVersionWithDefaults() *AppStoreVersionRelationshipsGameCenterAppVersion {
	this := AppStoreVersionRelationshipsGameCenterAppVersion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) GetData() AppStoreVersionRelationshipsGameCenterAppVersionData {
	if o == nil || IsNil(o.Data) {
		var ret AppStoreVersionRelationshipsGameCenterAppVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) GetDataOk() (*AppStoreVersionRelationshipsGameCenterAppVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreVersionRelationshipsGameCenterAppVersionData and assigns it to the Data field.
func (o *AppStoreVersionRelationshipsGameCenterAppVersion) SetData(v AppStoreVersionRelationshipsGameCenterAppVersionData) {
	o.Data = &v
}

func (o AppStoreVersionRelationshipsGameCenterAppVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsGameCenterAppVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsGameCenterAppVersion struct {
	value *AppStoreVersionRelationshipsGameCenterAppVersion
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsGameCenterAppVersion) Get() *AppStoreVersionRelationshipsGameCenterAppVersion {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsGameCenterAppVersion) Set(val *AppStoreVersionRelationshipsGameCenterAppVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsGameCenterAppVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsGameCenterAppVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsGameCenterAppVersion(val *AppStoreVersionRelationshipsGameCenterAppVersion) *NullableAppStoreVersionRelationshipsGameCenterAppVersion {
	return &NullableAppStoreVersionRelationshipsGameCenterAppVersion{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsGameCenterAppVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsGameCenterAppVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


