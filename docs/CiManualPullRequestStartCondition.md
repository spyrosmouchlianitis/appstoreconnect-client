# CiManualPullRequestStartCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**CiBranchPatterns**](CiBranchPatterns.md) |  | [optional] 
**Destination** | Pointer to [**CiBranchPatterns**](CiBranchPatterns.md) |  | [optional] 

## Methods

### NewCiManualPullRequestStartCondition

`func NewCiManualPullRequestStartCondition() *CiManualPullRequestStartCondition`

NewCiManualPullRequestStartCondition instantiates a new CiManualPullRequestStartCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiManualPullRequestStartConditionWithDefaults

`func NewCiManualPullRequestStartConditionWithDefaults() *CiManualPullRequestStartCondition`

NewCiManualPullRequestStartConditionWithDefaults instantiates a new CiManualPullRequestStartCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *CiManualPullRequestStartCondition) GetSource() CiBranchPatterns`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CiManualPullRequestStartCondition) GetSourceOk() (*CiBranchPatterns, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CiManualPullRequestStartCondition) SetSource(v CiBranchPatterns)`

SetSource sets Source field to given value.

### HasSource

`func (o *CiManualPullRequestStartCondition) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *CiManualPullRequestStartCondition) GetDestination() CiBranchPatterns`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CiManualPullRequestStartCondition) GetDestinationOk() (*CiBranchPatterns, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CiManualPullRequestStartCondition) SetDestination(v CiBranchPatterns)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CiManualPullRequestStartCondition) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


