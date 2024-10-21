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

type DDNSHostnameBlockModel struct {
	DdnsGenerateName    types.Bool   `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix types.String `tfsdk:"ddns_generated_prefix"`
}

var DDNSHostnameBlockAttrTypes = map[string]attr.Type{
	"ddns_generate_name":    types.BoolType,
	"ddns_generated_prefix": types.StringType,
}

var DDNSHostnameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates if DDNS should generate a hostname when not supplied by the client.",
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The prefix used in the generation of an FQDN.",
	},
}

func ExpandDDNSHostnameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DDNSHostnameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DDNSHostnameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DDNSHostnameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DDNSHostnameBlock {
	if m == nil {
		return nil
	}
	to := &ipam.DDNSHostnameBlock{
		DdnsGenerateName:    flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix: flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
	}
	return to
}

func FlattenDDNSHostnameBlock(ctx context.Context, from *ipam.DDNSHostnameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DDNSHostnameBlockAttrTypes)
	}
	m := DDNSHostnameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DDNSHostnameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DDNSHostnameBlockModel) Flatten(ctx context.Context, from *ipam.DDNSHostnameBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DDNSHostnameBlockModel{}
	}
	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
}
