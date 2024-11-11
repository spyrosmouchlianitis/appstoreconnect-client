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

// checks if the WinBackOfferAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WinBackOfferAttributes{}

// WinBackOfferAttributes struct for WinBackOfferAttributes
type WinBackOfferAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
	OfferId *string `json:"offerId,omitempty"`
	Duration *SubscriptionOfferDuration `json:"duration,omitempty"`
	OfferMode *SubscriptionOfferMode `json:"offerMode,omitempty"`
	PeriodCount *int32 `json:"periodCount,omitempty"`
	CustomerEligibilityPaidSubscriptionDurationInMonths *int32 `json:"customerEligibilityPaidSubscriptionDurationInMonths,omitempty"`
	CustomerEligibilityTimeSinceLastSubscribedInMonths *IntegerRange `json:"customerEligibilityTimeSinceLastSubscribedInMonths,omitempty"`
	CustomerEligibilityWaitBetweenOffersInMonths *int32 `json:"customerEligibilityWaitBetweenOffersInMonths,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	Priority *string `json:"priority,omitempty"`
	PromotionIntent *string `json:"promotionIntent,omitempty"`
}

// NewWinBackOfferAttributes instantiates a new WinBackOfferAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWinBackOfferAttributes() *WinBackOfferAttributes {
	this := WinBackOfferAttributes{}
	return &this
}

// NewWinBackOfferAttributesWithDefaults instantiates a new WinBackOfferAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWinBackOfferAttributesWithDefaults() *WinBackOfferAttributes {
	this := WinBackOfferAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *WinBackOfferAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetOfferId() string {
	if o == nil || IsNil(o.OfferId) {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.OfferId) {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasOfferId() bool {
	if o != nil && !IsNil(o.OfferId) {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *WinBackOfferAttributes) SetOfferId(v string) {
	o.OfferId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetDuration() SubscriptionOfferDuration {
	if o == nil || IsNil(o.Duration) {
		var ret SubscriptionOfferDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given SubscriptionOfferDuration and assigns it to the Duration field.
func (o *WinBackOfferAttributes) SetDuration(v SubscriptionOfferDuration) {
	o.Duration = &v
}

// GetOfferMode returns the OfferMode field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetOfferMode() SubscriptionOfferMode {
	if o == nil || IsNil(o.OfferMode) {
		var ret SubscriptionOfferMode
		return ret
	}
	return *o.OfferMode
}

// GetOfferModeOk returns a tuple with the OfferMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool) {
	if o == nil || IsNil(o.OfferMode) {
		return nil, false
	}
	return o.OfferMode, true
}

// HasOfferMode returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasOfferMode() bool {
	if o != nil && !IsNil(o.OfferMode) {
		return true
	}

	return false
}

// SetOfferMode gets a reference to the given SubscriptionOfferMode and assigns it to the OfferMode field.
func (o *WinBackOfferAttributes) SetOfferMode(v SubscriptionOfferMode) {
	o.OfferMode = &v
}

// GetPeriodCount returns the PeriodCount field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetPeriodCount() int32 {
	if o == nil || IsNil(o.PeriodCount) {
		var ret int32
		return ret
	}
	return *o.PeriodCount
}

// GetPeriodCountOk returns a tuple with the PeriodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetPeriodCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodCount) {
		return nil, false
	}
	return o.PeriodCount, true
}

// HasPeriodCount returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasPeriodCount() bool {
	if o != nil && !IsNil(o.PeriodCount) {
		return true
	}

	return false
}

// SetPeriodCount gets a reference to the given int32 and assigns it to the PeriodCount field.
func (o *WinBackOfferAttributes) SetPeriodCount(v int32) {
	o.PeriodCount = &v
}

// GetCustomerEligibilityPaidSubscriptionDurationInMonths returns the CustomerEligibilityPaidSubscriptionDurationInMonths field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonths() int32 {
	if o == nil || IsNil(o.CustomerEligibilityPaidSubscriptionDurationInMonths) {
		var ret int32
		return ret
	}
	return *o.CustomerEligibilityPaidSubscriptionDurationInMonths
}

// GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk returns a tuple with the CustomerEligibilityPaidSubscriptionDurationInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerEligibilityPaidSubscriptionDurationInMonths) {
		return nil, false
	}
	return o.CustomerEligibilityPaidSubscriptionDurationInMonths, true
}

// HasCustomerEligibilityPaidSubscriptionDurationInMonths returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasCustomerEligibilityPaidSubscriptionDurationInMonths() bool {
	if o != nil && !IsNil(o.CustomerEligibilityPaidSubscriptionDurationInMonths) {
		return true
	}

	return false
}

// SetCustomerEligibilityPaidSubscriptionDurationInMonths gets a reference to the given int32 and assigns it to the CustomerEligibilityPaidSubscriptionDurationInMonths field.
func (o *WinBackOfferAttributes) SetCustomerEligibilityPaidSubscriptionDurationInMonths(v int32) {
	o.CustomerEligibilityPaidSubscriptionDurationInMonths = &v
}

// GetCustomerEligibilityTimeSinceLastSubscribedInMonths returns the CustomerEligibilityTimeSinceLastSubscribedInMonths field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonths() IntegerRange {
	if o == nil || IsNil(o.CustomerEligibilityTimeSinceLastSubscribedInMonths) {
		var ret IntegerRange
		return ret
	}
	return *o.CustomerEligibilityTimeSinceLastSubscribedInMonths
}

// GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk returns a tuple with the CustomerEligibilityTimeSinceLastSubscribedInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk() (*IntegerRange, bool) {
	if o == nil || IsNil(o.CustomerEligibilityTimeSinceLastSubscribedInMonths) {
		return nil, false
	}
	return o.CustomerEligibilityTimeSinceLastSubscribedInMonths, true
}

// HasCustomerEligibilityTimeSinceLastSubscribedInMonths returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasCustomerEligibilityTimeSinceLastSubscribedInMonths() bool {
	if o != nil && !IsNil(o.CustomerEligibilityTimeSinceLastSubscribedInMonths) {
		return true
	}

	return false
}

// SetCustomerEligibilityTimeSinceLastSubscribedInMonths gets a reference to the given IntegerRange and assigns it to the CustomerEligibilityTimeSinceLastSubscribedInMonths field.
func (o *WinBackOfferAttributes) SetCustomerEligibilityTimeSinceLastSubscribedInMonths(v IntegerRange) {
	o.CustomerEligibilityTimeSinceLastSubscribedInMonths = &v
}

// GetCustomerEligibilityWaitBetweenOffersInMonths returns the CustomerEligibilityWaitBetweenOffersInMonths field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetCustomerEligibilityWaitBetweenOffersInMonths() int32 {
	if o == nil || IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		var ret int32
		return ret
	}
	return *o.CustomerEligibilityWaitBetweenOffersInMonths
}

// GetCustomerEligibilityWaitBetweenOffersInMonthsOk returns a tuple with the CustomerEligibilityWaitBetweenOffersInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetCustomerEligibilityWaitBetweenOffersInMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		return nil, false
	}
	return o.CustomerEligibilityWaitBetweenOffersInMonths, true
}

// HasCustomerEligibilityWaitBetweenOffersInMonths returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasCustomerEligibilityWaitBetweenOffersInMonths() bool {
	if o != nil && !IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		return true
	}

	return false
}

// SetCustomerEligibilityWaitBetweenOffersInMonths gets a reference to the given int32 and assigns it to the CustomerEligibilityWaitBetweenOffersInMonths field.
func (o *WinBackOfferAttributes) SetCustomerEligibilityWaitBetweenOffersInMonths(v int32) {
	o.CustomerEligibilityWaitBetweenOffersInMonths = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *WinBackOfferAttributes) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *WinBackOfferAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *WinBackOfferAttributes) SetPriority(v string) {
	o.Priority = &v
}

// GetPromotionIntent returns the PromotionIntent field value if set, zero value otherwise.
func (o *WinBackOfferAttributes) GetPromotionIntent() string {
	if o == nil || IsNil(o.PromotionIntent) {
		var ret string
		return ret
	}
	return *o.PromotionIntent
}

// GetPromotionIntentOk returns a tuple with the PromotionIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WinBackOfferAttributes) GetPromotionIntentOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionIntent) {
		return nil, false
	}
	return o.PromotionIntent, true
}

// HasPromotionIntent returns a boolean if a field has been set.
func (o *WinBackOfferAttributes) HasPromotionIntent() bool {
	if o != nil && !IsNil(o.PromotionIntent) {
		return true
	}

	return false
}

// SetPromotionIntent gets a reference to the given string and assigns it to the PromotionIntent field.
func (o *WinBackOfferAttributes) SetPromotionIntent(v string) {
	o.PromotionIntent = &v
}

func (o WinBackOfferAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WinBackOfferAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.OfferId) {
		toSerialize["offerId"] = o.OfferId
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.OfferMode) {
		toSerialize["offerMode"] = o.OfferMode
	}
	if !IsNil(o.PeriodCount) {
		toSerialize["periodCount"] = o.PeriodCount
	}
	if !IsNil(o.CustomerEligibilityPaidSubscriptionDurationInMonths) {
		toSerialize["customerEligibilityPaidSubscriptionDurationInMonths"] = o.CustomerEligibilityPaidSubscriptionDurationInMonths
	}
	if !IsNil(o.CustomerEligibilityTimeSinceLastSubscribedInMonths) {
		toSerialize["customerEligibilityTimeSinceLastSubscribedInMonths"] = o.CustomerEligibilityTimeSinceLastSubscribedInMonths
	}
	if !IsNil(o.CustomerEligibilityWaitBetweenOffersInMonths) {
		toSerialize["customerEligibilityWaitBetweenOffersInMonths"] = o.CustomerEligibilityWaitBetweenOffersInMonths
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.PromotionIntent) {
		toSerialize["promotionIntent"] = o.PromotionIntent
	}
	return toSerialize, nil
}

type NullableWinBackOfferAttributes struct {
	value *WinBackOfferAttributes
	isSet bool
}

func (v NullableWinBackOfferAttributes) Get() *WinBackOfferAttributes {
	return v.value
}

func (v *NullableWinBackOfferAttributes) Set(val *WinBackOfferAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableWinBackOfferAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableWinBackOfferAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWinBackOfferAttributes(val *WinBackOfferAttributes) *NullableWinBackOfferAttributes {
	return &NullableWinBackOfferAttributes{value: val, isSet: true}
}

func (v NullableWinBackOfferAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWinBackOfferAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


