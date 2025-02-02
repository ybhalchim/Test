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

type InheritanceInheritedIdentifierModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.String `tfsdk:"value"`
}

var InheritanceInheritedIdentifierAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.StringType,
}

var InheritanceInheritedIdentifierResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandInheritanceInheritedIdentifier(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedIdentifier {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedIdentifierModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritanceInheritedIdentifierModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedIdentifier {
	if m == nil {
		return nil
	}
	to := &ipam.InheritanceInheritedIdentifier{
		Action: m.Action.ValueStringPointer(),
	}
	return to
}

func FlattenInheritanceInheritedIdentifier(ctx context.Context, from *ipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedIdentifierAttrTypes)
	}
	m := InheritanceInheritedIdentifierModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedIdentifierAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedIdentifierModel) Flatten(ctx context.Context, from *ipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedIdentifierModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenStringPointer(from.Value)
}
