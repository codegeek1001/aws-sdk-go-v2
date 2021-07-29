// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new parameter group. A parameter group is a collection of parameters
// that you apply to all of the nodes in a DAX cluster.
func (c *Client) CreateParameterGroup(ctx context.Context, params *CreateParameterGroupInput, optFns ...func(*Options)) (*CreateParameterGroupOutput, error) {
	if params == nil {
		params = &CreateParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateParameterGroup", params, optFns, c.addOperationCreateParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateParameterGroupInput struct {

	// The name of the parameter group to apply to all of the clusters in this
	// replication group.
	//
	// This member is required.
	ParameterGroupName *string

	// A description of the parameter group.
	Description *string

	noSmithyDocumentSerde
}

type CreateParameterGroupOutput struct {

	// Represents the output of a CreateParameterGroup action.
	ParameterGroup *types.ParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateParameterGroup{}, middleware.After)
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
	if err = addOpCreateParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "CreateParameterGroup",
	}
}
