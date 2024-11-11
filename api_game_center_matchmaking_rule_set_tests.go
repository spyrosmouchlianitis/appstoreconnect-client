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
)


// GameCenterMatchmakingRuleSetTestsAPIService GameCenterMatchmakingRuleSetTestsAPI service
type GameCenterMatchmakingRuleSetTestsAPIService service

type ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *GameCenterMatchmakingRuleSetTestsAPIService
	gameCenterMatchmakingRuleSetTestCreateRequest *GameCenterMatchmakingRuleSetTestCreateRequest
}

// GameCenterMatchmakingRuleSetTest representation
func (r ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest) GameCenterMatchmakingRuleSetTestCreateRequest(gameCenterMatchmakingRuleSetTestCreateRequest GameCenterMatchmakingRuleSetTestCreateRequest) ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest {
	r.gameCenterMatchmakingRuleSetTestCreateRequest = &gameCenterMatchmakingRuleSetTestCreateRequest
	return r
}

func (r ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest) Execute() (*GameCenterMatchmakingRuleSetTestResponse, *http.Response, error) {
	return r.ApiService.GameCenterMatchmakingRuleSetTestsCreateInstanceExecute(r)
}

/*
GameCenterMatchmakingRuleSetTestsCreateInstance Method for GameCenterMatchmakingRuleSetTestsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest
*/
func (a *GameCenterMatchmakingRuleSetTestsAPIService) GameCenterMatchmakingRuleSetTestsCreateInstance(ctx context.Context) ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest {
	return ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GameCenterMatchmakingRuleSetTestResponse
func (a *GameCenterMatchmakingRuleSetTestsAPIService) GameCenterMatchmakingRuleSetTestsCreateInstanceExecute(r ApiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest) (*GameCenterMatchmakingRuleSetTestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameCenterMatchmakingRuleSetTestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterMatchmakingRuleSetTestsAPIService.GameCenterMatchmakingRuleSetTestsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterMatchmakingRuleSetTests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gameCenterMatchmakingRuleSetTestCreateRequest == nil {
		return localVarReturnValue, nil, reportError("gameCenterMatchmakingRuleSetTestCreateRequest is required and must be specified")
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
	localVarPostBody = r.gameCenterMatchmakingRuleSetTestCreateRequest
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
