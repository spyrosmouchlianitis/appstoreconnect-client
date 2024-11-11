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

// checks if the GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse{}

// GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse struct for GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse
type GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse struct {
	Data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse

// NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse instantiates a new GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse(data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, links PagedDocumentLinks) *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse {
	this := GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponseWithDefaults instantiates a new GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponseWithDefaults() *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse {
	this := GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse := _GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse(varGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse)

	return err
}

type NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse struct {
	value *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse
	isSet bool
}

func (v NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) Get() *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) Set(val *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse(val *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) *NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse {
	return &NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


