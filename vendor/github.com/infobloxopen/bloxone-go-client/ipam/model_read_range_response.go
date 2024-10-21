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

// checks if the ReadRangeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadRangeResponse{}

// ReadRangeResponse The response format to retrieve the __Range__ object.
type ReadRangeResponse struct {
	// The Range object.
	Result               *Range `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadRangeResponse ReadRangeResponse

// NewReadRangeResponse instantiates a new ReadRangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRangeResponse() *ReadRangeResponse {
	this := ReadRangeResponse{}
	return &this
}

// NewReadRangeResponseWithDefaults instantiates a new ReadRangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRangeResponseWithDefaults() *ReadRangeResponse {
	this := ReadRangeResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadRangeResponse) GetResult() Range {
	if o == nil || IsNil(o.Result) {
		var ret Range
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRangeResponse) GetResultOk() (*Range, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadRangeResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Range and assigns it to the Result field.
func (o *ReadRangeResponse) SetResult(v Range) {
	o.Result = &v
}

func (o ReadRangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadRangeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadRangeResponse) UnmarshalJSON(data []byte) (err error) {
	varReadRangeResponse := _ReadRangeResponse{}

	err = json.Unmarshal(data, &varReadRangeResponse)

	if err != nil {
		return err
	}

	*o = ReadRangeResponse(varReadRangeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadRangeResponse struct {
	value *ReadRangeResponse
	isSet bool
}

func (v NullableReadRangeResponse) Get() *ReadRangeResponse {
	return v.value
}

func (v *NullableReadRangeResponse) Set(val *ReadRangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRangeResponse(val *ReadRangeResponse) *NullableReadRangeResponse {
	return &NullableReadRangeResponse{value: val, isSet: true}
}

func (v NullableReadRangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}