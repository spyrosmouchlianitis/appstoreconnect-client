# MarketplaceWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**MarketplaceWebhookAttributes**](MarketplaceWebhookAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewMarketplaceWebhook

`func NewMarketplaceWebhook(type_ string, id string, ) *MarketplaceWebhook`

NewMarketplaceWebhook instantiates a new MarketplaceWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceWebhookWithDefaults

`func NewMarketplaceWebhookWithDefaults() *MarketplaceWebhook`

NewMarketplaceWebhookWithDefaults instantiates a new MarketplaceWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketplaceWebhook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceWebhook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceWebhook) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MarketplaceWebhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceWebhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceWebhook) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MarketplaceWebhook) GetAttributes() MarketplaceWebhookAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketplaceWebhook) GetAttributesOk() (*MarketplaceWebhookAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketplaceWebhook) SetAttributes(v MarketplaceWebhookAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MarketplaceWebhook) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *MarketplaceWebhook) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MarketplaceWebhook) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MarketplaceWebhook) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MarketplaceWebhook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


