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

// checks if the SubscriptionGroupLocalizationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationCreateRequestDataRelationships{}

// SubscriptionGroupLocalizationCreateRequestDataRelationships struct for SubscriptionGroupLocalizationCreateRequestDataRelationships
type SubscriptionGroupLocalizationCreateRequestDataRelationships struct {
	SubscriptionGroup SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup `json:"subscriptionGroup"`
}

type _SubscriptionGroupLocalizationCreateRequestDataRelationships SubscriptionGroupLocalizationCreateRequestDataRelationships

// NewSubscriptionGroupLocalizationCreateRequestDataRelationships instantiates a new SubscriptionGroupLocalizationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationCreateRequestDataRelationships(subscriptionGroup SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) *SubscriptionGroupLocalizationCreateRequestDataRelationships {
	this := SubscriptionGroupLocalizationCreateRequestDataRelationships{}
	this.SubscriptionGroup = subscriptionGroup
	return &this
}

// NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionGroupLocalizationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsWithDefaults() *SubscriptionGroupLocalizationCreateRequestDataRelationships {
	this := SubscriptionGroupLocalizationCreateRequestDataRelationships{}
	return &this
}

// GetSubscriptionGroup returns the SubscriptionGroup field value
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationships) GetSubscriptionGroup() SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup {
	if o == nil {
		var ret SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup
		return ret
	}

	return o.SubscriptionGroup
}

// GetSubscriptionGroupOk returns a tuple with the SubscriptionGroup field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationships) GetSubscriptionGroupOk() (*SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionGroup, true
}

// SetSubscriptionGroup sets field value
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationships) SetSubscriptionGroup(v SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) {
	o.SubscriptionGroup = v
}

func (o SubscriptionGroupLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionGroup"] = o.SubscriptionGroup
	return toSerialize, nil
}

func (o *SubscriptionGroupLocalizationCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionGroup",
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

	varSubscriptionGroupLocalizationCreateRequestDataRelationships := _SubscriptionGroupLocalizationCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionGroupLocalizationCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = SubscriptionGroupLocalizationCreateRequestDataRelationships(varSubscriptionGroupLocalizationCreateRequestDataRelationships)

	return err
}

type NullableSubscriptionGroupLocalizationCreateRequestDataRelationships struct {
	value *SubscriptionGroupLocalizationCreateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) Get() *SubscriptionGroupLocalizationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) Set(val *SubscriptionGroupLocalizationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationCreateRequestDataRelationships(val *SubscriptionGroupLocalizationCreateRequestDataRelationships) *NullableSubscriptionGroupLocalizationCreateRequestDataRelationships {
	return &NullableSubscriptionGroupLocalizationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


