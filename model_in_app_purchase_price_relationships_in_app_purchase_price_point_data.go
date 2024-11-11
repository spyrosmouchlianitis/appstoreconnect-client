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

// checks if the InAppPurchasePriceRelationshipsInAppPurchasePricePointData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceRelationshipsInAppPurchasePricePointData{}

// InAppPurchasePriceRelationshipsInAppPurchasePricePointData struct for InAppPurchasePriceRelationshipsInAppPurchasePricePointData
type InAppPurchasePriceRelationshipsInAppPurchasePricePointData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _InAppPurchasePriceRelationshipsInAppPurchasePricePointData InAppPurchasePriceRelationshipsInAppPurchasePricePointData

// NewInAppPurchasePriceRelationshipsInAppPurchasePricePointData instantiates a new InAppPurchasePriceRelationshipsInAppPurchasePricePointData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceRelationshipsInAppPurchasePricePointData(type_ string, id string) *InAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	this := InAppPurchasePriceRelationshipsInAppPurchasePricePointData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchasePriceRelationshipsInAppPurchasePricePointDataWithDefaults instantiates a new InAppPurchasePriceRelationshipsInAppPurchasePricePointData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceRelationshipsInAppPurchasePricePointDataWithDefaults() *InAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	this := InAppPurchasePriceRelationshipsInAppPurchasePricePointData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) SetId(v string) {
	o.Id = v
}

func (o InAppPurchasePriceRelationshipsInAppPurchasePricePointData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceRelationshipsInAppPurchasePricePointData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varInAppPurchasePriceRelationshipsInAppPurchasePricePointData := _InAppPurchasePriceRelationshipsInAppPurchasePricePointData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInAppPurchasePriceRelationshipsInAppPurchasePricePointData)

	if err != nil {
		return err
	}

	*o = InAppPurchasePriceRelationshipsInAppPurchasePricePointData(varInAppPurchasePriceRelationshipsInAppPurchasePricePointData)

	return err
}

type NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData struct {
	value *InAppPurchasePriceRelationshipsInAppPurchasePricePointData
	isSet bool
}

func (v NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) Get() *InAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	return v.value
}

func (v *NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) Set(val *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData(val *InAppPurchasePriceRelationshipsInAppPurchasePricePointData) *NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	return &NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceRelationshipsInAppPurchasePricePointData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


