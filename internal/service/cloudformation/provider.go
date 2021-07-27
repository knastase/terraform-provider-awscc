package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

// Provider is the interface implemented by AWS CloudFormation client providers.
type Provider interface {
	// CloudFormationClient returns an AWS CloudFormation client.
	CloudFormationClient(context.Context) *cloudformation.Client

	// RoleARN returns an AWS CloudFormation service role ARN.
	RoleARN(context.Context) string
}
