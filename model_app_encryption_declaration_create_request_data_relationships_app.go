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

// checks if the AppEncryptionDeclarationCreateRequestDataRelationshipsApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationCreateRequestDataRelationshipsApp{}

// AppEncryptionDeclarationCreateRequestDataRelationshipsApp struct for AppEncryptionDeclarationCreateRequestDataRelationshipsApp
type AppEncryptionDeclarationCreateRequestDataRelationshipsApp struct {
	Data AlternativeDistributionKeyCreateRequestDataRelationshipsAppData `json:"data"`
}

type _AppEncryptionDeclarationCreateRequestDataRelationshipsApp AppEncryptionDeclarationCreateRequestDataRelationshipsApp

// NewAppEncryptionDeclarationCreateRequestDataRelationshipsApp instantiates a new AppEncryptionDeclarationCreateRequestDataRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationCreateRequestDataRelationshipsApp(data AlternativeDistributionKeyCreateRequestDataRelationshipsAppData) *AppEncryptionDeclarationCreateRequestDataRelationshipsApp {
	this := AppEncryptionDeclarationCreateRequestDataRelationshipsApp{}
	this.Data = data
	return &this
}

// NewAppEncryptionDeclarationCreateRequestDataRelationshipsAppWithDefaults instantiates a new AppEncryptionDeclarationCreateRequestDataRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationCreateRequestDataRelationshipsAppWithDefaults() *AppEncryptionDeclarationCreateRequestDataRelationshipsApp {
	this := AppEncryptionDeclarationCreateRequestDataRelationshipsApp{}
	return &this
}

// GetData returns the Data field value
func (o *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) GetData() AlternativeDistributionKeyCreateRequestDataRelationshipsAppData {
	if o == nil {
		var ret AlternativeDistributionKeyCreateRequestDataRelationshipsAppData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) GetDataOk() (*AlternativeDistributionKeyCreateRequestDataRelationshipsAppData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) SetData(v AlternativeDistributionKeyCreateRequestDataRelationshipsAppData) {
	o.Data = v
}

func (o AppEncryptionDeclarationCreateRequestDataRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationCreateRequestDataRelationshipsApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAppEncryptionDeclarationCreateRequestDataRelationshipsApp := _AppEncryptionDeclarationCreateRequestDataRelationshipsApp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationCreateRequestDataRelationshipsApp)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationCreateRequestDataRelationshipsApp(varAppEncryptionDeclarationCreateRequestDataRelationshipsApp)

	return err
}

type NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp struct {
	value *AppEncryptionDeclarationCreateRequestDataRelationshipsApp
	isSet bool
}

func (v NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) Get() *AppEncryptionDeclarationCreateRequestDataRelationshipsApp {
	return v.value
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) Set(val *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp(val *AppEncryptionDeclarationCreateRequestDataRelationshipsApp) *NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp {
	return &NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


