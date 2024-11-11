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

// checks if the SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships{}

// SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships struct for SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships
type SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships struct {
	SandboxTesters SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters `json:"sandboxTesters"`
}

type _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships(sandboxTesters SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships{}
	this.SandboxTesters = sandboxTesters
	return &this
}

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsWithDefaults instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsWithDefaults() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships{}
	return &this
}

// GetSandboxTesters returns the SandboxTesters field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) GetSandboxTesters() SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters {
	if o == nil {
		var ret SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters
		return ret
	}

	return o.SandboxTesters
}

// GetSandboxTestersOk returns a tuple with the SandboxTesters field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) GetSandboxTestersOk() (*SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SandboxTesters, true
}

// SetSandboxTesters sets field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) SetSandboxTesters(v SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) {
	o.SandboxTesters = v
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sandboxTesters"] = o.SandboxTesters
	return toSerialize, nil
}

func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sandboxTesters",
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

	varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships := _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships(varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships)

	return err
}

type NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships struct {
	value *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships
	isSet bool
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) Get() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships {
	return v.value
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) Set(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships {
	return &NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


