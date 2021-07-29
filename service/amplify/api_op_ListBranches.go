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

// Lists the branches of an Amplify app.
func (c *Client) ListBranches(ctx context.Context, params *ListBranchesInput, optFns ...func(*Options)) (*ListBranchesOutput, error) {
	if params == nil {
		params = &ListBranchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBranches", params, optFns, c.addOperationListBranchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list branches request.
type ListBranchesInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The maximum number of records to list in a single response.
	MaxResults int32

	// A pagination token. Set to null to start listing branches from the start. If a
	// non-null pagination token is returned in a result, pass its value in here to
	// list more branches.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list branches request.
type ListBranchesOutput struct {

	// A list of branches for an Amplify app.
	//
	// This member is required.
	Branches []types.Branch

	// A pagination token. If a non-null pagination token is returned in a result, pass
	// its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBranchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBranches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBranches{}, middleware.After)
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
	if err = addOpListBranchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBranches(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListBranches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListBranches",
	}
}
