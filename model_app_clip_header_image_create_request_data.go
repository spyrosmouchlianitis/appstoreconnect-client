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

// checks if the AppClipHeaderImageCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipHeaderImageCreateRequestData{}

// AppClipHeaderImageCreateRequestData struct for AppClipHeaderImageCreateRequestData
type AppClipHeaderImageCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
	Relationships AppClipHeaderImageCreateRequestDataRelationships `json:"relationships"`
}

type _AppClipHeaderImageCreateRequestData AppClipHeaderImageCreateRequestData

// NewAppClipHeaderImageCreateRequestData instantiates a new AppClipHeaderImageCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipHeaderImageCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships AppClipHeaderImageCreateRequestDataRelationships) *AppClipHeaderImageCreateRequestData {
	this := AppClipHeaderImageCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppClipHeaderImageCreateRequestDataWithDefaults instantiates a new AppClipHeaderImageCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipHeaderImageCreateRequestDataWithDefaults() *AppClipHeaderImageCreateRequestData {
	this := AppClipHeaderImageCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipHeaderImageCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipHeaderImageCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppClipHeaderImageCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppClipHeaderImageCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppClipHeaderImageCreateRequestData) GetRelationships() AppClipHeaderImageCreateRequestDataRelationships {
	if o == nil {
		var ret AppClipHeaderImageCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageCreateRequestData) GetRelationshipsOk() (*AppClipHeaderImageCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppClipHeaderImageCreateRequestData) SetRelationships(v AppClipHeaderImageCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppClipHeaderImageCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipHeaderImageCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AppClipHeaderImageCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppClipHeaderImageCreateRequestData := _AppClipHeaderImageCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppClipHeaderImageCreateRequestData)

	if err != nil {
		return err
	}

	*o = AppClipHeaderImageCreateRequestData(varAppClipHeaderImageCreateRequestData)

	return err
}

type NullableAppClipHeaderImageCreateRequestData struct {
	value *AppClipHeaderImageCreateRequestData
	isSet bool
}

func (v NullableAppClipHeaderImageCreateRequestData) Get() *AppClipHeaderImageCreateRequestData {
	return v.value
}

func (v *NullableAppClipHeaderImageCreateRequestData) Set(val *AppClipHeaderImageCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipHeaderImageCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipHeaderImageCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipHeaderImageCreateRequestData(val *AppClipHeaderImageCreateRequestData) *NullableAppClipHeaderImageCreateRequestData {
	return &NullableAppClipHeaderImageCreateRequestData{value: val, isSet: true}
}

func (v NullableAppClipHeaderImageCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipHeaderImageCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


