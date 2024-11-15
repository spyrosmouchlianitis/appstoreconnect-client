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

// checks if the CiBuildRunRelationshipsWorkflowData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunRelationshipsWorkflowData{}

// CiBuildRunRelationshipsWorkflowData struct for CiBuildRunRelationshipsWorkflowData
type CiBuildRunRelationshipsWorkflowData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _CiBuildRunRelationshipsWorkflowData CiBuildRunRelationshipsWorkflowData

// NewCiBuildRunRelationshipsWorkflowData instantiates a new CiBuildRunRelationshipsWorkflowData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunRelationshipsWorkflowData(type_ string, id string) *CiBuildRunRelationshipsWorkflowData {
	this := CiBuildRunRelationshipsWorkflowData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiBuildRunRelationshipsWorkflowDataWithDefaults instantiates a new CiBuildRunRelationshipsWorkflowData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunRelationshipsWorkflowDataWithDefaults() *CiBuildRunRelationshipsWorkflowData {
	this := CiBuildRunRelationshipsWorkflowData{}
	return &this
}

// GetType returns the Type field value
func (o *CiBuildRunRelationshipsWorkflowData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationshipsWorkflowData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiBuildRunRelationshipsWorkflowData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiBuildRunRelationshipsWorkflowData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationshipsWorkflowData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiBuildRunRelationshipsWorkflowData) SetId(v string) {
	o.Id = v
}

func (o CiBuildRunRelationshipsWorkflowData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunRelationshipsWorkflowData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CiBuildRunRelationshipsWorkflowData) UnmarshalJSON(data []byte) (err error) {
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

	varCiBuildRunRelationshipsWorkflowData := _CiBuildRunRelationshipsWorkflowData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCiBuildRunRelationshipsWorkflowData)

	if err != nil {
		return err
	}

	*o = CiBuildRunRelationshipsWorkflowData(varCiBuildRunRelationshipsWorkflowData)

	return err
}

type NullableCiBuildRunRelationshipsWorkflowData struct {
	value *CiBuildRunRelationshipsWorkflowData
	isSet bool
}

func (v NullableCiBuildRunRelationshipsWorkflowData) Get() *CiBuildRunRelationshipsWorkflowData {
	return v.value
}

func (v *NullableCiBuildRunRelationshipsWorkflowData) Set(val *CiBuildRunRelationshipsWorkflowData) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunRelationshipsWorkflowData) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunRelationshipsWorkflowData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunRelationshipsWorkflowData(val *CiBuildRunRelationshipsWorkflowData) *NullableCiBuildRunRelationshipsWorkflowData {
	return &NullableCiBuildRunRelationshipsWorkflowData{value: val, isSet: true}
}

func (v NullableCiBuildRunRelationshipsWorkflowData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunRelationshipsWorkflowData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


