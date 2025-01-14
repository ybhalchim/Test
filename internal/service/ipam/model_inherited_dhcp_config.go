package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type InheritedDHCPConfigModel struct {
	AbandonedReclaimTime   types.Object `tfsdk:"abandoned_reclaim_time"`
	AbandonedReclaimTimeV6 types.Object `tfsdk:"abandoned_reclaim_time_v6"`
	AllowUnknown           types.Object `tfsdk:"allow_unknown"`
	AllowUnknownV6         types.Object `tfsdk:"allow_unknown_v6"`
	EchoClientId           types.Object `tfsdk:"echo_client_id"`
	Filters                types.Object `tfsdk:"filters"`
	FiltersV6              types.Object `tfsdk:"filters_v6"`
	IgnoreClientUid        types.Object `tfsdk:"ignore_client_uid"`
	IgnoreList             types.Object `tfsdk:"ignore_list"`
	LeaseTime              types.Object `tfsdk:"lease_time"`
	LeaseTimeV6            types.Object `tfsdk:"lease_time_v6"`
}

var InheritedDHCPConfigAttrTypes = map[string]attr.Type{
	"abandoned_reclaim_time":    types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"abandoned_reclaim_time_v6": types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"allow_unknown":             types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"allow_unknown_v6":          types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"echo_client_id":            types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"filters":                   types.ObjectType{AttrTypes: InheritedDHCPConfigFilterListAttrTypes},
	"filters_v6":                types.ObjectType{AttrTypes: InheritedDHCPConfigFilterListAttrTypes},
	"ignore_client_uid":         types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ignore_list":               types.ObjectType{AttrTypes: InheritedDHCPConfigIgnoreItemListAttrTypes},
	"lease_time":                types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"lease_time_v6":             types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
}

var InheritedDHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned_reclaim_time": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _abandoned_reclaim_time_ field from _DHCPConfig_ object.",
	},
	"abandoned_reclaim_time_v6": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _abandoned_reclaim_time_v6_ field from _DHCPConfig_ object.",
	},
	"allow_unknown": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _allow_unknown_ field from _DHCPConfig_ object.",
	},
	"allow_unknown_v6": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _allow_unknown_v6_ field from _DHCPConfig_ object.",
	},
	"echo_client_id": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _echo_client_id_ field from _DHCPConfig_ object.",
	},
	"filters": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPConfigFilterListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for filters field from _DHCPConfig_ object.",
	},
	"filters_v6": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPConfigFilterListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _filters_v6_ field from _DHCPConfig_ object.",
	},
	"ignore_client_uid": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ignore_client_uid_ field from _DHCPConfig_ object.",
	},
	"ignore_list": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPConfigIgnoreItemListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ignore_list_ field from _DHCPConfig_ object.",
	},
	"lease_time": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _lease_time_ field from _DHCPConfig_ object.",
	},
	"lease_time_v6": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _lease_time_v6_ field from _DHCPConfig_ object.",
	},
}

func ExpandInheritedDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDHCPConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPConfig {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDHCPConfig{
		AbandonedReclaimTime:   ExpandInheritanceInheritedUInt32(ctx, m.AbandonedReclaimTime, diags),
		AbandonedReclaimTimeV6: ExpandInheritanceInheritedUInt32(ctx, m.AbandonedReclaimTimeV6, diags),
		AllowUnknown:           ExpandInheritanceInheritedBool(ctx, m.AllowUnknown, diags),
		AllowUnknownV6:         ExpandInheritanceInheritedBool(ctx, m.AllowUnknownV6, diags),
		EchoClientId:           ExpandInheritanceInheritedBool(ctx, m.EchoClientId, diags),
		Filters:                ExpandInheritedDHCPConfigFilterList(ctx, m.Filters, diags),
		FiltersV6:              ExpandInheritedDHCPConfigFilterList(ctx, m.FiltersV6, diags),
		IgnoreClientUid:        ExpandInheritanceInheritedBool(ctx, m.IgnoreClientUid, diags),
		IgnoreList:             ExpandInheritedDHCPConfigIgnoreItemList(ctx, m.IgnoreList, diags),
		LeaseTime:              ExpandInheritanceInheritedUInt32(ctx, m.LeaseTime, diags),
		LeaseTimeV6:            ExpandInheritanceInheritedUInt32(ctx, m.LeaseTimeV6, diags),
	}
	return to
}

func FlattenInheritedDHCPConfig(ctx context.Context, from *ipam.InheritedDHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPConfigAttrTypes)
	}
	m := InheritedDHCPConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDHCPConfigModel) Flatten(ctx context.Context, from *ipam.InheritedDHCPConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDHCPConfigModel{}
	}
	m.AbandonedReclaimTime = FlattenInheritanceInheritedUInt32(ctx, from.AbandonedReclaimTime, diags)
	m.AbandonedReclaimTimeV6 = FlattenInheritanceInheritedUInt32(ctx, from.AbandonedReclaimTimeV6, diags)
	m.AllowUnknown = FlattenInheritanceInheritedBool(ctx, from.AllowUnknown, diags)
	m.AllowUnknownV6 = FlattenInheritanceInheritedBool(ctx, from.AllowUnknownV6, diags)
	m.EchoClientId = FlattenInheritanceInheritedBool(ctx, from.EchoClientId, diags)
	m.Filters = FlattenInheritedDHCPConfigFilterList(ctx, from.Filters, diags)
	m.FiltersV6 = FlattenInheritedDHCPConfigFilterList(ctx, from.FiltersV6, diags)
	m.IgnoreClientUid = FlattenInheritanceInheritedBool(ctx, from.IgnoreClientUid, diags)
	m.IgnoreList = FlattenInheritedDHCPConfigIgnoreItemList(ctx, from.IgnoreList, diags)
	m.LeaseTime = FlattenInheritanceInheritedUInt32(ctx, from.LeaseTime, diags)
	m.LeaseTimeV6 = FlattenInheritanceInheritedUInt32(ctx, from.LeaseTimeV6, diags)
}
