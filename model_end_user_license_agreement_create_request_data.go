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

// checks if the EndUserLicenseAgreementCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementCreateRequestData{}

// EndUserLicenseAgreementCreateRequestData struct for EndUserLicenseAgreementCreateRequestData
type EndUserLicenseAgreementCreateRequestData struct {
	Type string `json:"type"`
	Attributes EndUserLicenseAgreementCreateRequestDataAttributes `json:"attributes"`
	Relationships EndUserLicenseAgreementCreateRequestDataRelationships `json:"relationships"`
}

type _EndUserLicenseAgreementCreateRequestData EndUserLicenseAgreementCreateRequestData

// NewEndUserLicenseAgreementCreateRequestData instantiates a new EndUserLicenseAgreementCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementCreateRequestData(type_ string, attributes EndUserLicenseAgreementCreateRequestDataAttributes, relationships EndUserLicenseAgreementCreateRequestDataRelationships) *EndUserLicenseAgreementCreateRequestData {
	this := EndUserLicenseAgreementCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewEndUserLicenseAgreementCreateRequestDataWithDefaults instantiates a new EndUserLicenseAgreementCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementCreateRequestDataWithDefaults() *EndUserLicenseAgreementCreateRequestData {
	this := EndUserLicenseAgreementCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *EndUserLicenseAgreementCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndUserLicenseAgreementCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EndUserLicenseAgreementCreateRequestData) GetAttributes() EndUserLicenseAgreementCreateRequestDataAttributes {
	if o == nil {
		var ret EndUserLicenseAgreementCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestData) GetAttributesOk() (*EndUserLicenseAgreementCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EndUserLicenseAgreementCreateRequestData) SetAttributes(v EndUserLicenseAgreementCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *EndUserLicenseAgreementCreateRequestData) GetRelationships() EndUserLicenseAgreementCreateRequestDataRelationships {
	if o == nil {
		var ret EndUserLicenseAgreementCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequestData) GetRelationshipsOk() (*EndUserLicenseAgreementCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *EndUserLicenseAgreementCreateRequestData) SetRelationships(v EndUserLicenseAgreementCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o EndUserLicenseAgreementCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *EndUserLicenseAgreementCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varEndUserLicenseAgreementCreateRequestData := _EndUserLicenseAgreementCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndUserLicenseAgreementCreateRequestData)

	if err != nil {
		return err
	}

	*o = EndUserLicenseAgreementCreateRequestData(varEndUserLicenseAgreementCreateRequestData)

	return err
}

type NullableEndUserLicenseAgreementCreateRequestData struct {
	value *EndUserLicenseAgreementCreateRequestData
	isSet bool
}

func (v NullableEndUserLicenseAgreementCreateRequestData) Get() *EndUserLicenseAgreementCreateRequestData {
	return v.value
}

func (v *NullableEndUserLicenseAgreementCreateRequestData) Set(val *EndUserLicenseAgreementCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementCreateRequestData(val *EndUserLicenseAgreementCreateRequestData) *NullableEndUserLicenseAgreementCreateRequestData {
	return &NullableEndUserLicenseAgreementCreateRequestData{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


