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

type InheritedDHCPOptionListModel struct {
	Action types.String `tfsdk:"action"`
	Value  types.List   `tfsdk:"value"`
}

var InheritedDHCPOptionListAttrTypes = map[string]attr.Type{
	"action": types.StringType,
	"value":  types.ListType{ElemType: types.ObjectType{AttrTypes: InheritedDHCPOptionAttrTypes}},
}

var InheritedDHCPOptionListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don't use the inherited value.  Defaults to _inherit_.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritedDHCPOptionResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inherited DHCP option values.",
	},
}

func ExpandInheritedDHCPOptionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPOptionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDHCPOptionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionList {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDHCPOptionList{
		Action: m.Action.ValueStringPointer(),
		Value:  flex.ExpandFrameworkListNestedBlock(ctx, m.Value, diags, ExpandInheritedDHCPOption),
	}
	return to
}

func FlattenInheritedDHCPOptionList(ctx context.Context, from *ipam.InheritedDHCPOptionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPOptionListAttrTypes)
	}
	m := InheritedDHCPOptionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPOptionListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDHCPOptionListModel) Flatten(ctx context.Context, from *ipam.InheritedDHCPOptionList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDHCPOptionListModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, InheritedDHCPOptionAttrTypes, diags, FlattenInheritedDHCPOption)
}
