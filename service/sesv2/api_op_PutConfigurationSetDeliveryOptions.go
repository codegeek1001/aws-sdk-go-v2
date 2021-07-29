// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate a configuration set with a dedicated IP pool. You can use dedicated IP
// pools to create groups of dedicated IP addresses for sending specific types of
// email.
func (c *Client) PutConfigurationSetDeliveryOptions(ctx context.Context, params *PutConfigurationSetDeliveryOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetDeliveryOptionsOutput, error) {
	if params == nil {
		params = &PutConfigurationSetDeliveryOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationSetDeliveryOptions", params, optFns, c.addOperationPutConfigurationSetDeliveryOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationSetDeliveryOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to associate a configuration set with a dedicated IP pool.
type PutConfigurationSetDeliveryOptionsInput struct {

	// The name of the configuration set that you want to associate with a dedicated IP
	// pool.
	//
	// This member is required.
	ConfigurationSetName *string

	// The name of the dedicated IP pool that you want to associate with the
	// configuration set.
	SendingPoolName *string

	// Specifies whether messages that use the configuration set are required to use
	// Transport Layer Security (TLS). If the value is Require, messages are only
	// delivered if a TLS connection can be established. If the value is Optional,
	// messages can be delivered in plain text if a TLS connection can't be
	// established.
	TlsPolicy types.TlsPolicy

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetDeliveryOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConfigurationSetDeliveryOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetDeliveryOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetDeliveryOptions{}, middleware.After)
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
	if err = addOpPutConfigurationSetDeliveryOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetDeliveryOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationSetDeliveryOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetDeliveryOptions",
	}
}
