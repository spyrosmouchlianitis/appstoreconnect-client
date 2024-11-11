/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnectclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AnalyticsReportsAPIService AnalyticsReportsAPI service
type AnalyticsReportsAPIService service

type ApiAnalyticsReportsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AnalyticsReportsAPIService
	id string
	fieldsAnalyticsReports *[]string
}

// the fields to include for returned resources of type analyticsReports
func (r ApiAnalyticsReportsGetInstanceRequest) FieldsAnalyticsReports(fieldsAnalyticsReports []string) ApiAnalyticsReportsGetInstanceRequest {
	r.fieldsAnalyticsReports = &fieldsAnalyticsReports
	return r
}

func (r ApiAnalyticsReportsGetInstanceRequest) Execute() (*AnalyticsReportResponse, *http.Response, error) {
	return r.ApiService.AnalyticsReportsGetInstanceExecute(r)
}

/*
AnalyticsReportsGetInstance Method for AnalyticsReportsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAnalyticsReportsGetInstanceRequest
*/
func (a *AnalyticsReportsAPIService) AnalyticsReportsGetInstance(ctx context.Context, id string) ApiAnalyticsReportsGetInstanceRequest {
	return ApiAnalyticsReportsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AnalyticsReportResponse
func (a *AnalyticsReportsAPIService) AnalyticsReportsGetInstanceExecute(r ApiAnalyticsReportsGetInstanceRequest) (*AnalyticsReportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnalyticsReportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsReportsAPIService.AnalyticsReportsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analyticsReports/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAnalyticsReports != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[analyticsReports]", r.fieldsAnalyticsReports, "form", "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAnalyticsReportsInstancesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AnalyticsReportsAPIService
	id string
	filterGranularity *[]string
	filterProcessingDate *[]string
	fieldsAnalyticsReportInstances *[]string
	limit *int32
}

// filter by attribute &#39;granularity&#39;
func (r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) FilterGranularity(filterGranularity []string) ApiAnalyticsReportsInstancesGetToManyRelatedRequest {
	r.filterGranularity = &filterGranularity
	return r
}

// filter by attribute &#39;processingDate&#39;
func (r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) FilterProcessingDate(filterProcessingDate []string) ApiAnalyticsReportsInstancesGetToManyRelatedRequest {
	r.filterProcessingDate = &filterProcessingDate
	return r
}

// the fields to include for returned resources of type analyticsReportInstances
func (r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) FieldsAnalyticsReportInstances(fieldsAnalyticsReportInstances []string) ApiAnalyticsReportsInstancesGetToManyRelatedRequest {
	r.fieldsAnalyticsReportInstances = &fieldsAnalyticsReportInstances
	return r
}

// maximum resources per page
func (r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) Limit(limit int32) ApiAnalyticsReportsInstancesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) Execute() (*AnalyticsReportInstancesResponse, *http.Response, error) {
	return r.ApiService.AnalyticsReportsInstancesGetToManyRelatedExecute(r)
}

/*
AnalyticsReportsInstancesGetToManyRelated Method for AnalyticsReportsInstancesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAnalyticsReportsInstancesGetToManyRelatedRequest
*/
func (a *AnalyticsReportsAPIService) AnalyticsReportsInstancesGetToManyRelated(ctx context.Context, id string) ApiAnalyticsReportsInstancesGetToManyRelatedRequest {
	return ApiAnalyticsReportsInstancesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AnalyticsReportInstancesResponse
func (a *AnalyticsReportsAPIService) AnalyticsReportsInstancesGetToManyRelatedExecute(r ApiAnalyticsReportsInstancesGetToManyRelatedRequest) (*AnalyticsReportInstancesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnalyticsReportInstancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsReportsAPIService.AnalyticsReportsInstancesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analyticsReports/{id}/instances"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[granularity]", r.filterGranularity, "form", "csv")
	}
	if r.filterProcessingDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[processingDate]", r.filterProcessingDate, "form", "csv")
	}
	if r.fieldsAnalyticsReportInstances != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[analyticsReportInstances]", r.fieldsAnalyticsReportInstances, "form", "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
