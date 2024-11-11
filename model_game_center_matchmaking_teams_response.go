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

// checks if the GameCenterMatchmakingTeamsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamsResponse{}

// GameCenterMatchmakingTeamsResponse struct for GameCenterMatchmakingTeamsResponse
type GameCenterMatchmakingTeamsResponse struct {
	Data []GameCenterMatchmakingTeam `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _GameCenterMatchmakingTeamsResponse GameCenterMatchmakingTeamsResponse

// NewGameCenterMatchmakingTeamsResponse instantiates a new GameCenterMatchmakingTeamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamsResponse(data []GameCenterMatchmakingTeam, links PagedDocumentLinks) *GameCenterMatchmakingTeamsResponse {
	this := GameCenterMatchmakingTeamsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingTeamsResponseWithDefaults instantiates a new GameCenterMatchmakingTeamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamsResponseWithDefaults() *GameCenterMatchmakingTeamsResponse {
	this := GameCenterMatchmakingTeamsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingTeamsResponse) GetData() []GameCenterMatchmakingTeam {
	if o == nil {
		var ret []GameCenterMatchmakingTeam
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamsResponse) GetDataOk() ([]GameCenterMatchmakingTeam, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingTeamsResponse) SetData(v []GameCenterMatchmakingTeam) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingTeamsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingTeamsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTeamsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTeamsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterMatchmakingTeamsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterMatchmakingTeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *GameCenterMatchmakingTeamsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingTeamsResponse := _GameCenterMatchmakingTeamsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingTeamsResponse)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingTeamsResponse(varGameCenterMatchmakingTeamsResponse)

	return err
}

type NullableGameCenterMatchmakingTeamsResponse struct {
	value *GameCenterMatchmakingTeamsResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamsResponse) Get() *GameCenterMatchmakingTeamsResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamsResponse) Set(val *GameCenterMatchmakingTeamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamsResponse(val *GameCenterMatchmakingTeamsResponse) *NullableGameCenterMatchmakingTeamsResponse {
	return &NullableGameCenterMatchmakingTeamsResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

