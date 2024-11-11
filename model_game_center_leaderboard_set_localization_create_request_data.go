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

// checks if the GameCenterLeaderboardSetLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationCreateRequestData{}

// GameCenterLeaderboardSetLocalizationCreateRequestData struct for GameCenterLeaderboardSetLocalizationCreateRequestData
type GameCenterLeaderboardSetLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships `json:"relationships"`
}

type _GameCenterLeaderboardSetLocalizationCreateRequestData GameCenterLeaderboardSetLocalizationCreateRequestData

// NewGameCenterLeaderboardSetLocalizationCreateRequestData instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationCreateRequestData(type_ string, attributes GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes, relationships GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) *GameCenterLeaderboardSetLocalizationCreateRequestData {
	this := GameCenterLeaderboardSetLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGameCenterLeaderboardSetLocalizationCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationCreateRequestDataWithDefaults() *GameCenterLeaderboardSetLocalizationCreateRequestData {
	this := GameCenterLeaderboardSetLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetAttributes() GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetAttributesOk() (*GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) SetAttributes(v GameCenterLeaderboardSetLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetRelationships() GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) SetRelationships(v GameCenterLeaderboardSetLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *GameCenterLeaderboardSetLocalizationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"relationships",
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

	varGameCenterLeaderboardSetLocalizationCreateRequestData := _GameCenterLeaderboardSetLocalizationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardSetLocalizationCreateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardSetLocalizationCreateRequestData(varGameCenterLeaderboardSetLocalizationCreateRequestData)

	return err
}

type NullableGameCenterLeaderboardSetLocalizationCreateRequestData struct {
	value *GameCenterLeaderboardSetLocalizationCreateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestData) Get() *GameCenterLeaderboardSetLocalizationCreateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestData) Set(val *GameCenterLeaderboardSetLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationCreateRequestData(val *GameCenterLeaderboardSetLocalizationCreateRequestData) *NullableGameCenterLeaderboardSetLocalizationCreateRequestData {
	return &NullableGameCenterLeaderboardSetLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


