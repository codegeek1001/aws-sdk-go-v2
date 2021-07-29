// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// View a list of configurations stored in the AppConfig configuration store by
// version.
func (c *Client) ListHostedConfigurationVersions(ctx context.Context, params *ListHostedConfigurationVersionsInput, optFns ...func(*Options)) (*ListHostedConfigurationVersionsOutput, error) {
	if params == nil {
		params = &ListHostedConfigurationVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostedConfigurationVersions", params, optFns, c.addOperationListHostedConfigurationVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostedConfigurationVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHostedConfigurationVersionsInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHostedConfigurationVersionsOutput struct {

	// The elements from this collection.
	Items []types.HostedConfigurationVersionSummary

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHostedConfigurationVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHostedConfigurationVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHostedConfigurationVersions{}, middleware.After)
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
	if err = addOpListHostedConfigurationVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHostedConfigurationVersions(options.Region), middleware.Before); err != nil {
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

// ListHostedConfigurationVersionsAPIClient is a client that implements the
// ListHostedConfigurationVersions operation.
type ListHostedConfigurationVersionsAPIClient interface {
	ListHostedConfigurationVersions(context.Context, *ListHostedConfigurationVersionsInput, ...func(*Options)) (*ListHostedConfigurationVersionsOutput, error)
}

var _ ListHostedConfigurationVersionsAPIClient = (*Client)(nil)

// ListHostedConfigurationVersionsPaginatorOptions is the paginator options for
// ListHostedConfigurationVersions
type ListHostedConfigurationVersionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHostedConfigurationVersionsPaginator is a paginator for
// ListHostedConfigurationVersions
type ListHostedConfigurationVersionsPaginator struct {
	options   ListHostedConfigurationVersionsPaginatorOptions
	client    ListHostedConfigurationVersionsAPIClient
	params    *ListHostedConfigurationVersionsInput
	nextToken *string
	firstPage bool
}

// NewListHostedConfigurationVersionsPaginator returns a new
// ListHostedConfigurationVersionsPaginator
func NewListHostedConfigurationVersionsPaginator(client ListHostedConfigurationVersionsAPIClient, params *ListHostedConfigurationVersionsInput, optFns ...func(*ListHostedConfigurationVersionsPaginatorOptions)) *ListHostedConfigurationVersionsPaginator {
	if params == nil {
		params = &ListHostedConfigurationVersionsInput{}
	}

	options := ListHostedConfigurationVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHostedConfigurationVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHostedConfigurationVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListHostedConfigurationVersions page.
func (p *ListHostedConfigurationVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHostedConfigurationVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListHostedConfigurationVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHostedConfigurationVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListHostedConfigurationVersions",
	}
}
