// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the playback configurations defined in AWS Elemental
// MediaTailor. You can specify a maximum number of configurations to return at a
// time. The default maximum is 50. Results are returned in pagefuls. If
// MediaTailor has more configurations than the specified maximum, it provides
// parameters in the response that you can use to retrieve the next pageful.
func (c *Client) ListPlaybackConfigurations(ctx context.Context, params *ListPlaybackConfigurationsInput, optFns ...func(*Options)) (*ListPlaybackConfigurationsOutput, error) {
	if params == nil {
		params = &ListPlaybackConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlaybackConfigurations", params, optFns, c.addOperationListPlaybackConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPlaybackConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlaybackConfigurationsInput struct {

	// Maximum number of records to return.
	MaxResults int32

	// Pagination token returned by the GET list request when results exceed the
	// maximum allowed. Use the token to fetch the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPlaybackConfigurationsOutput struct {

	// Array of playback configurations. This might be all the available configurations
	// or a subset, depending on the settings that you provide and the total number of
	// configurations stored.
	Items []types.PlaybackConfiguration

	// Pagination token returned by the GET list request when results exceed the
	// maximum allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlaybackConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlaybackConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlaybackConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlaybackConfigurations(options.Region), middleware.Before); err != nil {
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

// ListPlaybackConfigurationsAPIClient is a client that implements the
// ListPlaybackConfigurations operation.
type ListPlaybackConfigurationsAPIClient interface {
	ListPlaybackConfigurations(context.Context, *ListPlaybackConfigurationsInput, ...func(*Options)) (*ListPlaybackConfigurationsOutput, error)
}

var _ ListPlaybackConfigurationsAPIClient = (*Client)(nil)

// ListPlaybackConfigurationsPaginatorOptions is the paginator options for
// ListPlaybackConfigurations
type ListPlaybackConfigurationsPaginatorOptions struct {
	// Maximum number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlaybackConfigurationsPaginator is a paginator for
// ListPlaybackConfigurations
type ListPlaybackConfigurationsPaginator struct {
	options   ListPlaybackConfigurationsPaginatorOptions
	client    ListPlaybackConfigurationsAPIClient
	params    *ListPlaybackConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListPlaybackConfigurationsPaginator returns a new
// ListPlaybackConfigurationsPaginator
func NewListPlaybackConfigurationsPaginator(client ListPlaybackConfigurationsAPIClient, params *ListPlaybackConfigurationsInput, optFns ...func(*ListPlaybackConfigurationsPaginatorOptions)) *ListPlaybackConfigurationsPaginator {
	if params == nil {
		params = &ListPlaybackConfigurationsInput{}
	}

	options := ListPlaybackConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPlaybackConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlaybackConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPlaybackConfigurations page.
func (p *ListPlaybackConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlaybackConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPlaybackConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlaybackConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListPlaybackConfigurations",
	}
}
