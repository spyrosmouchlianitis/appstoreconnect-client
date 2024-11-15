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

// checks if the GameCenterAppVersionRelationshipsCompatibilityVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionRelationshipsCompatibilityVersions{}

// GameCenterAppVersionRelationshipsCompatibilityVersions struct for GameCenterAppVersionRelationshipsCompatibilityVersions
type GameCenterAppVersionRelationshipsCompatibilityVersions struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppStoreVersionRelationshipsGameCenterAppVersionData `json:"data,omitempty"`
}

// NewGameCenterAppVersionRelationshipsCompatibilityVersions instantiates a new GameCenterAppVersionRelationshipsCompatibilityVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionRelationshipsCompatibilityVersions() *GameCenterAppVersionRelationshipsCompatibilityVersions {
	this := GameCenterAppVersionRelationshipsCompatibilityVersions{}
	return &this
}

// NewGameCenterAppVersionRelationshipsCompatibilityVersionsWithDefaults instantiates a new GameCenterAppVersionRelationshipsCompatibilityVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionRelationshipsCompatibilityVersionsWithDefaults() *GameCenterAppVersionRelationshipsCompatibilityVersions {
	this := GameCenterAppVersionRelationshipsCompatibilityVersions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetData() []AppStoreVersionRelationshipsGameCenterAppVersionData {
	if o == nil || IsNil(o.Data) {
		var ret []AppStoreVersionRelationshipsGameCenterAppVersionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) GetDataOk() ([]AppStoreVersionRelationshipsGameCenterAppVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppStoreVersionRelationshipsGameCenterAppVersionData and assigns it to the Data field.
func (o *GameCenterAppVersionRelationshipsCompatibilityVersions) SetData(v []AppStoreVersionRelationshipsGameCenterAppVersionData) {
	o.Data = v
}

func (o GameCenterAppVersionRelationshipsCompatibilityVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionRelationshipsCompatibilityVersions) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterAppVersionRelationshipsCompatibilityVersions struct {
	value *GameCenterAppVersionRelationshipsCompatibilityVersions
	isSet bool
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersions) Get() *GameCenterAppVersionRelationshipsCompatibilityVersions {
	return v.value
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersions) Set(val *GameCenterAppVersionRelationshipsCompatibilityVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionRelationshipsCompatibilityVersions(val *GameCenterAppVersionRelationshipsCompatibilityVersions) *NullableGameCenterAppVersionRelationshipsCompatibilityVersions {
	return &NullableGameCenterAppVersionRelationshipsCompatibilityVersions{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionRelationshipsCompatibilityVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionRelationshipsCompatibilityVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


