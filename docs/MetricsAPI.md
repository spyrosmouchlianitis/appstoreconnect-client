# \MetricsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsBetaTesterUsagesGetMetrics_0**](MetricsAPI.md#AppsBetaTesterUsagesGetMetrics_0) | **Get** /v1/apps/{id}/metrics/betaTesterUsages | 
[**BetaGroupsBetaTesterUsagesGetMetrics_0**](MetricsAPI.md#BetaGroupsBetaTesterUsagesGetMetrics_0) | **Get** /v1/betaGroups/{id}/metrics/betaTesterUsages | 
[**BetaTestersBetaTesterUsagesGetMetrics_0**](MetricsAPI.md#BetaTestersBetaTesterUsagesGetMetrics_0) | **Get** /v1/betaTesters/{id}/metrics/betaTesterUsages | 
[**BuildsBetaBuildUsagesGetMetrics_0**](MetricsAPI.md#BuildsBetaBuildUsagesGetMetrics_0) | **Get** /v1/builds/{id}/metrics/betaBuildUsages | 
[**GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0**](MetricsAPI.md#GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0) | **Get** /v1/gameCenterDetails/{id}/metrics/classicMatchmakingRequests | 
[**GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0**](MetricsAPI.md#GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0) | **Get** /v1/gameCenterDetails/{id}/metrics/ruleBasedMatchmakingRequests | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingRequests | 
[**GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingRequests | 
[**GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingSessions | 
[**GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingBooleanRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingNumberRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingRuleErrors | 



## AppsBetaTesterUsagesGetMetrics_0

> AppsBetaTesterUsagesV1MetricResponse AppsBetaTesterUsagesGetMetrics_0(ctx, id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()



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
	resp, r, err := apiClient.MetricsAPI.AppsBetaTesterUsagesGetMetrics_0(context.Background(), id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.AppsBetaTesterUsagesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsBetaTesterUsagesGetMetrics_0`: AppsBetaTesterUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.AppsBetaTesterUsagesGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaTesterUsagesGetMetricsRequest struct via the builder pattern


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


## BetaGroupsBetaTesterUsagesGetMetrics_0

> AppsBetaTesterUsagesV1MetricResponse BetaGroupsBetaTesterUsagesGetMetrics_0(ctx, id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()



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
	resp, r, err := apiClient.MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics_0(context.Background(), id).Period(period).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaGroupsBetaTesterUsagesGetMetrics_0`: AppsBetaTesterUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics_0`: %v\n", resp)
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


## BetaTestersBetaTesterUsagesGetMetrics_0

> BetaTesterUsagesV1MetricResponse BetaTestersBetaTesterUsagesGetMetrics_0(ctx, id).FilterApps(filterApps).Period(period).Limit(limit).Execute()



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
	resp, r, err := apiClient.MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics_0(context.Background(), id).FilterApps(filterApps).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BetaTestersBetaTesterUsagesGetMetrics_0`: BetaTesterUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics_0`: %v\n", resp)
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


## BuildsBetaBuildUsagesGetMetrics_0

> BetaBuildUsagesV1MetricResponse BuildsBetaBuildUsagesGetMetrics_0(ctx, id).Limit(limit).Execute()



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
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.BuildsBetaBuildUsagesGetMetrics_0(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BuildsBetaBuildUsagesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBetaBuildUsagesGetMetrics_0`: BetaBuildUsagesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BuildsBetaBuildUsagesGetMetrics_0`: %v\n", resp)
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


## GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()



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
	resp, r, err := apiClient.MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0`: GameCenterMatchmakingAppRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics_0`: %v\n", resp)
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


## GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()



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
	resp, r, err := apiClient.MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0`: GameCenterMatchmakingAppRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics_0`: %v\n", resp)
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


## GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0

> GameCenterMatchmakingQueueSizesV1MetricResponse GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0(ctx, id).Granularity(granularity).Sort(sort).Limit(limit).Execute()



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
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0(context.Background(), id).Granularity(granularity).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0`: GameCenterMatchmakingQueueSizesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingQueueSizesV1MetricResponse**](GameCenterMatchmakingQueueSizesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0

> GameCenterMatchmakingQueueRequestsV1MetricResponse GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Limit(limit).Execute()



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
	filterGameCenterDetail := "filterGameCenterDetail_example" // string | filter by 'gameCenterDetail' relationship dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0`: GameCenterMatchmakingQueueRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterDetail** | **string** | filter by &#39;gameCenterDetail&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingQueueRequestsV1MetricResponse**](GameCenterMatchmakingQueueRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0

> GameCenterMatchmakingQueueSizesV1MetricResponse GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0(ctx, id).Granularity(granularity).Sort(sort).Limit(limit).Execute()



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
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0(context.Background(), id).Granularity(granularity).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0`: GameCenterMatchmakingQueueSizesV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingQueueSizesV1MetricResponse**](GameCenterMatchmakingQueueSizesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0

> GameCenterMatchmakingQueueRequestsV1MetricResponse GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Limit(limit).Execute()



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
	filterGameCenterDetail := "filterGameCenterDetail_example" // string | filter by 'gameCenterDetail' relationship dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0`: GameCenterMatchmakingQueueRequestsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterDetail** | **string** | filter by &#39;gameCenterDetail&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingQueueRequestsV1MetricResponse**](GameCenterMatchmakingQueueRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0

> GameCenterMatchmakingSessionsV1MetricResponse GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0(ctx, id).Granularity(granularity).Sort(sort).Limit(limit).Execute()



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
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0(context.Background(), id).Granularity(granularity).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0`: GameCenterMatchmakingSessionsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingSessionsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingSessionsV1MetricResponse**](GameCenterMatchmakingSessionsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0

> GameCenterMatchmakingBooleanRuleResultsV1MetricResponse GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()



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
	filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0`: GameCenterMatchmakingBooleanRuleResultsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingBooleanRuleResultsV1MetricResponse**](GameCenterMatchmakingBooleanRuleResultsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0

> GameCenterMatchmakingNumberRuleResultsV1MetricResponse GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()



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
	filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0`: GameCenterMatchmakingNumberRuleResultsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingNumberRuleResultsV1MetricResponse**](GameCenterMatchmakingNumberRuleResultsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0

> GameCenterMatchmakingRuleErrorsV1MetricResponse GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0(ctx, id).Granularity(granularity).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()



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
	filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)
	limit := int32(56) // int32 | maximum number of groups to return per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0(context.Background(), id).Granularity(granularity).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0`: GameCenterMatchmakingRuleErrorsV1MetricResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **string** | the granularity of the per-group dataset | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 
 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**GameCenterMatchmakingRuleErrorsV1MetricResponse**](GameCenterMatchmakingRuleErrorsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

