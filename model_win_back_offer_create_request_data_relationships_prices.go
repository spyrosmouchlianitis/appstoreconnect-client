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

// checks if the WinBackOfferCreateRequestDataRelationshipsPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WinBackOfferCreateRequestDataRelationshipsPrices{}

// WinBackOfferCreateRequestDataRelationshipsPrices struct for WinBackOfferCreateRequestDataRelationshipsPrices
type WinBackOfferCreateRequestDataRelationshipsPrices struct {
	Data []WinBackOfferRelationshipsPricesDataInner `json:"data"`
}

type _WinBackOfferCreateRequestDataRelationshipsPrices WinBackOfferCreateRequestDataRelationshipsPrices

// NewWinBackOfferCreateRequestDataRelationshipsPrices instantiates a new WinBackOfferCreateRequestDataRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWinBackOfferCreateRequestDataRelationshipsPrices(data []WinBackOfferRelationshipsPricesDataInner) *WinBackOfferCreateRequestDataRelationshipsPrices {
	this := WinBackOfferCreateRequestDataRelationshipsPrices{}
	this.Data = data
	return &this
}

// NewWinBackOfferCreateRequestDataRelationshipsPricesWithDefaults instantiates a new WinBackOfferCreateRequestDataRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWinBackOfferCreateRequestDataRelationshipsPricesWithDefaults() *WinBackOfferCreateRequestDataRelationshipsPrices {
	this := WinBackOfferCreateRequestDataRelationshipsPrices{}
	return &this
}

// GetData returns the Data field value
func (o *WinBackOfferCreateRequestDataRelationshipsPrices) GetData() []WinBackOfferRelationshipsPricesDataInner {
	if o == nil {
		var ret []WinBackOfferRelationshipsPricesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataRelationshipsPrices) GetDataOk() ([]WinBackOfferRelationshipsPricesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WinBackOfferCreateRequestDataRelationshipsPrices) SetData(v []WinBackOfferRelationshipsPricesDataInner) {
	o.Data = v
}

func (o WinBackOfferCreateRequestDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WinBackOfferCreateRequestDataRelationshipsPrices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *WinBackOfferCreateRequestDataRelationshipsPrices) UnmarshalJSON(data []byte) (err error) {
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

	varWinBackOfferCreateRequestDataRelationshipsPrices := _WinBackOfferCreateRequestDataRelationshipsPrices{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWinBackOfferCreateRequestDataRelationshipsPrices)

	if err != nil {
		return err
	}

	*o = WinBackOfferCreateRequestDataRelationshipsPrices(varWinBackOfferCreateRequestDataRelationshipsPrices)

	return err
}

type NullableWinBackOfferCreateRequestDataRelationshipsPrices struct {
	value *WinBackOfferCreateRequestDataRelationshipsPrices
	isSet bool
}

func (v NullableWinBackOfferCreateRequestDataRelationshipsPrices) Get() *WinBackOfferCreateRequestDataRelationshipsPrices {
	return v.value
}

func (v *NullableWinBackOfferCreateRequestDataRelationshipsPrices) Set(val *WinBackOfferCreateRequestDataRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableWinBackOfferCreateRequestDataRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableWinBackOfferCreateRequestDataRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWinBackOfferCreateRequestDataRelationshipsPrices(val *WinBackOfferCreateRequestDataRelationshipsPrices) *NullableWinBackOfferCreateRequestDataRelationshipsPrices {
	return &NullableWinBackOfferCreateRequestDataRelationshipsPrices{value: val, isSet: true}
}

func (v NullableWinBackOfferCreateRequestDataRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWinBackOfferCreateRequestDataRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


