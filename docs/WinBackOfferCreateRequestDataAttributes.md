# WinBackOfferCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceName** | **string** |  | 
**OfferId** | **string** |  | 
**Duration** | [**SubscriptionOfferDuration**](SubscriptionOfferDuration.md) |  | 
**OfferMode** | [**SubscriptionOfferMode**](SubscriptionOfferMode.md) |  | 
**PeriodCount** | **int32** |  | 
**CustomerEligibilityPaidSubscriptionDurationInMonths** | **int32** |  | 
**CustomerEligibilityTimeSinceLastSubscribedInMonths** | [**IntegerRange**](IntegerRange.md) |  | 
**CustomerEligibilityWaitBetweenOffersInMonths** | Pointer to **int32** |  | [optional] 
**StartDate** | **string** |  | 
**EndDate** | Pointer to **string** |  | [optional] 
**Priority** | **string** |  | 
**PromotionIntent** | Pointer to **string** |  | [optional] 

## Methods

### NewWinBackOfferCreateRequestDataAttributes

`func NewWinBackOfferCreateRequestDataAttributes(referenceName string, offerId string, duration SubscriptionOfferDuration, offerMode SubscriptionOfferMode, periodCount int32, customerEligibilityPaidSubscriptionDurationInMonths int32, customerEligibilityTimeSinceLastSubscribedInMonths IntegerRange, startDate string, priority string, ) *WinBackOfferCreateRequestDataAttributes`

NewWinBackOfferCreateRequestDataAttributes instantiates a new WinBackOfferCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferCreateRequestDataAttributesWithDefaults

`func NewWinBackOfferCreateRequestDataAttributesWithDefaults() *WinBackOfferCreateRequestDataAttributes`

NewWinBackOfferCreateRequestDataAttributesWithDefaults instantiates a new WinBackOfferCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceName

`func (o *WinBackOfferCreateRequestDataAttributes) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *WinBackOfferCreateRequestDataAttributes) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.


### GetOfferId

`func (o *WinBackOfferCreateRequestDataAttributes) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *WinBackOfferCreateRequestDataAttributes) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.


### GetDuration

`func (o *WinBackOfferCreateRequestDataAttributes) GetDuration() SubscriptionOfferDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WinBackOfferCreateRequestDataAttributes) SetDuration(v SubscriptionOfferDuration)`

SetDuration sets Duration field to given value.


### GetOfferMode

`func (o *WinBackOfferCreateRequestDataAttributes) GetOfferMode() SubscriptionOfferMode`

GetOfferMode returns the OfferMode field if non-nil, zero value otherwise.

### GetOfferModeOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool)`

GetOfferModeOk returns a tuple with the OfferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferMode

`func (o *WinBackOfferCreateRequestDataAttributes) SetOfferMode(v SubscriptionOfferMode)`

SetOfferMode sets OfferMode field to given value.


### GetPeriodCount

`func (o *WinBackOfferCreateRequestDataAttributes) GetPeriodCount() int32`

GetPeriodCount returns the PeriodCount field if non-nil, zero value otherwise.

### GetPeriodCountOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetPeriodCountOk() (*int32, bool)`

GetPeriodCountOk returns a tuple with the PeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodCount

`func (o *WinBackOfferCreateRequestDataAttributes) SetPeriodCount(v int32)`

SetPeriodCount sets PeriodCount field to given value.


### GetCustomerEligibilityPaidSubscriptionDurationInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonths() int32`

GetCustomerEligibilityPaidSubscriptionDurationInMonths returns the CustomerEligibilityPaidSubscriptionDurationInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk() (*int32, bool)`

GetCustomerEligibilityPaidSubscriptionDurationInMonthsOk returns a tuple with the CustomerEligibilityPaidSubscriptionDurationInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityPaidSubscriptionDurationInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityPaidSubscriptionDurationInMonths(v int32)`

SetCustomerEligibilityPaidSubscriptionDurationInMonths sets CustomerEligibilityPaidSubscriptionDurationInMonths field to given value.


### GetCustomerEligibilityTimeSinceLastSubscribedInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonths() IntegerRange`

GetCustomerEligibilityTimeSinceLastSubscribedInMonths returns the CustomerEligibilityTimeSinceLastSubscribedInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk() (*IntegerRange, bool)`

GetCustomerEligibilityTimeSinceLastSubscribedInMonthsOk returns a tuple with the CustomerEligibilityTimeSinceLastSubscribedInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityTimeSinceLastSubscribedInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityTimeSinceLastSubscribedInMonths(v IntegerRange)`

SetCustomerEligibilityTimeSinceLastSubscribedInMonths sets CustomerEligibilityTimeSinceLastSubscribedInMonths field to given value.


### GetCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityWaitBetweenOffersInMonths() int32`

GetCustomerEligibilityWaitBetweenOffersInMonths returns the CustomerEligibilityWaitBetweenOffersInMonths field if non-nil, zero value otherwise.

### GetCustomerEligibilityWaitBetweenOffersInMonthsOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetCustomerEligibilityWaitBetweenOffersInMonthsOk() (*int32, bool)`

GetCustomerEligibilityWaitBetweenOffersInMonthsOk returns a tuple with the CustomerEligibilityWaitBetweenOffersInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) SetCustomerEligibilityWaitBetweenOffersInMonths(v int32)`

SetCustomerEligibilityWaitBetweenOffersInMonths sets CustomerEligibilityWaitBetweenOffersInMonths field to given value.

### HasCustomerEligibilityWaitBetweenOffersInMonths

`func (o *WinBackOfferCreateRequestDataAttributes) HasCustomerEligibilityWaitBetweenOffersInMonths() bool`

HasCustomerEligibilityWaitBetweenOffersInMonths returns a boolean if a field has been set.

### GetStartDate

`func (o *WinBackOfferCreateRequestDataAttributes) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WinBackOfferCreateRequestDataAttributes) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *WinBackOfferCreateRequestDataAttributes) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WinBackOfferCreateRequestDataAttributes) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *WinBackOfferCreateRequestDataAttributes) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriority

`func (o *WinBackOfferCreateRequestDataAttributes) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WinBackOfferCreateRequestDataAttributes) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetPromotionIntent

`func (o *WinBackOfferCreateRequestDataAttributes) GetPromotionIntent() string`

GetPromotionIntent returns the PromotionIntent field if non-nil, zero value otherwise.

### GetPromotionIntentOk

`func (o *WinBackOfferCreateRequestDataAttributes) GetPromotionIntentOk() (*string, bool)`

GetPromotionIntentOk returns a tuple with the PromotionIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionIntent

`func (o *WinBackOfferCreateRequestDataAttributes) SetPromotionIntent(v string)`

SetPromotionIntent sets PromotionIntent field to given value.

### HasPromotionIntent

`func (o *WinBackOfferCreateRequestDataAttributes) HasPromotionIntent() bool`

HasPromotionIntent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


