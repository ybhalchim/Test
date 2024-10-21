/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type AuthNsgAPI interface {
	/*
			Create Create the AuthNSG object.

			Use this method to create an AuthNSG object.
		The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AuthNsgAPICreateRequest
	*/
	Create(ctx context.Context) AuthNsgAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateAuthNSGResponse
	CreateExecute(r AuthNsgAPICreateRequest) (*CreateAuthNSGResponse, *http.Response, error)
	/*
			Delete Move the AuthNSG object to Recyclebin.

			Use this method to move an AuthNSG object to Recyclebin.
		The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return AuthNsgAPIDeleteRequest
	*/
	Delete(ctx context.Context, id string) AuthNsgAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r AuthNsgAPIDeleteRequest) (*http.Response, error)
	/*
			List List AuthNSG objects.

			Use this method to list AuthNSG objects.
		The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return AuthNsgAPIListRequest
	*/
	List(ctx context.Context) AuthNsgAPIListRequest

	// ListExecute executes the request
	//  @return ListAuthNSGResponse
	ListExecute(r AuthNsgAPIListRequest) (*ListAuthNSGResponse, *http.Response, error)
	/*
			Read Read the AuthNSG object.

			Use this method to read an AuthNSG object.
		The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return AuthNsgAPIReadRequest
	*/
	Read(ctx context.Context, id string) AuthNsgAPIReadRequest

	// ReadExecute executes the request
	//  @return ReadAuthNSGResponse
	ReadExecute(r AuthNsgAPIReadRequest) (*ReadAuthNSGResponse, *http.Response, error)
	/*
			Update Update the AuthNSG object.

			Use this method to update an AuthNSG object.
		The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return AuthNsgAPIUpdateRequest
	*/
	Update(ctx context.Context, id string) AuthNsgAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateAuthNSGResponse
	UpdateExecute(r AuthNsgAPIUpdateRequest) (*UpdateAuthNSGResponse, *http.Response, error)
}

// AuthNsgAPIService AuthNsgAPI service
type AuthNsgAPIService internal.Service

type AuthNsgAPICreateRequest struct {
	ctx        context.Context
	ApiService AuthNsgAPI
	body       *AuthNSG
}

func (r AuthNsgAPICreateRequest) Body(body AuthNSG) AuthNsgAPICreateRequest {
	r.body = &body
	return r
}

func (r AuthNsgAPICreateRequest) Execute() (*CreateAuthNSGResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create the AuthNSG object.

Use this method to create an AuthNSG object.
The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthNsgAPICreateRequest
*/
func (a *AuthNsgAPIService) Create(ctx context.Context) AuthNsgAPICreateRequest {
	return AuthNsgAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateAuthNSGResponse
func (a *AuthNsgAPIService) CreateExecute(r AuthNsgAPICreateRequest) (*CreateAuthNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateAuthNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AuthNsgAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/auth_nsg"

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
	if len(a.Client.Cfg.DefaultTags) > 0 && r.body != nil {
		if r.body.Tags == nil {
			r.body.Tags = make(map[string]interface{})
		}
		for k, v := range a.Client.Cfg.DefaultTags {
			if _, ok := r.body.Tags[k]; !ok {
				r.body.Tags[k] = v
			}
		}
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

type AuthNsgAPIDeleteRequest struct {
	ctx        context.Context
	ApiService AuthNsgAPI
	id         string
}

func (r AuthNsgAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Move the AuthNSG object to Recyclebin.

Use this method to move an AuthNSG object to Recyclebin.
The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return AuthNsgAPIDeleteRequest
*/
func (a *AuthNsgAPIService) Delete(ctx context.Context, id string) AuthNsgAPIDeleteRequest {
	return AuthNsgAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *AuthNsgAPIService) DeleteExecute(r AuthNsgAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AuthNsgAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/auth_nsg/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AuthNsgAPIListRequest struct {
	ctx        context.Context
	ApiService AuthNsgAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	tfilter    *string
	torderBy   *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r AuthNsgAPIListRequest) Fields(fields string) AuthNsgAPIListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r AuthNsgAPIListRequest) Filter(filter string) AuthNsgAPIListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r AuthNsgAPIListRequest) Offset(offset int32) AuthNsgAPIListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r AuthNsgAPIListRequest) Limit(limit int32) AuthNsgAPIListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r AuthNsgAPIListRequest) PageToken(pageToken string) AuthNsgAPIListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r AuthNsgAPIListRequest) OrderBy(orderBy string) AuthNsgAPIListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r AuthNsgAPIListRequest) Tfilter(tfilter string) AuthNsgAPIListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r AuthNsgAPIListRequest) TorderBy(torderBy string) AuthNsgAPIListRequest {
	r.torderBy = &torderBy
	return r
}

func (r AuthNsgAPIListRequest) Execute() (*ListAuthNSGResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List AuthNSG objects.

Use this method to list AuthNSG objects.
The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthNsgAPIListRequest
*/
func (a *AuthNsgAPIService) List(ctx context.Context) AuthNsgAPIListRequest {
	return AuthNsgAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListAuthNSGResponse
func (a *AuthNsgAPIService) ListExecute(r AuthNsgAPIListRequest) (*ListAuthNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListAuthNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AuthNsgAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/auth_nsg"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
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

type AuthNsgAPIReadRequest struct {
	ctx        context.Context
	ApiService AuthNsgAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r AuthNsgAPIReadRequest) Fields(fields string) AuthNsgAPIReadRequest {
	r.fields = &fields
	return r
}

func (r AuthNsgAPIReadRequest) Execute() (*ReadAuthNSGResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Read the AuthNSG object.

Use this method to read an AuthNSG object.
The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return AuthNsgAPIReadRequest
*/
func (a *AuthNsgAPIService) Read(ctx context.Context, id string) AuthNsgAPIReadRequest {
	return AuthNsgAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ReadAuthNSGResponse
func (a *AuthNsgAPIService) ReadExecute(r AuthNsgAPIReadRequest) (*ReadAuthNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadAuthNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AuthNsgAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/auth_nsg/{id}"
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

type AuthNsgAPIUpdateRequest struct {
	ctx        context.Context
	ApiService AuthNsgAPI
	id         string
	body       *AuthNSG
}

func (r AuthNsgAPIUpdateRequest) Body(body AuthNSG) AuthNsgAPIUpdateRequest {
	r.body = &body
	return r
}

func (r AuthNsgAPIUpdateRequest) Execute() (*UpdateAuthNSGResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update the AuthNSG object.

Use this method to update an AuthNSG object.
The _dns/auth_nsg_ object represents an Authoritative DNS Server Group for authoritative zones.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return AuthNsgAPIUpdateRequest
*/
func (a *AuthNsgAPIService) Update(ctx context.Context, id string) AuthNsgAPIUpdateRequest {
	return AuthNsgAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateAuthNSGResponse
func (a *AuthNsgAPIService) UpdateExecute(r AuthNsgAPIUpdateRequest) (*UpdateAuthNSGResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateAuthNSGResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AuthNsgAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/auth_nsg/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

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
	if len(a.Client.Cfg.DefaultTags) > 0 && r.body != nil {
		if r.body.Tags == nil {
			r.body.Tags = make(map[string]interface{})
		}
		for k, v := range a.Client.Cfg.DefaultTags {
			if _, ok := r.body.Tags[k]; !ok {
				r.body.Tags[k] = v
			}
		}
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
