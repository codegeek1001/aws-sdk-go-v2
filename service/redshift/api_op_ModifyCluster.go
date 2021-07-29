// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the settings for a cluster. You can also change node type and the
// number of nodes to scale up or down the cluster. When resizing a cluster, you
// must specify both the number of nodes and the node type even if one of the
// parameters does not change. You can add another security or parameter group, or
// change the admin user password. Resetting a cluster password or modifying the
// security groups associated with a cluster do not need a reboot. However,
// modifying a parameter group requires a reboot for parameters to take effect. For
// more information about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) ModifyCluster(ctx context.Context, params *ModifyClusterInput, optFns ...func(*Options)) (*ModifyClusterOutput, error) {
	if params == nil {
		params = &ModifyClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCluster", params, optFns, c.addOperationModifyClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyClusterInput struct {

	// The unique identifier of the cluster to be modified. Example: examplecluster
	//
	// This member is required.
	ClusterIdentifier *string

	// If true, major version upgrades will be applied automatically to the cluster
	// during the maintenance window. Default: false
	AllowVersionUpgrade *bool

	// The number of days that automated snapshots are retained. If the value is 0,
	// automated snapshots are disabled. Even if automated snapshots are disabled, you
	// can still create manual snapshots when you want with CreateClusterSnapshot. If
	// you decrease the automated snapshot retention period from its current value,
	// existing automated snapshots that fall outside of the new retention period will
	// be immediately deleted. You can't disable automated snapshots for RA3 node
	// types. Set the automated retention period from 1-35 days. Default: Uses existing
	// setting. Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *int32

	// The option to initiate relocation for an Amazon Redshift cluster to the target
	// Availability Zone.
	AvailabilityZone *string

	// The option to enable relocation for an Amazon Redshift cluster between
	// Availability Zones after the cluster modification is complete.
	AvailabilityZoneRelocation *bool

	// The name of the cluster parameter group to apply to this cluster. This change is
	// applied only after the cluster is rebooted. To reboot a cluster use
	// RebootCluster. Default: Uses existing setting. Constraints: The cluster
	// parameter group must be in the same parameter group family that matches the
	// cluster version.
	ClusterParameterGroupName *string

	// A list of cluster security groups to be authorized on this cluster. This change
	// is asynchronously applied as soon as possible. Security groups currently
	// associated with the cluster, and not in the list of groups to apply, will be
	// revoked from the cluster. Constraints:
	//
	// * Must be 1 to 255 alphanumeric
	// characters or hyphens
	//
	// * First character must be a letter
	//
	// * Cannot end with a
	// hyphen or contain two consecutive hyphens
	ClusterSecurityGroups []string

	// The new cluster type. When you submit your cluster resize request, your existing
	// cluster goes into a read-only mode. After Amazon Redshift provisions a new
	// cluster based on your resize requirements, there will be outage for a period
	// while the old cluster is deleted and your connection is switched to the new
	// cluster. You can use DescribeResize to track the progress of the resize request.
	// Valid Values:  multi-node | single-node
	ClusterType *string

	// The new version number of the Amazon Redshift engine to upgrade to. For major
	// version upgrades, if a non-default cluster parameter group is currently in use,
	// a new cluster parameter group in the cluster parameter group family for the new
	// version must be specified. The new cluster parameter group can be the default
	// for that cluster parameter group family. For more information about parameters
	// and parameter groups, go to Amazon Redshift Parameter Groups
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	// in the Amazon Redshift Cluster Management Guide. Example: 1.0
	ClusterVersion *string

	// The Elastic IP (EIP) address for the cluster. Constraints: The cluster must be
	// provisioned in EC2-VPC and publicly-accessible through an Internet gateway. For
	// more information about provisioning clusters in EC2-VPC, go to Supported
	// Platforms to Launch Your Cluster
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms)
	// in the Amazon Redshift Cluster Management Guide.
	ElasticIp *string

	// Indicates whether the cluster is encrypted. If the value is encrypted (true) and
	// you provide a value for the KmsKeyId parameter, we encrypt the cluster with the
	// provided KmsKeyId. If you don't provide a KmsKeyId, we encrypt with the default
	// key. If the value is not encrypted (false), then the cluster is decrypted.
	Encrypted *bool

	// An option that specifies whether to create the cluster with enhanced VPC routing
	// enabled. To create a cluster that uses enhanced VPC routing, the cluster must be
	// in a VPC. For more information, see Enhanced VPC Routing
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in
	// the Amazon Redshift Cluster Management Guide. If this option is true, enhanced
	// VPC routing is enabled. Default: false
	EnhancedVpcRouting *bool

	// Specifies the name of the HSM client certificate the Amazon Redshift cluster
	// uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string

	// Specifies the name of the HSM configuration that contains the information the
	// Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string

	// The Key Management Service (KMS) key ID of the encryption key that you want to
	// use to encrypt data in the cluster.
	KmsKeyId *string

	// The name for the maintenance track that you want to assign for the cluster. This
	// name change is asynchronous. The new track name stays in the
	// PendingModifiedValues for the cluster until the next maintenance window. When
	// the maintenance track changes, the cluster is switched to the latest cluster
	// release available for the maintenance track. At this point, the maintenance
	// track name is applied.
	MaintenanceTrackName *string

	// The default for number of days that a newly created manual snapshot is retained.
	// If the value is -1, the manual snapshot is retained indefinitely. This value
	// doesn't retroactively change the retention periods of existing manual snapshots.
	// The value must be either -1 or an integer between 1 and 3,653. The default value
	// is -1.
	ManualSnapshotRetentionPeriod *int32

	// The new password for the cluster admin user. This change is asynchronously
	// applied as soon as possible. Between the time of the request and the completion
	// of the request, the MasterUserPassword element exists in the
	// PendingModifiedValues element of the operation response. Operations never return
	// the password, so this operation provides a way to regain access to the admin
	// user account for a cluster if the password is lost. Default: Uses existing
	// setting. Constraints:
	//
	// * Must be between 8 and 64 characters in length.
	//
	// * Must
	// contain at least one uppercase letter.
	//
	// * Must contain at least one lowercase
	// letter.
	//
	// * Must contain one number.
	//
	// * Can be any printable ASCII character
	// (ASCII code 33 to 126) except ' (single quote), " (double quote), \, /, @, or
	// space.
	MasterUserPassword *string

	// The new identifier for the cluster. Constraints:
	//
	// * Must contain from 1 to 63
	// alphanumeric characters or hyphens.
	//
	// * Alphabetic characters must be
	// lowercase.
	//
	// * First character must be a letter.
	//
	// * Cannot end with a hyphen or
	// contain two consecutive hyphens.
	//
	// * Must be unique for all clusters within an
	// account.
	//
	// Example: examplecluster
	NewClusterIdentifier *string

	// The new node type of the cluster. If you specify a new node type, you must also
	// specify the number of nodes parameter. For more information about resizing
	// clusters, go to Resizing Clusters in Amazon Redshift
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/rs-resize-tutorial.html) in
	// the Amazon Redshift Cluster Management Guide. Valid Values: ds2.xlarge |
	// ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.xlplus |
	// ra3.4xlarge | ra3.16xlarge
	NodeType *string

	// The new number of nodes of the cluster. If you specify a new number of nodes,
	// you must also specify the node type parameter. For more information about
	// resizing clusters, go to Resizing Clusters in Amazon Redshift
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/rs-resize-tutorial.html) in
	// the Amazon Redshift Cluster Management Guide. Valid Values: Integer greater than
	// 0.
	NumberOfNodes *int32

	// The option to change the port of an Amazon Redshift cluster.
	Port *int32

	// The weekly time range (in UTC) during which system maintenance can occur, if
	// necessary. If system maintenance is necessary during the window, it may result
	// in an outage. This maintenance window change is made immediately. If the new
	// maintenance window indicates the current time, there must be at least 120
	// minutes between the current time and end of the window in order to ensure that
	// pending changes are applied. Default: Uses existing setting. Format:
	// ddd:hh24:mi-ddd:hh24:mi, for example wed:07:30-wed:08:00. Valid Days: Mon | Tue
	// | Wed | Thu | Fri | Sat | Sun Constraints: Must be at least 30 minutes.
	PreferredMaintenanceWindow *string

	// If true, the cluster can be accessed from a public network. Only clusters in
	// VPCs can be set to be publicly available.
	PubliclyAccessible *bool

	// A list of virtual private cloud (VPC) security groups to be associated with the
	// cluster. This change is asynchronously applied as soon as possible.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCluster{}, middleware.After)
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
	if err = addOpModifyClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyCluster",
	}
}
