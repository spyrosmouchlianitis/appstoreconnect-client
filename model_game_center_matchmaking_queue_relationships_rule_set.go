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

// checks if the GameCenterMatchmakingQueueRelationshipsRuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueRelationshipsRuleSet{}

// GameCenterMatchmakingQueueRelationshipsRuleSet struct for GameCenterMatchmakingQueueRelationshipsRuleSet
type GameCenterMatchmakingQueueRelationshipsRuleSet struct {
	Data *GameCenterMatchmakingQueueRelationshipsRuleSetData `json:"data,omitempty"`
}

// NewGameCenterMatchmakingQueueRelationshipsRuleSet instantiates a new GameCenterMatchmakingQueueRelationshipsRuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueRelationshipsRuleSet() *GameCenterMatchmakingQueueRelationshipsRuleSet {
	this := GameCenterMatchmakingQueueRelationshipsRuleSet{}
	return &this
}

// NewGameCenterMatchmakingQueueRelationshipsRuleSetWithDefaults instantiates a new GameCenterMatchmakingQueueRelationshipsRuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueRelationshipsRuleSetWithDefaults() *GameCenterMatchmakingQueueRelationshipsRuleSet {
	this := GameCenterMatchmakingQueueRelationshipsRuleSet{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueRelationshipsRuleSet) GetData() GameCenterMatchmakingQueueRelationshipsRuleSetData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterMatchmakingQueueRelationshipsRuleSetData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueRelationshipsRuleSet) GetDataOk() (*GameCenterMatchmakingQueueRelationshipsRuleSetData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueRelationshipsRuleSet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterMatchmakingQueueRelationshipsRuleSetData and assigns it to the Data field.
func (o *GameCenterMatchmakingQueueRelationshipsRuleSet) SetData(v GameCenterMatchmakingQueueRelationshipsRuleSetData) {
	o.Data = &v
}

func (o GameCenterMatchmakingQueueRelationshipsRuleSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueRelationshipsRuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueRelationshipsRuleSet struct {
	value *GameCenterMatchmakingQueueRelationshipsRuleSet
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueRelationshipsRuleSet) Get() *GameCenterMatchmakingQueueRelationshipsRuleSet {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueRelationshipsRuleSet) Set(val *GameCenterMatchmakingQueueRelationshipsRuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueRelationshipsRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueRelationshipsRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueRelationshipsRuleSet(val *GameCenterMatchmakingQueueRelationshipsRuleSet) *NullableGameCenterMatchmakingQueueRelationshipsRuleSet {
	return &NullableGameCenterMatchmakingQueueRelationshipsRuleSet{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueRelationshipsRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueRelationshipsRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


