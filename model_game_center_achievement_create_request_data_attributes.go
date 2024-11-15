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

// checks if the GameCenterAchievementCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementCreateRequestDataAttributes{}

// GameCenterAchievementCreateRequestDataAttributes struct for GameCenterAchievementCreateRequestDataAttributes
type GameCenterAchievementCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points int32 `json:"points"`
	ShowBeforeEarned bool `json:"showBeforeEarned"`
	Repeatable bool `json:"repeatable"`
}

type _GameCenterAchievementCreateRequestDataAttributes GameCenterAchievementCreateRequestDataAttributes

// NewGameCenterAchievementCreateRequestDataAttributes instantiates a new GameCenterAchievementCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementCreateRequestDataAttributes(referenceName string, vendorIdentifier string, points int32, showBeforeEarned bool, repeatable bool) *GameCenterAchievementCreateRequestDataAttributes {
	this := GameCenterAchievementCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	this.VendorIdentifier = vendorIdentifier
	this.Points = points
	this.ShowBeforeEarned = showBeforeEarned
	this.Repeatable = repeatable
	return &this
}

// NewGameCenterAchievementCreateRequestDataAttributesWithDefaults instantiates a new GameCenterAchievementCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementCreateRequestDataAttributesWithDefaults() *GameCenterAchievementCreateRequestDataAttributes {
	this := GameCenterAchievementCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *GameCenterAchievementCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GameCenterAchievementCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetVendorIdentifier returns the VendorIdentifier field value
func (o *GameCenterAchievementCreateRequestDataAttributes) GetVendorIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorIdentifier, true
}

// SetVendorIdentifier sets field value
func (o *GameCenterAchievementCreateRequestDataAttributes) SetVendorIdentifier(v string) {
	o.VendorIdentifier = v
}

// GetPoints returns the Points field value
func (o *GameCenterAchievementCreateRequestDataAttributes) GetPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataAttributes) GetPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value
func (o *GameCenterAchievementCreateRequestDataAttributes) SetPoints(v int32) {
	o.Points = v
}

// GetShowBeforeEarned returns the ShowBeforeEarned field value
func (o *GameCenterAchievementCreateRequestDataAttributes) GetShowBeforeEarned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowBeforeEarned
}

// GetShowBeforeEarnedOk returns a tuple with the ShowBeforeEarned field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataAttributes) GetShowBeforeEarnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowBeforeEarned, true
}

// SetShowBeforeEarned sets field value
func (o *GameCenterAchievementCreateRequestDataAttributes) SetShowBeforeEarned(v bool) {
	o.ShowBeforeEarned = v
}

// GetRepeatable returns the Repeatable field value
func (o *GameCenterAchievementCreateRequestDataAttributes) GetRepeatable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Repeatable
}

// GetRepeatableOk returns a tuple with the Repeatable field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataAttributes) GetRepeatableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repeatable, true
}

// SetRepeatable sets field value
func (o *GameCenterAchievementCreateRequestDataAttributes) SetRepeatable(v bool) {
	o.Repeatable = v
}

func (o GameCenterAchievementCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["vendorIdentifier"] = o.VendorIdentifier
	toSerialize["points"] = o.Points
	toSerialize["showBeforeEarned"] = o.ShowBeforeEarned
	toSerialize["repeatable"] = o.Repeatable
	return toSerialize, nil
}

func (o *GameCenterAchievementCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceName",
		"vendorIdentifier",
		"points",
		"showBeforeEarned",
		"repeatable",
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

	varGameCenterAchievementCreateRequestDataAttributes := _GameCenterAchievementCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementCreateRequestDataAttributes(varGameCenterAchievementCreateRequestDataAttributes)

	return err
}

type NullableGameCenterAchievementCreateRequestDataAttributes struct {
	value *GameCenterAchievementCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterAchievementCreateRequestDataAttributes) Get() *GameCenterAchievementCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterAchievementCreateRequestDataAttributes) Set(val *GameCenterAchievementCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementCreateRequestDataAttributes(val *GameCenterAchievementCreateRequestDataAttributes) *NullableGameCenterAchievementCreateRequestDataAttributes {
	return &NullableGameCenterAchievementCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAchievementCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


