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

// checks if the AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData{}

// AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData struct for AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData
type AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData

// NewAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData instantiates a new AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData(type_ string, id string) *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData {
	this := AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentDataWithDefaults instantiates a new AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentDataWithDefaults() *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData {
	this := AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) SetId(v string) {
	o.Id = v
}

func (o AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) UnmarshalJSON(data []byte) (err error) {
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

	varAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData := _AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData(varAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData)

	return err
}

type NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData struct {
	value *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData
	isSet bool
}

func (v NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) Get() *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData {
	return v.value
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) Set(val *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData(val *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) *NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData {
	return &NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocumentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


