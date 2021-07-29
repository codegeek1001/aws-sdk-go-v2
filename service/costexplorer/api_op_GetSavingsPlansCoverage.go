// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the Savings Plans covered for your account. This enables you to see
// how much of your cost is covered by a Savings Plan. An organization’s management
// account can see the coverage of the associated member accounts. This supports
// dimensions, Cost Categories, and nested expressions. For any time period, you
// can filter data for Savings Plans usage with the following dimensions:
//
// *
// LINKED_ACCOUNT
//
// * REGION
//
// * SERVICE
//
// * INSTANCE_FAMILY
//
// To determine valid
// values for a dimension, use the GetDimensionValues operation.
func (c *Client) GetSavingsPlansCoverage(ctx context.Context, params *GetSavingsPlansCoverageInput, optFns ...func(*Options)) (*GetSavingsPlansCoverageOutput, error) {
	if params == nil {
		params = &GetSavingsPlansCoverageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansCoverage", params, optFns, c.addOperationGetSavingsPlansCoverageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansCoverageInput struct {

	// The time period that you want the usage and costs for. The Start date must be
	// within 13 months. The End date must be after the Start date, and before the
	// current date. Future dates can't be used as an End date.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters Savings Plans coverage data by dimensions. You can filter data for
	// Savings Plans usage with the following dimensions:
	//
	// * LINKED_ACCOUNT
	//
	// *
	// REGION
	//
	// * SERVICE
	//
	// * INSTANCE_FAMILY
	//
	// GetSavingsPlansCoverage uses the same
	// Expression
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	// If there are multiple values for a dimension, they are OR'd together. Cost
	// category is also supported.
	Filter *types.Expression

	// The granularity of the Amazon Web Services cost data for your Savings Plans.
	// Granularity can't be set if GroupBy is set. The GetSavingsPlansCoverage
	// operation supports only DAILY and MONTHLY granularities.
	Granularity types.Granularity

	// You can group the data using the attributes INSTANCE_FAMILY, REGION, or SERVICE.
	GroupBy []types.GroupDefinition

	// The number of items to be returned in a response. The default is 20, with a
	// minimum value of 1.
	MaxResults int32

	// The measurement that you want your Savings Plans coverage reported in. The only
	// valid value is SpendCoveredBySavingsPlans.
	Metrics []string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// The value by which you want to sort the data. The following values are supported
	// for Key:
	//
	// * SpendCoveredBySavingsPlan
	//
	// * OnDemandCost
	//
	// * CoveragePercentage
	//
	// *
	// TotalCost
	//
	// * InstanceFamily
	//
	// * Region
	//
	// * Service
	//
	// Supported values for SortOrder
	// are ASCENDING or DESCENDING.
	SortBy *types.SortDefinition

	noSmithyDocumentSerde
}

type GetSavingsPlansCoverageOutput struct {

	// The amount of spend that your Savings Plans covered.
	//
	// This member is required.
	SavingsPlansCoverages []types.SavingsPlansCoverage

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSavingsPlansCoverageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansCoverage{}, middleware.After)
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
	if err = addOpGetSavingsPlansCoverageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansCoverage(options.Region), middleware.Before); err != nil {
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

// GetSavingsPlansCoverageAPIClient is a client that implements the
// GetSavingsPlansCoverage operation.
type GetSavingsPlansCoverageAPIClient interface {
	GetSavingsPlansCoverage(context.Context, *GetSavingsPlansCoverageInput, ...func(*Options)) (*GetSavingsPlansCoverageOutput, error)
}

var _ GetSavingsPlansCoverageAPIClient = (*Client)(nil)

// GetSavingsPlansCoveragePaginatorOptions is the paginator options for
// GetSavingsPlansCoverage
type GetSavingsPlansCoveragePaginatorOptions struct {
	// The number of items to be returned in a response. The default is 20, with a
	// minimum value of 1.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSavingsPlansCoveragePaginator is a paginator for GetSavingsPlansCoverage
type GetSavingsPlansCoveragePaginator struct {
	options   GetSavingsPlansCoveragePaginatorOptions
	client    GetSavingsPlansCoverageAPIClient
	params    *GetSavingsPlansCoverageInput
	nextToken *string
	firstPage bool
}

// NewGetSavingsPlansCoveragePaginator returns a new
// GetSavingsPlansCoveragePaginator
func NewGetSavingsPlansCoveragePaginator(client GetSavingsPlansCoverageAPIClient, params *GetSavingsPlansCoverageInput, optFns ...func(*GetSavingsPlansCoveragePaginatorOptions)) *GetSavingsPlansCoveragePaginator {
	if params == nil {
		params = &GetSavingsPlansCoverageInput{}
	}

	options := GetSavingsPlansCoveragePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSavingsPlansCoveragePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSavingsPlansCoveragePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetSavingsPlansCoverage page.
func (p *GetSavingsPlansCoveragePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSavingsPlansCoverageOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetSavingsPlansCoverage(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetSavingsPlansCoverage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansCoverage",
	}
}
