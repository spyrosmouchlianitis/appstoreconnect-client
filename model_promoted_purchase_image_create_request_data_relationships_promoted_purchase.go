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

// checks if the PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase{}

// PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase struct for PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase
type PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase struct {
	Data AppRelationshipsPromotedPurchasesDataInner `json:"data"`
}

type _PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase

// NewPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase instantiates a new PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase(data AppRelationshipsPromotedPurchasesDataInner) *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase {
	this := PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase{}
	this.Data = data
	return &this
}

// NewPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchaseWithDefaults instantiates a new PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchaseWithDefaults() *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase {
	this := PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase{}
	return &this
}

// GetData returns the Data field value
func (o *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) GetData() AppRelationshipsPromotedPurchasesDataInner {
	if o == nil {
		var ret AppRelationshipsPromotedPurchasesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) GetDataOk() (*AppRelationshipsPromotedPurchasesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) SetData(v AppRelationshipsPromotedPurchasesDataInner) {
	o.Data = v
}

func (o PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase := _PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase)

	if err != nil {
		return err
	}

	*o = PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase(varPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase)

	return err
}

type NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase struct {
	value *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase
	isSet bool
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) Get() *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase {
	return v.value
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) Set(val *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase(val *PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) *NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase {
	return &NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

