// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a physical device from a placement.
func (c *Client) DisassociateDeviceFromPlacement(ctx context.Context, params *DisassociateDeviceFromPlacementInput, optFns ...func(*Options)) (*DisassociateDeviceFromPlacementOutput, error) {
	if params == nil {
		params = &DisassociateDeviceFromPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDeviceFromPlacement", params, optFns, c.addOperationDisassociateDeviceFromPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDeviceFromPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDeviceFromPlacementInput struct {

	// The device ID that should be removed from the placement.
	//
	// This member is required.
	DeviceTemplateName *string

	// The name of the placement that the device should be removed from.
	//
	// This member is required.
	PlacementName *string

	// The name of the project that contains the placement.
	//
	// This member is required.
	ProjectName *string

	noSmithyDocumentSerde
}

type DisassociateDeviceFromPlacementOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateDeviceFromPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateDeviceFromPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateDeviceFromPlacement{}, middleware.After)
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
	if err = addOpDisassociateDeviceFromPlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDeviceFromPlacement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateDeviceFromPlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "DisassociateDeviceFromPlacement",
	}
}
