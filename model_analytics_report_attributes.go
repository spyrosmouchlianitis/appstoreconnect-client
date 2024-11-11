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

// checks if the AnalyticsReportAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsReportAttributes{}

// AnalyticsReportAttributes struct for AnalyticsReportAttributes
type AnalyticsReportAttributes struct {
	Name *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
}

// NewAnalyticsReportAttributes instantiates a new AnalyticsReportAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsReportAttributes() *AnalyticsReportAttributes {
	this := AnalyticsReportAttributes{}
	return &this
}

// NewAnalyticsReportAttributesWithDefaults instantiates a new AnalyticsReportAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsReportAttributesWithDefaults() *AnalyticsReportAttributes {
	this := AnalyticsReportAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnalyticsReportAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnalyticsReportAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnalyticsReportAttributes) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AnalyticsReportAttributes) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsReportAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AnalyticsReportAttributes) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AnalyticsReportAttributes) SetCategory(v string) {
	o.Category = &v
}

func (o AnalyticsReportAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsReportAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableAnalyticsReportAttributes struct {
	value *AnalyticsReportAttributes
	isSet bool
}

func (v NullableAnalyticsReportAttributes) Get() *AnalyticsReportAttributes {
	return v.value
}

func (v *NullableAnalyticsReportAttributes) Set(val *AnalyticsReportAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsReportAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsReportAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsReportAttributes(val *AnalyticsReportAttributes) *NullableAnalyticsReportAttributes {
	return &NullableAnalyticsReportAttributes{value: val, isSet: true}
}

func (v NullableAnalyticsReportAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsReportAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


