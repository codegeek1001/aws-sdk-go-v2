// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a CloudFormation extension publisher. If you do not
// supply a PublisherId, and you have registered as an extension publisher,
// DescribePublisher returns information about your own publisher account. For more
// information on registering as a publisher, see:
//
// * RegisterPublisher
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterPublisher.html)
//
// *
// Publishing extensions to make them available for public use
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html)
// in the CloudFormation CLI User Guide
func (c *Client) DescribePublisher(ctx context.Context, params *DescribePublisherInput, optFns ...func(*Options)) (*DescribePublisherOutput, error) {
	if params == nil {
		params = &DescribePublisherInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePublisher", params, optFns, c.addOperationDescribePublisherMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePublisherOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePublisherInput struct {

	// The ID of the extension publisher. If you do not supply a PublisherId, and you
	// have registered as an extension publisher, DescribePublisher returns information
	// about your own publisher account.
	PublisherId *string

	noSmithyDocumentSerde
}

type DescribePublisherOutput struct {

	// The type of account used as the identity provider when registering this
	// publisher with CloudFormation.
	IdentityProvider types.IdentityProvider

	// The ID of the extension publisher.
	PublisherId *string

	// The URL to the publisher's profile with the identity provider.
	PublisherProfile *string

	// Whether the publisher is verified. Currently, all registered publishers are
	// verified.
	PublisherStatus types.PublisherStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePublisherMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribePublisher{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribePublisher{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePublisher(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePublisher(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribePublisher",
	}
}
