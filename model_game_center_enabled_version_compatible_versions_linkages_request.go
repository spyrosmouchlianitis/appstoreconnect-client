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

// checks if the GameCenterEnabledVersionCompatibleVersionsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{}

// GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct for GameCenterEnabledVersionCompatibleVersionsLinkagesRequest
type GameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	Data []AppRelationshipsGameCenterEnabledVersionsDataInner `json:"data"`
}

type _GameCenterEnabledVersionCompatibleVersionsLinkagesRequest GameCenterEnabledVersionCompatibleVersionsLinkagesRequest

// NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequest instantiates a new GameCenterEnabledVersionCompatibleVersionsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequest(data []AppRelationshipsGameCenterEnabledVersionsDataInner) *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest {
	this := GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequestWithDefaults instantiates a new GameCenterEnabledVersionCompatibleVersionsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequestWithDefaults() *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest {
	this := GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) GetData() []AppRelationshipsGameCenterEnabledVersionsDataInner {
	if o == nil {
		var ret []AppRelationshipsGameCenterEnabledVersionsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) GetDataOk() ([]AppRelationshipsGameCenterEnabledVersionsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) SetData(v []AppRelationshipsGameCenterEnabledVersionsDataInner) {
	o.Data = v
}

func (o GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterEnabledVersionCompatibleVersionsLinkagesRequest := _GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterEnabledVersionCompatibleVersionsLinkagesRequest)

	if err != nil {
		return err
	}

	*o = GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(varGameCenterEnabledVersionCompatibleVersionsLinkagesRequest)

	return err
}

type NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest struct {
	value *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest
	isSet bool
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) Get() *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest {
	return v.value
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) Set(val *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest(val *GameCenterEnabledVersionCompatibleVersionsLinkagesRequest) *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest {
	return &NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest{value: val, isSet: true}
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


