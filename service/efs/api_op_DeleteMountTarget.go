// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified mount target. This operation forcibly breaks any mounts of
// the file system by using the mount target that is being deleted, which might
// disrupt instances or applications using those mounts. To avoid applications
// getting cut off abruptly, you might consider unmounting any mounts of the mount
// target, if feasible. The operation also deletes the associated network
// interface. Uncommitted writes might be lost, but breaking a mount target using
// this operation does not corrupt the file system itself. The file system you
// created remains. You can mount an EC2 instance in your VPC by using another
// mount target. This operation requires permissions for the following action on
// the file system:
//
// * elasticfilesystem:DeleteMountTarget
//
// The DeleteMountTarget
// call returns while the mount target state is still deleting. You can check the
// mount target deletion by calling the DescribeMountTargets operation, which
// returns a list of mount target descriptions for the given file system. The
// operation also requires permissions for the following Amazon EC2 action on the
// mount target's network interface:
//
// * ec2:DeleteNetworkInterface
func (c *Client) DeleteMountTarget(ctx context.Context, params *DeleteMountTargetInput, optFns ...func(*Options)) (*DeleteMountTargetOutput, error) {
	if params == nil {
		params = &DeleteMountTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMountTarget", params, optFns, c.addOperationDeleteMountTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMountTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteMountTargetInput struct {

	// The ID of the mount target to delete (String).
	//
	// This member is required.
	MountTargetId *string

	noSmithyDocumentSerde
}

type DeleteMountTargetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMountTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMountTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMountTarget{}, middleware.After)
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
	if err = addOpDeleteMountTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMountTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMountTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DeleteMountTarget",
	}
}
