// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify one of your existing queues.
func (c *Client) UpdateQueue(ctx context.Context, params *UpdateQueueInput, optFns ...func(*Options)) (*UpdateQueueOutput, error) {
	if params == nil {
		params = &UpdateQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQueue", params, optFns, c.addOperationUpdateQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQueueInput struct {

	// The name of the queue that you are modifying.
	//
	// This member is required.
	Name *string

	// The new description for the queue, if you are changing it.
	Description *string

	// The new details of your pricing plan for your reserved queue. When you set up a
	// new pricing plan to replace an expired one, you enter into another 12-month
	// commitment. When you add capacity to your queue by increasing the number of RTS,
	// you extend the term of your commitment to 12 months from when you add capacity.
	// After you make these commitments, you can't cancel them.
	ReservationPlanSettings *types.ReservationPlanSettings

	// Pause or activate a queue by changing its status between ACTIVE and PAUSED. If
	// you pause a queue, jobs in that queue won't begin. Jobs that are running when
	// you pause the queue continue to run until they finish or result in an error.
	Status types.QueueStatus

	noSmithyDocumentSerde
}

type UpdateQueueOutput struct {

	// You can use queues to manage the resources that are available to your AWS
	// account for running multiple transcoding jobs at the same time. If you don't
	// specify a queue, the service sends all jobs through the default queue. For more
	// information, see
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html.
	Queue *types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQueue{}, middleware.After)
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
	if err = addOpUpdateQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "UpdateQueue",
	}
}
