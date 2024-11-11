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

// checks if the SubscriptionIntroductoryOfferCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferCreateRequestDataRelationships{}

// SubscriptionIntroductoryOfferCreateRequestDataRelationships struct for SubscriptionIntroductoryOfferCreateRequestDataRelationships
type SubscriptionIntroductoryOfferCreateRequestDataRelationships struct {
	Subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription `json:"subscription"`
	Territory *AppPricePointV3RelationshipsTerritory `json:"territory,omitempty"`
	SubscriptionPricePoint *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint `json:"subscriptionPricePoint,omitempty"`
}

type _SubscriptionIntroductoryOfferCreateRequestDataRelationships SubscriptionIntroductoryOfferCreateRequestDataRelationships

// NewSubscriptionIntroductoryOfferCreateRequestDataRelationships instantiates a new SubscriptionIntroductoryOfferCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferCreateRequestDataRelationships(subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) *SubscriptionIntroductoryOfferCreateRequestDataRelationships {
	this := SubscriptionIntroductoryOfferCreateRequestDataRelationships{}
	this.Subscription = subscription
	return &this
}

// NewSubscriptionIntroductoryOfferCreateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionIntroductoryOfferCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferCreateRequestDataRelationshipsWithDefaults() *SubscriptionIntroductoryOfferCreateRequestDataRelationships {
	this := SubscriptionIntroductoryOfferCreateRequestDataRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetSubscription() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetSubscriptionOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) SetSubscription(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) {
	o.Subscription = v
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetTerritory() AppPricePointV3RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV3RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetTerritoryOk() (*AppPricePointV3RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV3RelationshipsTerritory and assigns it to the Territory field.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) SetTerritory(v AppPricePointV3RelationshipsTerritory) {
	o.Territory = &v
}

// GetSubscriptionPricePoint returns the SubscriptionPricePoint field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetSubscriptionPricePoint() SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		var ret SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint
		return ret
	}
	return *o.SubscriptionPricePoint
}

// GetSubscriptionPricePointOk returns a tuple with the SubscriptionPricePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) GetSubscriptionPricePointOk() (*SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint, bool) {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		return nil, false
	}
	return o.SubscriptionPricePoint, true
}

// HasSubscriptionPricePoint returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) HasSubscriptionPricePoint() bool {
	if o != nil && !IsNil(o.SubscriptionPricePoint) {
		return true
	}

	return false
}

// SetSubscriptionPricePoint gets a reference to the given SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint and assigns it to the SubscriptionPricePoint field.
func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) SetSubscriptionPricePoint(v SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) {
	o.SubscriptionPricePoint = &v
}

func (o SubscriptionIntroductoryOfferCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	if !IsNil(o.SubscriptionPricePoint) {
		toSerialize["subscriptionPricePoint"] = o.SubscriptionPricePoint
	}
	return toSerialize, nil
}

func (o *SubscriptionIntroductoryOfferCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscription",
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

	varSubscriptionIntroductoryOfferCreateRequestDataRelationships := _SubscriptionIntroductoryOfferCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionIntroductoryOfferCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = SubscriptionIntroductoryOfferCreateRequestDataRelationships(varSubscriptionIntroductoryOfferCreateRequestDataRelationships)

	return err
}

type NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships struct {
	value *SubscriptionIntroductoryOfferCreateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) Get() *SubscriptionIntroductoryOfferCreateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) Set(val *SubscriptionIntroductoryOfferCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferCreateRequestDataRelationships(val *SubscriptionIntroductoryOfferCreateRequestDataRelationships) *NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships {
	return &NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


