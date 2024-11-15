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

// checks if the GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage{}

// GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage struct for GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage
type GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData `json:"data,omitempty"`
}

// NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage instantiates a new GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage() *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage {
	this := GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage{}
	return &this
}

// NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageWithDefaults() *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage {
	this := GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) GetData() GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) GetDataOk() (*GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData and assigns it to the Data field.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) SetData(v GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) {
	o.Data = &v
}

func (o GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage struct {
	value *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) Get() *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) Set(val *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage(val *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage {
	return &NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


