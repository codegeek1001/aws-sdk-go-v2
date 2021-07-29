// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an AWS Identity and Access Management (AWS IAM) policy for sharing the
// rule group. You can use the policy to share the rule group using AWS Resource
// Access Manager (AWS RAM).
func (c *Client) PutFirewallRuleGroupPolicy(ctx context.Context, params *PutFirewallRuleGroupPolicyInput, optFns ...func(*Options)) (*PutFirewallRuleGroupPolicyOutput, error) {
	if params == nil {
		params = &PutFirewallRuleGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutFirewallRuleGroupPolicy", params, optFns, c.addOperationPutFirewallRuleGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutFirewallRuleGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFirewallRuleGroupPolicyInput struct {

	// The ARN (Amazon Resource Name) for the rule group that you want to share.
	//
	// This member is required.
	Arn *string

	// The AWS Identity and Access Management (AWS IAM) policy to attach to the rule
	// group.
	//
	// This member is required.
	FirewallRuleGroupPolicy *string

	noSmithyDocumentSerde
}

type PutFirewallRuleGroupPolicyOutput struct {

	//
	ReturnValue bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutFirewallRuleGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutFirewallRuleGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutFirewallRuleGroupPolicy{}, middleware.After)
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
	if err = addOpPutFirewallRuleGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutFirewallRuleGroupPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutFirewallRuleGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "PutFirewallRuleGroupPolicy",
	}
}
