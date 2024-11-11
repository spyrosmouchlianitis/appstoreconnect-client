# DiagnosticInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsightType** | Pointer to [**DiagnosticInsightType**](DiagnosticInsightType.md) |  | [optional] 
**Direction** | Pointer to [**DiagnosticInsightDirection**](DiagnosticInsightDirection.md) |  | [optional] 
**ReferenceVersions** | Pointer to [**[]DiagnosticInsightReferenceVersionsInner**](DiagnosticInsightReferenceVersionsInner.md) |  | [optional] 

## Methods

### NewDiagnosticInsight

`func NewDiagnosticInsight() *DiagnosticInsight`

NewDiagnosticInsight instantiates a new DiagnosticInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticInsightWithDefaults

`func NewDiagnosticInsightWithDefaults() *DiagnosticInsight`

NewDiagnosticInsightWithDefaults instantiates a new DiagnosticInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsightType

`func (o *DiagnosticInsight) GetInsightType() DiagnosticInsightType`

GetInsightType returns the InsightType field if non-nil, zero value otherwise.

### GetInsightTypeOk

`func (o *DiagnosticInsight) GetInsightTypeOk() (*DiagnosticInsightType, bool)`

GetInsightTypeOk returns a tuple with the InsightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightType

`func (o *DiagnosticInsight) SetInsightType(v DiagnosticInsightType)`

SetInsightType sets InsightType field to given value.

### HasInsightType

`func (o *DiagnosticInsight) HasInsightType() bool`

HasInsightType returns a boolean if a field has been set.

### GetDirection

`func (o *DiagnosticInsight) GetDirection() DiagnosticInsightDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DiagnosticInsight) GetDirectionOk() (*DiagnosticInsightDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DiagnosticInsight) SetDirection(v DiagnosticInsightDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *DiagnosticInsight) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetReferenceVersions

`func (o *DiagnosticInsight) GetReferenceVersions() []DiagnosticInsightReferenceVersionsInner`

GetReferenceVersions returns the ReferenceVersions field if non-nil, zero value otherwise.

### GetReferenceVersionsOk

`func (o *DiagnosticInsight) GetReferenceVersionsOk() (*[]DiagnosticInsightReferenceVersionsInner, bool)`

GetReferenceVersionsOk returns a tuple with the ReferenceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceVersions

`func (o *DiagnosticInsight) SetReferenceVersions(v []DiagnosticInsightReferenceVersionsInner)`

SetReferenceVersions sets ReferenceVersions field to given value.

### HasReferenceVersions

`func (o *DiagnosticInsight) HasReferenceVersions() bool`

HasReferenceVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


