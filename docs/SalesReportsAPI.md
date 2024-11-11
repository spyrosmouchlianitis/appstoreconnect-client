# \SalesReportsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SalesReportsGetCollection**](SalesReportsAPI.md#SalesReportsGetCollection) | **Get** /v1/salesReports | 



## SalesReportsGetCollection

> *os.File SalesReportsGetCollection(ctx).FilterVendorNumber(filterVendorNumber).FilterReportType(filterReportType).FilterReportSubType(filterReportSubType).FilterFrequency(filterFrequency).FilterReportDate(filterReportDate).FilterVersion(filterVersion).Execute()



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
	filterVendorNumber := []string{"Inner_example"} // []string | filter by attribute 'vendorNumber'
	filterReportType := []string{"FilterReportType_example"} // []string | filter by attribute 'reportType'
	filterReportSubType := []string{"FilterReportSubType_example"} // []string | filter by attribute 'reportSubType'
	filterFrequency := []string{"FilterFrequency_example"} // []string | filter by attribute 'frequency'
	filterReportDate := []string{"Inner_example"} // []string | filter by attribute 'reportDate' (optional)
	filterVersion := []string{"Inner_example"} // []string | filter by attribute 'version' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesReportsAPI.SalesReportsGetCollection(context.Background()).FilterVendorNumber(filterVendorNumber).FilterReportType(filterReportType).FilterReportSubType(filterReportSubType).FilterFrequency(filterFrequency).FilterReportDate(filterReportDate).FilterVersion(filterVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesReportsAPI.SalesReportsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalesReportsGetCollection`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SalesReportsAPI.SalesReportsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalesReportsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterVendorNumber** | **[]string** | filter by attribute &#39;vendorNumber&#39; | 
 **filterReportType** | **[]string** | filter by attribute &#39;reportType&#39; | 
 **filterReportSubType** | **[]string** | filter by attribute &#39;reportSubType&#39; | 
 **filterFrequency** | **[]string** | filter by attribute &#39;frequency&#39; | 
 **filterReportDate** | **[]string** | filter by attribute &#39;reportDate&#39; | 
 **filterVersion** | **[]string** | filter by attribute &#39;version&#39; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/a-gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

