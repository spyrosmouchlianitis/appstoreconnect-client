/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the GameCenterMatchmakingQueueAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueAttributes{}

// GameCenterMatchmakingQueueAttributes struct for GameCenterMatchmakingQueueAttributes
type GameCenterMatchmakingQueueAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
	ClassicMatchmakingBundleIds []string `json:"classicMatchmakingBundleIds,omitempty"`
}

// NewGameCenterMatchmakingQueueAttributes instantiates a new GameCenterMatchmakingQueueAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueAttributes() *GameCenterMatchmakingQueueAttributes {
	this := GameCenterMatchmakingQueueAttributes{}
	return &this
}

// NewGameCenterMatchmakingQueueAttributesWithDefaults instantiates a new GameCenterMatchmakingQueueAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueAttributesWithDefaults() *GameCenterMatchmakingQueueAttributes {
	this := GameCenterMatchmakingQueueAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterMatchmakingQueueAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetClassicMatchmakingBundleIds returns the ClassicMatchmakingBundleIds field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueAttributes) GetClassicMatchmakingBundleIds() []string {
	if o == nil || IsNil(o.ClassicMatchmakingBundleIds) {
		var ret []string
		return ret
	}
	return o.ClassicMatchmakingBundleIds
}

// GetClassicMatchmakingBundleIdsOk returns a tuple with the ClassicMatchmakingBundleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueAttributes) GetClassicMatchmakingBundleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassicMatchmakingBundleIds) {
		return nil, false
	}
	return o.ClassicMatchmakingBundleIds, true
}

// HasClassicMatchmakingBundleIds returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueAttributes) HasClassicMatchmakingBundleIds() bool {
	if o != nil && !IsNil(o.ClassicMatchmakingBundleIds) {
		return true
	}

	return false
}

// SetClassicMatchmakingBundleIds gets a reference to the given []string and assigns it to the ClassicMatchmakingBundleIds field.
func (o *GameCenterMatchmakingQueueAttributes) SetClassicMatchmakingBundleIds(v []string) {
	o.ClassicMatchmakingBundleIds = v
}

func (o GameCenterMatchmakingQueueAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.ClassicMatchmakingBundleIds) {
		toSerialize["classicMatchmakingBundleIds"] = o.ClassicMatchmakingBundleIds
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueAttributes struct {
	value *GameCenterMatchmakingQueueAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueAttributes) Get() *GameCenterMatchmakingQueueAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueAttributes) Set(val *GameCenterMatchmakingQueueAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueAttributes(val *GameCenterMatchmakingQueueAttributes) *NullableGameCenterMatchmakingQueueAttributes {
	return &NullableGameCenterMatchmakingQueueAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


