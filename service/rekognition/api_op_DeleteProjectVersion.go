// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon Rekognition Custom Labels model. You can't delete a model if
// it is running or if it is training. To check the status of a model, use the
// Status field returned from DescribeProjectVersions (). To stop a running model
// call StopProjectVersion (). If the model is training, wait until it finishes.
// This operation requires permissions to perform the
// rekognition:DeleteProjectVersion action.
func (c *Client) DeleteProjectVersion(ctx context.Context, params *DeleteProjectVersionInput, optFns ...func(*Options)) (*DeleteProjectVersionOutput, error) {
	stack := middleware.NewStack("DeleteProjectVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteProjectVersionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteProjectVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProjectVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteProjectVersion",
			Err:           err,
		}
	}
	out := result.(*DeleteProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProjectVersionInput struct {

	// The Amazon Resource Name (ARN) of the model version that you want to delete.
	//
	// This member is required.
	ProjectVersionArn *string
}

type DeleteProjectVersionOutput struct {

	// The status of the deletion operation.
	Status types.ProjectVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteProjectVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProjectVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProjectVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteProjectVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DeleteProjectVersion",
	}
}