/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the CiWorkflowRelationshipsXcodeVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowRelationshipsXcodeVersion{}

// CiWorkflowRelationshipsXcodeVersion struct for CiWorkflowRelationshipsXcodeVersion
type CiWorkflowRelationshipsXcodeVersion struct {
	Data *CiMacOsVersionRelationshipsXcodeVersionsDataInner `json:"data,omitempty"`
}

// NewCiWorkflowRelationshipsXcodeVersion instantiates a new CiWorkflowRelationshipsXcodeVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowRelationshipsXcodeVersion() *CiWorkflowRelationshipsXcodeVersion {
	this := CiWorkflowRelationshipsXcodeVersion{}
	return &this
}

// NewCiWorkflowRelationshipsXcodeVersionWithDefaults instantiates a new CiWorkflowRelationshipsXcodeVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowRelationshipsXcodeVersionWithDefaults() *CiWorkflowRelationshipsXcodeVersion {
	this := CiWorkflowRelationshipsXcodeVersion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiWorkflowRelationshipsXcodeVersion) GetData() CiMacOsVersionRelationshipsXcodeVersionsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret CiMacOsVersionRelationshipsXcodeVersionsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationshipsXcodeVersion) GetDataOk() (*CiMacOsVersionRelationshipsXcodeVersionsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiWorkflowRelationshipsXcodeVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiMacOsVersionRelationshipsXcodeVersionsDataInner and assigns it to the Data field.
func (o *CiWorkflowRelationshipsXcodeVersion) SetData(v CiMacOsVersionRelationshipsXcodeVersionsDataInner) {
	o.Data = &v
}

func (o CiWorkflowRelationshipsXcodeVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowRelationshipsXcodeVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiWorkflowRelationshipsXcodeVersion struct {
	value *CiWorkflowRelationshipsXcodeVersion
	isSet bool
}

func (v NullableCiWorkflowRelationshipsXcodeVersion) Get() *CiWorkflowRelationshipsXcodeVersion {
	return v.value
}

func (v *NullableCiWorkflowRelationshipsXcodeVersion) Set(val *CiWorkflowRelationshipsXcodeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowRelationshipsXcodeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowRelationshipsXcodeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowRelationshipsXcodeVersion(val *CiWorkflowRelationshipsXcodeVersion) *NullableCiWorkflowRelationshipsXcodeVersion {
	return &NullableCiWorkflowRelationshipsXcodeVersion{value: val, isSet: true}
}

func (v NullableCiWorkflowRelationshipsXcodeVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowRelationshipsXcodeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


