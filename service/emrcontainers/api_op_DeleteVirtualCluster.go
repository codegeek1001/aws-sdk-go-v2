// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a virtual cluster. Virtual cluster is a managed entity on Amazon EMR on
// EKS. You can create, describe, list and delete virtual clusters. They do not
// consume any additional resource in your system. A single virtual cluster maps to
// a single Kubernetes namespace. Given this relationship, you can model virtual
// clusters the same way you model Kubernetes namespaces to meet your requirements.
func (c *Client) DeleteVirtualCluster(ctx context.Context, params *DeleteVirtualClusterInput, optFns ...func(*Options)) (*DeleteVirtualClusterOutput, error) {
	if params == nil {
		params = &DeleteVirtualClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVirtualCluster", params, optFns, c.addOperationDeleteVirtualClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVirtualClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVirtualClusterInput struct {

	// The ID of the virtual cluster that will be deleted.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DeleteVirtualClusterOutput struct {

	// This output contains the ID of the virtual cluster that will be deleted.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVirtualClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVirtualCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVirtualCluster{}, middleware.After)
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
	if err = addOpDeleteVirtualClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVirtualCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVirtualCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-containers",
		OperationName: "DeleteVirtualCluster",
	}
}
