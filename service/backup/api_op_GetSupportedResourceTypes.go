// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the AWS resource types supported by AWS Backup.
func (c *Client) GetSupportedResourceTypes(ctx context.Context, params *GetSupportedResourceTypesInput, optFns ...func(*Options)) (*GetSupportedResourceTypesOutput, error) {
	if params == nil {
		params = &GetSupportedResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSupportedResourceTypes", params, optFns, c.addOperationGetSupportedResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSupportedResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSupportedResourceTypesInput struct {
	noSmithyDocumentSerde
}

type GetSupportedResourceTypesOutput struct {

	// Contains a string with the supported AWS resource types:
	//
	// * DynamoDB for Amazon
	// DynamoDB
	//
	// * EBS for Amazon Elastic Block Store
	//
	// * EC2 for Amazon Elastic Compute
	// Cloud
	//
	// * EFS for Amazon Elastic File System
	//
	// * RDS for Amazon Relational
	// Database Service
	//
	// * Aurora for Amazon Aurora
	//
	// * Storage Gateway for AWS Storage
	// Gateway
	ResourceTypes []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSupportedResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSupportedResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSupportedResourceTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSupportedResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSupportedResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "GetSupportedResourceTypes",
	}
}
