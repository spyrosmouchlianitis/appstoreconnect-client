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

// checks if the AppRelationshipsInAppPurchasesV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsInAppPurchasesV2{}

// AppRelationshipsInAppPurchasesV2 struct for AppRelationshipsInAppPurchasesV2
type AppRelationshipsInAppPurchasesV2 struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppRelationshipsInAppPurchasesDataInner `json:"data,omitempty"`
}

// NewAppRelationshipsInAppPurchasesV2 instantiates a new AppRelationshipsInAppPurchasesV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsInAppPurchasesV2() *AppRelationshipsInAppPurchasesV2 {
	this := AppRelationshipsInAppPurchasesV2{}
	return &this
}

// NewAppRelationshipsInAppPurchasesV2WithDefaults instantiates a new AppRelationshipsInAppPurchasesV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsInAppPurchasesV2WithDefaults() *AppRelationshipsInAppPurchasesV2 {
	this := AppRelationshipsInAppPurchasesV2{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsInAppPurchasesV2) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsInAppPurchasesV2) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsInAppPurchasesV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *AppRelationshipsInAppPurchasesV2) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppRelationshipsInAppPurchasesV2) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsInAppPurchasesV2) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppRelationshipsInAppPurchasesV2) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppRelationshipsInAppPurchasesV2) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsInAppPurchasesV2) GetData() []AppRelationshipsInAppPurchasesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppRelationshipsInAppPurchasesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsInAppPurchasesV2) GetDataOk() ([]AppRelationshipsInAppPurchasesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsInAppPurchasesV2) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppRelationshipsInAppPurchasesDataInner and assigns it to the Data field.
func (o *AppRelationshipsInAppPurchasesV2) SetData(v []AppRelationshipsInAppPurchasesDataInner) {
	o.Data = v
}

func (o AppRelationshipsInAppPurchasesV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsInAppPurchasesV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsInAppPurchasesV2 struct {
	value *AppRelationshipsInAppPurchasesV2
	isSet bool
}

func (v NullableAppRelationshipsInAppPurchasesV2) Get() *AppRelationshipsInAppPurchasesV2 {
	return v.value
}

func (v *NullableAppRelationshipsInAppPurchasesV2) Set(val *AppRelationshipsInAppPurchasesV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsInAppPurchasesV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsInAppPurchasesV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsInAppPurchasesV2(val *AppRelationshipsInAppPurchasesV2) *NullableAppRelationshipsInAppPurchasesV2 {
	return &NullableAppRelationshipsInAppPurchasesV2{value: val, isSet: true}
}

func (v NullableAppRelationshipsInAppPurchasesV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsInAppPurchasesV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


