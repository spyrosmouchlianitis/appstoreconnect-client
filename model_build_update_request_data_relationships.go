/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the BuildUpdateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildUpdateRequestDataRelationships{}

// BuildUpdateRequestDataRelationships struct for BuildUpdateRequestDataRelationships
type BuildUpdateRequestDataRelationships struct {
	AppEncryptionDeclaration *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration `json:"appEncryptionDeclaration,omitempty"`
}

// NewBuildUpdateRequestDataRelationships instantiates a new BuildUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildUpdateRequestDataRelationships() *BuildUpdateRequestDataRelationships {
	this := BuildUpdateRequestDataRelationships{}
	return &this
}

// NewBuildUpdateRequestDataRelationshipsWithDefaults instantiates a new BuildUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildUpdateRequestDataRelationshipsWithDefaults() *BuildUpdateRequestDataRelationships {
	this := BuildUpdateRequestDataRelationships{}
	return &this
}

// GetAppEncryptionDeclaration returns the AppEncryptionDeclaration field value if set, zero value otherwise.
func (o *BuildUpdateRequestDataRelationships) GetAppEncryptionDeclaration() BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration {
	if o == nil || IsNil(o.AppEncryptionDeclaration) {
		var ret BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration
		return ret
	}
	return *o.AppEncryptionDeclaration
}

// GetAppEncryptionDeclarationOk returns a tuple with the AppEncryptionDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestDataRelationships) GetAppEncryptionDeclarationOk() (*BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration, bool) {
	if o == nil || IsNil(o.AppEncryptionDeclaration) {
		return nil, false
	}
	return o.AppEncryptionDeclaration, true
}

// HasAppEncryptionDeclaration returns a boolean if a field has been set.
func (o *BuildUpdateRequestDataRelationships) HasAppEncryptionDeclaration() bool {
	if o != nil && !IsNil(o.AppEncryptionDeclaration) {
		return true
	}

	return false
}

// SetAppEncryptionDeclaration gets a reference to the given BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration and assigns it to the AppEncryptionDeclaration field.
func (o *BuildUpdateRequestDataRelationships) SetAppEncryptionDeclaration(v BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) {
	o.AppEncryptionDeclaration = &v
}

func (o BuildUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildUpdateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppEncryptionDeclaration) {
		toSerialize["appEncryptionDeclaration"] = o.AppEncryptionDeclaration
	}
	return toSerialize, nil
}

type NullableBuildUpdateRequestDataRelationships struct {
	value *BuildUpdateRequestDataRelationships
	isSet bool
}

func (v NullableBuildUpdateRequestDataRelationships) Get() *BuildUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableBuildUpdateRequestDataRelationships) Set(val *BuildUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildUpdateRequestDataRelationships(val *BuildUpdateRequestDataRelationships) *NullableBuildUpdateRequestDataRelationships {
	return &NullableBuildUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBuildUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


