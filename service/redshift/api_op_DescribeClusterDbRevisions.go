// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of ClusterDbRevision objects.
func (c *Client) DescribeClusterDbRevisions(ctx context.Context, params *DescribeClusterDbRevisionsInput, optFns ...func(*Options)) (*DescribeClusterDbRevisionsOutput, error) {
	if params == nil {
		params = &DescribeClusterDbRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterDbRevisions", params, optFns, c.addOperationDescribeClusterDbRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterDbRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterDbRevisionsInput struct {

	// A unique identifier for a cluster whose ClusterDbRevisions you are requesting.
	// This parameter is case sensitive. All clusters defined for an account are
	// returned by default.
	ClusterIdentifier *string

	// An optional parameter that specifies the starting point for returning a set of
	// response records. When the results of a DescribeClusterDbRevisions request
	// exceed the value specified in MaxRecords, Amazon Redshift returns a value in the
	// marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the marker parameter and retrying the
	// request. Constraints: You can specify either the ClusterIdentifier parameter, or
	// the marker parameter, but not both.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in the marker field of the response. You can retrieve the next set of
	// response records by providing the returned marker value in the marker parameter
	// and retrying the request. Default: 100 Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeClusterDbRevisionsOutput struct {

	// A list of revisions.
	ClusterDbRevisions []types.ClusterDbRevision

	// A string representing the starting point for the next set of revisions. If a
	// value is returned in a response, you can retrieve the next set of revisions by
	// providing the value in the marker parameter and retrying the command. If the
	// marker field is empty, all revisions have already been returned.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterDbRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterDbRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterDbRevisions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterDbRevisions(options.Region), middleware.Before); err != nil {
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

// DescribeClusterDbRevisionsAPIClient is a client that implements the
// DescribeClusterDbRevisions operation.
type DescribeClusterDbRevisionsAPIClient interface {
	DescribeClusterDbRevisions(context.Context, *DescribeClusterDbRevisionsInput, ...func(*Options)) (*DescribeClusterDbRevisionsOutput, error)
}

var _ DescribeClusterDbRevisionsAPIClient = (*Client)(nil)

// DescribeClusterDbRevisionsPaginatorOptions is the paginator options for
// DescribeClusterDbRevisions
type DescribeClusterDbRevisionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in the marker field of the response. You can retrieve the next set of
	// response records by providing the returned marker value in the marker parameter
	// and retrying the request. Default: 100 Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterDbRevisionsPaginator is a paginator for
// DescribeClusterDbRevisions
type DescribeClusterDbRevisionsPaginator struct {
	options   DescribeClusterDbRevisionsPaginatorOptions
	client    DescribeClusterDbRevisionsAPIClient
	params    *DescribeClusterDbRevisionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterDbRevisionsPaginator returns a new
// DescribeClusterDbRevisionsPaginator
func NewDescribeClusterDbRevisionsPaginator(client DescribeClusterDbRevisionsAPIClient, params *DescribeClusterDbRevisionsInput, optFns ...func(*DescribeClusterDbRevisionsPaginatorOptions)) *DescribeClusterDbRevisionsPaginator {
	if params == nil {
		params = &DescribeClusterDbRevisionsInput{}
	}

	options := DescribeClusterDbRevisionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterDbRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterDbRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeClusterDbRevisions page.
func (p *DescribeClusterDbRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterDbRevisionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeClusterDbRevisions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeClusterDbRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterDbRevisions",
	}
}
