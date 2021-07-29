// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an API destination.
func (c *Client) UpdateApiDestination(ctx context.Context, params *UpdateApiDestinationInput, optFns ...func(*Options)) (*UpdateApiDestinationOutput, error) {
	if params == nil {
		params = &UpdateApiDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApiDestination", params, optFns, c.addOperationUpdateApiDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApiDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApiDestinationInput struct {

	// The name of the API destination to update.
	//
	// This member is required.
	Name *string

	// The ARN of the connection to use for the API destination.
	ConnectionArn *string

	// The name of the API destination to update.
	Description *string

	// The method to use for the API destination.
	HttpMethod types.ApiDestinationHttpMethod

	// The URL to the endpoint to use for the API destination.
	InvocationEndpoint *string

	// The maximum number of invocations per second to send to the API destination.
	InvocationRateLimitPerSecond *int32

	noSmithyDocumentSerde
}

type UpdateApiDestinationOutput struct {

	// The ARN of the API destination that was updated.
	ApiDestinationArn *string

	// The state of the API destination that was updated.
	ApiDestinationState types.ApiDestinationState

	// A time stamp for the time that the API destination was created.
	CreationTime *time.Time

	// A time stamp for the time that the API destination was last modified.
	LastModifiedTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApiDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApiDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApiDestination{}, middleware.After)
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
	if err = addOpUpdateApiDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApiDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApiDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "UpdateApiDestination",
	}
}
