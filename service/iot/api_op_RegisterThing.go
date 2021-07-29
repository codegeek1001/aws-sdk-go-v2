// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provisions a thing in the device registry. RegisterThing calls other AWS IoT
// control plane APIs. These calls might exceed your account level  AWS IoT
// Throttling Limits
// (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_iot)
// and cause throttle errors. Please contact AWS Customer Support
// (https://console.aws.amazon.com/support/home) to raise your throttling limits if
// necessary.
func (c *Client) RegisterThing(ctx context.Context, params *RegisterThingInput, optFns ...func(*Options)) (*RegisterThingOutput, error) {
	if params == nil {
		params = &RegisterThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterThing", params, optFns, c.addOperationRegisterThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterThingInput struct {

	// The provisioning template. See Provisioning Devices That Have Device
	// Certificates
	// (https://docs.aws.amazon.com/iot/latest/developerguide/provision-w-cert.html)
	// for more information.
	//
	// This member is required.
	TemplateBody *string

	// The parameters for provisioning a thing. See Provisioning Templates
	// (https://docs.aws.amazon.com/iot/latest/developerguide/provision-template.html)
	// for more information.
	Parameters map[string]string

	noSmithyDocumentSerde
}

type RegisterThingOutput struct {

	// The certificate data, in PEM format.
	CertificatePem *string

	// ARNs for the generated resources.
	ResourceArns map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterThing{}, middleware.After)
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
	if err = addOpRegisterThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterThing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "RegisterThing",
	}
}
