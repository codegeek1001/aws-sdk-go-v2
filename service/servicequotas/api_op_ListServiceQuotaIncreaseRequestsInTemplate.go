// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the quota increase requests in the specified quota request template.
func (c *Client) ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, params *ListServiceQuotaIncreaseRequestsInTemplateInput, optFns ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	if params == nil {
		params = &ListServiceQuotaIncreaseRequestsInTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceQuotaIncreaseRequestsInTemplate", params, optFns, c.addOperationListServiceQuotaIncreaseRequestsInTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceQuotaIncreaseRequestsInTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceQuotaIncreaseRequestsInTemplateInput struct {

	// The AWS Region.
	AwsRegion *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The service identifier.
	ServiceCode *string

	noSmithyDocumentSerde
}

type ListServiceQuotaIncreaseRequestsInTemplateOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the quota increase requests.
	ServiceQuotaIncreaseRequestInTemplateList []types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceQuotaIncreaseRequestsInTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(options.Region), middleware.Before); err != nil {
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

// ListServiceQuotaIncreaseRequestsInTemplateAPIClient is a client that implements
// the ListServiceQuotaIncreaseRequestsInTemplate operation.
type ListServiceQuotaIncreaseRequestsInTemplateAPIClient interface {
	ListServiceQuotaIncreaseRequestsInTemplate(context.Context, *ListServiceQuotaIncreaseRequestsInTemplateInput, ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
}

var _ ListServiceQuotaIncreaseRequestsInTemplateAPIClient = (*Client)(nil)

// ListServiceQuotaIncreaseRequestsInTemplatePaginatorOptions is the paginator
// options for ListServiceQuotaIncreaseRequestsInTemplate
type ListServiceQuotaIncreaseRequestsInTemplatePaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceQuotaIncreaseRequestsInTemplatePaginator is a paginator for
// ListServiceQuotaIncreaseRequestsInTemplate
type ListServiceQuotaIncreaseRequestsInTemplatePaginator struct {
	options   ListServiceQuotaIncreaseRequestsInTemplatePaginatorOptions
	client    ListServiceQuotaIncreaseRequestsInTemplateAPIClient
	params    *ListServiceQuotaIncreaseRequestsInTemplateInput
	nextToken *string
	firstPage bool
}

// NewListServiceQuotaIncreaseRequestsInTemplatePaginator returns a new
// ListServiceQuotaIncreaseRequestsInTemplatePaginator
func NewListServiceQuotaIncreaseRequestsInTemplatePaginator(client ListServiceQuotaIncreaseRequestsInTemplateAPIClient, params *ListServiceQuotaIncreaseRequestsInTemplateInput, optFns ...func(*ListServiceQuotaIncreaseRequestsInTemplatePaginatorOptions)) *ListServiceQuotaIncreaseRequestsInTemplatePaginator {
	if params == nil {
		params = &ListServiceQuotaIncreaseRequestsInTemplateInput{}
	}

	options := ListServiceQuotaIncreaseRequestsInTemplatePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceQuotaIncreaseRequestsInTemplatePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceQuotaIncreaseRequestsInTemplatePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListServiceQuotaIncreaseRequestsInTemplate page.
func (p *ListServiceQuotaIncreaseRequestsInTemplatePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
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

	result, err := p.client.ListServiceQuotaIncreaseRequestsInTemplate(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListServiceQuotaIncreaseRequestsInTemplate",
	}
}
