/*
Hippo.Web

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hippo-openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// StorageApiService StorageApi service
type StorageApiService service

type ApiApiStorageGetRequest struct {
	ctx context.Context
	ApiService *StorageApiService
	searchText *string
	pageIndex *int32
	pageSize *int32
	isSortedAscending *bool
}

func (r ApiApiStorageGetRequest) SearchText(searchText string) ApiApiStorageGetRequest {
	r.searchText = &searchText
	return r
}

func (r ApiApiStorageGetRequest) PageIndex(pageIndex int32) ApiApiStorageGetRequest {
	r.pageIndex = &pageIndex
	return r
}

func (r ApiApiStorageGetRequest) PageSize(pageSize int32) ApiApiStorageGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiApiStorageGetRequest) IsSortedAscending(isSortedAscending bool) ApiApiStorageGetRequest {
	r.isSortedAscending = &isSortedAscending
	return r
}

func (r ApiApiStorageGetRequest) Execute() (*StringPage, *http.Response, error) {
	return r.ApiService.ApiStorageGetExecute(r)
}

/*
ApiStorageGet Method for ApiStorageGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiStorageGetRequest
*/
func (a *StorageApiService) ApiStorageGet(ctx context.Context) ApiApiStorageGetRequest {
	return ApiApiStorageGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringPage
func (a *StorageApiService) ApiStorageGetExecute(r ApiApiStorageGetRequest) (*StringPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageApiService.ApiStorageGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/storage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchText != nil {
		localVarQueryParams.Add("searchText", parameterToString(*r.searchText, ""))
	}
	if r.pageIndex != nil {
		localVarQueryParams.Add("pageIndex", parameterToString(*r.pageIndex, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.isSortedAscending != nil {
		localVarQueryParams.Add("IsSortedAscending", parameterToString(*r.isSortedAscending, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
