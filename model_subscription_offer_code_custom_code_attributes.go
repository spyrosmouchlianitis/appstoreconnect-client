/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"time"
)

// checks if the SubscriptionOfferCodeCustomCodeAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeAttributes{}

// SubscriptionOfferCodeCustomCodeAttributes struct for SubscriptionOfferCodeCustomCodeAttributes
type SubscriptionOfferCodeCustomCodeAttributes struct {
	CustomCode *string `json:"customCode,omitempty"`
	NumberOfCodes *int32 `json:"numberOfCodes,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewSubscriptionOfferCodeCustomCodeAttributes instantiates a new SubscriptionOfferCodeCustomCodeAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeAttributes() *SubscriptionOfferCodeCustomCodeAttributes {
	this := SubscriptionOfferCodeCustomCodeAttributes{}
	return &this
}

// NewSubscriptionOfferCodeCustomCodeAttributesWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeAttributesWithDefaults() *SubscriptionOfferCodeCustomCodeAttributes {
	this := SubscriptionOfferCodeCustomCodeAttributes{}
	return &this
}

// GetCustomCode returns the CustomCode field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetCustomCode() string {
	if o == nil || IsNil(o.CustomCode) {
		var ret string
		return ret
	}
	return *o.CustomCode
}

// GetCustomCodeOk returns a tuple with the CustomCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetCustomCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomCode) {
		return nil, false
	}
	return o.CustomCode, true
}

// HasCustomCode returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) HasCustomCode() bool {
	if o != nil && !IsNil(o.CustomCode) {
		return true
	}

	return false
}

// SetCustomCode gets a reference to the given string and assigns it to the CustomCode field.
func (o *SubscriptionOfferCodeCustomCodeAttributes) SetCustomCode(v string) {
	o.CustomCode = &v
}

// GetNumberOfCodes returns the NumberOfCodes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetNumberOfCodes() int32 {
	if o == nil || IsNil(o.NumberOfCodes) {
		var ret int32
		return ret
	}
	return *o.NumberOfCodes
}

// GetNumberOfCodesOk returns a tuple with the NumberOfCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetNumberOfCodesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfCodes) {
		return nil, false
	}
	return o.NumberOfCodes, true
}

// HasNumberOfCodes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) HasNumberOfCodes() bool {
	if o != nil && !IsNil(o.NumberOfCodes) {
		return true
	}

	return false
}

// SetNumberOfCodes gets a reference to the given int32 and assigns it to the NumberOfCodes field.
func (o *SubscriptionOfferCodeCustomCodeAttributes) SetNumberOfCodes(v int32) {
	o.NumberOfCodes = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *SubscriptionOfferCodeCustomCodeAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *SubscriptionOfferCodeCustomCodeAttributes) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeAttributes) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionOfferCodeCustomCodeAttributes) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionOfferCodeCustomCodeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomCode) {
		toSerialize["customCode"] = o.CustomCode
	}
	if !IsNil(o.NumberOfCodes) {
		toSerialize["numberOfCodes"] = o.NumberOfCodes
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeAttributes struct {
	value *SubscriptionOfferCodeCustomCodeAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeAttributes) Get() *SubscriptionOfferCodeCustomCodeAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeAttributes) Set(val *SubscriptionOfferCodeCustomCodeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeAttributes(val *SubscriptionOfferCodeCustomCodeAttributes) *NullableSubscriptionOfferCodeCustomCodeAttributes {
	return &NullableSubscriptionOfferCodeCustomCodeAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


