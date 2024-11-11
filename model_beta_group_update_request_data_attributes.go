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

// checks if the BetaGroupUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupUpdateRequestDataAttributes{}

// BetaGroupUpdateRequestDataAttributes struct for BetaGroupUpdateRequestDataAttributes
type BetaGroupUpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	PublicLinkEnabled *bool `json:"publicLinkEnabled,omitempty"`
	PublicLinkLimitEnabled *bool `json:"publicLinkLimitEnabled,omitempty"`
	PublicLinkLimit *int32 `json:"publicLinkLimit,omitempty"`
	FeedbackEnabled *bool `json:"feedbackEnabled,omitempty"`
	IosBuildsAvailableForAppleSiliconMac *bool `json:"iosBuildsAvailableForAppleSiliconMac,omitempty"`
}

// NewBetaGroupUpdateRequestDataAttributes instantiates a new BetaGroupUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupUpdateRequestDataAttributes() *BetaGroupUpdateRequestDataAttributes {
	this := BetaGroupUpdateRequestDataAttributes{}
	return &this
}

// NewBetaGroupUpdateRequestDataAttributesWithDefaults instantiates a new BetaGroupUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupUpdateRequestDataAttributesWithDefaults() *BetaGroupUpdateRequestDataAttributes {
	this := BetaGroupUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BetaGroupUpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublicLinkEnabled returns the PublicLinkEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkEnabled() bool {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkEnabled
}

// GetPublicLinkEnabledOk returns a tuple with the PublicLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		return nil, false
	}
	return o.PublicLinkEnabled, true
}

// HasPublicLinkEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkEnabled) {
		return true
	}

	return false
}

// SetPublicLinkEnabled gets a reference to the given bool and assigns it to the PublicLinkEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkEnabled(v bool) {
	o.PublicLinkEnabled = &v
}

// GetPublicLinkLimitEnabled returns the PublicLinkLimitEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitEnabled() bool {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkLimitEnabled
}

// GetPublicLinkLimitEnabledOk returns a tuple with the PublicLinkLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		return nil, false
	}
	return o.PublicLinkLimitEnabled, true
}

// HasPublicLinkLimitEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkLimitEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkLimitEnabled) {
		return true
	}

	return false
}

// SetPublicLinkLimitEnabled gets a reference to the given bool and assigns it to the PublicLinkLimitEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkLimitEnabled(v bool) {
	o.PublicLinkLimitEnabled = &v
}

// GetPublicLinkLimit returns the PublicLinkLimit field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimit() int32 {
	if o == nil || IsNil(o.PublicLinkLimit) {
		var ret int32
		return ret
	}
	return *o.PublicLinkLimit
}

// GetPublicLinkLimitOk returns a tuple with the PublicLinkLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetPublicLinkLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicLinkLimit) {
		return nil, false
	}
	return o.PublicLinkLimit, true
}

// HasPublicLinkLimit returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasPublicLinkLimit() bool {
	if o != nil && !IsNil(o.PublicLinkLimit) {
		return true
	}

	return false
}

// SetPublicLinkLimit gets a reference to the given int32 and assigns it to the PublicLinkLimit field.
func (o *BetaGroupUpdateRequestDataAttributes) SetPublicLinkLimit(v int32) {
	o.PublicLinkLimit = &v
}

// GetFeedbackEnabled returns the FeedbackEnabled field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetFeedbackEnabled() bool {
	if o == nil || IsNil(o.FeedbackEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackEnabled
}

// GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetFeedbackEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackEnabled) {
		return nil, false
	}
	return o.FeedbackEnabled, true
}

// HasFeedbackEnabled returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasFeedbackEnabled() bool {
	if o != nil && !IsNil(o.FeedbackEnabled) {
		return true
	}

	return false
}

// SetFeedbackEnabled gets a reference to the given bool and assigns it to the FeedbackEnabled field.
func (o *BetaGroupUpdateRequestDataAttributes) SetFeedbackEnabled(v bool) {
	o.FeedbackEnabled = &v
}

// GetIosBuildsAvailableForAppleSiliconMac returns the IosBuildsAvailableForAppleSiliconMac field value if set, zero value otherwise.
func (o *BetaGroupUpdateRequestDataAttributes) GetIosBuildsAvailableForAppleSiliconMac() bool {
	if o == nil || IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		var ret bool
		return ret
	}
	return *o.IosBuildsAvailableForAppleSiliconMac
}

// GetIosBuildsAvailableForAppleSiliconMacOk returns a tuple with the IosBuildsAvailableForAppleSiliconMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupUpdateRequestDataAttributes) GetIosBuildsAvailableForAppleSiliconMacOk() (*bool, bool) {
	if o == nil || IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		return nil, false
	}
	return o.IosBuildsAvailableForAppleSiliconMac, true
}

// HasIosBuildsAvailableForAppleSiliconMac returns a boolean if a field has been set.
func (o *BetaGroupUpdateRequestDataAttributes) HasIosBuildsAvailableForAppleSiliconMac() bool {
	if o != nil && !IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		return true
	}

	return false
}

// SetIosBuildsAvailableForAppleSiliconMac gets a reference to the given bool and assigns it to the IosBuildsAvailableForAppleSiliconMac field.
func (o *BetaGroupUpdateRequestDataAttributes) SetIosBuildsAvailableForAppleSiliconMac(v bool) {
	o.IosBuildsAvailableForAppleSiliconMac = &v
}

func (o BetaGroupUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicLinkEnabled) {
		toSerialize["publicLinkEnabled"] = o.PublicLinkEnabled
	}
	if !IsNil(o.PublicLinkLimitEnabled) {
		toSerialize["publicLinkLimitEnabled"] = o.PublicLinkLimitEnabled
	}
	if !IsNil(o.PublicLinkLimit) {
		toSerialize["publicLinkLimit"] = o.PublicLinkLimit
	}
	if !IsNil(o.FeedbackEnabled) {
		toSerialize["feedbackEnabled"] = o.FeedbackEnabled
	}
	if !IsNil(o.IosBuildsAvailableForAppleSiliconMac) {
		toSerialize["iosBuildsAvailableForAppleSiliconMac"] = o.IosBuildsAvailableForAppleSiliconMac
	}
	return toSerialize, nil
}

type NullableBetaGroupUpdateRequestDataAttributes struct {
	value *BetaGroupUpdateRequestDataAttributes
	isSet bool
}

func (v NullableBetaGroupUpdateRequestDataAttributes) Get() *BetaGroupUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) Set(val *BetaGroupUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupUpdateRequestDataAttributes(val *BetaGroupUpdateRequestDataAttributes) *NullableBetaGroupUpdateRequestDataAttributes {
	return &NullableBetaGroupUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaGroupUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

