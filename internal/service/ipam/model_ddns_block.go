package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DDNSBlockModel struct {
	ClientPrincipal       types.String `tfsdk:"client_principal"`
	DdnsDomain            types.String `tfsdk:"ddns_domain"`
	DdnsEnabled           types.Bool   `tfsdk:"ddns_enabled"`
	DdnsSendUpdates       types.Bool   `tfsdk:"ddns_send_updates"`
	DdnsZones             types.List   `tfsdk:"ddns_zones"`
	GssTsigFallback       types.Bool   `tfsdk:"gss_tsig_fallback"`
	KerberosKdc           types.String `tfsdk:"kerberos_kdc"`
	KerberosKeys          types.List   `tfsdk:"kerberos_keys"`
	KerberosRekeyInterval types.Int64  `tfsdk:"kerberos_rekey_interval"`
	KerberosRetryInterval types.Int64  `tfsdk:"kerberos_retry_interval"`
	KerberosTkeyLifetime  types.Int64  `tfsdk:"kerberos_tkey_lifetime"`
	KerberosTkeyProtocol  types.String `tfsdk:"kerberos_tkey_protocol"`
	ServerPrincipal       types.String `tfsdk:"server_principal"`
}

var DDNSBlockAttrTypes = map[string]attr.Type{
	"client_principal":        types.StringType,
	"ddns_domain":             types.StringType,
	"ddns_enabled":            types.BoolType,
	"ddns_send_updates":       types.BoolType,
	"ddns_zones":              types.ListType{ElemType: types.ObjectType{AttrTypes: DDNSZoneAttrTypes}},
	"gss_tsig_fallback":       types.BoolType,
	"kerberos_kdc":            types.StringType,
	"kerberos_keys":           types.ListType{ElemType: types.ObjectType{AttrTypes: KerberosKeyAttrTypes}},
	"kerberos_rekey_interval": types.Int64Type,
	"kerberos_retry_interval": types.Int64Type,
	"kerberos_tkey_lifetime":  types.Int64Type,
	"kerberos_tkey_protocol":  types.StringType,
	"server_principal":        types.StringType,
}

var DDNSBlockResourceSchemaAttributes = map[string]schema.Attribute{

	"client_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name. It uses the typical Kerberos notation: <SERVICE-NAME>/<server-domain-name>@<REALM>.  Defaults to empty.`,
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain name for DDNS.`,
	},
	"ddns_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if DDNS is enabled.`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Determines if DDNS updates are enabled at this level.`,
	},
	"ddns_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DDNSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DDNS zones.`,
	},
	"gss_tsig_fallback": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.`,
	},
	"kerberos_kdc": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Address of Kerberos Key Distribution Center.  Defaults to empty.`,
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.`,
	},
	"kerberos_rekey_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the rekey_interval value.  Defaults to 120 seconds.`,
	},
	"kerberos_retry_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds.`,
	},
	"kerberos_tkey_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds.`,
	},
	"kerberos_tkey_protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_.`,
	},
	"server_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty.`,
	},
}

func ExpandDDNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DDNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DDNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DDNSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DDNSBlock {
	if m == nil {
		return nil
	}
	to := &ipam.DDNSBlock{
		ClientPrincipal:       flex.ExpandStringPointer(m.ClientPrincipal),
		DdnsDomain:            flex.ExpandStringPointer(m.DdnsDomain),
		DdnsEnabled:           flex.ExpandBoolPointer(m.DdnsEnabled),
		DdnsSendUpdates:       flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsZones:             flex.ExpandFrameworkListNestedBlock(ctx, m.DdnsZones, diags, ExpandDDNSZone),
		GssTsigFallback:       flex.ExpandBoolPointer(m.GssTsigFallback),
		KerberosKdc:           flex.ExpandStringPointer(m.KerberosKdc),
		KerberosKeys:          flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandKerberosKey),
		KerberosRekeyInterval: flex.ExpandInt64Pointer(m.KerberosRekeyInterval),
		KerberosRetryInterval: flex.ExpandInt64Pointer(m.KerberosRetryInterval),
		KerberosTkeyLifetime:  flex.ExpandInt64Pointer(m.KerberosTkeyLifetime),
		KerberosTkeyProtocol:  flex.ExpandStringPointer(m.KerberosTkeyProtocol),
		ServerPrincipal:       flex.ExpandStringPointer(m.ServerPrincipal),
	}
	return to
}

func FlattenDDNSBlock(ctx context.Context, from *ipam.DDNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DDNSBlockAttrTypes)
	}
	m := DDNSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DDNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DDNSBlockModel) Flatten(ctx context.Context, from *ipam.DDNSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DDNSBlockModel{}
	}
	m.ClientPrincipal = flex.FlattenStringPointer(from.ClientPrincipal)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsEnabled = types.BoolPointerValue(from.DdnsEnabled)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.DdnsZones, DDNSZoneAttrTypes, diags, FlattenDDNSZone)
	m.GssTsigFallback = types.BoolPointerValue(from.GssTsigFallback)
	m.KerberosKdc = flex.FlattenStringPointer(from.KerberosKdc)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
	m.KerberosRekeyInterval = flex.FlattenInt64Pointer(from.KerberosRekeyInterval)
	m.KerberosRetryInterval = flex.FlattenInt64Pointer(from.KerberosRetryInterval)
	m.KerberosTkeyLifetime = flex.FlattenInt64Pointer(from.KerberosTkeyLifetime)
	m.KerberosTkeyProtocol = flex.FlattenStringPointer(from.KerberosTkeyProtocol)
	m.ServerPrincipal = flex.FlattenStringPointer(from.ServerPrincipal)
}
