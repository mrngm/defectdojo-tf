// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_products

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ProductsResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"authorization_groups": schema.ListAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
			},
			"business_criticality": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "* `very high` - Very High\n* `high` - High\n* `medium` - Medium\n* `low` - Low\n* `very low` - Very Low\n* `none` - None",
				MarkdownDescription: "* `very high` - Very High\n* `high` - High\n* `medium` - Medium\n* `low` - Low\n* `very low` - Very Low\n* `none` - None",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"very high",
						"high",
						"medium",
						"low",
						"very low",
						"none",
						"",
						"",
					),
				},
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 4000),
				},
			},
			"disable_sla_breach_notifications": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Disable SLA breach notifications if configured in the global settings",
				MarkdownDescription: "Disable SLA breach notifications if configured in the global settings",
			},
			"enable_full_risk_acceptance": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Allows full risk acceptance using a risk acceptance form, expiration date, uploaded proof, etc.",
				MarkdownDescription: "Allows full risk acceptance using a risk acceptance form, expiration date, uploaded proof, etc.",
			},
			"enable_product_tag_inheritance": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enables product tag inheritance. Any tags added on a product will automatically be added to all Engagements, Tests, and Findings",
				MarkdownDescription: "Enables product tag inheritance. Any tags added on a product will automatically be added to all Engagements, Tests, and Findings",
			},
			"enable_simple_risk_acceptance": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Allows simple risk acceptance by checking/unchecking a checkbox.",
				MarkdownDescription: "Allows simple risk acceptance by checking/unchecking a checkbox.",
			},
			"external_audience": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Specify if the application is used by people outside the organization.",
				MarkdownDescription: "Specify if the application is used by people outside the organization.",
			},
			"findings_count": schema.Int64Attribute{
				Computed: true,
			},
			"findings_list": schema.ListAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Computed:            true,
				Description:         "A unique integer value identifying this product.",
				MarkdownDescription: "A unique integer value identifying this product.",
			},
			"internet_accessible": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Specify if the application is accessible from the public internet.",
				MarkdownDescription: "Specify if the application is accessible from the public internet.",
			},
			"lifecycle": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "* `construction` - Construction\n* `production` - Production\n* `retirement` - Retirement",
				MarkdownDescription: "* `construction` - Construction\n* `production` - Production\n* `retirement` - Retirement",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"construction",
						"production",
						"retirement",
						"",
						"",
					),
				},
			},
			"members": schema.ListAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
			},
			"origin": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "* `third party library` - Third Party Library\n* `purchased` - Purchased\n* `contractor` - Contractor Developed\n* `internal` - Internally Developed\n* `open source` - Open Source\n* `outsourced` - Outsourced",
				MarkdownDescription: "* `third party library` - Third Party Library\n* `purchased` - Purchased\n* `contractor` - Contractor Developed\n* `internal` - Internally Developed\n* `open source` - Open Source\n* `outsourced` - Outsourced",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"third party library",
						"purchased",
						"contractor",
						"internal",
						"open source",
						"outsourced",
						"",
						"",
					),
				},
			},
			"platform": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "* `web service` - API\n* `desktop` - Desktop\n* `iot` - Internet of Things\n* `mobile` - Mobile\n* `web` - Web",
				MarkdownDescription: "* `web service` - API\n* `desktop` - Desktop\n* `iot` - Internet of Things\n* `mobile` - Mobile\n* `web` - Web",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"web service",
						"desktop",
						"iot",
						"mobile",
						"web",
						"",
						"",
					),
				},
			},
			"prod_numeric_grade": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			"prod_type": schema.Int64Attribute{
				Required: true,
			},
			"product_manager": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"product_meta": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
						},
						"value": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: ProductMetaType{
						ObjectType: types.ObjectType{
							AttrTypes: ProductMetaValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"regulations": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
			},
			"revenue": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Estimate the application's revenue.",
				MarkdownDescription: "Estimate the application's revenue.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^-?\\d{0,13}(?:\\.\\d{0,2})?$"), ""),
				},
			},
			"sla_configuration": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"tags": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"team_manager": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"technical_contact": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"user_records": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Estimate the number of user records within the application.",
				MarkdownDescription: "Estimate the number of user records within the application.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
		},
	}
}

