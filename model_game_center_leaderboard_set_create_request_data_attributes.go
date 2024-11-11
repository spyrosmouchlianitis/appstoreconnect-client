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

// checks if the GameCenterLeaderboardSetCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetCreateRequestDataAttributes{}

// GameCenterLeaderboardSetCreateRequestDataAttributes struct for GameCenterLeaderboardSetCreateRequestDataAttributes
type GameCenterLeaderboardSetCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
}

type _GameCenterLeaderboardSetCreateRequestDataAttributes GameCenterLeaderboardSetCreateRequestDataAttributes

// NewGameCenterLeaderboardSetCreateRequestDataAttributes instantiates a new GameCenterLeaderboardSetCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetCreateRequestDataAttributes(referenceName string, vendorIdentifier string) *GameCenterLeaderboardSetCreateRequestDataAttributes {
	this := GameCenterLeaderboardSetCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	this.VendorIdentifier = vendorIdentifier
	return &this
}

// NewGameCenterLeaderboardSetCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardSetCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardSetCreateRequestDataAttributes {
	this := GameCenterLeaderboardSetCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetVendorIdentifier returns the VendorIdentifier field value
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) GetVendorIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorIdentifier, true
}

// SetVendorIdentifier sets field value
func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) SetVendorIdentifier(v string) {
	o.VendorIdentifier = v
}

func (o GameCenterLeaderboardSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["vendorIdentifier"] = o.VendorIdentifier
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceName",
		"vendorIdentifier",
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

	varGameCenterLeaderboardSetCreateRequestDataAttributes := _GameCenterLeaderboardSetCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetCreateRequestDataAttributes(varGameCenterLeaderboardSetCreateRequestDataAttributes)

	return err
}

type NullableGameCenterLeaderboardSetCreateRequestDataAttributes struct {
	value *GameCenterLeaderboardSetCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardSetCreateRequestDataAttributes) Get() *GameCenterLeaderboardSetCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetCreateRequestDataAttributes) Set(val *GameCenterLeaderboardSetCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetCreateRequestDataAttributes(val *GameCenterLeaderboardSetCreateRequestDataAttributes) *NullableGameCenterLeaderboardSetCreateRequestDataAttributes {
	return &NullableGameCenterLeaderboardSetCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


