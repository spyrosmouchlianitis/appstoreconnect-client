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

// checks if the InAppPurchasePricePointRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePricePointRelationships{}

// InAppPurchasePricePointRelationships struct for InAppPurchasePricePointRelationships
type InAppPurchasePricePointRelationships struct {
	Territory *AppPricePointV3RelationshipsTerritory `json:"territory,omitempty"`
}

// NewInAppPurchasePricePointRelationships instantiates a new InAppPurchasePricePointRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePricePointRelationships() *InAppPurchasePricePointRelationships {
	this := InAppPurchasePricePointRelationships{}
	return &this
}

// NewInAppPurchasePricePointRelationshipsWithDefaults instantiates a new InAppPurchasePricePointRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePricePointRelationshipsWithDefaults() *InAppPurchasePricePointRelationships {
	this := InAppPurchasePricePointRelationships{}
	return &this
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *InAppPurchasePricePointRelationships) GetTerritory() AppPricePointV3RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV3RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePricePointRelationships) GetTerritoryOk() (*AppPricePointV3RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *InAppPurchasePricePointRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV3RelationshipsTerritory and assigns it to the Territory field.
func (o *InAppPurchasePricePointRelationships) SetTerritory(v AppPricePointV3RelationshipsTerritory) {
	o.Territory = &v
}

func (o InAppPurchasePricePointRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePricePointRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	return toSerialize, nil
}

type NullableInAppPurchasePricePointRelationships struct {
	value *InAppPurchasePricePointRelationships
	isSet bool
}

func (v NullableInAppPurchasePricePointRelationships) Get() *InAppPurchasePricePointRelationships {
	return v.value
}

func (v *NullableInAppPurchasePricePointRelationships) Set(val *InAppPurchasePricePointRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePricePointRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePricePointRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePricePointRelationships(val *InAppPurchasePricePointRelationships) *NullableInAppPurchasePricePointRelationships {
	return &NullableInAppPurchasePricePointRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchasePricePointRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePricePointRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


