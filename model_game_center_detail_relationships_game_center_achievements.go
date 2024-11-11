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

// checks if the GameCenterDetailRelationshipsGameCenterAchievements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsGameCenterAchievements{}

// GameCenterDetailRelationshipsGameCenterAchievements struct for GameCenterDetailRelationshipsGameCenterAchievements
type GameCenterDetailRelationshipsGameCenterAchievements struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData `json:"data,omitempty"`
}

// NewGameCenterDetailRelationshipsGameCenterAchievements instantiates a new GameCenterDetailRelationshipsGameCenterAchievements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsGameCenterAchievements() *GameCenterDetailRelationshipsGameCenterAchievements {
	this := GameCenterDetailRelationshipsGameCenterAchievements{}
	return &this
}

// NewGameCenterDetailRelationshipsGameCenterAchievementsWithDefaults instantiates a new GameCenterDetailRelationshipsGameCenterAchievements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsGameCenterAchievementsWithDefaults() *GameCenterDetailRelationshipsGameCenterAchievements {
	this := GameCenterDetailRelationshipsGameCenterAchievements{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetData() []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData {
	if o == nil || IsNil(o.Data) {
		var ret []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) GetDataOk() ([]GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData and assigns it to the Data field.
func (o *GameCenterDetailRelationshipsGameCenterAchievements) SetData(v []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) {
	o.Data = v
}

func (o GameCenterDetailRelationshipsGameCenterAchievements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsGameCenterAchievements) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterDetailRelationshipsGameCenterAchievements struct {
	value *GameCenterDetailRelationshipsGameCenterAchievements
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsGameCenterAchievements) Get() *GameCenterDetailRelationshipsGameCenterAchievements {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsGameCenterAchievements) Set(val *GameCenterDetailRelationshipsGameCenterAchievements) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsGameCenterAchievements) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsGameCenterAchievements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsGameCenterAchievements(val *GameCenterDetailRelationshipsGameCenterAchievements) *NullableGameCenterDetailRelationshipsGameCenterAchievements {
	return &NullableGameCenterDetailRelationshipsGameCenterAchievements{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsGameCenterAchievements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsGameCenterAchievements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


