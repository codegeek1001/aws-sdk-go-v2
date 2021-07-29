// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to confirm that the attachment has been uploaded using the pre-signed
// URL provided in StartAttachmentUpload API.
func (c *Client) CompleteAttachmentUpload(ctx context.Context, params *CompleteAttachmentUploadInput, optFns ...func(*Options)) (*CompleteAttachmentUploadOutput, error) {
	if params == nil {
		params = &CompleteAttachmentUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteAttachmentUpload", params, optFns, c.addOperationCompleteAttachmentUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteAttachmentUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteAttachmentUploadInput struct {

	// A list of unique identifiers for the attachments.
	//
	// This member is required.
	AttachmentIds []string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	//
	// This member is required.
	ClientToken *string

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	noSmithyDocumentSerde
}

type CompleteAttachmentUploadOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteAttachmentUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteAttachmentUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteAttachmentUpload{}, middleware.After)
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
	if err = addIdempotencyToken_opCompleteAttachmentUploadMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCompleteAttachmentUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteAttachmentUpload(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCompleteAttachmentUpload struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCompleteAttachmentUpload) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCompleteAttachmentUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CompleteAttachmentUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CompleteAttachmentUploadInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCompleteAttachmentUploadMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCompleteAttachmentUpload{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCompleteAttachmentUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CompleteAttachmentUpload",
	}
}
