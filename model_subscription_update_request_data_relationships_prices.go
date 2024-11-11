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

// checks if the SubscriptionUpdateRequestDataRelationshipsPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUpdateRequestDataRelationshipsPrices{}

// SubscriptionUpdateRequestDataRelationshipsPrices struct for SubscriptionUpdateRequestDataRelationshipsPrices
type SubscriptionUpdateRequestDataRelationshipsPrices struct {
	Data []SubscriptionRelationshipsPricesDataInner `json:"data,omitempty"`
}

// NewSubscriptionUpdateRequestDataRelationshipsPrices instantiates a new SubscriptionUpdateRequestDataRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUpdateRequestDataRelationshipsPrices() *SubscriptionUpdateRequestDataRelationshipsPrices {
	this := SubscriptionUpdateRequestDataRelationshipsPrices{}
	return &this
}

// NewSubscriptionUpdateRequestDataRelationshipsPricesWithDefaults instantiates a new SubscriptionUpdateRequestDataRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUpdateRequestDataRelationshipsPricesWithDefaults() *SubscriptionUpdateRequestDataRelationshipsPrices {
	this := SubscriptionUpdateRequestDataRelationshipsPrices{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequestDataRelationshipsPrices) GetData() []SubscriptionRelationshipsPricesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SubscriptionRelationshipsPricesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequestDataRelationshipsPrices) GetDataOk() ([]SubscriptionRelationshipsPricesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequestDataRelationshipsPrices) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubscriptionRelationshipsPricesDataInner and assigns it to the Data field.
func (o *SubscriptionUpdateRequestDataRelationshipsPrices) SetData(v []SubscriptionRelationshipsPricesDataInner) {
	o.Data = v
}

func (o SubscriptionUpdateRequestDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUpdateRequestDataRelationshipsPrices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionUpdateRequestDataRelationshipsPrices struct {
	value *SubscriptionUpdateRequestDataRelationshipsPrices
	isSet bool
}

func (v NullableSubscriptionUpdateRequestDataRelationshipsPrices) Get() *SubscriptionUpdateRequestDataRelationshipsPrices {
	return v.value
}

func (v *NullableSubscriptionUpdateRequestDataRelationshipsPrices) Set(val *SubscriptionUpdateRequestDataRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUpdateRequestDataRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUpdateRequestDataRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUpdateRequestDataRelationshipsPrices(val *SubscriptionUpdateRequestDataRelationshipsPrices) *NullableSubscriptionUpdateRequestDataRelationshipsPrices {
	return &NullableSubscriptionUpdateRequestDataRelationshipsPrices{value: val, isSet: true}
}

func (v NullableSubscriptionUpdateRequestDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUpdateRequestDataRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


