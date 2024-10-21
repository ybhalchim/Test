/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the ReadACLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadACLResponse{}

// ReadACLResponse The ACL object read response format.
type ReadACLResponse struct {
	// The ACL object.
	Result               *ACL `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadACLResponse ReadACLResponse

// NewReadACLResponse instantiates a new ReadACLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadACLResponse() *ReadACLResponse {
	this := ReadACLResponse{}
	return &this
}

// NewReadACLResponseWithDefaults instantiates a new ReadACLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadACLResponseWithDefaults() *ReadACLResponse {
	this := ReadACLResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadACLResponse) GetResult() ACL {
	if o == nil || IsNil(o.Result) {
		var ret ACL
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadACLResponse) GetResultOk() (*ACL, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadACLResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ACL and assigns it to the Result field.
func (o *ReadACLResponse) SetResult(v ACL) {
	o.Result = &v
}

func (o ReadACLResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadACLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadACLResponse) UnmarshalJSON(data []byte) (err error) {
	varReadACLResponse := _ReadACLResponse{}

	err = json.Unmarshal(data, &varReadACLResponse)

	if err != nil {
		return err
	}

	*o = ReadACLResponse(varReadACLResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadACLResponse struct {
	value *ReadACLResponse
	isSet bool
}

func (v NullableReadACLResponse) Get() *ReadACLResponse {
	return v.value
}

func (v *NullableReadACLResponse) Set(val *ReadACLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadACLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadACLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadACLResponse(val *ReadACLResponse) *NullableReadACLResponse {
	return &NullableReadACLResponse{value: val, isSet: true}
}

func (v NullableReadACLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadACLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
