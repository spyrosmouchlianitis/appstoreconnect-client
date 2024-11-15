/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the GameCenterLeaderboardLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationUpdateRequestDataAttributes{}

// GameCenterLeaderboardLocalizationUpdateRequestDataAttributes struct for GameCenterLeaderboardLocalizationUpdateRequestDataAttributes
type GameCenterLeaderboardLocalizationUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	FormatterOverride *GameCenterLeaderboardFormatter `json:"formatterOverride,omitempty"`
	FormatterSuffix *string `json:"formatterSuffix,omitempty"`
	FormatterSuffixSingular *string `json:"formatterSuffixSingular,omitempty"`
}

// NewGameCenterLeaderboardLocalizationUpdateRequestDataAttributes instantiates a new GameCenterLeaderboardLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationUpdateRequestDataAttributes() *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes {
	this := GameCenterLeaderboardLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewGameCenterLeaderboardLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationUpdateRequestDataAttributesWithDefaults() *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes {
	this := GameCenterLeaderboardLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetFormatterOverride returns the FormatterOverride field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterOverride() GameCenterLeaderboardFormatter {
	if o == nil || IsNil(o.FormatterOverride) {
		var ret GameCenterLeaderboardFormatter
		return ret
	}
	return *o.FormatterOverride
}

// GetFormatterOverrideOk returns a tuple with the FormatterOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterOverrideOk() (*GameCenterLeaderboardFormatter, bool) {
	if o == nil || IsNil(o.FormatterOverride) {
		return nil, false
	}
	return o.FormatterOverride, true
}

// HasFormatterOverride returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) HasFormatterOverride() bool {
	if o != nil && !IsNil(o.FormatterOverride) {
		return true
	}

	return false
}

// SetFormatterOverride gets a reference to the given GameCenterLeaderboardFormatter and assigns it to the FormatterOverride field.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) SetFormatterOverride(v GameCenterLeaderboardFormatter) {
	o.FormatterOverride = &v
}

// GetFormatterSuffix returns the FormatterSuffix field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterSuffix() string {
	if o == nil || IsNil(o.FormatterSuffix) {
		var ret string
		return ret
	}
	return *o.FormatterSuffix
}

// GetFormatterSuffixOk returns a tuple with the FormatterSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.FormatterSuffix) {
		return nil, false
	}
	return o.FormatterSuffix, true
}

// HasFormatterSuffix returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) HasFormatterSuffix() bool {
	if o != nil && !IsNil(o.FormatterSuffix) {
		return true
	}

	return false
}

// SetFormatterSuffix gets a reference to the given string and assigns it to the FormatterSuffix field.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) SetFormatterSuffix(v string) {
	o.FormatterSuffix = &v
}

// GetFormatterSuffixSingular returns the FormatterSuffixSingular field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterSuffixSingular() string {
	if o == nil || IsNil(o.FormatterSuffixSingular) {
		var ret string
		return ret
	}
	return *o.FormatterSuffixSingular
}

// GetFormatterSuffixSingularOk returns a tuple with the FormatterSuffixSingular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) GetFormatterSuffixSingularOk() (*string, bool) {
	if o == nil || IsNil(o.FormatterSuffixSingular) {
		return nil, false
	}
	return o.FormatterSuffixSingular, true
}

// HasFormatterSuffixSingular returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) HasFormatterSuffixSingular() bool {
	if o != nil && !IsNil(o.FormatterSuffixSingular) {
		return true
	}

	return false
}

// SetFormatterSuffixSingular gets a reference to the given string and assigns it to the FormatterSuffixSingular field.
func (o *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) SetFormatterSuffixSingular(v string) {
	o.FormatterSuffixSingular = &v
}

func (o GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FormatterOverride) {
		toSerialize["formatterOverride"] = o.FormatterOverride
	}
	if !IsNil(o.FormatterSuffix) {
		toSerialize["formatterSuffix"] = o.FormatterSuffix
	}
	if !IsNil(o.FormatterSuffixSingular) {
		toSerialize["formatterSuffixSingular"] = o.FormatterSuffixSingular
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes struct {
	value *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) Get() *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) Set(val *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes(val *GameCenterLeaderboardLocalizationUpdateRequestDataAttributes) *NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes {
	return &NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


