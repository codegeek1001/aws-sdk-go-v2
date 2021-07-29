// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes an asset, such as a dashboard, analysis, or dataset, from a folder.
func (c *Client) DeleteFolderMembership(ctx context.Context, params *DeleteFolderMembershipInput, optFns ...func(*Options)) (*DeleteFolderMembershipOutput, error) {
	if params == nil {
		params = &DeleteFolderMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFolderMembership", params, optFns, c.addOperationDeleteFolderMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFolderMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFolderMembershipInput struct {

	// The AWS Account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The Folder ID.
	//
	// This member is required.
	FolderId *string

	// The ID of the asset (the dashboard, analysis, or dataset) that you want to
	// delete.
	//
	// This member is required.
	MemberId *string

	// The type of the member, including DASHBOARD, ANALYSIS, and DATASET
	//
	// This member is required.
	MemberType types.MemberType

	noSmithyDocumentSerde
}

type DeleteFolderMembershipOutput struct {

	// The request ID.
	RequestId *string

	// The status of deleting the asset. If succeeded, the status is SC_OK (200).
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFolderMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFolderMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFolderMembership{}, middleware.After)
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
	if err = addOpDeleteFolderMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFolderMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFolderMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteFolderMembership",
	}
}
