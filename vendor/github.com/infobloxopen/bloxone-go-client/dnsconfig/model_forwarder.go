/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
	"fmt"
)

// checks if the Forwarder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Forwarder{}

// Forwarder External DNS server to forward to.
type Forwarder struct {
	// Server IP address.
	Address string `json:"address"`
	// Server FQDN.
	Fqdn *string `json:"fqdn,omitempty"`
	// Server FQDN in punycode.
	ProtocolFqdn         *string `json:"protocol_fqdn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Forwarder Forwarder

// NewForwarder instantiates a new Forwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwarder(address string) *Forwarder {
	this := Forwarder{}
	this.Address = address
	return &this
}

// NewForwarderWithDefaults instantiates a new Forwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwarderWithDefaults() *Forwarder {
	this := Forwarder{}
	return &this
}

// GetAddress returns the Address field value
func (o *Forwarder) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Forwarder) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Forwarder) SetAddress(v string) {
	o.Address = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *Forwarder) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forwarder) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *Forwarder) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *Forwarder) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *Forwarder) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forwarder) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *Forwarder) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *Forwarder) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

func (o Forwarder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Forwarder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Forwarder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varForwarder := _Forwarder{}

	err = json.Unmarshal(data, &varForwarder)

	if err != nil {
		return err
	}

	*o = Forwarder(varForwarder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "protocol_fqdn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForwarder struct {
	value *Forwarder
	isSet bool
}

func (v NullableForwarder) Get() *Forwarder {
	return v.value
}

func (v *NullableForwarder) Set(val *Forwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwarder(val *Forwarder) *NullableForwarder {
	return &NullableForwarder{value: val, isSet: true}
}

func (v NullableForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}