/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the GameCenterMatchmakingTeamAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamAssignment{}

// GameCenterMatchmakingTeamAssignment struct for GameCenterMatchmakingTeamAssignment
type GameCenterMatchmakingTeamAssignment struct {
	PlayerId *string `json:"playerId,omitempty"`
	Team *string `json:"team,omitempty"`
}

// NewGameCenterMatchmakingTeamAssignment instantiates a new GameCenterMatchmakingTeamAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamAssignment() *GameCenterMatchmakingTeamAssignment {
	this := GameCenterMatchmakingTeamAssignment{}
	return &this
}

// NewGameCenterMatchmakingTeamAssignmentWithDefaults instantiates a new GameCenterMatchmakingTeamAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamAssignmentWithDefaults() *GameCenterMatchmakingTeamAssignment {
	this := GameCenterMatchmakingTeamAssignment{}
	return &this
}

// GetPlayerId returns the PlayerId field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTeamAssignment) GetPlayerId() string {
	if o == nil || IsNil(o.PlayerId) {
		var ret string
		return ret
	}
	return *o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamAssignment) GetPlayerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlayerId) {
		return nil, false
	}
	return o.PlayerId, true
}

// HasPlayerId returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTeamAssignment) HasPlayerId() bool {
	if o != nil && !IsNil(o.PlayerId) {
		return true
	}

	return false
}

// SetPlayerId gets a reference to the given string and assigns it to the PlayerId field.
func (o *GameCenterMatchmakingTeamAssignment) SetPlayerId(v string) {
	o.PlayerId = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTeamAssignment) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamAssignment) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTeamAssignment) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *GameCenterMatchmakingTeamAssignment) SetTeam(v string) {
	o.Team = &v
}

func (o GameCenterMatchmakingTeamAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlayerId) {
		toSerialize["playerId"] = o.PlayerId
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingTeamAssignment struct {
	value *GameCenterMatchmakingTeamAssignment
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamAssignment) Get() *GameCenterMatchmakingTeamAssignment {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamAssignment) Set(val *GameCenterMatchmakingTeamAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamAssignment(val *GameCenterMatchmakingTeamAssignment) *NullableGameCenterMatchmakingTeamAssignment {
	return &NullableGameCenterMatchmakingTeamAssignment{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


