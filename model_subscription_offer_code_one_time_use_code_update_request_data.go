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

// checks if the SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData{}

// SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData struct for SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData
type SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData

// NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData instantiates a new SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData(type_ string, id string) *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData {
	this := SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestDataWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestDataWithDefaults() *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData {
	this := SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetAttributes() SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) GetAttributesOk() (*SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) SetAttributes(v SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData := _SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData)

	if err != nil {
		return err
	}

	*o = SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData(varSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData)

	return err
}

type NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData struct {
	value *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) Get() *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) Set(val *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData(val *SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData {
	return &NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


