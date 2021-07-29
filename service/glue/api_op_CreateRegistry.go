// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new registry which may be used to hold a collection of schemas.
func (c *Client) CreateRegistry(ctx context.Context, params *CreateRegistryInput, optFns ...func(*Options)) (*CreateRegistryOutput, error) {
	if params == nil {
		params = &CreateRegistryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegistry", params, optFns, c.addOperationCreateRegistryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegistryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegistryInput struct {

	// Name of the registry to be created of max length of 255, and may only contain
	// letters, numbers, hyphen, underscore, dollar sign, or hash mark. No whitespace.
	//
	// This member is required.
	RegistryName *string

	// A description of the registry. If description is not provided, there will not be
	// any default value for this.
	Description *string

	// Amazon Web Services tags that contain a key value pair and may be searched by
	// console, command line, or API.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRegistryOutput struct {

	// A description of the registry.
	Description *string

	// The Amazon Resource Name (ARN) of the newly created registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// The tags for the registry.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRegistryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegistry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegistry{}, middleware.After)
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
	if err = addOpCreateRegistryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegistry(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegistry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateRegistry",
	}
}
