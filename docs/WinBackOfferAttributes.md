# WinBackOfferAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceName** | Pointer to **string** |  | [optional] 
**OfferId** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to [**SubscriptionOfferDuration**](SubscriptionOfferDuration.md) |  | [optional] 
**OfferMode** | Pointer to [**SubscriptionOfferMode**](SubscriptionOfferMode.md) |  | [optional] 
**PeriodCount** | Pointer to **int32** |  | [optional] 
**CustomerEligibilityPaidSubscriptionDurationInMonths** | Pointer to **int32** |  | [optional] 
**CustomerEligibilityTimeSinceLastSubscribedInMonths** | Pointer to [**IntegerRange**](IntegerRange.md) |  | [optional] 
**CustomerEligibilityWaitBetweenOffersInMonths** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**PromotionIntent** | Pointer to **string** |  | [optional] 

## Methods

### NewWinBackOfferAttributes

`func NewWinBackOfferAttributes() *WinBackOfferAttributes`

NewWinBackOfferAttributes instantiates a new WinBackOfferAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferAttributesWithDefaults

`func NewWinBackOfferAttributesWithDefaults() *WinBackOfferAttributes`

NewWinBackOfferAttributesWithDefaults instantiates a new WinBackOfferAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceName

`func (o *WinBackOfferAttributes) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *WinBackOfferAttributes) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *WinBackOfferAttributes) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.

### HasReferenceName

`func (o *WinBackOfferAttributes) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### GetOfferId

`func (o *WinBackOfferAttributes) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *WinBackOfferAttributes) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *WinBackOfferAttributes) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *WinBackOfferAttributes) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetDuration

`func (o *WinBackOfferAttributes) GetDuration() SubscriptionOfferDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WinBackOfferAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WinBackOfferAttributes) SetDuration(v SubscriptionOfferDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WinBackOfferAttributes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetOfferMode

`func (o *WinBackOfferAttributes) GetOfferMode() SubscriptionOfferMode`

GetOfferMode returns the OfferMode field if non-nil, zero value otherwise.

### GetOfferModeOk

`func (o *WinBackOfferAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool)`

GetOfferModeOk returns a tuple with the OfferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferMode

`func (o *WinBackOfferAttributes) SetOfferMode(v SubscriptionOfferMode)`

SetOfferMode sets OfferMode field to given value.

### HasOfferMode

`func (o *WinBackOfferAttributes) HasOfferMode() bool`

HasOfferMode returns a boolean if a field has been set.

### GetPeriodCount

`func (o *WinBackOfferAttributes) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *WinBackOfferAttributes) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *WinBackOfferAttributes) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.

### HasPeriodCount

`func (o *WinBackOfferAttributes) HasPeriodCount() bool`

HasPeriodCount returns a boolean if a field has been set.

### GetCustomerEligibilityPaidSubscriptionDurationInMonths

`func (o *WinBackOfferAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonths() int32`

GetCustomerEligibilityPaidSubscriptionDurationInMonths returns the CustomerEligibilityPaidSubscriptionDurationInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk

`func (o *WinBackOfferAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk() (*int32, bool)`

GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk returns a tuple with the CustomerEligibilityPaidSubscriptionDurationInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityPaidSubscriptionDurationInMonths

`func (o *WinBackOfferAttributes) SetCustomerEligibilityPaidSubscriptionDurationInMonths(v int32)`

SetCustomerEligibilityPaidSubscriptionDurationInMonths sets CustomerEligibilityPaidSubscriptionDurationInMonths field to given value.

### HasCustomerEligibilityPaidSubscriptionDurationInMonths

`func (o *WinBackOfferAttributes) HasCustomerEligibilityPaidSubscriptionDurationInMonths() bool`

HasCustomerEligibilityPaidSubscriptionDurationInMonths returns a boolean if a field has been set.

### GetCustomerEligibilityTimeSinceLastSubscribedInMonths

`func (o *WinBackOfferAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonths() IntegerRange`

GetCustomerEligibilityTimeSinceLastSubscribedInMonths returns the CustomerEligibilityTimeSinceLastSubscribedInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk

`func (o *WinBackOfferAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk() (*IntegerRange, bool)`

GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk returns a tuple with the CustomerEligibilityTimeSinceLastSubscribedInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityTimeSinceLastSubscribedInMonths

`func (o *WinBackOfferAttributes) SetCustomerEligibilityTimeSinceLastSubscribedInMonths(v IntegerRange)`

SetCustomerEligibilityTimeSinceLastSubscribedInMonths sets CustomerEligibilityTimeSinceLastSubscribedInMonths field to given value.

### HasCustomerEligibilityTimeSinceLastSubscribedInMonths

`func (o *WinBackOfferAttributes) HasCustomerEligibilityTimeSinceLastSubscribedInMonths() bool`

HasCustomerEligibilityTimeSinceLastSubscribedInMonths returns a boolean if a field has been set.

### GetCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferAttributes) GetCustomerEligibilityWaitBetweenOffersInMonths() int32`

GetCustomerEligibilityWaitBetweenOffersInMonths returns the CustomerEligibilityWaitBetweenOffersInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityWaitBetweenOffersInMonthsOk

`func (o *WinBackOfferAttributes) GetCustomerEligibilityWaitBetweenOffersInMonthsOk() (*int32, bool)`

GetCustomerEligibilityWaitBetweenOffersInMonthsOk returns a tuple with the CustomerEligibilityWaitBetweenOffersInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferAttributes) SetCustomerEligibilityWaitBetweenOffersInMonths(v int32)`

SetCustomerEligibilityWaitBetweenOffersInMonths sets CustomerEligibilityWaitBetweenOffersInMonths field to given value.

### HasCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferAttributes) HasCustomerEligibilityWaitBetweenOffersInMonths() bool`

HasCustomerEligibilityWaitBetweenOffersInMonths returns a boolean if a field has been set.

### GetStartDate

`func (o *WinBackOfferAttributes) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WinBackOfferAttributes) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WinBackOfferAttributes) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *WinBackOfferAttributes) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *WinBackOfferAttributes) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WinBackOfferAttributes) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WinBackOfferAttributes) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *WinBackOfferAttributes) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriority

`func (o *WinBackOfferAttributes) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WinBackOfferAttributes) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WinBackOfferAttributes) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WinBackOfferAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPromotionIntent

`func (o *WinBackOfferAttributes) GetPromotionIntent() string`

GetPromotionIntent returns the PromotionIntent field if non-nil, zero value otherwise.

### GetPromotionIntentOk

`func (o *WinBackOfferAttributes) GetPromotionIntentOk() (*string, bool)`

GetPromotionIntentOk returns a tuple with the PromotionIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionIntent

`func (o *WinBackOfferAttributes) SetPromotionIntent(v string)`

SetPromotionIntent sets PromotionIntent field to given value.

### HasPromotionIntent

`func (o *WinBackOfferAttributes) HasPromotionIntent() bool`

HasPromotionIntent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


