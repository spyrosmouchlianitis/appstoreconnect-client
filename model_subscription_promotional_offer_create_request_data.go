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

// checks if the SubscriptionPromotionalOfferCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferCreateRequestData{}

// SubscriptionPromotionalOfferCreateRequestData struct for SubscriptionPromotionalOfferCreateRequestData
type SubscriptionPromotionalOfferCreateRequestData struct {
	Type string `json:"type"`
	Attributes SubscriptionPromotionalOfferInlineCreateAttributes `json:"attributes"`
	Relationships SubscriptionPromotionalOfferCreateRequestDataRelationships `json:"relationships"`
}

type _SubscriptionPromotionalOfferCreateRequestData SubscriptionPromotionalOfferCreateRequestData

// NewSubscriptionPromotionalOfferCreateRequestData instantiates a new SubscriptionPromotionalOfferCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferCreateRequestData(type_ string, attributes SubscriptionPromotionalOfferInlineCreateAttributes, relationships SubscriptionPromotionalOfferCreateRequestDataRelationships) *SubscriptionPromotionalOfferCreateRequestData {
	this := SubscriptionPromotionalOfferCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionPromotionalOfferCreateRequestDataWithDefaults instantiates a new SubscriptionPromotionalOfferCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferCreateRequestDataWithDefaults() *SubscriptionPromotionalOfferCreateRequestData {
	this := SubscriptionPromotionalOfferCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPromotionalOfferCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPromotionalOfferCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionPromotionalOfferCreateRequestData) GetAttributes() SubscriptionPromotionalOfferInlineCreateAttributes {
	if o == nil {
		var ret SubscriptionPromotionalOfferInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferCreateRequestData) GetAttributesOk() (*SubscriptionPromotionalOfferInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionPromotionalOfferCreateRequestData) SetAttributes(v SubscriptionPromotionalOfferInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionPromotionalOfferCreateRequestData) GetRelationships() SubscriptionPromotionalOfferCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionPromotionalOfferCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferCreateRequestData) GetRelationshipsOk() (*SubscriptionPromotionalOfferCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionPromotionalOfferCreateRequestData) SetRelationships(v SubscriptionPromotionalOfferCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionPromotionalOfferCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *SubscriptionPromotionalOfferCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
		"relationships",
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

	varSubscriptionPromotionalOfferCreateRequestData := _SubscriptionPromotionalOfferCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPromotionalOfferCreateRequestData)

	if err != nil {
		return err
	}

	*o = SubscriptionPromotionalOfferCreateRequestData(varSubscriptionPromotionalOfferCreateRequestData)

	return err
}

type NullableSubscriptionPromotionalOfferCreateRequestData struct {
	value *SubscriptionPromotionalOfferCreateRequestData
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferCreateRequestData) Get() *SubscriptionPromotionalOfferCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferCreateRequestData) Set(val *SubscriptionPromotionalOfferCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferCreateRequestData(val *SubscriptionPromotionalOfferCreateRequestData) *NullableSubscriptionPromotionalOfferCreateRequestData {
	return &NullableSubscriptionPromotionalOfferCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


