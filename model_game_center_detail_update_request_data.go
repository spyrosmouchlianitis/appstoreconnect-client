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

// checks if the GameCenterDetailUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailUpdateRequestData{}

// GameCenterDetailUpdateRequestData struct for GameCenterDetailUpdateRequestData
type GameCenterDetailUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterDetailCreateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *GameCenterDetailUpdateRequestDataRelationships `json:"relationships,omitempty"`
}

type _GameCenterDetailUpdateRequestData GameCenterDetailUpdateRequestData

// NewGameCenterDetailUpdateRequestData instantiates a new GameCenterDetailUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailUpdateRequestData(type_ string, id string) *GameCenterDetailUpdateRequestData {
	this := GameCenterDetailUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterDetailUpdateRequestDataWithDefaults instantiates a new GameCenterDetailUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailUpdateRequestDataWithDefaults() *GameCenterDetailUpdateRequestData {
	this := GameCenterDetailUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetailUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetailUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterDetailUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterDetailUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterDetailUpdateRequestData) GetAttributes() GameCenterDetailCreateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterDetailCreateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestData) GetAttributesOk() (*GameCenterDetailCreateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterDetailUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterDetailCreateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterDetailUpdateRequestData) SetAttributes(v GameCenterDetailCreateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterDetailUpdateRequestData) GetRelationships() GameCenterDetailUpdateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterDetailUpdateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailUpdateRequestData) GetRelationshipsOk() (*GameCenterDetailUpdateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterDetailUpdateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterDetailUpdateRequestDataRelationships and assigns it to the Relationships field.
func (o *GameCenterDetailUpdateRequestData) SetRelationships(v GameCenterDetailUpdateRequestDataRelationships) {
	o.Relationships = &v
}

func (o GameCenterDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *GameCenterDetailUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterDetailUpdateRequestData := _GameCenterDetailUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterDetailUpdateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterDetailUpdateRequestData(varGameCenterDetailUpdateRequestData)

	return err
}

type NullableGameCenterDetailUpdateRequestData struct {
	value *GameCenterDetailUpdateRequestData
	isSet bool
}

func (v NullableGameCenterDetailUpdateRequestData) Get() *GameCenterDetailUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterDetailUpdateRequestData) Set(val *GameCenterDetailUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailUpdateRequestData(val *GameCenterDetailUpdateRequestData) *NullableGameCenterDetailUpdateRequestData {
	return &NullableGameCenterDetailUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