type ProductsModel struct {
	AuthorizationGroups           types.List   `tfsdk:"authorization_groups"`
	BusinessCriticality           types.String `tfsdk:"business_criticality"`
	Created                       types.String `tfsdk:"created"`
	Description                   types.String `tfsdk:"description"`
	DisableSlaBreachNotifications types.Bool   `tfsdk:"disable_sla_breach_notifications"`
	EnableFullRiskAcceptance      types.Bool   `tfsdk:"enable_full_risk_acceptance"`
	EnableProductTagInheritance   types.Bool   `tfsdk:"enable_product_tag_inheritance"`
	EnableSimpleRiskAcceptance    types.Bool   `tfsdk:"enable_simple_risk_acceptance"`
	ExternalAudience              types.Bool   `tfsdk:"external_audience"`
	FindingsCount                 types.Int64  `tfsdk:"findings_count"`
	FindingsList                  types.List   `tfsdk:"findings_list"`
	Id                            types.Int64  `tfsdk:"id"`
	InternetAccessible            types.Bool   `tfsdk:"internet_accessible"`
	Lifecycle                     types.String `tfsdk:"lifecycle"`
	Members                       types.List   `tfsdk:"members"`
	Name                          types.String `tfsdk:"name"`
	Origin                        types.String `tfsdk:"origin"`
	Platform                      types.String `tfsdk:"platform"`
	ProdNumericGrade              types.Int64  `tfsdk:"prod_numeric_grade"`
	ProdType                      types.Int64  `tfsdk:"prod_type"`
	ProductManager                types.Int64  `tfsdk:"product_manager"`
	ProductMeta                   types.List   `tfsdk:"product_meta"`
	Regulations                   types.List   `tfsdk:"regulations"`
	Revenue                       types.String `tfsdk:"revenue"`
	SlaConfiguration              types.Int64  `tfsdk:"sla_configuration"`
	Tags                          types.List   `tfsdk:"tags"`
	TeamManager                   types.Int64  `tfsdk:"team_manager"`
	TechnicalContact              types.Int64  `tfsdk:"technical_contact"`
	UserRecords                   types.Int64  `tfsdk:"user_records"`
}

var _ basetypes.ObjectTypable = ProductMetaType{}

type ProductMetaType struct {
	basetypes.ObjectType
}

func (t ProductMetaType) Equal(o attr.Type) bool {
	other, ok := o.(ProductMetaType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ProductMetaType) String() string {
	return "ProductMetaType"
}

func (t ProductMetaType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	valueAttribute, ok := attributes["value"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`value is missing from object`)

		return nil, diags
	}

	valueVal, ok := valueAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`value expected to be basetypes.StringValue, was: %T`, valueAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ProductMetaValue{
		Name:  nameVal,
		Value: valueVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewProductMetaValueNull() ProductMetaValue {
	return ProductMetaValue{
		state: attr.ValueStateNull,
	}
}

func NewProductMetaValueUnknown() ProductMetaValue {
	return ProductMetaValue{
		state: attr.ValueStateUnknown,
	}
}

func NewProductMetaValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ProductMetaValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ProductMetaValue Attribute Value",
				"While creating a ProductMetaValue value, a missing attribute value was detected. "+
					"A ProductMetaValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProductMetaValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ProductMetaValue Attribute Type",
				"While creating a ProductMetaValue value, an invalid attribute value was detected. "+
					"A ProductMetaValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProductMetaValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ProductMetaValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ProductMetaValue Attribute Value",
				"While creating a ProductMetaValue value, an extra attribute value was detected. "+
					"A ProductMetaValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ProductMetaValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewProductMetaValueUnknown(), diags
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewProductMetaValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	valueAttribute, ok := attributes["value"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`value is missing from object`)

		return NewProductMetaValueUnknown(), diags
	}

	valueVal, ok := valueAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`value expected to be basetypes.StringValue, was: %T`, valueAttribute))
	}

	if diags.HasError() {
		return NewProductMetaValueUnknown(), diags
	}

	return ProductMetaValue{
		Name:  nameVal,
		Value: valueVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewProductMetaValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ProductMetaValue {
	object, diags := NewProductMetaValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewProductMetaValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ProductMetaType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewProductMetaValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewProductMetaValueUnknown(), nil
	}

	if in.IsNull() {
		return NewProductMetaValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewProductMetaValueMust(ProductMetaValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ProductMetaType) ValueType(ctx context.Context) attr.Value {
	return ProductMetaValue{}
}

var _ basetypes.ObjectValuable = ProductMetaValue{}

type ProductMetaValue struct {
	Name  basetypes.StringValue `tfsdk:"name"`
	Value basetypes.StringValue `tfsdk:"value"`
	state attr.ValueState
}

func (v ProductMetaValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["value"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Value.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["value"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ProductMetaValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ProductMetaValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ProductMetaValue) String() string {
	return "ProductMetaValue"
}

func (v ProductMetaValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"name":  v.Name,
			"value": v.Value,
		})

	return objVal, diags
}

func (v ProductMetaValue) Equal(o attr.Value) bool {
	other, ok := o.(ProductMetaValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Value.Equal(other.Value) {
		return false
	}

	return true
}

func (v ProductMetaValue) Type(ctx context.Context) attr.Type {
	return ProductMetaType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ProductMetaValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"name":  basetypes.StringType{},
		"value": basetypes.StringType{},
	}
}
