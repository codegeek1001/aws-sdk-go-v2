// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about recent sampling results for all sampling rules.
func (c *Client) GetSamplingStatisticSummaries(ctx context.Context, params *GetSamplingStatisticSummariesInput, optFns ...func(*Options)) (*GetSamplingStatisticSummariesOutput, error) {
	if params == nil {
		params = &GetSamplingStatisticSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSamplingStatisticSummaries", params, optFns, c.addOperationGetSamplingStatisticSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSamplingStatisticSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSamplingStatisticSummariesInput struct {

	// Pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSamplingStatisticSummariesOutput struct {

	// Pagination token.
	NextToken *string

	// Information about the number of requests instrumented for each sampling rule.
	SamplingStatisticSummaries []types.SamplingStatisticSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSamplingStatisticSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSamplingStatisticSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSamplingStatisticSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSamplingStatisticSummaries(options.Region), middleware.Before); err != nil {
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

// GetSamplingStatisticSummariesAPIClient is a client that implements the
// GetSamplingStatisticSummaries operation.
type GetSamplingStatisticSummariesAPIClient interface {
	GetSamplingStatisticSummaries(context.Context, *GetSamplingStatisticSummariesInput, ...func(*Options)) (*GetSamplingStatisticSummariesOutput, error)
}

var _ GetSamplingStatisticSummariesAPIClient = (*Client)(nil)

// GetSamplingStatisticSummariesPaginatorOptions is the paginator options for
// GetSamplingStatisticSummaries
type GetSamplingStatisticSummariesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSamplingStatisticSummariesPaginator is a paginator for
// GetSamplingStatisticSummaries
type GetSamplingStatisticSummariesPaginator struct {
	options   GetSamplingStatisticSummariesPaginatorOptions
	client    GetSamplingStatisticSummariesAPIClient
	params    *GetSamplingStatisticSummariesInput
	nextToken *string
	firstPage bool
}

// NewGetSamplingStatisticSummariesPaginator returns a new
// GetSamplingStatisticSummariesPaginator
func NewGetSamplingStatisticSummariesPaginator(client GetSamplingStatisticSummariesAPIClient, params *GetSamplingStatisticSummariesInput, optFns ...func(*GetSamplingStatisticSummariesPaginatorOptions)) *GetSamplingStatisticSummariesPaginator {
	if params == nil {
		params = &GetSamplingStatisticSummariesInput{}
	}

	options := GetSamplingStatisticSummariesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSamplingStatisticSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSamplingStatisticSummariesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetSamplingStatisticSummaries page.
func (p *GetSamplingStatisticSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSamplingStatisticSummariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetSamplingStatisticSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSamplingStatisticSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetSamplingStatisticSummaries",
	}
}
