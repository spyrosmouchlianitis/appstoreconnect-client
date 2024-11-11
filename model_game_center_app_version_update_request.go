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

// checks if the GameCenterAppVersionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionUpdateRequest{}

// GameCenterAppVersionUpdateRequest struct for GameCenterAppVersionUpdateRequest
type GameCenterAppVersionUpdateRequest struct {
	Data GameCenterAppVersionUpdateRequestData `json:"data"`
}

type _GameCenterAppVersionUpdateRequest GameCenterAppVersionUpdateRequest

// NewGameCenterAppVersionUpdateRequest instantiates a new GameCenterAppVersionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionUpdateRequest(data GameCenterAppVersionUpdateRequestData) *GameCenterAppVersionUpdateRequest {
	this := GameCenterAppVersionUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAppVersionUpdateRequestWithDefaults instantiates a new GameCenterAppVersionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionUpdateRequestWithDefaults() *GameCenterAppVersionUpdateRequest {
	this := GameCenterAppVersionUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionUpdateRequest) GetData() GameCenterAppVersionUpdateRequestData {
	if o == nil {
		var ret GameCenterAppVersionUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionUpdateRequest) GetDataOk() (*GameCenterAppVersionUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionUpdateRequest) SetData(v GameCenterAppVersionUpdateRequestData) {
	o.Data = v
}

func (o GameCenterAppVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterAppVersionUpdateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterAppVersionUpdateRequest := _GameCenterAppVersionUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAppVersionUpdateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterAppVersionUpdateRequest(varGameCenterAppVersionUpdateRequest)

	return err
}

type NullableGameCenterAppVersionUpdateRequest struct {
	value *GameCenterAppVersionUpdateRequest
	isSet bool
}

func (v NullableGameCenterAppVersionUpdateRequest) Get() *GameCenterAppVersionUpdateRequest {
	return v.value
}

func (v *NullableGameCenterAppVersionUpdateRequest) Set(val *GameCenterAppVersionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionUpdateRequest(val *GameCenterAppVersionUpdateRequest) *NullableGameCenterAppVersionUpdateRequest {
	return &NullableGameCenterAppVersionUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

