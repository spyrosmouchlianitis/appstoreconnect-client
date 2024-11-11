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

// checks if the SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription{}

// SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription struct for SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription
type SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription struct {
	Data PromotedPurchaseRelationshipsSubscriptionData `json:"data"`
}

type _SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription

// NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription instantiates a new SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(data PromotedPurchaseRelationshipsSubscriptionData) *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	this := SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription{}
	this.Data = data
	return &this
}

// NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscriptionWithDefaults instantiates a new SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscriptionWithDefaults() *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	this := SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) GetData() PromotedPurchaseRelationshipsSubscriptionData {
	if o == nil {
		var ret PromotedPurchaseRelationshipsSubscriptionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) GetDataOk() (*PromotedPurchaseRelationshipsSubscriptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) SetData(v PromotedPurchaseRelationshipsSubscriptionData) {
	o.Data = v
}

func (o SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription := _SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription)

	if err != nil {
		return err
	}

	*o = SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(varSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription)

	return err
}

type NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription struct {
	value *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription
	isSet bool
}

func (v NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) Get() *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	return v.value
}

func (v *NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) Set(val *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(val *SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) *NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	return &NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription{value: val, isSet: true}
}

func (v NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


