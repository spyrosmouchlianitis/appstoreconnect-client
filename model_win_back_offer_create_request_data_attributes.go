/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WinBackOfferCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WinBackOfferCreateRequestDataAttributes{}

// WinBackOfferCreateRequestDataAttributes struct for WinBackOfferCreateRequestDataAttributes
type WinBackOfferCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	OfferId string `json:"offerId"`
	Duration SubscriptionOfferDuration `json:"duration"`
	OfferMode SubscriptionOfferMode `json:"offerMode"`
	PeriodCount int32 `json:"periodCount"`
	CustomerEligibilityPaidSubscriptionDurationInMonths int32 `json:"customerEligibilityPaidSubscriptionDurationInMonths"`
	CustomerEligibilityTimeSinceLastSubscribedInMonths IntegerRange `json:"customerEligibilityTimeSinceLastSubscribedInMonths"`
	CustomerEligibilityWaitBetweenOffersInMonths *int32 `json:"customerEligibilityWaitBetweenOffersInMonths,omitempty"`
	StartDate string `json:"startDate"`
	EndDate *string `json:"endDate,omitempty"`
	Priority string `json:"priority"`
	PromotionIntent *string `json:"promotionIntent,omitempty"`
}

type _WinBackOfferCreateRequestDataAttributes WinBackOfferCreateRequestDataAttributes

// NewWinBackOfferCreateRequestDataAttributes instantiates a new WinBackOfferCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWinBackOfferCreateRequestDataAttributes(referenceName string, offerId string, duration SubscriptionOfferDuration, offerMode SubscriptionOfferMode, periodCount int32, customerEligibilityPaidSubscriptionDurationInMonths int32, customerEligibilityTimeSinceLastSubscribedInMonths IntegerRange, startDate string, priority string) *WinBackOfferCreateRequestDataAttributes {
	this := WinBackOfferCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	this.OfferId = offerId
	this.Duration = duration
	this.OfferMode = offerMode
	this.PeriodCount = periodCount
	this.CustomerEligibilityPaidSubscriptionDurationInMonths = customerEligibilityPaidSubscriptionDurationInMonths
	this.CustomerEligibilityTimeSinceLastSubscribedInMonths = customerEligibilityTimeSinceLastSubscribedInMonths
	this.StartDate = startDate
	this.Priority = priority
	return &this
}

// NewWinBackOfferCreateRequestDataAttributesWithDefaults instantiates a new WinBackOfferCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWinBackOfferCreateRequestDataAttributesWithDefaults() *WinBackOfferCreateRequestDataAttributes {
	this := WinBackOfferCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *WinBackOfferCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetOfferId returns the OfferId field value
func (o *WinBackOfferCreateRequestDataAttributes) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetOfferId(v string) {
	o.OfferId = v
}

// GetDuration returns the Duration field value
func (o *WinBackOfferCreateRequestDataAttributes) GetDuration() SubscriptionOfferDuration {
	if o == nil {
		var ret SubscriptionOfferDuration
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetDuration(v SubscriptionOfferDuration) {
	o.Duration = v
}

// GetOfferMode returns the OfferMode field value
func (o *WinBackOfferCreateRequestDataAttributes) GetOfferMode() SubscriptionOfferMode {
	if o == nil {
		var ret SubscriptionOfferMode
		return ret
	}

	return o.OfferMode
}

// GetOfferModeOk returns a tuple with the OfferMode field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferMode, true
}

// SetOfferMode sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetOfferMode(v SubscriptionOfferMode) {
	o.OfferMode = v
}

// GetPeriodCount returns the PeriodCount field value
func (o *WinBackOfferCreateRequestDataAttributes) GetPeriodCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeriodCount
}

// GetPeriodCountOk returns a tuple with the PeriodCount field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetPeriodCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodCount, true
}

// SetPeriodCount sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetPeriodCount(v int32) {
	o.PeriodCount = v
}

// GetCustomerEligibilityPaidSubscriptionDurationInMonths returns the CustomerEligibilityPaidSubscriptionDurationInMonths field value
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonths() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerEligibilityPaidSubscriptionDurationInMonths
}

// GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk returns a tuple with the CustomerEligibilityPaidSubscriptionDurationInMonths field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerEligibilityPaidSubscriptionDurationInMonths, true
}

// SetCustomerEligibilityPaidSubscriptionDurationInMonths sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityPaidSubscriptionDurationInMonths(v int32) {
	o.CustomerEligibilityPaidSubscriptionDurationInMonths = v
}

// GetCustomerEligibilityTimeSinceLastSubscribedInMonths returns the CustomerEligibilityTimeSinceLastSubscribedInMonths field value
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonths() IntegerRange {
	if o == nil {
		var ret IntegerRange
		return ret
	}

	return o.CustomerEligibilityTimeSinceLastSubscribedInMonths
}

// GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk returns a tuple with the CustomerEligibilityTimeSinceLastSubscribedInMonths field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk() (*IntegerRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerEligibilityTimeSinceLastSubscribedInMonths, true
}

