// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the mobile device access rules for the specified Amazon WorkMail
// organization.
func (c *Client) ListMobileDeviceAccessRules(ctx context.Context, params *ListMobileDeviceAccessRulesInput, optFns ...func(*Options)) (*ListMobileDeviceAccessRulesOutput, error) {
	if params == nil {
		params = &ListMobileDeviceAccessRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMobileDeviceAccessRules", params, optFns, c.addOperationListMobileDeviceAccessRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMobileDeviceAccessRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMobileDeviceAccessRulesInput struct {

	// The Amazon WorkMail organization for which to list the rules.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type ListMobileDeviceAccessRulesOutput struct {

	// The list of mobile device access rules that exist under the specified Amazon
	// WorkMail organization.
	Rules []types.MobileDeviceAccessRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMobileDeviceAccessRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMobileDeviceAccessRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMobileDeviceAccessRules{}, middleware.After)
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
	if err = addOpListMobileDeviceAccessRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMobileDeviceAccessRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListMobileDeviceAccessRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListMobileDeviceAccessRules",
	}
}
