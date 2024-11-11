# SubscriptionImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SubscriptionImage**](SubscriptionImage.md) |  | 
**Included** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewSubscriptionImageResponse

`func NewSubscriptionImageResponse(data SubscriptionImage, links DocumentLinks, ) *SubscriptionImageResponse`

NewSubscriptionImageResponse instantiates a new SubscriptionImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionImageResponseWithDefaults

`func NewSubscriptionImageResponseWithDefaults() *SubscriptionImageResponse`

NewSubscriptionImageResponseWithDefaults instantiates a new SubscriptionImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionImageResponse) GetData() SubscriptionImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionImageResponse) GetDataOk() (*SubscriptionImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionImageResponse) SetData(v SubscriptionImage)`

SetData sets Data field to given value.


### GetIncluded

`func (o *SubscriptionImageResponse) GetIncluded() []Subscription`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SubscriptionImageResponse) GetIncludedOk() (*[]Subscription, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SubscriptionImageResponse) SetIncluded(v []Subscription)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *SubscriptionImageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *SubscriptionImageResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SubscriptionImageResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SubscriptionImageResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


