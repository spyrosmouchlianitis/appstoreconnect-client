# SubscriptionOfferCodePricesResponseIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**SubscriptionPricePointAttributes**](SubscriptionPricePointAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 
**Relationships** | Pointer to [**SubscriptionPricePointRelationships**](SubscriptionPricePointRelationships.md) |  | [optional] 

## Methods

### NewSubscriptionOfferCodePricesResponseIncludedInner

`func NewSubscriptionOfferCodePricesResponseIncludedInner(type_ string, id string, ) *SubscriptionOfferCodePricesResponseIncludedInner`

NewSubscriptionOfferCodePricesResponseIncludedInner instantiates a new SubscriptionOfferCodePricesResponseIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionOfferCodePricesResponseIncludedInnerWithDefaults

`func NewSubscriptionOfferCodePricesResponseIncludedInnerWithDefaults() *SubscriptionOfferCodePricesResponseIncludedInner`

NewSubscriptionOfferCodePricesResponseIncludedInnerWithDefaults instantiates a new SubscriptionOfferCodePricesResponseIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetAttributes() SubscriptionPricePointAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetAttributesOk() (*SubscriptionPricePointAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) SetAttributes(v SubscriptionPricePointAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetRelationships() SubscriptionPricePointRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) GetRelationshipsOk() (*SubscriptionPricePointRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) SetRelationships(v SubscriptionPricePointRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SubscriptionOfferCodePricesResponseIncludedInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


