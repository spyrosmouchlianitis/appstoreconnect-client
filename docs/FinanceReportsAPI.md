# \FinanceReportsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinanceReportsGetCollection**](FinanceReportsAPI.md#FinanceReportsGetCollection) | **Get** /v1/financeReports | 



## FinanceReportsGetCollection

> *os.File FinanceReportsGetCollection(ctx).FilterVendorNumber(filterVendorNumber).FilterReportType(filterReportType).FilterRegionCode(filterRegionCode).FilterReportDate(filterReportDate).Execute()



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
	filterVendorNumber := []string{"Inner_example"} // []string | filter by attribute 'vendorNumber'
	filterReportType := []string{"FilterReportType_example"} // []string | filter by attribute 'reportType'
	filterRegionCode := []string{"Inner_example"} // []string | filter by attribute 'regionCode'
	filterReportDate := []string{"Inner_example"} // []string | filter by attribute 'reportDate'

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinanceReportsAPI.FinanceReportsGetCollection(context.Background()).FilterVendorNumber(filterVendorNumber).FilterReportType(filterReportType).FilterRegionCode(filterRegionCode).FilterReportDate(filterReportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceReportsAPI.FinanceReportsGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinanceReportsGetCollection`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FinanceReportsAPI.FinanceReportsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinanceReportsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterVendorNumber** | **[]string** | filter by attribute &#39;vendorNumber&#39; | 
 **filterReportType** | **[]string** | filter by attribute &#39;reportType&#39; | 
 **filterRegionCode** | **[]string** | filter by attribute &#39;regionCode&#39; | 
 **filterReportDate** | **[]string** | filter by attribute &#39;reportDate&#39; | 

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

