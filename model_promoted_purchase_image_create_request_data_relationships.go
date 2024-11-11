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

// checks if the PromotedPurchaseImageCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageCreateRequestDataRelationships{}

// PromotedPurchaseImageCreateRequestDataRelationships struct for PromotedPurchaseImageCreateRequestDataRelationships
type PromotedPurchaseImageCreateRequestDataRelationships struct {
	PromotedPurchase PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase `json:"promotedPurchase"`
}

type _PromotedPurchaseImageCreateRequestDataRelationships PromotedPurchaseImageCreateRequestDataRelationships

// NewPromotedPurchaseImageCreateRequestDataRelationships instantiates a new PromotedPurchaseImageCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageCreateRequestDataRelationships(promotedPurchase PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) *PromotedPurchaseImageCreateRequestDataRelationships {
	this := PromotedPurchaseImageCreateRequestDataRelationships{}
	this.PromotedPurchase = promotedPurchase
	return &this
}

// NewPromotedPurchaseImageCreateRequestDataRelationshipsWithDefaults instantiates a new PromotedPurchaseImageCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageCreateRequestDataRelationshipsWithDefaults() *PromotedPurchaseImageCreateRequestDataRelationships {
	this := PromotedPurchaseImageCreateRequestDataRelationships{}
	return &this
}

// GetPromotedPurchase returns the PromotedPurchase field value
func (o *PromotedPurchaseImageCreateRequestDataRelationships) GetPromotedPurchase() PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase {
	if o == nil {
		var ret PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase
		return ret
	}

	return o.PromotedPurchase
}

// GetPromotedPurchaseOk returns a tuple with the PromotedPurchase field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageCreateRequestDataRelationships) GetPromotedPurchaseOk() (*PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromotedPurchase, true
}

// SetPromotedPurchase sets field value
func (o *PromotedPurchaseImageCreateRequestDataRelationships) SetPromotedPurchase(v PromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase) {
	o.PromotedPurchase = v
}

func (o PromotedPurchaseImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promotedPurchase"] = o.PromotedPurchase
	return toSerialize, nil
}

func (o *PromotedPurchaseImageCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"promotedPurchase",
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

	varPromotedPurchaseImageCreateRequestDataRelationships := _PromotedPurchaseImageCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPromotedPurchaseImageCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = PromotedPurchaseImageCreateRequestDataRelationships(varPromotedPurchaseImageCreateRequestDataRelationships)

	return err
}

type NullablePromotedPurchaseImageCreateRequestDataRelationships struct {
	value *PromotedPurchaseImageCreateRequestDataRelationships
	isSet bool
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationships) Get() *PromotedPurchaseImageCreateRequestDataRelationships {
	return v.value
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationships) Set(val *PromotedPurchaseImageCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageCreateRequestDataRelationships(val *PromotedPurchaseImageCreateRequestDataRelationships) *NullablePromotedPurchaseImageCreateRequestDataRelationships {
	return &NullablePromotedPurchaseImageCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


