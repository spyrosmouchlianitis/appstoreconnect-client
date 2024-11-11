# \AppPricePointsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPricePointsV3EqualizationsGetToManyRelated**](AppPricePointsAPI.md#AppPricePointsV3EqualizationsGetToManyRelated) | **Get** /v3/appPricePoints/{id}/equalizations | 
[**AppPricePointsV3GetInstance**](AppPricePointsAPI.md#AppPricePointsV3GetInstance) | **Get** /v3/appPricePoints/{id} | 



## AppPricePointsV3EqualizationsGetToManyRelated

> AppPricePointsV3Response AppPricePointsV3EqualizationsGetToManyRelated(ctx, id).FilterTerritory(filterTerritory).FieldsAppPricePoints(fieldsAppPricePoints).FieldsApps(fieldsApps).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
	filterTerritory := []string{"Inner_example"} // []string | filter by id(s) of related 'territory' (optional)
	fieldsAppPricePoints := []string{"FieldsAppPricePoints_example"} // []string | the fields to include for returned resources of type appPricePoints (optional)
	fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
	fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPricePointsAPI.AppPricePointsV3EqualizationsGetToManyRelated(context.Background(), id).FilterTerritory(filterTerritory).FieldsAppPricePoints(fieldsAppPricePoints).FieldsApps(fieldsApps).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPricePointsAPI.AppPricePointsV3EqualizationsGetToManyRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPricePointsV3EqualizationsGetToManyRelated`: AppPricePointsV3Response
	fmt.Fprintf(os.Stdout, "Response from `AppPricePointsAPI.AppPricePointsV3EqualizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPricePointsV3EqualizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterTerritory** | **[]string** | filter by id(s) of related &#39;territory&#39; | 
 **fieldsAppPricePoints** | **[]string** | the fields to include for returned resources of type appPricePoints | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPricePointsV3Response**](AppPricePointsV3Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPricePointsV3GetInstance

> AppPricePointV3Response AppPricePointsV3GetInstance(ctx, id).FieldsAppPricePoints(fieldsAppPricePoints).Include(include).Execute()



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
	fieldsAppPricePoints := []string{"FieldsAppPricePoints_example"} // []string | the fields to include for returned resources of type appPricePoints (optional)
	include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPricePointsAPI.AppPricePointsV3GetInstance(context.Background(), id).FieldsAppPricePoints(fieldsAppPricePoints).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPricePointsAPI.AppPricePointsV3GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPricePointsV3GetInstance`: AppPricePointV3Response
	fmt.Fprintf(os.Stdout, "Response from `AppPricePointsAPI.AppPricePointsV3GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPricePointsV3GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPricePoints** | **[]string** | the fields to include for returned resources of type appPricePoints | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPricePointV3Response**](AppPricePointV3Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

