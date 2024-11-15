# AlternativeDistributionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AlternativeDistributionKeyAttributes**](AlternativeDistributionKeyAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAlternativeDistributionKey

`func NewAlternativeDistributionKey(type_ string, id string, ) *AlternativeDistributionKey`

NewAlternativeDistributionKey instantiates a new AlternativeDistributionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionKeyWithDefaults

`func NewAlternativeDistributionKeyWithDefaults() *AlternativeDistributionKey`

NewAlternativeDistributionKeyWithDefaults instantiates a new AlternativeDistributionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlternativeDistributionKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlternativeDistributionKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlternativeDistributionKey) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AlternativeDistributionKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlternativeDistributionKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlternativeDistributionKey) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AlternativeDistributionKey) GetAttributes() AlternativeDistributionKeyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AlternativeDistributionKey) GetAttributesOk() (*AlternativeDistributionKeyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AlternativeDistributionKey) SetAttributes(v AlternativeDistributionKeyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AlternativeDistributionKey) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionKey) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionKey) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionKey) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlternativeDistributionKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


