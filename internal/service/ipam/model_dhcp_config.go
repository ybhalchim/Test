package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DHCPConfigModel struct {
	AbandonedReclaimTime   types.Int64 `tfsdk:"abandoned_reclaim_time"`
	AbandonedReclaimTimeV6 types.Int64 `tfsdk:"abandoned_reclaim_time_v6"`
	AllowUnknown           types.Bool  `tfsdk:"allow_unknown"`
	AllowUnknownV6         types.Bool  `tfsdk:"allow_unknown_v6"`
	EchoClientId           types.Bool  `tfsdk:"echo_client_id"`
	Filters                types.List  `tfsdk:"filters"`
	FiltersLargeSelection  types.List  `tfsdk:"filters_large_selection"`
	FiltersV6              types.List  `tfsdk:"filters_v6"`
	IgnoreClientUid        types.Bool  `tfsdk:"ignore_client_uid"`
	IgnoreList             types.List  `tfsdk:"ignore_list"`
	LeaseTime              types.Int64 `tfsdk:"lease_time"`
	LeaseTimeV6            types.Int64 `tfsdk:"lease_time_v6"`
}

var DHCPConfigAttrTypes = map[string]attr.Type{
	"abandoned_reclaim_time":    types.Int64Type,
	"abandoned_reclaim_time_v6": types.Int64Type,
	"allow_unknown":             types.BoolType,
	"allow_unknown_v6":          types.BoolType,
	"echo_client_id":            types.BoolType,
	"filters":                   types.ListType{ElemType: types.StringType},
	"filters_large_selection":   types.ListType{ElemType: types.StringType},
	"filters_v6":                types.ListType{ElemType: types.StringType},
	"ignore_client_uid":         types.BoolType,
	"ignore_list":               types.ListType{ElemType: types.ObjectType{AttrTypes: IgnoreItemAttrTypes}},
	"lease_time":                types.Int64Type,
	"lease_time_v6":             types.Int64Type,
}

var DHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned_reclaim_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The abandoned reclaim time in seconds for IPV4 clients.",
	},
	"abandoned_reclaim_time_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The abandoned reclaim time in seconds for IPV6 clients.",
	},
	"allow_unknown": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured.",
	},
	"allow_unknown_v6": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured.",
	},
	"echo_client_id": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Enable/disable to include/exclude the client id when responding to discover or request.",
	},
	"filters": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"filters_large_selection": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"filters_v6": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ignore_client_uid": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase.",
	},
	"ignore_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IgnoreItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of clients to ignore requests from.",
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease duration in seconds.",
	},
	"lease_time_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(3600),
		MarkdownDescription: "The lease duration in seconds for IPV6 clients.",
	},
}

func ExpandDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DHCPConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DHCPConfig {
	if m == nil {
		return nil
	}
	to := &ipam.DHCPConfig{
		AbandonedReclaimTime:   flex.ExpandInt64Pointer(m.AbandonedReclaimTime),
		AbandonedReclaimTimeV6: flex.ExpandInt64Pointer(m.AbandonedReclaimTimeV6),
		AllowUnknown:           flex.ExpandBoolPointer(m.AllowUnknown),
		AllowUnknownV6:         flex.ExpandBoolPointer(m.AllowUnknownV6),
		EchoClientId:           flex.ExpandBoolPointer(m.EchoClientId),
		Filters:                flex.ExpandFrameworkListString(ctx, m.Filters, diags),
		FiltersLargeSelection:  flex.ExpandFrameworkListString(ctx, m.FiltersLargeSelection, diags),
		FiltersV6:              flex.ExpandFrameworkListString(ctx, m.FiltersV6, diags),
		IgnoreClientUid:        flex.ExpandBoolPointer(m.IgnoreClientUid),
		IgnoreList:             flex.ExpandFrameworkListNestedBlock(ctx, m.IgnoreList, diags, ExpandIgnoreItem),
		LeaseTime:              flex.ExpandInt64Pointer(m.LeaseTime),
		LeaseTimeV6:            flex.ExpandInt64Pointer(m.LeaseTimeV6),
	}
	return to
}

func FlattenDHCPConfig(ctx context.Context, from *ipam.DHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DHCPConfigAttrTypes)
	}
	m := DHCPConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DHCPConfigModel) Flatten(ctx context.Context, from *ipam.DHCPConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DHCPConfigModel{}
	}
	m.AbandonedReclaimTime = flex.FlattenInt64Pointer(from.AbandonedReclaimTime)
	m.AbandonedReclaimTimeV6 = flex.FlattenInt64Pointer(from.AbandonedReclaimTimeV6)
	m.AllowUnknown = types.BoolPointerValue(from.AllowUnknown)
	m.AllowUnknownV6 = types.BoolPointerValue(from.AllowUnknownV6)
	m.EchoClientId = types.BoolPointerValue(from.EchoClientId)
	m.Filters = flex.FlattenFrameworkListString(ctx, from.Filters, diags)
	m.FiltersLargeSelection = flex.FlattenFrameworkListString(ctx, from.FiltersLargeSelection, diags)
	m.FiltersV6 = flex.FlattenFrameworkListString(ctx, from.FiltersV6, diags)
	m.IgnoreClientUid = types.BoolPointerValue(from.IgnoreClientUid)
	m.IgnoreList = flex.FlattenFrameworkListNestedBlock(ctx, from.IgnoreList, IgnoreItemAttrTypes, diags, FlattenIgnoreItem)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.LeaseTimeV6 = flex.FlattenInt64Pointer(from.LeaseTimeV6)
}
