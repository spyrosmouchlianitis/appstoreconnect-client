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

// checks if the GameCenterMatchmakingTeamUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamUpdateRequest{}

// GameCenterMatchmakingTeamUpdateRequest struct for GameCenterMatchmakingTeamUpdateRequest
type GameCenterMatchmakingTeamUpdateRequest struct {
	Data GameCenterMatchmakingTeamUpdateRequestData `json:"data"`
}

type _GameCenterMatchmakingTeamUpdateRequest GameCenterMatchmakingTeamUpdateRequest

// NewGameCenterMatchmakingTeamUpdateRequest instantiates a new GameCenterMatchmakingTeamUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamUpdateRequest(data GameCenterMatchmakingTeamUpdateRequestData) *GameCenterMatchmakingTeamUpdateRequest {
	this := GameCenterMatchmakingTeamUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterMatchmakingTeamUpdateRequestWithDefaults instantiates a new GameCenterMatchmakingTeamUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamUpdateRequestWithDefaults() *GameCenterMatchmakingTeamUpdateRequest {
	this := GameCenterMatchmakingTeamUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingTeamUpdateRequest) GetData() GameCenterMatchmakingTeamUpdateRequestData {
	if o == nil {
		var ret GameCenterMatchmakingTeamUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamUpdateRequest) GetDataOk() (*GameCenterMatchmakingTeamUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingTeamUpdateRequest) SetData(v GameCenterMatchmakingTeamUpdateRequestData) {
	o.Data = v
}

func (o GameCenterMatchmakingTeamUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTeamUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingTeamUpdateRequest := _GameCenterMatchmakingTeamUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTeamUpdateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTeamUpdateRequest(varGameCenterMatchmakingTeamUpdateRequest)

	return err
}

type NullableGameCenterMatchmakingTeamUpdateRequest struct {
	value *GameCenterMatchmakingTeamUpdateRequest
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamUpdateRequest) Get() *GameCenterMatchmakingTeamUpdateRequest {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequest) Set(val *GameCenterMatchmakingTeamUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamUpdateRequest(val *GameCenterMatchmakingTeamUpdateRequest) *NullableGameCenterMatchmakingTeamUpdateRequest {
	return &NullableGameCenterMatchmakingTeamUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


