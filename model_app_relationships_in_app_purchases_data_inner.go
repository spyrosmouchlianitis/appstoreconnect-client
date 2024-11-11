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

// checks if the AppRelationshipsInAppPurchasesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsInAppPurchasesDataInner{}

// AppRelationshipsInAppPurchasesDataInner struct for AppRelationshipsInAppPurchasesDataInner
type AppRelationshipsInAppPurchasesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

type _AppRelationshipsInAppPurchasesDataInner AppRelationshipsInAppPurchasesDataInner

// NewAppRelationshipsInAppPurchasesDataInner instantiates a new AppRelationshipsInAppPurchasesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsInAppPurchasesDataInner(type_ string, id string) *AppRelationshipsInAppPurchasesDataInner {
	this := AppRelationshipsInAppPurchasesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsInAppPurchasesDataInnerWithDefaults instantiates a new AppRelationshipsInAppPurchasesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsInAppPurchasesDataInnerWithDefaults() *AppRelationshipsInAppPurchasesDataInner {
	this := AppRelationshipsInAppPurchasesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsInAppPurchasesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsInAppPurchasesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsInAppPurchasesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsInAppPurchasesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsInAppPurchasesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsInAppPurchasesDataInner) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsInAppPurchasesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsInAppPurchasesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AppRelationshipsInAppPurchasesDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varAppRelationshipsInAppPurchasesDataInner := _AppRelationshipsInAppPurchasesDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppRelationshipsInAppPurchasesDataInner)

	if err != nil {
		return err
	}

	*o = AppRelationshipsInAppPurchasesDataInner(varAppRelationshipsInAppPurchasesDataInner)

	return err
}

type NullableAppRelationshipsInAppPurchasesDataInner struct {
	value *AppRelationshipsInAppPurchasesDataInner
	isSet bool
}

func (v NullableAppRelationshipsInAppPurchasesDataInner) Get() *AppRelationshipsInAppPurchasesDataInner {
	return v.value
}

func (v *NullableAppRelationshipsInAppPurchasesDataInner) Set(val *AppRelationshipsInAppPurchasesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsInAppPurchasesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsInAppPurchasesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsInAppPurchasesDataInner(val *AppRelationshipsInAppPurchasesDataInner) *NullableAppRelationshipsInAppPurchasesDataInner {
	return &NullableAppRelationshipsInAppPurchasesDataInner{value: val, isSet: true}
}

func (v NullableAppRelationshipsInAppPurchasesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsInAppPurchasesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


