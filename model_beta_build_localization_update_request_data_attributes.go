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

// checks if the BetaBuildLocalizationUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationUpdateRequestDataAttributes{}

// BetaBuildLocalizationUpdateRequestDataAttributes struct for BetaBuildLocalizationUpdateRequestDataAttributes
type BetaBuildLocalizationUpdateRequestDataAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
}

// NewBetaBuildLocalizationUpdateRequestDataAttributes instantiates a new BetaBuildLocalizationUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationUpdateRequestDataAttributes() *BetaBuildLocalizationUpdateRequestDataAttributes {
	this := BetaBuildLocalizationUpdateRequestDataAttributes{}
	return &this
}

// NewBetaBuildLocalizationUpdateRequestDataAttributesWithDefaults instantiates a new BetaBuildLocalizationUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationUpdateRequestDataAttributesWithDefaults() *BetaBuildLocalizationUpdateRequestDataAttributes {
	this := BetaBuildLocalizationUpdateRequestDataAttributes{}
	return &this
}

// GetWhatsNew returns the WhatsNew field value if set, zero value otherwise.
func (o *BetaBuildLocalizationUpdateRequestDataAttributes) GetWhatsNew() string {
	if o == nil || IsNil(o.WhatsNew) {
		var ret string
		return ret
	}
	return *o.WhatsNew
}

// GetWhatsNewOk returns a tuple with the WhatsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationUpdateRequestDataAttributes) GetWhatsNewOk() (*string, bool) {
	if o == nil || IsNil(o.WhatsNew) {
		return nil, false
	}
	return o.WhatsNew, true
}

// HasWhatsNew returns a boolean if a field has been set.
func (o *BetaBuildLocalizationUpdateRequestDataAttributes) HasWhatsNew() bool {
	if o != nil && !IsNil(o.WhatsNew) {
		return true
	}

	return false
}

// SetWhatsNew gets a reference to the given string and assigns it to the WhatsNew field.
func (o *BetaBuildLocalizationUpdateRequestDataAttributes) SetWhatsNew(v string) {
	o.WhatsNew = &v
}

func (o BetaBuildLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WhatsNew) {
		toSerialize["whatsNew"] = o.WhatsNew
	}
	return toSerialize, nil
}

type NullableBetaBuildLocalizationUpdateRequestDataAttributes struct {
	value *BetaBuildLocalizationUpdateRequestDataAttributes
	isSet bool
}

func (v NullableBetaBuildLocalizationUpdateRequestDataAttributes) Get() *BetaBuildLocalizationUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaBuildLocalizationUpdateRequestDataAttributes) Set(val *BetaBuildLocalizationUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationUpdateRequestDataAttributes(val *BetaBuildLocalizationUpdateRequestDataAttributes) *NullableBetaBuildLocalizationUpdateRequestDataAttributes {
	return &NullableBetaBuildLocalizationUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


