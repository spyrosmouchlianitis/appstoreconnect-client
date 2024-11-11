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

// checks if the MarketplaceSearchDetailUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceSearchDetailUpdateRequestData{}

// MarketplaceSearchDetailUpdateRequestData struct for MarketplaceSearchDetailUpdateRequestData
type MarketplaceSearchDetailUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *MarketplaceSearchDetailAttributes `json:"attributes,omitempty"`
}

type _MarketplaceSearchDetailUpdateRequestData MarketplaceSearchDetailUpdateRequestData

// NewMarketplaceSearchDetailUpdateRequestData instantiates a new MarketplaceSearchDetailUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceSearchDetailUpdateRequestData(type_ string, id string) *MarketplaceSearchDetailUpdateRequestData {
	this := MarketplaceSearchDetailUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewMarketplaceSearchDetailUpdateRequestDataWithDefaults instantiates a new MarketplaceSearchDetailUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceSearchDetailUpdateRequestDataWithDefaults() *MarketplaceSearchDetailUpdateRequestData {
	this := MarketplaceSearchDetailUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *MarketplaceSearchDetailUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketplaceSearchDetailUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketplaceSearchDetailUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *MarketplaceSearchDetailUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MarketplaceSearchDetailUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MarketplaceSearchDetailUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MarketplaceSearchDetailUpdateRequestData) GetAttributes() MarketplaceSearchDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret MarketplaceSearchDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceSearchDetailUpdateRequestData) GetAttributesOk() (*MarketplaceSearchDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MarketplaceSearchDetailUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MarketplaceSearchDetailAttributes and assigns it to the Attributes field.
func (o *MarketplaceSearchDetailUpdateRequestData) SetAttributes(v MarketplaceSearchDetailAttributes) {
	o.Attributes = &v
}

func (o MarketplaceSearchDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceSearchDetailUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *MarketplaceSearchDetailUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varMarketplaceSearchDetailUpdateRequestData := _MarketplaceSearchDetailUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceSearchDetailUpdateRequestData)

	if err != nil {
		return err
	}

	*o = MarketplaceSearchDetailUpdateRequestData(varMarketplaceSearchDetailUpdateRequestData)

	return err
}

type NullableMarketplaceSearchDetailUpdateRequestData struct {
	value *MarketplaceSearchDetailUpdateRequestData
	isSet bool
}

func (v NullableMarketplaceSearchDetailUpdateRequestData) Get() *MarketplaceSearchDetailUpdateRequestData {
	return v.value
}

func (v *NullableMarketplaceSearchDetailUpdateRequestData) Set(val *MarketplaceSearchDetailUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceSearchDetailUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceSearchDetailUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceSearchDetailUpdateRequestData(val *MarketplaceSearchDetailUpdateRequestData) *NullableMarketplaceSearchDetailUpdateRequestData {
	return &NullableMarketplaceSearchDetailUpdateRequestData{value: val, isSet: true}
}

func (v NullableMarketplaceSearchDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceSearchDetailUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


