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

// checks if the GameCenterLeaderboardSetImageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetImageUpdateRequest{}

// GameCenterLeaderboardSetImageUpdateRequest struct for GameCenterLeaderboardSetImageUpdateRequest
type GameCenterLeaderboardSetImageUpdateRequest struct {
	Data GameCenterLeaderboardSetImageUpdateRequestData `json:"data"`
}

type _GameCenterLeaderboardSetImageUpdateRequest GameCenterLeaderboardSetImageUpdateRequest

// NewGameCenterLeaderboardSetImageUpdateRequest instantiates a new GameCenterLeaderboardSetImageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetImageUpdateRequest(data GameCenterLeaderboardSetImageUpdateRequestData) *GameCenterLeaderboardSetImageUpdateRequest {
	this := GameCenterLeaderboardSetImageUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetImageUpdateRequestWithDefaults instantiates a new GameCenterLeaderboardSetImageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetImageUpdateRequestWithDefaults() *GameCenterLeaderboardSetImageUpdateRequest {
	this := GameCenterLeaderboardSetImageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetImageUpdateRequest) GetData() GameCenterLeaderboardSetImageUpdateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardSetImageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageUpdateRequest) GetDataOk() (*GameCenterLeaderboardSetImageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetImageUpdateRequest) SetData(v GameCenterLeaderboardSetImageUpdateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetImageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetImageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetImageUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetImageUpdateRequest := _GameCenterLeaderboardSetImageUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetImageUpdateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetImageUpdateRequest(varGameCenterLeaderboardSetImageUpdateRequest)

	return err
}

type NullableGameCenterLeaderboardSetImageUpdateRequest struct {
	value *GameCenterLeaderboardSetImageUpdateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequest) Get() *GameCenterLeaderboardSetImageUpdateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequest) Set(val *GameCenterLeaderboardSetImageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetImageUpdateRequest(val *GameCenterLeaderboardSetImageUpdateRequest) *NullableGameCenterLeaderboardSetImageUpdateRequest {
	return &NullableGameCenterLeaderboardSetImageUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


