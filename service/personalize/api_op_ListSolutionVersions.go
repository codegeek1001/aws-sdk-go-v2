// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of solution versions for the given solution. When a solution is
// not specified, all the solution versions associated with the account are listed.
// The response provides the properties for each solution version, including the
// Amazon Resource Name (ARN). For more information on solutions, see
// CreateSolution.
func (c *Client) ListSolutionVersions(ctx context.Context, params *ListSolutionVersionsInput, optFns ...func(*Options)) (*ListSolutionVersionsOutput, error) {
	if params == nil {
		params = &ListSolutionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolutionVersions", params, optFns, c.addOperationListSolutionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolutionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolutionVersionsInput struct {

	// The maximum number of solution versions to return.
	MaxResults *int32

	// A token returned from the previous call to ListSolutionVersions for getting the
	// next set of solution versions (if they exist).
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution.
	SolutionArn *string

	noSmithyDocumentSerde
}

type ListSolutionVersionsOutput struct {

	// A token for getting the next set of solution versions (if they exist).
	NextToken *string

	// A list of solution versions describing the version properties.
	SolutionVersions []types.SolutionVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolutionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSolutionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSolutionVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolutionVersions(options.Region), middleware.Before); err != nil {
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

// ListSolutionVersionsAPIClient is a client that implements the
// ListSolutionVersions operation.
type ListSolutionVersionsAPIClient interface {
	ListSolutionVersions(context.Context, *ListSolutionVersionsInput, ...func(*Options)) (*ListSolutionVersionsOutput, error)
}

var _ ListSolutionVersionsAPIClient = (*Client)(nil)

// ListSolutionVersionsPaginatorOptions is the paginator options for
// ListSolutionVersions
type ListSolutionVersionsPaginatorOptions struct {
	// The maximum number of solution versions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolutionVersionsPaginator is a paginator for ListSolutionVersions
type ListSolutionVersionsPaginator struct {
	options   ListSolutionVersionsPaginatorOptions
	client    ListSolutionVersionsAPIClient
	params    *ListSolutionVersionsInput
	nextToken *string
	firstPage bool
}

// NewListSolutionVersionsPaginator returns a new ListSolutionVersionsPaginator
func NewListSolutionVersionsPaginator(client ListSolutionVersionsAPIClient, params *ListSolutionVersionsInput, optFns ...func(*ListSolutionVersionsPaginatorOptions)) *ListSolutionVersionsPaginator {
	if params == nil {
		params = &ListSolutionVersionsInput{}
	}

	options := ListSolutionVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolutionVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolutionVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSolutionVersions page.
func (p *ListSolutionVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolutionVersionsOutput, error) {
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

	result, err := p.client.ListSolutionVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSolutionVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListSolutionVersions",
	}
}
