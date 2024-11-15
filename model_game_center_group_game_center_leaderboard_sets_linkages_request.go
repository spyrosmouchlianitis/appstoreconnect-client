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

// checks if the GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest{}

// GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest struct for GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest
type GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest struct {
	Data []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner `json:"data"`
}

type _GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest

// NewGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest instantiates a new GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest(data []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest {
	this := GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewGameCenterGroupGameCenterLeaderboardSetsLinkagesRequestWithDefaults instantiates a new GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupGameCenterLeaderboardSetsLinkagesRequestWithDefaults() *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest {
	this := GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	if o == nil {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	o.Data = v
}

func (o GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest := _GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest)

	if err != nil {
		return err
	}

	*o = GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest(varGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest)

	return err
}

type NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest struct {
	value *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest
	isSet bool
}

func (v NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) Get() *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest {
	return v.value
}

func (v *NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) Set(val *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest(val *GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) *NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest {
	return &NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest{value: val, isSet: true}
}

func (v NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


