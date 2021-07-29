// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the snapshots for the specified WorkSpace.
func (c *Client) DescribeWorkspaceSnapshots(ctx context.Context, params *DescribeWorkspaceSnapshotsInput, optFns ...func(*Options)) (*DescribeWorkspaceSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceSnapshots", params, optFns, c.addOperationDescribeWorkspaceSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceSnapshotsInput struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type DescribeWorkspaceSnapshotsOutput struct {

	// Information about the snapshots that can be used to rebuild a WorkSpace. These
	// snapshots include the user volume.
	RebuildSnapshots []types.Snapshot

	// Information about the snapshots that can be used to restore a WorkSpace. These
	// snapshots include both the root volume and the user volume.
	RestoreSnapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkspaceSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceSnapshots{}, middleware.After)
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
	if err = addOpDescribeWorkspaceSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceSnapshots(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorkspaceSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaceSnapshots",
	}
}
