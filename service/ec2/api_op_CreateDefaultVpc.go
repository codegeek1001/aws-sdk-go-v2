// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a default VPC with a size /16 IPv4 CIDR block and a default subnet in
// each Availability Zone. For more information about the components of a default
// VPC, see Default VPC and Default Subnets
// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html) in the
// Amazon Virtual Private Cloud User Guide. You cannot specify the components of
// the default VPC yourself. If you deleted your previous default VPC, you can
// create a default VPC. You cannot have more than one default VPC per Region. If
// your account supports EC2-Classic, you cannot use this action to create a
// default VPC in a Region that supports EC2-Classic. If you want a default VPC in
// a Region that supports EC2-Classic, see "I really want a default VPC for my
// existing EC2 account. Is that possible?" in the Default VPCs FAQ
// (http://aws.amazon.com/vpc/faqs/#Default_VPCs).
func (c *Client) CreateDefaultVpc(ctx context.Context, params *CreateDefaultVpcInput, optFns ...func(*Options)) (*CreateDefaultVpcOutput, error) {
	if params == nil {
		params = &CreateDefaultVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDefaultVpc", params, optFns, c.addOperationCreateDefaultVpcMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDefaultVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDefaultVpcInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type CreateDefaultVpcOutput struct {

	// Information about the VPC.
	Vpc *types.Vpc

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDefaultVpcMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateDefaultVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDefaultVpc{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDefaultVpc(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDefaultVpc(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateDefaultVpc",
	}
}
