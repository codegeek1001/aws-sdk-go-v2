// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an application version.
func (c *Client) CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplicationVersion", params, optFns, c.addOperationCreateApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationVersionInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The semantic version of the new version.
	//
	// This member is required.
	SemanticVersion *string

	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.Maximum size 50 MB
	SourceCodeArchiveUrl *string

	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string

	// The raw packaged AWS SAM template of your application.
	TemplateBody *string

	// A link to the packaged AWS SAM template of your application.
	TemplateUrl *string

	noSmithyDocumentSerde
}

type CreateApplicationVersionOutput struct {

	// The application Amazon Resource Name (ARN).
	ApplicationId *string

	// The date and time this resource was created.
	CreationTime *string

	// An array of parameter types supported by the application.
	ParameterDefinitions []types.ParameterDefinition

	// A list of values that you must specify before you can deploy certain
	// applications. Some applications might include resources that can affect
	// permissions in your AWS account, for example, by creating new AWS Identity and
	// Access Management (IAM) users. For those applications, you must explicitly
	// acknowledge their capabilities by specifying this parameter.The only valid
	// values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, and
	// CAPABILITY_AUTO_EXPAND.The following resources require you to specify
	// CAPABILITY_IAM or CAPABILITY_NAMED_IAM: AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html),
	// AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html),
	// AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// and AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html).
	// If the application contains IAM resources, you can specify either CAPABILITY_IAM
	// or CAPABILITY_NAMED_IAM. If the application contains IAM resources with custom
	// names, you must specify CAPABILITY_NAMED_IAM.The following resources require you
	// to specify CAPABILITY_RESOURCE_POLICY: AWS::Lambda::Permission
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html),
	// AWS::IAM:Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// AWS::ApplicationAutoScaling::ScalingPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html),
	// AWS::S3::BucketPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html),
	// AWS::SQS::QueuePolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html),
	// and AWS::SNS::TopicPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html).Applications
	// that contain one or more nested applications require you to specify
	// CAPABILITY_AUTO_EXPAND.If your application template contains any of the above
	// resources, we recommend that you review all permissions associated with the
	// application before deploying. If you don't specify this parameter for an
	// application that requires capabilities, the call will fail.
	RequiredCapabilities []types.Capability

	// Whether all of the AWS resources contained in this application are supported in
	// the region in which it is being retrieved.
	ResourcesSupported bool

	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string

	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.Maximum size 50 MB
	SourceCodeArchiveUrl *string

	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string

	// A link to the packaged AWS SAM template of your application.
	TemplateUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplicationVersion{}, middleware.After)
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
	if err = addOpCreateApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "CreateApplicationVersion",
	}
}
