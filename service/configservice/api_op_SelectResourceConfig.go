// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts a structured query language (SQL) SELECT command, performs the
// corresponding search, and returns resource configurations matching the
// properties. For more information about query components, see the  Query
// Components
// (https://docs.aws.amazon.com/config/latest/developerguide/query-components.html)
// section in the AWS Config Developer Guide.
func (c *Client) SelectResourceConfig(ctx context.Context, params *SelectResourceConfigInput, optFns ...func(*Options)) (*SelectResourceConfigOutput, error) {
	if params == nil {
		params = &SelectResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SelectResourceConfig", params, optFns, c.addOperationSelectResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SelectResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SelectResourceConfigInput struct {

	// The SQL query SELECT command.
	//
	// This member is required.
	Expression *string

	// The maximum number of query results returned on each page.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type SelectResourceConfigOutput struct {

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Returns the QueryInfo object.
	QueryInfo *types.QueryInfo

	// Returns the results for the SQL query.
	Results []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSelectResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSelectResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSelectResourceConfig{}, middleware.After)
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
	if err = addOpSelectResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSelectResourceConfig(options.Region), middleware.Before); err != nil {
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

// SelectResourceConfigAPIClient is a client that implements the
// SelectResourceConfig operation.
type SelectResourceConfigAPIClient interface {
	SelectResourceConfig(context.Context, *SelectResourceConfigInput, ...func(*Options)) (*SelectResourceConfigOutput, error)
}

var _ SelectResourceConfigAPIClient = (*Client)(nil)

// SelectResourceConfigPaginatorOptions is the paginator options for
// SelectResourceConfig
type SelectResourceConfigPaginatorOptions struct {
	// The maximum number of query results returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SelectResourceConfigPaginator is a paginator for SelectResourceConfig
type SelectResourceConfigPaginator struct {
	options   SelectResourceConfigPaginatorOptions
	client    SelectResourceConfigAPIClient
	params    *SelectResourceConfigInput
	nextToken *string
	firstPage bool
}

// NewSelectResourceConfigPaginator returns a new SelectResourceConfigPaginator
func NewSelectResourceConfigPaginator(client SelectResourceConfigAPIClient, params *SelectResourceConfigInput, optFns ...func(*SelectResourceConfigPaginatorOptions)) *SelectResourceConfigPaginator {
	if params == nil {
		params = &SelectResourceConfigInput{}
	}

	options := SelectResourceConfigPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SelectResourceConfigPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SelectResourceConfigPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SelectResourceConfig page.
func (p *SelectResourceConfigPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SelectResourceConfigOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.SelectResourceConfig(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSelectResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "SelectResourceConfig",
	}
}
