# \BetaGroupsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaGroupsAppGetToOneRelated**](BetaGroupsAPI.md#BetaGroupsAppGetToOneRelated) | **Get** /v1/betaGroups/{id}/app | 
[**BetaGroupsBetaTesterUsagesGetMetrics**](BetaGroupsAPI.md#BetaGroupsBetaTesterUsagesGetMetrics) | **Get** /v1/betaGroups/{id}/metrics/betaTesterUsages | 
[**BetaGroupsBetaTestersCreateToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBetaTestersCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersDeleteToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBetaTestersDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelated**](BetaGroupsAPI.md#BetaGroupsBetaTestersGetToManyRelated) | **Get** /v1/betaGroups/{id}/betaTesters | 
[**BetaGroupsBetaTestersGetToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBetaTestersGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/betaTesters | 
[**BetaGroupsBuildsCreateToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBuildsCreateToManyRelationship) | **Post** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsDeleteToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBuildsDeleteToManyRelationship) | **Delete** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsBuildsGetToManyRelated**](BetaGroupsAPI.md#BetaGroupsBuildsGetToManyRelated) | **Get** /v1/betaGroups/{id}/builds | 
[**BetaGroupsBuildsGetToManyRelationship**](BetaGroupsAPI.md#BetaGroupsBuildsGetToManyRelationship) | **Get** /v1/betaGroups/{id}/relationships/builds | 
[**BetaGroupsCreateInstance**](BetaGroupsAPI.md#BetaGroupsCreateInstance) | **Post** /v1/betaGroups | 
[**BetaGroupsDeleteInstance**](BetaGroupsAPI.md#BetaGroupsDeleteInstance) | **Delete** /v1/betaGroups/{id} | 
[**BetaGroupsGetCollection**](BetaGroupsAPI.md#BetaGroupsGetCollection) | **Get** /v1/betaGroups | 
[**BetaGroupsGetInstance**](BetaGroupsAPI.md#BetaGroupsGetInstance) | **Get** /v1/betaGroups/{id} | 
[**BetaGroupsUpdateInstance**](BetaGroupsAPI.md#BetaGroupsUpdateInstance) | **Patch** /v1/betaGroups/{id} | 



## BetaGroupsAppGetToOneRelated

> AppWithoutIncludesResponse BetaGroupsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsAppGetToOneRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsAppGetToOneRelated`: AppWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsAppGetToOneRelatedRequest struct via the builder pattern


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


## BetaGroupsBetaTesterUsagesGetMetrics

> AppsBetaTesterUsagesV1MetricResponse BetaGroupsBetaTesterUsagesGetMetrics(ctx, id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()



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
	period := "P7D" // string | the duration of the reporting period (optional)
	groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
	filterBetaTesters := "filterBetaTesters_example" // string | filter by 'betaTesters' relationship dimension (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsBetaTesterUsagesGetMetrics(context.Background(), id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBetaTesterUsagesGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBetaTesterUsagesGetMetrics`: AppsBetaTesterUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsBetaTesterUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTesterUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | the duration of the reporting period | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterBetaTesters** | **string** | filter by &#39;betaTesters&#39; relationship dimension | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**AppsBetaTesterUsagesV1MetricResponse**](AppsBetaTesterUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTestersCreateToManyRelationship

> BetaGroupsBetaTestersCreateToManyRelationship(ctx, id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()



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
	betaGroupBetaTestersLinkagesRequest := *openapiclient.NewBetaGroupBetaTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersDataInner{*openapiclient.NewBetaGroupRelationshipsBetaTestersDataInner("Type_example", "Id_example")}) // BetaGroupBetaTestersLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaGroupsAPI.BetaGroupsBetaTestersCreateToManyRelationship(context.Background(), id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBetaTestersCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md) | List of related linkages | 

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


## BetaGroupsBetaTestersDeleteToManyRelationship

> BetaGroupsBetaTestersDeleteToManyRelationship(ctx, id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()



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
	betaGroupBetaTestersLinkagesRequest := *openapiclient.NewBetaGroupBetaTestersLinkagesRequest([]openapiclient.BetaGroupRelationshipsBetaTestersDataInner{*openapiclient.NewBetaGroupRelationshipsBetaTestersDataInner("Type_example", "Id_example")}) // BetaGroupBetaTestersLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaGroupsAPI.BetaGroupsBetaTestersDeleteToManyRelationship(context.Background(), id).BetaGroupBetaTestersLinkagesRequest(betaGroupBetaTestersLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBetaTestersDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBetaTestersLinkagesRequest** | [**BetaGroupBetaTestersLinkagesRequest**](BetaGroupBetaTestersLinkagesRequest.md) | List of related linkages | 

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


## BetaGroupsBetaTestersGetToManyRelated

> BetaTestersWithoutIncludesResponse BetaGroupsBetaTestersGetToManyRelated(ctx, id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()



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
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelated(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBetaTestersGetToManyRelated`: BetaTestersWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersGetToManyRelatedRequest struct via the builder pattern


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


## BetaGroupsBetaTestersGetToManyRelationship

> BetaGroupBetaTestersLinkagesResponse BetaGroupsBetaTestersGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBetaTestersGetToManyRelationship`: BetaGroupBetaTestersLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsBetaTestersGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTestersGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaGroupBetaTestersLinkagesResponse**](BetaGroupBetaTestersLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsCreateToManyRelationship

> BetaGroupsBuildsCreateToManyRelationship(ctx, id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()



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
	betaGroupBuildsLinkagesRequest := *openapiclient.NewBetaGroupBuildsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsBuildsDataInner{*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")}) // BetaGroupBuildsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaGroupsAPI.BetaGroupsBuildsCreateToManyRelationship(context.Background(), id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBuildsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBuildsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaGroupsBuildsDeleteToManyRelationship

> BetaGroupsBuildsDeleteToManyRelationship(ctx, id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()



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
	betaGroupBuildsLinkagesRequest := *openapiclient.NewBetaGroupBuildsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsBuildsDataInner{*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")}) // BetaGroupBuildsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaGroupsAPI.BetaGroupsBuildsDeleteToManyRelationship(context.Background(), id).BetaGroupBuildsLinkagesRequest(betaGroupBuildsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBuildsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsBuildsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupBuildsLinkagesRequest** | [**BetaGroupBuildsLinkagesRequest**](BetaGroupBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaGroupsBuildsGetToManyRelated

> BuildsWithoutIncludesResponse BetaGroupsBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBuildsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBuildsGetToManyRelated`: BuildsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBuildsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BuildsWithoutIncludesResponse**](BuildsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBuildsGetToManyRelationship

> BetaGroupBuildsLinkagesResponse BetaGroupsBuildsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsBuildsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsBuildsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBuildsGetToManyRelationship`: BetaGroupBuildsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsBuildsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBuildsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaGroupBuildsLinkagesResponse**](BetaGroupBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsCreateInstance

> BetaGroupResponse BetaGroupsCreateInstance(ctx).BetaGroupCreateRequest(betaGroupCreateRequest).Execute()



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
	betaGroupCreateRequest := *openapiclient.NewBetaGroupCreateRequest(*openapiclient.NewBetaGroupCreateRequestData("Type_example", *openapiclient.NewBetaGroupCreateRequestDataAttributes("Name_example"), *openapiclient.NewBetaGroupCreateRequestDataRelationships(*openapiclient.NewAnalyticsReportRequestCreateRequestDataRelationshipsApp(*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // BetaGroupCreateRequest | BetaGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsCreateInstance(context.Background()).BetaGroupCreateRequest(betaGroupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsCreateInstance`: BetaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaGroupCreateRequest** | [**BetaGroupCreateRequest**](BetaGroupCreateRequest.md) | BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsDeleteInstance

> BetaGroupsDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.BetaGroupsAPI.BetaGroupsDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaGroupsDeleteInstanceRequest struct via the builder pattern


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


## BetaGroupsGetCollection

> BetaGroupsResponse BetaGroupsGetCollection(ctx).FilterName(filterName).FilterIsInternalGroup(filterIsInternalGroup).FilterPublicLinkEnabled(filterPublicLinkEnabled).FilterPublicLinkLimitEnabled(filterPublicLinkLimitEnabled).FilterPublicLink(filterPublicLink).FilterApp(filterApp).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaGroups(fieldsBetaGroups).FieldsApps(fieldsApps).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Include(include).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()



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
	filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
	filterIsInternalGroup := []string{"Inner_example"} // []string | filter by attribute 'isInternalGroup' (optional)
	filterPublicLinkEnabled := []string{"Inner_example"} // []string | filter by attribute 'publicLinkEnabled' (optional)
	filterPublicLinkLimitEnabled := []string{"Inner_example"} // []string | filter by attribute 'publicLinkLimitEnabled' (optional)
	filterPublicLink := []string{"Inner_example"} // []string | filter by attribute 'publicLink' (optional)
	filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
	filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBetaTesters := int32(56) // int32 | maximum number of related betaTesters returned (when they are included) (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsGetCollection(context.Background()).FilterName(filterName).FilterIsInternalGroup(filterIsInternalGroup).FilterPublicLinkEnabled(filterPublicLinkEnabled).FilterPublicLinkLimitEnabled(filterPublicLinkLimitEnabled).FilterPublicLink(filterPublicLink).FilterApp(filterApp).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaGroups(fieldsBetaGroups).FieldsApps(fieldsApps).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).Limit(limit).Include(include).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsGetCollection`: BetaGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterIsInternalGroup** | **[]string** | filter by attribute &#39;isInternalGroup&#39; | 
 **filterPublicLinkEnabled** | **[]string** | filter by attribute &#39;publicLinkEnabled&#39; | 
 **filterPublicLinkLimitEnabled** | **[]string** | filter by attribute &#39;publicLinkLimitEnabled&#39; | 
 **filterPublicLink** | **[]string** | filter by attribute &#39;publicLink&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBetaTesters** | **int32** | maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupsResponse**](BetaGroupsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsGetInstance

> BetaGroupResponse BetaGroupsGetInstance(ctx, id).FieldsBetaGroups(fieldsBetaGroups).FieldsApps(fieldsApps).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).Include(include).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()



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
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitBetaTesters := int32(56) // int32 | maximum number of related betaTesters returned (when they are included) (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsGetInstance(context.Background(), id).FieldsBetaGroups(fieldsBetaGroups).FieldsApps(fieldsApps).FieldsBuilds(fieldsBuilds).FieldsBetaTesters(fieldsBetaTesters).Include(include).LimitBetaTesters(limitBetaTesters).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsGetInstance`: BetaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBetaTesters** | **int32** | maximum number of related betaTesters returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsUpdateInstance

> BetaGroupResponse BetaGroupsUpdateInstance(ctx, id).BetaGroupUpdateRequest(betaGroupUpdateRequest).Execute()



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
	betaGroupUpdateRequest := *openapiclient.NewBetaGroupUpdateRequest(*openapiclient.NewBetaGroupUpdateRequestData("Type_example", "Id_example")) // BetaGroupUpdateRequest | BetaGroup representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaGroupsAPI.BetaGroupsUpdateInstance(context.Background(), id).BetaGroupUpdateRequest(betaGroupUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaGroupsAPI.BetaGroupsUpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsUpdateInstance`: BetaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaGroupsAPI.BetaGroupsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaGroupUpdateRequest** | [**BetaGroupUpdateRequest**](BetaGroupUpdateRequest.md) | BetaGroup representation | 

### Return type

[**BetaGroupResponse**](BetaGroupResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