// SetCustomerEligibilityTimeSinceLastSubscribedInMonths sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityTimeSinceLastSubscribedInMonths(v IntegerRange) {
	o.CustomerEligibilityTimeSinceLastSubscribedInMonths = v
}

// GetCustomerEligibilityWaitBetweenOffersInMonths returns the CustomerEligibilityWaitBetweenOffersInMonths field value if set, zero value otherwise.
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityWaitBetweenOffersInMonths() int32 {
	if o == nil || IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		var ret int32
		return ret
	}
	return *o.CustomerEligibilityWaitBetweenOffersInMonths
}

// GetCustomerEligibilityWaitBetweenOffersInMonthsOk returns a tuple with the CustomerEligibilityWaitBetweenOffersInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityWaitBetweenOffersInMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		return nil, false
	}
	return o.CustomerEligibilityWaitBetweenOffersInMonths, true
}

// HasCustomerEligibilityWaitBetweenOffersInMonths returns a boolean if a field has been set.
func (o *WinBackOfferCreateRequestDataAttributes) HasCustomerEligibilityWaitBetweenOffersInMonths() bool {
	if o != nil && !IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		return true
	}

	return false
}

// SetCustomerEligibilityWaitBetweenOffersInMonths gets a reference to the given int32 and assigns it to the CustomerEligibilityWaitBetweenOffersInMonths field.
func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityWaitBetweenOffersInMonths(v int32) {
	o.CustomerEligibilityWaitBetweenOffersInMonths = &v
}

// GetStartDate returns the StartDate field value
func (o *WinBackOfferCreateRequestDataAttributes) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *WinBackOfferCreateRequestDataAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *WinBackOfferCreateRequestDataAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *WinBackOfferCreateRequestDataAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPriority returns the Priority field value
func (o *WinBackOfferCreateRequestDataAttributes) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *WinBackOfferCreateRequestDataAttributes) SetPriority(v string) {
	o.Priority = v
}

// GetPromotionIntent returns the PromotionIntent field value if set, zero value otherwise.
func (o *WinBackOfferCreateRequestDataAttributes) GetPromotionIntent() string {
	if o == nil || IsNil(o.PromotionIntent) {
		var ret string
		return ret
	}
	return *o.PromotionIntent
}

// GetPromotionIntentOk returns a tuple with the PromotionIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferCreateRequestDataAttributes) GetPromotionIntentOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionIntent) {
		return nil, false
	}
	return o.PromotionIntent, true
}

// HasPromotionIntent returns a boolean if a field has been set.
func (o *WinBackOfferCreateRequestDataAttributes) HasPromotionIntent() bool {
	if o != nil && !IsNil(o.PromotionIntent) {
		return true
	}

	return false
}

// SetPromotionIntent gets a reference to the given string and assigns it to the PromotionIntent field.
func (o *WinBackOfferCreateRequestDataAttributes) SetPromotionIntent(v string) {
	o.PromotionIntent = &v
}

func (o WinBackOfferCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WinBackOfferCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["offerId"] = o.OfferId
	toSerialize["duration"] = o.Duration
	toSerialize["offerMode"] = o.OfferMode
	toSerialize["periodCount"] = o.PeriodCount
	toSerialize["customerEligibilityPaidSubscriptionDurationInMonths"] = o.CustomerEligibilityPaidSubscriptionDurationInMonths
	toSerialize["customerEligibilityTimeSinceLastSubscribedInMonths"] = o.CustomerEligibilityTimeSinceLastSubscribedInMonths
	if !IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		toSerialize["customerEligibilityWaitBetweenOffersInMonths"] = o.CustomerEligibilityWaitBetweenOffersInMonths
	}
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.PromotionIntent) {
		toSerialize["promotionIntent"] = o.PromotionIntent
	}
	return toSerialize, nil
}

func (o *WinBackOfferCreateRequestDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenceName",
		"offerId",
		"duration",
		"offerMode",
		"periodCount",
		"customerEligibilityPaidSubscriptionDurationInMonths",
		"customerEligibilityTimeSinceLastSubscribedInMonths",
		"startDate",
		"priority",
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

	varWinBackOfferCreateRequestDataAttributes := _WinBackOfferCreateRequestDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWinBackOfferCreateRequestDataAttributes)

	if err != nil {
		return err
	}

	*o = WinBackOfferCreateRequestDataAttributes(varWinBackOfferCreateRequestDataAttributes)

	return err
}

type NullableWinBackOfferCreateRequestDataAttributes struct {
	value *WinBackOfferCreateRequestDataAttributes
	isSet bool
}

func (v NullableWinBackOfferCreateRequestDataAttributes) Get() *WinBackOfferCreateRequestDataAttributes {
	return v.value
}

func (v *NullableWinBackOfferCreateRequestDataAttributes) Set(val *WinBackOfferCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableWinBackOfferCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableWinBackOfferCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWinBackOfferCreateRequestDataAttributes(val *WinBackOfferCreateRequestDataAttributes) *NullableWinBackOfferCreateRequestDataAttributes {
	return &NullableWinBackOfferCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableWinBackOfferCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWinBackOfferCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


