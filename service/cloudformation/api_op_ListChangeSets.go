// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the ID and status of each active change set for a stack. For example,
// AWS CloudFormation lists change sets that are in the CREATE_IN_PROGRESS or
// CREATE_PENDING state.
func (c *Client) ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
	if params == nil {
		params = &ListChangeSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChangeSets", params, optFns, c.addOperationListChangeSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChangeSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListChangeSets action.
type ListChangeSetsInput struct {

	// The name or the Amazon Resource Name (ARN) of the stack for which you want to
	// list change sets.
	//
	// This member is required.
	StackName *string

	// A string (provided by the ListChangeSets response output) that identifies the
	// next page of change sets that you want to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

// The output for the ListChangeSets action.
type ListChangeSetsOutput struct {

	// If the output exceeds 1 MB, a string that identifies the next page of change
	// sets. If there is no additional page, this value is null.
	NextToken *string

	// A list of ChangeSetSummary structures that provides the ID and status of each
	// change set for the specified stack.
	Summaries []types.ChangeSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChangeSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListChangeSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListChangeSets{}, middleware.After)
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
	if err = addOpListChangeSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChangeSets(options.Region), middleware.Before); err != nil {
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

// ListChangeSetsAPIClient is a client that implements the ListChangeSets
// operation.
type ListChangeSetsAPIClient interface {
	ListChangeSets(context.Context, *ListChangeSetsInput, ...func(*Options)) (*ListChangeSetsOutput, error)
}

var _ ListChangeSetsAPIClient = (*Client)(nil)

// ListChangeSetsPaginatorOptions is the paginator options for ListChangeSets
type ListChangeSetsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChangeSetsPaginator is a paginator for ListChangeSets
type ListChangeSetsPaginator struct {
	options   ListChangeSetsPaginatorOptions
	client    ListChangeSetsAPIClient
	params    *ListChangeSetsInput
	nextToken *string
	firstPage bool
}

// NewListChangeSetsPaginator returns a new ListChangeSetsPaginator
func NewListChangeSetsPaginator(client ListChangeSetsAPIClient, params *ListChangeSetsInput, optFns ...func(*ListChangeSetsPaginatorOptions)) *ListChangeSetsPaginator {
	if params == nil {
		params = &ListChangeSetsInput{}
	}

	options := ListChangeSetsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChangeSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChangeSetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListChangeSets page.
func (p *ListChangeSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChangeSetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListChangeSets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListChangeSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListChangeSets",
	}
}
