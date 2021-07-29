// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an artifact.
func (c *Client) DescribeArtifact(ctx context.Context, params *DescribeArtifactInput, optFns ...func(*Options)) (*DescribeArtifactOutput, error) {
	if params == nil {
		params = &DescribeArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeArtifact", params, optFns, c.addOperationDescribeArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeArtifactInput struct {

	// The Amazon Resource Name (ARN) of the artifact to describe.
	//
	// This member is required.
	ArtifactArn *string

	noSmithyDocumentSerde
}

type DescribeArtifactOutput struct {

	// The Amazon Resource Name (ARN) of the artifact.
	ArtifactArn *string

	// The name of the artifact.
	ArtifactName *string

	// The type of the artifact.
	ArtifactType *string

	// Information about the user who created or modified an experiment, trial, or
	// trial component.
	CreatedBy *types.UserContext

	// When the artifact was created.
	CreationTime *time.Time

	// Information about the user who created or modified an experiment, trial, or
	// trial component.
	LastModifiedBy *types.UserContext

	// When the artifact was last modified.
	LastModifiedTime *time.Time

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// A list of the artifact's properties.
	Properties map[string]string

	// The source of the artifact.
	Source *types.ArtifactSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeArtifact{}, middleware.After)
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
	if err = addOpDescribeArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeArtifact",
	}
}
