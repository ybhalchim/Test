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

// checks if the CreateAuthNSGResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthNSGResponse{}

// CreateAuthNSGResponse The AuthNSG object create response format.
type CreateAuthNSGResponse struct {
	// The created AuthNSG object.
	Result               *AuthNSG `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthNSGResponse CreateAuthNSGResponse

// NewCreateAuthNSGResponse instantiates a new CreateAuthNSGResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthNSGResponse() *CreateAuthNSGResponse {
	this := CreateAuthNSGResponse{}
	return &this
}

// NewCreateAuthNSGResponseWithDefaults instantiates a new CreateAuthNSGResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthNSGResponseWithDefaults() *CreateAuthNSGResponse {
	this := CreateAuthNSGResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateAuthNSGResponse) GetResult() AuthNSG {
	if o == nil || IsNil(o.Result) {
		var ret AuthNSG
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthNSGResponse) GetResultOk() (*AuthNSG, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateAuthNSGResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AuthNSG and assigns it to the Result field.
func (o *CreateAuthNSGResponse) SetResult(v AuthNSG) {
	o.Result = &v
}

func (o CreateAuthNSGResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthNSGResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAuthNSGResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateAuthNSGResponse := _CreateAuthNSGResponse{}

	err = json.Unmarshal(data, &varCreateAuthNSGResponse)

	if err != nil {
		return err
	}

	*o = CreateAuthNSGResponse(varCreateAuthNSGResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthNSGResponse struct {
	value *CreateAuthNSGResponse
	isSet bool
}

func (v NullableCreateAuthNSGResponse) Get() *CreateAuthNSGResponse {
	return v.value
}

func (v *NullableCreateAuthNSGResponse) Set(val *CreateAuthNSGResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthNSGResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthNSGResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthNSGResponse(val *CreateAuthNSGResponse) *NullableCreateAuthNSGResponse {
	return &NullableCreateAuthNSGResponse{value: val, isSet: true}
}

func (v NullableCreateAuthNSGResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthNSGResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
