// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The example tests how requests serialize different timestamp formats in the URI
// path.
func (c *Client) HttpRequestWithLabelsAndTimestampFormat(ctx context.Context, params *HttpRequestWithLabelsAndTimestampFormatInput, optFns ...func(*Options)) (*HttpRequestWithLabelsAndTimestampFormatOutput, error) {
	if params == nil {
		params = &HttpRequestWithLabelsAndTimestampFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithLabelsAndTimestampFormat", params, optFns, c.addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithLabelsAndTimestampFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsAndTimestampFormatInput struct {

	// This member is required.
	DefaultFormat *time.Time

	// This member is required.
	MemberDateTime *time.Time

	// This member is required.
	MemberEpochSeconds *time.Time

	// This member is required.
	MemberHttpDate *time.Time

	// This member is required.
	TargetDateTime *time.Time

	// This member is required.
	TargetEpochSeconds *time.Time

	// This member is required.
	TargetHttpDate *time.Time

	noSmithyDocumentSerde
}

type HttpRequestWithLabelsAndTimestampFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabelsAndTimestampFormat",
	}
}
