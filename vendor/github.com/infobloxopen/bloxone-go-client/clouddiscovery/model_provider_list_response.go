/*
Discovery Configuration API V2

The Discovery configuration service is a BloxOne Service that provides configuration for accessing and syncing the Cloud assets   Base Paths:  1. provider: **_/api/cloud_discovery/v2/_**

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clouddiscovery

import (
	"encoding/json"
)

// checks if the ProviderListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderListResponse{}

// ProviderListResponse The Provider object List response format.
type ProviderListResponse struct {
	Page                 *ApiPageInfo      `json:"page,omitempty"`
	Results              []DiscoveryConfig `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProviderListResponse ProviderListResponse

// NewProviderListResponse instantiates a new ProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderListResponse() *ProviderListResponse {
	this := ProviderListResponse{}
	return &this
}

// NewProviderListResponseWithDefaults instantiates a new ProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderListResponseWithDefaults() *ProviderListResponse {
	this := ProviderListResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ProviderListResponse) GetPage() ApiPageInfo {
	if o == nil || IsNil(o.Page) {
		var ret ApiPageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderListResponse) GetPageOk() (*ApiPageInfo, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ProviderListResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given ApiPageInfo and assigns it to the Page field.
func (o *ProviderListResponse) SetPage(v ApiPageInfo) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProviderListResponse) GetResults() []DiscoveryConfig {
	if o == nil || IsNil(o.Results) {
		var ret []DiscoveryConfig
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderListResponse) GetResultsOk() ([]DiscoveryConfig, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProviderListResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DiscoveryConfig and assigns it to the Results field.
func (o *ProviderListResponse) SetResults(v []DiscoveryConfig) {
	o.Results = v
}

func (o ProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	varProviderListResponse := _ProviderListResponse{}

	err = json.Unmarshal(data, &varProviderListResponse)

	if err != nil {
		return err
	}

	*o = ProviderListResponse(varProviderListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProviderListResponse struct {
	value *ProviderListResponse
	isSet bool
}

func (v NullableProviderListResponse) Get() *ProviderListResponse {
	return v.value
}

func (v *NullableProviderListResponse) Set(val *ProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderListResponse(val *ProviderListResponse) *NullableProviderListResponse {
	return &NullableProviderListResponse{value: val, isSet: true}
}

func (v NullableProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}