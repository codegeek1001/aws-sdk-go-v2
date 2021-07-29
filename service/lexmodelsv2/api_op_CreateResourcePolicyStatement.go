// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a new resource policy statement to a bot or bot alias. If a resource policy
// exists, the statement is added to the current resource policy. If a policy
// doesn't exist, a new policy is created. You can't create a resource policy
// statement that allows cross-account access.
func (c *Client) CreateResourcePolicyStatement(ctx context.Context, params *CreateResourcePolicyStatementInput, optFns ...func(*Options)) (*CreateResourcePolicyStatementOutput, error) {
	if params == nil {
		params = &CreateResourcePolicyStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourcePolicyStatement", params, optFns, c.addOperationCreateResourcePolicyStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourcePolicyStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourcePolicyStatementInput struct {

	// The Amazon Lex action that this policy either allows or denies. The action must
	// apply to the resource type of the specified ARN. For more information, see
	// Actions, resources, and condition keys for Amazon Lex V2
	// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonlexv2.html).
	//
	// This member is required.
	Action []string

	// Determines whether the statement allows or denies access to the resource.
	//
	// This member is required.
	Effect types.Effect

	// An IAM principal, such as an IAM users, IAM roles, or AWS services that is
	// allowed or denied access to a resource. For more information, see AWS JSON
	// policy elements: Principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html).
	//
	// This member is required.
	Principal []types.Principal

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	//
	// This member is required.
	ResourceArn *string

	// The name of the statement. The ID is the same as the Sid IAM property. The
	// statement name must be unique within the policy. For more information, see IAM
	// JSON policy elements: Sid
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_sid.html).
	//
	// This member is required.
	StatementId *string

	// Specifies a condition when the policy is in effect. If the principal of the
	// policy is a service principal, you must provide two condition blocks, one with a
	// SourceAccount global condition key and one with a SourceArn global condition
	// key. For more information, see IAM JSON policy elements: Condition
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html).
	Condition map[string]map[string]string

	// The identifier of the revision of the policy to edit. If this revision ID
	// doesn't match the current revision ID, Amazon Lex throws an exception. If you
	// don't specify a revision, Amazon Lex overwrites the contents of the policy with
	// the new values.
	ExpectedRevisionId *string

	noSmithyDocumentSerde
}

type CreateResourcePolicyStatementOutput struct {

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	ResourceArn *string

	// The current revision of the resource policy. Use the revision ID to make sure
	// that you are updating the most current version of a resource policy when you add
	// a policy statement to a resource, delete a resource, or update a resource.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourcePolicyStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourcePolicyStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourcePolicyStatement{}, middleware.After)
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
	if err = addOpCreateResourcePolicyStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourcePolicyStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourcePolicyStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateResourcePolicyStatement",
	}
}
