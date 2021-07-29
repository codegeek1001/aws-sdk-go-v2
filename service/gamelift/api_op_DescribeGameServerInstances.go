// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the GameLift FleetIQ solution and game server
// groups. Retrieves status information about the Amazon EC2 instances associated
// with a GameLift FleetIQ game server group. Use this operation to detect when
// instances are active or not available to host new game servers. If you are
// looking for instance configuration information, call DescribeGameServerGroup or
// access the corresponding Auto Scaling group properties. To request status for
// all instances in the game server group, provide a game server group ID only. To
// request status for specific instances, provide the game server group ID and one
// or more instance IDs. Use the pagination parameters to retrieve results in
// sequential segments. If successful, a collection of GameServerInstance objects
// is returned. This operation is not designed to be called with every game server
// claim request; this practice can cause you to exceed your API limit, which
// results in errors. Instead, as a best practice, cache the results and refresh
// your cache no more than once every 10 seconds. Learn more GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related actions CreateGameServerGroup | ListGameServerGroups |
// DescribeGameServerGroup | UpdateGameServerGroup | DeleteGameServerGroup |
// ResumeGameServerGroup | SuspendGameServerGroup | DescribeGameServerInstances |
// All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/reference-awssdk-fleetiq.html)
func (c *Client) DescribeGameServerInstances(ctx context.Context, params *DescribeGameServerInstancesInput, optFns ...func(*Options)) (*DescribeGameServerInstancesOutput, error) {
	if params == nil {
		params = &DescribeGameServerInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameServerInstances", params, optFns, c.addOperationDescribeGameServerInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameServerInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameServerInstancesInput struct {

	// A unique identifier for the game server group. Use either the GameServerGroup
	// name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The EC2 instance IDs that you want to retrieve status on. EC2 instance IDs use a
	// 17-character format, for example: i-1234567890abcdef0. To retrieve all instances
	// in the game server group, leave this parameter empty.
	InstanceIds []string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeGameServerInstancesOutput struct {

	// The collection of requested game server instances.
	GameServerInstances []types.GameServerInstance

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGameServerInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameServerInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameServerInstances{}, middleware.After)
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
	if err = addOpDescribeGameServerInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameServerInstances(options.Region), middleware.Before); err != nil {
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

// DescribeGameServerInstancesAPIClient is a client that implements the
// DescribeGameServerInstances operation.
type DescribeGameServerInstancesAPIClient interface {
	DescribeGameServerInstances(context.Context, *DescribeGameServerInstancesInput, ...func(*Options)) (*DescribeGameServerInstancesOutput, error)
}

var _ DescribeGameServerInstancesAPIClient = (*Client)(nil)

// DescribeGameServerInstancesPaginatorOptions is the paginator options for
// DescribeGameServerInstances
type DescribeGameServerInstancesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGameServerInstancesPaginator is a paginator for
// DescribeGameServerInstances
type DescribeGameServerInstancesPaginator struct {
	options   DescribeGameServerInstancesPaginatorOptions
	client    DescribeGameServerInstancesAPIClient
	params    *DescribeGameServerInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeGameServerInstancesPaginator returns a new
// DescribeGameServerInstancesPaginator
func NewDescribeGameServerInstancesPaginator(client DescribeGameServerInstancesAPIClient, params *DescribeGameServerInstancesInput, optFns ...func(*DescribeGameServerInstancesPaginatorOptions)) *DescribeGameServerInstancesPaginator {
	if params == nil {
		params = &DescribeGameServerInstancesInput{}
	}

	options := DescribeGameServerInstancesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeGameServerInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGameServerInstancesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeGameServerInstances page.
func (p *DescribeGameServerInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGameServerInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeGameServerInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeGameServerInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameServerInstances",
	}
}
