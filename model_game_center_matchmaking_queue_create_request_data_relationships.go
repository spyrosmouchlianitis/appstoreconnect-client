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

// checks if the GameCenterMatchmakingQueueCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueCreateRequestDataRelationships{}

// GameCenterMatchmakingQueueCreateRequestDataRelationships struct for GameCenterMatchmakingQueueCreateRequestDataRelationships
type GameCenterMatchmakingQueueCreateRequestDataRelationships struct {
	RuleSet GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet `json:"ruleSet"`
	ExperimentRuleSet *GameCenterMatchmakingQueueRelationshipsRuleSet `json:"experimentRuleSet,omitempty"`
}

type _GameCenterMatchmakingQueueCreateRequestDataRelationships GameCenterMatchmakingQueueCreateRequestDataRelationships

// NewGameCenterMatchmakingQueueCreateRequestDataRelationships instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueCreateRequestDataRelationships(ruleSet GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) *GameCenterMatchmakingQueueCreateRequestDataRelationships {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationships{}
	this.RuleSet = ruleSet
	return &this
}

// NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsWithDefaults() *GameCenterMatchmakingQueueCreateRequestDataRelationships {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationships{}
	return &this
}

// GetRuleSet returns the RuleSet field value
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) GetRuleSet() GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet {
	if o == nil {
		var ret GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet
		return ret
	}

	return o.RuleSet
}

// GetRuleSetOk returns a tuple with the RuleSet field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) GetRuleSetOk() (*GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleSet, true
}

// SetRuleSet sets field value
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) SetRuleSet(v GameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet) {
	o.RuleSet = v
}

// GetExperimentRuleSet returns the ExperimentRuleSet field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) GetExperimentRuleSet() GameCenterMatchmakingQueueRelationshipsRuleSet {
	if o == nil || IsNil(o.ExperimentRuleSet) {
		var ret GameCenterMatchmakingQueueRelationshipsRuleSet
		return ret
	}
	return *o.ExperimentRuleSet
}

// GetExperimentRuleSetOk returns a tuple with the ExperimentRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) GetExperimentRuleSetOk() (*GameCenterMatchmakingQueueRelationshipsRuleSet, bool) {
	if o == nil || IsNil(o.ExperimentRuleSet) {
		return nil, false
	}
	return o.ExperimentRuleSet, true
}

// HasExperimentRuleSet returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) HasExperimentRuleSet() bool {
	if o != nil && !IsNil(o.ExperimentRuleSet) {
		return true
	}

	return false
}

// SetExperimentRuleSet gets a reference to the given GameCenterMatchmakingQueueRelationshipsRuleSet and assigns it to the ExperimentRuleSet field.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) SetExperimentRuleSet(v GameCenterMatchmakingQueueRelationshipsRuleSet) {
	o.ExperimentRuleSet = &v
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ruleSet"] = o.RuleSet
	if !IsNil(o.ExperimentRuleSet) {
		toSerialize["experimentRuleSet"] = o.ExperimentRuleSet
	}
	return toSerialize, nil
}

func (o *GameCenterMatchmakingQueueCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ruleSet",
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

	varGameCenterMatchmakingQueueCreateRequestDataRelationships := _GameCenterMatchmakingQueueCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingQueueCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingQueueCreateRequestDataRelationships(varGameCenterMatchmakingQueueCreateRequestDataRelationships)

	return err
}

type NullableGameCenterMatchmakingQueueCreateRequestDataRelationships struct {
	value *GameCenterMatchmakingQueueCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) Get() *GameCenterMatchmakingQueueCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) Set(val *GameCenterMatchmakingQueueCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueCreateRequestDataRelationships(val *GameCenterMatchmakingQueueCreateRequestDataRelationships) *NullableGameCenterMatchmakingQueueCreateRequestDataRelationships {
	return &NullableGameCenterMatchmakingQueueCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

