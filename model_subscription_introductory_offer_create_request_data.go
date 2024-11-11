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

// checks if the SubscriptionIntroductoryOfferCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferCreateRequestData{}

// SubscriptionIntroductoryOfferCreateRequestData struct for SubscriptionIntroductoryOfferCreateRequestData
type SubscriptionIntroductoryOfferCreateRequestData struct {
	Type string `json:"type"`
	Attributes SubscriptionIntroductoryOfferInlineCreateAttributes `json:"attributes"`
	Relationships SubscriptionIntroductoryOfferCreateRequestDataRelationships `json:"relationships"`
}

type _SubscriptionIntroductoryOfferCreateRequestData SubscriptionIntroductoryOfferCreateRequestData

// NewSubscriptionIntroductoryOfferCreateRequestData instantiates a new SubscriptionIntroductoryOfferCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferCreateRequestData(type_ string, attributes SubscriptionIntroductoryOfferInlineCreateAttributes, relationships SubscriptionIntroductoryOfferCreateRequestDataRelationships) *SubscriptionIntroductoryOfferCreateRequestData {
	this := SubscriptionIntroductoryOfferCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionIntroductoryOfferCreateRequestDataWithDefaults instantiates a new SubscriptionIntroductoryOfferCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferCreateRequestDataWithDefaults() *SubscriptionIntroductoryOfferCreateRequestData {
	this := SubscriptionIntroductoryOfferCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetAttributes() SubscriptionIntroductoryOfferInlineCreateAttributes {
	if o == nil {
		var ret SubscriptionIntroductoryOfferInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetAttributesOk() (*SubscriptionIntroductoryOfferInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) SetAttributes(v SubscriptionIntroductoryOfferInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetRelationships() SubscriptionIntroductoryOfferCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionIntroductoryOfferCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestData) GetRelationshipsOk() (*SubscriptionIntroductoryOfferCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionIntroductoryOfferCreateRequestData) SetRelationships(v SubscriptionIntroductoryOfferCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionIntroductoryOfferCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *SubscriptionIntroductoryOfferCreateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionIntroductoryOfferCreateRequestData := _SubscriptionIntroductoryOfferCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionIntroductoryOfferCreateRequestData)

	if err != nil {
		return err
	}

	*o = SubscriptionIntroductoryOfferCreateRequestData(varSubscriptionIntroductoryOfferCreateRequestData)

	return err
}

type NullableSubscriptionIntroductoryOfferCreateRequestData struct {
	value *SubscriptionIntroductoryOfferCreateRequestData
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestData) Get() *SubscriptionIntroductoryOfferCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestData) Set(val *SubscriptionIntroductoryOfferCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferCreateRequestData(val *SubscriptionIntroductoryOfferCreateRequestData) *NullableSubscriptionIntroductoryOfferCreateRequestData {
	return &NullableSubscriptionIntroductoryOfferCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


