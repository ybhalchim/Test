/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type AsmAPI interface {
	/*
			Create Update subnet and ranges for Automated Scope Management.

			Use this method to update the subnet and range for Automated Scope Management.
		The __ASM__ object generates and returns the suggestions from the ASM suggestion engine and allows for updating the subnet and range.
		This method attempts to expand the scope by expanding a range or adding a new range and, if necessary, expanding the subnet.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AsmAPICreateRequest
	*/
	Create(ctx context.Context) AsmAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateASMResponse
	CreateExecute(r AsmAPICreateRequest) (*CreateASMResponse, *http.Response, error)
	/*
			List Retrieve suggested updates for Automated Scope Management.

			Use this method to retrieve __ASM__ objects for Automated Scope Management.
		The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AsmAPIListRequest
	*/
	List(ctx context.Context) AsmAPIListRequest

	// ListExecute executes the request
	//  @return ListASMResponse
	ListExecute(r AsmAPIListRequest) (*ListASMResponse, *http.Response, error)
	/*
			Read Retrieve the suggested update for Automated Scope Management.

			Use this method to retrieve an __ASM__ object for Automated Scope Management.
		The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return AsmAPIReadRequest
	*/
	Read(ctx context.Context, id string) AsmAPIReadRequest

	// ReadExecute executes the request
	//  @return ReadASMResponse
	ReadExecute(r AsmAPIReadRequest) (*ReadASMResponse, *http.Response, error)
}

// AsmAPIService AsmAPI service
type AsmAPIService internal.Service

type AsmAPICreateRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	body       *ASM
}

func (r AsmAPICreateRequest) Body(body ASM) AsmAPICreateRequest {
	r.body = &body
	return r
}

func (r AsmAPICreateRequest) Execute() (*CreateASMResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Update subnet and ranges for Automated Scope Management.

Use this method to update the subnet and range for Automated Scope Management.
The __ASM__ object generates and returns the suggestions from the ASM suggestion engine and allows for updating the subnet and range.
This method attempts to expand the scope by expanding a range or adding a new range and, if necessary, expanding the subnet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AsmAPICreateRequest
*/
func (a *AsmAPIService) Create(ctx context.Context) AsmAPICreateRequest {
	return AsmAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateASMResponse
func (a *AsmAPIService) CreateExecute(r AsmAPICreateRequest) (*CreateASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type AsmAPIListRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	fields     *string
	subnetId   *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r AsmAPIListRequest) Fields(fields string) AsmAPIListRequest {
	r.fields = &fields
	return r
}

func (r AsmAPIListRequest) SubnetId(subnetId string) AsmAPIListRequest {
	r.subnetId = &subnetId
	return r
}

func (r AsmAPIListRequest) Execute() (*ListASMResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve suggested updates for Automated Scope Management.

Use this method to retrieve __ASM__ objects for Automated Scope Management.
The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AsmAPIListRequest
*/
func (a *AsmAPIService) List(ctx context.Context) AsmAPIListRequest {
	return AsmAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListASMResponse
func (a *AsmAPIService) ListExecute(r AsmAPIListRequest) (*ListASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.subnetId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "subnet_id", r.subnetId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type AsmAPIReadRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r AsmAPIReadRequest) Fields(fields string) AsmAPIReadRequest {
	r.fields = &fields
	return r
}

func (r AsmAPIReadRequest) Execute() (*ReadASMResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Retrieve the suggested update for Automated Scope Management.

Use this method to retrieve an __ASM__ object for Automated Scope Management.
The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return AsmAPIReadRequest
*/
func (a *AsmAPIService) Read(ctx context.Context, id string) AsmAPIReadRequest {
	return AsmAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ReadASMResponse
func (a *AsmAPIService) ReadExecute(r AsmAPIReadRequest) (*ReadASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}