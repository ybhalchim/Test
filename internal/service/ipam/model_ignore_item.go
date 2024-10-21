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

type IgnoreItemModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

var IgnoreItemAttrTypes = map[string]attr.Type{
	"type":  types.StringType,
	"value": types.StringType,
}

var IgnoreItemResourceSchemaAttributes = map[string]schema.Attribute{
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:  * _client_hex_,  * _client_text_,  * _hardware_.",
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Value to match.",
	},
}

func ExpandIgnoreItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IgnoreItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IgnoreItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IgnoreItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IgnoreItem {
	if m == nil {
		return nil
	}
	to := &ipam.IgnoreItem{
		Type:  m.Type.ValueString(),
		Value: m.Value.ValueString(),
	}
	return to
}

func FlattenIgnoreItem(ctx context.Context, from *ipam.IgnoreItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IgnoreItemAttrTypes)
	}
	m := IgnoreItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IgnoreItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IgnoreItemModel) Flatten(ctx context.Context, from *ipam.IgnoreItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IgnoreItemModel{}
	}
	m.Type = flex.FlattenString(from.Type)
	m.Value = flex.FlattenString(from.Value)
}
