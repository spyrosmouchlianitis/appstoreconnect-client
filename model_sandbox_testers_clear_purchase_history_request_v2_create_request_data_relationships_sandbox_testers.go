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

// checks if the SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters{}

// SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters struct for SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters
type SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters struct {
	Data []SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner `json:"data"`
}

type _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters(data []SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner) *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters{}
	this.Data = data
	return &this
}

// NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersWithDefaults instantiates a new SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersWithDefaults() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters {
	this := SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters{}
	return &this
}

// GetData returns the Data field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) GetData() []SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner {
	if o == nil {
		var ret []SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) GetDataOk() ([]SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) SetData(v []SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTestersDataInner) {
	o.Data = v
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) UnmarshalJSON(data []byte) (err error) {
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

	varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters := _SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters)

	if err != nil {
		return err
	}

	*o = SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters(varSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters)

	return err
}

type NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters struct {
	value *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters
	isSet bool
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) Get() *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters {
	return v.value
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) Set(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters(val *SandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters {
	return &NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters{value: val, isSet: true}
}

func (v NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTestersClearPurchaseHistoryRequestV2CreateRequestDataRelationshipsSandboxTesters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


