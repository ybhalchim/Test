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

type AsmGrowthBlockModel struct {
	GrowthFactor types.Int64  `tfsdk:"growth_factor"`
	GrowthType   types.String `tfsdk:"growth_type"`
}

var AsmGrowthBlockAttrTypes = map[string]attr.Type{
	"growth_factor": types.Int64Type,
	"growth_type":   types.StringType,
}

var AsmGrowthBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"growth_factor": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Either the number or percentage of addresses to grow by.",
	},
	"growth_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of factor to use: _percent_ or _count_.",
	},
}

func ExpandAsmGrowthBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.AsmGrowthBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AsmGrowthBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AsmGrowthBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.AsmGrowthBlock {
	if m == nil {
		return nil
	}
	to := &ipam.AsmGrowthBlock{
		GrowthFactor: flex.ExpandInt64Pointer(m.GrowthFactor),
		GrowthType:   flex.ExpandStringPointer(m.GrowthType),
	}
	return to
}

func FlattenAsmGrowthBlock(ctx context.Context, from *ipam.AsmGrowthBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AsmGrowthBlockAttrTypes)
	}
	m := AsmGrowthBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AsmGrowthBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AsmGrowthBlockModel) Flatten(ctx context.Context, from *ipam.AsmGrowthBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AsmGrowthBlockModel{}
	}
	m.GrowthFactor = flex.FlattenInt64Pointer(from.GrowthFactor)
	m.GrowthType = flex.FlattenStringPointer(from.GrowthType)
}
