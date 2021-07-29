// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an alarm model. Any alarms that were created based on the previous
// version are deleted and then created again as new data arrives.
func (c *Client) UpdateAlarmModel(ctx context.Context, params *UpdateAlarmModelInput, optFns ...func(*Options)) (*UpdateAlarmModelOutput, error) {
	if params == nil {
		params = &UpdateAlarmModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAlarmModel", params, optFns, c.addOperationUpdateAlarmModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAlarmModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAlarmModelInput struct {

	// The name of the alarm model.
	//
	// This member is required.
	AlarmModelName *string

	// Defines when your alarm is invoked.
	//
	// This member is required.
	AlarmRule *types.AlarmRule

	// The ARN of the IAM role that allows the alarm to perform actions and access AWS
	// resources. For more information, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	RoleArn *string

	// Contains the configuration information of alarm state changes.
	AlarmCapabilities *types.AlarmCapabilities

	// Contains information about one or more alarm actions.
	AlarmEventActions *types.AlarmEventActions

	// The description of the alarm model.
	AlarmModelDescription *string

	// Contains information about one or more notification actions.
	AlarmNotification *types.AlarmNotification

	// A non-negative integer that reflects the severity level of the alarm.
	Severity *int32

	noSmithyDocumentSerde
}

type UpdateAlarmModelOutput struct {

	// The ARN of the alarm model. For more information, see Amazon Resource Names
	// (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	AlarmModelArn *string

	// The version of the alarm model.
	AlarmModelVersion *string

	// The time the alarm model was created, in the Unix epoch format.
	CreationTime *time.Time

	// The time the alarm model was last updated, in the Unix epoch format.
	LastUpdateTime *time.Time

	// The status of the alarm model. The status can be one of the following values:
	//
	// *
	// ACTIVE - The alarm model is active and it's ready to evaluate data.
	//
	// *
	// ACTIVATING - AWS IoT Events is activating your alarm model. Activating an alarm
	// model can take up to a few minutes.
	//
	// * INACTIVE - The alarm model is inactive,
	// so it isn't ready to evaluate data. Check your alarm model information and
	// update the alarm model.
	//
	// * FAILED - You couldn't create or update the alarm
	// model. Check your alarm model information and try again.
	Status types.AlarmModelVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAlarmModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAlarmModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAlarmModel{}, middleware.After)
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
	if err = addOpUpdateAlarmModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlarmModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAlarmModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "UpdateAlarmModel",
	}
}
