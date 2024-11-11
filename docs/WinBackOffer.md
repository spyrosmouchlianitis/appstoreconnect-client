# WinBackOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**WinBackOfferAttributes**](WinBackOfferAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**WinBackOfferRelationships**](WinBackOfferRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewWinBackOffer

`func NewWinBackOffer(type_ string, id string, ) *WinBackOffer`

NewWinBackOffer instantiates a new WinBackOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWinBackOfferWithDefaults

`func NewWinBackOfferWithDefaults() *WinBackOffer`

NewWinBackOfferWithDefaults instantiates a new WinBackOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WinBackOffer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WinBackOffer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WinBackOffer) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WinBackOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WinBackOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WinBackOffer) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *WinBackOffer) GetAttributes() WinBackOfferAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WinBackOffer) GetAttributesOk() (*WinBackOfferAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WinBackOffer) SetAttributes(v WinBackOfferAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WinBackOffer) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WinBackOffer) GetRelationships() WinBackOfferRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WinBackOffer) GetRelationshipsOk() (*WinBackOfferRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WinBackOffer) SetRelationships(v WinBackOfferRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WinBackOffer) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *WinBackOffer) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WinBackOffer) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WinBackOffer) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WinBackOffer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


