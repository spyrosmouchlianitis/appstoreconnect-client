# AnalyticsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AnalyticsReportAttributes**](AnalyticsReportAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**AnalyticsReportRelationships**](AnalyticsReportRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAnalyticsReport

`func NewAnalyticsReport(type_ string, id string, ) *AnalyticsReport`

NewAnalyticsReport instantiates a new AnalyticsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsReportWithDefaults

`func NewAnalyticsReportWithDefaults() *AnalyticsReport`

NewAnalyticsReportWithDefaults instantiates a new AnalyticsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnalyticsReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyticsReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyticsReport) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AnalyticsReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsReport) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AnalyticsReport) GetAttributes() AnalyticsReportAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AnalyticsReport) GetAttributesOk() (*AnalyticsReportAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AnalyticsReport) SetAttributes(v AnalyticsReportAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AnalyticsReport) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AnalyticsReport) GetRelationships() AnalyticsReportRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AnalyticsReport) GetRelationshipsOk() (*AnalyticsReportRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AnalyticsReport) SetRelationships(v AnalyticsReportRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AnalyticsReport) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AnalyticsReport) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnalyticsReport) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnalyticsReport) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AnalyticsReport) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


