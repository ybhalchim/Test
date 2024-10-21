package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type OptionItemModel struct {
	Group       types.String `tfsdk:"group"`
	OptionCode  types.String `tfsdk:"option_code"`
	OptionValue types.String `tfsdk:"option_value"`
	Type        types.String `tfsdk:"type"`
}

var OptionItemAttrTypes = map[string]attr.Type{
	"group":        types.StringType,
	"option_code":  types.StringType,
	"option_value": types.StringType,
	"type":         types.StringType,
}

var OptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("option_code"), path.MatchRelative().AtParent().AtName("group")),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"option_code": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("option_code"), path.MatchRelative().AtParent().AtName("group")),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"option_value": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("option_code")),
		},
		MarkdownDescription: "The option value.",
	},
	"type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("group", "option"),
		},
		MarkdownDescription: "The type of item.  Valid values are: * _group_ * _option_",
	},
}

func ExpandOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.OptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *OptionItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.OptionItem {
	if m == nil {
		return nil
	}
	to := &ipam.OptionItem{
		Group:       m.Group.ValueStringPointer(),
		OptionCode:  m.OptionCode.ValueStringPointer(),
		OptionValue: m.OptionValue.ValueStringPointer(),
		Type:        m.Type.ValueStringPointer(),
	}
	return to
}

func FlattenOptionItem(ctx context.Context, from *ipam.OptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OptionItemAttrTypes)
	}
	m := OptionItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *OptionItemModel) Flatten(ctx context.Context, from *ipam.OptionItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = OptionItemModel{}
	}
	m.Group = flex.FlattenStringPointer(from.Group)
	m.OptionCode = flex.FlattenStringPointer(from.OptionCode)
	m.OptionValue = flex.FlattenStringPointer(from.OptionValue)
	m.Type = flex.FlattenStringPointer(from.Type)
}
