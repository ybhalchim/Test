/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type ApplicationFiltersAPI interface {
	/*
			CreateApplicationFilter Create Application Filter.

			Use this method to create a Application Filter object.

		Required:
		- name
		- criteria


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApplicationFiltersAPICreateApplicationFilterRequest
	*/
	CreateApplicationFilter(ctx context.Context) ApplicationFiltersAPICreateApplicationFilterRequest

	// CreateApplicationFilterExecute executes the request
	//  @return ApplicationFilterCreateResponse
	CreateApplicationFilterExecute(r ApplicationFiltersAPICreateApplicationFilterRequest) (*ApplicationFilterCreateResponse, *http.Response, error)
	/*
			DeleteApplicationFilters Delete Application Filters.

			Use this method to delete Application Filter objects. Deletion of multiple lists is an all-or-nothing operation (if any of the specified lists can not be deleted then none of the specified lists will be deleted).

		Required:
		- ids



			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApplicationFiltersAPIDeleteApplicationFiltersRequest
	*/
	DeleteApplicationFilters(ctx context.Context) ApplicationFiltersAPIDeleteApplicationFiltersRequest

	// DeleteApplicationFiltersExecute executes the request
	DeleteApplicationFiltersExecute(r ApplicationFiltersAPIDeleteApplicationFiltersRequest) (*http.Response, error)
	/*
		DeleteSingleApplicationFilters Delete Application Filter Object by ID.

		Use this method to delete single Application filter object by id.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id The Application Filter object identifier.
		@return ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest
	*/
	DeleteSingleApplicationFilters(ctx context.Context, id int32) ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest

	// DeleteSingleApplicationFiltersExecute executes the request
	DeleteSingleApplicationFiltersExecute(r ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest) (*http.Response, error)
	/*
		ListApplicationFilters List Application Filters.

		Use this method to retrieve information on all Application Filter objects for the account.



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApplicationFiltersAPIListApplicationFiltersRequest
	*/
	ListApplicationFilters(ctx context.Context) ApplicationFiltersAPIListApplicationFiltersRequest

	// ListApplicationFiltersExecute executes the request
	//  @return ApplicationFilterMultiResponse
	ListApplicationFiltersExecute(r ApplicationFiltersAPIListApplicationFiltersRequest) (*ApplicationFilterMultiResponse, *http.Response, error)
	/*
			ReadApplicationFilter Read Application Filter.

			Use this method to retrieve information on the specified Application Filter object.

		Required:
		- id


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id The Application Filter object identifier.
			@return ApplicationFiltersAPIReadApplicationFilterRequest
	*/
	ReadApplicationFilter(ctx context.Context, id int32) ApplicationFiltersAPIReadApplicationFilterRequest

	// ReadApplicationFilterExecute executes the request
	//  @return ApplicationFilterReadResponse
	ReadApplicationFilterExecute(r ApplicationFiltersAPIReadApplicationFilterRequest) (*ApplicationFilterReadResponse, *http.Response, error)
	/*
			UpdateApplicationFilter Update Application Filter.

			Use this method to update the specified Application Filter object.

		Category filters are content categorization rules that BloxOne Cloud uses to detect and filter specific internet content. Based on your configuration, specific actions such as Allow or Block, will be taken on the detected content. BloxOne Cloud provides the following content categories from which you can build your category filters: Drugs, Risk/Fraud/Crime, Entertainment/Culture, Purchasing, Information/Communication, Business/Services, Information Technology, Lifestyle, Society/Education/Religion, Mature/Violent, Games/Gambling, Pornography/Nudity and Uncategorized. Each of these categories contains sub-categories that further define the respective content. When you configure your category filter, you can add as many sub-categories as you need. You then add the category filter to your security policy and assign the Block action for the filter.

		Required:
		- id
		- name
		- criteria


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id The Application Filter object identifier.
			@return ApplicationFiltersAPIUpdateApplicationFilterRequest
	*/
	UpdateApplicationFilter(ctx context.Context, id int32) ApplicationFiltersAPIUpdateApplicationFilterRequest

	// UpdateApplicationFilterExecute executes the request
	//  @return ApplicationFilterUpdateResponse
	UpdateApplicationFilterExecute(r ApplicationFiltersAPIUpdateApplicationFilterRequest) (*ApplicationFilterUpdateResponse, *http.Response, error)
}

