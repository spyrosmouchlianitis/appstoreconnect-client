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

// checks if the GameCenterAchievementRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementRelease{}

// GameCenterAchievementRelease struct for GameCenterAchievementRelease
type GameCenterAchievementRelease struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterAchievementReleaseAttributes `json:"attributes,omitempty"`
	Relationships *GameCenterAchievementReleaseRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

type _GameCenterAchievementRelease GameCenterAchievementRelease

// NewGameCenterAchievementRelease instantiates a new GameCenterAchievementRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementRelease(type_ string, id string) *GameCenterAchievementRelease {
	this := GameCenterAchievementRelease{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAchievementReleaseWithDefaults instantiates a new GameCenterAchievementRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleaseWithDefaults() *GameCenterAchievementRelease {
	this := GameCenterAchievementRelease{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementRelease) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelease) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementRelease) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAchievementRelease) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelease) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAchievementRelease) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterAchievementRelease) GetAttributes() GameCenterAchievementReleaseAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAchievementReleaseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterAchievementRelease) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAchievementReleaseAttributes and assigns it to the Attributes field.
func (o *GameCenterAchievementRelease) SetAttributes(v GameCenterAchievementReleaseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterAchievementRelease) GetRelationships() GameCenterAchievementReleaseRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterAchievementReleaseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelease) GetRelationshipsOk() (*GameCenterAchievementReleaseRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterAchievementRelease) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterAchievementReleaseRelationships and assigns it to the Relationships field.
func (o *GameCenterAchievementRelease) SetRelationships(v GameCenterAchievementReleaseRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterAchievementRelease) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementRelease) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterAchievementRelease) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterAchievementRelease) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterAchievementRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementRelease) ToMap() (map[string]interface{}, error) {
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

func (o *GameCenterAchievementRelease) UnmarshalJSON(data []byte) (err error) {
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

	varGameCenterAchievementRelease := _GameCenterAchievementRelease{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementRelease)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementRelease(varGameCenterAchievementRelease)

	return err
}

type NullableGameCenterAchievementRelease struct {
	value *GameCenterAchievementRelease
	isSet bool
}

func (v NullableGameCenterAchievementRelease) Get() *GameCenterAchievementRelease {
	return v.value
}

func (v *NullableGameCenterAchievementRelease) Set(val *GameCenterAchievementRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementRelease(val *GameCenterAchievementRelease) *NullableGameCenterAchievementRelease {
	return &NullableGameCenterAchievementRelease{value: val, isSet: true}
}

func (v NullableGameCenterAchievementRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


