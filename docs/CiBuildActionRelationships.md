# CiBuildActionRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildRun** | Pointer to [**CiBuildActionRelationshipsBuildRun**](CiBuildActionRelationshipsBuildRun.md) |  | [optional] 
**Artifacts** | Pointer to [**AnalyticsReportInstanceRelationshipsSegments**](AnalyticsReportInstanceRelationshipsSegments.md) |  | [optional] 
**Issues** | Pointer to [**AnalyticsReportInstanceRelationshipsSegments**](AnalyticsReportInstanceRelationshipsSegments.md) |  | [optional] 
**TestResults** | Pointer to [**AnalyticsReportInstanceRelationshipsSegments**](AnalyticsReportInstanceRelationshipsSegments.md) |  | [optional] 

## Methods

### NewCiBuildActionRelationships

`func NewCiBuildActionRelationships() *CiBuildActionRelationships`

NewCiBuildActionRelationships instantiates a new CiBuildActionRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiBuildActionRelationshipsWithDefaults

`func NewCiBuildActionRelationshipsWithDefaults() *CiBuildActionRelationships`

NewCiBuildActionRelationshipsWithDefaults instantiates a new CiBuildActionRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildRun

`func (o *CiBuildActionRelationships) GetBuildRun() CiBuildActionRelationshipsBuildRun`

GetBuildRun returns the BuildRun field if non-nil, zero value otherwise.

### GetBuildRunOk

`func (o *CiBuildActionRelationships) GetBuildRunOk() (*CiBuildActionRelationshipsBuildRun, bool)`

GetBuildRunOk returns a tuple with the BuildRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRun

`func (o *CiBuildActionRelationships) SetBuildRun(v CiBuildActionRelationshipsBuildRun)`

SetBuildRun sets BuildRun field to given value.

### HasBuildRun

`func (o *CiBuildActionRelationships) HasBuildRun() bool`

HasBuildRun returns a boolean if a field has been set.

### GetArtifacts

`func (o *CiBuildActionRelationships) GetArtifacts() AnalyticsReportInstanceRelationshipsSegments`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *CiBuildActionRelationships) GetArtifactsOk() (*AnalyticsReportInstanceRelationshipsSegments, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *CiBuildActionRelationships) SetArtifacts(v AnalyticsReportInstanceRelationshipsSegments)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *CiBuildActionRelationships) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetIssues

`func (o *CiBuildActionRelationships) GetIssues() AnalyticsReportInstanceRelationshipsSegments`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CiBuildActionRelationships) GetIssuesOk() (*AnalyticsReportInstanceRelationshipsSegments, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CiBuildActionRelationships) SetIssues(v AnalyticsReportInstanceRelationshipsSegments)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CiBuildActionRelationships) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetTestResults

`func (o *CiBuildActionRelationships) GetTestResults() AnalyticsReportInstanceRelationshipsSegments`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *CiBuildActionRelationships) GetTestResultsOk() (*AnalyticsReportInstanceRelationshipsSegments, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *CiBuildActionRelationships) SetTestResults(v AnalyticsReportInstanceRelationshipsSegments)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *CiBuildActionRelationships) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


