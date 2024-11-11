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

// checks if the GameCenterMatchmakingRuleSetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetsResponse{}

// GameCenterMatchmakingRuleSetsResponse struct for GameCenterMatchmakingRuleSetsResponse
type GameCenterMatchmakingRuleSetsResponse struct {
	Data []GameCenterMatchmakingRuleSet `json:"data"`
	Included []GameCenterMatchmakingRuleSetsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

type _GameCenterMatchmakingRuleSetsResponse GameCenterMatchmakingRuleSetsResponse

// NewGameCenterMatchmakingRuleSetsResponse instantiates a new GameCenterMatchmakingRuleSetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetsResponse(data []GameCenterMatchmakingRuleSet, links PagedDocumentLinks) *GameCenterMatchmakingRuleSetsResponse {
	this := GameCenterMatchmakingRuleSetsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGameCenterMatchmakingRuleSetsResponseWithDefaults instantiates a new GameCenterMatchmakingRuleSetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetsResponseWithDefaults() *GameCenterMatchmakingRuleSetsResponse {
	this := GameCenterMatchmakingRuleSetsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterMatchmakingRuleSetsResponse) GetData() []GameCenterMatchmakingRuleSet {
	if o == nil {
		var ret []GameCenterMatchmakingRuleSet
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) GetDataOk() ([]GameCenterMatchmakingRuleSet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterMatchmakingRuleSetsResponse) SetData(v []GameCenterMatchmakingRuleSet) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetsResponse) GetIncluded() []GameCenterMatchmakingRuleSetsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GameCenterMatchmakingRuleSetsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) GetIncludedOk() ([]GameCenterMatchmakingRuleSetsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GameCenterMatchmakingRuleSetsResponseIncludedInner and assigns it to the Included field.
func (o *GameCenterMatchmakingRuleSetsResponse) SetIncluded(v []GameCenterMatchmakingRuleSetsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *GameCenterMatchmakingRuleSetsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GameCenterMatchmakingRuleSetsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterMatchmakingRuleSetsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o GameCenterMatchmakingRuleSetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

func (o *GameCenterMatchmakingRuleSetsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterMatchmakingRuleSetsResponse := _GameCenterMatchmakingRuleSetsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterMatchmakingRuleSetsResponse)

	if err != nil {
		return err
	}

	*o = GameCenterMatchmakingRuleSetsResponse(varGameCenterMatchmakingRuleSetsResponse)

	return err
}

type NullableGameCenterMatchmakingRuleSetsResponse struct {
	value *GameCenterMatchmakingRuleSetsResponse
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetsResponse) Get() *GameCenterMatchmakingRuleSetsResponse {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetsResponse) Set(val *GameCenterMatchmakingRuleSetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetsResponse(val *GameCenterMatchmakingRuleSetsResponse) *NullableGameCenterMatchmakingRuleSetsResponse {
	return &NullableGameCenterMatchmakingRuleSetsResponse{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


