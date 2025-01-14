package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamHostModel struct {
	Addresses           types.List        `tfsdk:"addresses"`
	AutoGenerateRecords types.Bool        `tfsdk:"auto_generate_records"`
	Comment             types.String      `tfsdk:"comment"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	HostNames           types.List        `tfsdk:"host_names"`
	Id                  types.String      `tfsdk:"id"`
	Name                types.String      `tfsdk:"name"`
	Tags                types.Map         `tfsdk:"tags"`
	TagsAll             types.Map         `tfsdk:"tags_all"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamHostAttrTypes = map[string]attr.Type{
	"addresses":             types.ListType{ElemType: types.ObjectType{AttrTypes: HostAddressAttrTypes}},
	"auto_generate_records": types.BoolType,
	"comment":               types.StringType,
	"created_at":            timetypes.RFC3339Type{},
	"host_names":            types.ListType{ElemType: types.ObjectType{AttrTypes: HostNameAttrTypes}},
	"id":                    types.StringType,
	"name":                  types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"tags_all":              types.MapType{ElemType: types.StringType},
	"updated_at":            timetypes.RFC3339Type{},
}

var IpamHostResourceSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HostAddressResourceSchemaAttributes,
		},
		Optional: true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The list of all addresses associated with the IPAM host, which may be in different IP spaces.",
	},
	"auto_generate_records": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag specifies if resource records have to be auto generated for the host.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"host_names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: HostNameResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The name records to be generated for the host.  This field is required if _auto_generate_records_ is true.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the IPAM host in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: `The tags for the IPAM host in JSON format including default tags.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandIpamHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamHost {
	if m == nil {
		return nil
	}
	to := &ipam.IpamHost{
		Addresses:           flex.ExpandFrameworkListNestedBlock(ctx, m.Addresses, diags, ExpandHostAddress),
		AutoGenerateRecords: flex.ExpandBoolPointer(m.AutoGenerateRecords),
		Comment:             flex.ExpandStringPointer(m.Comment),
		HostNames:           flex.ExpandFrameworkListNestedBlock(ctx, m.HostNames, diags, ExpandHostName),
		Name:                flex.ExpandString(m.Name),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenIpamHost(ctx context.Context, from *ipam.IpamHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamHostAttrTypes)
	}
	m := IpamHostModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, IpamHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamHostModel) Flatten(ctx context.Context, from *ipam.IpamHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamHostModel{}
	}
	var i []HostAddressModel

	m.Addresses = flex.FlattenFrameworkListsNestedBlock(ctx, from.Addresses, flex.ExpandList(ctx, m.Addresses, i, diags), HostAddressAttrTypes, diags, FlattenHostAddress)
	m.AutoGenerateRecords = types.BoolPointerValue(from.AutoGenerateRecords)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.HostNames = flex.FlattenFrameworkListNestedBlock(ctx, from.HostNames, HostNameAttrTypes, diags, FlattenHostName)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
