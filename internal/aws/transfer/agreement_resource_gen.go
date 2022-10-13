// Code generated by generators/resource/main.go; DO NOT EDIT.

package transfer

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_transfer_agreement", agreementResource)
}

// agreementResource returns the Terraform awscc_transfer_agreement resource.
// This Terraform resource corresponds to the CloudFormation AWS::Transfer::Agreement resource.
func agreementResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_role": {
			// Property: AccessRole
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the access role for the agreement.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:.*role/.*",
			//   "type": "string"
			// }
			Description: "Specifies the access role for the agreement.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:.*role/.*"), ""),
			},
		},
		"agreement_id": {
			// Property: AgreementId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the agreement.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^a-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the agreement.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
			//   "maxLength": 1600,
			//   "minLength": 20,
			//   "pattern": "arn:.*",
			//   "type": "string"
			// }
			Description: "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"base_directory": {
			// Property: BaseDirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the base directory for the agreement.",
			//   "maxLength": 1024,
			//   "pattern": "^$|/.*",
			//   "type": "string"
			// }
			Description: "Specifies the base directory for the agreement.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1024),
				validate.StringMatch(regexp.MustCompile("^$|/.*"), ""),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A textual description for the agreement.",
			//   "maxLength": 200,
			//   "minLength": 1,
			//   "pattern": "^[\\w\\- ]*$",
			//   "type": "string"
			// }
			Description: "A textual description for the agreement.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 200),
				validate.StringMatch(regexp.MustCompile("^[\\w\\- ]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"local_profile_id": {
			// Property: LocalProfileId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the local profile.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^p-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the local profile.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(19, 19),
				validate.StringMatch(regexp.MustCompile("^p-([0-9a-f]{17})$"), ""),
			},
		},
		"partner_profile_id": {
			// Property: PartnerProfileId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the partner profile.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^p-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the partner profile.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(19, 19),
				validate.StringMatch(regexp.MustCompile("^p-([0-9a-f]{17})$"), ""),
			},
		},
		"server_id": {
			// Property: ServerId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the server.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^s-([0-9a-f]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the server.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(19, 19),
				validate.StringMatch(regexp.MustCompile("^s-([0-9a-f]{17})$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the status of the agreement.",
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the status of the agreement.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE",
					"INACTIVE",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Creates a key-value pair for a specific resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The name assigned to the tag that you create.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Contains one or more values that you assigned to the key name you create.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name assigned to the tag that you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "Contains one or more values that you assigned to the key name you create.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Transfer::Agreement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Agreement").WithTerraformTypeName("awscc_transfer_agreement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role":        "AccessRole",
		"agreement_id":       "AgreementId",
		"arn":                "Arn",
		"base_directory":     "BaseDirectory",
		"description":        "Description",
		"key":                "Key",
		"local_profile_id":   "LocalProfileId",
		"partner_profile_id": "PartnerProfileId",
		"server_id":          "ServerId",
		"status":             "Status",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}