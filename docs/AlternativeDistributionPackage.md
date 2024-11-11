# AlternativeDistributionPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Relationships** | Pointer to [**AlternativeDistributionPackageRelationships**](AlternativeDistributionPackageRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAlternativeDistributionPackage

`func NewAlternativeDistributionPackage(type_ string, id string, ) *AlternativeDistributionPackage`

NewAlternativeDistributionPackage instantiates a new AlternativeDistributionPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeDistributionPackageWithDefaults

`func NewAlternativeDistributionPackageWithDefaults() *AlternativeDistributionPackage`

NewAlternativeDistributionPackageWithDefaults instantiates a new AlternativeDistributionPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlternativeDistributionPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlternativeDistributionPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlternativeDistributionPackage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AlternativeDistributionPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlternativeDistributionPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlternativeDistributionPackage) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *AlternativeDistributionPackage) GetRelationships() AlternativeDistributionPackageRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AlternativeDistributionPackage) GetRelationshipsOk() (*AlternativeDistributionPackageRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AlternativeDistributionPackage) SetRelationships(v AlternativeDistributionPackageRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AlternativeDistributionPackage) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AlternativeDistributionPackage) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlternativeDistributionPackage) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlternativeDistributionPackage) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlternativeDistributionPackage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


