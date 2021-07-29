// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the properties for one or more game session queues. When requesting
// multiple queues, use the pagination parameters to retrieve results as a set of
// sequential pages. If successful, a GameSessionQueue object is returned for each
// requested queue. When specifying a list of queues, objects are returned only for
// queues that currently exist in the Region. Learn more  View Your Queues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-console.html)
// Related actions CreateGameSessionQueue | DescribeGameSessionQueues |
// UpdateGameSessionQueue | DeleteGameSessionQueue | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DescribeGameSessionQueues(ctx context.Context, params *DescribeGameSessionQueuesInput, optFns ...func(*Options)) (*DescribeGameSessionQueuesOutput, error) {
	if params == nil {
		params = &DescribeGameSessionQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessionQueues", params, optFns, c.addOperationDescribeGameSessionQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeGameSessionQueuesInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. You can request up to 50 results.
	Limit *int32

	// A list of queue names to retrieve information for. You can use either the queue
	// ID or ARN value. To request settings for all queues, leave this parameter empty.
	Names []string

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type DescribeGameSessionQueuesOutput struct {

	// A collection of objects that describe the requested game session queues.
	GameSessionQueues []types.GameSessionQueue

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGameSessionQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessionQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessionQueues{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessionQueues(options.Region), middleware.Before); err != nil {
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

// DescribeGameSessionQueuesAPIClient is a client that implements the
// DescribeGameSessionQueues operation.
type DescribeGameSessionQueuesAPIClient interface {
	DescribeGameSessionQueues(context.Context, *DescribeGameSessionQueuesInput, ...func(*Options)) (*DescribeGameSessionQueuesOutput, error)
}

var _ DescribeGameSessionQueuesAPIClient = (*Client)(nil)

// DescribeGameSessionQueuesPaginatorOptions is the paginator options for
// DescribeGameSessionQueues
type DescribeGameSessionQueuesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. You can request up to 50 results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGameSessionQueuesPaginator is a paginator for DescribeGameSessionQueues
type DescribeGameSessionQueuesPaginator struct {
	options   DescribeGameSessionQueuesPaginatorOptions
	client    DescribeGameSessionQueuesAPIClient
	params    *DescribeGameSessionQueuesInput
	nextToken *string
	firstPage bool
}

// NewDescribeGameSessionQueuesPaginator returns a new
// DescribeGameSessionQueuesPaginator
func NewDescribeGameSessionQueuesPaginator(client DescribeGameSessionQueuesAPIClient, params *DescribeGameSessionQueuesInput, optFns ...func(*DescribeGameSessionQueuesPaginatorOptions)) *DescribeGameSessionQueuesPaginator {
	if params == nil {
		params = &DescribeGameSessionQueuesInput{}
	}

	options := DescribeGameSessionQueuesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeGameSessionQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGameSessionQueuesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeGameSessionQueues page.
func (p *DescribeGameSessionQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGameSessionQueuesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeGameSessionQueues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeGameSessionQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessionQueues",
	}
}
