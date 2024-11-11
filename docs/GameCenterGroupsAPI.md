# \GameCenterGroupsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterGroupsCreateInstance**](GameCenterGroupsAPI.md#GameCenterGroupsCreateInstance) | **Post** /v1/gameCenterGroups | 
[**GameCenterGroupsDeleteInstance**](GameCenterGroupsAPI.md#GameCenterGroupsDeleteInstance) | **Delete** /v1/gameCenterGroups/{id} | 
[**GameCenterGroupsGameCenterAchievementsGetToManyRelated**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterAchievementsGetToManyRelated) | **Get** /v1/gameCenterGroups/{id}/gameCenterAchievements | 
[**GameCenterGroupsGameCenterAchievementsGetToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterAchievementsGetToManyRelationship) | **Get** /v1/gameCenterGroups/{id}/relationships/gameCenterAchievements | 
[**GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship) | **Patch** /v1/gameCenterGroups/{id}/relationships/gameCenterAchievements | 
[**GameCenterGroupsGameCenterDetailsGetToManyRelated**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterDetailsGetToManyRelated) | **Get** /v1/gameCenterGroups/{id}/gameCenterDetails | 
[**GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated) | **Get** /v1/gameCenterGroups/{id}/gameCenterLeaderboardSets | 
[**GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship) | **Get** /v1/gameCenterGroups/{id}/relationships/gameCenterLeaderboardSets | 
[**GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship) | **Patch** /v1/gameCenterGroups/{id}/relationships/gameCenterLeaderboardSets | 
[**GameCenterGroupsGameCenterLeaderboardsGetToManyRelated**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardsGetToManyRelated) | **Get** /v1/gameCenterGroups/{id}/gameCenterLeaderboards | 
[**GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship) | **Get** /v1/gameCenterGroups/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship**](GameCenterGroupsAPI.md#GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship) | **Patch** /v1/gameCenterGroups/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterGroupsGetCollection**](GameCenterGroupsAPI.md#GameCenterGroupsGetCollection) | **Get** /v1/gameCenterGroups | 
[**GameCenterGroupsGetInstance**](GameCenterGroupsAPI.md#GameCenterGroupsGetInstance) | **Get** /v1/gameCenterGroups/{id} | 
[**GameCenterGroupsUpdateInstance**](GameCenterGroupsAPI.md#GameCenterGroupsUpdateInstance) | **Patch** /v1/gameCenterGroups/{id} | 



## GameCenterGroupsCreateInstance

> GameCenterGroupResponse GameCenterGroupsCreateInstance(ctx).GameCenterGroupCreateRequest(gameCenterGroupCreateRequest).Execute()



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
	gameCenterGroupCreateRequest := *openapiclient.NewGameCenterGroupCreateRequest(*openapiclient.NewGameCenterGroupCreateRequestData("Type_example")) // GameCenterGroupCreateRequest | GameCenterGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsCreateInstance(context.Background()).GameCenterGroupCreateRequest(gameCenterGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsCreateInstance`: GameCenterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterGroupCreateRequest** | [**GameCenterGroupCreateRequest**](GameCenterGroupCreateRequest.md) | GameCenterGroup representation | 

### Return type

[**GameCenterGroupResponse**](GameCenterGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsDeleteInstance

> GameCenterGroupsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterGroupsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterGroupsGameCenterAchievementsGetToManyRelated

> GameCenterAchievementsResponse GameCenterGroupsGameCenterAchievementsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterAchievementsGetToManyRelated`: GameCenterAchievementsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterAchievementsGetToManyRelatedRequest struct via the builder pattern


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


## GameCenterGroupsGameCenterAchievementsGetToManyRelationship

> GameCenterGroupGameCenterAchievementsLinkagesResponse GameCenterGroupsGameCenterAchievementsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterAchievementsGetToManyRelationship`: GameCenterGroupGameCenterAchievementsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterAchievementsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterGroupGameCenterAchievementsLinkagesResponse**](GameCenterGroupGameCenterAchievementsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship

> GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship(ctx, id).GameCenterGroupGameCenterAchievementsLinkagesRequest(gameCenterGroupGameCenterAchievementsLinkagesRequest).Execute()



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
	gameCenterGroupGameCenterAchievementsLinkagesRequest := *openapiclient.NewGameCenterGroupGameCenterAchievementsLinkagesRequest([]openapiclient.GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData{*openapiclient.NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementData("Type_example", "Id_example")}) // GameCenterGroupGameCenterAchievementsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship(context.Background(), id).GameCenterGroupGameCenterAchievementsLinkagesRequest(gameCenterGroupGameCenterAchievementsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterAchievementsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterAchievementsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterGroupGameCenterAchievementsLinkagesRequest** | [**GameCenterGroupGameCenterAchievementsLinkagesRequest**](GameCenterGroupGameCenterAchievementsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterGroupsGameCenterDetailsGetToManyRelated

> GameCenterDetailsResponse GameCenterGroupsGameCenterDetailsGetToManyRelated(ctx, id).FilterGameCenterAppVersionsEnabled(filterGameCenterAppVersionsEnabled).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsApps(fieldsApps).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitGameCenterAppVersions(limitGameCenterAppVersions).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterAchievements(limitGameCenterAchievements).LimitAchievementReleases(limitAchievementReleases).LimitLeaderboardReleases(limitLeaderboardReleases).LimitLeaderboardSetReleases(limitLeaderboardSetReleases).Execute()



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
	filterGameCenterAppVersionsEnabled := []string{"Inner_example"} // []string | filter by attribute 'gameCenterAppVersions.enabled' (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
	fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
	fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitGameCenterAppVersions := int32(56) // int32 | maximum number of related gameCenterAppVersions returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitGameCenterAchievements := int32(56) // int32 | maximum number of related gameCenterAchievements returned (when they are included) (optional)
	limitAchievementReleases := int32(56) // int32 | maximum number of related achievementReleases returned (when they are included) (optional)
	limitLeaderboardReleases := int32(56) // int32 | maximum number of related leaderboardReleases returned (when they are included) (optional)
	limitLeaderboardSetReleases := int32(56) // int32 | maximum number of related leaderboardSetReleases returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterDetailsGetToManyRelated(context.Background(), id).FilterGameCenterAppVersionsEnabled(filterGameCenterAppVersionsEnabled).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsApps(fieldsApps).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitGameCenterAppVersions(limitGameCenterAppVersions).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterAchievements(limitGameCenterAchievements).LimitAchievementReleases(limitAchievementReleases).LimitLeaderboardReleases(limitLeaderboardReleases).LimitLeaderboardSetReleases(limitLeaderboardSetReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterDetailsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterDetailsGetToManyRelated`: GameCenterDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterDetailsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterDetailsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterGameCenterAppVersionsEnabled** | **[]string** | filter by attribute &#39;gameCenterAppVersions.enabled&#39; | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitGameCenterAppVersions** | **int32** | maximum number of related gameCenterAppVersions returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitGameCenterAchievements** | **int32** | maximum number of related gameCenterAchievements returned (when they are included) | 
 **limitAchievementReleases** | **int32** | maximum number of related achievementReleases returned (when they are included) | 
 **limitLeaderboardReleases** | **int32** | maximum number of related leaderboardReleases returned (when they are included) | 
 **limitLeaderboardSetReleases** | **int32** | maximum number of related leaderboardSetReleases returned (when they are included) | 

### Return type

[**GameCenterDetailsResponse**](GameCenterDetailsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated

> GameCenterLeaderboardSetsResponse GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Limit(limit).Include(include).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated`: GameCenterLeaderboardSetsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardSetsGetToManyRelatedRequest struct via the builder pattern


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


## GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship

> GameCenterGroupGameCenterLeaderboardSetsLinkagesResponse GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship`: GameCenterGroupGameCenterLeaderboardSetsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardSetsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterGroupGameCenterLeaderboardSetsLinkagesResponse**](GameCenterGroupGameCenterLeaderboardSetsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship

> GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship(ctx, id).GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest(gameCenterGroupGameCenterLeaderboardSetsLinkagesRequest).Execute()



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
	gameCenterGroupGameCenterLeaderboardSetsLinkagesRequest := *openapiclient.NewGameCenterGroupGameCenterLeaderboardSetsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example")}) // GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship(context.Background(), id).GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest(gameCenterGroupGameCenterLeaderboardSetsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardSetsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterGroupGameCenterLeaderboardSetsLinkagesRequest** | [**GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest**](GameCenterGroupGameCenterLeaderboardSetsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterGroupsGameCenterLeaderboardsGetToManyRelated

> GameCenterLeaderboardsResponse GameCenterGroupsGameCenterLeaderboardsGetToManyRelated(ctx, id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelated(context.Background(), id).FilterReferenceName(filterReferenceName).FilterArchived(filterArchived).FilterId(filterId).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterLeaderboardsGetToManyRelated`: GameCenterLeaderboardsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardsGetToManyRelatedRequest struct via the builder pattern


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


## GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship

> GameCenterGroupGameCenterLeaderboardsLinkagesResponse GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship`: GameCenterGroupGameCenterLeaderboardsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterGroupGameCenterLeaderboardsLinkagesResponse**](GameCenterGroupGameCenterLeaderboardsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship

> GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship(ctx, id).GameCenterGroupGameCenterLeaderboardsLinkagesRequest(gameCenterGroupGameCenterLeaderboardsLinkagesRequest).Execute()



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
	gameCenterGroupGameCenterLeaderboardsLinkagesRequest := *openapiclient.NewGameCenterGroupGameCenterLeaderboardsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")}) // GameCenterGroupGameCenterLeaderboardsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship(context.Background(), id).GameCenterGroupGameCenterLeaderboardsLinkagesRequest(gameCenterGroupGameCenterLeaderboardsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterGroupsGameCenterLeaderboardsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterGroupGameCenterLeaderboardsLinkagesRequest** | [**GameCenterGroupGameCenterLeaderboardsLinkagesRequest**](GameCenterGroupGameCenterLeaderboardsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterGroupsGetCollection

> GameCenterGroupsResponse GameCenterGroupsGetCollection(ctx).FilterGameCenterDetails(filterGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).Execute()



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
	filterGameCenterDetails := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterDetails' (optional)
	fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
	fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
	fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
	fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
	fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitGameCenterAchievements := int32(56) // int32 | maximum number of related gameCenterAchievements returned (when they are included) (optional)
	limitGameCenterDetails := int32(56) // int32 | maximum number of related gameCenterDetails returned (when they are included) (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGetCollection(context.Background()).FilterGameCenterDetails(filterGameCenterDetails).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGetCollection`: GameCenterGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGameCenterDetails** | **[]string** | filter by id(s) of related &#39;gameCenterDetails&#39; | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitGameCenterAchievements** | **int32** | maximum number of related gameCenterAchievements returned (when they are included) | 
 **limitGameCenterDetails** | **int32** | maximum number of related gameCenterDetails returned (when they are included) | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 

### Return type

[**GameCenterGroupsResponse**](GameCenterGroupsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterGroupsGetInstance

> GameCenterGroupResponse GameCenterGroupsGetInstance(ctx, id).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).Execute()



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
	limitGameCenterAchievements := int32(56) // int32 | maximum number of related gameCenterAchievements returned (when they are included) (optional)
	limitGameCenterDetails := int32(56) // int32 | maximum number of related gameCenterDetails returned (when they are included) (optional)
	limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
	limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsGetInstance(context.Background(), id).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).LimitGameCenterAchievements(limitGameCenterAchievements).LimitGameCenterDetails(limitGameCenterDetails).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsGetInstance`: GameCenterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitGameCenterAchievements** | **int32** | maximum number of related gameCenterAchievements returned (when they are included) | 
 **limitGameCenterDetails** | **int32** | maximum number of related gameCenterDetails returned (when they are included) | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 

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


## GameCenterGroupsUpdateInstance

> GameCenterGroupResponse GameCenterGroupsUpdateInstance(ctx, id).GameCenterGroupUpdateRequest(gameCenterGroupUpdateRequest).Execute()



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
	gameCenterGroupUpdateRequest := *openapiclient.NewGameCenterGroupUpdateRequest(*openapiclient.NewGameCenterGroupUpdateRequestData("Type_example", "Id_example")) // GameCenterGroupUpdateRequest | GameCenterGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterGroupsAPI.GameCenterGroupsUpdateInstance(context.Background(), id).GameCenterGroupUpdateRequest(gameCenterGroupUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterGroupsAPI.GameCenterGroupsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterGroupsUpdateInstance`: GameCenterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterGroupsAPI.GameCenterGroupsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterGroupsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterGroupUpdateRequest** | [**GameCenterGroupUpdateRequest**](GameCenterGroupUpdateRequest.md) | GameCenterGroup representation | 

### Return type

[**GameCenterGroupResponse**](GameCenterGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

