# \CertificatesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesCreateInstance**](CertificatesAPI.md#CertificatesCreateInstance) | **Post** /v1/certificates | 
[**CertificatesDeleteInstance**](CertificatesAPI.md#CertificatesDeleteInstance) | **Delete** /v1/certificates/{id} | 
[**CertificatesGetCollection**](CertificatesAPI.md#CertificatesGetCollection) | **Get** /v1/certificates | 
[**CertificatesGetInstance**](CertificatesAPI.md#CertificatesGetInstance) | **Get** /v1/certificates/{id} | 



## CertificatesCreateInstance

> CertificateResponse CertificatesCreateInstance(ctx).CertificateCreateRequest(certificateCreateRequest).Execute()



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
	certificateCreateRequest := *openapiclient.NewCertificateCreateRequest(*openapiclient.NewCertificateCreateRequestData("Type_example", *openapiclient.NewCertificateCreateRequestDataAttributes("CsrContent_example", openapiclient.CertificateType("IOS_DEVELOPMENT")))) // CertificateCreateRequest | Certificate representation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesCreateInstance(context.Background()).CertificateCreateRequest(certificateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesCreateInstance`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificateCreateRequest** | [**CertificateCreateRequest**](CertificateCreateRequest.md) | Certificate representation | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesDeleteInstance

> CertificatesDeleteInstance(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CertificatesAPI.CertificatesDeleteInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCertificatesDeleteInstanceRequest struct via the builder pattern


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


## CertificatesGetCollection

> CertificatesResponse CertificatesGetCollection(ctx).FilterDisplayName(filterDisplayName).FilterCertificateType(filterCertificateType).FilterSerialNumber(filterSerialNumber).FilterId(filterId).Sort(sort).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()



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
	filterDisplayName := []string{"Inner_example"} // []string | filter by attribute 'displayName' (optional)
	filterCertificateType := []string{"FilterCertificateType_example"} // []string | filter by attribute 'certificateType' (optional)
	filterSerialNumber := []string{"Inner_example"} // []string | filter by attribute 'serialNumber' (optional)
	filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
	sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
	fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)
	limit := int32(56) // int32 | maximum resources per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesGetCollection(context.Background()).FilterDisplayName(filterDisplayName).FilterCertificateType(filterCertificateType).FilterSerialNumber(filterSerialNumber).FilterId(filterId).Sort(sort).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesGetCollection`: CertificatesResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterDisplayName** | **[]string** | filter by attribute &#39;displayName&#39; | 
 **filterCertificateType** | **[]string** | filter by attribute &#39;certificateType&#39; | 
 **filterSerialNumber** | **[]string** | filter by attribute &#39;serialNumber&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CertificatesResponse**](CertificatesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificatesGetInstance

> CertificateResponse CertificatesGetInstance(ctx, id).FieldsCertificates(fieldsCertificates).Execute()



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
	fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificatesAPI.CertificatesGetInstance(context.Background(), id).FieldsCertificates(fieldsCertificates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificatesAPI.CertificatesGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificatesGetInstance`: CertificateResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificatesAPI.CertificatesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 

### Return type

[**CertificateResponse**](CertificateResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

