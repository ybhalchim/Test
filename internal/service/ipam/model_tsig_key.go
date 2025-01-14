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

type TSIGKeyModel struct {
	Algorithm    types.String `tfsdk:"algorithm"`
	Comment      types.String `tfsdk:"comment"`
	Key          types.String `tfsdk:"key"`
	Name         types.String `tfsdk:"name"`
	ProtocolName types.String `tfsdk:"protocol_name"`
	Secret       types.String `tfsdk:"secret"`
}

var TSIGKeyAttrTypes = map[string]attr.Type{
	"algorithm":     types.StringType,
	"comment":       types.StringType,
	"key":           types.StringType,
	"name":          types.StringType,
	"protocol_name": types.StringType,
	"secret":        types.StringType,
}

var TSIGKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "TSIG key algorithm.  Valid values are:  * _hmac_sha256_  * _hmac_sha1_  * _hmac_sha224_  * _hmac_sha384_  * _hmac_sha512_",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TSIG key name, FQDN.",
	},
	"protocol_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key name in punycode.",
	},
	"secret": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The TSIG key secret, base64 string.",
	},
}

func ExpandTSIGKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.TSIGKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TSIGKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TSIGKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.TSIGKey {
	if m == nil {
		return nil
	}
	to := &ipam.TSIGKey{
		Algorithm: m.Algorithm.ValueStringPointer(),
		Comment:   m.Comment.ValueStringPointer(),
		Key:       m.Key.ValueString(),
		Name:      m.Name.ValueStringPointer(),
		Secret:    m.Secret.ValueStringPointer(),
	}
	return to
}

func FlattenTSIGKey(ctx context.Context, from *ipam.TSIGKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TSIGKeyAttrTypes)
	}
	m := TSIGKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TSIGKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TSIGKeyModel) Flatten(ctx context.Context, from *ipam.TSIGKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TSIGKeyModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Key = flex.FlattenString(from.Key)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ProtocolName = flex.FlattenStringPointer(from.ProtocolName)
	m.Secret = flex.FlattenStringPointer(from.Secret)
}
