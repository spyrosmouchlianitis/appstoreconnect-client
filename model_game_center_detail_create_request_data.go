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

// checks if the GameCenterDetailCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailCreateRequestData{}

// GameCenterDetailCreateRequestData struct for GameCenterDetailCreateRequestData
type GameCenterDetailCreateRequestData struct {
	Type string `json:"type"`
	Attributes *GameCenterDetailCreateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships AnalyticsReportRequestCreateRequestDataRelationships `json:"relationships"`
}

type _GameCenterDetailCreateRequestData GameCenterDetailCreateRequestData

// NewGameCenterDetailCreateRequestData instantiates a new GameCenterDetailCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailCreateRequestData(type_ string, relationships AnalyticsReportRequestCreateRequestDataRelationships) *GameCenterDetailCreateRequestData {
	this := GameCenterDetailCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewGameCenterDetailCreateRequestDataWithDefaults instantiates a new GameCenterDetailCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailCreateRequestDataWithDefaults() *GameCenterDetailCreateRequestData {
	this := GameCenterDetailCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetailCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetailCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterDetailCreateRequestData) GetAttributes() GameCenterDetailCreateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterDetailCreateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailCreateRequestData) GetAttributesOk() (*GameCenterDetailCreateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterDetailCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterDetailCreateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterDetailCreateRequestData) SetAttributes(v GameCenterDetailCreateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterDetailCreateRequestData) GetRelationships() AnalyticsReportRequestCreateRequestDataRelationships {
	if o == nil {
		var ret AnalyticsReportRequestCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailCreateRequestData) GetRelationshipsOk() (*AnalyticsReportRequestCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterDetailCreateRequestData) SetRelationships(v AnalyticsReportRequestCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *GameCenterDetailCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varGameCenterDetailCreateRequestData := _GameCenterDetailCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterDetailCreateRequestData)

	if err != nil {
		return err
	}

	*o = GameCenterDetailCreateRequestData(varGameCenterDetailCreateRequestData)

	return err
}

type NullableGameCenterDetailCreateRequestData struct {
	value *GameCenterDetailCreateRequestData
	isSet bool
}

func (v NullableGameCenterDetailCreateRequestData) Get() *GameCenterDetailCreateRequestData {
	return v.value
}

func (v *NullableGameCenterDetailCreateRequestData) Set(val *GameCenterDetailCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailCreateRequestData(val *GameCenterDetailCreateRequestData) *NullableGameCenterDetailCreateRequestData {
	return &NullableGameCenterDetailCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


