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

// checks if the UpdateACLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateACLResponse{}

// UpdateACLResponse The ACL object update response format.
type UpdateACLResponse struct {
	// The updated ACL object.
	Result               *ACL `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateACLResponse UpdateACLResponse

// NewUpdateACLResponse instantiates a new UpdateACLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateACLResponse() *UpdateACLResponse {
	this := UpdateACLResponse{}
	return &this
}

// NewUpdateACLResponseWithDefaults instantiates a new UpdateACLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateACLResponseWithDefaults() *UpdateACLResponse {
	this := UpdateACLResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateACLResponse) GetResult() ACL {
	if o == nil || IsNil(o.Result) {
		var ret ACL
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateACLResponse) GetResultOk() (*ACL, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateACLResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ACL and assigns it to the Result field.
func (o *UpdateACLResponse) SetResult(v ACL) {
	o.Result = &v
}

func (o UpdateACLResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateACLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateACLResponse) UnmarshalJSON(data []byte) (err error) {
	varUpdateACLResponse := _UpdateACLResponse{}

	err = json.Unmarshal(data, &varUpdateACLResponse)

	if err != nil {
		return err
	}

	*o = UpdateACLResponse(varUpdateACLResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateACLResponse struct {
	value *UpdateACLResponse
	isSet bool
}

func (v NullableUpdateACLResponse) Get() *UpdateACLResponse {
	return v.value
}

func (v *NullableUpdateACLResponse) Set(val *UpdateACLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateACLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateACLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateACLResponse(val *UpdateACLResponse) *NullableUpdateACLResponse {
	return &NullableUpdateACLResponse{value: val, isSet: true}
}

func (v NullableUpdateACLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateACLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}