// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_medialive_multiplex", multiplexResource)
}

// multiplexResource returns the Terraform awscc_medialive_multiplex resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaLive::Multiplex resource.
func multiplexResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique arn of the multiplex.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique arn of the multiplex.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZones
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of availability zones for the multiplex.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"availability_zones": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of availability zones for the multiplex.",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Destinations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of the multiplex output destinations.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Multiplex MediaConnect output destination settings.",
		//	    "properties": {
		//	      "MultiplexMediaConnectOutputDestinationSettings": {
		//	        "additionalProperties": false,
		//	        "description": "Multiplex MediaConnect output destination settings.",
		//	        "properties": {
		//	          "EntitlementArn": {
		//	            "description": "The MediaConnect entitlement ARN available as a Flow source.",
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        }
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"destinations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: MultiplexMediaConnectOutputDestinationSettings
					"multiplex_media_connect_output_destination_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: EntitlementArn
							"entitlement_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The MediaConnect entitlement ARN available as a Flow source.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtLeast(1),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Multiplex MediaConnect output destination settings.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of the multiplex output destinations.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique id of the multiplex.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique id of the multiplex.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MultiplexSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for a multiplex event.",
		//	  "properties": {
		//	    "MaximumVideoBufferDelayMilliseconds": {
		//	      "description": "Maximum video buffer delay in milliseconds.",
		//	      "maximum": 3000,
		//	      "minimum": 800,
		//	      "type": "integer"
		//	    },
		//	    "TransportStreamBitrate": {
		//	      "description": "Transport stream bit rate.",
		//	      "maximum": 100000000,
		//	      "minimum": 1000000,
		//	      "type": "integer"
		//	    },
		//	    "TransportStreamId": {
		//	      "description": "Transport stream ID.",
		//	      "maximum": 65535,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "TransportStreamReservedBitrate": {
		//	      "description": "Transport stream reserved bit rate.",
		//	      "maximum": 100000000,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "TransportStreamBitrate",
		//	    "TransportStreamId"
		//	  ],
		//	  "type": "object"
		//	}
		"multiplex_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaximumVideoBufferDelayMilliseconds
				"maximum_video_buffer_delay_milliseconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Maximum video buffer delay in milliseconds.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(800, 3000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TransportStreamBitrate
				"transport_stream_bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Transport stream bit rate.",
					Required:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1000000, 100000000),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: TransportStreamId
				"transport_stream_id": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Transport stream ID.",
					Required:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 65535),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: TransportStreamReservedBitrate
				"transport_stream_reserved_bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Transport stream reserved bit rate.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 100000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for a multiplex event.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of multiplex.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of multiplex.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PipelinesRunningCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of currently healthy pipelines.",
		//	  "type": "integer"
		//	}
		"pipelines_running_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of currently healthy pipelines.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProgramCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of programs in the multiplex.",
		//	  "type": "integer"
		//	}
		"program_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of programs in the multiplex.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATING",
		//	    "CREATE_FAILED",
		//	    "IDLE",
		//	    "STARTING",
		//	    "RUNNING",
		//	    "RECOVERING",
		//	    "STOPPING",
		//	    "DELETING",
		//	    "DELETED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of key-value pairs.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of key-value pairs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaLive::Multiplex",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaLive::Multiplex").WithTerraformTypeName("awscc_medialive_multiplex")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"availability_zones": "AvailabilityZones",
		"destinations":       "Destinations",
		"entitlement_arn":    "EntitlementArn",
		"id":                 "Id",
		"key":                "Key",
		"maximum_video_buffer_delay_milliseconds":             "MaximumVideoBufferDelayMilliseconds",
		"multiplex_media_connect_output_destination_settings": "MultiplexMediaConnectOutputDestinationSettings",
		"multiplex_settings":                                  "MultiplexSettings",
		"name":                                                "Name",
		"pipelines_running_count":                             "PipelinesRunningCount",
		"program_count":                                       "ProgramCount",
		"state":                                               "State",
		"tags":                                                "Tags",
		"transport_stream_bitrate":                            "TransportStreamBitrate",
		"transport_stream_id":                                 "TransportStreamId",
		"transport_stream_reserved_bitrate":                   "TransportStreamReservedBitrate",
		"value":                                               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
