// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get interpolated values for an asset property for a specified time interval,
// during a period of time. For example, you can use the this operation to return
// the interpolated temperature values for a wind turbine every 24 hours over a
// duration of 7 days. To identify an asset property, you must specify one of the
// following:
//
// * The assetId and propertyId of an asset property.
//
// * A
// propertyAlias, which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature). To define an asset property's alias,
// see UpdateAssetProperty
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
func (c *Client) GetInterpolatedAssetPropertyValues(ctx context.Context, params *GetInterpolatedAssetPropertyValuesInput, optFns ...func(*Options)) (*GetInterpolatedAssetPropertyValuesOutput, error) {
	if params == nil {
		params = &GetInterpolatedAssetPropertyValuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInterpolatedAssetPropertyValues", params, optFns, c.addOperationGetInterpolatedAssetPropertyValuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInterpolatedAssetPropertyValuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInterpolatedAssetPropertyValuesInput struct {

	// The inclusive end of the range from which to interpolate data, expressed in
	// seconds in Unix epoch time.
	//
	// This member is required.
	EndTimeInSeconds *int64

	// The time interval in seconds over which to interpolate data. Each interval
	// starts when the previous one ends.
	//
	// This member is required.
	IntervalInSeconds *int64

	// The quality of the asset property value. You can use this parameter as a filter
	// to choose only the asset property values that have a specific quality.
	//
	// This member is required.
	Quality types.Quality

	// The exclusive start of the range from which to interpolate data, expressed in
	// seconds in Unix epoch time.
	//
	// This member is required.
	StartTimeInSeconds *int64

	// The interpolation type. Valid values: LINEAR_INTERPOLATION
	//
	// This member is required.
	Type *string

	// The ID of the asset.
	AssetId *string

	// The nanosecond offset converted from endTimeInSeconds.
	EndTimeOffsetInNanos *int32

	// The maximum number of results to return for each paginated request. If not
	// specified, the default value is 10.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The alias that identifies the property, such as an OPC-UA server data stream
	// path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping industrial data streams to asset properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property.
	PropertyId *string

	// The nanosecond offset converted from startTimeInSeconds.
	StartTimeOffsetInNanos *int32

	noSmithyDocumentSerde
}

type GetInterpolatedAssetPropertyValuesOutput struct {

	// The requested interpolated values.
	//
	// This member is required.
	InterpolatedAssetPropertyValues []types.InterpolatedAssetPropertyValue

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInterpolatedAssetPropertyValuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInterpolatedAssetPropertyValues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInterpolatedAssetPropertyValues{}, middleware.After)
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
	if err = addEndpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetInterpolatedAssetPropertyValuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInterpolatedAssetPropertyValues(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware struct {
}

func (*endpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetInterpolatedAssetPropertyValuesMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetInterpolatedAssetPropertyValuesAPIClient is a client that implements the
// GetInterpolatedAssetPropertyValues operation.
type GetInterpolatedAssetPropertyValuesAPIClient interface {
	GetInterpolatedAssetPropertyValues(context.Context, *GetInterpolatedAssetPropertyValuesInput, ...func(*Options)) (*GetInterpolatedAssetPropertyValuesOutput, error)
}

var _ GetInterpolatedAssetPropertyValuesAPIClient = (*Client)(nil)

// GetInterpolatedAssetPropertyValuesPaginatorOptions is the paginator options for
// GetInterpolatedAssetPropertyValues
type GetInterpolatedAssetPropertyValuesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. If not
	// specified, the default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInterpolatedAssetPropertyValuesPaginator is a paginator for
// GetInterpolatedAssetPropertyValues
type GetInterpolatedAssetPropertyValuesPaginator struct {
	options   GetInterpolatedAssetPropertyValuesPaginatorOptions
	client    GetInterpolatedAssetPropertyValuesAPIClient
	params    *GetInterpolatedAssetPropertyValuesInput
	nextToken *string
	firstPage bool
}

// NewGetInterpolatedAssetPropertyValuesPaginator returns a new
// GetInterpolatedAssetPropertyValuesPaginator
func NewGetInterpolatedAssetPropertyValuesPaginator(client GetInterpolatedAssetPropertyValuesAPIClient, params *GetInterpolatedAssetPropertyValuesInput, optFns ...func(*GetInterpolatedAssetPropertyValuesPaginatorOptions)) *GetInterpolatedAssetPropertyValuesPaginator {
	if params == nil {
		params = &GetInterpolatedAssetPropertyValuesInput{}
	}

	options := GetInterpolatedAssetPropertyValuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInterpolatedAssetPropertyValuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInterpolatedAssetPropertyValuesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetInterpolatedAssetPropertyValues page.
func (p *GetInterpolatedAssetPropertyValuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInterpolatedAssetPropertyValuesOutput, error) {
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

	result, err := p.client.GetInterpolatedAssetPropertyValues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInterpolatedAssetPropertyValues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "GetInterpolatedAssetPropertyValues",
	}
}
