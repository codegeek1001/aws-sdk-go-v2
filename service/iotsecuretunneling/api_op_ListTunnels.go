// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsecuretunneling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all tunnels for an AWS account. Tunnels are listed by creation time in
// descending order, newer tunnels will be listed before older tunnels.
func (c *Client) ListTunnels(ctx context.Context, params *ListTunnelsInput, optFns ...func(*Options)) (*ListTunnelsOutput, error) {
	if params == nil {
		params = &ListTunnelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTunnels", params, optFns, c.addOperationListTunnelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTunnelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTunnelsInput struct {

	// The maximum number of results to return at once.
	MaxResults int32

	// A token to retrieve the next set of results.
	NextToken *string

	// The name of the IoT thing associated with the destination device.
	ThingName *string

	noSmithyDocumentSerde
}

type ListTunnelsOutput struct {

	// A token to used to retrieve the next set of results.
	NextToken *string

	// A short description of the tunnels in an AWS account.
	TunnelSummaries []types.TunnelSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTunnelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTunnels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTunnels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTunnels(options.Region), middleware.Before); err != nil {
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

// ListTunnelsAPIClient is a client that implements the ListTunnels operation.
type ListTunnelsAPIClient interface {
	ListTunnels(context.Context, *ListTunnelsInput, ...func(*Options)) (*ListTunnelsOutput, error)
}

var _ ListTunnelsAPIClient = (*Client)(nil)

// ListTunnelsPaginatorOptions is the paginator options for ListTunnels
type ListTunnelsPaginatorOptions struct {
	// The maximum number of results to return at once.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTunnelsPaginator is a paginator for ListTunnels
type ListTunnelsPaginator struct {
	options   ListTunnelsPaginatorOptions
	client    ListTunnelsAPIClient
	params    *ListTunnelsInput
	nextToken *string
	firstPage bool
}

// NewListTunnelsPaginator returns a new ListTunnelsPaginator
func NewListTunnelsPaginator(client ListTunnelsAPIClient, params *ListTunnelsInput, optFns ...func(*ListTunnelsPaginatorOptions)) *ListTunnelsPaginator {
	if params == nil {
		params = &ListTunnelsInput{}
	}

	options := ListTunnelsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTunnelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTunnelsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTunnels page.
func (p *ListTunnelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTunnelsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListTunnels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTunnels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsecuredtunneling",
		OperationName: "ListTunnels",
	}
}
