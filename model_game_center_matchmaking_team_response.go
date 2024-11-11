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

// checks if the GameCenterMatchmakingTeamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamResponse{}

// GameCenterMatchmakingTeamResponse struct for GameCenterMatchmakingTeamResponse
type GameCenterMatchmakingTeamResponse struct {
	Data GameCenterMatchmakingTeam `json:"data"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterMatchmakingTeamResponse GameCenterMatchmakingTeamResponse

// NewGameCenterMatchmakingTeamResponse instantiates a new GameCenterMatchmakingTeamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamResponse(data GameCenterMatchmakingTeam, links DocumentLinks) *GameCenterMatchmakingTeamResponse {
	this := GameCenterMatchmakingTeamResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingTeamResponseWithDefaults instantiates a new GameCenterMatchmakingTeamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamResponseWithDefaults() *GameCenterMatchmakingTeamResponse {
	this := GameCenterMatchmakingTeamResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingTeamResponse) GetData() GameCenterMatchmakingTeam {
	if o == nil {
		var ret GameCenterMatchmakingTeam
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamResponse) GetDataOk() (*GameCenterMatchmakingTeam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingTeamResponse) SetData(v GameCenterMatchmakingTeam) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingTeamResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingTeamResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterMatchmakingTeamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTeamResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varGameCenterMatchmakingTeamResponse := _GameCenterMatchmakingTeamResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTeamResponse)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTeamResponse(varGameCenterMatchmakingTeamResponse)

	return err
}

type NullableGameCenterMatchmakingTeamResponse struct {
	value *GameCenterMatchmakingTeamResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamResponse) Get() *GameCenterMatchmakingTeamResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamResponse) Set(val *GameCenterMatchmakingTeamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamResponse(val *GameCenterMatchmakingTeamResponse) *NullableGameCenterMatchmakingTeamResponse {
	return &NullableGameCenterMatchmakingTeamResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


