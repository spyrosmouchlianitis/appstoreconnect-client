/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"encoding/json"
)

// checks if the AppAvailabilityV2Relationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2Relationships{}

// AppAvailabilityV2Relationships struct for AppAvailabilityV2Relationships
type AppAvailabilityV2Relationships struct {
	TerritoryAvailabilities *AppAvailabilityV2RelationshipsTerritoryAvailabilities `json:"territoryAvailabilities,omitempty"`
}

// NewAppAvailabilityV2Relationships instantiates a new AppAvailabilityV2Relationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2Relationships() *AppAvailabilityV2Relationships {
	this := AppAvailabilityV2Relationships{}
	return &this
}

// NewAppAvailabilityV2RelationshipsWithDefaults instantiates a new AppAvailabilityV2Relationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2RelationshipsWithDefaults() *AppAvailabilityV2Relationships {
	this := AppAvailabilityV2Relationships{}
	return &this
}

// GetTerritoryAvailabilities returns the TerritoryAvailabilities field value if set, zero value otherwise.
func (o *AppAvailabilityV2Relationships) GetTerritoryAvailabilities() AppAvailabilityV2RelationshipsTerritoryAvailabilities {
	if o == nil || IsNil(o.TerritoryAvailabilities) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilities
		return ret
	}
	return *o.TerritoryAvailabilities
}

// GetTerritoryAvailabilitiesOk returns a tuple with the TerritoryAvailabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2Relationships) GetTerritoryAvailabilitiesOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilities, bool) {
	if o == nil || IsNil(o.TerritoryAvailabilities) {
		return nil, false
	}
	return o.TerritoryAvailabilities, true
}

// HasTerritoryAvailabilities returns a boolean if a field has been set.
func (o *AppAvailabilityV2Relationships) HasTerritoryAvailabilities() bool {
	if o != nil && !IsNil(o.TerritoryAvailabilities) {
		return true
	}

	return false
}

// SetTerritoryAvailabilities gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilities and assigns it to the TerritoryAvailabilities field.
func (o *AppAvailabilityV2Relationships) SetTerritoryAvailabilities(v AppAvailabilityV2RelationshipsTerritoryAvailabilities) {
	o.TerritoryAvailabilities = &v
}

func (o AppAvailabilityV2Relationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2Relationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TerritoryAvailabilities) {
		toSerialize["territoryAvailabilities"] = o.TerritoryAvailabilities
	}
	return toSerialize, nil
}

type NullableAppAvailabilityV2Relationships struct {
	value *AppAvailabilityV2Relationships
	isSet bool
}

func (v NullableAppAvailabilityV2Relationships) Get() *AppAvailabilityV2Relationships {
	return v.value
}

func (v *NullableAppAvailabilityV2Relationships) Set(val *AppAvailabilityV2Relationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2Relationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2Relationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2Relationships(val *AppAvailabilityV2Relationships) *NullableAppAvailabilityV2Relationships {
	return &NullableAppAvailabilityV2Relationships{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2Relationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2Relationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


