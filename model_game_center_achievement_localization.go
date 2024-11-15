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

// checks if the GameCenterAchievementLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalization{}

// GameCenterAchievementLocalization struct for GameCenterAchievementLocalization
type GameCenterAchievementLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterAchievementLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *GameCenterAchievementLocalizationRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _GameCenterAchievementLocalization GameCenterAchievementLocalization

// NewGameCenterAchievementLocalization instantiates a new GameCenterAchievementLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalization(type_ string, id string) *GameCenterAchievementLocalization {
	this := GameCenterAchievementLocalization{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAchievementLocalizationWithDefaults instantiates a new GameCenterAchievementLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationWithDefaults() *GameCenterAchievementLocalization {
	this := GameCenterAchievementLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAchievementLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAchievementLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalization) GetAttributes() GameCenterAchievementLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAchievementLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalization) GetAttributesOk() (*GameCenterAchievementLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAchievementLocalizationAttributes and assigns it to the Attributes field.
func (o *GameCenterAchievementLocalization) SetAttributes(v GameCenterAchievementLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalization) GetRelationships() GameCenterAchievementLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterAchievementLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalization) GetRelationshipsOk() (*GameCenterAchievementLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterAchievementLocalizationRelationships and assigns it to the Relationships field.
func (o *GameCenterAchievementLocalization) SetRelationships(v GameCenterAchievementLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalization) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterAchievementLocalization) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterAchievementLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GameCenterAchievementLocalization) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterAchievementLocalization := _GameCenterAchievementLocalization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementLocalization)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementLocalization(varGameCenterAchievementLocalization)

	return err
}

type NullableGameCenterAchievementLocalization struct {
	value *GameCenterAchievementLocalization
	isSet bool
}

func (v NullableGameCenterAchievementLocalization) Get() *GameCenterAchievementLocalization {
	return v.value
}

func (v *NullableGameCenterAchievementLocalization) Set(val *GameCenterAchievementLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalization(val *GameCenterAchievementLocalization) *NullableGameCenterAchievementLocalization {
	return &NullableGameCenterAchievementLocalization{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


