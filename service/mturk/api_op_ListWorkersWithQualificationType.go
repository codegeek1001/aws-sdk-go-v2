// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListWorkersWithQualificationType operation returns all of the Workers that
// have been associated with a given Qualification type.
func (c *Client) ListWorkersWithQualificationType(ctx context.Context, params *ListWorkersWithQualificationTypeInput, optFns ...func(*Options)) (*ListWorkersWithQualificationTypeOutput, error) {
	if params == nil {
		params = &ListWorkersWithQualificationTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkersWithQualificationType", params, optFns, c.addOperationListWorkersWithQualificationTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkersWithQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkersWithQualificationTypeInput struct {

	// The ID of the Qualification type of the Qualifications to return.
	//
	// This member is required.
	QualificationTypeId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination Token
	NextToken *string

	// The status of the Qualifications to return. Can be Granted | Revoked.
	Status types.QualificationStatus

	noSmithyDocumentSerde
}

type ListWorkersWithQualificationTypeOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of Qualifications on this page in the filtered results list,
	// equivalent to the number of Qualifications being returned by this call.
	NumResults *int32

	// The list of Qualification elements returned by this call.
	Qualifications []types.Qualification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkersWithQualificationTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkersWithQualificationType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkersWithQualificationType{}, middleware.After)
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
	if err = addOpListWorkersWithQualificationTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkersWithQualificationType(options.Region), middleware.Before); err != nil {
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

// ListWorkersWithQualificationTypeAPIClient is a client that implements the
// ListWorkersWithQualificationType operation.
type ListWorkersWithQualificationTypeAPIClient interface {
	ListWorkersWithQualificationType(context.Context, *ListWorkersWithQualificationTypeInput, ...func(*Options)) (*ListWorkersWithQualificationTypeOutput, error)
}

var _ ListWorkersWithQualificationTypeAPIClient = (*Client)(nil)

// ListWorkersWithQualificationTypePaginatorOptions is the paginator options for
// ListWorkersWithQualificationType
type ListWorkersWithQualificationTypePaginatorOptions struct {
	// Limit the number of results returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkersWithQualificationTypePaginator is a paginator for
// ListWorkersWithQualificationType
type ListWorkersWithQualificationTypePaginator struct {
	options   ListWorkersWithQualificationTypePaginatorOptions
	client    ListWorkersWithQualificationTypeAPIClient
	params    *ListWorkersWithQualificationTypeInput
	nextToken *string
	firstPage bool
}

// NewListWorkersWithQualificationTypePaginator returns a new
// ListWorkersWithQualificationTypePaginator
func NewListWorkersWithQualificationTypePaginator(client ListWorkersWithQualificationTypeAPIClient, params *ListWorkersWithQualificationTypeInput, optFns ...func(*ListWorkersWithQualificationTypePaginatorOptions)) *ListWorkersWithQualificationTypePaginator {
	if params == nil {
		params = &ListWorkersWithQualificationTypeInput{}
	}

	options := ListWorkersWithQualificationTypePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkersWithQualificationTypePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkersWithQualificationTypePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListWorkersWithQualificationType page.
func (p *ListWorkersWithQualificationTypePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkersWithQualificationTypeOutput, error) {
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

	result, err := p.client.ListWorkersWithQualificationType(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkersWithQualificationType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListWorkersWithQualificationType",
	}
}
