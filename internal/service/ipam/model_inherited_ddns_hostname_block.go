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

type InheritedDDNSHostnameBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var InheritedDDNSHostnameBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: DDNSHostnameBlockAttrTypes},
}

var InheritedDDNSHostnameBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          utils.ToComputedAttributeMap(DDNSHostnameBlockResourceSchemaAttributes),
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

func ExpandInheritedDDNSHostnameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDDNSHostnameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDDNSHostnameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDDNSHostnameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDDNSHostnameBlock {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDDNSHostnameBlock{
		Action: m.Action.ValueStringPointer(),
	}
	return to
}

func FlattenInheritedDDNSHostnameBlock(ctx context.Context, from *ipam.InheritedDDNSHostnameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDDNSHostnameBlockAttrTypes)
	}
	m := InheritedDDNSHostnameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDDNSHostnameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDDNSHostnameBlockModel) Flatten(ctx context.Context, from *ipam.InheritedDDNSHostnameBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDDNSHostnameBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenDDNSHostnameBlock(ctx, from.Value, diags)
}
