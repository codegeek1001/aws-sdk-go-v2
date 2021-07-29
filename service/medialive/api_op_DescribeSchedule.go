// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a channel schedule
func (c *Client) DescribeSchedule(ctx context.Context, params *DescribeScheduleInput, optFns ...func(*Options)) (*DescribeScheduleOutput, error) {
	if params == nil {
		params = &DescribeScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSchedule", params, optFns, c.addOperationDescribeScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeScheduleRequest
type DescribeScheduleInput struct {

	// Id of the channel whose schedule is being updated.
	//
	// This member is required.
	ChannelId *string

	// Placeholder documentation for MaxResults
	MaxResults int32

	// Placeholder documentation for __string
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DescribeScheduleResponse
type DescribeScheduleOutput struct {

	// The next token; for use in pagination.
	NextToken *string

	// The list of actions in the schedule.
	ScheduleActions []types.ScheduleAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSchedule{}, middleware.After)
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
	if err = addOpDescribeScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSchedule(options.Region), middleware.Before); err != nil {
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

// DescribeScheduleAPIClient is a client that implements the DescribeSchedule
// operation.
type DescribeScheduleAPIClient interface {
	DescribeSchedule(context.Context, *DescribeScheduleInput, ...func(*Options)) (*DescribeScheduleOutput, error)
}

var _ DescribeScheduleAPIClient = (*Client)(nil)

// DescribeSchedulePaginatorOptions is the paginator options for DescribeSchedule
type DescribeSchedulePaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSchedulePaginator is a paginator for DescribeSchedule
type DescribeSchedulePaginator struct {
	options   DescribeSchedulePaginatorOptions
	client    DescribeScheduleAPIClient
	params    *DescribeScheduleInput
	nextToken *string
	firstPage bool
}

// NewDescribeSchedulePaginator returns a new DescribeSchedulePaginator
func NewDescribeSchedulePaginator(client DescribeScheduleAPIClient, params *DescribeScheduleInput, optFns ...func(*DescribeSchedulePaginatorOptions)) *DescribeSchedulePaginator {
	if params == nil {
		params = &DescribeScheduleInput{}
	}

	options := DescribeSchedulePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSchedulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSchedulePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSchedule page.
func (p *DescribeSchedulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScheduleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeSchedule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeSchedule",
	}
}
