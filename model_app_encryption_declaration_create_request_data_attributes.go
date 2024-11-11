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

// checks if the AppEncryptionDeclarationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationCreateRequestDataAttributes{}

// AppEncryptionDeclarationCreateRequestDataAttributes struct for AppEncryptionDeclarationCreateRequestDataAttributes
type AppEncryptionDeclarationCreateRequestDataAttributes struct {
	AppDescription string `json:"appDescription"`
	ContainsProprietaryCryptography bool `json:"containsProprietaryCryptography"`
	ContainsThirdPartyCryptography bool `json:"containsThirdPartyCryptography"`
	AvailableOnFrenchStore bool `json:"availableOnFrenchStore"`
}

type _AppEncryptionDeclarationCreateRequestDataAttributes AppEncryptionDeclarationCreateRequestDataAttributes

// NewAppEncryptionDeclarationCreateRequestDataAttributes instantiates a new AppEncryptionDeclarationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationCreateRequestDataAttributes(appDescription string, containsProprietaryCryptography bool, containsThirdPartyCryptography bool, availableOnFrenchStore bool) *AppEncryptionDeclarationCreateRequestDataAttributes {
	this := AppEncryptionDeclarationCreateRequestDataAttributes{}
	this.AppDescription = appDescription
	this.ContainsProprietaryCryptography = containsProprietaryCryptography
	this.ContainsThirdPartyCryptography = containsThirdPartyCryptography
	this.AvailableOnFrenchStore = availableOnFrenchStore
	return &this
}

// NewAppEncryptionDeclarationCreateRequestDataAttributesWithDefaults instantiates a new AppEncryptionDeclarationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationCreateRequestDataAttributesWithDefaults() *AppEncryptionDeclarationCreateRequestDataAttributes {
	this := AppEncryptionDeclarationCreateRequestDataAttributes{}
	return &this
}

// GetAppDescription returns the AppDescription field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetAppDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppDescription
}

// GetAppDescriptionOk returns a tuple with the AppDescription field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetAppDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppDescription, true
}

// SetAppDescription sets field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) SetAppDescription(v string) {
	o.AppDescription = v
}

// GetContainsProprietaryCryptography returns the ContainsProprietaryCryptography field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetContainsProprietaryCryptography() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContainsProprietaryCryptography
}

// GetContainsProprietaryCryptographyOk returns a tuple with the ContainsProprietaryCryptography field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetContainsProprietaryCryptographyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainsProprietaryCryptography, true
}

// SetContainsProprietaryCryptography sets field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) SetContainsProprietaryCryptography(v bool) {
	o.ContainsProprietaryCryptography = v
}

// GetContainsThirdPartyCryptography returns the ContainsThirdPartyCryptography field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetContainsThirdPartyCryptography() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContainsThirdPartyCryptography
}

// GetContainsThirdPartyCryptographyOk returns a tuple with the ContainsThirdPartyCryptography field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetContainsThirdPartyCryptographyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainsThirdPartyCryptography, true
}

// SetContainsThirdPartyCryptography sets field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) SetContainsThirdPartyCryptography(v bool) {
	o.ContainsThirdPartyCryptography = v
}

// GetAvailableOnFrenchStore returns the AvailableOnFrenchStore field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetAvailableOnFrenchStore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AvailableOnFrenchStore
}

// GetAvailableOnFrenchStoreOk returns a tuple with the AvailableOnFrenchStore field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) GetAvailableOnFrenchStoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableOnFrenchStore, true
}

// SetAvailableOnFrenchStore sets field value
func (o *AppEncryptionDeclarationCreateRequestDataAttributes) SetAvailableOnFrenchStore(v bool) {
	o.AvailableOnFrenchStore = v
}

func (o AppEncryptionDeclarationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appDescription"] = o.AppDescription
	toSerialize["containsProprietaryCryptography"] = o.ContainsProprietaryCryptography
	toSerialize["containsThirdPartyCryptography"] = o.ContainsThirdPartyCryptography
	toSerialize["availableOnFrenchStore"] = o.AvailableOnFrenchStore
	return toSerialize, nil
}

func (o *AppEncryptionDeclarationCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appDescription",
		"containsProprietaryCryptography",
		"containsThirdPartyCryptography",
		"availableOnFrenchStore",
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

	varAppEncryptionDeclarationCreateRequestDataAttributes := _AppEncryptionDeclarationCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppEncryptionDeclarationCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = AppEncryptionDeclarationCreateRequestDataAttributes(varAppEncryptionDeclarationCreateRequestDataAttributes)

	return err
}

type NullableAppEncryptionDeclarationCreateRequestDataAttributes struct {
	value *AppEncryptionDeclarationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppEncryptionDeclarationCreateRequestDataAttributes) Get() *AppEncryptionDeclarationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataAttributes) Set(val *AppEncryptionDeclarationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationCreateRequestDataAttributes(val *AppEncryptionDeclarationCreateRequestDataAttributes) *NullableAppEncryptionDeclarationCreateRequestDataAttributes {
	return &NullableAppEncryptionDeclarationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


