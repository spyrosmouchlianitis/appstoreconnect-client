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

// checks if the GameCenterLeaderboardSetLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationResponse{}

// GameCenterLeaderboardSetLocalizationResponse struct for GameCenterLeaderboardSetLocalizationResponse
type GameCenterLeaderboardSetLocalizationResponse struct {
	Data GameCenterLeaderboardSetLocalization `json:"data"`
	Included []GameCenterLeaderboardSetLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterLeaderboardSetLocalizationResponse GameCenterLeaderboardSetLocalizationResponse

// NewGameCenterLeaderboardSetLocalizationResponse instantiates a new GameCenterLeaderboardSetLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationResponse(data GameCenterLeaderboardSetLocalization, links DocumentLinks) *GameCenterLeaderboardSetLocalizationResponse {
	this := GameCenterLeaderboardSetLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetLocalizationResponseWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationResponseWithDefaults() *GameCenterLeaderboardSetLocalizationResponse {
	this := GameCenterLeaderboardSetLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetLocalizationResponse) GetData() GameCenterLeaderboardSetLocalization {
	if o == nil {
		var ret GameCenterLeaderboardSetLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationResponse) GetDataOk() (*GameCenterLeaderboardSetLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetLocalizationResponse) SetData(v GameCenterLeaderboardSetLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalizationResponse) GetIncluded() []GameCenterLeaderboardSetLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationResponse) GetIncludedOk() ([]GameCenterLeaderboardSetLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardSetLocalizationResponse) SetIncluded(v []GameCenterLeaderboardSetLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardSetLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetLocalizationResponse := _GameCenterLeaderboardSetLocalizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationResponse)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationResponse(varGameCenterLeaderboardSetLocalizationResponse)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationResponse struct {
	value *GameCenterLeaderboardSetLocalizationResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationResponse) Get() *GameCenterLeaderboardSetLocalizationResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationResponse) Set(val *GameCenterLeaderboardSetLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationResponse(val *GameCenterLeaderboardSetLocalizationResponse) *NullableGameCenterLeaderboardSetLocalizationResponse {
	return &NullableGameCenterLeaderboardSetLocalizationResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


