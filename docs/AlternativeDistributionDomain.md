# AlternativeDistributionDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AlternativeDistributionDomainAttributes**](AlternativeDistributionDomainAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAlternativeDistributionDomain

`func NewAlternativeDistributionDomain(type_ string, id string, ) *AlternativeDistributionDomain`

NewAlternativeDistributionDomain instantiates a new AlternativeDistributionDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionDomainWithDefaults

`func NewAlternativeDistributionDomainWithDefaults() *AlternativeDistributionDomain`

NewAlternativeDistributionDomainWithDefaults instantiates a new AlternativeDistributionDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlternativeDistributionDomain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlternativeDistributionDomain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlternativeDistributionDomain) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AlternativeDistributionDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlternativeDistributionDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlternativeDistributionDomain) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AlternativeDistributionDomain) GetAttributes() AlternativeDistributionDomainAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AlternativeDistributionDomain) GetAttributesOk() (*AlternativeDistributionDomainAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AlternativeDistributionDomain) SetAttributes(v AlternativeDistributionDomainAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AlternativeDistributionDomain) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionDomain) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionDomain) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionDomain) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlternativeDistributionDomain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


