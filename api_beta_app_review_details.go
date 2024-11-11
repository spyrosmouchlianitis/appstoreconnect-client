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


// BetaAppReviewDetailsAPIService BetaAppReviewDetailsAPI service
type BetaAppReviewDetailsAPIService service

type ApiBetaAppReviewDetailsAppGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewDetailsAPIService
	id string
	fieldsApps *[]string
}

// the fields to include for returned resources of type apps
func (r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsAppGetToOneRelatedRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) Execute() (*AppWithoutIncludesResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewDetailsAppGetToOneRelatedExecute(r)
}

/*
BetaAppReviewDetailsAppGetToOneRelated Method for BetaAppReviewDetailsAppGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaAppReviewDetailsAppGetToOneRelatedRequest
*/
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsAppGetToOneRelated(ctx context.Context, id string) ApiBetaAppReviewDetailsAppGetToOneRelatedRequest {
	return ApiBetaAppReviewDetailsAppGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppWithoutIncludesResponse
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsAppGetToOneRelatedExecute(r ApiBetaAppReviewDetailsAppGetToOneRelatedRequest) (*AppWithoutIncludesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppWithoutIncludesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsAPIService.BetaAppReviewDetailsAppGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}/app"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "form", "csv")
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

type ApiBetaAppReviewDetailsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewDetailsAPIService
	filterApp *[]string
	fieldsBetaAppReviewDetails *[]string
	fieldsApps *[]string
	limit *int32
	include *[]string
}

// filter by id(s) of related &#39;app&#39;
func (r ApiBetaAppReviewDetailsGetCollectionRequest) FilterApp(filterApp []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.filterApp = &filterApp
	return r
}

// the fields to include for returned resources of type betaAppReviewDetails
func (r ApiBetaAppReviewDetailsGetCollectionRequest) FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.fieldsBetaAppReviewDetails = &fieldsBetaAppReviewDetails
	return r
}

// the fields to include for returned resources of type apps
func (r ApiBetaAppReviewDetailsGetCollectionRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.fieldsApps = &fieldsApps
	return r
}

// maximum resources per page
func (r ApiBetaAppReviewDetailsGetCollectionRequest) Limit(limit int32) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaAppReviewDetailsGetCollectionRequest) Include(include []string) ApiBetaAppReviewDetailsGetCollectionRequest {
	r.include = &include
	return r
}

func (r ApiBetaAppReviewDetailsGetCollectionRequest) Execute() (*BetaAppReviewDetailsResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewDetailsGetCollectionExecute(r)
}

/*
BetaAppReviewDetailsGetCollection Method for BetaAppReviewDetailsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBetaAppReviewDetailsGetCollectionRequest
*/
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsGetCollection(ctx context.Context) ApiBetaAppReviewDetailsGetCollectionRequest {
	return ApiBetaAppReviewDetailsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaAppReviewDetailsResponse
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsGetCollectionExecute(r ApiBetaAppReviewDetailsGetCollectionRequest) (*BetaAppReviewDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsAPIService.BetaAppReviewDetailsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filterApp == nil {
		return localVarReturnValue, nil, reportError("filterApp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[app]", r.filterApp, "form", "csv")
	if r.fieldsBetaAppReviewDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaAppReviewDetails]", r.fieldsBetaAppReviewDetails, "form", "csv")
	}
	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "form", "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
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

type ApiBetaAppReviewDetailsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewDetailsAPIService
	id string
	fieldsBetaAppReviewDetails *[]string
	fieldsApps *[]string
	include *[]string
}

// the fields to include for returned resources of type betaAppReviewDetails
func (r ApiBetaAppReviewDetailsGetInstanceRequest) FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.fieldsBetaAppReviewDetails = &fieldsBetaAppReviewDetails
	return r
}

// the fields to include for returned resources of type apps
func (r ApiBetaAppReviewDetailsGetInstanceRequest) FieldsApps(fieldsApps []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.fieldsApps = &fieldsApps
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaAppReviewDetailsGetInstanceRequest) Include(include []string) ApiBetaAppReviewDetailsGetInstanceRequest {
	r.include = &include
	return r
}

func (r ApiBetaAppReviewDetailsGetInstanceRequest) Execute() (*BetaAppReviewDetailResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewDetailsGetInstanceExecute(r)
}

/*
BetaAppReviewDetailsGetInstance Method for BetaAppReviewDetailsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaAppReviewDetailsGetInstanceRequest
*/
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsGetInstance(ctx context.Context, id string) ApiBetaAppReviewDetailsGetInstanceRequest {
	return ApiBetaAppReviewDetailsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaAppReviewDetailResponse
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsGetInstanceExecute(r ApiBetaAppReviewDetailsGetInstanceRequest) (*BetaAppReviewDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsAPIService.BetaAppReviewDetailsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBetaAppReviewDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaAppReviewDetails]", r.fieldsBetaAppReviewDetails, "form", "csv")
	}
	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "form", "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
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

type ApiBetaAppReviewDetailsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewDetailsAPIService
	id string
	betaAppReviewDetailUpdateRequest *BetaAppReviewDetailUpdateRequest
}

// BetaAppReviewDetail representation
func (r ApiBetaAppReviewDetailsUpdateInstanceRequest) BetaAppReviewDetailUpdateRequest(betaAppReviewDetailUpdateRequest BetaAppReviewDetailUpdateRequest) ApiBetaAppReviewDetailsUpdateInstanceRequest {
	r.betaAppReviewDetailUpdateRequest = &betaAppReviewDetailUpdateRequest
	return r
}

func (r ApiBetaAppReviewDetailsUpdateInstanceRequest) Execute() (*BetaAppReviewDetailResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewDetailsUpdateInstanceExecute(r)
}

/*
BetaAppReviewDetailsUpdateInstance Method for BetaAppReviewDetailsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaAppReviewDetailsUpdateInstanceRequest
*/
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsUpdateInstance(ctx context.Context, id string) ApiBetaAppReviewDetailsUpdateInstanceRequest {
	return ApiBetaAppReviewDetailsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaAppReviewDetailResponse
func (a *BetaAppReviewDetailsAPIService) BetaAppReviewDetailsUpdateInstanceExecute(r ApiBetaAppReviewDetailsUpdateInstanceRequest) (*BetaAppReviewDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewDetailsAPIService.BetaAppReviewDetailsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaAppReviewDetailUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaAppReviewDetailUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.betaAppReviewDetailUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 422 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
