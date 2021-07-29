// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a constraint. A delegated admin is authorized to invoke this command.
func (c *Client) CreateConstraint(ctx context.Context, params *CreateConstraintInput, optFns ...func(*Options)) (*CreateConstraintOutput, error) {
	if params == nil {
		params = &CreateConstraintInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConstraint", params, optFns, c.addOperationCreateConstraintMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConstraintOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConstraintInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The constraint parameters, in JSON format. The syntax depends on the constraint
	// type as follows: LAUNCH You are required to specify either the RoleArn or the
	// LocalRoleName but can't use both. Specify the RoleArn property as follows:
	// {"RoleArn" : "arn:aws:iam::123456789012:role/LaunchRole"} Specify the
	// LocalRoleName property as follows: {"LocalRoleName": "SCBasicLaunchRole"} If you
	// specify the LocalRoleName property, when an account uses the launch constraint,
	// the IAM role with that name in the account will be used. This allows launch-role
	// constraints to be account-agnostic so the administrator can create fewer
	// resources per shared account. The given role name must exist in the account used
	// to create the launch constraint and the account of the user who launches a
	// product with this launch constraint. You cannot have both a LAUNCH and a
	// STACKSET constraint. You also cannot have more than one LAUNCH constraint on a
	// product and portfolio. NOTIFICATION Specify the NotificationArns property as
	// follows: {"NotificationArns" : ["arn:aws:sns:us-east-1:123456789012:Topic"]}
	// RESOURCE_UPDATE Specify the TagUpdatesOnProvisionedProduct property as follows:
	// {"Version":"2.0","Properties":{"TagUpdateOnProvisionedProduct":"String"}} The
	// TagUpdatesOnProvisionedProduct property accepts a string value of ALLOWED or
	// NOT_ALLOWED. STACKSET Specify the Parameters property as follows: {"Version":
	// "String", "Properties": {"AccountList": [ "String" ], "RegionList": [ "String"
	// ], "AdminRole": "String", "ExecutionRole": "String"}} You cannot have both a
	// LAUNCH and a STACKSET constraint. You also cannot have more than one STACKSET
	// constraint on a product and portfolio. Products with a STACKSET constraint will
	// launch an AWS CloudFormation stack set. TEMPLATE Specify the Rules property. For
	// more information, see Template Constraint Rules
	// (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html).
	//
	// This member is required.
	Parameters *string

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The product identifier.
	//
	// This member is required.
	ProductId *string

	// The type of constraint.
	//
	// * LAUNCH
	//
	// * NOTIFICATION
	//
	// * RESOURCE_UPDATE
	//
	// *
	// STACKSET
	//
	// * TEMPLATE
	//
	// This member is required.
	Type *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The description of the constraint.
	Description *string

	noSmithyDocumentSerde
}

type CreateConstraintOutput struct {

	// Information about the constraint.
	ConstraintDetail *types.ConstraintDetail

	// The constraint parameters.
	ConstraintParameters *string

	// The status of the current request.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConstraintMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConstraint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConstraint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateConstraintMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateConstraintValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConstraint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateConstraint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConstraint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConstraint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConstraintInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConstraintInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateConstraintMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateConstraint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConstraint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateConstraint",
	}
}
