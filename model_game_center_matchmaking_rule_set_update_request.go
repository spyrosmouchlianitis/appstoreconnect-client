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

// checks if the GameCenterMatchmakingRuleSetUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetUpdateRequest{}

// GameCenterMatchmakingRuleSetUpdateRequest struct for GameCenterMatchmakingRuleSetUpdateRequest
type GameCenterMatchmakingRuleSetUpdateRequest struct {
	Data GameCenterMatchmakingRuleSetUpdateRequestData `json:"data"`
}

type _GameCenterMatchmakingRuleSetUpdateRequest GameCenterMatchmakingRuleSetUpdateRequest

// NewGameCenterMatchmakingRuleSetUpdateRequest instantiates a new GameCenterMatchmakingRuleSetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetUpdateRequest(data GameCenterMatchmakingRuleSetUpdateRequestData) *GameCenterMatchmakingRuleSetUpdateRequest {
	this := GameCenterMatchmakingRuleSetUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterMatchmakingRuleSetUpdateRequestWithDefaults instantiates a new GameCenterMatchmakingRuleSetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetUpdateRequestWithDefaults() *GameCenterMatchmakingRuleSetUpdateRequest {
	this := GameCenterMatchmakingRuleSetUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingRuleSetUpdateRequest) GetData() GameCenterMatchmakingRuleSetUpdateRequestData {
	if o == nil {
		var ret GameCenterMatchmakingRuleSetUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetUpdateRequest) GetDataOk() (*GameCenterMatchmakingRuleSetUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingRuleSetUpdateRequest) SetData(v GameCenterMatchmakingRuleSetUpdateRequestData) {
	o.Data = v
}

func (o GameCenterMatchmakingRuleSetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterMatchmakingRuleSetUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingRuleSetUpdateRequest := _GameCenterMatchmakingRuleSetUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingRuleSetUpdateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingRuleSetUpdateRequest(varGameCenterMatchmakingRuleSetUpdateRequest)

	return err
}

type NullableGameCenterMatchmakingRuleSetUpdateRequest struct {
	value *GameCenterMatchmakingRuleSetUpdateRequest
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequest) Get() *GameCenterMatchmakingRuleSetUpdateRequest {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequest) Set(val *GameCenterMatchmakingRuleSetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetUpdateRequest(val *GameCenterMatchmakingRuleSetUpdateRequest) *NullableGameCenterMatchmakingRuleSetUpdateRequest {
	return &NullableGameCenterMatchmakingRuleSetUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


