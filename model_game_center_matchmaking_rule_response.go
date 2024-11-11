/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameCenterMatchmakingRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleResponse{}

// GameCenterMatchmakingRuleResponse struct for GameCenterMatchmakingRuleResponse
type GameCenterMatchmakingRuleResponse struct {
	Data GameCenterMatchmakingRule `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterMatchmakingRuleResponse GameCenterMatchmakingRuleResponse

// NewGameCenterMatchmakingRuleResponse instantiates a new GameCenterMatchmakingRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleResponse(data GameCenterMatchmakingRule, links DocumentLinks) *GameCenterMatchmakingRuleResponse {
	this := GameCenterMatchmakingRuleResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingRuleResponseWithDefaults instantiates a new GameCenterMatchmakingRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleResponseWithDefaults() *GameCenterMatchmakingRuleResponse {
	this := GameCenterMatchmakingRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingRuleResponse) GetData() GameCenterMatchmakingRule {
	if o == nil {
		var ret GameCenterMatchmakingRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleResponse) GetDataOk() (*GameCenterMatchmakingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingRuleResponse) SetData(v GameCenterMatchmakingRule) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingRuleResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingRuleResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterMatchmakingRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterMatchmakingRuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varGameCenterMatchmakingRuleResponse := _GameCenterMatchmakingRuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingRuleResponse)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingRuleResponse(varGameCenterMatchmakingRuleResponse)

	return err
}

type NullableGameCenterMatchmakingRuleResponse struct {
	value *GameCenterMatchmakingRuleResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleResponse) Get() *GameCenterMatchmakingRuleResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleResponse) Set(val *GameCenterMatchmakingRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleResponse(val *GameCenterMatchmakingRuleResponse) *NullableGameCenterMatchmakingRuleResponse {
	return &NullableGameCenterMatchmakingRuleResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


