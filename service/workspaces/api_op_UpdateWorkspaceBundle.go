// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a WorkSpace bundle with a new image. For more information about updating
// WorkSpace bundles, see  Update a Custom WorkSpaces Bundle
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-custom-bundle.html).
// Existing WorkSpaces aren't automatically updated when you update the bundle that
// they're based on. To update existing WorkSpaces that are based on a bundle that
// you've updated, you must either rebuild the WorkSpaces or delete and recreate
// them.
func (c *Client) UpdateWorkspaceBundle(ctx context.Context, params *UpdateWorkspaceBundleInput, optFns ...func(*Options)) (*UpdateWorkspaceBundleOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspaceBundle", params, optFns, c.addOperationUpdateWorkspaceBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceBundleInput struct {

	// The identifier of the bundle.
	BundleId *string

	// The identifier of the image.
	ImageId *string

	noSmithyDocumentSerde
}

type UpdateWorkspaceBundleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkspaceBundle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkspaceBundle{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspaceBundle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspaceBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "UpdateWorkspaceBundle",
	}
}
