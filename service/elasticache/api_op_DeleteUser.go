// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For Redis engine version 6.x onwards: Deletes a user. The user will be removed
// from all user groups and in turn removed from all replication groups. For more
// information, see Using Role Based Access Control (RBAC)
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html).
func (c *Client) DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) {
	if params == nil {
		params = &DeleteUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUser", params, optFns, c.addOperationDeleteUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserInput struct {

	// The ID of the user.
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type DeleteUserOutput struct {

	// The Amazon Resource Name (ARN) of the user.
	ARN *string

	// Access permissions string used for this user.
	AccessString *string

	// Denotes whether the user requires a password to authenticate.
	Authentication *types.Authentication

	// The current supported value is Redis.
	Engine *string

	// Indicates the user status. Can be "active", "modifying" or "deleting".
	Status *string

	// Returns a list of the user group IDs the user belongs to.
	UserGroupIds []string

	// The ID of the user.
	UserId *string

	// The username of the user.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteUser{}, middleware.After)
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
	if err = addOpDeleteUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteUser",
	}
}
