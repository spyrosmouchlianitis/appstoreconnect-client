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

// checks if the GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet{}

// GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet struct for GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet
type GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner `json:"data,omitempty"`
}

// NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet instantiates a new GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet {
	this := GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet{}
	return &this
}

// NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSetWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSetWithDefaults() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet {
	this := GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) GetData() GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) GetDataOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner and assigns it to the Data field.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) SetData(v GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	o.Data = &v
}

func (o GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet struct {
	value *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet
	isSet bool
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) Get() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) Set(val *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet(val *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet {
	return &NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


