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

// checks if the GameCenterAchievementLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationCreateRequestDataAttributes{}

// GameCenterAchievementLocalizationCreateRequestDataAttributes struct for GameCenterAchievementLocalizationCreateRequestDataAttributes
type GameCenterAchievementLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
	Name string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription string `json:"afterEarnedDescription"`
}

type _GameCenterAchievementLocalizationCreateRequestDataAttributes GameCenterAchievementLocalizationCreateRequestDataAttributes

// NewGameCenterAchievementLocalizationCreateRequestDataAttributes instantiates a new GameCenterAchievementLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationCreateRequestDataAttributes(locale string, name string, beforeEarnedDescription string, afterEarnedDescription string) *GameCenterAchievementLocalizationCreateRequestDataAttributes {
	this := GameCenterAchievementLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	this.Name = name
	this.BeforeEarnedDescription = beforeEarnedDescription
	this.AfterEarnedDescription = afterEarnedDescription
	return &this
}

// NewGameCenterAchievementLocalizationCreateRequestDataAttributesWithDefaults instantiates a new GameCenterAchievementLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationCreateRequestDataAttributesWithDefaults() *GameCenterAchievementLocalizationCreateRequestDataAttributes {
	this := GameCenterAchievementLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetName returns the Name field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetBeforeEarnedDescription returns the BeforeEarnedDescription field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetBeforeEarnedDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeforeEarnedDescription
}

// GetBeforeEarnedDescriptionOk returns a tuple with the BeforeEarnedDescription field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetBeforeEarnedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeforeEarnedDescription, true
}

// SetBeforeEarnedDescription sets field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) SetBeforeEarnedDescription(v string) {
	o.BeforeEarnedDescription = v
}

// GetAfterEarnedDescription returns the AfterEarnedDescription field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetAfterEarnedDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfterEarnedDescription
}

// GetAfterEarnedDescriptionOk returns a tuple with the AfterEarnedDescription field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) GetAfterEarnedDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfterEarnedDescription, true
}

// SetAfterEarnedDescription sets field value
func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) SetAfterEarnedDescription(v string) {
	o.AfterEarnedDescription = v
}

func (o GameCenterAchievementLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	toSerialize["name"] = o.Name
	toSerialize["beforeEarnedDescription"] = o.BeforeEarnedDescription
	toSerialize["afterEarnedDescription"] = o.AfterEarnedDescription
	return toSerialize, nil
}

func (o *GameCenterAchievementLocalizationCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"locale",
		"name",
		"beforeEarnedDescription",
		"afterEarnedDescription",
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

	varGameCenterAchievementLocalizationCreateRequestDataAttributes := _GameCenterAchievementLocalizationCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameCenterAchievementLocalizationCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = GameCenterAchievementLocalizationCreateRequestDataAttributes(varGameCenterAchievementLocalizationCreateRequestDataAttributes)

	return err
}

type NullableGameCenterAchievementLocalizationCreateRequestDataAttributes struct {
	value *GameCenterAchievementLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) Get() *GameCenterAchievementLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) Set(val *GameCenterAchievementLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationCreateRequestDataAttributes(val *GameCenterAchievementLocalizationCreateRequestDataAttributes) *NullableGameCenterAchievementLocalizationCreateRequestDataAttributes {
	return &NullableGameCenterAchievementLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


