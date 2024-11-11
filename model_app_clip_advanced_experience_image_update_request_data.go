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

// checks if the AppClipAdvancedExperienceImageUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImageUpdateRequestData{}

// AppClipAdvancedExperienceImageUpdateRequestData struct for AppClipAdvancedExperienceImageUpdateRequestData
type AppClipAdvancedExperienceImageUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceImageUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _AppClipAdvancedExperienceImageUpdateRequestData AppClipAdvancedExperienceImageUpdateRequestData

// NewAppClipAdvancedExperienceImageUpdateRequestData instantiates a new AppClipAdvancedExperienceImageUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImageUpdateRequestData(type_ string, id string) *AppClipAdvancedExperienceImageUpdateRequestData {
	this := AppClipAdvancedExperienceImageUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipAdvancedExperienceImageUpdateRequestDataWithDefaults instantiates a new AppClipAdvancedExperienceImageUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageUpdateRequestDataWithDefaults() *AppClipAdvancedExperienceImageUpdateRequestData {
	this := AppClipAdvancedExperienceImageUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceImageUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipAdvancedExperienceImageUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetAttributes() AppClipAdvancedExperienceImageUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppClipAdvancedExperienceImageUpdateRequestData) SetAttributes(v AppClipAdvancedExperienceImageUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppClipAdvancedExperienceImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImageUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *AppClipAdvancedExperienceImageUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipAdvancedExperienceImageUpdateRequestData := _AppClipAdvancedExperienceImageUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipAdvancedExperienceImageUpdateRequestData)

	if err != nil {
		return err
	}

	*o = AppClipAdvancedExperienceImageUpdateRequestData(varAppClipAdvancedExperienceImageUpdateRequestData)

	return err
}

type NullableAppClipAdvancedExperienceImageUpdateRequestData struct {
	value *AppClipAdvancedExperienceImageUpdateRequestData
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequestData) Get() *AppClipAdvancedExperienceImageUpdateRequestData {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequestData) Set(val *AppClipAdvancedExperienceImageUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImageUpdateRequestData(val *AppClipAdvancedExperienceImageUpdateRequestData) *NullableAppClipAdvancedExperienceImageUpdateRequestData {
	return &NullableAppClipAdvancedExperienceImageUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


