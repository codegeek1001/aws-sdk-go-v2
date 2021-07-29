// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a backend environment for an Amplify app.
func (c *Client) GetBackendEnvironment(ctx context.Context, params *GetBackendEnvironmentInput, optFns ...func(*Options)) (*GetBackendEnvironmentOutput, error) {
	if params == nil {
		params = &GetBackendEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBackendEnvironment", params, optFns, c.addOperationGetBackendEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBackendEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the get backend environment request.
type GetBackendEnvironmentInput struct {

	// The unique id for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the backend environment.
	//
	// This member is required.
	EnvironmentName *string

	noSmithyDocumentSerde
}

// The result structure for the get backend environment result.
type GetBackendEnvironmentOutput struct {

	// Describes the backend environment for an Amplify app.
	//
	// This member is required.
	BackendEnvironment *types.BackendEnvironment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBackendEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBackendEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBackendEnvironment{}, middleware.After)
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
	if err = addOpGetBackendEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBackendEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBackendEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "GetBackendEnvironment",
	}
}
