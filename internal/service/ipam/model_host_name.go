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

type HostNameModel struct {
	Alias       types.Bool   `tfsdk:"alias"`
	Name        types.String `tfsdk:"name"`
	PrimaryName types.Bool   `tfsdk:"primary_name"`
	Zone        types.String `tfsdk:"zone"`
}

var HostNameAttrTypes = map[string]attr.Type{
	"alias":        types.BoolType,
	"name":         types.StringType,
	"primary_name": types.BoolType,
	"zone":         types.StringType,
}

var HostNameResourceSchemaAttributes = map[string]schema.Attribute{
	"alias": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "When _true_, the name is treated as an alias.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "A name for the host.",
	},
	"primary_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "When _true_, the name field is treated as primary name. There must be one and only one primary name in the list of host names. The primary name will be treated as the canonical name for all the aliases. PTR record will be generated only for the primary name.",
	},
	"zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandHostName(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.HostName {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HostNameModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HostNameModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.HostName {
	if m == nil {
		return nil
	}
	to := &ipam.HostName{
		Alias:       flex.ExpandBoolPointer(m.Alias),
		Name:        flex.ExpandString(m.Name),
		PrimaryName: flex.ExpandBoolPointer(m.PrimaryName),
		Zone:        flex.ExpandString(m.Zone),
	}
	return to
}

func FlattenHostName(ctx context.Context, from *ipam.HostName, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostNameAttrTypes)
	}
	m := HostNameModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HostNameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HostNameModel) Flatten(ctx context.Context, from *ipam.HostName, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HostNameModel{}
	}
	m.Alias = types.BoolPointerValue(from.Alias)
	m.Name = flex.FlattenString(from.Name)
	m.PrimaryName = types.BoolPointerValue(from.PrimaryName)
	m.Zone = flex.FlattenString(from.Zone)
}
