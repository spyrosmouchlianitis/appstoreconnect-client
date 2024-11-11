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

// checks if the GameCenterEnabledVersionCompatibleVersionsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterEnabledVersionCompatibleVersionsLinkagesResponse{}

// GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct for GameCenterEnabledVersionCompatibleVersionsLinkagesResponse
type GameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	Data []AppRelationshipsGameCenterEnabledVersionsDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _GameCenterEnabledVersionCompatibleVersionsLinkagesResponse GameCenterEnabledVersionCompatibleVersionsLinkagesResponse

// NewGameCenterEnabledVersionCompatibleVersionsLinkagesResponse instantiates a new GameCenterEnabledVersionCompatibleVersionsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterEnabledVersionCompatibleVersionsLinkagesResponse(data []AppRelationshipsGameCenterEnabledVersionsDataInner, links PagedDocumentLinks) *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse {
	this := GameCenterEnabledVersionCompatibleVersionsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterEnabledVersionCompatibleVersionsLinkagesResponseWithDefaults instantiates a new GameCenterEnabledVersionCompatibleVersionsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterEnabledVersionCompatibleVersionsLinkagesResponseWithDefaults() *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse {
	this := GameCenterEnabledVersionCompatibleVersionsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetData() []AppRelationshipsGameCenterEnabledVersionsDataInner {
	if o == nil {
		var ret []AppRelationshipsGameCenterEnabledVersionsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetDataOk() ([]AppRelationshipsGameCenterEnabledVersionsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) SetData(v []AppRelationshipsGameCenterEnabledVersionsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterEnabledVersionCompatibleVersionsLinkagesResponse := _GameCenterEnabledVersionCompatibleVersionsLinkagesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterEnabledVersionCompatibleVersionsLinkagesResponse)

	if err != nil {
		return err
	}

	*o = GameCenterEnabledVersionCompatibleVersionsLinkagesResponse(varGameCenterEnabledVersionCompatibleVersionsLinkagesResponse)

	return err
}

type NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse struct {
	value *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse
	isSet bool
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) Get() *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse {
	return v.value
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) Set(val *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse(val *GameCenterEnabledVersionCompatibleVersionsLinkagesResponse) *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse {
	return &NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse{value: val, isSet: true}
}

func (v NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterEnabledVersionCompatibleVersionsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


