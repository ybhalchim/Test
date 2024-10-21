package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type HAGroupModel struct {
	AnycastConfigId types.String      `tfsdk:"anycast_config_id"`
	Comment         types.String      `tfsdk:"comment"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Hosts           types.List        `tfsdk:"hosts"`
	Id              types.String      `tfsdk:"id"`
	IpSpace         types.String      `tfsdk:"ip_space"`
	Mode            types.String      `tfsdk:"mode"`
	Name            types.String      `tfsdk:"name"`
	Status          types.String      `tfsdk:"status"`
	StatusV6        types.String      `tfsdk:"status_v6"`
	Tags            types.Map         `tfsdk:"tags"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var HAGroupAttrTypes = map[string]attr.Type{
	"anycast_config_id": types.StringType,
	"comment":           types.StringType,
	"created_at":        timetypes.RFC3339Type{},
	"hosts":             types.ListType{ElemType: types.ObjectType{AttrTypes: HAGroupHostAttrTypes}},
	"id":                types.StringType,
	"ip_space":          types.StringType,
	"mode":              types.StringType,
	"name":              types.StringType,
	"status":            types.StringType,
	"status_v6":         types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"updated_at":        timetypes.RFC3339Type{},
}

var HAGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HAGroupHostResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "The list of hosts.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.",
	},
	"status_v6": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Status of the DHCPv6 HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the HA group.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandHAGroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.HAGroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HAGroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HAGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.HAGroup {
	if m == nil {
		return nil
	}
	to := &ipam.HAGroup{
		AnycastConfigId: flex.ExpandStringPointer(m.AnycastConfigId),
		Comment:         flex.ExpandStringPointer(m.Comment),
		Hosts:           flex.ExpandFrameworkListNestedBlock(ctx, m.Hosts, diags, ExpandHAGroupHost),
		IpSpace:         flex.ExpandStringPointer(m.IpSpace),
		Mode:            flex.ExpandStringPointer(m.Mode),
		Name:            flex.ExpandString(m.Name),
		Status:          flex.ExpandStringPointer(m.Status),
		StatusV6:        flex.ExpandStringPointer(m.StatusV6),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenHAGroup(ctx context.Context, from *ipam.HAGroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HAGroupAttrTypes)
	}
	m := HAGroupModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HAGroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HAGroupModel) Flatten(ctx context.Context, from *ipam.HAGroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HAGroupModel{}
	}
	m.AnycastConfigId = flex.FlattenStringPointer(from.AnycastConfigId)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Hosts = flex.FlattenFrameworkListNestedBlock(ctx, from.Hosts, HAGroupHostAttrTypes, diags, FlattenHAGroupHost)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Mode = flex.FlattenStringPointer(from.Mode)
	m.Name = flex.FlattenString(from.Name)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusV6 = flex.FlattenStringPointer(from.StatusV6)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
