// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the settings that were used when your QuickSight subscription was
// first created in this AWS Account.
func (c *Client) DescribeAccountSettings(ctx context.Context, params *DescribeAccountSettingsInput, optFns ...func(*Options)) (*DescribeAccountSettingsOutput, error) {
	stack := middleware.NewStack("DescribeAccountSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAccountSettingsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribeAccountSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeAccountSettings",
			Err:           err,
		}
	}
	out := result.(*DescribeAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountSettingsInput struct {

	// The ID for the AWS account that contains the QuickSight namespaces that you want
	// to list.
	//
	// This member is required.
	AwsAccountId *string
}

type DescribeAccountSettingsOutput struct {

	// The settings associated with the QuickSight subscription associated with this
	// AWS account. This information includes the edition of Amazon QuickSight that you
	// subscribed to (Standard or Enterprise) and the notification email for the
	// QuickSight subscription. The QuickSight console, the QuickSight subscription is
	// sometimes referred to as a QuickSight "account" even though it is technically
	// not an account, but a subscription in your AWS account.
	AccountSettings *types.AccountSettings

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAccountSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountSettings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccountSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAccountSettings",
	}
}