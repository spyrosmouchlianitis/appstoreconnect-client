# \AlternativeDistributionPackageVariantsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlternativeDistributionPackageVariantsGetInstance**](AlternativeDistributionPackageVariantsAPI.md#AlternativeDistributionPackageVariantsGetInstance) | **Get** /v1/alternativeDistributionPackageVariants/{id} | 



## AlternativeDistributionPackageVariantsGetInstance

> AlternativeDistributionPackageVariantResponse AlternativeDistributionPackageVariantsGetInstance(ctx, id).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).Execute()



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
	fieldsAlternativeDistributionPackageVariants := []string{"FieldsAlternativeDistributionPackageVariants_example"} // []string | the fields to include for returned resources of type alternativeDistributionPackageVariants (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeDistributionPackageVariantsAPI.AlternativeDistributionPackageVariantsGetInstance(context.Background(), id).FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeDistributionPackageVariantsAPI.AlternativeDistributionPackageVariantsGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlternativeDistributionPackageVariantsGetInstance`: AlternativeDistributionPackageVariantResponse
	fmt.Fprintf(os.Stdout, "Response from `AlternativeDistributionPackageVariantsAPI.AlternativeDistributionPackageVariantsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlternativeDistributionPackageVariantsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAlternativeDistributionPackageVariants** | **[]string** | the fields to include for returned resources of type alternativeDistributionPackageVariants | 

### Return type

[**AlternativeDistributionPackageVariantResponse**](AlternativeDistributionPackageVariantResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

