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

// checks if the AppCategoryRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCategoryRelationships{}

// AppCategoryRelationships struct for AppCategoryRelationships
type AppCategoryRelationships struct {
	Subcategories *AppCategoryRelationshipsSubcategories `json:"subcategories,omitempty"`
	Parent *AppCategoryRelationshipsParent `json:"parent,omitempty"`
}

// NewAppCategoryRelationships instantiates a new AppCategoryRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryRelationships() *AppCategoryRelationships {
	this := AppCategoryRelationships{}
	return &this
}

// NewAppCategoryRelationshipsWithDefaults instantiates a new AppCategoryRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryRelationshipsWithDefaults() *AppCategoryRelationships {
	this := AppCategoryRelationships{}
	return &this
}

// GetSubcategories returns the Subcategories field value if set, zero value otherwise.
func (o *AppCategoryRelationships) GetSubcategories() AppCategoryRelationshipsSubcategories {
	if o == nil || IsNil(o.Subcategories) {
		var ret AppCategoryRelationshipsSubcategories
		return ret
	}
	return *o.Subcategories
}

// GetSubcategoriesOk returns a tuple with the Subcategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationships) GetSubcategoriesOk() (*AppCategoryRelationshipsSubcategories, bool) {
	if o == nil || IsNil(o.Subcategories) {
		return nil, false
	}
	return o.Subcategories, true
}

// HasSubcategories returns a boolean if a field has been set.
func (o *AppCategoryRelationships) HasSubcategories() bool {
	if o != nil && !IsNil(o.Subcategories) {
		return true
	}

	return false
}

// SetSubcategories gets a reference to the given AppCategoryRelationshipsSubcategories and assigns it to the Subcategories field.
func (o *AppCategoryRelationships) SetSubcategories(v AppCategoryRelationshipsSubcategories) {
	o.Subcategories = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *AppCategoryRelationships) GetParent() AppCategoryRelationshipsParent {
	if o == nil || IsNil(o.Parent) {
		var ret AppCategoryRelationshipsParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationships) GetParentOk() (*AppCategoryRelationshipsParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *AppCategoryRelationships) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given AppCategoryRelationshipsParent and assigns it to the Parent field.
func (o *AppCategoryRelationships) SetParent(v AppCategoryRelationshipsParent) {
	o.Parent = &v
}

func (o AppCategoryRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCategoryRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subcategories) {
		toSerialize["subcategories"] = o.Subcategories
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

type NullableAppCategoryRelationships struct {
	value *AppCategoryRelationships
	isSet bool
}

func (v NullableAppCategoryRelationships) Get() *AppCategoryRelationships {
	return v.value
}

func (v *NullableAppCategoryRelationships) Set(val *AppCategoryRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryRelationships(val *AppCategoryRelationships) *NullableAppCategoryRelationships {
	return &NullableAppCategoryRelationships{value: val, isSet: true}
}

func (v NullableAppCategoryRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


