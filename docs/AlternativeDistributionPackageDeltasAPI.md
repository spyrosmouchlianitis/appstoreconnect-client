# \AlternativeDistributionPackageDeltasAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionPackageDeltasGetInstance**](AlternativeDistributionPackageDeltasAPI.md#AlternativeDistributionPackageDeltasGetInstance) | **Get** /v1/alternativeDistributionPackageDeltas/{id} | 



## AlternativeDistributionPackageDeltasGetInstance

> AlternativeDistributionPackageDeltaResponse AlternativeDistributionPackageDeltasGetInstance(ctx, id).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Execute()



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
	fieldsAlternativeDistributionPackageDeltas := []string{"FieldsAlternativeDistributionPackageDeltas_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageDeltas (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackageDeltasAPI.AlternativeDistributionPackageDeltasGetInstance(context.Background(), id).FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackageDeltasAPI.AlternativeDistributionPackageDeltasGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackageDeltasGetInstance`: AlternativeDistributionPackageDeltaResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackageDeltasAPI.AlternativeDistributionPackageDeltasGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackageDeltasGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackageDeltas** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageDeltas | 

### Return type

[**AlternativeDistributionPackageDeltaResponse**](AlternativeDistributionPackageDeltaResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

