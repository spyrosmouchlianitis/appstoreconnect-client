# \AppStoreVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppStoreVersionsAgeRatingDeclarationGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAgeRatingDeclarationGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/ageRatingDeclaration | 
[**AppStoreVersionsAlternativeDistributionPackageGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAlternativeDistributionPackageGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/alternativeDistributionPackage | 
[**AppStoreVersionsAppClipDefaultExperienceGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppClipDefaultExperienceGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appClipDefaultExperience | 
[**AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship**](AppStoreVersionsAPI.md#AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship) | **Get** /v1/appStoreVersions/{id}/relationships/appClipDefaultExperience | 
[**AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship**](AppStoreVersionsAPI.md#AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship) | **Patch** /v1/appStoreVersions/{id}/relationships/appClipDefaultExperience | 
[**AppStoreVersionsAppStoreReviewDetailGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreReviewDetailGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreReviewDetail | 
[**AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionExperiments | 
[**AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionExperimentsV2 | 
[**AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionLocalizations | 
[**AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionPhasedRelease | 
[**AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/appStoreVersionSubmission | 
[**AppStoreVersionsBuildGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsBuildGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/build | 
[**AppStoreVersionsBuildGetToOneRelationship**](AppStoreVersionsAPI.md#AppStoreVersionsBuildGetToOneRelationship) | **Get** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsBuildUpdateToOneRelationship**](AppStoreVersionsAPI.md#AppStoreVersionsBuildUpdateToOneRelationship) | **Patch** /v1/appStoreVersions/{id}/relationships/build | 
[**AppStoreVersionsCreateInstance**](AppStoreVersionsAPI.md#AppStoreVersionsCreateInstance) | **Post** /v1/appStoreVersions | 
[**AppStoreVersionsCustomerReviewsGetToManyRelated**](AppStoreVersionsAPI.md#AppStoreVersionsCustomerReviewsGetToManyRelated) | **Get** /v1/appStoreVersions/{id}/customerReviews | 
[**AppStoreVersionsDeleteInstance**](AppStoreVersionsAPI.md#AppStoreVersionsDeleteInstance) | **Delete** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsGameCenterAppVersionGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsGameCenterAppVersionGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/gameCenterAppVersion | 
[**AppStoreVersionsGetInstance**](AppStoreVersionsAPI.md#AppStoreVersionsGetInstance) | **Get** /v1/appStoreVersions/{id} | 
[**AppStoreVersionsRoutingAppCoverageGetToOneRelated**](AppStoreVersionsAPI.md#AppStoreVersionsRoutingAppCoverageGetToOneRelated) | **Get** /v1/appStoreVersions/{id}/routingAppCoverage | 
[**AppStoreVersionsUpdateInstance**](AppStoreVersionsAPI.md#AppStoreVersionsUpdateInstance) | **Patch** /v1/appStoreVersions/{id} | 



## AppStoreVersionsAgeRatingDeclarationGetToOneRelated

> AgeRatingDeclarationWithoutIncludesResponse AppStoreVersionsAgeRatingDeclarationGetToOneRelated(ctx, id).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).Execute()



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
	fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAgeRatingDeclarationGetToOneRelated(context.Background(), id).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAgeRatingDeclarationGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAgeRatingDeclarationGetToOneRelated`: AgeRatingDeclarationWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAgeRatingDeclarationGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAgeRatingDeclarationGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 

### Return type

[**AgeRatingDeclarationWithoutIncludesResponse**](AgeRatingDeclarationWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAlternativeDistributionPackageGetToOneRelated

> AlternativeDistributionPackageResponse AppStoreVersionsAlternativeDistributionPackageGetToOneRelated(ctx, id).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).Include(include).LimitVersions(limitVersions).Execute()



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
	fieldsAlternativeDistributionPackages := []string{"FieldsAlternativeDistributionPackages_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackages (optional)
	fieldsAlternativeDistributionPackageVersions := []string{"FieldsAlternativeDistributionPackageVersions_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitVersions := int32(56) // int32 | maximum number of related versions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAlternativeDistributionPackageGetToOneRelated(context.Background(), id).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions).Include(include).LimitVersions(limitVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAlternativeDistributionPackageGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAlternativeDistributionPackageGetToOneRelated`: AlternativeDistributionPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAlternativeDistributionPackageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAlternativeDistributionPackageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackages** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackages | 
 **fieldsAlternativeDistributionPackageVersions** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitVersions** | **int32** | maximum number of related versions returned (when they are included) | 

### Return type

[**AlternativeDistributionPackageResponse**](AlternativeDistributionPackageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppClipDefaultExperienceGetToOneRelated

> AppClipDefaultExperienceResponse AppStoreVersionsAppClipDefaultExperienceGetToOneRelated(ctx, id).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClips(fieldsAppClips).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).Include(include).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()



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
	fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
	fieldsAppClips := []string{"FieldsAppClips_example"} // []string | the fields to include for returned resources of type appClips (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppClipDefaultExperienceLocalizations := []string{"FieldsAppClipDefaultExperienceLocalizations_example"} // []string | the fields to include for returned resources of type appClipDefaultExperienceLocalizations (optional)
	fieldsAppClipAppStoreReviewDetails := []string{"FieldsAppClipAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appClipAppStoreReviewDetails (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppClipDefaultExperienceLocalizations := int32(56) // int32 | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelated(context.Background(), id).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClips(fieldsAppClips).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).Include(include).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppClipDefaultExperienceGetToOneRelated`: AppClipDefaultExperienceResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppClipDefaultExperienceGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **fieldsAppClips** | **[]string** | the fields to include for returned resources of type appClips | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppClipDefaultExperienceLocalizations** | **[]string** | the fields to include for returned resources of type appClipDefaultExperienceLocalizations | 
 **fieldsAppClipAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appClipAppStoreReviewDetails | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppClipDefaultExperienceLocalizations** | **int32** | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) | 

### Return type

[**AppClipDefaultExperienceResponse**](AppClipDefaultExperienceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship

> AppStoreVersionAppClipDefaultExperienceLinkageResponse AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship(ctx, id).Execute()



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
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship`: AppStoreVersionAppClipDefaultExperienceLinkageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppClipDefaultExperienceGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppStoreVersionAppClipDefaultExperienceLinkageResponse**](AppStoreVersionAppClipDefaultExperienceLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship

> AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship(ctx, id).AppStoreVersionAppClipDefaultExperienceLinkageRequest(appStoreVersionAppClipDefaultExperienceLinkageRequest).Execute()



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
	appStoreVersionAppClipDefaultExperienceLinkageRequest := *openapiclient.NewAppStoreVersionAppClipDefaultExperienceLinkageRequest(*openapiclient.NewAppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData("Type_example", "Id_example")) // AppStoreVersionAppClipDefaultExperienceLinkageRequest | Related linkage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship(context.Background(), id).AppStoreVersionAppClipDefaultExperienceLinkageRequest(appStoreVersionAppClipDefaultExperienceLinkageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionsAppClipDefaultExperienceUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionAppClipDefaultExperienceLinkageRequest** | [**AppStoreVersionAppClipDefaultExperienceLinkageRequest**](AppStoreVersionAppClipDefaultExperienceLinkageRequest.md) | Related linkage | 

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


## AppStoreVersionsAppStoreReviewDetailGetToOneRelated

> AppStoreReviewDetailResponse AppStoreVersionsAppStoreReviewDetailGetToOneRelated(ctx, id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).LimitAppStoreReviewAttachments(limitAppStoreReviewAttachments).Execute()



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
	fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppStoreReviewAttachments := []string{"FieldsAppStoreReviewAttachments_example"} // []string | the fields to include for returned resources of type appStoreReviewAttachments (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppStoreReviewAttachments := int32(56) // int32 | maximum number of related appStoreReviewAttachments returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreReviewDetailGetToOneRelated(context.Background(), id).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments).Include(include).LimitAppStoreReviewAttachments(limitAppStoreReviewAttachments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreReviewDetailGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreReviewDetailGetToOneRelated`: AppStoreReviewDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreReviewDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreReviewDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreReviewAttachments** | **[]string** | the fields to include for returned resources of type appStoreReviewAttachments | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppStoreReviewAttachments** | **int32** | maximum number of related appStoreReviewAttachments returned (when they are included) | 

### Return type

[**AppStoreReviewDetailResponse**](AppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated

> AppStoreVersionExperimentsResponse AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated(ctx, id).FilterState(filterState).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).Limit(limit).Include(include).LimitAppStoreVersionExperimentTreatments(limitAppStoreVersionExperimentTreatments).Execute()



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
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppStoreVersionExperimentTreatments := []string{"FieldsAppStoreVersionExperimentTreatments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatments (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppStoreVersionExperimentTreatments := int32(56) // int32 | maximum number of related appStoreVersionExperimentTreatments returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated(context.Background(), id).FilterState(filterState).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).Limit(limit).Include(include).LimitAppStoreVersionExperimentTreatments(limitAppStoreVersionExperimentTreatments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated`: AppStoreVersionExperimentsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionExperimentsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreVersionExperimentTreatments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatments | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppStoreVersionExperimentTreatments** | **int32** | maximum number of related appStoreVersionExperimentTreatments returned (when they are included) | 

### Return type

[**AppStoreVersionExperimentsResponse**](AppStoreVersionExperimentsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated

> AppStoreVersionExperimentsV2Response AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated(ctx, id).FilterState(filterState).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsApps(fieldsApps).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).Limit(limit).Include(include).LimitControlVersions(limitControlVersions).LimitAppStoreVersionExperimentTreatments(limitAppStoreVersionExperimentTreatments).Execute()



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
	filterState := []string{"FilterState_example"} // []string | filter by attribute 'state' (optional)
	fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppStoreVersionExperimentTreatments := []string{"FieldsAppStoreVersionExperimentTreatments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperimentTreatments (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitControlVersions := int32(56) // int32 | maximum number of related controlVersions returned (when they are included) (optional)
	limitAppStoreVersionExperimentTreatments := int32(56) // int32 | maximum number of related appStoreVersionExperimentTreatments returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated(context.Background(), id).FilterState(filterState).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsApps(fieldsApps).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppStoreVersionExperimentTreatments(fieldsAppStoreVersionExperimentTreatments).Limit(limit).Include(include).LimitControlVersions(limitControlVersions).LimitAppStoreVersionExperimentTreatments(limitAppStoreVersionExperimentTreatments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated`: AppStoreVersionExperimentsV2Response
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionExperimentsV2GetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterState** | **[]string** | filter by attribute &#39;state&#39; | 
 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppStoreVersionExperimentTreatments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperimentTreatments | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitControlVersions** | **int32** | maximum number of related controlVersions returned (when they are included) | 
 **limitAppStoreVersionExperimentTreatments** | **int32** | maximum number of related appStoreVersionExperimentTreatments returned (when they are included) | 

### Return type

[**AppStoreVersionExperimentsV2Response**](AppStoreVersionExperimentsV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated

> AppStoreVersionLocalizationsResponse AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated(ctx, id).FilterLocale(filterLocale).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()



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
	filterLocale := []string{"Inner_example"} // []string | filter by attribute 'locale' (optional)
	fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsAppScreenshotSets := []string{"FieldsAppScreenshotSets_example"} // []string | the fields to include for returned resources of type appScreenshotSets (optional)
	fieldsAppPreviewSets := []string{"FieldsAppPreviewSets_example"} // []string | the fields to include for returned resources of type appPreviewSets (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppScreenshotSets := int32(56) // int32 | maximum number of related appScreenshotSets returned (when they are included) (optional)
	limitAppPreviewSets := int32(56) // int32 | maximum number of related appPreviewSets returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated(context.Background(), id).FilterLocale(filterLocale).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppScreenshotSets(fieldsAppScreenshotSets).FieldsAppPreviewSets(fieldsAppPreviewSets).Limit(limit).Include(include).LimitAppScreenshotSets(limitAppScreenshotSets).LimitAppPreviewSets(limitAppPreviewSets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated`: AppStoreVersionLocalizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppScreenshotSets** | **[]string** | the fields to include for returned resources of type appScreenshotSets | 
 **fieldsAppPreviewSets** | **[]string** | the fields to include for returned resources of type appPreviewSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppScreenshotSets** | **int32** | maximum number of related appScreenshotSets returned (when they are included) | 
 **limitAppPreviewSets** | **int32** | maximum number of related appPreviewSets returned (when they are included) | 

### Return type

[**AppStoreVersionLocalizationsResponse**](AppStoreVersionLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated

> AppStoreVersionPhasedReleaseWithoutIncludesResponse AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated(ctx, id).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).Execute()



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
	fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated(context.Background(), id).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated`: AppStoreVersionPhasedReleaseWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionPhasedReleaseGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 

### Return type

[**AppStoreVersionPhasedReleaseWithoutIncludesResponse**](AppStoreVersionPhasedReleaseWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated

> AppStoreVersionSubmissionResponse AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated(ctx, id).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).Execute()



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
	fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated(context.Background(), id).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated`: AppStoreVersionSubmissionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsAppStoreVersionSubmissionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsAppStoreVersionSubmissionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionSubmissionResponse**](AppStoreVersionSubmissionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildGetToOneRelated

> BuildWithoutIncludesResponse AppStoreVersionsBuildGetToOneRelated(ctx, id).FieldsBuilds(fieldsBuilds).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsBuildGetToOneRelated`: BuildWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BuildWithoutIncludesResponse**](BuildWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildGetToOneRelationship

> AppStoreVersionBuildLinkageResponse AppStoreVersionsBuildGetToOneRelationship(ctx, id).Execute()



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
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelationship(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsBuildGetToOneRelationship`: AppStoreVersionBuildLinkageResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsBuildGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppStoreVersionBuildLinkageResponse**](AppStoreVersionBuildLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsBuildUpdateToOneRelationship

> AppStoreVersionsBuildUpdateToOneRelationship(ctx, id).AppStoreVersionBuildLinkageRequest(appStoreVersionBuildLinkageRequest).Execute()



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
	appStoreVersionBuildLinkageRequest := *openapiclient.NewAppStoreVersionBuildLinkageRequest(*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")) // AppStoreVersionBuildLinkageRequest | Related linkage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsBuildUpdateToOneRelationship(context.Background(), id).AppStoreVersionBuildLinkageRequest(appStoreVersionBuildLinkageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsBuildUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionsBuildUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionBuildLinkageRequest** | [**AppStoreVersionBuildLinkageRequest**](AppStoreVersionBuildLinkageRequest.md) | Related linkage | 

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


## AppStoreVersionsCreateInstance

> AppStoreVersionResponse AppStoreVersionsCreateInstance(ctx).AppStoreVersionCreateRequest(appStoreVersionCreateRequest).Execute()



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
	appStoreVersionCreateRequest := *openapiclient.NewAppStoreVersionCreateRequest(*openapiclient.NewAppStoreVersionCreateRequestData("Type_example", *openapiclient.NewAppStoreVersionCreateRequestDataAttributes(openapiclient.Platform("IOS"), "VersionString_example"), *openapiclient.NewAppStoreVersionCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // AppStoreVersionCreateRequest | AppStoreVersion representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsCreateInstance(context.Background()).AppStoreVersionCreateRequest(appStoreVersionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsCreateInstance`: AppStoreVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreVersionCreateRequest** | [**AppStoreVersionCreateRequest**](AppStoreVersionCreateRequest.md) | AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsCustomerReviewsGetToManyRelated

> CustomerReviewsResponse AppStoreVersionsCustomerReviewsGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FilterRating(filterRating).ExistsPublishedResponse(existsPublishedResponse).Sort(sort).FieldsCustomerReviews(fieldsCustomerReviews).FieldsCustomerReviewResponses(fieldsCustomerReviewResponses).Limit(limit).Include(include).Execute()



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
	filterTerritory := []string{"FilterTerritory_example"} // []string | filter by attribute 'territory' (optional)
	filterRating := []string{"Inner_example"} // []string | filter by attribute 'rating' (optional)
	existsPublishedResponse := true // bool | filter by publishedResponse (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsCustomerReviews := []string{"FieldsCustomerReviews_example"} // []string | the fields to include for returned resources of type customerReviews (optional)
	fieldsCustomerReviewResponses := []string{"FieldsCustomerReviewResponses_example"} // []string | the fields to include for returned resources of type customerReviewResponses (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsCustomerReviewsGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FilterRating(filterRating).ExistsPublishedResponse(existsPublishedResponse).Sort(sort).FieldsCustomerReviews(fieldsCustomerReviews).FieldsCustomerReviewResponses(fieldsCustomerReviewResponses).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsCustomerReviewsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsCustomerReviewsGetToManyRelated`: CustomerReviewsResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsCustomerReviewsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsCustomerReviewsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by attribute &#39;territory&#39; | 
 **filterRating** | **[]string** | filter by attribute &#39;rating&#39; | 
 **existsPublishedResponse** | **bool** | filter by publishedResponse | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsCustomerReviews** | **[]string** | the fields to include for returned resources of type customerReviews | 
 **fieldsCustomerReviewResponses** | **[]string** | the fields to include for returned resources of type customerReviewResponses | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**CustomerReviewsResponse**](CustomerReviewsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsDeleteInstance

> AppStoreVersionsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppStoreVersionsDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsGameCenterAppVersionGetToOneRelated

> GameCenterAppVersionResponse AppStoreVersionsGameCenterAppVersionGetToOneRelated(ctx, id).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()



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
	fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitCompatibilityVersions := int32(56) // int32 | maximum number of related compatibilityVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsGameCenterAppVersionGetToOneRelated(context.Background(), id).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).Include(include).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsGameCenterAppVersionGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsGameCenterAppVersionGetToOneRelated`: GameCenterAppVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsGameCenterAppVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsGameCenterAppVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitCompatibilityVersions** | **int32** | maximum number of related compatibilityVersions returned (when they are included) | 

### Return type

[**GameCenterAppVersionResponse**](GameCenterAppVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsGetInstance

> AppStoreVersionResponse AppStoreVersionsGetInstance(ctx, id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Include(include).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).Execute()



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
	limitAppStoreVersionExperiments := int32(56) // int32 | maximum number of related appStoreVersionExperiments returned (when they are included) (optional)
	limitAppStoreVersionExperimentsV2 := int32(56) // int32 | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) (optional)
	limitAppStoreVersionLocalizations := int32(56) // int32 | maximum number of related appStoreVersionLocalizations returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsGetInstance(context.Background(), id).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAlternativeDistributionPackages(fieldsAlternativeDistributionPackages).Include(include).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsGetInstance`: AppStoreVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
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
 **limitAppStoreVersionExperiments** | **int32** | maximum number of related appStoreVersionExperiments returned (when they are included) | 
 **limitAppStoreVersionExperimentsV2** | **int32** | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) | 
 **limitAppStoreVersionLocalizations** | **int32** | maximum number of related appStoreVersionLocalizations returned (when they are included) | 

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


## AppStoreVersionsRoutingAppCoverageGetToOneRelated

> RoutingAppCoverageWithoutIncludesResponse AppStoreVersionsRoutingAppCoverageGetToOneRelated(ctx, id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Execute()



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
	fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsRoutingAppCoverageGetToOneRelated(context.Background(), id).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsRoutingAppCoverageGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsRoutingAppCoverageGetToOneRelated`: RoutingAppCoverageWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsRoutingAppCoverageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsRoutingAppCoverageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 

### Return type

[**RoutingAppCoverageWithoutIncludesResponse**](RoutingAppCoverageWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppStoreVersionsUpdateInstance

> AppStoreVersionResponse AppStoreVersionsUpdateInstance(ctx, id).AppStoreVersionUpdateRequest(appStoreVersionUpdateRequest).Execute()



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
	appStoreVersionUpdateRequest := *openapiclient.NewAppStoreVersionUpdateRequest(*openapiclient.NewAppStoreVersionUpdateRequestData("Type_example", "Id_example")) // AppStoreVersionUpdateRequest | AppStoreVersion representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppStoreVersionsAPI.AppStoreVersionsUpdateInstance(context.Background(), id).AppStoreVersionUpdateRequest(appStoreVersionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppStoreVersionsAPI.AppStoreVersionsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppStoreVersionsUpdateInstance`: AppStoreVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `AppStoreVersionsAPI.AppStoreVersionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppStoreVersionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appStoreVersionUpdateRequest** | [**AppStoreVersionUpdateRequest**](AppStoreVersionUpdateRequest.md) | AppStoreVersion representation | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

