// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets statistics about synchronizing Amazon Kendra with a data source.
func (c *Client) ListDataSourceSyncJobs(ctx context.Context, params *ListDataSourceSyncJobsInput, optFns ...func(*Options)) (*ListDataSourceSyncJobsOutput, error) {
	if params == nil {
		params = &ListDataSourceSyncJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSourceSyncJobs", params, optFns, c.addOperationListDataSourceSyncJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSourceSyncJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSourceSyncJobsInput struct {

	// The identifier of the data source.
	//
	// This member is required.
	Id *string

	// The identifier of the index that contains the data source.
	//
	// This member is required.
	IndexId *string

	// The maximum number of synchronization jobs to return in the response. If there
	// are fewer results in the list, this response contains only the actual results.
	MaxResults *int32

	// If the result of the previous request to GetDataSourceSyncJobHistory was
	// truncated, include the NextToken to fetch the next set of jobs.
	NextToken *string

	// When specified, the synchronization jobs returned in the list are limited to
	// jobs between the specified dates.
	StartTimeFilter *types.TimeRange

	// When specified, only returns synchronization jobs with the Status field equal to
	// the specified status.
	StatusFilter types.DataSourceSyncJobStatus

	noSmithyDocumentSerde
}

type ListDataSourceSyncJobsOutput struct {

	// A history of synchronization jobs for the data source.
	History []types.DataSourceSyncJob

	// The GetDataSourceSyncJobHistory operation returns a page of vocabularies at a
	// time. The maximum size of the page is set by the MaxResults parameter. If there
	// are more jobs in the list than the page size, Amazon Kendra returns the NextPage
	// token. Include the token in the next request to the GetDataSourceSyncJobHistory
	// operation to return in the next page of jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSourceSyncJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataSourceSyncJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataSourceSyncJobs{}, middleware.After)
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
	if err = addOpListDataSourceSyncJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSourceSyncJobs(options.Region), middleware.Before); err != nil {
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

// ListDataSourceSyncJobsAPIClient is a client that implements the
// ListDataSourceSyncJobs operation.
type ListDataSourceSyncJobsAPIClient interface {
	ListDataSourceSyncJobs(context.Context, *ListDataSourceSyncJobsInput, ...func(*Options)) (*ListDataSourceSyncJobsOutput, error)
}

var _ ListDataSourceSyncJobsAPIClient = (*Client)(nil)

// ListDataSourceSyncJobsPaginatorOptions is the paginator options for
// ListDataSourceSyncJobs
type ListDataSourceSyncJobsPaginatorOptions struct {
	// The maximum number of synchronization jobs to return in the response. If there
	// are fewer results in the list, this response contains only the actual results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSourceSyncJobsPaginator is a paginator for ListDataSourceSyncJobs
type ListDataSourceSyncJobsPaginator struct {
	options   ListDataSourceSyncJobsPaginatorOptions
	client    ListDataSourceSyncJobsAPIClient
	params    *ListDataSourceSyncJobsInput
	nextToken *string
	firstPage bool
}

// NewListDataSourceSyncJobsPaginator returns a new ListDataSourceSyncJobsPaginator
func NewListDataSourceSyncJobsPaginator(client ListDataSourceSyncJobsAPIClient, params *ListDataSourceSyncJobsInput, optFns ...func(*ListDataSourceSyncJobsPaginatorOptions)) *ListDataSourceSyncJobsPaginator {
	if params == nil {
		params = &ListDataSourceSyncJobsInput{}
	}

	options := ListDataSourceSyncJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSourceSyncJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSourceSyncJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDataSourceSyncJobs page.
func (p *ListDataSourceSyncJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSourceSyncJobsOutput, error) {
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

	result, err := p.client.ListDataSourceSyncJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSourceSyncJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListDataSourceSyncJobs",
	}
}
