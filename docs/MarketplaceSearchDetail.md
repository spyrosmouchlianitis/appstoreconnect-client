# MarketplaceSearchDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**MarketplaceSearchDetailAttributes**](MarketplaceSearchDetailAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewMarketplaceSearchDetail

`func NewMarketplaceSearchDetail(type_ string, id string, ) *MarketplaceSearchDetail`

NewMarketplaceSearchDetail instantiates a new MarketplaceSearchDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceSearchDetailWithDefaults

`func NewMarketplaceSearchDetailWithDefaults() *MarketplaceSearchDetail`

NewMarketplaceSearchDetailWithDefaults instantiates a new MarketplaceSearchDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketplaceSearchDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceSearchDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceSearchDetail) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MarketplaceSearchDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceSearchDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceSearchDetail) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MarketplaceSearchDetail) GetAttributes() MarketplaceSearchDetailAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketplaceSearchDetail) GetAttributesOk() (*MarketplaceSearchDetailAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketplaceSearchDetail) SetAttributes(v MarketplaceSearchDetailAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MarketplaceSearchDetail) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *MarketplaceSearchDetail) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketplaceSearchDetail) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketplaceSearchDetail) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MarketplaceSearchDetail) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


