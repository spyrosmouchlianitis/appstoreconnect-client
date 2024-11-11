# \GameCenterEnabledVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship**](GameCenterEnabledVersionsAPI.md#GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship) | **Post** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship**](GameCenterEnabledVersionsAPI.md#GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship) | **Delete** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated**](GameCenterEnabledVersionsAPI.md#GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated) | **Get** /v1/gameCenterEnabledVersions/{id}/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship**](GameCenterEnabledVersionsAPI.md#GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship) | **Get** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 
[**GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship**](GameCenterEnabledVersionsAPI.md#GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship) | **Patch** /v1/gameCenterEnabledVersions/{id}/relationships/compatibleVersions | 



## GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship(ctx, id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()



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
	gameCenterEnabledVersionCompatibleVersionsLinkagesRequest := *openapiclient.NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequest([]openapiclient.AppRelationshipsGameCenterEnabledVersionsDataInner{*openapiclient.NewAppRelationshipsGameCenterEnabledVersionsDataInner("Type_example", "Id_example")}) // GameCenterEnabledVersionCompatibleVersionsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship(context.Background(), id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterEnabledVersionsCompatibleVersionsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship(ctx, id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()



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
	gameCenterEnabledVersionCompatibleVersionsLinkagesRequest := *openapiclient.NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequest([]openapiclient.AppRelationshipsGameCenterEnabledVersionsDataInner{*openapiclient.NewAppRelationshipsGameCenterEnabledVersionsDataInner("Type_example", "Id_example")}) // GameCenterEnabledVersionCompatibleVersionsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship(context.Background(), id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterEnabledVersionsCompatibleVersionsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated

> GameCenterEnabledVersionsResponse GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated(ctx, id).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterApp(filterApp).FilterId(filterId).Sort(sort).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsApps(fieldsApps).Limit(limit).Include(include).LimitCompatibleVersions(limitCompatibleVersions).Execute()



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
	filterVersionString := []string{"Inner_example"} // []string | filter by attribute 'versionString' (optional)
	filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsGameCenterEnabledVersions := []string{"FieldsGameCenterEnabledVersions_example"} // []string | the fields to include for returned resources of type gameCenterEnabledVersions (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitCompatibleVersions := int32(56) // int32 | maximum number of related compatibleVersions returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated(context.Background(), id).FilterPlatform(filterPlatform).FilterVersionString(filterVersionString).FilterApp(filterApp).FilterId(filterId).Sort(sort).FieldsGameCenterEnabledVersions(fieldsGameCenterEnabledVersions).FieldsApps(fieldsApps).Limit(limit).Include(include).LimitCompatibleVersions(limitCompatibleVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated`: GameCenterEnabledVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterEnabledVersionsCompatibleVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterVersionString** | **[]string** | filter by attribute &#39;versionString&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsGameCenterEnabledVersions** | **[]string** | the fields to include for returned resources of type gameCenterEnabledVersions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitCompatibleVersions** | **int32** | maximum number of related compatibleVersions returned (when they are included) | 

### Return type

[**GameCenterEnabledVersionsResponse**](GameCenterEnabledVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship

> GameCenterEnabledVersionCompatibleVersionsLinkagesResponse GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship`: GameCenterEnabledVersionCompatibleVersionsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterEnabledVersionsCompatibleVersionsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterEnabledVersionCompatibleVersionsLinkagesResponse**](GameCenterEnabledVersionCompatibleVersionsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship

> GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship(ctx, id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()



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
	gameCenterEnabledVersionCompatibleVersionsLinkagesRequest := *openapiclient.NewGameCenterEnabledVersionCompatibleVersionsLinkagesRequest([]openapiclient.AppRelationshipsGameCenterEnabledVersionsDataInner{*openapiclient.NewAppRelationshipsGameCenterEnabledVersionsDataInner("Type_example", "Id_example")}) // GameCenterEnabledVersionCompatibleVersionsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship(context.Background(), id).GameCenterEnabledVersionCompatibleVersionsLinkagesRequest(gameCenterEnabledVersionCompatibleVersionsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GameCenterEnabledVersionsAPI.GameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterEnabledVersionsCompatibleVersionsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterEnabledVersionCompatibleVersionsLinkagesRequest** | [**GameCenterEnabledVersionCompatibleVersionsLinkagesRequest**](GameCenterEnabledVersionCompatibleVersionsLinkagesRequest.md) | List of related linkages | 

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

