// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a notification rule for a resource. You can change the events that
// trigger the notification rule, the status of the rule, and the targets that
// receive the notifications. To add or remove tags for a notification rule, you
// must use TagResource and UntagResource.
func (c *Client) UpdateNotificationRule(ctx context.Context, params *UpdateNotificationRuleInput, optFns ...func(*Options)) (*UpdateNotificationRuleOutput, error) {
	if params == nil {
		params = &UpdateNotificationRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotificationRule", params, optFns, c.addOperationUpdateNotificationRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotificationRuleInput struct {

	// The Amazon Resource Name (ARN) of the notification rule.
	//
	// This member is required.
	Arn *string

	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in AWS
	// CloudWatch. FULL will include any supplemental information provided by AWS
	// CodeStar Notifications and/or the service for the resource for which the
	// notification is created.
	DetailType types.DetailType

	// A list of event types associated with this notification rule.
	EventTypeIds []string

	// The name of the notification rule.
	Name *string

	// The status of the notification rule. Valid statuses include enabled (sending
	// notifications) or disabled (not sending notifications).
	Status types.NotificationRuleStatus

	// The address and type of the targets to receive notifications from this
	// notification rule.
	Targets []types.Target

	noSmithyDocumentSerde
}

type UpdateNotificationRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNotificationRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNotificationRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNotificationRule{}, middleware.After)
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
	if err = addOpUpdateNotificationRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotificationRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNotificationRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "UpdateNotificationRule",
	}
}
