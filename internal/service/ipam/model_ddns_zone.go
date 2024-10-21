package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DDNSZoneModel struct {
	Fqdn           types.String `tfsdk:"fqdn"`
	GssTsigEnabled types.Bool   `tfsdk:"gss_tsig_enabled"`
	Nameservers    types.List   `tfsdk:"nameservers"`
	TsigEnabled    types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey        types.Object `tfsdk:"tsig_key"`
	View           types.String `tfsdk:"view"`
	ViewName       types.String `tfsdk:"view_name"`
	Zone           types.String `tfsdk:"zone"`
}

var DDNSZoneAttrTypes = map[string]attr.Type{
	"fqdn":             types.StringType,
	"gss_tsig_enabled": types.BoolType,
	"nameservers":      types.ListType{ElemType: types.ObjectType{AttrTypes: NameserverAttrTypes}},
	"tsig_enabled":     types.BoolType,
	"tsig_key":         types.ObjectType{AttrTypes: TSIGKeyAttrTypes},
	"view":             types.StringType,
	"view_name":        types.StringType,
	"zone":             types.StringType,
}

var DDNSZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"fqdn": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Zone FQDN.  If _zone_ is defined, the _fqdn_ field must be empty.",
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.",
	},
	"nameservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NameserverResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The Nameservers in the zone.  Each nameserver IP should be unique across the list of nameservers.",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if TSIG key should be used for the update.  Defaults to _false_.",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes:          TSIGKeyResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The TSIG key. Required if _tsig_enabled_ is _true_.  Defaults to empty.",
	},
	"view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"view_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the view.",
	},
	"zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandDDNSZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DDNSZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DDNSZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DDNSZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DDNSZone {
	if m == nil {
		return nil
	}
	to := &ipam.DDNSZone{
		Fqdn:           flex.ExpandStringPointer(m.Fqdn),
		GssTsigEnabled: flex.ExpandBoolPointer(m.GssTsigEnabled),
		Nameservers:    flex.ExpandFrameworkListNestedBlock(ctx, m.Nameservers, diags, ExpandNameserver),
		TsigEnabled:    flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:        ExpandTSIGKey(ctx, m.TsigKey, diags),
		View:           flex.ExpandStringPointer(m.View),
		Zone:           flex.ExpandString(m.Zone),
	}
	return to
}

func FlattenDDNSZone(ctx context.Context, from *ipam.DDNSZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DDNSZoneAttrTypes)
	}
	m := DDNSZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DDNSZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DDNSZoneModel) Flatten(ctx context.Context, from *ipam.DDNSZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DDNSZoneModel{}
	}
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Nameservers = flex.FlattenFrameworkListNestedBlock(ctx, from.Nameservers, NameserverAttrTypes, diags, FlattenNameserver)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = FlattenTSIGKey(ctx, from.TsigKey, diags)
	m.View = flex.FlattenStringPointer(from.View)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.Zone = flex.FlattenString(from.Zone)
}
