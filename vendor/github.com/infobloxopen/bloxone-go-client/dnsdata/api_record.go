/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsdata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type RecordAPI interface {
	/*
			Create Create the DNS resource record.

			Use this method to create a DNS __Record__ object.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return RecordAPICreateRequest
	*/
	Create(ctx context.Context) RecordAPICreateRequest

	// CreateExecute executes the request
	//  @return CreateRecordResponse
	CreateExecute(r RecordAPICreateRequest) (*CreateRecordResponse, *http.Response, error)
	/*
			Delete Move the DNS resource record to recycle bin.

			Use this method to move a DNS __Record__ object to the recycle bin.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return RecordAPIDeleteRequest
	*/
	Delete(ctx context.Context, id string) RecordAPIDeleteRequest

	// DeleteExecute executes the request
	DeleteExecute(r RecordAPIDeleteRequest) (*http.Response, error)
	/*
			List Retrieve DNS resource records.

			Use this method to retrieve DNS __Record__ objects.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return RecordAPIListRequest
	*/
	List(ctx context.Context) RecordAPIListRequest

	// ListExecute executes the request
	//  @return ListRecordResponse
	ListExecute(r RecordAPIListRequest) (*ListRecordResponse, *http.Response, error)
	/*
			Read Retrieve the DNS resource record.

			Use this method to retrieve a DNS __Record__ object.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return RecordAPIReadRequest
	*/
	Read(ctx context.Context, id string) RecordAPIReadRequest

	// ReadExecute executes the request
	//  @return ReadRecordResponse
	ReadExecute(r RecordAPIReadRequest) (*ReadRecordResponse, *http.Response, error)
	/*
			SOASerialIncrement Increment serial number for the SOA record.

			Use this method to increment the serial number for an SOA (Start of Authority) _Record_ object.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return RecordAPISOASerialIncrementRequest
	*/
	SOASerialIncrement(ctx context.Context, id string) RecordAPISOASerialIncrementRequest

	// SOASerialIncrementExecute executes the request
	//  @return SOASerialIncrementResponse
	SOASerialIncrementExecute(r RecordAPISOASerialIncrementRequest) (*SOASerialIncrementResponse, *http.Response, error)
	/*
			Update Update the DNS resource record.

			Use this method to update a DNS __Record__ object.
		A __Record__ object represents a DNS resource record in an authoritative zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return RecordAPIUpdateRequest
	*/
	Update(ctx context.Context, id string) RecordAPIUpdateRequest

	// UpdateExecute executes the request
	//  @return UpdateRecordResponse
	UpdateExecute(r RecordAPIUpdateRequest) (*UpdateRecordResponse, *http.Response, error)
}

// RecordAPIService RecordAPI service
type RecordAPIService internal.Service

type RecordAPICreateRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	body       *Record
	inherit    *string
}

func (r RecordAPICreateRequest) Body(body Record) RecordAPICreateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.
func (r RecordAPICreateRequest) Inherit(inherit string) RecordAPICreateRequest {
	r.inherit = &inherit
	return r
}

func (r RecordAPICreateRequest) Execute() (*CreateRecordResponse, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create the DNS resource record.

Use this method to create a DNS __Record__ object.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordAPICreateRequest
*/
func (a *RecordAPIService) Create(ctx context.Context) RecordAPICreateRequest {
	return RecordAPICreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateRecordResponse
func (a *RecordAPIService) CreateExecute(r RecordAPICreateRequest) (*CreateRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *CreateRecordResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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

type RecordAPIDeleteRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	id         string
}

func (r RecordAPIDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Move the DNS resource record to recycle bin.

Use this method to move a DNS __Record__ object to the recycle bin.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return RecordAPIDeleteRequest
*/
func (a *RecordAPIService) Delete(ctx context.Context, id string) RecordAPIDeleteRequest {
	return RecordAPIDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *RecordAPIService) DeleteExecute(r RecordAPIDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.Delete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record/{id}"
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

type RecordAPIListRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	tfilter    *string
	torderBy   *string
	inherit    *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r RecordAPIListRequest) Fields(fields string) RecordAPIListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r RecordAPIListRequest) Filter(filter string) RecordAPIListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r RecordAPIListRequest) Offset(offset int32) RecordAPIListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r RecordAPIListRequest) Limit(limit int32) RecordAPIListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r RecordAPIListRequest) PageToken(pageToken string) RecordAPIListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r RecordAPIListRequest) OrderBy(orderBy string) RecordAPIListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r RecordAPIListRequest) Tfilter(tfilter string) RecordAPIListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r RecordAPIListRequest) TorderBy(torderBy string) RecordAPIListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for getting inheritance_sources.
