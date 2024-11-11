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

// checks if the PromotedPurchaseRelationshipsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseRelationshipsSubscription{}

// PromotedPurchaseRelationshipsSubscription struct for PromotedPurchaseRelationshipsSubscription
type PromotedPurchaseRelationshipsSubscription struct {
	Data *PromotedPurchaseRelationshipsSubscriptionData `json:"data,omitempty"`
}

// NewPromotedPurchaseRelationshipsSubscription instantiates a new PromotedPurchaseRelationshipsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseRelationshipsSubscription() *PromotedPurchaseRelationshipsSubscription {
	this := PromotedPurchaseRelationshipsSubscription{}
	return &this
}

// NewPromotedPurchaseRelationshipsSubscriptionWithDefaults instantiates a new PromotedPurchaseRelationshipsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseRelationshipsSubscriptionWithDefaults() *PromotedPurchaseRelationshipsSubscription {
	this := PromotedPurchaseRelationshipsSubscription{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PromotedPurchaseRelationshipsSubscription) GetData() PromotedPurchaseRelationshipsSubscriptionData {
	if o == nil || IsNil(o.Data) {
		var ret PromotedPurchaseRelationshipsSubscriptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsSubscription) GetDataOk() (*PromotedPurchaseRelationshipsSubscriptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PromotedPurchaseRelationshipsSubscription) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PromotedPurchaseRelationshipsSubscriptionData and assigns it to the Data field.
func (o *PromotedPurchaseRelationshipsSubscription) SetData(v PromotedPurchaseRelationshipsSubscriptionData) {
	o.Data = &v
}

func (o PromotedPurchaseRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseRelationshipsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseRelationshipsSubscription struct {
	value *PromotedPurchaseRelationshipsSubscription
	isSet bool
}

func (v NullablePromotedPurchaseRelationshipsSubscription) Get() *PromotedPurchaseRelationshipsSubscription {
	return v.value
}

func (v *NullablePromotedPurchaseRelationshipsSubscription) Set(val *PromotedPurchaseRelationshipsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseRelationshipsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseRelationshipsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseRelationshipsSubscription(val *PromotedPurchaseRelationshipsSubscription) *NullablePromotedPurchaseRelationshipsSubscription {
	return &NullablePromotedPurchaseRelationshipsSubscription{value: val, isSet: true}
}

func (v NullablePromotedPurchaseRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseRelationshipsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

