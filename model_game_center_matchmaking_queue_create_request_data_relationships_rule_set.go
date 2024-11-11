/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet{}

// GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet struct for GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet
type GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet struct {
	Data GameCenterMatchmakingQueueRelationshipsRuleSetData `json:"data"`
}

type _GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet

// NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(data GameCenterMatchmakingQueueRelationshipsRuleSetData) *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet{}
	this.Data = data
	return &this
}

// NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSetWithDefaults instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSetWithDefaults() *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) GetData() GameCenterMatchmakingQueueRelationshipsRuleSetData {
	if o == nil {
		var ret GameCenterMatchmakingQueueRelationshipsRuleSetData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) GetDataOk() (*GameCenterMatchmakingQueueRelationshipsRuleSetData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) SetData(v GameCenterMatchmakingQueueRelationshipsRuleSetData) {
	o.Data = v
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet := _GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(varGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet)

	return err
}

type NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet struct {
	value *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) Get() *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) Set(val *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(val *GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet {
	return &NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


