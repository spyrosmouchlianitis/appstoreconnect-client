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

// checks if the SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes{}

// SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes struct for SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes
type SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes struct {
	NumberOfCodes int32 `json:"numberOfCodes"`
	ExpirationDate string `json:"expirationDate"`
}

type _SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes

// NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes instantiates a new SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes(numberOfCodes int32, expirationDate string) *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes{}
	this.NumberOfCodes = numberOfCodes
	this.ExpirationDate = expirationDate
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributesWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributesWithDefaults() *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes{}
	return &this
}

// GetNumberOfCodes returns the NumberOfCodes field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) GetNumberOfCodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCodes
}

// GetNumberOfCodesOk returns a tuple with the NumberOfCodes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) GetNumberOfCodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfCodes, true
}

// SetNumberOfCodes sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) SetNumberOfCodes(v int32) {
	o.NumberOfCodes = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

func (o SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numberOfCodes"] = o.NumberOfCodes
	toSerialize["expirationDate"] = o.ExpirationDate
	return toSerialize, nil
}

func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numberOfCodes",
		"expirationDate",
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

	varSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes := _SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes(varSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes)

	return err
}

type NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes struct {
	value *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) Get() *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) Set(val *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes(val *SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes {
	return &NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