// ApplicationFiltersAPIService ApplicationFiltersAPI service
type ApplicationFiltersAPIService internal.Service

type ApplicationFiltersAPICreateApplicationFilterRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	body       *ApplicationFilter
}

// The Application Filter object.
func (r ApplicationFiltersAPICreateApplicationFilterRequest) Body(body ApplicationFilter) ApplicationFiltersAPICreateApplicationFilterRequest {
	r.body = &body
	return r
}

func (r ApplicationFiltersAPICreateApplicationFilterRequest) Execute() (*ApplicationFilterCreateResponse, *http.Response, error) {
	return r.ApiService.CreateApplicationFilterExecute(r)
}

/*
CreateApplicationFilter Create Application Filter.

Use this method to create a Application Filter object.

Required:
- name
- criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApplicationFiltersAPICreateApplicationFilterRequest
*/
func (a *ApplicationFiltersAPIService) CreateApplicationFilter(ctx context.Context) ApplicationFiltersAPICreateApplicationFilterRequest {
	return ApplicationFiltersAPICreateApplicationFilterRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApplicationFilterCreateResponse
func (a *ApplicationFiltersAPIService) CreateApplicationFilterExecute(r ApplicationFiltersAPICreateApplicationFilterRequest) (*ApplicationFilterCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ApplicationFilterCreateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.CreateApplicationFilter")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters"

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
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApplicationFiltersAPIDeleteApplicationFiltersRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	body       *ApplicationFiltersDeleteRequest
}

func (r ApplicationFiltersAPIDeleteApplicationFiltersRequest) Body(body ApplicationFiltersDeleteRequest) ApplicationFiltersAPIDeleteApplicationFiltersRequest {
	r.body = &body
	return r
}

func (r ApplicationFiltersAPIDeleteApplicationFiltersRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationFiltersExecute(r)
}

/*
DeleteApplicationFilters Delete Application Filters.

Use this method to delete Application Filter objects. Deletion of multiple lists is an all-or-nothing operation (if any of the specified lists can not be deleted then none of the specified lists will be deleted).

Required:
- ids

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApplicationFiltersAPIDeleteApplicationFiltersRequest
*/
func (a *ApplicationFiltersAPIService) DeleteApplicationFilters(ctx context.Context) ApplicationFiltersAPIDeleteApplicationFiltersRequest {
	return ApplicationFiltersAPIDeleteApplicationFiltersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ApplicationFiltersAPIService) DeleteApplicationFiltersExecute(r ApplicationFiltersAPIDeleteApplicationFiltersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.DeleteApplicationFilters")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, internal.ReportError("body is required and must be specified")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApplicationFiltersDeleteApplicationFilters400Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	id         int32
}

func (r ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSingleApplicationFiltersExecute(r)
}

/*
DeleteSingleApplicationFilters Delete Application Filter Object by ID.

Use this method to delete single Application filter object by id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Application Filter object identifier.
	@return ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest
*/
func (a *ApplicationFiltersAPIService) DeleteSingleApplicationFilters(ctx context.Context, id int32) ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest {
	return ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ApplicationFiltersAPIService) DeleteSingleApplicationFiltersExecute(r ApplicationFiltersAPIDeleteSingleApplicationFiltersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.DeleteSingleApplicationFilters")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters/{id}"
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApplicationFiltersDeleteSingleApplicationFilters400Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApplicationFiltersAPIListApplicationFiltersRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	filter     *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	tfilter    *string
	torderBy   *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; |  In addition grouping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Allowed sets of parameters that can be groupped in one query:  - name  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;app_list1&#39;)or(name~&#39;app_list2&#39;))\&quot; &#x60;&#x60;&#x60;
func (r ApplicationFiltersAPIListApplicationFiltersRequest) Filter(filter string) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) Fields(fields string) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) Offset(offset int32) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) Limit(limit int32) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) PageToken(pageToken string) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.pageToken = &pageToken
	return r
}

// Filtering by tags.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) Tfilter(tfilter string) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.tfilter = &tfilter
	return r
}

// Sorting by tags.
func (r ApplicationFiltersAPIListApplicationFiltersRequest) TorderBy(torderBy string) ApplicationFiltersAPIListApplicationFiltersRequest {
	r.torderBy = &torderBy
	return r
}

func (r ApplicationFiltersAPIListApplicationFiltersRequest) Execute() (*ApplicationFilterMultiResponse, *http.Response, error) {
	return r.ApiService.ListApplicationFiltersExecute(r)
}

