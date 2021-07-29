// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of PolicyComplianceStatus objects. Use PolicyComplianceStatus
// to get a summary of which member accounts are protected by the specified policy.
func (c *Client) ListComplianceStatus(ctx context.Context, params *ListComplianceStatusInput, optFns ...func(*Options)) (*ListComplianceStatusOutput, error) {
	if params == nil {
		params = &ListComplianceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComplianceStatus", params, optFns, c.addOperationListComplianceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComplianceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComplianceStatusInput struct {

	// The ID of the Firewall Manager policy that you want the details for.
	//
	// This member is required.
	PolicyId *string

	// Specifies the number of PolicyComplianceStatus objects that you want Firewall
	// Manager to return for this request. If you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, the response includes a
	// NextToken value that you can use to get another batch of PolicyComplianceStatus
	// objects.
	MaxResults *int32

	// If you specify a value for MaxResults and you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, Firewall Manager
	// returns a NextToken value in the response that allows you to list another group
	// of PolicyComplianceStatus objects. For the second and subsequent
	// ListComplianceStatus requests, specify the value of NextToken from the previous
	// response to get information about another batch of PolicyComplianceStatus
	// objects.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComplianceStatusOutput struct {

	// If you have more PolicyComplianceStatus objects than the number that you
	// specified for MaxResults in the request, the response includes a NextToken
	// value. To list more PolicyComplianceStatus objects, submit another
	// ListComplianceStatus request, and specify the NextToken value from the response
	// in the NextToken value in the next request.
	NextToken *string

	// An array of PolicyComplianceStatus objects.
	PolicyComplianceStatusList []types.PolicyComplianceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComplianceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListComplianceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListComplianceStatus{}, middleware.After)
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
	if err = addOpListComplianceStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComplianceStatus(options.Region), middleware.Before); err != nil {
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

// ListComplianceStatusAPIClient is a client that implements the
// ListComplianceStatus operation.
type ListComplianceStatusAPIClient interface {
	ListComplianceStatus(context.Context, *ListComplianceStatusInput, ...func(*Options)) (*ListComplianceStatusOutput, error)
}

var _ ListComplianceStatusAPIClient = (*Client)(nil)

// ListComplianceStatusPaginatorOptions is the paginator options for
// ListComplianceStatus
type ListComplianceStatusPaginatorOptions struct {
	// Specifies the number of PolicyComplianceStatus objects that you want Firewall
	// Manager to return for this request. If you have more PolicyComplianceStatus
	// objects than the number that you specify for MaxResults, the response includes a
	// NextToken value that you can use to get another batch of PolicyComplianceStatus
	// objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComplianceStatusPaginator is a paginator for ListComplianceStatus
type ListComplianceStatusPaginator struct {
	options   ListComplianceStatusPaginatorOptions
	client    ListComplianceStatusAPIClient
	params    *ListComplianceStatusInput
	nextToken *string
	firstPage bool
}

// NewListComplianceStatusPaginator returns a new ListComplianceStatusPaginator
func NewListComplianceStatusPaginator(client ListComplianceStatusAPIClient, params *ListComplianceStatusInput, optFns ...func(*ListComplianceStatusPaginatorOptions)) *ListComplianceStatusPaginator {
	if params == nil {
		params = &ListComplianceStatusInput{}
	}

	options := ListComplianceStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComplianceStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComplianceStatusPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListComplianceStatus page.
func (p *ListComplianceStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComplianceStatusOutput, error) {
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

	result, err := p.client.ListComplianceStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComplianceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListComplianceStatus",
	}
}
