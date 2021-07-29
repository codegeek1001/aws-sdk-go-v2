// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an approval rule from a specified pull request. Approval rules can be
// deleted from a pull request only if the pull request is open, and if the
// approval rule was created specifically for a pull request and not generated from
// an approval rule template associated with the repository where the pull request
// was created. You cannot delete an approval rule from a merged or closed pull
// request.
func (c *Client) DeletePullRequestApprovalRule(ctx context.Context, params *DeletePullRequestApprovalRuleInput, optFns ...func(*Options)) (*DeletePullRequestApprovalRuleOutput, error) {
	if params == nil {
		params = &DeletePullRequestApprovalRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePullRequestApprovalRule", params, optFns, c.addOperationDeletePullRequestApprovalRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePullRequestApprovalRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePullRequestApprovalRuleInput struct {

	// The name of the approval rule you want to delete.
	//
	// This member is required.
	ApprovalRuleName *string

	// The system-generated ID of the pull request that contains the approval rule you
	// want to delete.
	//
	// This member is required.
	PullRequestId *string

	noSmithyDocumentSerde
}

type DeletePullRequestApprovalRuleOutput struct {

	// The ID of the deleted approval rule. If the approval rule was deleted in an
	// earlier API call, the response is 200 OK without content.
	//
	// This member is required.
	ApprovalRuleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePullRequestApprovalRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePullRequestApprovalRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePullRequestApprovalRule{}, middleware.After)
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
	if err = addOpDeletePullRequestApprovalRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePullRequestApprovalRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePullRequestApprovalRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DeletePullRequestApprovalRule",
	}
}
