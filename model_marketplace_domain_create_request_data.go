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

// checks if the MarketplaceDomainCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceDomainCreateRequestData{}

// MarketplaceDomainCreateRequestData struct for MarketplaceDomainCreateRequestData
type MarketplaceDomainCreateRequestData struct {
	Type string `json:"type"`
	Attributes AlternativeDistributionDomainCreateRequestDataAttributes `json:"attributes"`
}

type _MarketplaceDomainCreateRequestData MarketplaceDomainCreateRequestData

// NewMarketplaceDomainCreateRequestData instantiates a new MarketplaceDomainCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceDomainCreateRequestData(type_ string, attributes AlternativeDistributionDomainCreateRequestDataAttributes) *MarketplaceDomainCreateRequestData {
	this := MarketplaceDomainCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMarketplaceDomainCreateRequestDataWithDefaults instantiates a new MarketplaceDomainCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceDomainCreateRequestDataWithDefaults() *MarketplaceDomainCreateRequestData {
	this := MarketplaceDomainCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *MarketplaceDomainCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketplaceDomainCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketplaceDomainCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MarketplaceDomainCreateRequestData) GetAttributes() AlternativeDistributionDomainCreateRequestDataAttributes {
	if o == nil {
		var ret AlternativeDistributionDomainCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MarketplaceDomainCreateRequestData) GetAttributesOk() (*AlternativeDistributionDomainCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MarketplaceDomainCreateRequestData) SetAttributes(v AlternativeDistributionDomainCreateRequestDataAttributes) {
	o.Attributes = v
}

func (o MarketplaceDomainCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceDomainCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *MarketplaceDomainCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varMarketplaceDomainCreateRequestData := _MarketplaceDomainCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceDomainCreateRequestData)

	if err != nil {
		return err
	}

	*o = MarketplaceDomainCreateRequestData(varMarketplaceDomainCreateRequestData)

	return err
}

type NullableMarketplaceDomainCreateRequestData struct {
	value *MarketplaceDomainCreateRequestData
	isSet bool
}

func (v NullableMarketplaceDomainCreateRequestData) Get() *MarketplaceDomainCreateRequestData {
	return v.value
}

func (v *NullableMarketplaceDomainCreateRequestData) Set(val *MarketplaceDomainCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceDomainCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceDomainCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceDomainCreateRequestData(val *MarketplaceDomainCreateRequestData) *NullableMarketplaceDomainCreateRequestData {
	return &NullableMarketplaceDomainCreateRequestData{value: val, isSet: true}
}

func (v NullableMarketplaceDomainCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceDomainCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


