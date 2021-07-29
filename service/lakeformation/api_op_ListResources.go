// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resources registered to be managed by the Data Catalog.
func (c *Client) ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if params == nil {
		params = &ListResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResources", params, optFns, c.addOperationListResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInput struct {

	// Any applicable row-level and/or column-level filtering conditions for the
	// resources.
	FilterConditionList []types.FilterCondition

	// The maximum number of resource results.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve these resources.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourcesOutput struct {

	// A continuation token, if this is not the first call to retrieve these resources.
	NextToken *string

	// A summary of the data lake resources.
	ResourceInfoList []types.ResourceInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResources{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResources(options.Region), middleware.Before); err != nil {
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

// ListResourcesAPIClient is a client that implements the ListResources operation.
type ListResourcesAPIClient interface {
	ListResources(context.Context, *ListResourcesInput, ...func(*Options)) (*ListResourcesOutput, error)
}

var _ ListResourcesAPIClient = (*Client)(nil)

// ListResourcesPaginatorOptions is the paginator options for ListResources
type ListResourcesPaginatorOptions struct {
	// The maximum number of resource results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesPaginator is a paginator for ListResources
type ListResourcesPaginator struct {
	options   ListResourcesPaginatorOptions
	client    ListResourcesAPIClient
	params    *ListResourcesInput
	nextToken *string
	firstPage bool
}

// NewListResourcesPaginator returns a new ListResourcesPaginator
func NewListResourcesPaginator(client ListResourcesAPIClient, params *ListResourcesInput, optFns ...func(*ListResourcesPaginatorOptions)) *ListResourcesPaginator {
	if params == nil {
		params = &ListResourcesInput{}
	}

	options := ListResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListResources page.
func (p *ListResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesOutput, error) {
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

	result, err := p.client.ListResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "ListResources",
	}
}
