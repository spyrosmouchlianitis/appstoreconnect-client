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

// checks if the SubscriptionPriceCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPriceCreateRequestDataRelationships{}

// SubscriptionPriceCreateRequestDataRelationships struct for SubscriptionPriceCreateRequestDataRelationships
type SubscriptionPriceCreateRequestDataRelationships struct {
	Subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription `json:"subscription"`
	Territory *AppPricePointV3RelationshipsTerritory `json:"territory,omitempty"`
	SubscriptionPricePoint SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint `json:"subscriptionPricePoint"`
}

type _SubscriptionPriceCreateRequestDataRelationships SubscriptionPriceCreateRequestDataRelationships

// NewSubscriptionPriceCreateRequestDataRelationships instantiates a new SubscriptionPriceCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPriceCreateRequestDataRelationships(subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription, subscriptionPricePoint SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) *SubscriptionPriceCreateRequestDataRelationships {
	this := SubscriptionPriceCreateRequestDataRelationships{}
	this.Subscription = subscription
	this.SubscriptionPricePoint = subscriptionPricePoint
	return &this
}

// NewSubscriptionPriceCreateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionPriceCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPriceCreateRequestDataRelationshipsWithDefaults() *SubscriptionPriceCreateRequestDataRelationships {
	this := SubscriptionPriceCreateRequestDataRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *SubscriptionPriceCreateRequestDataRelationships) GetSubscription() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestDataRelationships) GetSubscriptionOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SubscriptionPriceCreateRequestDataRelationships) SetSubscription(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) {
	o.Subscription = v
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *SubscriptionPriceCreateRequestDataRelationships) GetTerritory() AppPricePointV3RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV3RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestDataRelationships) GetTerritoryOk() (*AppPricePointV3RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *SubscriptionPriceCreateRequestDataRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV3RelationshipsTerritory and assigns it to the Territory field.
func (o *SubscriptionPriceCreateRequestDataRelationships) SetTerritory(v AppPricePointV3RelationshipsTerritory) {
	o.Territory = &v
}

// GetSubscriptionPricePoint returns the SubscriptionPricePoint field value
func (o *SubscriptionPriceCreateRequestDataRelationships) GetSubscriptionPricePoint() SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint {
	if o == nil {
		var ret SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint
		return ret
	}

	return o.SubscriptionPricePoint
}

// GetSubscriptionPricePointOk returns a tuple with the SubscriptionPricePoint field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestDataRelationships) GetSubscriptionPricePointOk() (*SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionPricePoint, true
}

// SetSubscriptionPricePoint sets field value
func (o *SubscriptionPriceCreateRequestDataRelationships) SetSubscriptionPricePoint(v SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) {
	o.SubscriptionPricePoint = v
}

func (o SubscriptionPriceCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPriceCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	toSerialize["subscriptionPricePoint"] = o.SubscriptionPricePoint
	return toSerialize, nil
}

func (o *SubscriptionPriceCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscription",
		"subscriptionPricePoint",
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

	varSubscriptionPriceCreateRequestDataRelationships := _SubscriptionPriceCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPriceCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = SubscriptionPriceCreateRequestDataRelationships(varSubscriptionPriceCreateRequestDataRelationships)

	return err
}

type NullableSubscriptionPriceCreateRequestDataRelationships struct {
	value *SubscriptionPriceCreateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionPriceCreateRequestDataRelationships) Get() *SubscriptionPriceCreateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationships) Set(val *SubscriptionPriceCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPriceCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPriceCreateRequestDataRelationships(val *SubscriptionPriceCreateRequestDataRelationships) *NullableSubscriptionPriceCreateRequestDataRelationships {
	return &NullableSubscriptionPriceCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionPriceCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

