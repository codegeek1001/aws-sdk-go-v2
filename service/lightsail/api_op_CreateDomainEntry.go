// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one of the following domain name system (DNS) records in a domain DNS
// zone: Address (A), canonical name (CNAME), mail exchanger (MX), name server
// (NS), start of authority (SOA), service locator (SRV), or text (TXT). The create
// domain entry operation supports tag-based access control via resource tags
// applied to the resource identified by domain name. For more information, see the
// Amazon Lightsail Developer Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateDomainEntry(ctx context.Context, params *CreateDomainEntryInput, optFns ...func(*Options)) (*CreateDomainEntryOutput, error) {
	if params == nil {
		params = &CreateDomainEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainEntry", params, optFns, c.addOperationCreateDomainEntryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainEntryInput struct {

	// An array of key-value pairs containing information about the domain entry
	// request.
	//
	// This member is required.
	DomainEntry *types.DomainEntry

	// The domain name (e.g., example.com) for which you want to create the domain
	// entry.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type CreateDomainEntryOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainEntryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDomainEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDomainEntry{}, middleware.After)
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
	if err = addOpCreateDomainEntryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainEntry(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomainEntry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDomainEntry",
	}
}
