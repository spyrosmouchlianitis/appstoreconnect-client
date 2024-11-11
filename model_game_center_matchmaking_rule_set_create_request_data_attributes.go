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

// checks if the GameCenterMatchmakingRuleSetCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetCreateRequestDataAttributes{}

// GameCenterMatchmakingRuleSetCreateRequestDataAttributes struct for GameCenterMatchmakingRuleSetCreateRequestDataAttributes
type GameCenterMatchmakingRuleSetCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	RuleLanguageVersion int32 `json:"ruleLanguageVersion"`
	MinPlayers int32 `json:"minPlayers"`
	MaxPlayers int32 `json:"maxPlayers"`
}

type _GameCenterMatchmakingRuleSetCreateRequestDataAttributes GameCenterMatchmakingRuleSetCreateRequestDataAttributes

// NewGameCenterMatchmakingRuleSetCreateRequestDataAttributes instantiates a new GameCenterMatchmakingRuleSetCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetCreateRequestDataAttributes(referenceName string, ruleLanguageVersion int32, minPlayers int32, maxPlayers int32) *GameCenterMatchmakingRuleSetCreateRequestDataAttributes {
	this := GameCenterMatchmakingRuleSetCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	this.RuleLanguageVersion = ruleLanguageVersion
	this.MinPlayers = minPlayers
	this.MaxPlayers = maxPlayers
	return &this
}

// NewGameCenterMatchmakingRuleSetCreateRequestDataAttributesWithDefaults instantiates a new GameCenterMatchmakingRuleSetCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetCreateRequestDataAttributesWithDefaults() *GameCenterMatchmakingRuleSetCreateRequestDataAttributes {
	this := GameCenterMatchmakingRuleSetCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetRuleLanguageVersion returns the RuleLanguageVersion field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetRuleLanguageVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleLanguageVersion
}

// GetRuleLanguageVersionOk returns a tuple with the RuleLanguageVersion field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetRuleLanguageVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleLanguageVersion, true
}

// SetRuleLanguageVersion sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) SetRuleLanguageVersion(v int32) {
	o.RuleLanguageVersion = v
}

// GetMinPlayers returns the MinPlayers field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetMinPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinPlayers
}

// GetMinPlayersOk returns a tuple with the MinPlayers field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetMinPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPlayers, true
}

// SetMinPlayers sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) SetMinPlayers(v int32) {
	o.MinPlayers = v
}

// GetMaxPlayers returns the MaxPlayers field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetMaxPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPlayers
}

// GetMaxPlayersOk returns a tuple with the MaxPlayers field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) GetMaxPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPlayers, true
}

// SetMaxPlayers sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) SetMaxPlayers(v int32) {
	o.MaxPlayers = v
}

func (o GameCenterMatchmakingRuleSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["ruleLanguageVersion"] = o.RuleLanguageVersion
	toSerialize["minPlayers"] = o.MinPlayers
	toSerialize["maxPlayers"] = o.MaxPlayers
	return toSerialize, nil
}

func (o *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceName",
		"ruleLanguageVersion",
		"minPlayers",
		"maxPlayers",
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

	varGameCenterMatchmakingRuleSetCreateRequestDataAttributes := _GameCenterMatchmakingRuleSetCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingRuleSetCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingRuleSetCreateRequestDataAttributes(varGameCenterMatchmakingRuleSetCreateRequestDataAttributes)

	return err
}

type NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes struct {
	value *GameCenterMatchmakingRuleSetCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) Get() *GameCenterMatchmakingRuleSetCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) Set(val *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes(val *GameCenterMatchmakingRuleSetCreateRequestDataAttributes) *NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes {
	return &NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

