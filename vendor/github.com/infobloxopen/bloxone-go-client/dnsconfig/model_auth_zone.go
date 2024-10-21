/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
	"time"
)

// checks if the AuthZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthZone{}

// AuthZone Authoritative zone.
type AuthZone struct {
	// Optional. Comment for zone configuration.
	Comment *string `json:"comment,omitempty"`
	// Time when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	Disabled *bool `json:"disabled,omitempty"`
	// Optional. DNS primaries external to BloxOne DDI. Order is not significant.
	ExternalPrimaries []ExternalPrimary `json:"external_primaries,omitempty"`
	// list of external providers for the auth zone.
	ExternalProviders []AuthZoneExternalProvider `json:"external_providers,omitempty"`
	// DNS secondaries external to BloxOne DDI. Order is not significant.
	ExternalSecondaries []ExternalSecondary `json:"external_secondaries,omitempty"`
	// Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
	Fqdn *string `json:"fqdn,omitempty"`
	// _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.
	GssTsigEnabled *bool `json:"gss_tsig_enabled,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The list of the inheritance assigned hosts of the object.
	InheritanceAssignedHosts []Inheritance2AssignedHost `json:"inheritance_assigned_hosts,omitempty"`
	// Optional. Inheritance configuration.
	InheritanceSources *AuthZoneInheritance `json:"inheritance_sources,omitempty"`
	// On-create-only. SOA serial is allowed to be set when the authoritative zone is created.
	InitialSoaSerial *int64 `json:"initial_soa_serial,omitempty"`
	// Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.
	InternalSecondaries []InternalSecondary `json:"internal_secondaries,omitempty"`
	// Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.
	MappedSubnet *string `json:"mapped_subnet,omitempty"`
	// Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward.
	Mapping *string `json:"mapping,omitempty"`
	// Also notify all external secondary DNS servers if enabled.  Defaults to _false_.
	Notify *bool `json:"notify,omitempty"`
	// The resource identifier.
	Nsgs []string `json:"nsgs,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host.
	PrimaryType *string `json:"primary_type,omitempty"`
	// Zone FQDN in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.
	QueryAcl []ACLItem `json:"query_acl,omitempty"`
	// Tagging specifics.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Optional. Clients must match this ACL to receive zone transfers.
	TransferAcl []ACLItem `json:"transfer_acl,omitempty"`
	// Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.
	UpdateAcl []ACLItem `json:"update_acl,omitempty"`
	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.
	UseForwardersForSubzones *bool `json:"use_forwarders_for_subzones,omitempty"`
	// The resource identifier.
	View *string `json:"view,omitempty"`
	// The list of an auth zone warnings.
	Warnings []Warning `json:"warnings,omitempty"`
	// Optional. ZoneAuthority.
	ZoneAuthority        *ZoneAuthority `json:"zone_authority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthZone AuthZone

// NewAuthZone instantiates a new AuthZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthZone() *AuthZone {
	this := AuthZone{}
	return &this
}

// NewAuthZoneWithDefaults instantiates a new AuthZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthZoneWithDefaults() *AuthZone {
	this := AuthZone{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AuthZone) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AuthZone) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AuthZone) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuthZone) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuthZone) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuthZone) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *AuthZone) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *AuthZone) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *AuthZone) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetExternalPrimaries returns the ExternalPrimaries field value if set, zero value otherwise.
func (o *AuthZone) GetExternalPrimaries() []ExternalPrimary {
	if o == nil || IsNil(o.ExternalPrimaries) {
		var ret []ExternalPrimary
		return ret
	}
	return o.ExternalPrimaries
}

// GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetExternalPrimariesOk() ([]ExternalPrimary, bool) {
	if o == nil || IsNil(o.ExternalPrimaries) {
		return nil, false
	}
	return o.ExternalPrimaries, true
}

// HasExternalPrimaries returns a boolean if a field has been set.
func (o *AuthZone) HasExternalPrimaries() bool {
	if o != nil && !IsNil(o.ExternalPrimaries) {
		return true
	}

	return false
}

// SetExternalPrimaries gets a reference to the given []ExternalPrimary and assigns it to the ExternalPrimaries field.
func (o *AuthZone) SetExternalPrimaries(v []ExternalPrimary) {
	o.ExternalPrimaries = v
}

// GetExternalProviders returns the ExternalProviders field value if set, zero value otherwise.
func (o *AuthZone) GetExternalProviders() []AuthZoneExternalProvider {
	if o == nil || IsNil(o.ExternalProviders) {
		var ret []AuthZoneExternalProvider
		return ret
	}
	return o.ExternalProviders
}

// GetExternalProvidersOk returns a tuple with the ExternalProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetExternalProvidersOk() ([]AuthZoneExternalProvider, bool) {
	if o == nil || IsNil(o.ExternalProviders) {
		return nil, false
	}
	return o.ExternalProviders, true
}

// HasExternalProviders returns a boolean if a field has been set.
func (o *AuthZone) HasExternalProviders() bool {
	if o != nil && !IsNil(o.ExternalProviders) {
		return true
	}

	return false
}

// SetExternalProviders gets a reference to the given []AuthZoneExternalProvider and assigns it to the ExternalProviders field.
func (o *AuthZone) SetExternalProviders(v []AuthZoneExternalProvider) {
	o.ExternalProviders = v
}

// GetExternalSecondaries returns the ExternalSecondaries field value if set, zero value otherwise.
func (o *AuthZone) GetExternalSecondaries() []ExternalSecondary {
	if o == nil || IsNil(o.ExternalSecondaries) {
		var ret []ExternalSecondary
		return ret
	}
	return o.ExternalSecondaries
}

// GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetExternalSecondariesOk() ([]ExternalSecondary, bool) {
	if o == nil || IsNil(o.ExternalSecondaries) {
		return nil, false
	}
	return o.ExternalSecondaries, true
}

// HasExternalSecondaries returns a boolean if a field has been set.
func (o *AuthZone) HasExternalSecondaries() bool {
	if o != nil && !IsNil(o.ExternalSecondaries) {
		return true
	}

	return false
}

// SetExternalSecondaries gets a reference to the given []ExternalSecondary and assigns it to the ExternalSecondaries field.
func (o *AuthZone) SetExternalSecondaries(v []ExternalSecondary) {
	o.ExternalSecondaries = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *AuthZone) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *AuthZone) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *AuthZone) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetGssTsigEnabled returns the GssTsigEnabled field value if set, zero value otherwise.
func (o *AuthZone) GetGssTsigEnabled() bool {
	if o == nil || IsNil(o.GssTsigEnabled) {
		var ret bool
		return ret
	}
	return *o.GssTsigEnabled
}

// GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetGssTsigEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GssTsigEnabled) {
		return nil, false
	}
	return o.GssTsigEnabled, true
}

// HasGssTsigEnabled returns a boolean if a field has been set.
func (o *AuthZone) HasGssTsigEnabled() bool {
	if o != nil && !IsNil(o.GssTsigEnabled) {
		return true
	}

	return false
}

// SetGssTsigEnabled gets a reference to the given bool and assigns it to the GssTsigEnabled field.
func (o *AuthZone) SetGssTsigEnabled(v bool) {
	o.GssTsigEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthZone) SetId(v string) {
	o.Id = &v
}

// GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field value if set, zero value otherwise.
func (o *AuthZone) GetInheritanceAssignedHosts() []Inheritance2AssignedHost {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		var ret []Inheritance2AssignedHost
		return ret
	}
	return o.InheritanceAssignedHosts
}

// GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetInheritanceAssignedHostsOk() ([]Inheritance2AssignedHost, bool) {
	if o == nil || IsNil(o.InheritanceAssignedHosts) {
		return nil, false
	}
	return o.InheritanceAssignedHosts, true
}

// HasInheritanceAssignedHosts returns a boolean if a field has been set.
func (o *AuthZone) HasInheritanceAssignedHosts() bool {
	if o != nil && !IsNil(o.InheritanceAssignedHosts) {
		return true
	}

	return false
}

// SetInheritanceAssignedHosts gets a reference to the given []Inheritance2AssignedHost and assigns it to the InheritanceAssignedHosts field.
func (o *AuthZone) SetInheritanceAssignedHosts(v []Inheritance2AssignedHost) {
	o.InheritanceAssignedHosts = v
}

// GetInheritanceSources returns the InheritanceSources field value if set, zero value otherwise.
func (o *AuthZone) GetInheritanceSources() AuthZoneInheritance {
	if o == nil || IsNil(o.InheritanceSources) {
		var ret AuthZoneInheritance
		return ret
	}
	return *o.InheritanceSources
}

// GetInheritanceSourcesOk returns a tuple with the InheritanceSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetInheritanceSourcesOk() (*AuthZoneInheritance, bool) {
	if o == nil || IsNil(o.InheritanceSources) {
		return nil, false
	}
	return o.InheritanceSources, true
}

// HasInheritanceSources returns a boolean if a field has been set.
func (o *AuthZone) HasInheritanceSources() bool {
	if o != nil && !IsNil(o.InheritanceSources) {
		return true
	}

	return false
}

// SetInheritanceSources gets a reference to the given AuthZoneInheritance and assigns it to the InheritanceSources field.
func (o *AuthZone) SetInheritanceSources(v AuthZoneInheritance) {
	o.InheritanceSources = &v
}

// GetInitialSoaSerial returns the InitialSoaSerial field value if set, zero value otherwise.
func (o *AuthZone) GetInitialSoaSerial() int64 {
	if o == nil || IsNil(o.InitialSoaSerial) {
		var ret int64
		return ret
	}
	return *o.InitialSoaSerial
}

// GetInitialSoaSerialOk returns a tuple with the InitialSoaSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetInitialSoaSerialOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialSoaSerial) {
		return nil, false
	}
	return o.InitialSoaSerial, true
}

// HasInitialSoaSerial returns a boolean if a field has been set.
func (o *AuthZone) HasInitialSoaSerial() bool {
	if o != nil && !IsNil(o.InitialSoaSerial) {
		return true
	}

	return false
}

// SetInitialSoaSerial gets a reference to the given int64 and assigns it to the InitialSoaSerial field.
func (o *AuthZone) SetInitialSoaSerial(v int64) {
	o.InitialSoaSerial = &v
}

// GetInternalSecondaries returns the InternalSecondaries field value if set, zero value otherwise.
func (o *AuthZone) GetInternalSecondaries() []InternalSecondary {
	if o == nil || IsNil(o.InternalSecondaries) {
		var ret []InternalSecondary
		return ret
	}
	return o.InternalSecondaries
}

// GetInternalSecondariesOk returns a tuple with the InternalSecondaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetInternalSecondariesOk() ([]InternalSecondary, bool) {
	if o == nil || IsNil(o.InternalSecondaries) {
		return nil, false
	}
	return o.InternalSecondaries, true
}

// HasInternalSecondaries returns a boolean if a field has been set.
func (o *AuthZone) HasInternalSecondaries() bool {
	if o != nil && !IsNil(o.InternalSecondaries) {
		return true
	}

	return false
}

// SetInternalSecondaries gets a reference to the given []InternalSecondary and assigns it to the InternalSecondaries field.
func (o *AuthZone) SetInternalSecondaries(v []InternalSecondary) {
	o.InternalSecondaries = v
}

// GetMappedSubnet returns the MappedSubnet field value if set, zero value otherwise.
func (o *AuthZone) GetMappedSubnet() string {
	if o == nil || IsNil(o.MappedSubnet) {
		var ret string
		return ret
	}
	return *o.MappedSubnet
}

// GetMappedSubnetOk returns a tuple with the MappedSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetMappedSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.MappedSubnet) {
		return nil, false
	}
	return o.MappedSubnet, true
}

// HasMappedSubnet returns a boolean if a field has been set.
func (o *AuthZone) HasMappedSubnet() bool {
	if o != nil && !IsNil(o.MappedSubnet) {
		return true
	}

	return false
}

// SetMappedSubnet gets a reference to the given string and assigns it to the MappedSubnet field.
func (o *AuthZone) SetMappedSubnet(v string) {
	o.MappedSubnet = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *AuthZone) GetMapping() string {
	if o == nil || IsNil(o.Mapping) {
		var ret string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetMappingOk() (*string, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *AuthZone) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given string and assigns it to the Mapping field.
func (o *AuthZone) SetMapping(v string) {
	o.Mapping = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *AuthZone) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *AuthZone) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *AuthZone) SetNotify(v bool) {
	o.Notify = &v
}

// GetNsgs returns the Nsgs field value if set, zero value otherwise.
func (o *AuthZone) GetNsgs() []string {
	if o == nil || IsNil(o.Nsgs) {
		var ret []string
		return ret
	}
	return o.Nsgs
}

// GetNsgsOk returns a tuple with the Nsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetNsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nsgs) {
		return nil, false
	}
	return o.Nsgs, true
}

// HasNsgs returns a boolean if a field has been set.
func (o *AuthZone) HasNsgs() bool {
	if o != nil && !IsNil(o.Nsgs) {
		return true
	}

	return false
}

// SetNsgs gets a reference to the given []string and assigns it to the Nsgs field.
func (o *AuthZone) SetNsgs(v []string) {
	o.Nsgs = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *AuthZone) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *AuthZone) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *AuthZone) SetParent(v string) {
	o.Parent = &v
}

// GetPrimaryType returns the PrimaryType field value if set, zero value otherwise.
func (o *AuthZone) GetPrimaryType() string {
	if o == nil || IsNil(o.PrimaryType) {
		var ret string
		return ret
	}
	return *o.PrimaryType
}

// GetPrimaryTypeOk returns a tuple with the PrimaryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetPrimaryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryType) {
		return nil, false
	}
	return o.PrimaryType, true
}

// HasPrimaryType returns a boolean if a field has been set.
func (o *AuthZone) HasPrimaryType() bool {
	if o != nil && !IsNil(o.PrimaryType) {
		return true
	}

	return false
}

// SetPrimaryType gets a reference to the given string and assigns it to the PrimaryType field.
func (o *AuthZone) SetPrimaryType(v string) {
	o.PrimaryType = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *AuthZone) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *AuthZone) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *AuthZone) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetQueryAcl returns the QueryAcl field value if set, zero value otherwise.
func (o *AuthZone) GetQueryAcl() []ACLItem {
	if o == nil || IsNil(o.QueryAcl) {
		var ret []ACLItem
		return ret
	}
	return o.QueryAcl
}

// GetQueryAclOk returns a tuple with the QueryAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetQueryAclOk() ([]ACLItem, bool) {
	if o == nil || IsNil(o.QueryAcl) {
		return nil, false
	}
	return o.QueryAcl, true
}

// HasQueryAcl returns a boolean if a field has been set.
func (o *AuthZone) HasQueryAcl() bool {
	if o != nil && !IsNil(o.QueryAcl) {
		return true
	}

	return false
}

// SetQueryAcl gets a reference to the given []ACLItem and assigns it to the QueryAcl field.
func (o *AuthZone) SetQueryAcl(v []ACLItem) {
	o.QueryAcl = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AuthZone) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AuthZone) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *AuthZone) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetTransferAcl returns the TransferAcl field value if set, zero value otherwise.
func (o *AuthZone) GetTransferAcl() []ACLItem {
	if o == nil || IsNil(o.TransferAcl) {
		var ret []ACLItem
		return ret
	}
	return o.TransferAcl
}

// GetTransferAclOk returns a tuple with the TransferAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetTransferAclOk() ([]ACLItem, bool) {
	if o == nil || IsNil(o.TransferAcl) {
		return nil, false
	}
	return o.TransferAcl, true
}

// HasTransferAcl returns a boolean if a field has been set.
func (o *AuthZone) HasTransferAcl() bool {
	if o != nil && !IsNil(o.TransferAcl) {
		return true
	}

	return false
}

// SetTransferAcl gets a reference to the given []ACLItem and assigns it to the TransferAcl field.
func (o *AuthZone) SetTransferAcl(v []ACLItem) {
	o.TransferAcl = v
}

// GetUpdateAcl returns the UpdateAcl field value if set, zero value otherwise.
func (o *AuthZone) GetUpdateAcl() []ACLItem {
	if o == nil || IsNil(o.UpdateAcl) {
		var ret []ACLItem
		return ret
	}
	return o.UpdateAcl
}

// GetUpdateAclOk returns a tuple with the UpdateAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetUpdateAclOk() ([]ACLItem, bool) {
	if o == nil || IsNil(o.UpdateAcl) {
		return nil, false
	}
	return o.UpdateAcl, true
}

// HasUpdateAcl returns a boolean if a field has been set.
func (o *AuthZone) HasUpdateAcl() bool {
	if o != nil && !IsNil(o.UpdateAcl) {
		return true
	}

	return false
}

// SetUpdateAcl gets a reference to the given []ACLItem and assigns it to the UpdateAcl field.
func (o *AuthZone) SetUpdateAcl(v []ACLItem) {
	o.UpdateAcl = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthZone) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthZone) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AuthZone) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUseForwardersForSubzones returns the UseForwardersForSubzones field value if set, zero value otherwise.
func (o *AuthZone) GetUseForwardersForSubzones() bool {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		var ret bool
		return ret
	}
	return *o.UseForwardersForSubzones
}

// GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetUseForwardersForSubzonesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForwardersForSubzones) {
		return nil, false
	}
	return o.UseForwardersForSubzones, true
}

// HasUseForwardersForSubzones returns a boolean if a field has been set.
func (o *AuthZone) HasUseForwardersForSubzones() bool {
	if o != nil && !IsNil(o.UseForwardersForSubzones) {
		return true
	}

	return false
}

// SetUseForwardersForSubzones gets a reference to the given bool and assigns it to the UseForwardersForSubzones field.
func (o *AuthZone) SetUseForwardersForSubzones(v bool) {
	o.UseForwardersForSubzones = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *AuthZone) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *AuthZone) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *AuthZone) SetView(v string) {
	o.View = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AuthZone) GetWarnings() []Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetWarningsOk() ([]Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AuthZone) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Warning and assigns it to the Warnings field.
func (o *AuthZone) SetWarnings(v []Warning) {
	o.Warnings = v
}

// GetZoneAuthority returns the ZoneAuthority field value if set, zero value otherwise.
func (o *AuthZone) GetZoneAuthority() ZoneAuthority {
	if o == nil || IsNil(o.ZoneAuthority) {
		var ret ZoneAuthority
		return ret
	}
	return *o.ZoneAuthority
}

// GetZoneAuthorityOk returns a tuple with the ZoneAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZone) GetZoneAuthorityOk() (*ZoneAuthority, bool) {
	if o == nil || IsNil(o.ZoneAuthority) {
		return nil, false
	}
	return o.ZoneAuthority, true
}

// HasZoneAuthority returns a boolean if a field has been set.
func (o *AuthZone) HasZoneAuthority() bool {
	if o != nil && !IsNil(o.ZoneAuthority) {
		return true
	}

	return false
}

// SetZoneAuthority gets a reference to the given ZoneAuthority and assigns it to the ZoneAuthority field.
func (o *AuthZone) SetZoneAuthority(v ZoneAuthority) {
	o.ZoneAuthority = &v
}

func (o AuthZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.ExternalPrimaries) {
		toSerialize["external_primaries"] = o.ExternalPrimaries
	}
	if !IsNil(o.ExternalProviders) {
		toSerialize["external_providers"] = o.ExternalProviders
	}
	if !IsNil(o.ExternalSecondaries) {
		toSerialize["external_secondaries"] = o.ExternalSecondaries
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.GssTsigEnabled) {
		toSerialize["gss_tsig_enabled"] = o.GssTsigEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InheritanceAssignedHosts) {
		toSerialize["inheritance_assigned_hosts"] = o.InheritanceAssignedHosts
	}
	if !IsNil(o.InheritanceSources) {
		toSerialize["inheritance_sources"] = o.InheritanceSources
	}
	if !IsNil(o.InitialSoaSerial) {
		toSerialize["initial_soa_serial"] = o.InitialSoaSerial
	}
	if !IsNil(o.InternalSecondaries) {
		toSerialize["internal_secondaries"] = o.InternalSecondaries
	}
	if !IsNil(o.MappedSubnet) {
		toSerialize["mapped_subnet"] = o.MappedSubnet
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.Nsgs) {
		toSerialize["nsgs"] = o.Nsgs
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.PrimaryType) {
		toSerialize["primary_type"] = o.PrimaryType
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.QueryAcl) {
		toSerialize["query_acl"] = o.QueryAcl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TransferAcl) {
		toSerialize["transfer_acl"] = o.TransferAcl
	}
	if !IsNil(o.UpdateAcl) {
		toSerialize["update_acl"] = o.UpdateAcl
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UseForwardersForSubzones) {
		toSerialize["use_forwarders_for_subzones"] = o.UseForwardersForSubzones
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.ZoneAuthority) {
		toSerialize["zone_authority"] = o.ZoneAuthority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthZone) UnmarshalJSON(data []byte) (err error) {
	varAuthZone := _AuthZone{}

	err = json.Unmarshal(data, &varAuthZone)

	if err != nil {
		return err
	}

	*o = AuthZone(varAuthZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "external_primaries")
		delete(additionalProperties, "external_providers")
		delete(additionalProperties, "external_secondaries")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "gss_tsig_enabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "inheritance_assigned_hosts")
		delete(additionalProperties, "inheritance_sources")
		delete(additionalProperties, "initial_soa_serial")
		delete(additionalProperties, "internal_secondaries")
		delete(additionalProperties, "mapped_subnet")
		delete(additionalProperties, "mapping")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "nsgs")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "primary_type")
		delete(additionalProperties, "protocol_fqdn")
		delete(additionalProperties, "query_acl")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "transfer_acl")
		delete(additionalProperties, "update_acl")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "use_forwarders_for_subzones")
		delete(additionalProperties, "view")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "zone_authority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthZone struct {
	value *AuthZone
	isSet bool
}

func (v NullableAuthZone) Get() *AuthZone {
	return v.value
}

func (v *NullableAuthZone) Set(val *AuthZone) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthZone) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthZone(val *AuthZone) *NullableAuthZone {
	return &NullableAuthZone{value: val, isSet: true}
}

func (v NullableAuthZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}