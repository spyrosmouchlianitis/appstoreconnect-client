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

// checks if the BetaBuildLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationUpdateRequestData{}

// BetaBuildLocalizationUpdateRequestData struct for BetaBuildLocalizationUpdateRequestData
type BetaBuildLocalizationUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaBuildLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

type _BetaBuildLocalizationUpdateRequestData BetaBuildLocalizationUpdateRequestData

// NewBetaBuildLocalizationUpdateRequestData instantiates a new BetaBuildLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationUpdateRequestData(type_ string, id string) *BetaBuildLocalizationUpdateRequestData {
	this := BetaBuildLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaBuildLocalizationUpdateRequestDataWithDefaults instantiates a new BetaBuildLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationUpdateRequestDataWithDefaults() *BetaBuildLocalizationUpdateRequestData {
	this := BetaBuildLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaBuildLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaBuildLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaBuildLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaBuildLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaBuildLocalizationUpdateRequestData) GetAttributes() BetaBuildLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaBuildLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationUpdateRequestData) GetAttributesOk() (*BetaBuildLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaBuildLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaBuildLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *BetaBuildLocalizationUpdateRequestData) SetAttributes(v BetaBuildLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o BetaBuildLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *BetaBuildLocalizationUpdateRequestData) UnmarshalJSON(data []byte) (err error) {
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

	varBetaBuildLocalizationUpdateRequestData := _BetaBuildLocalizationUpdateRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBetaBuildLocalizationUpdateRequestData)

	if err != nil {
		return err
	}

	*o = BetaBuildLocalizationUpdateRequestData(varBetaBuildLocalizationUpdateRequestData)

	return err
}

type NullableBetaBuildLocalizationUpdateRequestData struct {
	value *BetaBuildLocalizationUpdateRequestData
	isSet bool
}

func (v NullableBetaBuildLocalizationUpdateRequestData) Get() *BetaBuildLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableBetaBuildLocalizationUpdateRequestData) Set(val *BetaBuildLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationUpdateRequestData(val *BetaBuildLocalizationUpdateRequestData) *NullableBetaBuildLocalizationUpdateRequestData {
	return &NullableBetaBuildLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


