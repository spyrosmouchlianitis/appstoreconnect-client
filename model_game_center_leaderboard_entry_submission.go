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

// checks if the GameCenterLeaderboardEntrySubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardEntrySubmission{}

// GameCenterLeaderboardEntrySubmission struct for GameCenterLeaderboardEntrySubmission
type GameCenterLeaderboardEntrySubmission struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterLeaderboardEntrySubmissionAttributes `json:"attributes,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _GameCenterLeaderboardEntrySubmission GameCenterLeaderboardEntrySubmission

// NewGameCenterLeaderboardEntrySubmission instantiates a new GameCenterLeaderboardEntrySubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardEntrySubmission(type_ string, id string) *GameCenterLeaderboardEntrySubmission {
	this := GameCenterLeaderboardEntrySubmission{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardEntrySubmissionWithDefaults instantiates a new GameCenterLeaderboardEntrySubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardEntrySubmissionWithDefaults() *GameCenterLeaderboardEntrySubmission {
	this := GameCenterLeaderboardEntrySubmission{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardEntrySubmission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardEntrySubmission) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardEntrySubmission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardEntrySubmission) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardEntrySubmission) GetAttributes() GameCenterLeaderboardEntrySubmissionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterLeaderboardEntrySubmissionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmission) GetAttributesOk() (*GameCenterLeaderboardEntrySubmissionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardEntrySubmission) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterLeaderboardEntrySubmissionAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardEntrySubmission) SetAttributes(v GameCenterLeaderboardEntrySubmissionAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardEntrySubmission) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardEntrySubmission) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardEntrySubmission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardEntrySubmission) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterLeaderboardEntrySubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardEntrySubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GameCenterLeaderboardEntrySubmission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGameCenterLeaderboardEntrySubmission := _GameCenterLeaderboardEntrySubmission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterLeaderboardEntrySubmission)

	if err != nil {
		return err
	}

	*o = GameCenterLeaderboardEntrySubmission(varGameCenterLeaderboardEntrySubmission)

	return err
}

type NullableGameCenterLeaderboardEntrySubmission struct {
	value *GameCenterLeaderboardEntrySubmission
	isSet bool
}

func (v NullableGameCenterLeaderboardEntrySubmission) Get() *GameCenterLeaderboardEntrySubmission {
	return v.value
}

func (v *NullableGameCenterLeaderboardEntrySubmission) Set(val *GameCenterLeaderboardEntrySubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardEntrySubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardEntrySubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardEntrySubmission(val *GameCenterLeaderboardEntrySubmission) *NullableGameCenterLeaderboardEntrySubmission {
	return &NullableGameCenterLeaderboardEntrySubmission{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardEntrySubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardEntrySubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


