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

// checks if the AppStoreVersionReleaseRequestCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionReleaseRequestCreateRequestData{}

// AppStoreVersionReleaseRequestCreateRequestData struct for AppStoreVersionReleaseRequestCreateRequestData
type AppStoreVersionReleaseRequestCreateRequestData struct {
	Type string `json:"type"`
	Relationships AlternativeDistributionPackageCreateRequestDataRelationships `json:"relationships"`
}

type _AppStoreVersionReleaseRequestCreateRequestData AppStoreVersionReleaseRequestCreateRequestData

// NewAppStoreVersionReleaseRequestCreateRequestData instantiates a new AppStoreVersionReleaseRequestCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionReleaseRequestCreateRequestData(type_ string, relationships AlternativeDistributionPackageCreateRequestDataRelationships) *AppStoreVersionReleaseRequestCreateRequestData {
	this := AppStoreVersionReleaseRequestCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionReleaseRequestCreateRequestDataWithDefaults instantiates a new AppStoreVersionReleaseRequestCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionReleaseRequestCreateRequestDataWithDefaults() *AppStoreVersionReleaseRequestCreateRequestData {
	this := AppStoreVersionReleaseRequestCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionReleaseRequestCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionReleaseRequestCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionReleaseRequestCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionReleaseRequestCreateRequestData) GetRelationships() AlternativeDistributionPackageCreateRequestDataRelationships {
	if o == nil {
		var ret AlternativeDistributionPackageCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionReleaseRequestCreateRequestData) GetRelationshipsOk() (*AlternativeDistributionPackageCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionReleaseRequestCreateRequestData) SetRelationships(v AlternativeDistributionPackageCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionReleaseRequestCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionReleaseRequestCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AppStoreVersionReleaseRequestCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreVersionReleaseRequestCreateRequestData := _AppStoreVersionReleaseRequestCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreVersionReleaseRequestCreateRequestData)

	if err != nil {
		return err
	}

	*o = AppStoreVersionReleaseRequestCreateRequestData(varAppStoreVersionReleaseRequestCreateRequestData)

	return err
}

type NullableAppStoreVersionReleaseRequestCreateRequestData struct {
	value *AppStoreVersionReleaseRequestCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionReleaseRequestCreateRequestData) Get() *AppStoreVersionReleaseRequestCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionReleaseRequestCreateRequestData) Set(val *AppStoreVersionReleaseRequestCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionReleaseRequestCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionReleaseRequestCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionReleaseRequestCreateRequestData(val *AppStoreVersionReleaseRequestCreateRequestData) *NullableAppStoreVersionReleaseRequestCreateRequestData {
	return &NullableAppStoreVersionReleaseRequestCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionReleaseRequestCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionReleaseRequestCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


