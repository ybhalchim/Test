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

type UtilizationThresholdModel struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	High    types.Int64 `tfsdk:"high"`
	Low     types.Int64 `tfsdk:"low"`
}

var UtilizationThresholdAttrTypes = map[string]attr.Type{
	"enabled": types.BoolType,
	"high":    types.Int64Type,
	"low":     types.Int64Type,
}

var UtilizationThresholdResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Indicates whether the utilization threshold for IP addresses is enabled or not.",
	},
	"high": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.",
	},
	"low": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.",
	},
}

func ExpandUtilizationThreshold(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.UtilizationThreshold {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UtilizationThresholdModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UtilizationThresholdModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.UtilizationThreshold {
	if m == nil {
		return nil
	}
	to := &ipam.UtilizationThreshold{
		Enabled: m.Enabled.ValueBool(),
		High:    int64(m.High.ValueInt64()),
		Low:     int64(m.Low.ValueInt64()),
	}
	return to
}

func FlattenUtilizationThreshold(ctx context.Context, from *ipam.UtilizationThreshold, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UtilizationThresholdAttrTypes)
	}
	m := UtilizationThresholdModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UtilizationThresholdAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UtilizationThresholdModel) Flatten(ctx context.Context, from *ipam.UtilizationThreshold, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = UtilizationThresholdModel{}
	}
	m.Enabled = types.BoolValue(from.Enabled)
	m.High = flex.FlattenInt64(int64(from.High))
	m.Low = flex.FlattenInt64(int64(from.Low))
}
