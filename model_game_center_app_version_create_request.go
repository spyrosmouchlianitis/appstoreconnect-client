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

// checks if the GameCenterAppVersionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionCreateRequest{}

// GameCenterAppVersionCreateRequest struct for GameCenterAppVersionCreateRequest
type GameCenterAppVersionCreateRequest struct {
	Data GameCenterAppVersionCreateRequestData `json:"data"`
}

type _GameCenterAppVersionCreateRequest GameCenterAppVersionCreateRequest

// NewGameCenterAppVersionCreateRequest instantiates a new GameCenterAppVersionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionCreateRequest(data GameCenterAppVersionCreateRequestData) *GameCenterAppVersionCreateRequest {
	this := GameCenterAppVersionCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAppVersionCreateRequestWithDefaults instantiates a new GameCenterAppVersionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionCreateRequestWithDefaults() *GameCenterAppVersionCreateRequest {
	this := GameCenterAppVersionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionCreateRequest) GetData() GameCenterAppVersionCreateRequestData {
	if o == nil {
		var ret GameCenterAppVersionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionCreateRequest) GetDataOk() (*GameCenterAppVersionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionCreateRequest) SetData(v GameCenterAppVersionCreateRequestData) {
	o.Data = v
}

func (o GameCenterAppVersionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterAppVersionCreateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterAppVersionCreateRequest := _GameCenterAppVersionCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAppVersionCreateRequest)

	if err != nil {
		return err
	}

	*o = GameCenterAppVersionCreateRequest(varGameCenterAppVersionCreateRequest)

	return err
}

type NullableGameCenterAppVersionCreateRequest struct {
	value *GameCenterAppVersionCreateRequest
	isSet bool
}

func (v NullableGameCenterAppVersionCreateRequest) Get() *GameCenterAppVersionCreateRequest {
	return v.value
}

func (v *NullableGameCenterAppVersionCreateRequest) Set(val *GameCenterAppVersionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionCreateRequest(val *GameCenterAppVersionCreateRequest) *NullableGameCenterAppVersionCreateRequest {
	return &NullableGameCenterAppVersionCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


