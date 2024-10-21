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

type InheritedDDNSUpdateBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var InheritedDDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: DDNSUpdateBlockAttrTypes},
}

var InheritedDDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          utils.ToComputedAttributeMap(DDNSUpdateBlockResourceSchemaAttributes),
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

func ExpandInheritedDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDDNSUpdateBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDDNSUpdateBlock {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDDNSUpdateBlock{
		Action: m.Action.ValueStringPointer(),
	}
	return to
}

func FlattenInheritedDDNSUpdateBlock(ctx context.Context, from *ipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDDNSUpdateBlockAttrTypes)
	}
	m := InheritedDDNSUpdateBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDDNSUpdateBlockModel) Flatten(ctx context.Context, from *ipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDDNSUpdateBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenDDNSUpdateBlock(ctx, from.Value, diags)
}
