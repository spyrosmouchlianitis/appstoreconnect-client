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

// checks if the GameCenterLeaderboardSetLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationCreateRequest{}

// GameCenterLeaderboardSetLocalizationCreateRequest struct for GameCenterLeaderboardSetLocalizationCreateRequest
type GameCenterLeaderboardSetLocalizationCreateRequest struct {
	Data GameCenterLeaderboardSetLocalizationCreateRequestData `json:"data"`
}

type _GameCenterLeaderboardSetLocalizationCreateRequest GameCenterLeaderboardSetLocalizationCreateRequest

// NewGameCenterLeaderboardSetLocalizationCreateRequest instantiates a new GameCenterLeaderboardSetLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationCreateRequest(data GameCenterLeaderboardSetLocalizationCreateRequestData) *GameCenterLeaderboardSetLocalizationCreateRequest {
	this := GameCenterLeaderboardSetLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetLocalizationCreateRequestWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationCreateRequestWithDefaults() *GameCenterLeaderboardSetLocalizationCreateRequest {
	this := GameCenterLeaderboardSetLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequest) GetData() GameCenterLeaderboardSetLocalizationCreateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardSetLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequest) GetDataOk() (*GameCenterLeaderboardSetLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequest) SetData(v GameCenterLeaderboardSetLocalizationCreateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetLocalizationCreateRequest := _GameCenterLeaderboardSetLocalizationCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationCreateRequest(varGameCenterLeaderboardSetLocalizationCreateRequest)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationCreateRequest struct {
	value *GameCenterLeaderboardSetLocalizationCreateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequest) Get() *GameCenterLeaderboardSetLocalizationCreateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequest) Set(val *GameCenterLeaderboardSetLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationCreateRequest(val *GameCenterLeaderboardSetLocalizationCreateRequest) *NullableGameCenterLeaderboardSetLocalizationCreateRequest {
	return &NullableGameCenterLeaderboardSetLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


