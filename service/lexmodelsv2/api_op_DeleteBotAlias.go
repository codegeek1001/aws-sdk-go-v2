// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified bot alias.
func (c *Client) DeleteBotAlias(ctx context.Context, params *DeleteBotAliasInput, optFns ...func(*Options)) (*DeleteBotAliasOutput, error) {
	if params == nil {
		params = &DeleteBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBotAlias", params, optFns, c.addOperationDeleteBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotAliasInput struct {

	// The unique identifier of the bot alias to delete.
	//
	// This member is required.
	BotAliasId *string

	// The unique identifier of the bot associated with the alias to delete.
	//
	// This member is required.
	BotId *string

	// When this parameter is true, Amazon Lex doesn't check to see if any other
	// resource is using the alias before it is deleted.
	SkipResourceInUseCheck bool

	noSmithyDocumentSerde
}

type DeleteBotAliasOutput struct {

	// The unique identifier of the bot alias to delete.
	BotAliasId *string

	// The current status of the alias. The status is Deleting while the alias is in
	// the process of being deleted. Once the alias is deleted, it will no longer
	// appear in the list of aliases returned by the ListBotAliases operation.
	BotAliasStatus types.BotAliasStatus

	// The unique identifier of the bot that contains the alias to delete.
	BotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBotAlias{}, middleware.After)
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
	if err = addOpDeleteBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBotAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBotAlias",
	}
}
