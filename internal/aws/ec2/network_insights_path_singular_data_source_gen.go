// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_network_insights_path", networkInsightsPathDataSource)
}

// networkInsightsPathDataSource returns the Terraform awscc_ec2_network_insights_path data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::NetworkInsightsPath resource.
func networkInsightsPathDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"created_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Destination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"destination_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"destination_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FilterAtDestination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DestinationAddress": {
		//	      "type": "string"
		//	    },
		//	    "DestinationPortRange": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "FromPort": {
		//	          "type": "integer"
		//	        },
		//	        "ToPort": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SourceAddress": {
		//	      "type": "string"
		//	    },
		//	    "SourcePortRange": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "FromPort": {
		//	          "type": "integer"
		//	        },
		//	        "ToPort": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"filter_at_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DestinationAddress
				"destination_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DestinationPortRange
				"destination_port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: FromPort
						"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ToPort
						"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourceAddress
				"source_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourcePortRange
				"source_port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: FromPort
						"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ToPort
						"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FilterAtSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DestinationAddress": {
		//	      "type": "string"
		//	    },
		//	    "DestinationPortRange": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "FromPort": {
		//	          "type": "integer"
		//	        },
		//	        "ToPort": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SourceAddress": {
		//	      "type": "string"
		//	    },
		//	    "SourcePortRange": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "FromPort": {
		//	          "type": "integer"
		//	        },
		//	        "ToPort": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"filter_at_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DestinationAddress
				"destination_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DestinationPortRange
				"destination_port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: FromPort
						"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ToPort
						"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourceAddress
				"source_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourcePortRange
				"source_port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: FromPort
						"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ToPort
						"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsPathArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_path_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsPathId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_path_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "tcp",
		//	    "udp"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Source
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SourceIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"source_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::NetworkInsightsPath",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsPath").WithTerraformTypeName("awscc_ec2_network_insights_path")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_date":              "CreatedDate",
		"destination":               "Destination",
		"destination_address":       "DestinationAddress",
		"destination_arn":           "DestinationArn",
		"destination_ip":            "DestinationIp",
		"destination_port":          "DestinationPort",
		"destination_port_range":    "DestinationPortRange",
		"filter_at_destination":     "FilterAtDestination",
		"filter_at_source":          "FilterAtSource",
		"from_port":                 "FromPort",
		"key":                       "Key",
		"network_insights_path_arn": "NetworkInsightsPathArn",
		"network_insights_path_id":  "NetworkInsightsPathId",
		"protocol":                  "Protocol",
		"source":                    "Source",
		"source_address":            "SourceAddress",
		"source_arn":                "SourceArn",
		"source_ip":                 "SourceIp",
		"source_port_range":         "SourcePortRange",
		"tags":                      "Tags",
		"to_port":                   "ToPort",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
