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

// checks if the ACLItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACLItem{}

// ACLItem Element in an ACL.   Error if both _acl_ and _address_ are given.
type ACLItem struct {
	// Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.
	Access string `json:"access"`
	// The resource identifier.
	Acl *string `json:"acl,omitempty"`
	// Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.
	Address *string `json:"address,omitempty"`
	// Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.
	Element string `json:"element"`
	// Optional. TSIG key.  Must be empty if _element_ is not _tsig_key_.
	TsigKey              *TSIGKey `json:"tsig_key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ACLItem ACLItem

// NewACLItem instantiates a new ACLItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLItem(access string, element string) *ACLItem {
	this := ACLItem{}
	this.Access = access
	this.Element = element
	return &this
}

// NewACLItemWithDefaults instantiates a new ACLItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLItemWithDefaults() *ACLItem {
	this := ACLItem{}
	return &this
}

// GetAccess returns the Access field value
func (o *ACLItem) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ACLItem) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ACLItem) SetAccess(v string) {
	o.Access = v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *ACLItem) GetAcl() string {
	if o == nil || IsNil(o.Acl) {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLItem) GetAclOk() (*string, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *ACLItem) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *ACLItem) SetAcl(v string) {
	o.Acl = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ACLItem) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLItem) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ACLItem) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ACLItem) SetAddress(v string) {
	o.Address = &v
}

// GetElement returns the Element field value
func (o *ACLItem) GetElement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *ACLItem) GetElementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element, true
}

// SetElement sets field value
func (o *ACLItem) SetElement(v string) {
	o.Element = v
}

// GetTsigKey returns the TsigKey field value if set, zero value otherwise.
func (o *ACLItem) GetTsigKey() TSIGKey {
	if o == nil || IsNil(o.TsigKey) {
		var ret TSIGKey
		return ret
	}
	return *o.TsigKey
}

// GetTsigKeyOk returns a tuple with the TsigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLItem) GetTsigKeyOk() (*TSIGKey, bool) {
	if o == nil || IsNil(o.TsigKey) {
		return nil, false
	}
	return o.TsigKey, true
}

// HasTsigKey returns a boolean if a field has been set.
func (o *ACLItem) HasTsigKey() bool {
	if o != nil && !IsNil(o.TsigKey) {
		return true
	}

	return false
}

// SetTsigKey gets a reference to the given TSIGKey and assigns it to the TsigKey field.
func (o *ACLItem) SetTsigKey(v TSIGKey) {
	o.TsigKey = &v
}

func (o ACLItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACLItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access"] = o.Access
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["element"] = o.Element
	if !IsNil(o.TsigKey) {
		toSerialize["tsig_key"] = o.TsigKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ACLItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access",
		"element",
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

	varACLItem := _ACLItem{}

	err = json.Unmarshal(data, &varACLItem)

	if err != nil {
		return err
	}

	*o = ACLItem(varACLItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "acl")
		delete(additionalProperties, "address")
		delete(additionalProperties, "element")
		delete(additionalProperties, "tsig_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableACLItem struct {
	value *ACLItem
	isSet bool
}

func (v NullableACLItem) Get() *ACLItem {
	return v.value
}

func (v *NullableACLItem) Set(val *ACLItem) {
	v.value = val
	v.isSet = true
}

func (v NullableACLItem) IsSet() bool {
	return v.isSet
}

func (v *NullableACLItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLItem(val *ACLItem) *NullableACLItem {
	return &NullableACLItem{value: val, isSet: true}
}

func (v NullableACLItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}