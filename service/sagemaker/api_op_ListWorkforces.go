// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to list all private and vendor workforces in an Amazon Web
// Services Region. Note that you can only have one private workforce per Amazon
// Web Services Region.
func (c *Client) ListWorkforces(ctx context.Context, params *ListWorkforcesInput, optFns ...func(*Options)) (*ListWorkforcesOutput, error) {
	if params == nil {
		params = &ListWorkforcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkforces", params, optFns, c.addOperationListWorkforcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkforcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkforcesInput struct {

	// The maximum number of workforces returned in the response.
	MaxResults *int32

	// A filter you can use to search for workforces using part of the workforce name.
	NameContains *string

	// A token to resume pagination.
	NextToken *string

	// Sort workforces using the workforce name or creation date.
	SortBy types.ListWorkforcesSortByOptions

	// Sort workforces in ascending or descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListWorkforcesOutput struct {

	// A list containing information about your workforce.
	//
	// This member is required.
	Workforces []types.Workforce

	// A token to resume pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkforcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkforces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkforces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkforces(options.Region), middleware.Before); err != nil {
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

// ListWorkforcesAPIClient is a client that implements the ListWorkforces
// operation.
type ListWorkforcesAPIClient interface {
	ListWorkforces(context.Context, *ListWorkforcesInput, ...func(*Options)) (*ListWorkforcesOutput, error)
}

var _ ListWorkforcesAPIClient = (*Client)(nil)

// ListWorkforcesPaginatorOptions is the paginator options for ListWorkforces
type ListWorkforcesPaginatorOptions struct {
	// The maximum number of workforces returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkforcesPaginator is a paginator for ListWorkforces
type ListWorkforcesPaginator struct {
	options   ListWorkforcesPaginatorOptions
	client    ListWorkforcesAPIClient
	params    *ListWorkforcesInput
	nextToken *string
	firstPage bool
}

// NewListWorkforcesPaginator returns a new ListWorkforcesPaginator
func NewListWorkforcesPaginator(client ListWorkforcesAPIClient, params *ListWorkforcesInput, optFns ...func(*ListWorkforcesPaginatorOptions)) *ListWorkforcesPaginator {
	if params == nil {
		params = &ListWorkforcesInput{}
	}

	options := ListWorkforcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkforcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkforcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListWorkforces page.
func (p *ListWorkforcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkforcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListWorkforces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkforces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListWorkforces",
	}
}
