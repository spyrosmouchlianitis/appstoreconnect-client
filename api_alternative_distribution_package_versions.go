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


// AlternativeDistributionPackageVersionsAPIService AlternativeDistributionPackageVersionsAPI service
type AlternativeDistributionPackageVersionsAPIService service

type ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AlternativeDistributionPackageVersionsAPIService
	id string
	fieldsAlternativeDistributionPackageDeltas *[]string
	limit *int32
}

// the fields to include for returned resources of type alternativeDistributionPackageDeltas
func (r ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest) FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas []string) ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest {
	r.fieldsAlternativeDistributionPackageDeltas = &fieldsAlternativeDistributionPackageDeltas
	return r
}

// maximum resources per page
func (r ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest) Limit(limit int32) ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest) Execute() (*AlternativeDistributionPackageDeltasResponse, *http.Response, error) {
	return r.ApiService.AlternativeDistributionPackageVersionsDeltasGetToManyRelatedExecute(r)
}

/*
AlternativeDistributionPackageVersionsDeltasGetToManyRelated Method for AlternativeDistributionPackageVersionsDeltasGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest
*/
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsDeltasGetToManyRelated(ctx context.Context, id string) ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest {
	return ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AlternativeDistributionPackageDeltasResponse
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsDeltasGetToManyRelatedExecute(r ApiAlternativeDistributionPackageVersionsDeltasGetToManyRelatedRequest) (*AlternativeDistributionPackageDeltasResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlternativeDistributionPackageDeltasResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternativeDistributionPackageVersionsAPIService.AlternativeDistributionPackageVersionsDeltasGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alternativeDistributionPackageVersions/{id}/deltas"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAlternativeDistributionPackageDeltas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[alternativeDistributionPackageDeltas]", r.fieldsAlternativeDistributionPackageDeltas, "form", "csv")
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

type ApiAlternativeDistributionPackageVersionsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AlternativeDistributionPackageVersionsAPIService
	id string
	fieldsAlternativeDistributionPackageVersions *[]string
	fieldsAlternativeDistributionPackageVariants *[]string
	fieldsAlternativeDistributionPackageDeltas *[]string
	include *[]string
	limitDeltas *int32
	limitVariants *int32
}

// the fields to include for returned resources of type alternativeDistributionPackageVersions
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) FieldsAlternativeDistributionPackageVersions(fieldsAlternativeDistributionPackageVersions []string) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.fieldsAlternativeDistributionPackageVersions = &fieldsAlternativeDistributionPackageVersions
	return r
}

// the fields to include for returned resources of type alternativeDistributionPackageVariants
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants []string) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.fieldsAlternativeDistributionPackageVariants = &fieldsAlternativeDistributionPackageVariants
	return r
}

// the fields to include for returned resources of type alternativeDistributionPackageDeltas
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) FieldsAlternativeDistributionPackageDeltas(fieldsAlternativeDistributionPackageDeltas []string) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.fieldsAlternativeDistributionPackageDeltas = &fieldsAlternativeDistributionPackageDeltas
	return r
}

// comma-separated list of relationships to include
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) Include(include []string) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.include = &include
	return r
}

// maximum number of related deltas returned (when they are included)
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) LimitDeltas(limitDeltas int32) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.limitDeltas = &limitDeltas
	return r
}

// maximum number of related variants returned (when they are included)
func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) LimitVariants(limitVariants int32) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	r.limitVariants = &limitVariants
	return r
}

func (r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) Execute() (*AlternativeDistributionPackageVersionResponse, *http.Response, error) {
	return r.ApiService.AlternativeDistributionPackageVersionsGetInstanceExecute(r)
}

/*
AlternativeDistributionPackageVersionsGetInstance Method for AlternativeDistributionPackageVersionsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAlternativeDistributionPackageVersionsGetInstanceRequest
*/
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsGetInstance(ctx context.Context, id string) ApiAlternativeDistributionPackageVersionsGetInstanceRequest {
	return ApiAlternativeDistributionPackageVersionsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AlternativeDistributionPackageVersionResponse
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsGetInstanceExecute(r ApiAlternativeDistributionPackageVersionsGetInstanceRequest) (*AlternativeDistributionPackageVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlternativeDistributionPackageVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternativeDistributionPackageVersionsAPIService.AlternativeDistributionPackageVersionsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alternativeDistributionPackageVersions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAlternativeDistributionPackageVersions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[alternativeDistributionPackageVersions]", r.fieldsAlternativeDistributionPackageVersions, "form", "csv")
	}
	if r.fieldsAlternativeDistributionPackageVariants != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[alternativeDistributionPackageVariants]", r.fieldsAlternativeDistributionPackageVariants, "form", "csv")
	}
	if r.fieldsAlternativeDistributionPackageDeltas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[alternativeDistributionPackageDeltas]", r.fieldsAlternativeDistributionPackageDeltas, "form", "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
	}
	if r.limitDeltas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[deltas]", r.limitDeltas, "form", "")
	}
	if r.limitVariants != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[variants]", r.limitVariants, "form", "")
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

type ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AlternativeDistributionPackageVersionsAPIService
	id string
	fieldsAlternativeDistributionPackageVariants *[]string
	limit *int32
}

// the fields to include for returned resources of type alternativeDistributionPackageVariants
func (r ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest) FieldsAlternativeDistributionPackageVariants(fieldsAlternativeDistributionPackageVariants []string) ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest {
	r.fieldsAlternativeDistributionPackageVariants = &fieldsAlternativeDistributionPackageVariants
	return r
}

// maximum resources per page
func (r ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest) Limit(limit int32) ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest) Execute() (*AlternativeDistributionPackageVariantsResponse, *http.Response, error) {
	return r.ApiService.AlternativeDistributionPackageVersionsVariantsGetToManyRelatedExecute(r)
}

/*
AlternativeDistributionPackageVersionsVariantsGetToManyRelated Method for AlternativeDistributionPackageVersionsVariantsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest
*/
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsVariantsGetToManyRelated(ctx context.Context, id string) ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest {
	return ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AlternativeDistributionPackageVariantsResponse
func (a *AlternativeDistributionPackageVersionsAPIService) AlternativeDistributionPackageVersionsVariantsGetToManyRelatedExecute(r ApiAlternativeDistributionPackageVersionsVariantsGetToManyRelatedRequest) (*AlternativeDistributionPackageVariantsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlternativeDistributionPackageVariantsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlternativeDistributionPackageVersionsAPIService.AlternativeDistributionPackageVersionsVariantsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/alternativeDistributionPackageVersions/{id}/variants"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAlternativeDistributionPackageVariants != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[alternativeDistributionPackageVariants]", r.fieldsAlternativeDistributionPackageVariants, "form", "csv")
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
