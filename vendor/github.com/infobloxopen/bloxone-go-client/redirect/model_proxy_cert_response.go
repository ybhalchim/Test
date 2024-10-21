/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the ProxyCertResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyCertResponse{}

// ProxyCertResponse struct for ProxyCertResponse
type ProxyCertResponse struct {
	// Infoblox anycast dns client certificate URL.
	AnycastDnsCertificateUrl *string `json:"anycast_dns_certificate_url,omitempty"`
	// The certificate URL.
	CertificateUrl       *string `json:"certificate_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProxyCertResponse ProxyCertResponse

// NewProxyCertResponse instantiates a new ProxyCertResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyCertResponse() *ProxyCertResponse {
	this := ProxyCertResponse{}
	return &this
}

// NewProxyCertResponseWithDefaults instantiates a new ProxyCertResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyCertResponseWithDefaults() *ProxyCertResponse {
	this := ProxyCertResponse{}
	return &this
}

// GetAnycastDnsCertificateUrl returns the AnycastDnsCertificateUrl field value if set, zero value otherwise.
func (o *ProxyCertResponse) GetAnycastDnsCertificateUrl() string {
	if o == nil || IsNil(o.AnycastDnsCertificateUrl) {
		var ret string
		return ret
	}
	return *o.AnycastDnsCertificateUrl
}

// GetAnycastDnsCertificateUrlOk returns a tuple with the AnycastDnsCertificateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyCertResponse) GetAnycastDnsCertificateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AnycastDnsCertificateUrl) {
		return nil, false
	}
	return o.AnycastDnsCertificateUrl, true
}

// HasAnycastDnsCertificateUrl returns a boolean if a field has been set.
func (o *ProxyCertResponse) HasAnycastDnsCertificateUrl() bool {
	if o != nil && !IsNil(o.AnycastDnsCertificateUrl) {
		return true
	}

	return false
}

// SetAnycastDnsCertificateUrl gets a reference to the given string and assigns it to the AnycastDnsCertificateUrl field.
func (o *ProxyCertResponse) SetAnycastDnsCertificateUrl(v string) {
	o.AnycastDnsCertificateUrl = &v
}

// GetCertificateUrl returns the CertificateUrl field value if set, zero value otherwise.
func (o *ProxyCertResponse) GetCertificateUrl() string {
	if o == nil || IsNil(o.CertificateUrl) {
		var ret string
		return ret
	}
	return *o.CertificateUrl
}

// GetCertificateUrlOk returns a tuple with the CertificateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyCertResponse) GetCertificateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateUrl) {
		return nil, false
	}
	return o.CertificateUrl, true
}

// HasCertificateUrl returns a boolean if a field has been set.
func (o *ProxyCertResponse) HasCertificateUrl() bool {
	if o != nil && !IsNil(o.CertificateUrl) {
		return true
	}

	return false
}

// SetCertificateUrl gets a reference to the given string and assigns it to the CertificateUrl field.
func (o *ProxyCertResponse) SetCertificateUrl(v string) {
	o.CertificateUrl = &v
}

func (o ProxyCertResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyCertResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnycastDnsCertificateUrl) {
		toSerialize["anycast_dns_certificate_url"] = o.AnycastDnsCertificateUrl
	}
	if !IsNil(o.CertificateUrl) {
		toSerialize["certificate_url"] = o.CertificateUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProxyCertResponse) UnmarshalJSON(data []byte) (err error) {
	varProxyCertResponse := _ProxyCertResponse{}

	err = json.Unmarshal(data, &varProxyCertResponse)

	if err != nil {
		return err
	}

	*o = ProxyCertResponse(varProxyCertResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "anycast_dns_certificate_url")
		delete(additionalProperties, "certificate_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProxyCertResponse struct {
	value *ProxyCertResponse
	isSet bool
}

func (v NullableProxyCertResponse) Get() *ProxyCertResponse {
	return v.value
}

func (v *NullableProxyCertResponse) Set(val *ProxyCertResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyCertResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyCertResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyCertResponse(val *ProxyCertResponse) *NullableProxyCertResponse {
	return &NullableProxyCertResponse{value: val, isSet: true}
}

func (v NullableProxyCertResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyCertResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
