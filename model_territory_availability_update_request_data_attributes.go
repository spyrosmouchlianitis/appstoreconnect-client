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

// checks if the TerritoryAvailabilityUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryAvailabilityUpdateRequestDataAttributes{}

// TerritoryAvailabilityUpdateRequestDataAttributes struct for TerritoryAvailabilityUpdateRequestDataAttributes
type TerritoryAvailabilityUpdateRequestDataAttributes struct {
	Available *bool `json:"available,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
	PreOrderEnabled *bool `json:"preOrderEnabled,omitempty"`
}

// NewTerritoryAvailabilityUpdateRequestDataAttributes instantiates a new TerritoryAvailabilityUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryAvailabilityUpdateRequestDataAttributes() *TerritoryAvailabilityUpdateRequestDataAttributes {
	this := TerritoryAvailabilityUpdateRequestDataAttributes{}
	return &this
}

// NewTerritoryAvailabilityUpdateRequestDataAttributesWithDefaults instantiates a new TerritoryAvailabilityUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryAvailabilityUpdateRequestDataAttributesWithDefaults() *TerritoryAvailabilityUpdateRequestDataAttributes {
	this := TerritoryAvailabilityUpdateRequestDataAttributes{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) SetAvailable(v bool) {
	o.Available = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetPreOrderEnabled returns the PreOrderEnabled field value if set, zero value otherwise.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetPreOrderEnabled() bool {
	if o == nil || IsNil(o.PreOrderEnabled) {
		var ret bool
		return ret
	}
	return *o.PreOrderEnabled
}

// GetPreOrderEnabledOk returns a tuple with the PreOrderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) GetPreOrderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PreOrderEnabled) {
		return nil, false
	}
	return o.PreOrderEnabled, true
}

// HasPreOrderEnabled returns a boolean if a field has been set.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) HasPreOrderEnabled() bool {
	if o != nil && !IsNil(o.PreOrderEnabled) {
		return true
	}

	return false
}

// SetPreOrderEnabled gets a reference to the given bool and assigns it to the PreOrderEnabled field.
func (o *TerritoryAvailabilityUpdateRequestDataAttributes) SetPreOrderEnabled(v bool) {
	o.PreOrderEnabled = &v
}

func (o TerritoryAvailabilityUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryAvailabilityUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.PreOrderEnabled) {
		toSerialize["preOrderEnabled"] = o.PreOrderEnabled
	}
	return toSerialize, nil
}

type NullableTerritoryAvailabilityUpdateRequestDataAttributes struct {
	value *TerritoryAvailabilityUpdateRequestDataAttributes
	isSet bool
}

func (v NullableTerritoryAvailabilityUpdateRequestDataAttributes) Get() *TerritoryAvailabilityUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableTerritoryAvailabilityUpdateRequestDataAttributes) Set(val *TerritoryAvailabilityUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryAvailabilityUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryAvailabilityUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryAvailabilityUpdateRequestDataAttributes(val *TerritoryAvailabilityUpdateRequestDataAttributes) *NullableTerritoryAvailabilityUpdateRequestDataAttributes {
	return &NullableTerritoryAvailabilityUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableTerritoryAvailabilityUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryAvailabilityUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


