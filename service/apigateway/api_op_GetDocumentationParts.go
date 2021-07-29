// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) GetDocumentationParts(ctx context.Context, params *GetDocumentationPartsInput, optFns ...func(*Options)) (*GetDocumentationPartsOutput, error) {
	if params == nil {
		params = &GetDocumentationPartsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentationParts", params, optFns, c.addOperationGetDocumentationPartsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentationPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the documentation parts of an API. The result may be filtered by the type,
// name, or path of API entities (targets).
type GetDocumentationPartsInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The status of the API documentation parts to retrieve. Valid values are
	// DOCUMENTED for retrieving DocumentationPart resources with content and
	// UNDOCUMENTED for DocumentationPart resources without content.
	LocationStatus types.LocationStatusType

	// The name of API entities of the to-be-retrieved documentation parts.
	NameQuery *string

	// The path of API entities of the to-be-retrieved documentation parts.
	Path *string

	// The current pagination position in the paged result set.
	Position *string

	// The type of API entities of the to-be-retrieved documentation parts.
	Type types.DocumentationPartType

	noSmithyDocumentSerde
}

// The collection of documentation parts of an API. Documenting an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api.html),
// DocumentationPart
type GetDocumentationPartsOutput struct {

	// The current page of elements from this collection.
	Items []types.DocumentationPart

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDocumentationPartsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDocumentationParts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDocumentationParts{}, middleware.After)
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
	if err = addOpGetDocumentationPartsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentationParts(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetDocumentationParts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDocumentationParts",
	}
}
