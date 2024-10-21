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

type AsmEnableBlockModel struct {
	Enable             types.Bool        `tfsdk:"enable"`
	EnableNotification types.Bool        `tfsdk:"enable_notification"`
	ReenableDate       timetypes.RFC3339 `tfsdk:"reenable_date"`
}

var AsmEnableBlockAttrTypes = map[string]attr.Type{
	"enable":              types.BoolType,
	"enable_notification": types.BoolType,
	"reenable_date":       timetypes.RFC3339Type{},
}

var AsmEnableBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether Automated Scope Management is enabled or not.",
	},
	"enable_notification": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether sending notifications to the users is enabled or not.",
	},
	"reenable_date": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Optional:            true,
		MarkdownDescription: "The date at which notifications will be re-enabled automatically.",
	},
}

func ExpandAsmEnableBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.AsmEnableBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AsmEnableBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AsmEnableBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.AsmEnableBlock {
	if m == nil {
		return nil
	}
	to := &ipam.AsmEnableBlock{
		Enable:             flex.ExpandBoolPointer(m.Enable),
		EnableNotification: flex.ExpandBoolPointer(m.EnableNotification),
		ReenableDate:       flex.ExpandTimePointer(ctx, m.ReenableDate, diags),
	}
	return to
}

func FlattenAsmEnableBlock(ctx context.Context, from *ipam.AsmEnableBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AsmEnableBlockAttrTypes)
	}
	m := AsmEnableBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AsmEnableBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AsmEnableBlockModel) Flatten(ctx context.Context, from *ipam.AsmEnableBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AsmEnableBlockModel{}
	}
	m.Enable = types.BoolPointerValue(from.Enable)
	m.EnableNotification = types.BoolPointerValue(from.EnableNotification)
	m.ReenableDate = timetypes.NewRFC3339TimePointerValue(from.ReenableDate)
}
