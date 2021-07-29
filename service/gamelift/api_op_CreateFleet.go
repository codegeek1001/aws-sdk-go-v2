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

// Creates a fleet of Amazon Elastic Compute Cloud (Amazon EC2) instances to host
// your custom game server or Realtime Servers. Use this operation to configure the
// computing resources for your fleet and provide instructions for running game
// servers on each instance. Most GameLift fleets can deploy instances to multiple
// locations, including the home Region (where the fleet is created) and an
// optional set of remote locations. Fleets that are created in the following AWS
// Regions support multiple locations: us-east-1 (N. Virginia), us-west-2 (Oregon),
// eu-central-1 (Frankfurt), eu-west-1 (Ireland), ap-southeast-2 (Sydney),
// ap-northeast-1 (Tokyo), and ap-northeast-2 (Seoul). Fleets that are created in
// other GameLift Regions can deploy instances in the fleet's home Region only. All
// fleet instances use the same configuration regardless of location; however, you
// can adjust capacity settings and turn auto-scaling on/off for each location. To
// create a fleet, choose the hardware for your instances, specify a game server
// build or Realtime script to deploy, and provide a runtime configuration to
// direct GameLift how to start and run game servers on each instance in the fleet.
// Set permissions for inbound traffic to your game servers, and enable optional
// features as needed. When creating a multi-location fleet, provide a list of
// additional remote locations. If successful, this operation creates a new Fleet
// resource and places it in NEW status, which prompts GameLift to initiate the
// fleet creation workflow
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creation-workflow.html).
// You can track fleet creation by checking fleet status using
// DescribeFleetAttributes and DescribeFleetLocationAttributes/, or by monitoring
// fleet creation events using DescribeFleetEvents. As soon as the fleet status
// changes to ACTIVE, you can enable automatic scaling for the fleet with
// PutScalingPolicy and set capacity for the home Region with UpdateFleetCapacity.
// When the status of each remote location reaches ACTIVE, you can set capacity by
// location using UpdateFleetCapacity. Learn more Setting up fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)Debug
// fleet creation issues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html#fleets-creating-debug-creation)Multi-location
// fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related actions CreateFleet | UpdateFleetCapacity | PutScalingPolicy |
// DescribeEC2InstanceLimits | DescribeFleetAttributes |
// DescribeFleetLocationAttributes | UpdateFleetAttributes | StopFleetActions |
// DeleteFleet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) {
	if params == nil {
		params = &CreateFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleet", params, optFns, c.addOperationCreateFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type CreateFleetInput struct {

	// The GameLift-supported EC2 instance type to use for all fleet instances.
	// Instance type determines the computing resources that will be used to host your
	// game servers, including CPU, memory, storage, and networking capacity. See
	// Amazon EC2 Instance Types (http://aws.amazon.com/ec2/instance-types/) for
	// detailed descriptions of EC2 instance types.
	//
	// This member is required.
	EC2InstanceType types.EC2InstanceType

	// A descriptive label that is associated with a fleet. Fleet names do not need to
	// be unique.
	//
	// This member is required.
	Name *string

	// The unique identifier for a custom game server build to be deployed on fleet
	// instances. You can use either the build ID or ARN. The build must be uploaded to
	// GameLift and in READY status. This fleet property cannot be changed later.
	BuildId *string

	// Prompts GameLift to generate a TLS/SSL certificate for the fleet. TLS
	// certificates are used for encrypting traffic between game clients and the game
	// servers that are running on GameLift. By default, the CertificateConfiguration
	// is set to DISABLED. Learn more at Securing Client/Server Communication
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-howitworks.html#gamelift-howitworks-security).
	// This property cannot be changed after the fleet is created. Note: This feature
	// requires the AWS Certificate Manager (ACM) service, which is not available in
	// all AWS regions. When working in a region that does not support this feature, a
	// fleet creation request with certificate generation fails with a 4xx error.
	CertificateConfiguration *types.CertificateConfiguration

	// A human-readable description of the fleet.
	Description *string

	// The allowed IP address ranges and port settings that allow inbound traffic to
	// access game sessions on this fleet. If the fleet is hosting a custom game build,
	// this property must be set before players can connect to game sessions. For
	// Realtime Servers fleets, GameLift automatically sets TCP and UDP ranges.
	EC2InboundPermissions []types.IpPermission

	// Indicates whether to use On-Demand or Spot instances for this fleet. By default,
	// this property is set to ON_DEMAND. Learn more about when to use  On-Demand
	// versus Spot Instances
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-ec2-instances.html#gamelift-ec2-instances-spot).
	// This property cannot be changed after the fleet is created.
	FleetType types.FleetType

	// A unique identifier for an AWS IAM role that manages access to your AWS
	// services. With an instance role ARN set, any application that runs on an
	// instance in this fleet can assume the role, including install scripts, server
	// processes, and daemons (background processes). Create a role or look up a role's
	// ARN by using the IAM dashboard (https://console.aws.amazon.com/iam/) in the AWS
	// Management Console. Learn more about using on-box credentials for your game
	// servers at  Access external resources from a game server
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-resources.html).
	// This property cannot be changed after the fleet is created.
	InstanceRoleArn *string

	// A set of remote locations to deploy additional instances to and manage as part
	// of the fleet. This parameter can only be used when creating fleets in AWS
	// Regions that support multiple locations. You can add any GameLift-supported AWS
	// Region as a remote location, in the form of an AWS Region code such as
	// us-west-2. To create a fleet with instances in the home Region only, omit this
	// parameter.
	Locations []types.LocationConfiguration

	// This parameter is no longer used. To specify where GameLift should store log
	// files once a server process shuts down, use the GameLift server API
	// ProcessReady() and specify one or more directory paths in logParameters. See
	// more information in the Server API Reference
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api-ref.html#gamelift-sdk-server-api-ref-dataypes-process).
	LogPaths []string

	// The name of an AWS CloudWatch metric group to add this fleet to. A metric group
	// is used to aggregate the metrics for multiple fleets. You can specify an
	// existing metric group name or set a new name to create a new metric group. A
	// fleet can be included in only one metric group at a time.
	MetricGroups []string

	// The status of termination protection for active game sessions on the fleet. By
	// default, this property is set to NoProtection. You can also set game session
	// protection for an individual game session by calling UpdateGameSession.
	//
	// *
	// NoProtection - Game sessions can be terminated during active gameplay as a
	// result of a scale-down event.
	//
	// * FullProtection - Game sessions in ACTIVE status
	// cannot be terminated during a scale-down event.
	NewGameSessionProtectionPolicy types.ProtectionPolicy

	// Used when peering your GameLift fleet with a VPC, the unique identifier for the
	// AWS account that owns the VPC. You can find your account ID in the AWS
	// Management Console under account settings.
	PeerVpcAwsAccountId *string

	// A unique identifier for a VPC with resources to be accessed by your GameLift
	// fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID,
	// use the VPC Dashboard (https://console.aws.amazon.com/vpc/) in the AWS
	// Management Console. Learn more about VPC peering in VPC Peering with GameLift
	// Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string

	// A policy that limits the number of game sessions that an individual player can
	// create on instances in this fleet within a specified span of time.
	ResourceCreationLimitPolicy *types.ResourceCreationLimitPolicy

	// Instructions for how to launch and maintain server processes on instances in the
	// fleet. The runtime configuration defines one or more server process
	// configurations, each identifying a build executable or Realtime script file and
	// the number of processes of that type to run concurrently. The
	// RuntimeConfiguration parameter is required unless the fleet is being configured
	// using the older parameters ServerLaunchPath and ServerLaunchParameters, which
	// are still supported for backward compatibility.
	RuntimeConfiguration *types.RuntimeConfiguration

	// The unique identifier for a Realtime configuration script to be deployed on
	// fleet instances. You can use either the script ID or ARN. Scripts must be
	// uploaded to GameLift prior to creating the fleet. This fleet property cannot be
	// changed later.
	ScriptId *string

	// This parameter is no longer used. Specify server launch parameters using the
	// RuntimeConfiguration parameter. Requests that use this parameter instead
	// continue to be valid.
	ServerLaunchParameters *string

	// This parameter is no longer used. Specify a server launch path using the
	// RuntimeConfiguration parameter. Requests that use this parameter instead
	// continue to be valid.
	ServerLaunchPath *string

	// A list of labels to assign to the new fleet resource. Tags are developer-defined
	// key-value pairs. Tagging AWS resources are useful for resource management,
	// access management and cost allocation. For more information, see  Tagging AWS
	// Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in
	// the AWS General Reference. Once the fleet is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type CreateFleetOutput struct {

	// The properties for the new fleet, including the current status. All fleets are
	// placed in NEW status on creation.
	FleetAttributes *types.FleetAttributes

	// The fleet's locations and life-cycle status of each location. For new fleets,
	// the status of all locations is set to NEW. During fleet creation, GameLift
	// updates each location status as instances are deployed there and prepared for
	// game hosting. This list includes an entry for the fleet's home Region. For
	// fleets with no remote locations, only one entry, representing the home Region,
	// is returned.
	LocationStates []types.LocationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleet{}, middleware.After)
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
	if err = addOpCreateFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateFleet",
	}
}
