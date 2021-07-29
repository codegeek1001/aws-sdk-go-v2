// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a fleet scaling policy. Once deleted, the policy is no longer in force
// and GameLift removes all record of it. To delete a scaling policy, specify both
// the scaling policy name and the fleet ID it is associated with. To temporarily
// suspend scaling policies, call StopFleetActions. This operation suspends all
// policies for the fleet. Related actions DescribeFleetCapacity |
// UpdateFleetCapacity | DescribeEC2InstanceLimits | PutScalingPolicy |
// DescribeScalingPolicies | DeleteScalingPolicy | StopFleetActions |
// StartFleetActions | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteScalingPolicy(ctx context.Context, params *DeleteScalingPolicyInput, optFns ...func(*Options)) (*DeleteScalingPolicyOutput, error) {
	if params == nil {
		params = &DeleteScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteScalingPolicy", params, optFns, c.addOperationDeleteScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteScalingPolicyInput struct {

	// A unique identifier for the fleet to be deleted. You can use either the fleet ID
	// or ARN value.
	//
	// This member is required.
	FleetId *string

	// A descriptive label that is associated with a fleet's scaling policy. Policy
	// names do not need to be unique.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteScalingPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteScalingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteScalingPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteScalingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteScalingPolicy",
	}
}
