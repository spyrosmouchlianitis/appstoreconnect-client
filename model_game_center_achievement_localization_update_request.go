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

// checks if the GameCenterAchievementLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationUpdateRequest{}

// GameCenterAchievementLocalizationUpdateRequest struct for GameCenterAchievementLocalizationUpdateRequest
type GameCenterAchievementLocalizationUpdateRequest struct {
	Data GameCenterAchievementLocalizationUpdateRequestData `json:"data"`
}

type _GameCenterAchievementLocalizationUpdateRequest GameCenterAchievementLocalizationUpdateRequest

// NewGameCenterAchievementLocalizationUpdateRequest instantiates a new GameCenterAchievementLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationUpdateRequest(data GameCenterAchievementLocalizationUpdateRequestData) *GameCenterAchievementLocalizationUpdateRequest {
	this := GameCenterAchievementLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAchievementLocalizationUpdateRequestWithDefaults instantiates a new GameCenterAchievementLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationUpdateRequestWithDefaults() *GameCenterAchievementLocalizationUpdateRequest {
	this := GameCenterAchievementLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementLocalizationUpdateRequest) GetData() GameCenterAchievementLocalizationUpdateRequestData {
	if o == nil {
		var ret GameCenterAchievementLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationUpdateRequest) GetDataOk() (*GameCenterAchievementLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementLocalizationUpdateRequest) SetData(v GameCenterAchievementLocalizationUpdateRequestData) {
	o.Data = v
}

func (o GameCenterAchievementLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterAchievementLocalizationUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterAchievementLocalizationUpdateRequest := _GameCenterAchievementLocalizationUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementLocalizationUpdateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementLocalizationUpdateRequest(varGameCenterAchievementLocalizationUpdateRequest)

	return err
}

type NullableGameCenterAchievementLocalizationUpdateRequest struct {
	value *GameCenterAchievementLocalizationUpdateRequest
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationUpdateRequest) Get() *GameCenterAchievementLocalizationUpdateRequest {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequest) Set(val *GameCenterAchievementLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationUpdateRequest(val *GameCenterAchievementLocalizationUpdateRequest) *NullableGameCenterAchievementLocalizationUpdateRequest {
	return &NullableGameCenterAchievementLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


