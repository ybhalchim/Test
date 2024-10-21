/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the ReadOptionCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadOptionCodeResponse{}

// ReadOptionCodeResponse The response format to retrieve the __OptionCode__ object.
type ReadOptionCodeResponse struct {
	// The OptionCode object.
	Result               *OptionCode `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadOptionCodeResponse ReadOptionCodeResponse

// NewReadOptionCodeResponse instantiates a new ReadOptionCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOptionCodeResponse() *ReadOptionCodeResponse {
	this := ReadOptionCodeResponse{}
	return &this
}

// NewReadOptionCodeResponseWithDefaults instantiates a new ReadOptionCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOptionCodeResponseWithDefaults() *ReadOptionCodeResponse {
	this := ReadOptionCodeResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadOptionCodeResponse) GetResult() OptionCode {
	if o == nil || IsNil(o.Result) {
		var ret OptionCode
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOptionCodeResponse) GetResultOk() (*OptionCode, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadOptionCodeResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OptionCode and assigns it to the Result field.
func (o *ReadOptionCodeResponse) SetResult(v OptionCode) {
	o.Result = &v
}

func (o ReadOptionCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadOptionCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadOptionCodeResponse) UnmarshalJSON(data []byte) (err error) {
	varReadOptionCodeResponse := _ReadOptionCodeResponse{}

	err = json.Unmarshal(data, &varReadOptionCodeResponse)

	if err != nil {
		return err
	}

	*o = ReadOptionCodeResponse(varReadOptionCodeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadOptionCodeResponse struct {
	value *ReadOptionCodeResponse
	isSet bool
}

func (v NullableReadOptionCodeResponse) Get() *ReadOptionCodeResponse {
	return v.value
}

func (v *NullableReadOptionCodeResponse) Set(val *ReadOptionCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOptionCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOptionCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOptionCodeResponse(val *ReadOptionCodeResponse) *NullableReadOptionCodeResponse {
	return &NullableReadOptionCodeResponse{value: val, isSet: true}
}

func (v NullableReadOptionCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOptionCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}