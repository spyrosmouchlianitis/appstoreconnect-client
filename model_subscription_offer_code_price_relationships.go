/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
)

// checks if the SubscriptionOfferCodePriceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodePriceRelationships{}

// SubscriptionOfferCodePriceRelationships struct for SubscriptionOfferCodePriceRelationships
type SubscriptionOfferCodePriceRelationships struct {
	Territory *AppPricePointV3RelationshipsTerritory `json:"territory,omitempty"`
	SubscriptionPricePoint *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint `json:"subscriptionPricePoint,omitempty"`
}

// NewSubscriptionOfferCodePriceRelationships instantiates a new SubscriptionOfferCodePriceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodePriceRelationships() *SubscriptionOfferCodePriceRelationships {
	this := SubscriptionOfferCodePriceRelationships{}
	return &this
}

// NewSubscriptionOfferCodePriceRelationshipsWithDefaults instantiates a new SubscriptionOfferCodePriceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodePriceRelationshipsWithDefaults() *SubscriptionOfferCodePriceRelationships {
	this := SubscriptionOfferCodePriceRelationships{}
	return &this
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *SubscriptionOfferCodePriceRelationships) GetTerritory() AppPricePointV3RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV3RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodePriceRelationships) GetTerritoryOk() (*AppPricePointV3RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *SubscriptionOfferCodePriceRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV3RelationshipsTerritory and assigns it to the Territory field.
func (o *SubscriptionOfferCodePriceRelationships) SetTerritory(v AppPricePointV3RelationshipsTerritory) {
	o.Territory = &v
}

// GetSubscriptionPricePoint returns the SubscriptionPricePoint field value if set, zero value otherwise.
func (o *SubscriptionOfferCodePriceRelationships) GetSubscriptionPricePoint() SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		var ret SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint
		return ret
	}
	return *o.SubscriptionPricePoint
}

// GetSubscriptionPricePointOk returns a tuple with the SubscriptionPricePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodePriceRelationships) GetSubscriptionPricePointOk() (*SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint, bool) {
	if o == nil || IsNil(o.SubscriptionPricePoint) {
		return nil, false
	}
	return o.SubscriptionPricePoint, true
}

// HasSubscriptionPricePoint returns a boolean if a field has been set.
func (o *SubscriptionOfferCodePriceRelationships) HasSubscriptionPricePoint() bool {
	if o != nil && !IsNil(o.SubscriptionPricePoint) {
		return true
	}

	return false
}

// SetSubscriptionPricePoint gets a reference to the given SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint and assigns it to the SubscriptionPricePoint field.
func (o *SubscriptionOfferCodePriceRelationships) SetSubscriptionPricePoint(v SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) {
	o.SubscriptionPricePoint = &v
}

func (o SubscriptionOfferCodePriceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodePriceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	if !IsNil(o.SubscriptionPricePoint) {
		toSerialize["subscriptionPricePoint"] = o.SubscriptionPricePoint
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodePriceRelationships struct {
	value *SubscriptionOfferCodePriceRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodePriceRelationships) Get() *SubscriptionOfferCodePriceRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodePriceRelationships) Set(val *SubscriptionOfferCodePriceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodePriceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodePriceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodePriceRelationships(val *SubscriptionOfferCodePriceRelationships) *NullableSubscriptionOfferCodePriceRelationships {
	return &NullableSubscriptionOfferCodePriceRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodePriceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodePriceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


