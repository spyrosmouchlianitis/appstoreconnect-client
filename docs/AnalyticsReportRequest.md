# AnalyticsReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AnalyticsReportRequestAttributes**](AnalyticsReportRequestAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**AnalyticsReportRequestRelationships**](AnalyticsReportRequestRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAnalyticsReportRequest

`func NewAnalyticsReportRequest(type_ string, id string, ) *AnalyticsReportRequest`

NewAnalyticsReportRequest instantiates a new AnalyticsReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsReportRequestWithDefaults

`func NewAnalyticsReportRequestWithDefaults() *AnalyticsReportRequest`

NewAnalyticsReportRequestWithDefaults instantiates a new AnalyticsReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnalyticsReportRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyticsReportRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyticsReportRequest) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AnalyticsReportRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsReportRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsReportRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AnalyticsReportRequest) GetAttributes() AnalyticsReportRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AnalyticsReportRequest) GetAttributesOk() (*AnalyticsReportRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AnalyticsReportRequest) SetAttributes(v AnalyticsReportRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AnalyticsReportRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AnalyticsReportRequest) GetRelationships() AnalyticsReportRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AnalyticsReportRequest) GetRelationshipsOk() (*AnalyticsReportRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AnalyticsReportRequest) SetRelationships(v AnalyticsReportRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AnalyticsReportRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AnalyticsReportRequest) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyticsReportRequest) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyticsReportRequest) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AnalyticsReportRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


