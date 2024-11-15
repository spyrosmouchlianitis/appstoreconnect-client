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

// checks if the GameCenterPlayerAchievementSubmissionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterPlayerAchievementSubmissionCreateRequest{}

// GameCenterPlayerAchievementSubmissionCreateRequest struct for GameCenterPlayerAchievementSubmissionCreateRequest
type GameCenterPlayerAchievementSubmissionCreateRequest struct {
	Data GameCenterPlayerAchievementSubmissionCreateRequestData `json:"data"`
}

type _GameCenterPlayerAchievementSubmissionCreateRequest GameCenterPlayerAchievementSubmissionCreateRequest

// NewGameCenterPlayerAchievementSubmissionCreateRequest instantiates a new GameCenterPlayerAchievementSubmissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterPlayerAchievementSubmissionCreateRequest(data GameCenterPlayerAchievementSubmissionCreateRequestData) *GameCenterPlayerAchievementSubmissionCreateRequest {
	this := GameCenterPlayerAchievementSubmissionCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterPlayerAchievementSubmissionCreateRequestWithDefaults instantiates a new GameCenterPlayerAchievementSubmissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterPlayerAchievementSubmissionCreateRequestWithDefaults() *GameCenterPlayerAchievementSubmissionCreateRequest {
	this := GameCenterPlayerAchievementSubmissionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterPlayerAchievementSubmissionCreateRequest) GetData() GameCenterPlayerAchievementSubmissionCreateRequestData {
	if o == nil {
		var ret GameCenterPlayerAchievementSubmissionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterPlayerAchievementSubmissionCreateRequest) GetDataOk() (*GameCenterPlayerAchievementSubmissionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterPlayerAchievementSubmissionCreateRequest) SetData(v GameCenterPlayerAchievementSubmissionCreateRequestData) {
	o.Data = v
}

func (o GameCenterPlayerAchievementSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterPlayerAchievementSubmissionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterPlayerAchievementSubmissionCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterPlayerAchievementSubmissionCreateRequest := _GameCenterPlayerAchievementSubmissionCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterPlayerAchievementSubmissionCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterPlayerAchievementSubmissionCreateRequest(varGameCenterPlayerAchievementSubmissionCreateRequest)

	return err
}

type NullableGameCenterPlayerAchievementSubmissionCreateRequest struct {
	value *GameCenterPlayerAchievementSubmissionCreateRequest
	isSet bool
}

func (v NullableGameCenterPlayerAchievementSubmissionCreateRequest) Get() *GameCenterPlayerAchievementSubmissionCreateRequest {
	return v.value
}

func (v *NullableGameCenterPlayerAchievementSubmissionCreateRequest) Set(val *GameCenterPlayerAchievementSubmissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterPlayerAchievementSubmissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterPlayerAchievementSubmissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterPlayerAchievementSubmissionCreateRequest(val *GameCenterPlayerAchievementSubmissionCreateRequest) *NullableGameCenterPlayerAchievementSubmissionCreateRequest {
	return &NullableGameCenterPlayerAchievementSubmissionCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterPlayerAchievementSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterPlayerAchievementSubmissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


