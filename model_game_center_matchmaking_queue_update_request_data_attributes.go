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

// checks if the GameCenterMatchmakingQueueUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueUpdateRequestDataAttributes{}

// GameCenterMatchmakingQueueUpdateRequestDataAttributes struct for GameCenterMatchmakingQueueUpdateRequestDataAttributes
type GameCenterMatchmakingQueueUpdateRequestDataAttributes struct {
	ClassicMatchmakingBundleIds []string `json:"classicMatchmakingBundleIds,omitempty"`
}

// NewGameCenterMatchmakingQueueUpdateRequestDataAttributes instantiates a new GameCenterMatchmakingQueueUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueUpdateRequestDataAttributes() *GameCenterMatchmakingQueueUpdateRequestDataAttributes {
	this := GameCenterMatchmakingQueueUpdateRequestDataAttributes{}
	return &this
}

// NewGameCenterMatchmakingQueueUpdateRequestDataAttributesWithDefaults instantiates a new GameCenterMatchmakingQueueUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueUpdateRequestDataAttributesWithDefaults() *GameCenterMatchmakingQueueUpdateRequestDataAttributes {
	this := GameCenterMatchmakingQueueUpdateRequestDataAttributes{}
	return &this
}

// GetClassicMatchmakingBundleIds returns the ClassicMatchmakingBundleIds field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueUpdateRequestDataAttributes) GetClassicMatchmakingBundleIds() []string {
	if o == nil || IsNil(o.ClassicMatchmakingBundleIds) {
		var ret []string
		return ret
	}
	return o.ClassicMatchmakingBundleIds
}

// GetClassicMatchmakingBundleIdsOk returns a tuple with the ClassicMatchmakingBundleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueUpdateRequestDataAttributes) GetClassicMatchmakingBundleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassicMatchmakingBundleIds) {
		return nil, false
	}
	return o.ClassicMatchmakingBundleIds, true
}

// HasClassicMatchmakingBundleIds returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueUpdateRequestDataAttributes) HasClassicMatchmakingBundleIds() bool {
	if o != nil && !IsNil(o.ClassicMatchmakingBundleIds) {
		return true
	}

	return false
}

// SetClassicMatchmakingBundleIds gets a reference to the given []string and assigns it to the ClassicMatchmakingBundleIds field.
func (o *GameCenterMatchmakingQueueUpdateRequestDataAttributes) SetClassicMatchmakingBundleIds(v []string) {
	o.ClassicMatchmakingBundleIds = v
}

func (o GameCenterMatchmakingQueueUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClassicMatchmakingBundleIds) {
		toSerialize["classicMatchmakingBundleIds"] = o.ClassicMatchmakingBundleIds
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes struct {
	value *GameCenterMatchmakingQueueUpdateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) Get() *GameCenterMatchmakingQueueUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) Set(val *GameCenterMatchmakingQueueUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueUpdateRequestDataAttributes(val *GameCenterMatchmakingQueueUpdateRequestDataAttributes) *NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes {
	return &NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

