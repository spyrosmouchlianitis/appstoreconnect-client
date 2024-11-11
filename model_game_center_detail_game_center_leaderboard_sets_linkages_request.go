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

// checks if the GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest{}

// GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest struct for GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest
type GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest struct {
	Data []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner `json:"data"`
}

type _GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest

// NewGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest instantiates a new GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest(data []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest {
	this := GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewGameCenterDetailGameCenterLeaderboardSetsLinkagesRequestWithDefaults instantiates a new GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailGameCenterLeaderboardSetsLinkagesRequestWithDefaults() *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest {
	this := GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	if o == nil {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	o.Data = v
}

func (o GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest := _GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest)

	if err != nil {
		return err
	}

	*o = GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest(varGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest)

	return err
}

type NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest struct {
	value *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest
	isSet bool
}

func (v NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) Get() *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest {
	return v.value
}

func (v *NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) Set(val *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest(val *GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) *NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest {
	return &NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest{value: val, isSet: true}
}

func (v NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


