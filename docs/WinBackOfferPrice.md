# WinBackOfferPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**SubscriptionOfferCodePriceRelationships**](SubscriptionOfferCodePriceRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewWinBackOfferPrice

`func NewWinBackOfferPrice(type_ string, id string, ) *WinBackOfferPrice`

NewWinBackOfferPrice instantiates a new WinBackOfferPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferPriceWithDefaults

`func NewWinBackOfferPriceWithDefaults() *WinBackOfferPrice`

NewWinBackOfferPriceWithDefaults instantiates a new WinBackOfferPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WinBackOfferPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WinBackOfferPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WinBackOfferPrice) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WinBackOfferPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WinBackOfferPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WinBackOfferPrice) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *WinBackOfferPrice) GetRelationships() SubscriptionOfferCodePriceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WinBackOfferPrice) GetRelationshipsOk() (*SubscriptionOfferCodePriceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WinBackOfferPrice) SetRelationships(v SubscriptionOfferCodePriceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WinBackOfferPrice) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *WinBackOfferPrice) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WinBackOfferPrice) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WinBackOfferPrice) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WinBackOfferPrice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


