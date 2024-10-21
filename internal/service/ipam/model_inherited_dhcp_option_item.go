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

type InheritedDHCPOptionItemModel struct {
	Option          types.Object `tfsdk:"option"`
	OverridingGroup types.String `tfsdk:"overriding_group"`
}

var InheritedDHCPOptionItemAttrTypes = map[string]attr.Type{
	"option":           types.ObjectType{AttrTypes: OptionItemAttrTypes},
	"overriding_group": types.StringType,
}

var InheritedDHCPOptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"option": schema.SingleNestedAttribute{
		Attributes:          OptionItemResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Option inherited from the ancestor.",
	},
	"overriding_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandInheritedDHCPOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPOptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDHCPOptionItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionItem {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDHCPOptionItem{
		Option:          ExpandOptionItem(ctx, m.Option, diags),
		OverridingGroup: m.OverridingGroup.ValueStringPointer(),
	}
	return to
}

func FlattenInheritedDHCPOptionItem(ctx context.Context, from *ipam.InheritedDHCPOptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPOptionItemAttrTypes)
	}
	m := InheritedDHCPOptionItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPOptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDHCPOptionItemModel) Flatten(ctx context.Context, from *ipam.InheritedDHCPOptionItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDHCPOptionItemModel{}
	}
	m.Option = FlattenOptionItem(ctx, from.Option, diags)
	m.OverridingGroup = flex.FlattenStringPointer(from.OverridingGroup)
}