/*
ListApplicationFilters List Application Filters.

Use this method to retrieve information on all Application Filter objects for the account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApplicationFiltersAPIListApplicationFiltersRequest
*/
func (a *ApplicationFiltersAPIService) ListApplicationFilters(ctx context.Context) ApplicationFiltersAPIListApplicationFiltersRequest {
	return ApplicationFiltersAPIListApplicationFiltersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApplicationFilterMultiResponse
func (a *ApplicationFiltersAPIService) ListApplicationFiltersExecute(r ApplicationFiltersAPIListApplicationFiltersRequest) (*ApplicationFilterMultiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ApplicationFilterMultiResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.ListApplicationFilters")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApplicationFiltersAPIReadApplicationFilterRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	id         int32
	fields     *string
	name       *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApplicationFiltersAPIReadApplicationFilterRequest) Fields(fields string) ApplicationFiltersAPIReadApplicationFilterRequest {
	r.fields = &fields
	return r
}

// The name of the application filter.
func (r ApplicationFiltersAPIReadApplicationFilterRequest) Name(name string) ApplicationFiltersAPIReadApplicationFilterRequest {
	r.name = &name
	return r
}

func (r ApplicationFiltersAPIReadApplicationFilterRequest) Execute() (*ApplicationFilterReadResponse, *http.Response, error) {
	return r.ApiService.ReadApplicationFilterExecute(r)
}

/*
ReadApplicationFilter Read Application Filter.

Use this method to retrieve information on the specified Application Filter object.

Required:
- id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Application Filter object identifier.
	@return ApplicationFiltersAPIReadApplicationFilterRequest
*/
func (a *ApplicationFiltersAPIService) ReadApplicationFilter(ctx context.Context, id int32) ApplicationFiltersAPIReadApplicationFilterRequest {
	return ApplicationFiltersAPIReadApplicationFilterRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ApplicationFilterReadResponse
func (a *ApplicationFiltersAPIService) ReadApplicationFilterExecute(r ApplicationFiltersAPIReadApplicationFilterRequest) (*ApplicationFilterReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ApplicationFilterReadResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.ReadApplicationFilter")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.name != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApplicationFiltersAPIUpdateApplicationFilterRequest struct {
	ctx        context.Context
	ApiService ApplicationFiltersAPI
	id         int32
	body       *ApplicationFilter
}

// The Application Filter object.
func (r ApplicationFiltersAPIUpdateApplicationFilterRequest) Body(body ApplicationFilter) ApplicationFiltersAPIUpdateApplicationFilterRequest {
	r.body = &body
	return r
}

func (r ApplicationFiltersAPIUpdateApplicationFilterRequest) Execute() (*ApplicationFilterUpdateResponse, *http.Response, error) {
	return r.ApiService.UpdateApplicationFilterExecute(r)
}

/*
UpdateApplicationFilter Update Application Filter.

Use this method to update the specified Application Filter object.

Category filters are content categorization rules that BloxOne Cloud uses to detect and filter specific internet content. Based on your configuration, specific actions such as Allow or Block, will be taken on the detected content. BloxOne Cloud provides the following content categories from which you can build your category filters: Drugs, Risk/Fraud/Crime, Entertainment/Culture, Purchasing, Information/Communication, Business/Services, Information Technology, Lifestyle, Society/Education/Religion, Mature/Violent, Games/Gambling, Pornography/Nudity and Uncategorized. Each of these categories contains sub-categories that further define the respective content. When you configure your category filter, you can add as many sub-categories as you need. You then add the category filter to your security policy and assign the Block action for the filter.

Required:
- id
- name
- criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Application Filter object identifier.
	@return ApplicationFiltersAPIUpdateApplicationFilterRequest
*/
func (a *ApplicationFiltersAPIService) UpdateApplicationFilter(ctx context.Context, id int32) ApplicationFiltersAPIUpdateApplicationFilterRequest {
	return ApplicationFiltersAPIUpdateApplicationFilterRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ApplicationFilterUpdateResponse
func (a *ApplicationFiltersAPIService) UpdateApplicationFilterExecute(r ApplicationFiltersAPIUpdateApplicationFilterRequest) (*ApplicationFilterUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ApplicationFilterUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ApplicationFiltersAPIService.UpdateApplicationFilter")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/application_filters/{id}"
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v AccessCodesListAccessCodes500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
