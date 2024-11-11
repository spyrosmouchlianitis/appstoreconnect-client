# \BuildsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildsAppEncryptionDeclarationGetToOneRelated**](BuildsAPI.md#BuildsAppEncryptionDeclarationGetToOneRelated) | **Get** /v1/builds/{id}/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationGetToOneRelationship**](BuildsAPI.md#BuildsAppEncryptionDeclarationGetToOneRelationship) | **Get** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppEncryptionDeclarationUpdateToOneRelationship**](BuildsAPI.md#BuildsAppEncryptionDeclarationUpdateToOneRelationship) | **Patch** /v1/builds/{id}/relationships/appEncryptionDeclaration | 
[**BuildsAppGetToOneRelated**](BuildsAPI.md#BuildsAppGetToOneRelated) | **Get** /v1/builds/{id}/app | 
[**BuildsAppStoreVersionGetToOneRelated**](BuildsAPI.md#BuildsAppStoreVersionGetToOneRelated) | **Get** /v1/builds/{id}/appStoreVersion | 
[**BuildsBetaAppReviewSubmissionGetToOneRelated**](BuildsAPI.md#BuildsBetaAppReviewSubmissionGetToOneRelated) | **Get** /v1/builds/{id}/betaAppReviewSubmission | 
[**BuildsBetaBuildLocalizationsGetToManyRelated**](BuildsAPI.md#BuildsBetaBuildLocalizationsGetToManyRelated) | **Get** /v1/builds/{id}/betaBuildLocalizations | 
[**BuildsBetaBuildUsagesGetMetrics**](BuildsAPI.md#BuildsBetaBuildUsagesGetMetrics) | **Get** /v1/builds/{id}/metrics/betaBuildUsages | 
[**BuildsBetaGroupsCreateToManyRelationship**](BuildsAPI.md#BuildsBetaGroupsCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBetaGroupsDeleteToManyRelationship**](BuildsAPI.md#BuildsBetaGroupsDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/betaGroups | 
[**BuildsBuildBetaDetailGetToOneRelated**](BuildsAPI.md#BuildsBuildBetaDetailGetToOneRelated) | **Get** /v1/builds/{id}/buildBetaDetail | 
[**BuildsDiagnosticSignaturesGetToManyRelated**](BuildsAPI.md#BuildsDiagnosticSignaturesGetToManyRelated) | **Get** /v1/builds/{id}/diagnosticSignatures | 
[**BuildsGetCollection**](BuildsAPI.md#BuildsGetCollection) | **Get** /v1/builds | 
[**BuildsGetInstance**](BuildsAPI.md#BuildsGetInstance) | **Get** /v1/builds/{id} | 
[**BuildsIconsGetToManyRelated**](BuildsAPI.md#BuildsIconsGetToManyRelated) | **Get** /v1/builds/{id}/icons | 
[**BuildsIndividualTestersCreateToManyRelationship**](BuildsAPI.md#BuildsIndividualTestersCreateToManyRelationship) | **Post** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersDeleteToManyRelationship**](BuildsAPI.md#BuildsIndividualTestersDeleteToManyRelationship) | **Delete** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsIndividualTestersGetToManyRelated**](BuildsAPI.md#BuildsIndividualTestersGetToManyRelated) | **Get** /v1/builds/{id}/individualTesters | 
[**BuildsIndividualTestersGetToManyRelationship**](BuildsAPI.md#BuildsIndividualTestersGetToManyRelationship) | **Get** /v1/builds/{id}/relationships/individualTesters | 
[**BuildsPerfPowerMetricsGetToManyRelated**](BuildsAPI.md#BuildsPerfPowerMetricsGetToManyRelated) | **Get** /v1/builds/{id}/perfPowerMetrics | 
[**BuildsPreReleaseVersionGetToOneRelated**](BuildsAPI.md#BuildsPreReleaseVersionGetToOneRelated) | **Get** /v1/builds/{id}/preReleaseVersion | 
[**BuildsUpdateInstance**](BuildsAPI.md#BuildsUpdateInstance) | **Patch** /v1/builds/{id} | 



## BuildsAppEncryptionDeclarationGetToOneRelated

> AppEncryptionDeclarationWithoutIncludesResponse BuildsAppEncryptionDeclarationGetToOneRelated(ctx, id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelated(context.Background(), id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsAppEncryptionDeclarationGetToOneRelated`: AppEncryptionDeclarationWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 

### Return type

[**AppEncryptionDeclarationWithoutIncludesResponse**](AppEncryptionDeclarationWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationGetToOneRelationship

> BuildAppEncryptionDeclarationLinkageResponse BuildsAppEncryptionDeclarationGetToOneRelationship(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelationship(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsAppEncryptionDeclarationGetToOneRelationship`: BuildAppEncryptionDeclarationLinkageResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsAppEncryptionDeclarationGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuildAppEncryptionDeclarationLinkageResponse**](BuildAppEncryptionDeclarationLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppEncryptionDeclarationUpdateToOneRelationship

> BuildsAppEncryptionDeclarationUpdateToOneRelationship(ctx, id).BuildAppEncryptionDeclarationLinkageRequest(buildAppEncryptionDeclarationLinkageRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildAppEncryptionDeclarationLinkageRequest := *openapiclient.NewBuildAppEncryptionDeclarationLinkageRequest(*openapiclient.NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData("Type_example", "Id_example")) // BuildAppEncryptionDeclarationLinkageRequest | Related linkage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAPI.BuildsAppEncryptionDeclarationUpdateToOneRelationship(context.Background(), id).BuildAppEncryptionDeclarationLinkageRequest(buildAppEncryptionDeclarationLinkageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsAppEncryptionDeclarationUpdateToOneRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppEncryptionDeclarationUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildAppEncryptionDeclarationLinkageRequest** | [**BuildAppEncryptionDeclarationLinkageRequest**](BuildAppEncryptionDeclarationLinkageRequest.md) | Related linkage | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppGetToOneRelated

> AppWithoutIncludesResponse BuildsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsAppGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsAppGetToOneRelated`: AppWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**AppWithoutIncludesResponse**](AppWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsAppStoreVersionGetToOneRelated

> AppStoreVersionResponse BuildsAppStoreVersionGetToOneRelated(ctx, id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Include(include).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)
	fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)
	fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
	fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)
	fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
	fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
	fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
	fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
	fieldsAlternativeDistributionPackages := []string{"FieldsAlternativeDistributionPackages_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackages (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppStoreVersionLocalizations := int32(56) // int32 | maximum number of related appStoreVersionLocalizations returned (when they are included) (optional)
	limitAppStoreVersionExperiments := int32(56) // int32 | maximum number of related appStoreVersionExperiments returned (when they are included) (optional)
	limitAppStoreVersionExperimentsV2 := int32(56) // int32 | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsAppStoreVersionGetToOneRelated(context.Background(), id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Include(include).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsAppStoreVersionGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsAppStoreVersionGetToOneRelated`: AppStoreVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsAppStoreVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsAppStoreVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **fieldsAlternativeDistributionPackages** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackages | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppStoreVersionLocalizations** | **int32** | maximum number of related appStoreVersionLocalizations returned (when they are included) | 
 **limitAppStoreVersionExperiments** | **int32** | maximum number of related appStoreVersionExperiments returned (when they are included) | 
 **limitAppStoreVersionExperimentsV2** | **int32** | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaAppReviewSubmissionGetToOneRelated

> BetaAppReviewSubmissionWithoutIncludesResponse BuildsBetaAppReviewSubmissionGetToOneRelated(ctx, id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBetaAppReviewSubmissionGetToOneRelated(context.Background(), id).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBetaAppReviewSubmissionGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBetaAppReviewSubmissionGetToOneRelated`: BetaAppReviewSubmissionWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBetaAppReviewSubmissionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaAppReviewSubmissionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 

### Return type

[**BetaAppReviewSubmissionWithoutIncludesResponse**](BetaAppReviewSubmissionWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaBuildLocalizationsGetToManyRelated

> BetaBuildLocalizationsWithoutIncludesResponse BuildsBetaBuildLocalizationsGetToManyRelated(ctx, id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBetaBuildLocalizationsGetToManyRelated(context.Background(), id).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBetaBuildLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBetaBuildLocalizationsGetToManyRelated`: BetaBuildLocalizationsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBetaBuildLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaBuildLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaBuildLocalizationsWithoutIncludesResponse**](BetaBuildLocalizationsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaBuildUsagesGetMetrics

> BetaBuildUsagesV1MetricResponse BuildsBetaBuildUsagesGetMetrics(ctx, id).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBetaBuildUsagesGetMetrics(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBetaBuildUsagesGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBetaBuildUsagesGetMetrics`: BetaBuildUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBetaBuildUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaBuildUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**BetaBuildUsagesV1MetricResponse**](BetaBuildUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaGroupsCreateToManyRelationship

> BuildsBetaGroupsCreateToManyRelationship(ctx, id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildBetaGroupsLinkagesRequest := *openapiclient.NewBuildBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsDataInner{*openapiclient.NewAppRelationshipsBetaGroupsDataInner("Type_example", "Id_example")}) // BuildBetaGroupsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAPI.BuildsBetaGroupsCreateToManyRelationship(context.Background(), id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBetaGroupsCreateToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaGroupsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaGroupsDeleteToManyRelationship

> BuildsBetaGroupsDeleteToManyRelationship(ctx, id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildBetaGroupsLinkagesRequest := *openapiclient.NewBuildBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsDataInner{*openapiclient.NewAppRelationshipsBetaGroupsDataInner("Type_example", "Id_example")}) // BuildBetaGroupsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAPI.BuildsBetaGroupsDeleteToManyRelationship(context.Background(), id).BuildBetaGroupsLinkagesRequest(buildBetaGroupsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBetaGroupsDeleteToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaGroupsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildBetaGroupsLinkagesRequest** | [**BuildBetaGroupsLinkagesRequest**](BuildBetaGroupsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBuildBetaDetailGetToOneRelated

> BuildBetaDetailResponse BuildsBuildBetaDetailGetToOneRelated(ctx, id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuilds(fieldsBuilds).Include(include).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBuildBetaDetailGetToOneRelated(context.Background(), id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsBuilds(fieldsBuilds).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBuildBetaDetailGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildBetaDetailGetToOneRelated`: BuildBetaDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBuildBetaDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildBetaDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsDiagnosticSignaturesGetToManyRelated

> DiagnosticSignaturesResponse BuildsDiagnosticSignaturesGetToManyRelated(ctx, id).FilterDiagnosticType(filterDiagnosticType).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	filterDiagnosticType := []string{"FilterDiagnosticType_example"} // []string | filter by attribute 'diagnosticType' (optional)
	fieldsDiagnosticSignatures := []string{"FieldsDiagnosticSignatures_example"} // []string | the fields to include for returned resources of type diagnosticSignatures (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsDiagnosticSignaturesGetToManyRelated(context.Background(), id).FilterDiagnosticType(filterDiagnosticType).FieldsDiagnosticSignatures(fieldsDiagnosticSignatures).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsDiagnosticSignaturesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsDiagnosticSignaturesGetToManyRelated`: DiagnosticSignaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsDiagnosticSignaturesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsDiagnosticSignaturesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterDiagnosticType** | **[]string** | filter by attribute &#39;diagnosticType&#39; | 
 **fieldsDiagnosticSignatures** | **[]string** | the fields to include for returned resources of type diagnosticSignatures | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**DiagnosticSignaturesResponse**](DiagnosticSignaturesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsGetCollection

> BuildsResponse BuildsGetCollection(ctx).FilterVersion(filterVersion).FilterExpired(filterExpired).FilterProcessingState(filterProcessingState).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterBuildAudienceType(filterBuildAudienceType).FilterPreReleaseVersion(filterPreReleaseVersion).FilterApp(filterApp).FilterBetaGroups(filterBetaGroups).FilterAppStoreVersion(filterAppStoreVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Include(include).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuildBundles(limitBuildBundles).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
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
	fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
	fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
	fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBetaBuildLocalizations := int32(56) // int32 | maximum number of related betaBuildLocalizations returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitBuildBundles := int32(56) // int32 | maximum number of related buildBundles returned (when they are included) (optional)
	limitIcons := int32(56) // int32 | maximum number of related icons returned (when they are included) (optional)
	limitIndividualTesters := int32(56) // int32 | maximum number of related individualTesters returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsGetCollection(context.Background()).FilterVersion(filterVersion).FilterExpired(filterExpired).FilterProcessingState(filterProcessingState).FilterBetaAppReviewSubmissionBetaReviewState(filterBetaAppReviewSubmissionBetaReviewState).FilterUsesNonExemptEncryption(filterUsesNonExemptEncryption).FilterPreReleaseVersionVersion(filterPreReleaseVersionVersion).FilterPreReleaseVersionPlatform(filterPreReleaseVersionPlatform).FilterBuildAudienceType(filterBuildAudienceType).FilterPreReleaseVersion(filterPreReleaseVersion).FilterApp(filterApp).FilterBetaGroups(filterBetaGroups).FilterAppStoreVersion(filterAppStoreVersion).FilterId(filterId).Sort(sort).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Include(include).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuildBundles(limitBuildBundles).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsGetCollection`: BuildsResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildsGetCollectionRequest struct via the builder pattern


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
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBetaBuildLocalizations** | **int32** | maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuildBundles** | **int32** | maximum number of related buildBundles returned (when they are included) | 
 **limitIcons** | **int32** | maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **int32** | maximum number of related individualTesters returned (when they are included) | 

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


## BuildsGetInstance

> BuildResponse BuildsGetInstance(ctx, id).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).Include(include).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuildBundles(limitBuildBundles).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	fieldsBetaBuildLocalizations := []string{"FieldsBetaBuildLocalizations_example"} // []string | the fields to include for returned resources of type betaBuildLocalizations (optional)
	fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
	fieldsBetaAppReviewSubmissions := []string{"FieldsBetaAppReviewSubmissions_example"} // []string | the fields to include for returned resources of type betaAppReviewSubmissions (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBetaBuildLocalizations := int32(56) // int32 | maximum number of related betaBuildLocalizations returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitBuildBundles := int32(56) // int32 | maximum number of related buildBundles returned (when they are included) (optional)
	limitIcons := int32(56) // int32 | maximum number of related icons returned (when they are included) (optional)
	limitIndividualTesters := int32(56) // int32 | maximum number of related individualTesters returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsGetInstance(context.Background(), id).FieldsBuilds(fieldsBuilds).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaTesters(fieldsBetaTesters).FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions).FieldsApps(fieldsApps).FieldsBuildBetaDetails(fieldsBuildBetaDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsBuildIcons(fieldsBuildIcons).Include(include).LimitBetaBuildLocalizations(limitBetaBuildLocalizations).LimitBetaGroups(limitBetaGroups).LimitBuildBundles(limitBuildBundles).LimitIcons(limitIcons).LimitIndividualTesters(limitIndividualTesters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsGetInstance`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsBetaBuildLocalizations** | **[]string** | the fields to include for returned resources of type betaBuildLocalizations | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsBetaAppReviewSubmissions** | **[]string** | the fields to include for returned resources of type betaAppReviewSubmissions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBetaBuildLocalizations** | **int32** | maximum number of related betaBuildLocalizations returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuildBundles** | **int32** | maximum number of related buildBundles returned (when they are included) | 
 **limitIcons** | **int32** | maximum number of related icons returned (when they are included) | 
 **limitIndividualTesters** | **int32** | maximum number of related individualTesters returned (when they are included) | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIconsGetToManyRelated

> BuildIconsWithoutIncludesResponse BuildsIconsGetToManyRelated(ctx, id).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBuildIcons := []string{"FieldsBuildIcons_example"} // []string | the fields to include for returned resources of type buildIcons (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsIconsGetToManyRelated(context.Background(), id).FieldsBuildIcons(fieldsBuildIcons).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsIconsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsIconsGetToManyRelated`: BuildIconsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsIconsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIconsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildIcons** | **[]string** | the fields to include for returned resources of type buildIcons | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildIconsWithoutIncludesResponse**](BuildIconsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersCreateToManyRelationship

> BuildsIndividualTestersCreateToManyRelationship(ctx, id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildIndividualTestersLinkagesRequest := *openapiclient.NewBuildIndividualTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersDataInner{*openapiclient.NewBetaGroupRelationshipsBetaTestersDataInner("Type_example", "Id_example")}) // BuildIndividualTestersLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAPI.BuildsIndividualTestersCreateToManyRelationship(context.Background(), id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsIndividualTestersCreateToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersDeleteToManyRelationship

> BuildsIndividualTestersDeleteToManyRelationship(ctx, id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildIndividualTestersLinkagesRequest := *openapiclient.NewBuildIndividualTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersDataInner{*openapiclient.NewBetaGroupRelationshipsBetaTestersDataInner("Type_example", "Id_example")}) // BuildIndividualTestersLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAPI.BuildsIndividualTestersDeleteToManyRelationship(context.Background(), id).BuildIndividualTestersLinkagesRequest(buildIndividualTestersLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsIndividualTestersDeleteToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildIndividualTestersLinkagesRequest** | [**BuildIndividualTestersLinkagesRequest**](BuildIndividualTestersLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersGetToManyRelated

> BetaTestersWithoutIncludesResponse BuildsIndividualTestersGetToManyRelated(ctx, id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsIndividualTestersGetToManyRelated(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsIndividualTestersGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsIndividualTestersGetToManyRelated`: BetaTestersWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsIndividualTestersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTestersWithoutIncludesResponse**](BetaTestersWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsIndividualTestersGetToManyRelationship

> BuildIndividualTestersLinkagesResponse BuildsIndividualTestersGetToManyRelationship(ctx, id).Limit(limit).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsIndividualTestersGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsIndividualTestersGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsIndividualTestersGetToManyRelationship`: BuildIndividualTestersLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsIndividualTestersGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsIndividualTestersGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildIndividualTestersLinkagesResponse**](BuildIndividualTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPerfPowerMetricsGetToManyRelated

> XcodeMetrics BuildsPerfPowerMetricsGetToManyRelated(ctx, id).FilterPlatform(filterPlatform).FilterMetricType(filterMetricType).FilterDeviceType(filterDeviceType).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
	filterMetricType := []string{"FilterMetricType_example"} // []string | filter by attribute 'metricType' (optional)
	filterDeviceType := []string{"Inner_example"} // []string | filter by attribute 'deviceType' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsPerfPowerMetricsGetToManyRelated(context.Background(), id).FilterPlatform(filterPlatform).FilterMetricType(filterMetricType).FilterDeviceType(filterDeviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsPerfPowerMetricsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsPerfPowerMetricsGetToManyRelated`: XcodeMetrics
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsPerfPowerMetricsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsPerfPowerMetricsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterMetricType** | **[]string** | filter by attribute &#39;metricType&#39; | 
 **filterDeviceType** | **[]string** | filter by attribute &#39;deviceType&#39; | 

### Return type

[**XcodeMetrics**](XcodeMetrics.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.apple.xcode-metrics+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsPreReleaseVersionGetToOneRelated

> PrereleaseVersionWithoutIncludesResponse BuildsPreReleaseVersionGetToOneRelated(ctx, id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsPreReleaseVersionGetToOneRelated(context.Background(), id).FieldsPreReleaseVersions(fieldsPreReleaseVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsPreReleaseVersionGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsPreReleaseVersionGetToOneRelated`: PrereleaseVersionWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsPreReleaseVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsPreReleaseVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 

### Return type

[**PrereleaseVersionWithoutIncludesResponse**](PrereleaseVersionWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsUpdateInstance

> BuildResponse BuildsUpdateInstance(ctx, id).BuildUpdateRequest(buildUpdateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/spyrosmouchlianitis/appstoreconnect-client"
)

func main() {
	id := "id_example" // string | the id of the requested resource
	buildUpdateRequest := *openapiclient.NewBuildUpdateRequest(*openapiclient.NewBuildUpdateRequestData("Type_example", "Id_example")) // BuildUpdateRequest | Build representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsUpdateInstance(context.Background(), id).BuildUpdateRequest(buildUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsUpdateInstance`: BuildResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildUpdateRequest** | [**BuildUpdateRequest**](BuildUpdateRequest.md) | Build representation | 

### Return type

[**BuildResponse**](BuildResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

