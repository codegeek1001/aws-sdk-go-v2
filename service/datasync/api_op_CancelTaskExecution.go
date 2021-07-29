// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels execution of a task. When you cancel a task execution, the transfer of
// some files is abruptly interrupted. The contents of files that are transferred
// to the destination might be incomplete or inconsistent with the source files.
// However, if you start a new task execution on the same task and you allow the
// task execution to complete, file content on the destination is complete and
// consistent. This applies to other unexpected failures that interrupt a task
// execution. In all of these cases, AWS DataSync successfully complete the
// transfer when you start the next task execution.
func (c *Client) CancelTaskExecution(ctx context.Context, params *CancelTaskExecutionInput, optFns ...func(*Options)) (*CancelTaskExecutionOutput, error) {
	if params == nil {
		params = &CancelTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelTaskExecution", params, optFns, c.addOperationCancelTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CancelTaskExecutionRequest
type CancelTaskExecutionInput struct {

	// The Amazon Resource Name (ARN) of the task execution to cancel.
	//
	// This member is required.
	TaskExecutionArn *string

	noSmithyDocumentSerde
}

type CancelTaskExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelTaskExecution{}, middleware.After)
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
	if err = addOpCancelTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelTaskExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CancelTaskExecution",
	}
}
