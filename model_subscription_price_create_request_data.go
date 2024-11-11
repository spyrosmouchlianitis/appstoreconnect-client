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

// checks if the SubscriptionPriceCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPriceCreateRequestData{}

// SubscriptionPriceCreateRequestData struct for SubscriptionPriceCreateRequestData
type SubscriptionPriceCreateRequestData struct {
	Type string `json:"type"`
	Attributes *SubscriptionPriceInlineCreateAttributes `json:"attributes,omitempty"`
	Relationships SubscriptionPriceCreateRequestDataRelationships `json:"relationships"`
}

type _SubscriptionPriceCreateRequestData SubscriptionPriceCreateRequestData

// NewSubscriptionPriceCreateRequestData instantiates a new SubscriptionPriceCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPriceCreateRequestData(type_ string, relationships SubscriptionPriceCreateRequestDataRelationships) *SubscriptionPriceCreateRequestData {
	this := SubscriptionPriceCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewSubscriptionPriceCreateRequestDataWithDefaults instantiates a new SubscriptionPriceCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPriceCreateRequestDataWithDefaults() *SubscriptionPriceCreateRequestData {
	this := SubscriptionPriceCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPriceCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPriceCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionPriceCreateRequestData) GetAttributes() SubscriptionPriceInlineCreateAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionPriceInlineCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestData) GetAttributesOk() (*SubscriptionPriceInlineCreateAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionPriceCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionPriceInlineCreateAttributes and assigns it to the Attributes field.
func (o *SubscriptionPriceCreateRequestData) SetAttributes(v SubscriptionPriceInlineCreateAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionPriceCreateRequestData) GetRelationships() SubscriptionPriceCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionPriceCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestData) GetRelationshipsOk() (*SubscriptionPriceCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionPriceCreateRequestData) SetRelationships(v SubscriptionPriceCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionPriceCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPriceCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *SubscriptionPriceCreateRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSubscriptionPriceCreateRequestData := _SubscriptionPriceCreateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPriceCreateRequestData)

	if err != nil {
		return err
	}

	*o = SubscriptionPriceCreateRequestData(varSubscriptionPriceCreateRequestData)

	return err
}

type NullableSubscriptionPriceCreateRequestData struct {
	value *SubscriptionPriceCreateRequestData
	isSet bool
}

func (v NullableSubscriptionPriceCreateRequestData) Get() *SubscriptionPriceCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionPriceCreateRequestData) Set(val *SubscriptionPriceCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPriceCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPriceCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPriceCreateRequestData(val *SubscriptionPriceCreateRequestData) *NullableSubscriptionPriceCreateRequestData {
	return &NullableSubscriptionPriceCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionPriceCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPriceCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


