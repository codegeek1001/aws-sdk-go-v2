// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of AWS SSO identity store attributes that have been configured
// to work with attributes-based access control (ABAC) for the specified AWS SSO
// instance. This will not return attributes configured and sent by an external
// identity provider. For more information about ABAC, see Attribute-Based Access
// Control in the AWS SSO User Guide.
func (c *Client) DescribeInstanceAccessControlAttributeConfiguration(ctx context.Context, params *DescribeInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*DescribeInstanceAccessControlAttributeConfigurationOutput, error) {
	if params == nil {
		params = &DescribeInstanceAccessControlAttributeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceAccessControlAttributeConfiguration", params, optFns, c.addOperationDescribeInstanceAccessControlAttributeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceAccessControlAttributeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceAccessControlAttributeConfigurationInput struct {

	// The ARN of the SSO instance under which the operation will be executed.
	//
	// This member is required.
	InstanceArn *string

	noSmithyDocumentSerde
}

type DescribeInstanceAccessControlAttributeConfigurationOutput struct {

	// Gets the list of AWS SSO identity store attributes added to your ABAC
	// configuration.
	InstanceAccessControlAttributeConfiguration *types.InstanceAccessControlAttributeConfiguration

	// The status of the attribute configuration process.
	Status types.InstanceAccessControlAttributeConfigurationStatus

	// Provides more details about the current status of the specified attribute.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceAccessControlAttributeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstanceAccessControlAttributeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstanceAccessControlAttributeConfiguration{}, middleware.After)
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
	if err = addOpDescribeInstanceAccessControlAttributeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceAccessControlAttributeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceAccessControlAttributeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DescribeInstanceAccessControlAttributeConfiguration",
	}
}
