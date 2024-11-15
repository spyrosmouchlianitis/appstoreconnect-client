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

// checks if the GameCenterMatchmakingQueueCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueCreateRequest{}

// GameCenterMatchmakingQueueCreateRequest struct for GameCenterMatchmakingQueueCreateRequest
type GameCenterMatchmakingQueueCreateRequest struct {
	Data GameCenterMatchmakingQueueCreateRequestData `json:"data"`
}

type _GameCenterMatchmakingQueueCreateRequest GameCenterMatchmakingQueueCreateRequest

// NewGameCenterMatchmakingQueueCreateRequest instantiates a new GameCenterMatchmakingQueueCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueCreateRequest(data GameCenterMatchmakingQueueCreateRequestData) *GameCenterMatchmakingQueueCreateRequest {
	this := GameCenterMatchmakingQueueCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterMatchmakingQueueCreateRequestWithDefaults instantiates a new GameCenterMatchmakingQueueCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueCreateRequestWithDefaults() *GameCenterMatchmakingQueueCreateRequest {
	this := GameCenterMatchmakingQueueCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingQueueCreateRequest) GetData() GameCenterMatchmakingQueueCreateRequestData {
	if o == nil {
		var ret GameCenterMatchmakingQueueCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequest) GetDataOk() (*GameCenterMatchmakingQueueCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingQueueCreateRequest) SetData(v GameCenterMatchmakingQueueCreateRequestData) {
	o.Data = v
}

func (o GameCenterMatchmakingQueueCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterMatchmakingQueueCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingQueueCreateRequest := _GameCenterMatchmakingQueueCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingQueueCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingQueueCreateRequest(varGameCenterMatchmakingQueueCreateRequest)

	return err
}

type NullableGameCenterMatchmakingQueueCreateRequest struct {
	value *GameCenterMatchmakingQueueCreateRequest
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueCreateRequest) Get() *GameCenterMatchmakingQueueCreateRequest {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueCreateRequest) Set(val *GameCenterMatchmakingQueueCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueCreateRequest(val *GameCenterMatchmakingQueueCreateRequest) *NullableGameCenterMatchmakingQueueCreateRequest {
	return &NullableGameCenterMatchmakingQueueCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


