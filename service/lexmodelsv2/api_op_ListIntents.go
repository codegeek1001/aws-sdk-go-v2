// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of intents that meet the specified criteria.
func (c *Client) ListIntents(ctx context.Context, params *ListIntentsInput, optFns ...func(*Options)) (*ListIntentsOutput, error) {
	if params == nil {
		params = &ListIntentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIntents", params, optFns, c.addOperationListIntentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIntentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIntentsInput struct {

	// The unique identifier of the bot that contains the intent.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the intent.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the intents to list. The string
	// must match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// Provides the specification of a filter used to limit the intents in the response
	// to only those that match the filter specification. You can only specify one
	// filter and only one string to filter on.
	Filters []types.IntentFilter

	// The maximum number of intents to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListIntents operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token in the nextToken parameter to return the next page of results.
	NextToken *string

	// Determines the sort order for the response from the ListIntents operation. You
	// can choose to sort by the intent name or last updated date in either ascending
	// or descending order.
	SortBy *types.IntentSortBy

	noSmithyDocumentSerde
}

type ListIntentsOutput struct {

	// The identifier of the bot that contains the intent.
	BotId *string

	// The version of the bot that contains the intent.
	BotVersion *string

	// Summary information for the intents that meet the filter criteria specified in
	// the request. The length of the list is specified in the maxResults parameter of
	// the request. If there are more intents available, the nextToken field contains a
	// token to get the next page of results.
	IntentSummaries []types.IntentSummary

	// The language and locale of the intents in the list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response to
	// the ListIntents operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListIntents operation request to get
	// the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIntentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIntents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIntents{}, middleware.After)
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
	if err = addOpListIntentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIntents(options.Region), middleware.Before); err != nil {
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

// ListIntentsAPIClient is a client that implements the ListIntents operation.
type ListIntentsAPIClient interface {
	ListIntents(context.Context, *ListIntentsInput, ...func(*Options)) (*ListIntentsOutput, error)
}

var _ ListIntentsAPIClient = (*Client)(nil)

// ListIntentsPaginatorOptions is the paginator options for ListIntents
type ListIntentsPaginatorOptions struct {
	// The maximum number of intents to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIntentsPaginator is a paginator for ListIntents
type ListIntentsPaginator struct {
	options   ListIntentsPaginatorOptions
	client    ListIntentsAPIClient
	params    *ListIntentsInput
	nextToken *string
	firstPage bool
}

// NewListIntentsPaginator returns a new ListIntentsPaginator
func NewListIntentsPaginator(client ListIntentsAPIClient, params *ListIntentsInput, optFns ...func(*ListIntentsPaginatorOptions)) *ListIntentsPaginator {
	if params == nil {
		params = &ListIntentsInput{}
	}

	options := ListIntentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIntentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIntentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListIntents page.
func (p *ListIntentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIntentsOutput, error) {
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

	result, err := p.client.ListIntents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIntents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListIntents",
	}
}
