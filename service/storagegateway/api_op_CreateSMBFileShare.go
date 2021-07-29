// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Server Message Block (SMB) file share on an existing S3 File Gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway exposes file shares using an SMB interface.
// This operation is only supported for S3 File Gateways. S3 File Gateways require
// Security Token Service (STS) to be activated to enable you to create a file
// share. Make sure that STS is activated in the Region you are creating your S3
// File Gateway in. If STS is not activated in this Region, activate it. For
// information about how to activate STS, see Activating and deactivating STS in an
// Region
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the Identity and Access Management User Guide. File gateways don't support
// creating hard or symbolic links on a file share.
func (c *Client) CreateSMBFileShare(ctx context.Context, params *CreateSMBFileShareInput, optFns ...func(*Options)) (*CreateSMBFileShareOutput, error) {
	if params == nil {
		params = &CreateSMBFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSMBFileShare", params, optFns, c.addOperationCreateSMBFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSMBFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateSMBFileShareInput
type CreateSMBFileShareInput struct {

	// A unique string value that you supply that is used by S3 File Gateway to ensure
	// idempotent file share creation.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the S3 File Gateway on which you want to create a file share.
	//
	// This member is required.
	GatewayARN *string

	// The ARN of the backend storage used for storing file data. A prefix name can be
	// added to the S3 bucket name. It must end with a "/". You can specify a bucket
	// attached to an access point using a complete ARN that includes the bucket region
	// as shown: arn:aws:s3:region:account-id:accesspoint/access-point-name  If you
	// specify a bucket attached to an access point, the bucket policy must be
	// configured to delegate access control to the access point. For information, see
	// Delegating access control to access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-policies.html#access-points-delegating-control)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	LocationARN *string

	// The ARN of the Identity and Access Management (IAM) role that an S3 File Gateway
	// assumes when it accesses the underlying storage.
	//
	// This member is required.
	Role *string

	// The files and folders on this share will only be visible to users with read
	// access.
	AccessBasedEnumeration *bool

	// A list of users or groups in the Active Directory that will be granted
	// administrator privileges on the file share. These users can do all file
	// operations as the super-user. Acceptable formats include: DOMAIN\User1, user1,
	// @group1, and @DOMAIN\group1. Use this option very carefully, because any user in
	// this list can do anything they like on the file share, regardless of file
	// permissions.
	AdminUserList []string

	// The Amazon Resource Name (ARN) of the storage used for audit logs.
	AuditDestinationARN *string

	// The authentication method that users use to access the file share. The default
	// is ActiveDirectory. Valid Values: ActiveDirectory | GuestAccess
	Authentication *string

	// Specifies the Region of the S3 bucket where the SMB file share stores files.
	// This parameter is required for SMB file shares that connect to Amazon S3 through
	// a VPC endpoint, a VPC access point, or an access point alias that points to a
	// VPC access point.
	BucketRegion *string

	// Specifies refresh cache information for the file share.
	CacheAttributes *types.CacheAttributes

	// The case of an object name in an Amazon S3 bucket. For ClientSpecified, the
	// client determines the case sensitivity. For CaseSensitive, the gateway
	// determines the case sensitivity. The default value is ClientSpecified.
	CaseSensitivity types.CaseSensitivity

	// The default storage class for objects put into an Amazon S3 bucket by the S3
	// File Gateway. The default value is S3_INTELLIGENT_TIERING. Optional. Valid
	// Values: S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional. FileShareName must be set if an S3 prefix
	// name is set in LocationARN.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false. The default value is true. Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// A list of users or groups in the Active Directory that are not allowed to access
	// the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only be
	// set if Authentication is set to ActiveDirectory.
	InvalidUserList []string

	// Set to true to use Amazon S3 server-side encryption with your own KMS key, or
	// false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The notification policy of the file share. SettlingTimeInSeconds controls the
	// number of seconds to wait after the last point in time a client wrote to a file
	// before generating an ObjectUploaded notification. Because clients can make many
	// small writes to files, it's best to set this parameter for as long as possible
	// to avoid generating multiple notifications for the same file in a small time
	// period. SettlingTimeInSeconds has no effect on the timing of the object
	// uploading to Amazon S3, only the timing of the notification. The following
	// example sets NotificationPolicy on with SettlingTimeInSeconds set to 60.
	// {"Upload": {"SettlingTimeInSeconds": 60}} The following example sets
	// NotificationPolicy off. {}
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the S3
	// bucket that a S3 File Gateway puts objects into. The default value is private.
	ObjectACL types.ObjectACL

	// Specifies whether opportunistic locking is enabled for the SMB file share.
	// Enabling opportunistic locking on case-sensitive shares is not recommended for
	// workloads that involve access to files with the same name in different case.
	// Valid Values: true | false
	OplocksEnabled *bool

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false. Valid Values: true |
	// false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true, the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data. RequesterPays is a configuration for
	// the S3 bucket that backs the file share, so make sure that the configuration on
	// the file share is the same as the S3 bucket configuration. Valid Values: true |
	// false
	RequesterPays *bool

	// Set this value to true to enable access control list (ACL) on the SMB file
	// share. Set it to false to map file and directory permissions to the POSIX
	// permissions. For more information, see Using Microsoft Windows ACLs to control
	// access to an SMB file share
	// (https://docs.aws.amazon.com/storagegateway/latest/userguide/smb-acl.html) in
	// the Storage Gateway User Guide. Valid Values: true | false
	SMBACLEnabled *bool

	// A list of up to 50 tags that can be assigned to the NFS file share. Each tag is
	// a key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag

	// Specifies the DNS name for the VPC endpoint that the SMB file share uses to
	// connect to Amazon S3. This parameter is required for SMB file shares that
	// connect to Amazon S3 through a VPC endpoint, a VPC access point, or an access
	// point alias that points to a VPC access point.
	VPCEndpointDNSName *string

	// A list of users or groups in the Active Directory that are allowed to access the
	// file  share. A group must be prefixed with the @ character. Acceptable formats
	// include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only be set if
	// Authentication is set to ActiveDirectory.
	ValidUserList []string

	noSmithyDocumentSerde
}

// CreateSMBFileShareOutput
type CreateSMBFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSMBFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSMBFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSMBFileShare{}, middleware.After)
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
	if err = addOpCreateSMBFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSMBFileShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSMBFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateSMBFileShare",
	}
}
