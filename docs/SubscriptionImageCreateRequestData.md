# SubscriptionImageCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppClipAdvancedExperienceImageCreateRequestDataAttributes**](AppClipAdvancedExperienceImageCreateRequestDataAttributes.md) |  | 
**Relationships** | [**SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships**](SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships.md) |  | 

## Methods

### NewSubscriptionImageCreateRequestData

`func NewSubscriptionImageCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships, ) *SubscriptionImageCreateRequestData`

NewSubscriptionImageCreateRequestData instantiates a new SubscriptionImageCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionImageCreateRequestDataWithDefaults

`func NewSubscriptionImageCreateRequestDataWithDefaults() *SubscriptionImageCreateRequestData`

NewSubscriptionImageCreateRequestDataWithDefaults instantiates a new SubscriptionImageCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionImageCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionImageCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionImageCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionImageCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionImageCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionImageCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionImageCreateRequestData) GetRelationships() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionImageCreateRequestData) GetRelationshipsOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionImageCreateRequestData) SetRelationships(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


