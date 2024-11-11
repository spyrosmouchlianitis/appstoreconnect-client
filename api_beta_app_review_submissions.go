/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreconnect-client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BetaAppReviewSubmissionsAPIService BetaAppReviewSubmissionsAPI service
type BetaAppReviewSubmissionsAPIService service

type ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewSubmissionsAPIService
	id string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type builds
func (r ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest) Execute() (*BuildWithoutIncludesResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewSubmissionsBuildGetToOneRelatedExecute(r)
}

/*
BetaAppReviewSubmissionsBuildGetToOneRelated Method for BetaAppReviewSubmissionsBuildGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest
*/
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsBuildGetToOneRelated(ctx context.Context, id string) ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest {
	return ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildWithoutIncludesResponse
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsBuildGetToOneRelatedExecute(r ApiBetaAppReviewSubmissionsBuildGetToOneRelatedRequest) (*BuildWithoutIncludesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildWithoutIncludesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewSubmissionsAPIService.BetaAppReviewSubmissionsBuildGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewSubmissions/{id}/build"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "form", "csv")
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

type ApiBetaAppReviewSubmissionsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewSubmissionsAPIService
	betaAppReviewSubmissionCreateRequest *BetaAppReviewSubmissionCreateRequest
}

// BetaAppReviewSubmission representation
func (r ApiBetaAppReviewSubmissionsCreateInstanceRequest) BetaAppReviewSubmissionCreateRequest(betaAppReviewSubmissionCreateRequest BetaAppReviewSubmissionCreateRequest) ApiBetaAppReviewSubmissionsCreateInstanceRequest {
	r.betaAppReviewSubmissionCreateRequest = &betaAppReviewSubmissionCreateRequest
	return r
}

func (r ApiBetaAppReviewSubmissionsCreateInstanceRequest) Execute() (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewSubmissionsCreateInstanceExecute(r)
}

/*
BetaAppReviewSubmissionsCreateInstance Method for BetaAppReviewSubmissionsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBetaAppReviewSubmissionsCreateInstanceRequest
*/
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsCreateInstance(ctx context.Context) ApiBetaAppReviewSubmissionsCreateInstanceRequest {
	return ApiBetaAppReviewSubmissionsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaAppReviewSubmissionResponse
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsCreateInstanceExecute(r ApiBetaAppReviewSubmissionsCreateInstanceRequest) (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewSubmissionsAPIService.BetaAppReviewSubmissionsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewSubmissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaAppReviewSubmissionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("betaAppReviewSubmissionCreateRequest is required and must be specified")
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
	localVarPostBody = r.betaAppReviewSubmissionCreateRequest
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

type ApiBetaAppReviewSubmissionsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewSubmissionsAPIService
	filterBuild *[]string
	filterBetaReviewState *[]string
	fieldsBetaAppReviewSubmissions *[]string
	fieldsBuilds *[]string
	limit *int32
	include *[]string
}

// filter by id(s) of related &#39;build&#39;
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) FilterBuild(filterBuild []string) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.filterBuild = &filterBuild
	return r
}

// filter by attribute &#39;betaReviewState&#39;
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) FilterBetaReviewState(filterBetaReviewState []string) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.filterBetaReviewState = &filterBetaReviewState
	return r
}

// the fields to include for returned resources of type betaAppReviewSubmissions
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions []string) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.fieldsBetaAppReviewSubmissions = &fieldsBetaAppReviewSubmissions
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

// maximum resources per page
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) Limit(limit int32) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) Include(include []string) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	r.include = &include
	return r
}

func (r ApiBetaAppReviewSubmissionsGetCollectionRequest) Execute() (*BetaAppReviewSubmissionsResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewSubmissionsGetCollectionExecute(r)
}

/*
BetaAppReviewSubmissionsGetCollection Method for BetaAppReviewSubmissionsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBetaAppReviewSubmissionsGetCollectionRequest
*/
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsGetCollection(ctx context.Context) ApiBetaAppReviewSubmissionsGetCollectionRequest {
	return ApiBetaAppReviewSubmissionsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaAppReviewSubmissionsResponse
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsGetCollectionExecute(r ApiBetaAppReviewSubmissionsGetCollectionRequest) (*BetaAppReviewSubmissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewSubmissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewSubmissionsAPIService.BetaAppReviewSubmissionsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewSubmissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.filterBuild == nil {
		return localVarReturnValue, nil, reportError("filterBuild is required and must be specified")
	}

	if r.filterBetaReviewState != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[betaReviewState]", r.filterBetaReviewState, "form", "csv")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter[build]", r.filterBuild, "form", "csv")
	if r.fieldsBetaAppReviewSubmissions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaAppReviewSubmissions]", r.fieldsBetaAppReviewSubmissions, "form", "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "form", "csv")
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

type ApiBetaAppReviewSubmissionsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BetaAppReviewSubmissionsAPIService
	id string
	fieldsBetaAppReviewSubmissions *[]string
	fieldsBuilds *[]string
	include *[]string
}

// the fields to include for returned resources of type betaAppReviewSubmissions
func (r ApiBetaAppReviewSubmissionsGetInstanceRequest) FieldsBetaAppReviewSubmissions(fieldsBetaAppReviewSubmissions []string) ApiBetaAppReviewSubmissionsGetInstanceRequest {
	r.fieldsBetaAppReviewSubmissions = &fieldsBetaAppReviewSubmissions
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBetaAppReviewSubmissionsGetInstanceRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaAppReviewSubmissionsGetInstanceRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaAppReviewSubmissionsGetInstanceRequest) Include(include []string) ApiBetaAppReviewSubmissionsGetInstanceRequest {
	r.include = &include
	return r
}

func (r ApiBetaAppReviewSubmissionsGetInstanceRequest) Execute() (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	return r.ApiService.BetaAppReviewSubmissionsGetInstanceExecute(r)
}

/*
BetaAppReviewSubmissionsGetInstance Method for BetaAppReviewSubmissionsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaAppReviewSubmissionsGetInstanceRequest
*/
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsGetInstance(ctx context.Context, id string) ApiBetaAppReviewSubmissionsGetInstanceRequest {
	return ApiBetaAppReviewSubmissionsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaAppReviewSubmissionResponse
func (a *BetaAppReviewSubmissionsAPIService) BetaAppReviewSubmissionsGetInstanceExecute(r ApiBetaAppReviewSubmissionsGetInstanceRequest) (*BetaAppReviewSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaAppReviewSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAppReviewSubmissionsAPIService.BetaAppReviewSubmissionsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaAppReviewSubmissions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBetaAppReviewSubmissions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaAppReviewSubmissions]", r.fieldsBetaAppReviewSubmissions, "form", "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "form", "csv")
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
