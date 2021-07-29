// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a member association in Security Hub between the specified accounts and
// the account used to make the request, which is the administrator account. If you
// are integrated with Organizations, then the administrator account is designated
// by the organization management account. CreateMembers is always used to add
// accounts that are not organization members. For accounts that are part of an
// organization, CreateMembers is only used in the following cases:
//
// * Security Hub
// is not configured to automatically add new accounts in an organization.
//
// * The
// account was disassociated or deleted in Security Hub.
//
// This action can only be
// used by an account that has Security Hub enabled. To enable Security Hub, you
// can use the EnableSecurityHub operation. For accounts that are not organization
// members, you create the account association and then send an invitation to the
// member account. To send the invitation, you use the InviteMembers operation. If
// the account owner accepts the invitation, the account becomes a member account
// in Security Hub. Accounts that are part of an organization do not receive an
// invitation. They automatically become a member account in Security Hub. A
// permissions policy is added that permits the administrator account to view the
// findings generated in the member account. When Security Hub is enabled in a
// member account, the member account findings are also visible to the
// administrator account. To remove the association between the administrator and
// member accounts, use the DisassociateFromMasterAccount or DisassociateMembers
// operation.
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	if params == nil {
		params = &CreateMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembers", params, optFns, c.addOperationCreateMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// The list of accounts to associate with the Security Hub administrator account.
	// For each account, the list includes the account ID and optionally the email
	// address.
	//
	// This member is required.
	AccountDetails []types.AccountDetails

	noSmithyDocumentSerde
}

type CreateMembersOutput struct {

	// The list of AWS accounts that were not processed. For each account, the list
	// includes the account ID and the email address.
	UnprocessedAccounts []types.Result

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
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
	if err = addOpCreateMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "CreateMembers",
	}
}
