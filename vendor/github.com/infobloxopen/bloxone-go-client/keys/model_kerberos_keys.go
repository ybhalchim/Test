/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"encoding/json"
)

// checks if the KerberosKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosKeys{}

// KerberosKeys The list of __Key__ objects.
type KerberosKeys struct {
	Items                []KerberosKey `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KerberosKeys KerberosKeys

// NewKerberosKeys instantiates a new KerberosKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosKeys() *KerberosKeys {
	this := KerberosKeys{}
	return &this
}

// NewKerberosKeysWithDefaults instantiates a new KerberosKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosKeysWithDefaults() *KerberosKeys {
	this := KerberosKeys{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *KerberosKeys) GetItems() []KerberosKey {
	if o == nil || IsNil(o.Items) {
		var ret []KerberosKey
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKeys) GetItemsOk() ([]KerberosKey, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *KerberosKeys) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KerberosKey and assigns it to the Items field.
func (o *KerberosKeys) SetItems(v []KerberosKey) {
	o.Items = v
}

func (o KerberosKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KerberosKeys) UnmarshalJSON(data []byte) (err error) {
	varKerberosKeys := _KerberosKeys{}

	err = json.Unmarshal(data, &varKerberosKeys)

	if err != nil {
		return err
	}

	*o = KerberosKeys(varKerberosKeys)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKerberosKeys struct {
	value *KerberosKeys
	isSet bool
}

func (v NullableKerberosKeys) Get() *KerberosKeys {
	return v.value
}

func (v *NullableKerberosKeys) Set(val *KerberosKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosKeys(val *KerberosKeys) *NullableKerberosKeys {
	return &NullableKerberosKeys{value: val, isSet: true}
}

func (v NullableKerberosKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}