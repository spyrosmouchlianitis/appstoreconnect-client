# \GameCenterDetailsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterDetailsAchievementReleasesGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsAchievementReleasesGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/achievementReleases | 
[**GameCenterDetailsClassicMatchmakingRequestsGetMetrics**](GameCenterDetailsAPI.md#GameCenterDetailsClassicMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterDetails/{id}/metrics/classicMatchmakingRequests | 
[**GameCenterDetailsCreateInstance**](GameCenterDetailsAPI.md#GameCenterDetailsCreateInstance) | **Post** /v1/gameCenterDetails | 
[**GameCenterDetailsGameCenterAchievementsGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterAchievementsGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/gameCenterAchievements | 
[**GameCenterDetailsGameCenterAchievementsGetToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterAchievementsGetToManyRelationship) | **Get** /v1/gameCenterDetails/{id}/relationships/gameCenterAchievements | 
[**GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship) | **Patch** /v1/gameCenterDetails/{id}/relationships/gameCenterAchievements | 
[**GameCenterDetailsGameCenterAppVersionsGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterAppVersionsGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/gameCenterAppVersions | 
[**GameCenterDetailsGameCenterGroupGetToOneRelated**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterGroupGetToOneRelated) | **Get** /v1/gameCenterDetails/{id}/gameCenterGroup | 
[**GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/gameCenterLeaderboardSets | 
[**GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship) | **Get** /v1/gameCenterDetails/{id}/relationships/gameCenterLeaderboardSets | 
[**GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship) | **Patch** /v1/gameCenterDetails/{id}/relationships/gameCenterLeaderboardSets | 
[**GameCenterDetailsGameCenterLeaderboardsGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardsGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/gameCenterLeaderboards | 
[**GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship) | **Get** /v1/gameCenterDetails/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship**](GameCenterDetailsAPI.md#GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship) | **Patch** /v1/gameCenterDetails/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterDetailsGetInstance**](GameCenterDetailsAPI.md#GameCenterDetailsGetInstance) | **Get** /v1/gameCenterDetails/{id} | 
[**GameCenterDetailsLeaderboardReleasesGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsLeaderboardReleasesGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/leaderboardReleases | 
[**GameCenterDetailsLeaderboardSetReleasesGetToManyRelated**](GameCenterDetailsAPI.md#GameCenterDetailsLeaderboardSetReleasesGetToManyRelated) | **Get** /v1/gameCenterDetails/{id}/leaderboardSetReleases | 
[**GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics**](GameCenterDetailsAPI.md#GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterDetails/{id}/metrics/ruleBasedMatchmakingRequests | 
[**GameCenterDetailsUpdateInstance**](GameCenterDetailsAPI.md#GameCenterDetailsUpdateInstance) | **Patch** /v1/gameCenterDetails/{id} | 



## GameCenterDetailsAchievementReleasesGetToManyRelated

> GameCenterAchievementReleasesResponse GameCenterDetailsAchievementReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterAchievement(filterGameCenterAchievement).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()



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
	filterLive := []string{"Inner_example"} // []string | filter by attribute 'live' (optional)
	filterGameCenterAchievement := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterAchievement' (optional)
	fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsAchievementReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterAchievement(filterGameCenterAchievement).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsAchievementReleasesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsAchievementReleasesGetToManyRelated`: GameCenterAchievementReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsAchievementReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsAchievementReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterAchievement** | **[]string** | filter by id(s) of related &#39;gameCenterAchievement&#39; | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementReleasesResponse**](GameCenterAchievementReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsClassicMatchmakingRequestsGetMetrics

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsClassicMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()



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
	granularity := "P7D" // string | the granularity of the per-group dataset
	groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
	filterResult := "filterResult_example" // string | filter by 'result' attribute dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsClassicMatchmakingRequestsGetMetrics`: GameCenterMatchmakingAppRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsClassicMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingAppRequestsV1MetricResponse**](GameCenterMatchmakingAppRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsCreateInstance

> GameCenterDetailResponse GameCenterDetailsCreateInstance(ctx).GameCenterDetailCreateRequest(gameCenterDetailCreateRequest).Execute()



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
	gameCenterDetailCreateRequest := *openapiclient.NewGameCenterDetailCreateRequest(*openapiclient.NewGameCenterDetailCreateRequestData("Type_example", *openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // GameCenterDetailCreateRequest | GameCenterDetail representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsCreateInstance(context.Background()).GameCenterDetailCreateRequest(gameCenterDetailCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsCreateInstance`: GameCenterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterDetailCreateRequest** | [**GameCenterDetailCreateRequest**](GameCenterDetailCreateRequest.md) | GameCenterDetail representation | 

### Return type

[**GameCenterDetailResponse**](GameCenterDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterAchievementsGetToManyRelated

> GameCenterAchievementsResponse GameCenterDetailsGameCenterAchievementsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
	filterReferenceName := []string{"Inner_example"} // []string | filter by attribute 'referenceName' (optional)
	filterArchived := []string{"Inner_example"} // []string | filter by attribute 'archived' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterAchievementLocalizations := []string{"FieldsGameCenterAchievementLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterAchievementLocalizations (optional)
	fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
	limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterAchievementsGetToManyRelated`: GameCenterAchievementsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterAchievementsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterReferenceName** | **[]string** | filter by attribute &#39;referenceName&#39; | 
 **filterArchived** | **[]string** | filter by attribute &#39;archived&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

### Return type

[**GameCenterAchievementsResponse**](GameCenterAchievementsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterAchievementsGetToManyRelationship

> GameCenterDetailGameCenterAchievementsLinkagesResponse GameCenterDetailsGameCenterAchievementsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterAchievementsGetToManyRelationship`: GameCenterDetailGameCenterAchievementsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterAchievementsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterDetailGameCenterAchievementsLinkagesResponse**](GameCenterDetailGameCenterAchievementsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship

> GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship(ctx, id).GameCenterDetailGameCenterAchievementsLinkagesRequest(gameCenterDetailGameCenterAchievementsLinkagesRequest).Execute()



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
	gameCenterDetailGameCenterAchievementsLinkagesRequest := *openapiclient.NewGameCenterDetailGameCenterAchievementsLinkagesRequest([]openapiclient.GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData{*openapiclient.NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementData("Type_example", "Id_example")}) // GameCenterDetailGameCenterAchievementsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship(context.Background(), id).GameCenterDetailGameCenterAchievementsLinkagesRequest(gameCenterDetailGameCenterAchievementsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterAchievementsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterAchievementsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterDetailGameCenterAchievementsLinkagesRequest** | [**GameCenterDetailGameCenterAchievementsLinkagesRequest**](GameCenterDetailGameCenterAchievementsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterDetailsGameCenterAppVersionsGetToManyRelated

> GameCenterAppVersionsResponse GameCenterDetailsGameCenterAppVersionsGetToManyRelated(ctx, id).FilterEnabled(filterEnabled).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).Limit(limit).Include(include).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()



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
	filterEnabled := []string{"Inner_example"} // []string | filter by attribute 'enabled' (optional)
	fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
	fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitCompatibilityVersions := int32(56) // int32 | maximum number of related compatibilityVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterAppVersionsGetToManyRelated(context.Background(), id).FilterEnabled(filterEnabled).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsAppStoreVersions(fieldsAppStoreVersions).Limit(limit).Include(include).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterAppVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterAppVersionsGetToManyRelated`: GameCenterAppVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterAppVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterAppVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterEnabled** | **[]string** | filter by attribute &#39;enabled&#39; | 
 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitCompatibilityVersions** | **int32** | maximum number of related compatibilityVersions returned (when they are included) | 

### Return type

[**GameCenterAppVersionsResponse**](GameCenterAppVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterGroupGetToOneRelated

> GameCenterGroupResponse GameCenterDetailsGameCenterGroupGetToOneRelated(ctx, id).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterAchievements(limitGameCenterAchievements).Execute()



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
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitGameCenterDetails := int32(56) // int32 | maximum number of related gameCenterDetails returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitGameCenterAchievements := int32(56) // int32 | maximum number of related gameCenterAchievements returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterGroupGetToOneRelated(context.Background(), id).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterAchievements(limitGameCenterAchievements).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterGroupGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterGroupGetToOneRelated`: GameCenterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterGroupGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterGroupGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitGameCenterDetails** | **int32** | maximum number of related gameCenterDetails returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitGameCenterAchievements** | **int32** | maximum number of related gameCenterAchievements returned (when they are included) | 

### Return type

[**GameCenterGroupResponse**](GameCenterGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated

> GameCenterLeaderboardSetsResponse GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Execute()



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
	filterReferenceName := []string{"Inner_example"} // []string | filter by attribute 'referenceName' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterLeaderboardSetLocalizations := []string{"FieldsGameCenterLeaderboardSetLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
	limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated`: GameCenterLeaderboardSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardSetsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterReferenceName** | **[]string** | filter by attribute &#39;referenceName&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

### Return type

[**GameCenterLeaderboardSetsResponse**](GameCenterLeaderboardSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship

> GameCenterDetailGameCenterLeaderboardSetsLinkagesResponse GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship`: GameCenterDetailGameCenterLeaderboardSetsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardSetsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterDetailGameCenterLeaderboardSetsLinkagesResponse**](GameCenterDetailGameCenterLeaderboardSetsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship

> GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship(ctx, id).GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest(gameCenterDetailGameCenterLeaderboardSetsLinkagesRequest).Execute()



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
	gameCenterDetailGameCenterLeaderboardSetsLinkagesRequest := *openapiclient.NewGameCenterDetailGameCenterLeaderboardSetsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example")}) // GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship(context.Background(), id).GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest(gameCenterDetailGameCenterLeaderboardSetsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardSetsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterDetailGameCenterLeaderboardSetsLinkagesRequest** | [**GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest**](GameCenterDetailGameCenterLeaderboardSetsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterDetailsGameCenterLeaderboardsGetToManyRelated

> GameCenterLeaderboardsResponse GameCenterDetailsGameCenterLeaderboardsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
	filterReferenceName := []string{"Inner_example"} // []string | filter by attribute 'referenceName' (optional)
	filterArchived := []string{"Inner_example"} // []string | filter by attribute 'archived' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterLeaderboardLocalizations := []string{"FieldsGameCenterLeaderboardLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardLocalizations (optional)
	fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
	limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterLeaderboardsGetToManyRelated`: GameCenterLeaderboardsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterReferenceName** | **[]string** | filter by attribute &#39;referenceName&#39; | 
 **filterArchived** | **[]string** | filter by attribute &#39;archived&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

### Return type

[**GameCenterLeaderboardsResponse**](GameCenterLeaderboardsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship

> GameCenterDetailGameCenterLeaderboardsLinkagesResponse GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship`: GameCenterDetailGameCenterLeaderboardsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterDetailGameCenterLeaderboardsLinkagesResponse**](GameCenterDetailGameCenterLeaderboardsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship

> GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship(ctx, id).GameCenterDetailGameCenterLeaderboardsLinkagesRequest(gameCenterDetailGameCenterLeaderboardsLinkagesRequest).Execute()



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
	gameCenterDetailGameCenterLeaderboardsLinkagesRequest := *openapiclient.NewGameCenterDetailGameCenterLeaderboardsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")}) // GameCenterDetailGameCenterLeaderboardsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship(context.Background(), id).GameCenterDetailGameCenterLeaderboardsLinkagesRequest(gameCenterDetailGameCenterLeaderboardsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterDetailsGameCenterLeaderboardsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterDetailGameCenterLeaderboardsLinkagesRequest** | [**GameCenterDetailGameCenterLeaderboardsLinkagesRequest**](GameCenterDetailGameCenterLeaderboardsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterDetailsGetInstance

> GameCenterDetailResponse GameCenterDetailsGetInstance(ctx, id).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Include(include).LimitAchievementReleases(limitAchievementReleases).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterAppVersions(limitGameCenterAppVersions).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitLeaderboardReleases(limitLeaderboardReleases).LimitLeaderboardSetReleases(limitLeaderboardSetReleases).Execute()



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
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
	fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
	fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitAchievementReleases := int32(56) // int32 | maximum number of related achievementReleases returned (when they are included) (optional)
	limitGameCenterAchievements := int32(56) // int32 | maximum number of related gameCenterAchievements returned (when they are included) (optional)
	limitGameCenterAppVersions := int32(56) // int32 | maximum number of related gameCenterAppVersions returned (when they are included) (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
	limitLeaderboardReleases := int32(56) // int32 | maximum number of related leaderboardReleases returned (when they are included) (optional)
	limitLeaderboardSetReleases := int32(56) // int32 | maximum number of related leaderboardSetReleases returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsGetInstance(context.Background(), id).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Include(include).LimitAchievementReleases(limitAchievementReleases).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterAppVersions(limitGameCenterAppVersions).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitLeaderboardReleases(limitLeaderboardReleases).LimitLeaderboardSetReleases(limitLeaderboardSetReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsGetInstance`: GameCenterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitAchievementReleases** | **int32** | maximum number of related achievementReleases returned (when they are included) | 
 **limitGameCenterAchievements** | **int32** | maximum number of related gameCenterAchievements returned (when they are included) | 
 **limitGameCenterAppVersions** | **int32** | maximum number of related gameCenterAppVersions returned (when they are included) | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitLeaderboardReleases** | **int32** | maximum number of related leaderboardReleases returned (when they are included) | 
 **limitLeaderboardSetReleases** | **int32** | maximum number of related leaderboardSetReleases returned (when they are included) | 

### Return type

[**GameCenterDetailResponse**](GameCenterDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsLeaderboardReleasesGetToManyRelated

> GameCenterLeaderboardReleasesResponse GameCenterDetailsLeaderboardReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterLeaderboard(filterGameCenterLeaderboard).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Limit(limit).Include(include).Execute()



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
	filterLive := []string{"Inner_example"} // []string | filter by attribute 'live' (optional)
	filterGameCenterLeaderboard := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterLeaderboard' (optional)
	fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsLeaderboardReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterLeaderboard(filterGameCenterLeaderboard).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsLeaderboardReleasesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsLeaderboardReleasesGetToManyRelated`: GameCenterLeaderboardReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsLeaderboardReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsLeaderboardReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterLeaderboard** | **[]string** | filter by id(s) of related &#39;gameCenterLeaderboard&#39; | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardReleasesResponse**](GameCenterLeaderboardReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsLeaderboardSetReleasesGetToManyRelated

> GameCenterLeaderboardSetReleasesResponse GameCenterDetailsLeaderboardSetReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterLeaderboardSet(filterGameCenterLeaderboardSet).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).Limit(limit).Include(include).Execute()



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
	filterLive := []string{"Inner_example"} // []string | filter by attribute 'live' (optional)
	filterGameCenterLeaderboardSet := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterLeaderboardSet' (optional)
	fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsLeaderboardSetReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterLeaderboardSet(filterGameCenterLeaderboardSet).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsLeaderboardSetReleasesGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsLeaderboardSetReleasesGetToManyRelated`: GameCenterLeaderboardSetReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsLeaderboardSetReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsLeaderboardSetReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterLeaderboardSet** | **[]string** | filter by id(s) of related &#39;gameCenterLeaderboardSet&#39; | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetReleasesResponse**](GameCenterLeaderboardSetReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()



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
	granularity := "P7D" // string | the granularity of the per-group dataset
	groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
	filterResult := "filterResult_example" // string | filter by 'result' attribute dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics`: GameCenterMatchmakingAppRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsRuleBasedMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingAppRequestsV1MetricResponse**](GameCenterMatchmakingAppRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsUpdateInstance

> GameCenterDetailResponse GameCenterDetailsUpdateInstance(ctx, id).GameCenterDetailUpdateRequest(gameCenterDetailUpdateRequest).Execute()



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
	gameCenterDetailUpdateRequest := *openapiclient.NewGameCenterDetailUpdateRequest(*openapiclient.NewGameCenterDetailUpdateRequestData("Type_example", "Id_example")) // GameCenterDetailUpdateRequest | GameCenterDetail representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterDetailsAPI.GameCenterDetailsUpdateInstance(context.Background(), id).GameCenterDetailUpdateRequest(gameCenterDetailUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterDetailsAPI.GameCenterDetailsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsUpdateInstance`: GameCenterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterDetailsAPI.GameCenterDetailsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterDetailUpdateRequest** | [**GameCenterDetailUpdateRequest**](GameCenterDetailUpdateRequest.md) | GameCenterDetail representation | 

### Return type

[**GameCenterDetailResponse**](GameCenterDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

