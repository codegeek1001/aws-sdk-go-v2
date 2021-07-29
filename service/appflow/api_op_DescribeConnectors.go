// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the connectors vended by Amazon AppFlow for specified connector types.
// If you don't specify a connector type, this operation describes all connectors
// vended by Amazon AppFlow. If there are more connectors than can be returned in
// one page, the response contains a nextToken object, which can be be passed in to
// the next call to the DescribeConnectors API operation to retrieve the next page.
func (c *Client) DescribeConnectors(ctx context.Context, params *DescribeConnectorsInput, optFns ...func(*Options)) (*DescribeConnectorsOutput, error) {
	if params == nil {
		params = &DescribeConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectors", params, optFns, c.addOperationDescribeConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorsInput struct {

	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorTypes []types.ConnectorType

	// The pagination token for the next page of data.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConnectorsOutput struct {

	// The configuration that is applied to the connectors used in the flow.
	ConnectorConfigurations map[string]types.ConnectorConfiguration

	// The pagination token for the next page of data.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectors(options.Region), middleware.Before); err != nil {
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

// DescribeConnectorsAPIClient is a client that implements the DescribeConnectors
// operation.
type DescribeConnectorsAPIClient interface {
	DescribeConnectors(context.Context, *DescribeConnectorsInput, ...func(*Options)) (*DescribeConnectorsOutput, error)
}

var _ DescribeConnectorsAPIClient = (*Client)(nil)

// DescribeConnectorsPaginatorOptions is the paginator options for
// DescribeConnectors
type DescribeConnectorsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConnectorsPaginator is a paginator for DescribeConnectors
type DescribeConnectorsPaginator struct {
	options   DescribeConnectorsPaginatorOptions
	client    DescribeConnectorsAPIClient
	params    *DescribeConnectorsInput
	nextToken *string
	firstPage bool
}

// NewDescribeConnectorsPaginator returns a new DescribeConnectorsPaginator
func NewDescribeConnectorsPaginator(client DescribeConnectorsAPIClient, params *DescribeConnectorsInput, optFns ...func(*DescribeConnectorsPaginatorOptions)) *DescribeConnectorsPaginator {
	if params == nil {
		params = &DescribeConnectorsInput{}
	}

	options := DescribeConnectorsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeConnectors page.
func (p *DescribeConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConnectorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeConnectors",
	}
}
