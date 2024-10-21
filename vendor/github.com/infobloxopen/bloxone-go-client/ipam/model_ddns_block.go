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

// checks if the DDNSBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DDNSBlock{}

// DDNSBlock The dynamic DNS configuration.
type DDNSBlock struct {
	// The Kerberos principal name. It uses the typical Kerberos notation: <SERVICE-NAME>/<server-domain-name>@<REALM>.  Defaults to empty.
	ClientPrincipal *string `json:"client_principal,omitempty"`
	// The domain name for DDNS.
	DdnsDomain *string `json:"ddns_domain,omitempty"`
	// Indicates if DDNS is enabled.
	DdnsEnabled *bool `json:"ddns_enabled,omitempty"`
	// Determines if DDNS updates are enabled at this level.
	DdnsSendUpdates *bool `json:"ddns_send_updates,omitempty"`
	// The list of DDNS zones.
	DdnsZones []DDNSZone `json:"ddns_zones,omitempty"`
	// The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.
	GssTsigFallback *bool `json:"gss_tsig_fallback,omitempty"`
	// Address of Kerberos Key Distribution Center.  Defaults to empty.
	KerberosKdc *string `json:"kerberos_kdc,omitempty"`
	// _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.
	KerberosKeys []KerberosKey `json:"kerberos_keys,omitempty"`
	// Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the rekey_interval value.  Defaults to 120 seconds.
	KerberosRekeyInterval *int64 `json:"kerberos_rekey_interval,omitempty"`
	// Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds.
	KerberosRetryInterval *int64 `json:"kerberos_retry_interval,omitempty"`
	// Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds.
	KerberosTkeyLifetime *int64 `json:"kerberos_tkey_lifetime,omitempty"`
	// Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_.
	KerberosTkeyProtocol *string `json:"kerberos_tkey_protocol,omitempty"`
	// The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty.
	ServerPrincipal      *string `json:"server_principal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DDNSBlock DDNSBlock

// NewDDNSBlock instantiates a new DDNSBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDDNSBlock() *DDNSBlock {
	this := DDNSBlock{}
	return &this
}

// NewDDNSBlockWithDefaults instantiates a new DDNSBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDDNSBlockWithDefaults() *DDNSBlock {
	this := DDNSBlock{}
	return &this
}

// GetClientPrincipal returns the ClientPrincipal field value if set, zero value otherwise.
func (o *DDNSBlock) GetClientPrincipal() string {
	if o == nil || IsNil(o.ClientPrincipal) {
		var ret string
		return ret
	}
	return *o.ClientPrincipal
}

// GetClientPrincipalOk returns a tuple with the ClientPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetClientPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.ClientPrincipal) {
		return nil, false
	}
	return o.ClientPrincipal, true
}

// HasClientPrincipal returns a boolean if a field has been set.
func (o *DDNSBlock) HasClientPrincipal() bool {
	if o != nil && !IsNil(o.ClientPrincipal) {
		return true
	}

	return false
}

// SetClientPrincipal gets a reference to the given string and assigns it to the ClientPrincipal field.
func (o *DDNSBlock) SetClientPrincipal(v string) {
	o.ClientPrincipal = &v
}

// GetDdnsDomain returns the DdnsDomain field value if set, zero value otherwise.
func (o *DDNSBlock) GetDdnsDomain() string {
	if o == nil || IsNil(o.DdnsDomain) {
		var ret string
		return ret
	}
	return *o.DdnsDomain
}

// GetDdnsDomainOk returns a tuple with the DdnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetDdnsDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DdnsDomain) {
		return nil, false
	}
	return o.DdnsDomain, true
}

// HasDdnsDomain returns a boolean if a field has been set.
func (o *DDNSBlock) HasDdnsDomain() bool {
	if o != nil && !IsNil(o.DdnsDomain) {
		return true
	}

	return false
}

// SetDdnsDomain gets a reference to the given string and assigns it to the DdnsDomain field.
func (o *DDNSBlock) SetDdnsDomain(v string) {
	o.DdnsDomain = &v
}

// GetDdnsEnabled returns the DdnsEnabled field value if set, zero value otherwise.
func (o *DDNSBlock) GetDdnsEnabled() bool {
	if o == nil || IsNil(o.DdnsEnabled) {
		var ret bool
		return ret
	}
	return *o.DdnsEnabled
}

// GetDdnsEnabledOk returns a tuple with the DdnsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetDdnsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DdnsEnabled) {
		return nil, false
	}
	return o.DdnsEnabled, true
}

// HasDdnsEnabled returns a boolean if a field has been set.
func (o *DDNSBlock) HasDdnsEnabled() bool {
	if o != nil && !IsNil(o.DdnsEnabled) {
		return true
	}

	return false
}

// SetDdnsEnabled gets a reference to the given bool and assigns it to the DdnsEnabled field.
func (o *DDNSBlock) SetDdnsEnabled(v bool) {
	o.DdnsEnabled = &v
}

// GetDdnsSendUpdates returns the DdnsSendUpdates field value if set, zero value otherwise.
func (o *DDNSBlock) GetDdnsSendUpdates() bool {
	if o == nil || IsNil(o.DdnsSendUpdates) {
		var ret bool
		return ret
	}
	return *o.DdnsSendUpdates
}

// GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetDdnsSendUpdatesOk() (*bool, bool) {
	if o == nil || IsNil(o.DdnsSendUpdates) {
		return nil, false
	}
	return o.DdnsSendUpdates, true
}

// HasDdnsSendUpdates returns a boolean if a field has been set.
func (o *DDNSBlock) HasDdnsSendUpdates() bool {
	if o != nil && !IsNil(o.DdnsSendUpdates) {
		return true
	}

	return false
}

// SetDdnsSendUpdates gets a reference to the given bool and assigns it to the DdnsSendUpdates field.
func (o *DDNSBlock) SetDdnsSendUpdates(v bool) {
	o.DdnsSendUpdates = &v
}

// GetDdnsZones returns the DdnsZones field value if set, zero value otherwise.
func (o *DDNSBlock) GetDdnsZones() []DDNSZone {
	if o == nil || IsNil(o.DdnsZones) {
		var ret []DDNSZone
		return ret
	}
	return o.DdnsZones
}

// GetDdnsZonesOk returns a tuple with the DdnsZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetDdnsZonesOk() ([]DDNSZone, bool) {
	if o == nil || IsNil(o.DdnsZones) {
		return nil, false
	}
	return o.DdnsZones, true
}

// HasDdnsZones returns a boolean if a field has been set.
func (o *DDNSBlock) HasDdnsZones() bool {
	if o != nil && !IsNil(o.DdnsZones) {
		return true
	}

	return false
}

// SetDdnsZones gets a reference to the given []DDNSZone and assigns it to the DdnsZones field.
func (o *DDNSBlock) SetDdnsZones(v []DDNSZone) {
	o.DdnsZones = v
}

// GetGssTsigFallback returns the GssTsigFallback field value if set, zero value otherwise.
func (o *DDNSBlock) GetGssTsigFallback() bool {
	if o == nil || IsNil(o.GssTsigFallback) {
		var ret bool
		return ret
	}
	return *o.GssTsigFallback
}

// GetGssTsigFallbackOk returns a tuple with the GssTsigFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetGssTsigFallbackOk() (*bool, bool) {
	if o == nil || IsNil(o.GssTsigFallback) {
		return nil, false
	}
	return o.GssTsigFallback, true
}

// HasGssTsigFallback returns a boolean if a field has been set.
func (o *DDNSBlock) HasGssTsigFallback() bool {
	if o != nil && !IsNil(o.GssTsigFallback) {
		return true
	}

	return false
}

// SetGssTsigFallback gets a reference to the given bool and assigns it to the GssTsigFallback field.
func (o *DDNSBlock) SetGssTsigFallback(v bool) {
	o.GssTsigFallback = &v
}

// GetKerberosKdc returns the KerberosKdc field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosKdc() string {
	if o == nil || IsNil(o.KerberosKdc) {
		var ret string
		return ret
	}
	return *o.KerberosKdc
}

// GetKerberosKdcOk returns a tuple with the KerberosKdc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosKdcOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosKdc) {
		return nil, false
	}
	return o.KerberosKdc, true
}

// HasKerberosKdc returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosKdc() bool {
	if o != nil && !IsNil(o.KerberosKdc) {
		return true
	}

	return false
}

// SetKerberosKdc gets a reference to the given string and assigns it to the KerberosKdc field.
func (o *DDNSBlock) SetKerberosKdc(v string) {
	o.KerberosKdc = &v
}

// GetKerberosKeys returns the KerberosKeys field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosKeys() []KerberosKey {
	if o == nil || IsNil(o.KerberosKeys) {
		var ret []KerberosKey
		return ret
	}
	return o.KerberosKeys
}

// GetKerberosKeysOk returns a tuple with the KerberosKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosKeysOk() ([]KerberosKey, bool) {
	if o == nil || IsNil(o.KerberosKeys) {
		return nil, false
	}
	return o.KerberosKeys, true
}

// HasKerberosKeys returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosKeys() bool {
	if o != nil && !IsNil(o.KerberosKeys) {
		return true
	}

	return false
}

// SetKerberosKeys gets a reference to the given []KerberosKey and assigns it to the KerberosKeys field.
func (o *DDNSBlock) SetKerberosKeys(v []KerberosKey) {
	o.KerberosKeys = v
}

// GetKerberosRekeyInterval returns the KerberosRekeyInterval field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosRekeyInterval() int64 {
	if o == nil || IsNil(o.KerberosRekeyInterval) {
		var ret int64
		return ret
	}
	return *o.KerberosRekeyInterval
}

// GetKerberosRekeyIntervalOk returns a tuple with the KerberosRekeyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosRekeyIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.KerberosRekeyInterval) {
		return nil, false
	}
	return o.KerberosRekeyInterval, true
}

// HasKerberosRekeyInterval returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosRekeyInterval() bool {
	if o != nil && !IsNil(o.KerberosRekeyInterval) {
		return true
	}

	return false
}

// SetKerberosRekeyInterval gets a reference to the given int64 and assigns it to the KerberosRekeyInterval field.
func (o *DDNSBlock) SetKerberosRekeyInterval(v int64) {
	o.KerberosRekeyInterval = &v
}

// GetKerberosRetryInterval returns the KerberosRetryInterval field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosRetryInterval() int64 {
	if o == nil || IsNil(o.KerberosRetryInterval) {
		var ret int64
		return ret
	}
	return *o.KerberosRetryInterval
}

// GetKerberosRetryIntervalOk returns a tuple with the KerberosRetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosRetryIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.KerberosRetryInterval) {
		return nil, false
	}
	return o.KerberosRetryInterval, true
}

// HasKerberosRetryInterval returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosRetryInterval() bool {
	if o != nil && !IsNil(o.KerberosRetryInterval) {
		return true
	}

	return false
}

// SetKerberosRetryInterval gets a reference to the given int64 and assigns it to the KerberosRetryInterval field.
func (o *DDNSBlock) SetKerberosRetryInterval(v int64) {
	o.KerberosRetryInterval = &v
}

// GetKerberosTkeyLifetime returns the KerberosTkeyLifetime field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosTkeyLifetime() int64 {
	if o == nil || IsNil(o.KerberosTkeyLifetime) {
		var ret int64
		return ret
	}
	return *o.KerberosTkeyLifetime
}

// GetKerberosTkeyLifetimeOk returns a tuple with the KerberosTkeyLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosTkeyLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.KerberosTkeyLifetime) {
		return nil, false
	}
	return o.KerberosTkeyLifetime, true
}

// HasKerberosTkeyLifetime returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosTkeyLifetime() bool {
	if o != nil && !IsNil(o.KerberosTkeyLifetime) {
		return true
	}

	return false
}

// SetKerberosTkeyLifetime gets a reference to the given int64 and assigns it to the KerberosTkeyLifetime field.
func (o *DDNSBlock) SetKerberosTkeyLifetime(v int64) {
	o.KerberosTkeyLifetime = &v
}

// GetKerberosTkeyProtocol returns the KerberosTkeyProtocol field value if set, zero value otherwise.
func (o *DDNSBlock) GetKerberosTkeyProtocol() string {
	if o == nil || IsNil(o.KerberosTkeyProtocol) {
		var ret string
		return ret
	}
	return *o.KerberosTkeyProtocol
}

// GetKerberosTkeyProtocolOk returns a tuple with the KerberosTkeyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetKerberosTkeyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosTkeyProtocol) {
		return nil, false
	}
	return o.KerberosTkeyProtocol, true
}

// HasKerberosTkeyProtocol returns a boolean if a field has been set.
func (o *DDNSBlock) HasKerberosTkeyProtocol() bool {
	if o != nil && !IsNil(o.KerberosTkeyProtocol) {
		return true
	}

	return false
}

// SetKerberosTkeyProtocol gets a reference to the given string and assigns it to the KerberosTkeyProtocol field.
func (o *DDNSBlock) SetKerberosTkeyProtocol(v string) {
	o.KerberosTkeyProtocol = &v
}

// GetServerPrincipal returns the ServerPrincipal field value if set, zero value otherwise.
func (o *DDNSBlock) GetServerPrincipal() string {
	if o == nil || IsNil(o.ServerPrincipal) {
		var ret string
		return ret
	}
	return *o.ServerPrincipal
}

// GetServerPrincipalOk returns a tuple with the ServerPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DDNSBlock) GetServerPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.ServerPrincipal) {
		return nil, false
	}
	return o.ServerPrincipal, true
}

// HasServerPrincipal returns a boolean if a field has been set.
func (o *DDNSBlock) HasServerPrincipal() bool {
	if o != nil && !IsNil(o.ServerPrincipal) {
		return true
	}

	return false
}

// SetServerPrincipal gets a reference to the given string and assigns it to the ServerPrincipal field.
func (o *DDNSBlock) SetServerPrincipal(v string) {
	o.ServerPrincipal = &v
}

func (o DDNSBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DDNSBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientPrincipal) {
		toSerialize["client_principal"] = o.ClientPrincipal
	}
	if !IsNil(o.DdnsDomain) {
		toSerialize["ddns_domain"] = o.DdnsDomain
	}
	if !IsNil(o.DdnsEnabled) {
		toSerialize["ddns_enabled"] = o.DdnsEnabled
	}
	if !IsNil(o.DdnsSendUpdates) {
		toSerialize["ddns_send_updates"] = o.DdnsSendUpdates
	}
	if !IsNil(o.DdnsZones) {
		toSerialize["ddns_zones"] = o.DdnsZones
	}
	if !IsNil(o.GssTsigFallback) {
		toSerialize["gss_tsig_fallback"] = o.GssTsigFallback
	}
	if !IsNil(o.KerberosKdc) {
		toSerialize["kerberos_kdc"] = o.KerberosKdc
	}
	if !IsNil(o.KerberosKeys) {
		toSerialize["kerberos_keys"] = o.KerberosKeys
	}
	if !IsNil(o.KerberosRekeyInterval) {
		toSerialize["kerberos_rekey_interval"] = o.KerberosRekeyInterval
	}
	if !IsNil(o.KerberosRetryInterval) {
		toSerialize["kerberos_retry_interval"] = o.KerberosRetryInterval
	}
	if !IsNil(o.KerberosTkeyLifetime) {
		toSerialize["kerberos_tkey_lifetime"] = o.KerberosTkeyLifetime
	}
	if !IsNil(o.KerberosTkeyProtocol) {
		toSerialize["kerberos_tkey_protocol"] = o.KerberosTkeyProtocol
	}
	if !IsNil(o.ServerPrincipal) {
		toSerialize["server_principal"] = o.ServerPrincipal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DDNSBlock) UnmarshalJSON(data []byte) (err error) {
	varDDNSBlock := _DDNSBlock{}

	err = json.Unmarshal(data, &varDDNSBlock)

	if err != nil {
		return err
	}

	*o = DDNSBlock(varDDNSBlock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_principal")
		delete(additionalProperties, "ddns_domain")
		delete(additionalProperties, "ddns_enabled")
		delete(additionalProperties, "ddns_send_updates")
		delete(additionalProperties, "ddns_zones")
		delete(additionalProperties, "gss_tsig_fallback")
		delete(additionalProperties, "kerberos_kdc")
		delete(additionalProperties, "kerberos_keys")
		delete(additionalProperties, "kerberos_rekey_interval")
		delete(additionalProperties, "kerberos_retry_interval")
		delete(additionalProperties, "kerberos_tkey_lifetime")
		delete(additionalProperties, "kerberos_tkey_protocol")
		delete(additionalProperties, "server_principal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDDNSBlock struct {
	value *DDNSBlock
	isSet bool
}

func (v NullableDDNSBlock) Get() *DDNSBlock {
	return v.value
}

func (v *NullableDDNSBlock) Set(val *DDNSBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableDDNSBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableDDNSBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDDNSBlock(val *DDNSBlock) *NullableDDNSBlock {
	return &NullableDDNSBlock{value: val, isSet: true}
}

func (v NullableDDNSBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDDNSBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
