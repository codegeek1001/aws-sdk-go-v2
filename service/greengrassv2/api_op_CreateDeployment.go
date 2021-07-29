// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a continuous deployment for a target, which is a AWS IoT Greengrass core
// device or group of core devices. When you add a new core device to a group of
// core devices that has a deployment, AWS IoT Greengrass deploys that group's
// deployment to the new device. You can define one deployment for each target.
// When you create a new deployment for a target that has an existing deployment,
// you replace the previous deployment. AWS IoT Greengrass applies the new
// deployment to the target devices. Every deployment has a revision number that
// indicates how many deployment revisions you define for a target. Use this
// operation to create a new revision of an existing deployment. This operation
// returns the revision number of the new deployment when you create it. For more
// information, see the Create deployments
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html)
// in the AWS IoT Greengrass V2 Developer Guide.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	if params == nil {
		params = &CreateDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeployment", params, optFns, c.addOperationCreateDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentInput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the target AWS IoT thing or thing group.
	//
	// This member is required.
	TargetArn *string

	// The components to deploy. This is a dictionary, where each key is the name of a
	// component, and each key's value is the version and configuration to deploy for
	// that component.
	Components map[string]types.ComponentDeploymentSpecification

	// The name of the deployment. You can create deployments without names. If you
	// create a deployment without a name, the AWS IoT Greengrass V2 console shows the
	// deployment name as :, where targetType and targetName are the type and name of
	// the deployment target.
	DeploymentName *string

	// The deployment policies for the deployment. These policies define how the
	// deployment updates components and handles failure.
	DeploymentPolicies *types.DeploymentPolicies

	// The job configuration for the deployment configuration. The job configuration
	// specifies the rollout, timeout, and stop configurations for the deployment
	// configuration.
	IotJobConfiguration *types.DeploymentIoTJobConfiguration

	// A list of key-value pairs that contain metadata for the resource. For more
	// information, see Tag your resources
	// (https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in
	// the AWS IoT Greengrass V2 Developer Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDeploymentOutput struct {

	// The ID of the deployment.
	DeploymentId *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the AWS IoT job that applies the deployment to target devices.
	IotJobArn *string

	// The ID of the AWS IoT job that applies the deployment to target devices.
	IotJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeployment{}, middleware.After)
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
	if err = addOpCreateDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateDeployment",
	}
}
