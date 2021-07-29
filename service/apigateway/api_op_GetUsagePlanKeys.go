// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all the usage plan keys representing the API keys added to a specified
// usage plan.
func (c *Client) GetUsagePlanKeys(ctx context.Context, params *GetUsagePlanKeysInput, optFns ...func(*Options)) (*GetUsagePlanKeysOutput, error) {
	if params == nil {
		params = &GetUsagePlanKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsagePlanKeys", params, optFns, c.addOperationGetUsagePlanKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsagePlanKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get all the usage plan keys representing the API keys added
// to a specified usage plan.
type GetUsagePlanKeysInput struct {

	// [Required] The Id of the UsagePlan resource representing the usage plan
	// containing the to-be-retrieved UsagePlanKey resource representing a plan
	// customer.
	//
	// This member is required.
	UsagePlanId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// A query parameter specifying the name of the to-be-returned usage plan keys.
	NameQuery *string

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents the collection of usage plan keys added to usage plans for the
// associated API keys and, possibly, other types of keys. Create and Use Usage
// Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanKeysOutput struct {

	// The current page of elements from this collection.
	Items []types.UsagePlanKey

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsagePlanKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlanKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlanKeys{}, middleware.After)
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
	if err = addOpGetUsagePlanKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlanKeys(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetUsagePlanKeysAPIClient is a client that implements the GetUsagePlanKeys
// operation.
type GetUsagePlanKeysAPIClient interface {
	GetUsagePlanKeys(context.Context, *GetUsagePlanKeysInput, ...func(*Options)) (*GetUsagePlanKeysOutput, error)
}

var _ GetUsagePlanKeysAPIClient = (*Client)(nil)

// GetUsagePlanKeysPaginatorOptions is the paginator options for GetUsagePlanKeys
type GetUsagePlanKeysPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUsagePlanKeysPaginator is a paginator for GetUsagePlanKeys
type GetUsagePlanKeysPaginator struct {
	options   GetUsagePlanKeysPaginatorOptions
	client    GetUsagePlanKeysAPIClient
	params    *GetUsagePlanKeysInput
	nextToken *string
	firstPage bool
}

// NewGetUsagePlanKeysPaginator returns a new GetUsagePlanKeysPaginator
func NewGetUsagePlanKeysPaginator(client GetUsagePlanKeysAPIClient, params *GetUsagePlanKeysInput, optFns ...func(*GetUsagePlanKeysPaginatorOptions)) *GetUsagePlanKeysPaginator {
	if params == nil {
		params = &GetUsagePlanKeysInput{}
	}

	options := GetUsagePlanKeysPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUsagePlanKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUsagePlanKeysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetUsagePlanKeys page.
func (p *GetUsagePlanKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUsagePlanKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetUsagePlanKeys(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetUsagePlanKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlanKeys",
	}
}
