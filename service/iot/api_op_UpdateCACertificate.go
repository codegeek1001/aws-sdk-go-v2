// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a registered CA certificate.
func (c *Client) UpdateCACertificate(ctx context.Context, params *UpdateCACertificateInput, optFns ...func(*Options)) (*UpdateCACertificateOutput, error) {
	if params == nil {
		params = &UpdateCACertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCACertificate", params, optFns, c.addOperationUpdateCACertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCACertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the UpdateCACertificate operation.
type UpdateCACertificateInput struct {

	// The CA certificate identifier.
	//
	// This member is required.
	CertificateId *string

	// The new value for the auto registration status. Valid values are: "ENABLE" or
	// "DISABLE".
	NewAutoRegistrationStatus types.AutoRegistrationStatus

	// The updated status of the CA certificate. Note: The status value
	// REGISTER_INACTIVE is deprecated and should not be used.
	NewStatus types.CACertificateStatus

	// Information about the registration configuration.
	RegistrationConfig *types.RegistrationConfig

	// If true, removes auto registration.
	RemoveAutoRegistration bool

	noSmithyDocumentSerde
}

type UpdateCACertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCACertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCACertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCACertificate{}, middleware.After)
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
	if err = addOpUpdateCACertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCACertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCACertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateCACertificate",
	}
}
