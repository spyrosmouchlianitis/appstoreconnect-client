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

// checks if the SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData{}

// SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData struct for SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData
type SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData struct {
	Type string `json:"type"`
	Relationships SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships `json:"relationships"`
}

type _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData(type_ string, relationships SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataWithDefaults instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataWithDefaults() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) GetRelationships() SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships {
	if o == nil {
		var ret SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) GetRelationshipsOk() (*SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) SetRelationships(v SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData := _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData)

	if err != nil {
		return err
	}

	*o = SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData(varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData)

	return err
}

type NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData struct {
	value *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData
	isSet bool
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) Get() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData {
	return v.value
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) Set(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData {
	return &NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData{value: val, isSet: true}
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


