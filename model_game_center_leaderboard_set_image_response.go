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

// checks if the GameCenterLeaderboardSetImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetImageResponse{}

// GameCenterLeaderboardSetImageResponse struct for GameCenterLeaderboardSetImageResponse
type GameCenterLeaderboardSetImageResponse struct {
	Data GameCenterLeaderboardSetImage `json:"data"`
	Included []GameCenterLeaderboardSetLocalization `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterLeaderboardSetImageResponse GameCenterLeaderboardSetImageResponse

// NewGameCenterLeaderboardSetImageResponse instantiates a new GameCenterLeaderboardSetImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetImageResponse(data GameCenterLeaderboardSetImage, links DocumentLinks) *GameCenterLeaderboardSetImageResponse {
	this := GameCenterLeaderboardSetImageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetImageResponseWithDefaults instantiates a new GameCenterLeaderboardSetImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetImageResponseWithDefaults() *GameCenterLeaderboardSetImageResponse {
	this := GameCenterLeaderboardSetImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetImageResponse) GetData() GameCenterLeaderboardSetImage {
	if o == nil {
		var ret GameCenterLeaderboardSetImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageResponse) GetDataOk() (*GameCenterLeaderboardSetImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetImageResponse) SetData(v GameCenterLeaderboardSetImage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetImageResponse) GetIncluded() []GameCenterLeaderboardSetLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageResponse) GetIncludedOk() ([]GameCenterLeaderboardSetLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetImageResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetLocalization and assigns it to the Included field.
func (o *GameCenterLeaderboardSetImageResponse) SetIncluded(v []GameCenterLeaderboardSetLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetImageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetImageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardSetImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetImageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetImageResponse := _GameCenterLeaderboardSetImageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetImageResponse)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetImageResponse(varGameCenterLeaderboardSetImageResponse)

	return err
}

type NullableGameCenterLeaderboardSetImageResponse struct {
	value *GameCenterLeaderboardSetImageResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetImageResponse) Get() *GameCenterLeaderboardSetImageResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetImageResponse) Set(val *GameCenterLeaderboardSetImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetImageResponse(val *GameCenterLeaderboardSetImageResponse) *NullableGameCenterLeaderboardSetImageResponse {
	return &NullableGameCenterLeaderboardSetImageResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

