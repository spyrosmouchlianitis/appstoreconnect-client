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

// checks if the PromotedPurchaseImageRelationshipsPromotedPurchase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageRelationshipsPromotedPurchase{}

// PromotedPurchaseImageRelationshipsPromotedPurchase struct for PromotedPurchaseImageRelationshipsPromotedPurchase
type PromotedPurchaseImageRelationshipsPromotedPurchase struct {
	Data *AppRelationshipsPromotedPurchasesDataInner `json:"data,omitempty"`
}

// NewPromotedPurchaseImageRelationshipsPromotedPurchase instantiates a new PromotedPurchaseImageRelationshipsPromotedPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageRelationshipsPromotedPurchase() *PromotedPurchaseImageRelationshipsPromotedPurchase {
	this := PromotedPurchaseImageRelationshipsPromotedPurchase{}
	return &this
}

// NewPromotedPurchaseImageRelationshipsPromotedPurchaseWithDefaults instantiates a new PromotedPurchaseImageRelationshipsPromotedPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageRelationshipsPromotedPurchaseWithDefaults() *PromotedPurchaseImageRelationshipsPromotedPurchase {
	this := PromotedPurchaseImageRelationshipsPromotedPurchase{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PromotedPurchaseImageRelationshipsPromotedPurchase) GetData() AppRelationshipsPromotedPurchasesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsPromotedPurchasesDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageRelationshipsPromotedPurchase) GetDataOk() (*AppRelationshipsPromotedPurchasesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PromotedPurchaseImageRelationshipsPromotedPurchase) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsPromotedPurchasesDataInner and assigns it to the Data field.
func (o *PromotedPurchaseImageRelationshipsPromotedPurchase) SetData(v AppRelationshipsPromotedPurchasesDataInner) {
	o.Data = &v
}

func (o PromotedPurchaseImageRelationshipsPromotedPurchase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageRelationshipsPromotedPurchase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseImageRelationshipsPromotedPurchase struct {
	value *PromotedPurchaseImageRelationshipsPromotedPurchase
	isSet bool
}

func (v NullablePromotedPurchaseImageRelationshipsPromotedPurchase) Get() *PromotedPurchaseImageRelationshipsPromotedPurchase {
	return v.value
}

func (v *NullablePromotedPurchaseImageRelationshipsPromotedPurchase) Set(val *PromotedPurchaseImageRelationshipsPromotedPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageRelationshipsPromotedPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageRelationshipsPromotedPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageRelationshipsPromotedPurchase(val *PromotedPurchaseImageRelationshipsPromotedPurchase) *NullablePromotedPurchaseImageRelationshipsPromotedPurchase {
	return &NullablePromotedPurchaseImageRelationshipsPromotedPurchase{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageRelationshipsPromotedPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageRelationshipsPromotedPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


