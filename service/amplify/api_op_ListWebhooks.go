// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of webhooks for an Amplify app.
func (c *Client) ListWebhooks(ctx context.Context, params *ListWebhooksInput, optFns ...func(*Options)) (*ListWebhooksOutput, error) {
	if params == nil {
		params = &ListWebhooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebhooks", params, optFns, c.addOperationListWebhooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebhooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list webhooks request.
type ListWebhooksInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The maximum number of records to list in a single response.
	MaxResults int32

	// A pagination token. Set to null to start listing webhooks from the start. If
	// non-null,the pagination token is returned in a result. Pass its value in here to
	// list more webhooks.
	NextToken *string

	noSmithyDocumentSerde
}

// The result structure for the list webhooks request.
type ListWebhooksOutput struct {

	// A list of webhooks.
	//
	// This member is required.
	Webhooks []types.Webhook

	// A pagination token. If non-null, the pagination token is returned in a result.
	// Pass its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWebhooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWebhooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebhooks{}, middleware.After)
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
	if err = addOpListWebhooksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebhooks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListWebhooks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListWebhooks",
	}
}
