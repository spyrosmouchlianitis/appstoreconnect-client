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

// checks if the GameCenterDetailRelationshipsGameCenterGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsGameCenterGroup{}

// GameCenterDetailRelationshipsGameCenterGroup struct for GameCenterDetailRelationshipsGameCenterGroup
type GameCenterDetailRelationshipsGameCenterGroup struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *GameCenterAchievementRelationshipsGameCenterGroupData `json:"data,omitempty"`
}

// NewGameCenterDetailRelationshipsGameCenterGroup instantiates a new GameCenterDetailRelationshipsGameCenterGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsGameCenterGroup() *GameCenterDetailRelationshipsGameCenterGroup {
	this := GameCenterDetailRelationshipsGameCenterGroup{}
	return &this
}

// NewGameCenterDetailRelationshipsGameCenterGroupWithDefaults instantiates a new GameCenterDetailRelationshipsGameCenterGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsGameCenterGroupWithDefaults() *GameCenterDetailRelationshipsGameCenterGroup {
	this := GameCenterDetailRelationshipsGameCenterGroup{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterGroup) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterGroup) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterDetailRelationshipsGameCenterGroup) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterGroup) GetData() GameCenterAchievementRelationshipsGameCenterGroupData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterAchievementRelationshipsGameCenterGroupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterGroup) GetDataOk() (*GameCenterAchievementRelationshipsGameCenterGroupData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterGroup) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterAchievementRelationshipsGameCenterGroupData and assigns it to the Data field.
func (o *GameCenterDetailRelationshipsGameCenterGroup) SetData(v GameCenterAchievementRelationshipsGameCenterGroupData) {
	o.Data = &v
}

func (o GameCenterDetailRelationshipsGameCenterGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsGameCenterGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterDetailRelationshipsGameCenterGroup struct {
	value *GameCenterDetailRelationshipsGameCenterGroup
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsGameCenterGroup) Get() *GameCenterDetailRelationshipsGameCenterGroup {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsGameCenterGroup) Set(val *GameCenterDetailRelationshipsGameCenterGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsGameCenterGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsGameCenterGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsGameCenterGroup(val *GameCenterDetailRelationshipsGameCenterGroup) *NullableGameCenterDetailRelationshipsGameCenterGroup {
	return &NullableGameCenterDetailRelationshipsGameCenterGroup{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsGameCenterGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsGameCenterGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


