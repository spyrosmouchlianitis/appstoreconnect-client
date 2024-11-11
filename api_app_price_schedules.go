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


// AppPriceSchedulesAPIService AppPriceSchedulesAPI service
type AppPriceSchedulesAPIService service

type ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppPriceSchedulesAPIService
	id string
	filterStartDate *[]string
	filterEndDate *[]string
	filterTerritory *[]string
	fieldsAppPrices *[]string
	fieldsAppPricePoints *[]string
	fieldsTerritories *[]string
	limit *int32
	include *[]string
}

// filter by attribute &#39;startDate&#39;
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FilterStartDate(filterStartDate []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.filterStartDate = &filterStartDate
	return r
}

// filter by attribute &#39;endDate&#39;
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FilterEndDate(filterEndDate []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.filterEndDate = &filterEndDate
	return r
}

// filter by id(s) of related &#39;territory&#39;
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type appPrices
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsAppPrices(fieldsAppPrices []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsAppPrices = &fieldsAppPrices
	return r
}

// the fields to include for returned resources of type appPricePoints
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsAppPricePoints(fieldsAppPricePoints []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsAppPricePoints = &fieldsAppPricePoints
	return r
}

// the fields to include for returned resources of type territories
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) Limit(limit int32) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) Include(include []string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) Execute() (*AppPricesV2Response, *http.Response, error) {
	return r.ApiService.AppPriceSchedulesAutomaticPricesGetToManyRelatedExecute(r)
}

/*
AppPriceSchedulesAutomaticPricesGetToManyRelated Method for AppPriceSchedulesAutomaticPricesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest
*/
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesAutomaticPricesGetToManyRelated(ctx context.Context, id string) ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest {
	return ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppPricesV2Response
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesAutomaticPricesGetToManyRelatedExecute(r ApiAppPriceSchedulesAutomaticPricesGetToManyRelatedRequest) (*AppPricesV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppPricesV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceSchedulesAPIService.AppPriceSchedulesAutomaticPricesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceSchedules/{id}/automaticPrices"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterStartDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[startDate]", r.filterStartDate, "form", "csv")
	}
	if r.filterEndDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[endDate]", r.filterEndDate, "form", "csv")
	}
	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "form", "csv")
	}
	if r.fieldsAppPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPrices]", r.fieldsAppPrices, "form", "csv")
	}
	if r.fieldsAppPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPricePoints]", r.fieldsAppPricePoints, "form", "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "form", "csv")
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

type ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *AppPriceSchedulesAPIService
	id string
	fieldsTerritories *[]string
}

