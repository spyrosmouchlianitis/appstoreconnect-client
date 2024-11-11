# \BetaTestersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaTestersAppsDeleteToManyRelationship**](BetaTestersAPI.md#BetaTestersAppsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersAppsGetToManyRelated**](BetaTestersAPI.md#BetaTestersAppsGetToManyRelated) | **Get** /v1/betaTesters/{id}/apps | 
[**BetaTestersAppsGetToManyRelationship**](BetaTestersAPI.md#BetaTestersAppsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/apps | 
[**BetaTestersBetaGroupsCreateToManyRelationship**](BetaTestersAPI.md#BetaTestersBetaGroupsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsDeleteToManyRelationship**](BetaTestersAPI.md#BetaTestersBetaGroupsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelated**](BetaTestersAPI.md#BetaTestersBetaGroupsGetToManyRelated) | **Get** /v1/betaTesters/{id}/betaGroups | 
[**BetaTestersBetaGroupsGetToManyRelationship**](BetaTestersAPI.md#BetaTestersBetaGroupsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/betaGroups | 
[**BetaTestersBetaTesterUsagesGetMetrics**](BetaTestersAPI.md#BetaTestersBetaTesterUsagesGetMetrics) | **Get** /v1/betaTesters/{id}/metrics/betaTesterUsages | 
[**BetaTestersBuildsCreateToManyRelationship**](BetaTestersAPI.md#BetaTestersBuildsCreateToManyRelationship) | **Post** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsDeleteToManyRelationship**](BetaTestersAPI.md#BetaTestersBuildsDeleteToManyRelationship) | **Delete** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersBuildsGetToManyRelated**](BetaTestersAPI.md#BetaTestersBuildsGetToManyRelated) | **Get** /v1/betaTesters/{id}/builds | 
[**BetaTestersBuildsGetToManyRelationship**](BetaTestersAPI.md#BetaTestersBuildsGetToManyRelationship) | **Get** /v1/betaTesters/{id}/relationships/builds | 
[**BetaTestersCreateInstance**](BetaTestersAPI.md#BetaTestersCreateInstance) | **Post** /v1/betaTesters | 
[**BetaTestersDeleteInstance**](BetaTestersAPI.md#BetaTestersDeleteInstance) | **Delete** /v1/betaTesters/{id} | 
[**BetaTestersGetCollection**](BetaTestersAPI.md#BetaTestersGetCollection) | **Get** /v1/betaTesters | 
[**BetaTestersGetInstance**](BetaTestersAPI.md#BetaTestersGetInstance) | **Get** /v1/betaTesters/{id} | 



## BetaTestersAppsDeleteToManyRelationship

> BetaTestersAppsDeleteToManyRelationship(ctx, id).BetaTesterAppsLinkagesRequest(betaTesterAppsLinkagesRequest).Execute()



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
	betaTesterAppsLinkagesRequest := *openapiclient.NewBetaTesterAppsLinkagesRequest([]openapiclient.AlternativeDistributionKeyCreateRequestDataRelationshipsAppData{*openapiclient.NewAlternativeDistributionKeyCreateRequestDataRelationshipsAppData("Type_example", "Id_example")}) // BetaTesterAppsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaTestersAPI.BetaTestersAppsDeleteToManyRelationship(context.Background(), id).BetaTesterAppsLinkagesRequest(betaTesterAppsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersAppsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersAppsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterAppsLinkagesRequest** | [**BetaTesterAppsLinkagesRequest**](BetaTesterAppsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersAppsGetToManyRelated

> AppsWithoutIncludesResponse BetaTestersAppsGetToManyRelated(ctx, id).FieldsApps(fieldsApps).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersAppsGetToManyRelated(context.Background(), id).FieldsApps(fieldsApps).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersAppsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersAppsGetToManyRelated`: AppsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersAppsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersAppsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppsWithoutIncludesResponse**](AppsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersAppsGetToManyRelationship

> BetaTesterAppsLinkagesResponse BetaTestersAppsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersAppsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersAppsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersAppsGetToManyRelationship`: BetaTesterAppsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersAppsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersAppsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterAppsLinkagesResponse**](BetaTesterAppsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaGroupsCreateToManyRelationship

> BetaTestersBetaGroupsCreateToManyRelationship(ctx, id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()



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
	betaTesterBetaGroupsLinkagesRequest := *openapiclient.NewBetaTesterBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsDataInner{*openapiclient.NewAppRelationshipsBetaGroupsDataInner("Type_example", "Id_example")}) // BetaTesterBetaGroupsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaTestersAPI.BetaTestersBetaGroupsCreateToManyRelationship(context.Background(), id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBetaGroupsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBetaGroupsDeleteToManyRelationship

> BetaTestersBetaGroupsDeleteToManyRelationship(ctx, id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()



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
	betaTesterBetaGroupsLinkagesRequest := *openapiclient.NewBetaTesterBetaGroupsLinkagesRequest([]openapiclient.AppRelationshipsBetaGroupsDataInner{*openapiclient.NewAppRelationshipsBetaGroupsDataInner("Type_example", "Id_example")}) // BetaTesterBetaGroupsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaTestersAPI.BetaTestersBetaGroupsDeleteToManyRelationship(context.Background(), id).BetaTesterBetaGroupsLinkagesRequest(betaTesterBetaGroupsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBetaGroupsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBetaGroupsLinkagesRequest** | [**BetaTesterBetaGroupsLinkagesRequest**](BetaTesterBetaGroupsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBetaGroupsGetToManyRelated

> BetaGroupsWithoutIncludesResponse BetaTestersBetaGroupsGetToManyRelated(ctx, id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelated(context.Background(), id).FieldsBetaGroups(fieldsBetaGroups).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBetaGroupsGetToManyRelated`: BetaGroupsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaGroupsWithoutIncludesResponse**](BetaGroupsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaGroupsGetToManyRelationship

> BetaTesterBetaGroupsLinkagesResponse BetaTestersBetaGroupsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBetaGroupsGetToManyRelationship`: BetaTesterBetaGroupsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersBetaGroupsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaGroupsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterBetaGroupsLinkagesResponse**](BetaTesterBetaGroupsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaTesterUsagesGetMetrics

> BetaTesterUsagesV1MetricResponse BetaTestersBetaTesterUsagesGetMetrics(ctx, id).FilterApps(filterApps).Period(period).Limit(limit).Execute()



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
	filterApps := "filterApps_example" // string | filter by 'apps' relationship dimension
	period := "P7D" // string | the duration of the reporting period (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersBetaTesterUsagesGetMetrics(context.Background(), id).FilterApps(filterApps).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBetaTesterUsagesGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBetaTesterUsagesGetMetrics`: BetaTesterUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersBetaTesterUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaTesterUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterApps** | **string** | filter by &#39;apps&#39; relationship dimension | 
 **period** | **string** | the duration of the reporting period | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**BetaTesterUsagesV1MetricResponse**](BetaTesterUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBuildsCreateToManyRelationship

> BetaTestersBuildsCreateToManyRelationship(ctx, id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()



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
	betaTesterBuildsLinkagesRequest := *openapiclient.NewBetaTesterBuildsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsBuildsDataInner{*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")}) // BetaTesterBuildsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaTestersAPI.BetaTestersBuildsCreateToManyRelationship(context.Background(), id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBuildsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBuildsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBuildsDeleteToManyRelationship

> BetaTestersBuildsDeleteToManyRelationship(ctx, id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()



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
	betaTesterBuildsLinkagesRequest := *openapiclient.NewBetaTesterBuildsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsBuildsDataInner{*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")}) // BetaTesterBuildsLinkagesRequest | List of related linkages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BetaTestersAPI.BetaTestersBuildsDeleteToManyRelationship(context.Background(), id).BetaTesterBuildsLinkagesRequest(betaTesterBuildsLinkagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBuildsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersBuildsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaTesterBuildsLinkagesRequest** | [**BetaTesterBuildsLinkagesRequest**](BetaTesterBuildsLinkagesRequest.md) | List of related linkages | 

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


## BetaTestersBuildsGetToManyRelated

> BuildsWithoutIncludesResponse BetaTestersBuildsGetToManyRelated(ctx, id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersBuildsGetToManyRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBuildsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBuildsGetToManyRelated`: BuildsWithoutIncludesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersBuildsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBuildsGetToManyRelatedRequest struct via the builder pattern


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


## BetaTestersBuildsGetToManyRelationship

> BetaTesterBuildsLinkagesResponse BetaTestersBuildsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersBuildsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersBuildsGetToManyRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBuildsGetToManyRelationship`: BetaTesterBuildsLinkagesResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersBuildsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBuildsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**BetaTesterBuildsLinkagesResponse**](BetaTesterBuildsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersCreateInstance

> BetaTesterResponse BetaTestersCreateInstance(ctx).BetaTesterCreateRequest(betaTesterCreateRequest).Execute()



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
	betaTesterCreateRequest := *openapiclient.NewBetaTesterCreateRequest(*openapiclient.NewBetaTesterCreateRequestData("Type_example", *openapiclient.NewBetaTesterCreateRequestDataAttributes("Email_example"))) // BetaTesterCreateRequest | BetaTester representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersCreateInstance(context.Background()).BetaTesterCreateRequest(betaTesterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersCreateInstance`: BetaTesterResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaTesterCreateRequest** | [**BetaTesterCreateRequest**](BetaTesterCreateRequest.md) | BetaTester representation | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersDeleteInstance

> BetaTestersDeleteInstance(ctx, id).Execute()



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
	r, err := apiClient.BetaTestersAPI.BetaTestersDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaTestersDeleteInstanceRequest struct via the builder pattern


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


## BetaTestersGetCollection

> BetaTestersResponse BetaTestersGetCollection(ctx).FilterFirstName(filterFirstName).FilterLastName(filterLastName).FilterEmail(filterEmail).FilterInviteType(filterInviteType).FilterApps(filterApps).FilterBetaGroups(filterBetaGroups).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).Limit(limit).Include(include).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()



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
	filterFirstName := []string{"Inner_example"} // []string | filter by attribute 'firstName' (optional)
	filterLastName := []string{"Inner_example"} // []string | filter by attribute 'lastName' (optional)
	filterEmail := []string{"Inner_example"} // []string | filter by attribute 'email' (optional)
	filterInviteType := []string{"FilterInviteType_example"} // []string | filter by attribute 'inviteType' (optional)
	filterApps := []string{"Inner_example"} // []string | filter by id(s) of related 'apps' (optional)
	filterBetaGroups := []string{"Inner_example"} // []string | filter by id(s) of related 'betaGroups' (optional)
	filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsBetaTesters := []string{"FieldsBetaTesters_example"} // []string | the fields to include for returned resources of type betaTesters (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersGetCollection(context.Background()).FilterFirstName(filterFirstName).FilterLastName(filterLastName).FilterEmail(filterEmail).FilterInviteType(filterInviteType).FilterApps(filterApps).FilterBetaGroups(filterBetaGroups).FilterBuilds(filterBuilds).FilterId(filterId).Sort(sort).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).Limit(limit).Include(include).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersGetCollection`: BetaTestersResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterFirstName** | **[]string** | filter by attribute &#39;firstName&#39; | 
 **filterLastName** | **[]string** | filter by attribute &#39;lastName&#39; | 
 **filterEmail** | **[]string** | filter by attribute &#39;email&#39; | 
 **filterInviteType** | **[]string** | filter by attribute &#39;inviteType&#39; | 
 **filterApps** | **[]string** | filter by id(s) of related &#39;apps&#39; | 
 **filterBetaGroups** | **[]string** | filter by id(s) of related &#39;betaGroups&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaTestersResponse**](BetaTestersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersGetInstance

> BetaTesterResponse BetaTestersGetInstance(ctx, id).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).Include(include).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()



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
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsBetaGroups := []string{"FieldsBetaGroups_example"} // []string | the fields to include for returned resources of type betaGroups (optional)
	fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
	limitApps := int32(56) // int32 | maximum number of related apps returned (when they are included) (optional)
	limitBetaGroups := int32(56) // int32 | maximum number of related betaGroups returned (when they are included) (optional)
	limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BetaTestersAPI.BetaTestersGetInstance(context.Background(), id).FieldsBetaTesters(fieldsBetaTesters).FieldsApps(fieldsApps).FieldsBetaGroups(fieldsBetaGroups).FieldsBuilds(fieldsBuilds).Include(include).LimitApps(limitApps).LimitBetaGroups(limitBetaGroups).LimitBuilds(limitBuilds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BetaTestersAPI.BetaTestersGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersGetInstance`: BetaTesterResponse
	fmt.Fprintf(os.Stdout, "Response from `BetaTestersAPI.BetaTestersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaTesters** | **[]string** | the fields to include for returned resources of type betaTesters | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsBetaGroups** | **[]string** | the fields to include for returned resources of type betaGroups | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitApps** | **int32** | maximum number of related apps returned (when they are included) | 
 **limitBetaGroups** | **int32** | maximum number of related betaGroups returned (when they are included) | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**BetaTesterResponse**](BetaTesterResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

