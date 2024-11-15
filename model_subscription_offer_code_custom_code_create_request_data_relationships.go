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

// checks if the SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships{}

// SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships struct for SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships
type SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships struct {
	OfferCode SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode `json:"offerCode"`
}

type _SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships(offerCode SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships{}
	this.OfferCode = offerCode
	return &this
}

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsWithDefaults() *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships{}
	return &this
}

// GetOfferCode returns the OfferCode field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) GetOfferCode() SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode {
	if o == nil {
		var ret SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode
		return ret
	}

	return o.OfferCode
}

// GetOfferCodeOk returns a tuple with the OfferCode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) GetOfferCodeOk() (*SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferCode, true
}

// SetOfferCode sets field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) SetOfferCode(v SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) {
	o.OfferCode = v
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerCode"] = o.OfferCode
	return toSerialize, nil
}

func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerCode",
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

	varSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships := _SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships(varSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships)

	return err
}

type NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships struct {
	value *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) Get() *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) Set(val *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships(val *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships {
	return &NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


