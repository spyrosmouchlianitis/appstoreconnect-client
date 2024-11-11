# \CiProductsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CiProductsAdditionalRepositoriesGetToManyRelated**](CiProductsAPI.md#CiProductsAdditionalRepositoriesGetToManyRelated) | **Get** /v1/ciProducts/{id}/additionalRepositories | 
[**CiProductsAppGetToOneRelated**](CiProductsAPI.md#CiProductsAppGetToOneRelated) | **Get** /v1/ciProducts/{id}/app | 
[**CiProductsBuildRunsGetToManyRelated**](CiProductsAPI.md#CiProductsBuildRunsGetToManyRelated) | **Get** /v1/ciProducts/{id}/buildRuns | 
[**CiProductsDeleteInstance**](CiProductsAPI.md#CiProductsDeleteInstance) | **Delete** /v1/ciProducts/{id} | 
[**CiProductsGetCollection**](CiProductsAPI.md#CiProductsGetCollection) | **Get** /v1/ciProducts | 
[**CiProductsGetInstance**](CiProductsAPI.md#CiProductsGetInstance) | **Get** /v1/ciProducts/{id} | 
[**CiProductsPrimaryRepositoriesGetToManyRelated**](CiProductsAPI.md#CiProductsPrimaryRepositoriesGetToManyRelated) | **Get** /v1/ciProducts/{id}/primaryRepositories | 
[**CiProductsWorkflowsGetToManyRelated**](CiProductsAPI.md#CiProductsWorkflowsGetToManyRelated) | **Get** /v1/ciProducts/{id}/workflows | 



## CiProductsAdditionalRepositoriesGetToManyRelated

> ScmRepositoriesResponse CiProductsAdditionalRepositoriesGetToManyRelated(ctx, id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()



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
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	fieldsScmProviders := []string{"FieldsScmProviders_example"} // []string | the fields to include for returned resources of type scmProviders (optional)
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsAdditionalRepositoriesGetToManyRelated(context.Background(), id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsAdditionalRepositoriesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsAdditionalRepositoriesGetToManyRelated`: ScmRepositoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsAdditionalRepositoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsAdditionalRepositoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterId** | **[]string** | filter by id(s) | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **fieldsScmProviders** | **[]string** | the fields to include for returned resources of type scmProviders | 
 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmRepositoriesResponse**](ScmRepositoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsAppGetToOneRelated

> AppResponse CiProductsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsCiProducts(fieldsCiProducts).FieldsBetaGroups(fieldsBetaGroups).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBuilds(fieldsBuilds).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsAppInfos(fieldsAppInfos).FieldsAppClips(fieldsAppClips).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).FieldsAppPreOrders(fieldsAppPreOrders).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsAppEvents(fieldsAppEvents).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsSubscriptionGracePeriods(fieldsSubscriptionGracePeriods).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).Include(include).LimitAppEncryptionDeclarations(limitAppEncryptionDeclarations).LimitBetaGroups(limitBetaGroups).LimitAppStoreVersions(limitAppStoreVersions).LimitPreReleaseVersions(limitPreReleaseVersions).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBuilds(limitBuilds).LimitAppInfos(limitAppInfos).LimitAppClips(limitAppClips).LimitInAppPurchases(limitInAppPurchases).LimitSubscriptionGroups(limitSubscriptionGroups).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitAppCustomProductPages(limitAppCustomProductPages).LimitInAppPurchasesV2(limitInAppPurchasesV2).LimitPromotedPurchases(limitPromotedPurchases).LimitAppEvents(limitAppEvents).LimitReviewSubmissions(limitReviewSubmissions).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Execute()



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
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	fieldsPreReleaseVersions := []string{"FieldsPreReleaseVersions_example"} // []string | the fields to include for returned resources of type preReleaseVersions (optional)
	fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsBetaLicenseAgreements := []string{"FieldsBetaLicenseAgreements_example"} // []string | the fields to include for returned resources of type betaLicenseAgreements (optional)
	fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)
	fieldsAppInfos := []string{"FieldsAppInfos_example"} // []string | the fields to include for returned resources of type appInfos (optional)
	fieldsAppClips := []string{"FieldsAppClips_example"} // []string | the fields to include for returned resources of type appClips (optional)
	fieldsEndUserLicenseAgreements := []string{"FieldsEndUserLicenseAgreements_example"} // []string | the fields to include for returned resources of type endUserLicenseAgreements (optional)
	fieldsAppPreOrders := []string{"FieldsAppPreOrders_example"} // []string | the fields to include for returned resources of type appPreOrders (optional)
	fieldsInAppPurchases := []string{"FieldsInAppPurchases_example"} // []string | the fields to include for returned resources of type inAppPurchases (optional)
	fieldsSubscriptionGroups := []string{"FieldsSubscriptionGroups_example"} // []string | the fields to include for returned resources of type subscriptionGroups (optional)
	fieldsGameCenterEnabledVersions := []string{"FieldsGameCenterEnabledVersions_example"} // []string | the fields to include for returned resources of type gameCenterEnabledVersions (optional)
	fieldsAppCustomProductPages := []string{"FieldsAppCustomProductPages_example"} // []string | the fields to include for returned resources of type appCustomProductPages (optional)
	fieldsPromotedPurchases := []string{"FieldsPromotedPurchases_example"} // []string | the fields to include for returned resources of type promotedPurchases (optional)
	fieldsAppEvents := []string{"FieldsAppEvents_example"} // []string | the fields to include for returned resources of type appEvents (optional)
	fieldsReviewSubmissions := []string{"FieldsReviewSubmissions_example"} // []string | the fields to include for returned resources of type reviewSubmissions (optional)
	fieldsSubscriptionGracePeriods := []string{"FieldsSubscriptionGracePeriods_example"} // []string | the fields to include for returned resources of type subscriptionGracePeriods (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAppEncryptionDeclarations := int32(56) // int32 | maximum number of related appEncryptionDeclarations returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitAppStoreVersions := int32(56) // int32 | maximum number of related appStoreVersions returned (when they are included) (optional)
	limitPreReleaseVersions := int32(56) // int32 | maximum number of related preReleaseVersions returned (when they are included) (optional)
	limitBetaAppLocalizations := int32(56) // int32 | maximum number of related betaAppLocalizations returned (when they are included) (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)
	limitAppInfos := int32(56) // int32 | maximum number of related appInfos returned (when they are included) (optional)
	limitAppClips := int32(56) // int32 | maximum number of related appClips returned (when they are included) (optional)
	limitInAppPurchases := int32(56) // int32 | maximum number of related inAppPurchases returned (when they are included) (optional)
	limitSubscriptionGroups := int32(56) // int32 | maximum number of related subscriptionGroups returned (when they are included) (optional)
	limitGameCenterEnabledVersions := int32(56) // int32 | maximum number of related gameCenterEnabledVersions returned (when they are included) (optional)
	limitAppCustomProductPages := int32(56) // int32 | maximum number of related appCustomProductPages returned (when they are included) (optional)
	limitInAppPurchasesV2 := int32(56) // int32 | maximum number of related inAppPurchasesV2 returned (when they are included) (optional)
	limitPromotedPurchases := int32(56) // int32 | maximum number of related promotedPurchases returned (when they are included) (optional)
	limitAppEvents := int32(56) // int32 | maximum number of related appEvents returned (when they are included) (optional)
	limitReviewSubmissions := int32(56) // int32 | maximum number of related reviewSubmissions returned (when they are included) (optional)
	limitAppStoreVersionExperimentsV2 := int32(56) // int32 | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).FieldsCiProducts(fieldsCiProducts).FieldsBetaGroups(fieldsBetaGroups).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsPreReleaseVersions(fieldsPreReleaseVersions).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).FieldsBuilds(fieldsBuilds).FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).FieldsAppInfos(fieldsAppInfos).FieldsAppClips(fieldsAppClips).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).FieldsAppPreOrders(fieldsAppPreOrders).FieldsInAppPurchases(fieldsInAppPurchases).FieldsSubscriptionGroups(fieldsSubscriptionGroups).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsAppCustomProductPages(fieldsAppCustomProductPages).FieldsPromotedPurchases(fieldsPromotedPurchases).FieldsAppEvents(fieldsAppEvents).FieldsReviewSubmissions(fieldsReviewSubmissions).FieldsSubscriptionGracePeriods(fieldsSubscriptionGracePeriods).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).Include(include).LimitAppEncryptionDeclarations(limitAppEncryptionDeclarations).LimitBetaGroups(limitBetaGroups).LimitAppStoreVersions(limitAppStoreVersions).LimitPreReleaseVersions(limitPreReleaseVersions).LimitBetaAppLocalizations(limitBetaAppLocalizations).LimitBuilds(limitBuilds).LimitAppInfos(limitAppInfos).LimitAppClips(limitAppClips).LimitInAppPurchases(limitInAppPurchases).LimitSubscriptionGroups(limitSubscriptionGroups).LimitGameCenterEnabledVersions(limitGameCenterEnabledVersions).LimitAppCustomProductPages(limitAppCustomProductPages).LimitInAppPurchasesV2(limitInAppPurchasesV2).LimitPromotedPurchases(limitPromotedPurchases).LimitAppEvents(limitAppEvents).LimitReviewSubmissions(limitReviewSubmissions).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsAppGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsAppGetToOneRelated`: AppResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsAppGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsPreReleaseVersions** | **[]string** | the fields to include for returned resources of type preReleaseVersions | 
 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaLicenseAgreements** | **[]string** | the fields to include for returned resources of type betaLicenseAgreements | 
 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 
 **fieldsAppInfos** | **[]string** | the fields to include for returned resources of type appInfos | 
 **fieldsAppClips** | **[]string** | the fields to include for returned resources of type appClips | 
 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 
 **fieldsAppPreOrders** | **[]string** | the fields to include for returned resources of type appPreOrders | 
 **fieldsInAppPurchases** | **[]string** | the fields to include for returned resources of type inAppPurchases | 
 **fieldsSubscriptionGroups** | **[]string** | the fields to include for returned resources of type subscriptionGroups | 
 **fieldsGameCenterEnabledVersions** | **[]string** | the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsAppCustomProductPages** | **[]string** | the fields to include for returned resources of type appCustomProductPages | 
 **fieldsPromotedPurchases** | **[]string** | the fields to include for returned resources of type promotedPurchases | 
 **fieldsAppEvents** | **[]string** | the fields to include for returned resources of type appEvents | 
 **fieldsReviewSubmissions** | **[]string** | the fields to include for returned resources of type reviewSubmissions | 
 **fieldsSubscriptionGracePeriods** | **[]string** | the fields to include for returned resources of type subscriptionGracePeriods | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAppEncryptionDeclarations** | **int32** | maximum number of related appEncryptionDeclarations returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitAppStoreVersions** | **int32** | maximum number of related appStoreVersions returned (when they are included) | 
 **limitPreReleaseVersions** | **int32** | maximum number of related preReleaseVersions returned (when they are included) | 
 **limitBetaAppLocalizations** | **int32** | maximum number of related betaAppLocalizations returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 
 **limitAppInfos** | **int32** | maximum number of related appInfos returned (when they are included) | 
 **limitAppClips** | **int32** | maximum number of related appClips returned (when they are included) | 
 **limitInAppPurchases** | **int32** | maximum number of related inAppPurchases returned (when they are included) | 
 **limitSubscriptionGroups** | **int32** | maximum number of related subscriptionGroups returned (when they are included) | 
 **limitGameCenterEnabledVersions** | **int32** | maximum number of related gameCenterEnabledVersions returned (when they are included) | 
 **limitAppCustomProductPages** | **int32** | maximum number of related appCustomProductPages returned (when they are included) | 
 **limitInAppPurchasesV2** | **int32** | maximum number of related inAppPurchasesV2 returned (when they are included) | 
 **limitPromotedPurchases** | **int32** | maximum number of related promotedPurchases returned (when they are included) | 
 **limitAppEvents** | **int32** | maximum number of related appEvents returned (when they are included) | 
 **limitReviewSubmissions** | **int32** | maximum number of related reviewSubmissions returned (when they are included) | 
 **limitAppStoreVersionExperimentsV2** | **int32** | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) | 

### Return type

[**AppResponse**](AppResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsBuildRunsGetToManyRelated

> CiBuildRunsResponse CiProductsBuildRunsGetToManyRelated(ctx, id).FilterBuilds(filterBuilds).Sort(sort).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmPullRequests(fieldsScmPullRequests).Limit(limit).Include(include).LimitBuilds(limitBuilds).Execute()



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
	filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsCiBuildRuns := []string{"FieldsCiBuildRuns_example"} // []string | the fields to include for returned resources of type ciBuildRuns (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsCiWorkflows := []string{"FieldsCiWorkflows_example"} // []string | the fields to include for returned resources of type ciWorkflows (optional)
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	fieldsScmPullRequests := []string{"FieldsScmPullRequests_example"} // []string | the fields to include for returned resources of type scmPullRequests (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsBuildRunsGetToManyRelated(context.Background(), id).FilterBuilds(filterBuilds).Sort(sort).FieldsCiBuildRuns(fieldsCiBuildRuns).FieldsBuilds(fieldsBuilds).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmGitReferences(fieldsScmGitReferences).FieldsScmPullRequests(fieldsScmPullRequests).Limit(limit).Include(include).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsBuildRunsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsBuildRunsGetToManyRelated`: CiBuildRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsBuildRunsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsBuildRunsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsCiBuildRuns** | **[]string** | the fields to include for returned resources of type ciBuildRuns | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsCiWorkflows** | **[]string** | the fields to include for returned resources of type ciWorkflows | 
 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **fieldsScmPullRequests** | **[]string** | the fields to include for returned resources of type scmPullRequests | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**CiBuildRunsResponse**](CiBuildRunsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsDeleteInstance

> CiProductsDeleteInstance(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CiProductsAPI.CiProductsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCiProductsDeleteInstanceRequest struct via the builder pattern


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


## CiProductsGetCollection

> CiProductsResponse CiProductsGetCollection(ctx).FilterProductType(filterProductType).FilterApp(filterApp).FieldsCiProducts(fieldsCiProducts).FieldsApps(fieldsApps).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).LimitPrimaryRepositories(limitPrimaryRepositories).Execute()



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
	filterProductType := []string{"FilterProductType_example"} // []string | filter by attribute 'productType' (optional)
	filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPrimaryRepositories := int32(56) // int32 | maximum number of related primaryRepositories returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsGetCollection(context.Background()).FilterProductType(filterProductType).FilterApp(filterApp).FieldsCiProducts(fieldsCiProducts).FieldsApps(fieldsApps).FieldsScmRepositories(fieldsScmRepositories).Limit(limit).Include(include).LimitPrimaryRepositories(limitPrimaryRepositories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsGetCollection`: CiProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterProductType** | **[]string** | filter by attribute &#39;productType&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPrimaryRepositories** | **int32** | maximum number of related primaryRepositories returned (when they are included) | 

### Return type

[**CiProductsResponse**](CiProductsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsGetInstance

> CiProductResponse CiProductsGetInstance(ctx, id).FieldsCiProducts(fieldsCiProducts).FieldsApps(fieldsApps).FieldsScmRepositories(fieldsScmRepositories).Include(include).LimitPrimaryRepositories(limitPrimaryRepositories).Execute()



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
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitPrimaryRepositories := int32(56) // int32 | maximum number of related primaryRepositories returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsGetInstance(context.Background(), id).FieldsCiProducts(fieldsCiProducts).FieldsApps(fieldsApps).FieldsScmRepositories(fieldsScmRepositories).Include(include).LimitPrimaryRepositories(limitPrimaryRepositories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsGetInstance`: CiProductResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitPrimaryRepositories** | **int32** | maximum number of related primaryRepositories returned (when they are included) | 

### Return type

[**CiProductResponse**](CiProductResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsPrimaryRepositoriesGetToManyRelated

> ScmRepositoriesResponse CiProductsPrimaryRepositoriesGetToManyRelated(ctx, id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()



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
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	fieldsScmProviders := []string{"FieldsScmProviders_example"} // []string | the fields to include for returned resources of type scmProviders (optional)
	fieldsScmGitReferences := []string{"FieldsScmGitReferences_example"} // []string | the fields to include for returned resources of type scmGitReferences (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsPrimaryRepositoriesGetToManyRelated(context.Background(), id).FilterId(filterId).FieldsScmRepositories(fieldsScmRepositories).FieldsScmProviders(fieldsScmProviders).FieldsScmGitReferences(fieldsScmGitReferences).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsPrimaryRepositoriesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsPrimaryRepositoriesGetToManyRelated`: ScmRepositoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsPrimaryRepositoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsPrimaryRepositoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterId** | **[]string** | filter by id(s) | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **fieldsScmProviders** | **[]string** | the fields to include for returned resources of type scmProviders | 
 **fieldsScmGitReferences** | **[]string** | the fields to include for returned resources of type scmGitReferences | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**ScmRepositoriesResponse**](ScmRepositoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CiProductsWorkflowsGetToManyRelated

> CiWorkflowsResponse CiProductsWorkflowsGetToManyRelated(ctx, id).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmRepositories(fieldsScmRepositories).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).Execute()



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
	fieldsCiWorkflows := []string{"FieldsCiWorkflows_example"} // []string | the fields to include for returned resources of type ciWorkflows (optional)
	fieldsCiProducts := []string{"FieldsCiProducts_example"} // []string | the fields to include for returned resources of type ciProducts (optional)
	fieldsScmRepositories := []string{"FieldsScmRepositories_example"} // []string | the fields to include for returned resources of type scmRepositories (optional)
	fieldsCiXcodeVersions := []string{"FieldsCiXcodeVersions_example"} // []string | the fields to include for returned resources of type ciXcodeVersions (optional)
	fieldsCiMacOsVersions := []string{"FieldsCiMacOsVersions_example"} // []string | the fields to include for returned resources of type ciMacOsVersions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiProductsAPI.CiProductsWorkflowsGetToManyRelated(context.Background(), id).FieldsCiWorkflows(fieldsCiWorkflows).FieldsCiProducts(fieldsCiProducts).FieldsScmRepositories(fieldsScmRepositories).FieldsCiXcodeVersions(fieldsCiXcodeVersions).FieldsCiMacOsVersions(fieldsCiMacOsVersions).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiProductsAPI.CiProductsWorkflowsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CiProductsWorkflowsGetToManyRelated`: CiWorkflowsResponse
	fmt.Fprintf(os.Stdout, "Response from `CiProductsAPI.CiProductsWorkflowsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCiProductsWorkflowsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCiWorkflows** | **[]string** | the fields to include for returned resources of type ciWorkflows | 
 **fieldsCiProducts** | **[]string** | the fields to include for returned resources of type ciProducts | 
 **fieldsScmRepositories** | **[]string** | the fields to include for returned resources of type scmRepositories | 
 **fieldsCiXcodeVersions** | **[]string** | the fields to include for returned resources of type ciXcodeVersions | 
 **fieldsCiMacOsVersions** | **[]string** | the fields to include for returned resources of type ciMacOsVersions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**CiWorkflowsResponse**](CiWorkflowsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

