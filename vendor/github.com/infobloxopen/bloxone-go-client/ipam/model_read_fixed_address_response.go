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

// checks if the ReadFixedAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadFixedAddressResponse{}

// ReadFixedAddressResponse The response format to retrieve the __FixedAddress__ object.
type ReadFixedAddressResponse struct {
	// The FixedAddress object.
	Result               *FixedAddress `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadFixedAddressResponse ReadFixedAddressResponse

// NewReadFixedAddressResponse instantiates a new ReadFixedAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadFixedAddressResponse() *ReadFixedAddressResponse {
	this := ReadFixedAddressResponse{}
	return &this
}

// NewReadFixedAddressResponseWithDefaults instantiates a new ReadFixedAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadFixedAddressResponseWithDefaults() *ReadFixedAddressResponse {
	this := ReadFixedAddressResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ReadFixedAddressResponse) GetResult() FixedAddress {
	if o == nil || IsNil(o.Result) {
		var ret FixedAddress
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadFixedAddressResponse) GetResultOk() (*FixedAddress, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ReadFixedAddressResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FixedAddress and assigns it to the Result field.
func (o *ReadFixedAddressResponse) SetResult(v FixedAddress) {
	o.Result = &v
}

func (o ReadFixedAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadFixedAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadFixedAddressResponse) UnmarshalJSON(data []byte) (err error) {
	varReadFixedAddressResponse := _ReadFixedAddressResponse{}

	err = json.Unmarshal(data, &varReadFixedAddressResponse)

	if err != nil {
		return err
	}

	*o = ReadFixedAddressResponse(varReadFixedAddressResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadFixedAddressResponse struct {
	value *ReadFixedAddressResponse
	isSet bool
}

func (v NullableReadFixedAddressResponse) Get() *ReadFixedAddressResponse {
	return v.value
}

func (v *NullableReadFixedAddressResponse) Set(val *ReadFixedAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadFixedAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadFixedAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadFixedAddressResponse(val *ReadFixedAddressResponse) *NullableReadFixedAddressResponse {
	return &NullableReadFixedAddressResponse{value: val, isSet: true}
}

func (v NullableReadFixedAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadFixedAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}