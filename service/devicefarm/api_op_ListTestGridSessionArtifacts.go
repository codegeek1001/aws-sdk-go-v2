// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of artifacts created during the session.
func (c *Client) ListTestGridSessionArtifacts(ctx context.Context, params *ListTestGridSessionArtifactsInput, optFns ...func(*Options)) (*ListTestGridSessionArtifactsOutput, error) {
	if params == nil {
		params = &ListTestGridSessionArtifactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestGridSessionArtifacts", params, optFns, c.addOperationListTestGridSessionArtifactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestGridSessionArtifactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestGridSessionArtifactsInput struct {

	// The ARN of a TestGridSession.
	//
	// This member is required.
	SessionArn *string

	// The maximum number of results to be returned by a request.
	MaxResult *int32

	// Pagination token.
	NextToken *string

	// Limit results to a specified type of artifact.
	Type types.TestGridSessionArtifactCategory

	noSmithyDocumentSerde
}

type ListTestGridSessionArtifactsOutput struct {

	// A list of test grid session artifacts for a TestGridSession.
	Artifacts []types.TestGridSessionArtifact

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestGridSessionArtifactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTestGridSessionArtifacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTestGridSessionArtifacts{}, middleware.After)
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
	if err = addOpListTestGridSessionArtifactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestGridSessionArtifacts(options.Region), middleware.Before); err != nil {
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

// ListTestGridSessionArtifactsAPIClient is a client that implements the
// ListTestGridSessionArtifacts operation.
type ListTestGridSessionArtifactsAPIClient interface {
	ListTestGridSessionArtifacts(context.Context, *ListTestGridSessionArtifactsInput, ...func(*Options)) (*ListTestGridSessionArtifactsOutput, error)
}

var _ ListTestGridSessionArtifactsAPIClient = (*Client)(nil)

// ListTestGridSessionArtifactsPaginatorOptions is the paginator options for
// ListTestGridSessionArtifacts
type ListTestGridSessionArtifactsPaginatorOptions struct {
	// The maximum number of results to be returned by a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestGridSessionArtifactsPaginator is a paginator for
// ListTestGridSessionArtifacts
type ListTestGridSessionArtifactsPaginator struct {
	options   ListTestGridSessionArtifactsPaginatorOptions
	client    ListTestGridSessionArtifactsAPIClient
	params    *ListTestGridSessionArtifactsInput
	nextToken *string
	firstPage bool
}

// NewListTestGridSessionArtifactsPaginator returns a new
// ListTestGridSessionArtifactsPaginator
func NewListTestGridSessionArtifactsPaginator(client ListTestGridSessionArtifactsAPIClient, params *ListTestGridSessionArtifactsInput, optFns ...func(*ListTestGridSessionArtifactsPaginatorOptions)) *ListTestGridSessionArtifactsPaginator {
	if params == nil {
		params = &ListTestGridSessionArtifactsInput{}
	}

	options := ListTestGridSessionArtifactsPaginatorOptions{}
	if params.MaxResult != nil {
		options.Limit = *params.MaxResult
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestGridSessionArtifactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestGridSessionArtifactsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTestGridSessionArtifacts page.
func (p *ListTestGridSessionArtifactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestGridSessionArtifactsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResult = limit

	result, err := p.client.ListTestGridSessionArtifacts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestGridSessionArtifacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListTestGridSessionArtifacts",
	}
}
