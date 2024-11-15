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

// checks if the GameCenterLeaderboardSetMemberLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetMemberLocalizationResponse{}

// GameCenterLeaderboardSetMemberLocalizationResponse struct for GameCenterLeaderboardSetMemberLocalizationResponse
type GameCenterLeaderboardSetMemberLocalizationResponse struct {
	Data GameCenterLeaderboardSetMemberLocalization `json:"data"`
	Included []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

type _GameCenterLeaderboardSetMemberLocalizationResponse GameCenterLeaderboardSetMemberLocalizationResponse

// NewGameCenterLeaderboardSetMemberLocalizationResponse instantiates a new GameCenterLeaderboardSetMemberLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetMemberLocalizationResponse(data GameCenterLeaderboardSetMemberLocalization, links DocumentLinks) *GameCenterLeaderboardSetMemberLocalizationResponse {
	this := GameCenterLeaderboardSetMemberLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetMemberLocalizationResponseWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetMemberLocalizationResponseWithDefaults() *GameCenterLeaderboardSetMemberLocalizationResponse {
	this := GameCenterLeaderboardSetMemberLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetData() GameCenterLeaderboardSetMemberLocalization {
	if o == nil {
		var ret GameCenterLeaderboardSetMemberLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetDataOk() (*GameCenterLeaderboardSetMemberLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) SetData(v GameCenterLeaderboardSetMemberLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetIncluded() []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetIncludedOk() ([]GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) SetIncluded(v []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetMemberLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o GameCenterLeaderboardSetMemberLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetMemberLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetMemberLocalizationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetMemberLocalizationResponse := _GameCenterLeaderboardSetMemberLocalizationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetMemberLocalizationResponse)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetMemberLocalizationResponse(varGameCenterLeaderboardSetMemberLocalizationResponse)

	return err
}

type NullableGameCenterLeaderboardSetMemberLocalizationResponse struct {
	value *GameCenterLeaderboardSetMemberLocalizationResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationResponse) Get() *GameCenterLeaderboardSetMemberLocalizationResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationResponse) Set(val *GameCenterLeaderboardSetMemberLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetMemberLocalizationResponse(val *GameCenterLeaderboardSetMemberLocalizationResponse) *NullableGameCenterLeaderboardSetMemberLocalizationResponse {
	return &NullableGameCenterLeaderboardSetMemberLocalizationResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


