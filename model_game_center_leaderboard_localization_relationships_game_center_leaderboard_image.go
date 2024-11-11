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

// checks if the GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage{}

// GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage struct for GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage
type GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData `json:"data,omitempty"`
}

// NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage instantiates a new GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage() *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage {
	this := GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage{}
	return &this
}

// NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageWithDefaults instantiates a new GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageWithDefaults() *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage {
	this := GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) GetData() GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) GetDataOk() (*GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData and assigns it to the Data field.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) SetData(v GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) {
	o.Data = &v
}

func (o GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage struct {
	value *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) Get() *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) Set(val *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage(val *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage {
	return &NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


