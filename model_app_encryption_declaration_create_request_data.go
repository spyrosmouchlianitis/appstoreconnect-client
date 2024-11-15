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

// checks if the AppEncryptionDeclarationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationCreateRequestData{}

// AppEncryptionDeclarationCreateRequestData struct for AppEncryptionDeclarationCreateRequestData
type AppEncryptionDeclarationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppEncryptionDeclarationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppEncryptionDeclarationCreateRequestDataRelationships `json:"relationships"`
}

type _AppEncryptionDeclarationCreateRequestData AppEncryptionDeclarationCreateRequestData

// NewAppEncryptionDeclarationCreateRequestData instantiates a new AppEncryptionDeclarationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationCreateRequestData(type_ string, attributes AppEncryptionDeclarationCreateRequestDataAttributes, relationships AppEncryptionDeclarationCreateRequestDataRelationships) *AppEncryptionDeclarationCreateRequestData {
	this := AppEncryptionDeclarationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppEncryptionDeclarationCreateRequestDataWithDefaults instantiates a new AppEncryptionDeclarationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationCreateRequestDataWithDefaults() *AppEncryptionDeclarationCreateRequestData {
	this := AppEncryptionDeclarationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEncryptionDeclarationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEncryptionDeclarationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppEncryptionDeclarationCreateRequestData) GetAttributes() AppEncryptionDeclarationCreateRequestDataAttributes {
	if o == nil {
		var ret AppEncryptionDeclarationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestData) GetAttributesOk() (*AppEncryptionDeclarationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppEncryptionDeclarationCreateRequestData) SetAttributes(v AppEncryptionDeclarationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppEncryptionDeclarationCreateRequestData) GetRelationships() AppEncryptionDeclarationCreateRequestDataRelationships {
	if o == nil {
		var ret AppEncryptionDeclarationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestData) GetRelationshipsOk() (*AppEncryptionDeclarationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppEncryptionDeclarationCreateRequestData) SetRelationships(v AppEncryptionDeclarationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppEncryptionDeclarationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppEncryptionDeclarationCreateRequestData := _AppEncryptionDeclarationCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationCreateRequestData)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationCreateRequestData(varAppEncryptionDeclarationCreateRequestData)

	return err
}

type NullableAppEncryptionDeclarationCreateRequestData struct {
	value *AppEncryptionDeclarationCreateRequestData
	isSet bool
}

func (v NullableAppEncryptionDeclarationCreateRequestData) Get() *AppEncryptionDeclarationCreateRequestData {
	return v.value
}

func (v *NullableAppEncryptionDeclarationCreateRequestData) Set(val *AppEncryptionDeclarationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationCreateRequestData(val *AppEncryptionDeclarationCreateRequestData) *NullableAppEncryptionDeclarationCreateRequestData {
	return &NullableAppEncryptionDeclarationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


