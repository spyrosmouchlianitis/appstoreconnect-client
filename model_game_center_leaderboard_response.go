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

// checks if the GameCenterLeaderboardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardResponse{}

// GameCenterLeaderboardResponse struct for GameCenterLeaderboardResponse
type GameCenterLeaderboardResponse struct {
	Data GameCenterLeaderboard `json:"data"`
	Included []GameCenterLeaderboardsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterLeaderboardResponse GameCenterLeaderboardResponse

// NewGameCenterLeaderboardResponse instantiates a new GameCenterLeaderboardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardResponse(data GameCenterLeaderboard, links DocumentLinks) *GameCenterLeaderboardResponse {
	this := GameCenterLeaderboardResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardResponseWithDefaults instantiates a new GameCenterLeaderboardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardResponseWithDefaults() *GameCenterLeaderboardResponse {
	this := GameCenterLeaderboardResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardResponse) GetData() GameCenterLeaderboard {
	if o == nil {
		var ret GameCenterLeaderboard
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardResponse) GetDataOk() (*GameCenterLeaderboard, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardResponse) SetData(v GameCenterLeaderboard) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardResponse) GetIncluded() []GameCenterLeaderboardsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardResponse) GetIncludedOk() ([]GameCenterLeaderboardsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardResponse) SetIncluded(v []GameCenterLeaderboardsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterLeaderboardResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardResponse := _GameCenterLeaderboardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardResponse)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardResponse(varGameCenterLeaderboardResponse)

	return err
}

type NullableGameCenterLeaderboardResponse struct {
	value *GameCenterLeaderboardResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardResponse) Get() *GameCenterLeaderboardResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardResponse) Set(val *GameCenterLeaderboardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardResponse(val *GameCenterLeaderboardResponse) *NullableGameCenterLeaderboardResponse {
	return &NullableGameCenterLeaderboardResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


