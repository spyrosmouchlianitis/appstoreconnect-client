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

// checks if the MarketplaceWebhookCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceWebhookCreateRequestData{}

// MarketplaceWebhookCreateRequestData struct for MarketplaceWebhookCreateRequestData
type MarketplaceWebhookCreateRequestData struct {
	Type string `json:"type"`
	Attributes MarketplaceWebhookCreateRequestDataAttributes `json:"attributes"`
}

type _MarketplaceWebhookCreateRequestData MarketplaceWebhookCreateRequestData

// NewMarketplaceWebhookCreateRequestData instantiates a new MarketplaceWebhookCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceWebhookCreateRequestData(type_ string, attributes MarketplaceWebhookCreateRequestDataAttributes) *MarketplaceWebhookCreateRequestData {
	this := MarketplaceWebhookCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMarketplaceWebhookCreateRequestDataWithDefaults instantiates a new MarketplaceWebhookCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceWebhookCreateRequestDataWithDefaults() *MarketplaceWebhookCreateRequestData {
	this := MarketplaceWebhookCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *MarketplaceWebhookCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketplaceWebhookCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketplaceWebhookCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MarketplaceWebhookCreateRequestData) GetAttributes() MarketplaceWebhookCreateRequestDataAttributes {
	if o == nil {
		var ret MarketplaceWebhookCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MarketplaceWebhookCreateRequestData) GetAttributesOk() (*MarketplaceWebhookCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MarketplaceWebhookCreateRequestData) SetAttributes(v MarketplaceWebhookCreateRequestDataAttributes) {
	o.Attributes = v
}

func (o MarketplaceWebhookCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceWebhookCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *MarketplaceWebhookCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varMarketplaceWebhookCreateRequestData := _MarketplaceWebhookCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketplaceWebhookCreateRequestData)

	if err != nil {
		return err
	}

	*o = MarketplaceWebhookCreateRequestData(varMarketplaceWebhookCreateRequestData)

	return err
}

type NullableMarketplaceWebhookCreateRequestData struct {
	value *MarketplaceWebhookCreateRequestData
	isSet bool
}

func (v NullableMarketplaceWebhookCreateRequestData) Get() *MarketplaceWebhookCreateRequestData {
	return v.value
}

func (v *NullableMarketplaceWebhookCreateRequestData) Set(val *MarketplaceWebhookCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceWebhookCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceWebhookCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceWebhookCreateRequestData(val *MarketplaceWebhookCreateRequestData) *NullableMarketplaceWebhookCreateRequestData {
	return &NullableMarketplaceWebhookCreateRequestData{value: val, isSet: true}
}

func (v NullableMarketplaceWebhookCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceWebhookCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

