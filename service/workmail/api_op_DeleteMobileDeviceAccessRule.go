// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a mobile device access rule for the specified Amazon WorkMail
// organization.
func (c *Client) DeleteMobileDeviceAccessRule(ctx context.Context, params *DeleteMobileDeviceAccessRuleInput, optFns ...func(*Options)) (*DeleteMobileDeviceAccessRuleOutput, error) {
	if params == nil {
		params = &DeleteMobileDeviceAccessRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMobileDeviceAccessRule", params, optFns, c.addOperationDeleteMobileDeviceAccessRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMobileDeviceAccessRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMobileDeviceAccessRuleInput struct {

	// The identifier of the rule to be deleted.
	//
	// This member is required.
	MobileDeviceAccessRuleId *string

	// The Amazon WorkMail organization under which the rule will be deleted.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type DeleteMobileDeviceAccessRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMobileDeviceAccessRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMobileDeviceAccessRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMobileDeviceAccessRule{}, middleware.After)
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
	if err = addOpDeleteMobileDeviceAccessRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMobileDeviceAccessRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMobileDeviceAccessRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DeleteMobileDeviceAccessRule",
	}
}
