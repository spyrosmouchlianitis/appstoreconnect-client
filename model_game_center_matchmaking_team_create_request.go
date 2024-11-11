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

// checks if the GameCenterMatchmakingTeamCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamCreateRequest{}

// GameCenterMatchmakingTeamCreateRequest struct for GameCenterMatchmakingTeamCreateRequest
type GameCenterMatchmakingTeamCreateRequest struct {
	Data GameCenterMatchmakingTeamCreateRequestData `json:"data"`
}

type _GameCenterMatchmakingTeamCreateRequest GameCenterMatchmakingTeamCreateRequest

// NewGameCenterMatchmakingTeamCreateRequest instantiates a new GameCenterMatchmakingTeamCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamCreateRequest(data GameCenterMatchmakingTeamCreateRequestData) *GameCenterMatchmakingTeamCreateRequest {
	this := GameCenterMatchmakingTeamCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterMatchmakingTeamCreateRequestWithDefaults instantiates a new GameCenterMatchmakingTeamCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamCreateRequestWithDefaults() *GameCenterMatchmakingTeamCreateRequest {
	this := GameCenterMatchmakingTeamCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingTeamCreateRequest) GetData() GameCenterMatchmakingTeamCreateRequestData {
	if o == nil {
		var ret GameCenterMatchmakingTeamCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamCreateRequest) GetDataOk() (*GameCenterMatchmakingTeamCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingTeamCreateRequest) SetData(v GameCenterMatchmakingTeamCreateRequestData) {
	o.Data = v
}

func (o GameCenterMatchmakingTeamCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTeamCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingTeamCreateRequest := _GameCenterMatchmakingTeamCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTeamCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTeamCreateRequest(varGameCenterMatchmakingTeamCreateRequest)

	return err
}

type NullableGameCenterMatchmakingTeamCreateRequest struct {
	value *GameCenterMatchmakingTeamCreateRequest
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamCreateRequest) Get() *GameCenterMatchmakingTeamCreateRequest {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamCreateRequest) Set(val *GameCenterMatchmakingTeamCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamCreateRequest(val *GameCenterMatchmakingTeamCreateRequest) *NullableGameCenterMatchmakingTeamCreateRequest {
	return &NullableGameCenterMatchmakingTeamCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


