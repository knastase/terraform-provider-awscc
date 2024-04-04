// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package securityhub

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_securityhub_product_subscription", productSubscriptionResource)
}

// productSubscriptionResource returns the Terraform awscc_securityhub_product_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::SecurityHub::ProductSubscription resource.
func productSubscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ProductArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The generic ARN of the product being subscribed to",
		//	  "pattern": "arn:aws\\S*:securityhub:\\S*",
		//	  "type": "string"
		//	}
		"product_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The generic ARN of the product being subscribed to",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws\\S*:securityhub:\\S*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProductSubscriptionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the product subscription for the account",
		//	  "pattern": "arn:aws\\S*:securityhub:\\S*",
		//	  "type": "string"
		//	}
		"product_subscription_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the product subscription for the account",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::SecurityHub::ProductSubscription resource represents a subscription to a service that is allowed to generate findings for your Security Hub account. One product subscription resource is created for each product enabled.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityHub::ProductSubscription").WithTerraformTypeName("awscc_securityhub_product_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"product_arn":              "ProductArn",
		"product_subscription_arn": "ProductSubscriptionArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}