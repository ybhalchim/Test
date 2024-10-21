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
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type InheritedDDNSBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var InheritedDDNSBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: DDNSBlockAttrTypes},
}

var InheritedDDNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes:          utils.ToComputedAttributeMap(DDNSBlockResourceSchemaAttributes),
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

func ExpandInheritedDDNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDDNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDDNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDDNSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDDNSBlock {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDDNSBlock{
		Action: m.Action.ValueStringPointer(),
	}
	return to
}

func FlattenInheritedDDNSBlock(ctx context.Context, from *ipam.InheritedDDNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDDNSBlockAttrTypes)
	}
	m := InheritedDDNSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDDNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDDNSBlockModel) Flatten(ctx context.Context, from *ipam.InheritedDDNSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDDNSBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenDDNSBlock(ctx, from.Value, diags)
}
