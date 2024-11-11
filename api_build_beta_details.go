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


// BuildBetaDetailsAPIService BuildBetaDetailsAPI service
type BuildBetaDetailsAPIService service

type ApiBuildBetaDetailsBuildGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsAPIService
	id string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type builds
func (r ApiBuildBetaDetailsBuildGetToOneRelatedRequest) FieldsBuilds(fieldsBuilds []string) ApiBuildBetaDetailsBuildGetToOneRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r ApiBuildBetaDetailsBuildGetToOneRelatedRequest) Execute() (*BuildWithoutIncludesResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsBuildGetToOneRelatedExecute(r)
}

/*
BuildBetaDetailsBuildGetToOneRelated Method for BuildBetaDetailsBuildGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBuildBetaDetailsBuildGetToOneRelatedRequest
*/
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsBuildGetToOneRelated(ctx context.Context, id string) ApiBuildBetaDetailsBuildGetToOneRelatedRequest {
	return ApiBuildBetaDetailsBuildGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildWithoutIncludesResponse
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsBuildGetToOneRelatedExecute(r ApiBuildBetaDetailsBuildGetToOneRelatedRequest) (*BuildWithoutIncludesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildWithoutIncludesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsAPIService.BuildBetaDetailsBuildGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}/build"
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

type ApiBuildBetaDetailsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsAPIService
	filterBuild *[]string
	filterId *[]string
	fieldsBuildBetaDetails *[]string
	fieldsBuilds *[]string
	limit *int32
	include *[]string
}

// filter by id(s) of related &#39;build&#39;
func (r ApiBuildBetaDetailsGetCollectionRequest) FilterBuild(filterBuild []string) ApiBuildBetaDetailsGetCollectionRequest {
	r.filterBuild = &filterBuild
	return r
}

// filter by id(s)
func (r ApiBuildBetaDetailsGetCollectionRequest) FilterId(filterId []string) ApiBuildBetaDetailsGetCollectionRequest {
	r.filterId = &filterId
	return r
}

// the fields to include for returned resources of type buildBetaDetails
func (r ApiBuildBetaDetailsGetCollectionRequest) FieldsBuildBetaDetails(fieldsBuildBetaDetails []string) ApiBuildBetaDetailsGetCollectionRequest {
	r.fieldsBuildBetaDetails = &fieldsBuildBetaDetails
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBuildBetaDetailsGetCollectionRequest) FieldsBuilds(fieldsBuilds []string) ApiBuildBetaDetailsGetCollectionRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

// maximum resources per page
func (r ApiBuildBetaDetailsGetCollectionRequest) Limit(limit int32) ApiBuildBetaDetailsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiBuildBetaDetailsGetCollectionRequest) Include(include []string) ApiBuildBetaDetailsGetCollectionRequest {
	r.include = &include
	return r
}

func (r ApiBuildBetaDetailsGetCollectionRequest) Execute() (*BuildBetaDetailsResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsGetCollectionExecute(r)
}

/*
BuildBetaDetailsGetCollection Method for BuildBetaDetailsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBuildBetaDetailsGetCollectionRequest
*/
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsGetCollection(ctx context.Context) ApiBuildBetaDetailsGetCollectionRequest {
	return ApiBuildBetaDetailsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BuildBetaDetailsResponse
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsGetCollectionExecute(r ApiBuildBetaDetailsGetCollectionRequest) (*BuildBetaDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsAPIService.BuildBetaDetailsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBuild != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[build]", r.filterBuild, "form", "csv")
	}
	if r.filterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[id]", r.filterId, "form", "csv")
	}
	if r.fieldsBuildBetaDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[buildBetaDetails]", r.fieldsBuildBetaDetails, "form", "csv")
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

type ApiBuildBetaDetailsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsAPIService
	id string
	fieldsBuildBetaDetails *[]string
	fieldsBuilds *[]string
	include *[]string
}

// the fields to include for returned resources of type buildBetaDetails
func (r ApiBuildBetaDetailsGetInstanceRequest) FieldsBuildBetaDetails(fieldsBuildBetaDetails []string) ApiBuildBetaDetailsGetInstanceRequest {
	r.fieldsBuildBetaDetails = &fieldsBuildBetaDetails
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBuildBetaDetailsGetInstanceRequest) FieldsBuilds(fieldsBuilds []string) ApiBuildBetaDetailsGetInstanceRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

// comma-separated list of relationships to include
func (r ApiBuildBetaDetailsGetInstanceRequest) Include(include []string) ApiBuildBetaDetailsGetInstanceRequest {
	r.include = &include
	return r
}

func (r ApiBuildBetaDetailsGetInstanceRequest) Execute() (*BuildBetaDetailResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsGetInstanceExecute(r)
}

/*
BuildBetaDetailsGetInstance Method for BuildBetaDetailsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBuildBetaDetailsGetInstanceRequest
*/
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsGetInstance(ctx context.Context, id string) ApiBuildBetaDetailsGetInstanceRequest {
	return ApiBuildBetaDetailsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildBetaDetailResponse
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsGetInstanceExecute(r ApiBuildBetaDetailsGetInstanceRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsAPIService.BuildBetaDetailsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBuildBetaDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[buildBetaDetails]", r.fieldsBuildBetaDetails, "form", "csv")
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

type ApiBuildBetaDetailsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *BuildBetaDetailsAPIService
	id string
	buildBetaDetailUpdateRequest *BuildBetaDetailUpdateRequest
}

// BuildBetaDetail representation
func (r ApiBuildBetaDetailsUpdateInstanceRequest) BuildBetaDetailUpdateRequest(buildBetaDetailUpdateRequest BuildBetaDetailUpdateRequest) ApiBuildBetaDetailsUpdateInstanceRequest {
	r.buildBetaDetailUpdateRequest = &buildBetaDetailUpdateRequest
	return r
}

func (r ApiBuildBetaDetailsUpdateInstanceRequest) Execute() (*BuildBetaDetailResponse, *http.Response, error) {
	return r.ApiService.BuildBetaDetailsUpdateInstanceExecute(r)
}

/*
BuildBetaDetailsUpdateInstance Method for BuildBetaDetailsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBuildBetaDetailsUpdateInstanceRequest
*/
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsUpdateInstance(ctx context.Context, id string) ApiBuildBetaDetailsUpdateInstanceRequest {
	return ApiBuildBetaDetailsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BuildBetaDetailResponse
func (a *BuildBetaDetailsAPIService) BuildBetaDetailsUpdateInstanceExecute(r ApiBuildBetaDetailsUpdateInstanceRequest) (*BuildBetaDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BuildBetaDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildBetaDetailsAPIService.BuildBetaDetailsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buildBetaDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.buildBetaDetailUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("buildBetaDetailUpdateRequest is required and must be specified")
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
	localVarPostBody = r.buildBetaDetailUpdateRequest
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
