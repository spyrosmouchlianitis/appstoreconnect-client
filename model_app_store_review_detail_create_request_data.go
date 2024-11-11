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

// checks if the AppStoreReviewDetailCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailCreateRequestData{}

// AppStoreReviewDetailCreateRequestData struct for AppStoreReviewDetailCreateRequestData
type AppStoreReviewDetailCreateRequestData struct {
	Type string `json:"type"`
	Attributes *AppStoreReviewDetailAttributes `json:"attributes,omitempty"`
	Relationships AlternativeDistributionPackageCreateRequestDataRelationships `json:"relationships"`
}

type _AppStoreReviewDetailCreateRequestData AppStoreReviewDetailCreateRequestData

// NewAppStoreReviewDetailCreateRequestData instantiates a new AppStoreReviewDetailCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailCreateRequestData(type_ string, relationships AlternativeDistributionPackageCreateRequestDataRelationships) *AppStoreReviewDetailCreateRequestData {
	this := AppStoreReviewDetailCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppStoreReviewDetailCreateRequestDataWithDefaults instantiates a new AppStoreReviewDetailCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailCreateRequestDataWithDefaults() *AppStoreReviewDetailCreateRequestData {
	this := AppStoreReviewDetailCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewDetailCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewDetailCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreReviewDetailCreateRequestData) GetAttributes() AppStoreReviewDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreReviewDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequestData) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreReviewDetailCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreReviewDetailAttributes and assigns it to the Attributes field.
func (o *AppStoreReviewDetailCreateRequestData) SetAttributes(v AppStoreReviewDetailAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreReviewDetailCreateRequestData) GetRelationships() AlternativeDistributionPackageCreateRequestDataRelationships {
	if o == nil {
		var ret AlternativeDistributionPackageCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequestData) GetRelationshipsOk() (*AlternativeDistributionPackageCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreReviewDetailCreateRequestData) SetRelationships(v AlternativeDistributionPackageCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreReviewDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *AppStoreReviewDetailCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varAppStoreReviewDetailCreateRequestData := _AppStoreReviewDetailCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppStoreReviewDetailCreateRequestData)

	if err != nil {
		return err
	}

	*o = AppStoreReviewDetailCreateRequestData(varAppStoreReviewDetailCreateRequestData)

	return err
}

type NullableAppStoreReviewDetailCreateRequestData struct {
	value *AppStoreReviewDetailCreateRequestData
	isSet bool
}

func (v NullableAppStoreReviewDetailCreateRequestData) Get() *AppStoreReviewDetailCreateRequestData {
	return v.value
}

func (v *NullableAppStoreReviewDetailCreateRequestData) Set(val *AppStoreReviewDetailCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailCreateRequestData(val *AppStoreReviewDetailCreateRequestData) *NullableAppStoreReviewDetailCreateRequestData {
	return &NullableAppStoreReviewDetailCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


