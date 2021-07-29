// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops the specified pipeline execution. You choose to either stop the pipeline
// execution by completing in-progress actions without starting subsequent actions,
// or by abandoning in-progress actions. While completing or abandoning in-progress
// actions, the pipeline execution is in a Stopping state. After all in-progress
// actions are completed or abandoned, the pipeline execution is in a Stopped
// state.
func (c *Client) StopPipelineExecution(ctx context.Context, params *StopPipelineExecutionInput, optFns ...func(*Options)) (*StopPipelineExecutionOutput, error) {
	if params == nil {
		params = &StopPipelineExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopPipelineExecution", params, optFns, c.addOperationStopPipelineExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopPipelineExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopPipelineExecutionInput struct {

	// The ID of the pipeline execution to be stopped in the current stage. Use the
	// GetPipelineState action to retrieve the current pipelineExecutionId.
	//
	// This member is required.
	PipelineExecutionId *string

	// The name of the pipeline to stop.
	//
	// This member is required.
	PipelineName *string

	// Use this option to stop the pipeline execution by abandoning, rather than
	// finishing, in-progress actions. This option can lead to failed or
	// out-of-sequence tasks.
	Abandon bool

	// Use this option to enter comments, such as the reason the pipeline was stopped.
	Reason *string

	noSmithyDocumentSerde
}

type StopPipelineExecutionOutput struct {

	// The unique system-generated ID of the pipeline execution that was stopped.
	PipelineExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopPipelineExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopPipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopPipelineExecution{}, middleware.After)
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
	if err = addOpStopPipelineExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopPipelineExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopPipelineExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "StopPipelineExecution",
	}
}
