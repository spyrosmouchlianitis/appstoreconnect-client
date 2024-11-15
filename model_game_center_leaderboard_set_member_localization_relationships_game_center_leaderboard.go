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

// checks if the GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard{}

// GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard struct for GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard
type GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data,omitempty"`
}

// NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard instantiates a new GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard{}
	return &this
}

// NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboardWithDefaults() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) GetData() GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) GetDataOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner and assigns it to the Data field.
func (o *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) SetData(v GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = &v
}

func (o GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard struct {
	value *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard
	isSet bool
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) Get() *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) Set(val *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard(val *GameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard {
	return &NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationRelationshipsGameCenterLeaderboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


