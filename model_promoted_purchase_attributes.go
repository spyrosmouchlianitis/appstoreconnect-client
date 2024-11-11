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

// checks if the PromotedPurchaseAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseAttributes{}

// PromotedPurchaseAttributes struct for PromotedPurchaseAttributes
type PromotedPurchaseAttributes struct {
	VisibleForAllUsers *bool `json:"visibleForAllUsers,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewPromotedPurchaseAttributes instantiates a new PromotedPurchaseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseAttributes() *PromotedPurchaseAttributes {
	this := PromotedPurchaseAttributes{}
	return &this
}

// NewPromotedPurchaseAttributesWithDefaults instantiates a new PromotedPurchaseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseAttributesWithDefaults() *PromotedPurchaseAttributes {
	this := PromotedPurchaseAttributes{}
	return &this
}

// GetVisibleForAllUsers returns the VisibleForAllUsers field value if set, zero value otherwise.
func (o *PromotedPurchaseAttributes) GetVisibleForAllUsers() bool {
	if o == nil || IsNil(o.VisibleForAllUsers) {
		var ret bool
		return ret
	}
	return *o.VisibleForAllUsers
}

// GetVisibleForAllUsersOk returns a tuple with the VisibleForAllUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseAttributes) GetVisibleForAllUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.VisibleForAllUsers) {
		return nil, false
	}
	return o.VisibleForAllUsers, true
}

// HasVisibleForAllUsers returns a boolean if a field has been set.
func (o *PromotedPurchaseAttributes) HasVisibleForAllUsers() bool {
	if o != nil && !IsNil(o.VisibleForAllUsers) {
		return true
	}

	return false
}

// SetVisibleForAllUsers gets a reference to the given bool and assigns it to the VisibleForAllUsers field.
func (o *PromotedPurchaseAttributes) SetVisibleForAllUsers(v bool) {
	o.VisibleForAllUsers = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PromotedPurchaseAttributes) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PromotedPurchaseAttributes) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PromotedPurchaseAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PromotedPurchaseAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PromotedPurchaseAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PromotedPurchaseAttributes) SetState(v string) {
	o.State = &v
}

func (o PromotedPurchaseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VisibleForAllUsers) {
		toSerialize["visibleForAllUsers"] = o.VisibleForAllUsers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseAttributes struct {
	value *PromotedPurchaseAttributes
	isSet bool
}

func (v NullablePromotedPurchaseAttributes) Get() *PromotedPurchaseAttributes {
	return v.value
}

func (v *NullablePromotedPurchaseAttributes) Set(val *PromotedPurchaseAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseAttributes(val *PromotedPurchaseAttributes) *NullablePromotedPurchaseAttributes {
	return &NullablePromotedPurchaseAttributes{value: val, isSet: true}
}

func (v NullablePromotedPurchaseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


