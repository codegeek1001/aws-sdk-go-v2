// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an estimate of the monthly cost for DevOps Guru to analyze your AWS
// resources. For more information, see Estimate your Amazon DevOps Guru costs
// (https://docs.aws.amazon.com/devops-guru/latest/userguide/cost-estimate.html)
// and Amazon DevOps Guru pricing (http://aws.amazon.com/devops-guru/pricing/).
func (c *Client) GetCostEstimation(ctx context.Context, params *GetCostEstimationInput, optFns ...func(*Options)) (*GetCostEstimationOutput, error) {
	if params == nil {
		params = &GetCostEstimationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostEstimation", params, optFns, c.addOperationGetCostEstimationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostEstimationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostEstimationInput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCostEstimationOutput struct {

	// An array of ResourceCost objects that each contains details about the monthly
	// cost estimate to analyze one of your AWS resources.
	Costs []types.ServiceResourceCost

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// The collection of the AWS resources used to create your monthly DevOps Guru cost
	// estimate.
	ResourceCollection *types.CostEstimationResourceCollectionFilter

	// The status of creating this cost estimate. If it's still in progress, the status
	// ONGOING is returned. If it is finished, the status COMPLETED is returned.
	Status types.CostEstimationStatus

	// The start and end time of the cost estimation.
	TimeRange *types.CostEstimationTimeRange

	// The estimated monthly cost to analyze the AWS resources. This value is the sum
	// of the estimated costs to analyze each resource in the Costs object in this
	// response.
	TotalCost float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostEstimationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCostEstimation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCostEstimation{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostEstimation(options.Region), middleware.Before); err != nil {
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

// GetCostEstimationAPIClient is a client that implements the GetCostEstimation
// operation.
type GetCostEstimationAPIClient interface {
	GetCostEstimation(context.Context, *GetCostEstimationInput, ...func(*Options)) (*GetCostEstimationOutput, error)
}

var _ GetCostEstimationAPIClient = (*Client)(nil)

// GetCostEstimationPaginatorOptions is the paginator options for GetCostEstimation
type GetCostEstimationPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCostEstimationPaginator is a paginator for GetCostEstimation
type GetCostEstimationPaginator struct {
	options   GetCostEstimationPaginatorOptions
	client    GetCostEstimationAPIClient
	params    *GetCostEstimationInput
	nextToken *string
	firstPage bool
}

// NewGetCostEstimationPaginator returns a new GetCostEstimationPaginator
func NewGetCostEstimationPaginator(client GetCostEstimationAPIClient, params *GetCostEstimationInput, optFns ...func(*GetCostEstimationPaginatorOptions)) *GetCostEstimationPaginator {
	if params == nil {
		params = &GetCostEstimationInput{}
	}

	options := GetCostEstimationPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCostEstimationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCostEstimationPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetCostEstimation page.
func (p *GetCostEstimationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCostEstimationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetCostEstimation(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCostEstimation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "GetCostEstimation",
	}
}
