// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns current default log-levels, or log levels by resource types, could be
// for wireless device log options or wireless gateway log options.
func (c *Client) GetLogLevelsByResourceTypes(ctx context.Context, params *GetLogLevelsByResourceTypesInput, optFns ...func(*Options)) (*GetLogLevelsByResourceTypesOutput, error) {
	if params == nil {
		params = &GetLogLevelsByResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLogLevelsByResourceTypes", params, optFns, c.addOperationGetLogLevelsByResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLogLevelsByResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLogLevelsByResourceTypesInput struct {
	noSmithyDocumentSerde
}

type GetLogLevelsByResourceTypesOutput struct {

	// The log level for a log message.
	DefaultLogLevel types.LogLevel

	// The list of wireless device log options.
	WirelessDeviceLogOptions []types.WirelessDeviceLogOption

	// The list of wireless gateway log options.
	WirelessGatewayLogOptions []types.WirelessGatewayLogOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLogLevelsByResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLogLevelsByResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLogLevelsByResourceTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLogLevelsByResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLogLevelsByResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetLogLevelsByResourceTypes",
	}
}