func (r RecordAPIListRequest) Inherit(inherit string) RecordAPIListRequest {
	r.inherit = &inherit
	return r
}

func (r RecordAPIListRequest) Execute() (*ListRecordResponse, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List Retrieve DNS resource records.

Use this method to retrieve DNS __Record__ objects.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RecordAPIListRequest
*/
func (a *RecordAPIService) List(ctx context.Context) RecordAPIListRequest {
	return RecordAPIListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListRecordResponse
func (a *RecordAPIService) ListExecute(r RecordAPIListRequest) (*ListRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ListRecordResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.List")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record"

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
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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

type RecordAPIReadRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	id         string
	fields     *string
	inherit    *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r RecordAPIReadRequest) Fields(fields string) RecordAPIReadRequest {
	r.fields = &fields
	return r
}

// This parameter is used for getting inheritance_sources.
func (r RecordAPIReadRequest) Inherit(inherit string) RecordAPIReadRequest {
	r.inherit = &inherit
	return r
}

func (r RecordAPIReadRequest) Execute() (*ReadRecordResponse, *http.Response, error) {
	return r.ApiService.ReadExecute(r)
}

/*
Read Retrieve the DNS resource record.

Use this method to retrieve a DNS __Record__ object.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return RecordAPIReadRequest
*/
func (a *RecordAPIService) Read(ctx context.Context, id string) RecordAPIReadRequest {
	return RecordAPIReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ReadRecordResponse
func (a *RecordAPIService) ReadExecute(r RecordAPIReadRequest) (*ReadRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ReadRecordResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.Read")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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

type RecordAPISOASerialIncrementRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	id         string
	body       *SOASerialIncrementRequest
}

func (r RecordAPISOASerialIncrementRequest) Body(body SOASerialIncrementRequest) RecordAPISOASerialIncrementRequest {
	r.body = &body
	return r
}

func (r RecordAPISOASerialIncrementRequest) Execute() (*SOASerialIncrementResponse, *http.Response, error) {
	return r.ApiService.SOASerialIncrementExecute(r)
}

/*
SOASerialIncrement Increment serial number for the SOA record.

Use this method to increment the serial number for an SOA (Start of Authority) _Record_ object.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return RecordAPISOASerialIncrementRequest
*/
func (a *RecordAPIService) SOASerialIncrement(ctx context.Context, id string) RecordAPISOASerialIncrementRequest {
	return RecordAPISOASerialIncrementRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SOASerialIncrementResponse
func (a *RecordAPIService) SOASerialIncrementExecute(r RecordAPISOASerialIncrementRequest) (*SOASerialIncrementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *SOASerialIncrementResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.SOASerialIncrement")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record/{id}/serial_increment"
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

type RecordAPIUpdateRequest struct {
	ctx        context.Context
	ApiService RecordAPI
	id         string
	body       *Record
	inherit    *string
}

func (r RecordAPIUpdateRequest) Body(body Record) RecordAPIUpdateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.
func (r RecordAPIUpdateRequest) Inherit(inherit string) RecordAPIUpdateRequest {
	r.inherit = &inherit
	return r
}

func (r RecordAPIUpdateRequest) Execute() (*UpdateRecordResponse, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update the DNS resource record.

Use this method to update a DNS __Record__ object.
A __Record__ object represents a DNS resource record in an authoritative zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return RecordAPIUpdateRequest
*/
func (a *RecordAPIService) Update(ctx context.Context, id string) RecordAPIUpdateRequest {
	return RecordAPIUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UpdateRecordResponse
func (a *RecordAPIService) UpdateExecute(r RecordAPIUpdateRequest) (*UpdateRecordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *UpdateRecordResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "RecordAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/record/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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
