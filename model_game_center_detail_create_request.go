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

// checks if the GameCenterDetailCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailCreateRequest{}

// GameCenterDetailCreateRequest struct for GameCenterDetailCreateRequest
type GameCenterDetailCreateRequest struct {
	Data GameCenterDetailCreateRequestData `json:"data"`
}

type _GameCenterDetailCreateRequest GameCenterDetailCreateRequest

// NewGameCenterDetailCreateRequest instantiates a new GameCenterDetailCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailCreateRequest(data GameCenterDetailCreateRequestData) *GameCenterDetailCreateRequest {
	this := GameCenterDetailCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterDetailCreateRequestWithDefaults instantiates a new GameCenterDetailCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailCreateRequestWithDefaults() *GameCenterDetailCreateRequest {
	this := GameCenterDetailCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterDetailCreateRequest) GetData() GameCenterDetailCreateRequestData {
	if o == nil {
		var ret GameCenterDetailCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailCreateRequest) GetDataOk() (*GameCenterDetailCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterDetailCreateRequest) SetData(v GameCenterDetailCreateRequestData) {
	o.Data = v
}

func (o GameCenterDetailCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterDetailCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterDetailCreateRequest := _GameCenterDetailCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterDetailCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterDetailCreateRequest(varGameCenterDetailCreateRequest)

	return err
}

type NullableGameCenterDetailCreateRequest struct {
	value *GameCenterDetailCreateRequest
	isSet bool
}

func (v NullableGameCenterDetailCreateRequest) Get() *GameCenterDetailCreateRequest {
	return v.value
}

func (v *NullableGameCenterDetailCreateRequest) Set(val *GameCenterDetailCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailCreateRequest(val *GameCenterDetailCreateRequest) *NullableGameCenterDetailCreateRequest {
	return &NullableGameCenterDetailCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterDetailCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

