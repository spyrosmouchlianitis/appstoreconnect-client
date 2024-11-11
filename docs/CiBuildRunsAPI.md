# \CiBuildRunsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CiBuildRunsActionsGetToManyRelated**](CiBuildRunsAPI.md#CiBuildRunsActionsGetToManyRelated) | **Get** /v1/ciBuildRuns/{id}/actions | 
[**CiBuildRunsBuildsGetToManyRelated**](CiBuildRunsAPI.md#CiBuildRunsBuildsGetToManyRelated) | **Get** /v1/ciBuildRuns/{id}/builds | 
[**CiBuildRunsCreateInstance**](CiBuildRunsAPI.md#CiBuildRunsCreateInstance) | **Post** /v1/ciBuildRuns | 
[**CiBuildRunsGetInstance**](CiBuildRunsAPI.md#CiBuildRunsGetInstance) | **Get** /v1/ciBuildRuns/{id} | 



## CiBuildRunsActionsGetToManyRelated

> CiBuildActionsResponse CiBuildRunsActionsGetToManyRelated(ctx, id).FieldsCiBuildActions(fieldsCiBuildActions).FieldsCiBuildRuns(fieldsCiBuildRuns).Limit(limit).Include(include).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsCiBuildActions := []string{"FieldsCiBuildActions_example"} // []string | the fields to include for returned resources of type ciBuildActions (optional)
	fieldsCiBuildRuns := []string{"FieldsCiBuildRuns_example"} // []string | the fields to include for returned resources of type ciBuildRuns (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildRunsAPI.CiBuildRunsActionsGetToManyRelated(context.Background(), id).FieldsCiBuildActions(fieldsCiBuildActions).FieldsCiBuildRuns(fieldsCiBuildRuns).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildRunsAPI.CiBuildRunsActionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildRunsActionsGetToManyRelated`: CiBuildActionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildRunsAPI.CiBuildRunsActionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildRunsActionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiBuildActions** | **[]string** | the fields to include for returned resources of type ciBuildActions | 
 **fieldsCiBuildRuns** | **[]string** | the fields to include for returned resources of type ciBuildRuns | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**CiBuildActionsResponse**](CiBuildActionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildRunsBuildsGetToManyRelated

> BuildsResponse CiBuildRunsBuildsGetToManyRelated(ctx, id).FilterVersion(filterVersion).FilterExpired(filterExpired).FilterProcessingState(filterProcessingState).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterBuildAudienceType(filterBuildAudienceType).FilterPreReleaseVersion(filterPreReleaseVersion).FilterApp(filterApp).FilterBetaGroups(filterBetaGroups).FilterAppStoreVersion(filterAppStoreVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaGroups(fieldsBetaGroups).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).FieldsBuildBundles(fieldsBuildBundles).Limit(limit).Include(include).LimitIndividualTesters(limitIndividualTesters).LimitBetaGroups(limitBetaGroups).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitBuildBundles(limitBuildBundles).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	filterVersion := []string{"Inner_example"} // []string | filter by attribute 'version' (optional)
	filterExpired := []string{"Inner_example"} // []string | filter by attribute 'expired' (optional)
	filterProcessingState := []string{"FilterProcessingState_example"} // []string | filter by attribute 'processingState' (optional)
	filterBetaAppReviewSubmissionBetaReviewState := []string{"FilterBetaAppReviewSubmissionBetaReviewState_example"} // []string | filter by attribute 'betaAppReviewSubmission.betaReviewState' (optional)
	filterUsesNonExemptEncryption := []string{"Inner_example"} // []string | filter by attribute 'usesNonExemptEncryption' (optional)
	filterPreReleaseVersionVersion := []string{"Inner_example"} // []string | filter by attribute 'preReleaseVersion.version' (optional)
	filterPreReleaseVersionPlatform := []string{"FilterPreReleaseVersionPlatform_example"} // []string | filter by attribute 'preReleaseVersion.platform' (optional)
	filterBuildAudienceType := []string{"FilterBuildAudienceType_example"} // []string | filter by attribute 'buildAudienceType' (optional)
	filterPreReleaseVersion := []string{"Inner_example"} // []string | filter by id(s) of related 'preReleaseVersion' (optional)
	filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
	filterBetaGroups := []string{"Inner_example"} // []string | filter by id(s) of related 'betaGroups' (optional)
	filterAppStoreVersion := []string{"Inner_example"} // []string | filter by id(s) of related 'appStoreVersion' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
	fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
	fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
	fieldsBuildBundles := []string{"FieldsBuildBundles_example"} // []string | the fields to include for returned resources of type buildBundles (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitIndividualTesters := int32(56) // int32 | maximum number of related individualTesters returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitBetaBuildLocalizations := int32(56) // int32 | maximum number of related betaBuildLocalizations returned (when they are included) (optional)
	limitIcons := int32(56) // int32 | maximum number of related icons returned (when they are included) (optional)
	limitBuildBundles := int32(56) // int32 | maximum number of related buildBundles returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildRunsAPI.CiBuildRunsBuildsGetToManyRelated(context.Background(), id).FilterVersion(filterVersion).FilterExpired(filterExpired).FilterProcessingState(filterProcessingState).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterBuildAudienceType(filterBuildAudienceType).FilterPreReleaseVersion(filterPreReleaseVersion).FilterApp(filterApp).FilterBetaGroups(filterBetaGroups).FilterAppStoreVersion(filterAppStoreVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaGroups(fieldsBetaGroups).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).FieldsBuildBundles(fieldsBuildBundles).Limit(limit).Include(include).LimitIndividualTesters(limitIndividualTesters).LimitBetaGroups(limitBetaGroups).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitIcons(limitIcons).LimitBuildBundles(limitBuildBundles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildRunsAPI.CiBuildRunsBuildsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildRunsBuildsGetToManyRelated`: BuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildRunsAPI.CiBuildRunsBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildRunsBuildsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterVersion** | **[]string** | filter by attribute &#39;version&#39; | 
 **filterExpired** | **[]string** | filter by attribute &#39;expired&#39; | 
 **filterProcessingState** | **[]string** | filter by attribute &#39;processingState&#39; | 
 **filterBetaAppReviewSubmissionBetaReviewState** | **[]string** | filter by attribute &#39;betaAppReviewSubmission.betaReviewState&#39; | 
 **filterUsesNonExemptEncryption** | **[]string** | filter by attribute &#39;usesNonExemptEncryption&#39; | 
 **filterPreReleaseVersionVersion** | **[]string** | filter by attribute &#39;preReleaseVersion.version&#39; | 
 **filterPreReleaseVersionPlatform** | **[]string** | filter by attribute &#39;preReleaseVersion.platform&#39; | 
 **filterBuildAudienceType** | **[]string** | filter by attribute &#39;buildAudienceType&#39; | 
 **filterPreReleaseVersion** | **[]string** | filter by id(s) of related &#39;preReleaseVersion&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterBetaGroups** | **[]string** | filter by id(s) of related &#39;betaGroups&#39; | 
 **filterAppStoreVersion** | **[]string** | filter by id(s) of related &#39;appStoreVersion&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **fieldsBuildBundles** | **[]string** | the fields to include for returned resources of type buildBundles | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitIndividualTesters** | **int32** | maximum number of related individualTesters returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBetaBuildLocalizations** | **int32** | maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitIcons** | **int32** | maximum number of related icons returned (when they are included) | 
 **limitBuildBundles** | **int32** | maximum number of related buildBundles returned (when they are included) | 

### Return type

[**BuildsResponse**](BuildsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildRunsCreateInstance

> CiBuildRunResponse CiBuildRunsCreateInstance(ctx).CiBuildRunCreateRequest(ciBuildRunCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ciBuildRunCreateRequest := *openapiclient.NewCiBuildRunCreateRequest(*openapiclient.NewCiBuildRunCreateRequestData("Type_example")) // CiBuildRunCreateRequest | CiBuildRun representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildRunsAPI.CiBuildRunsCreateInstance(context.Background()).CiBuildRunCreateRequest(ciBuildRunCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildRunsAPI.CiBuildRunsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildRunsCreateInstance`: CiBuildRunResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildRunsAPI.CiBuildRunsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildRunsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ciBuildRunCreateRequest** | [**CiBuildRunCreateRequest**](CiBuildRunCreateRequest.md) | CiBuildRun representation | 

### Return type

[**CiBuildRunResponse**](CiBuildRunResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiBuildRunsGetInstance

> CiBuildRunResponse CiBuildRunsGetInstance(ctx, id).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).Include(include).LimitBuilds(limitBuilds).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsCiBuildRuns := []string{"FieldsCiBuildRuns_example"} // []string | the fields to include for returned resources of type ciBuildRuns (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiBuildRunsAPI.CiBuildRunsGetInstance(context.Background(), id).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).Include(include).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiBuildRunsAPI.CiBuildRunsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiBuildRunsGetInstance`: CiBuildRunResponse
	fmt.Fprintf(os.Stdout, "Response from `CiBuildRunsAPI.CiBuildRunsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiBuildRunsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiBuildRuns** | **[]string** | the fields to include for returned resources of type ciBuildRuns | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**CiBuildRunResponse**](CiBuildRunResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

