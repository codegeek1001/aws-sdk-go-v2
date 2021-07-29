// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed status for each member account within an organization for a
// given organization conformance pack.
func (c *Client) GetOrganizationConformancePackDetailedStatus(ctx context.Context, params *GetOrganizationConformancePackDetailedStatusInput, optFns ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error) {
	if params == nil {
		params = &GetOrganizationConformancePackDetailedStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOrganizationConformancePackDetailedStatus", params, optFns, c.addOperationGetOrganizationConformancePackDetailedStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOrganizationConformancePackDetailedStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOrganizationConformancePackDetailedStatusInput struct {

	// The name of organization conformance pack for which you want status details for
	// member accounts.
	//
	// This member is required.
	OrganizationConformancePackName *string

	// An OrganizationResourceDetailedStatusFilters object.
	Filters *types.OrganizationResourceDetailedStatusFilters

	// The maximum number of OrganizationConformancePackDetailedStatuses returned on
	// each page. If you do not specify a number, AWS Config uses the default. The
	// default is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetOrganizationConformancePackDetailedStatusOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// A list of OrganizationConformancePackDetailedStatus objects.
	OrganizationConformancePackDetailedStatuses []types.OrganizationConformancePackDetailedStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOrganizationConformancePackDetailedStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOrganizationConformancePackDetailedStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOrganizationConformancePackDetailedStatus{}, middleware.After)
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
	if err = addOpGetOrganizationConformancePackDetailedStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(options.Region), middleware.Before); err != nil {
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

// GetOrganizationConformancePackDetailedStatusAPIClient is a client that
// implements the GetOrganizationConformancePackDetailedStatus operation.
type GetOrganizationConformancePackDetailedStatusAPIClient interface {
	GetOrganizationConformancePackDetailedStatus(context.Context, *GetOrganizationConformancePackDetailedStatusInput, ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error)
}

var _ GetOrganizationConformancePackDetailedStatusAPIClient = (*Client)(nil)

// GetOrganizationConformancePackDetailedStatusPaginatorOptions is the paginator
// options for GetOrganizationConformancePackDetailedStatus
type GetOrganizationConformancePackDetailedStatusPaginatorOptions struct {
	// The maximum number of OrganizationConformancePackDetailedStatuses returned on
	// each page. If you do not specify a number, AWS Config uses the default. The
	// default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOrganizationConformancePackDetailedStatusPaginator is a paginator for
// GetOrganizationConformancePackDetailedStatus
type GetOrganizationConformancePackDetailedStatusPaginator struct {
	options   GetOrganizationConformancePackDetailedStatusPaginatorOptions
	client    GetOrganizationConformancePackDetailedStatusAPIClient
	params    *GetOrganizationConformancePackDetailedStatusInput
	nextToken *string
	firstPage bool
}

// NewGetOrganizationConformancePackDetailedStatusPaginator returns a new
// GetOrganizationConformancePackDetailedStatusPaginator
func NewGetOrganizationConformancePackDetailedStatusPaginator(client GetOrganizationConformancePackDetailedStatusAPIClient, params *GetOrganizationConformancePackDetailedStatusInput, optFns ...func(*GetOrganizationConformancePackDetailedStatusPaginatorOptions)) *GetOrganizationConformancePackDetailedStatusPaginator {
	if params == nil {
		params = &GetOrganizationConformancePackDetailedStatusInput{}
	}

	options := GetOrganizationConformancePackDetailedStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOrganizationConformancePackDetailedStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOrganizationConformancePackDetailedStatusPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetOrganizationConformancePackDetailedStatus page.
func (p *GetOrganizationConformancePackDetailedStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOrganizationConformancePackDetailedStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetOrganizationConformancePackDetailedStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetOrganizationConformancePackDetailedStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetOrganizationConformancePackDetailedStatus",
	}
}
