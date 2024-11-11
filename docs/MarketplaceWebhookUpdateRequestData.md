# MarketplaceWebhookUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**MarketplaceWebhookUpdateRequestDataAttributes**](MarketplaceWebhookUpdateRequestDataAttributes.md) |  | [optional] 

## Methods

### NewMarketplaceWebhookUpdateRequestData

`func NewMarketplaceWebhookUpdateRequestData(type_ string, id string, ) *MarketplaceWebhookUpdateRequestData`

NewMarketplaceWebhookUpdateRequestData instantiates a new MarketplaceWebhookUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceWebhookUpdateRequestDataWithDefaults

`func NewMarketplaceWebhookUpdateRequestDataWithDefaults() *MarketplaceWebhookUpdateRequestData`

NewMarketplaceWebhookUpdateRequestDataWithDefaults instantiates a new MarketplaceWebhookUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketplaceWebhookUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceWebhookUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceWebhookUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MarketplaceWebhookUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceWebhookUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceWebhookUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MarketplaceWebhookUpdateRequestData) GetAttributes() MarketplaceWebhookUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketplaceWebhookUpdateRequestData) GetAttributesOk() (*MarketplaceWebhookUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketplaceWebhookUpdateRequestData) SetAttributes(v MarketplaceWebhookUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MarketplaceWebhookUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


