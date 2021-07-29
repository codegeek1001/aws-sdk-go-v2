// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Merges two branches using the fast-forward merge strategy.
func (c *Client) MergeBranchesByFastForward(ctx context.Context, params *MergeBranchesByFastForwardInput, optFns ...func(*Options)) (*MergeBranchesByFastForwardOutput, error) {
	if params == nil {
		params = &MergeBranchesByFastForwardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergeBranchesByFastForward", params, optFns, c.addOperationMergeBranchesByFastForwardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergeBranchesByFastForwardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MergeBranchesByFastForwardInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The name of the repository where you want to merge two branches.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The branch where the merge is applied.
	TargetBranch *string

	noSmithyDocumentSerde
}

type MergeBranchesByFastForwardOutput struct {

	// The commit ID of the merge in the destination or target branch.
	CommitId *string

	// The tree ID of the merge in the destination or target branch.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMergeBranchesByFastForwardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergeBranchesByFastForward{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeBranchesByFastForward{}, middleware.After)
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
	if err = addOpMergeBranchesByFastForwardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergeBranchesByFastForward(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergeBranchesByFastForward(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "MergeBranchesByFastForward",
	}
}
