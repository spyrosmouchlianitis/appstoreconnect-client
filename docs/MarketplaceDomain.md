# MarketplaceDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AlternativeDistributionDomainAttributes**](AlternativeDistributionDomainAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewMarketplaceDomain

`func NewMarketplaceDomain(type_ string, id string, ) *MarketplaceDomain`

NewMarketplaceDomain instantiates a new MarketplaceDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceDomainWithDefaults

`func NewMarketplaceDomainWithDefaults() *MarketplaceDomain`

NewMarketplaceDomainWithDefaults instantiates a new MarketplaceDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketplaceDomain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceDomain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceDomain) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MarketplaceDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceDomain) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MarketplaceDomain) GetAttributes() AlternativeDistributionDomainAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketplaceDomain) GetAttributesOk() (*AlternativeDistributionDomainAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketplaceDomain) SetAttributes(v AlternativeDistributionDomainAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MarketplaceDomain) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *MarketplaceDomain) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketplaceDomain) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketplaceDomain) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MarketplaceDomain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


