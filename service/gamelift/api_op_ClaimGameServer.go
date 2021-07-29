// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is used with the GameLift FleetIQ solution and game server
// groups. Locates an available game server and temporarily reserves it to host
// gameplay and players. This operation is called from a game client or client
// service (such as a matchmaker) to request hosting resources for a new game
// session. In response, GameLift FleetIQ locates an available game server, places
// it in CLAIMED status for 60 seconds, and returns connection information that
// players can use to connect to the game server. To claim a game server, identify
// a game server group. You can also specify a game server ID, although this
// approach bypasses GameLift FleetIQ placement optimization. Optionally, include
// game data to pass to the game server at the start of a game session, such as a
// game map or player information. When a game server is successfully claimed,
// connection information is returned. A claimed game server's utilization status
// remains AVAILABLE while the claim status is set to CLAIMED for up to 60 seconds.
// This time period gives the game server time to update its status to UTILIZED
// (using UpdateGameServer) once players join. If the game server's status is not
// updated within 60 seconds, the game server reverts to unclaimed status and is
// available to be claimed by another request. The claim time period is a fixed
// value and is not configurable. If you try to claim a specific game server, this
// request will fail in the following cases:
//
// * If the game server utilization
// status is UTILIZED.
//
// * If the game server claim status is CLAIMED.
//
// When
// claiming a specific game server, this request will succeed even if the game
// server is running on an instance in DRAINING status. To avoid this, first check
// the instance status by calling DescribeGameServerInstances. Learn more GameLift
// FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related actions RegisterGameServer | ListGameServers | ClaimGameServer |
// DescribeGameServer | UpdateGameServer | DeregisterGameServer | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/reference-awssdk-fleetiq.html)
func (c *Client) ClaimGameServer(ctx context.Context, params *ClaimGameServerInput, optFns ...func(*Options)) (*ClaimGameServerOutput, error) {
	if params == nil {
		params = &ClaimGameServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ClaimGameServer", params, optFns, c.addOperationClaimGameServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ClaimGameServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ClaimGameServerInput struct {

	// A unique identifier for the game server group where the game server is running.
	// Use either the GameServerGroup name or ARN value. If you are not specifying a
	// game server to claim, this value identifies where you want GameLift FleetIQ to
	// look for an available game server to claim.
	//
	// This member is required.
	GameServerGroupName *string

	// A set of custom game server properties, formatted as a single string value. This
	// data is passed to a game client or service when it requests information on game
	// servers using ListGameServers or ClaimGameServer.
	GameServerData *string

	// A custom string that uniquely identifies the game server to claim. If this
	// parameter is left empty, GameLift FleetIQ searches for an available game server
	// in the specified game server group.
	GameServerId *string

	noSmithyDocumentSerde
}

type ClaimGameServerOutput struct {

	// Object that describes the newly claimed game server.
	GameServer *types.GameServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationClaimGameServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpClaimGameServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpClaimGameServer{}, middleware.After)
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
	if err = addOpClaimGameServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opClaimGameServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opClaimGameServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ClaimGameServer",
	}
}