// the fields to include for returned resources of type territories
func (r ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

func (r ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest) Execute() (*TerritoryResponse, *http.Response, error) {
	return r.ApiService.AppPriceSchedulesBaseTerritoryGetToOneRelatedExecute(r)
}

/*
AppPriceSchedulesBaseTerritoryGetToOneRelated Method for AppPriceSchedulesBaseTerritoryGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest
*/
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesBaseTerritoryGetToOneRelated(ctx context.Context, id string) ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest {
	return ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TerritoryResponse
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesBaseTerritoryGetToOneRelatedExecute(r ApiAppPriceSchedulesBaseTerritoryGetToOneRelatedRequest) (*TerritoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TerritoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceSchedulesAPIService.AppPriceSchedulesBaseTerritoryGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceSchedules/{id}/baseTerritory"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "form", "csv")
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

type ApiAppPriceSchedulesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppPriceSchedulesAPIService
	appPriceScheduleCreateRequest *AppPriceScheduleCreateRequest
}

// AppPriceSchedule representation
func (r ApiAppPriceSchedulesCreateInstanceRequest) AppPriceScheduleCreateRequest(appPriceScheduleCreateRequest AppPriceScheduleCreateRequest) ApiAppPriceSchedulesCreateInstanceRequest {
	r.appPriceScheduleCreateRequest = &appPriceScheduleCreateRequest
	return r
}

func (r ApiAppPriceSchedulesCreateInstanceRequest) Execute() (*AppPriceScheduleResponse, *http.Response, error) {
	return r.ApiService.AppPriceSchedulesCreateInstanceExecute(r)
}

/*
AppPriceSchedulesCreateInstance Method for AppPriceSchedulesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppPriceSchedulesCreateInstanceRequest
*/
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesCreateInstance(ctx context.Context) ApiAppPriceSchedulesCreateInstanceRequest {
	return ApiAppPriceSchedulesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppPriceScheduleResponse
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesCreateInstanceExecute(r ApiAppPriceSchedulesCreateInstanceRequest) (*AppPriceScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppPriceScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceSchedulesAPIService.AppPriceSchedulesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceSchedules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appPriceScheduleCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appPriceScheduleCreateRequest is required and must be specified")
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
	localVarPostBody = r.appPriceScheduleCreateRequest
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

type ApiAppPriceSchedulesGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppPriceSchedulesAPIService
	id string
	fieldsAppPriceSchedules *[]string
	fieldsTerritories *[]string
	fieldsAppPrices *[]string
	include *[]string
	limitAutomaticPrices *int32
	limitManualPrices *int32
}

// the fields to include for returned resources of type appPriceSchedules
func (r ApiAppPriceSchedulesGetInstanceRequest) FieldsAppPriceSchedules(fieldsAppPriceSchedules []string) ApiAppPriceSchedulesGetInstanceRequest {
	r.fieldsAppPriceSchedules = &fieldsAppPriceSchedules
	return r
}

// the fields to include for returned resources of type territories
func (r ApiAppPriceSchedulesGetInstanceRequest) FieldsTerritories(fieldsTerritories []string) ApiAppPriceSchedulesGetInstanceRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// the fields to include for returned resources of type appPrices
func (r ApiAppPriceSchedulesGetInstanceRequest) FieldsAppPrices(fieldsAppPrices []string) ApiAppPriceSchedulesGetInstanceRequest {
	r.fieldsAppPrices = &fieldsAppPrices
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceSchedulesGetInstanceRequest) Include(include []string) ApiAppPriceSchedulesGetInstanceRequest {
	r.include = &include
	return r
}

// maximum number of related automaticPrices returned (when they are included)
func (r ApiAppPriceSchedulesGetInstanceRequest) LimitAutomaticPrices(limitAutomaticPrices int32) ApiAppPriceSchedulesGetInstanceRequest {
	r.limitAutomaticPrices = &limitAutomaticPrices
	return r
}

// maximum number of related manualPrices returned (when they are included)
func (r ApiAppPriceSchedulesGetInstanceRequest) LimitManualPrices(limitManualPrices int32) ApiAppPriceSchedulesGetInstanceRequest {
	r.limitManualPrices = &limitManualPrices
	return r
}

func (r ApiAppPriceSchedulesGetInstanceRequest) Execute() (*AppPriceScheduleResponse, *http.Response, error) {
	return r.ApiService.AppPriceSchedulesGetInstanceExecute(r)
}

/*
AppPriceSchedulesGetInstance Method for AppPriceSchedulesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppPriceSchedulesGetInstanceRequest
*/
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesGetInstance(ctx context.Context, id string) ApiAppPriceSchedulesGetInstanceRequest {
	return ApiAppPriceSchedulesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppPriceScheduleResponse
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesGetInstanceExecute(r ApiAppPriceSchedulesGetInstanceRequest) (*AppPriceScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppPriceScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceSchedulesAPIService.AppPriceSchedulesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceSchedules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppPriceSchedules != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPriceSchedules]", r.fieldsAppPriceSchedules, "form", "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "form", "csv")
	}
	if r.fieldsAppPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPrices]", r.fieldsAppPrices, "form", "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
	}
	if r.limitAutomaticPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[automaticPrices]", r.limitAutomaticPrices, "form", "")
	}
	if r.limitManualPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[manualPrices]", r.limitManualPrices, "form", "")
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

type ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppPriceSchedulesAPIService
	id string
	filterStartDate *[]string
	filterEndDate *[]string
	filterTerritory *[]string
	fieldsAppPrices *[]string
	fieldsAppPricePoints *[]string
	fieldsTerritories *[]string
	limit *int32
	include *[]string
}

// filter by attribute &#39;startDate&#39;
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FilterStartDate(filterStartDate []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.filterStartDate = &filterStartDate
	return r
}

// filter by attribute &#39;endDate&#39;
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FilterEndDate(filterEndDate []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.filterEndDate = &filterEndDate
	return r
}

// filter by id(s) of related &#39;territory&#39;
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type appPrices
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FieldsAppPrices(fieldsAppPrices []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsAppPrices = &fieldsAppPrices
	return r
}

// the fields to include for returned resources of type appPricePoints
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FieldsAppPricePoints(fieldsAppPricePoints []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsAppPricePoints = &fieldsAppPricePoints
	return r
}

// the fields to include for returned resources of type territories
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) Limit(limit int32) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) Include(include []string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) Execute() (*AppPricesV2Response, *http.Response, error) {
	return r.ApiService.AppPriceSchedulesManualPricesGetToManyRelatedExecute(r)
}

/*
AppPriceSchedulesManualPricesGetToManyRelated Method for AppPriceSchedulesManualPricesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest
*/
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesManualPricesGetToManyRelated(ctx context.Context, id string) ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest {
	return ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppPricesV2Response
func (a *AppPriceSchedulesAPIService) AppPriceSchedulesManualPricesGetToManyRelatedExecute(r ApiAppPriceSchedulesManualPricesGetToManyRelatedRequest) (*AppPricesV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppPricesV2Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceSchedulesAPIService.AppPriceSchedulesManualPricesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceSchedules/{id}/manualPrices"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterStartDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[startDate]", r.filterStartDate, "form", "csv")
	}
	if r.filterEndDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[endDate]", r.filterEndDate, "form", "csv")
	}
	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "form", "csv")
	}
	if r.fieldsAppPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPrices]", r.fieldsAppPrices, "form", "csv")
	}
	if r.fieldsAppPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPricePoints]", r.fieldsAppPricePoints, "form", "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "form", "csv")
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
