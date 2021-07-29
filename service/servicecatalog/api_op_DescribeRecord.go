// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the specified request operation. Use this operation after
// calling a request operation (for example, ProvisionProduct,
// TerminateProvisionedProduct, or UpdateProvisionedProduct). If a provisioned
// product was transferred to a new owner using UpdateProvisionedProductProperties,
// the new owner will be able to describe all past records for that product. The
// previous owner will no longer be able to describe the records, but will be able
// to use ListRecordHistory to see the product's history from when he was the
// owner.
func (c *Client) DescribeRecord(ctx context.Context, params *DescribeRecordInput, optFns ...func(*Options)) (*DescribeRecordOutput, error) {
	if params == nil {
		params = &DescribeRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecord", params, optFns, c.addOperationDescribeRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecordInput struct {

	// The record identifier of the provisioned product. This identifier is returned by
	// the request operation.
	//
	// This member is required.
	Id *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	noSmithyDocumentSerde
}

type DescribeRecordOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the product.
	RecordDetail *types.RecordDetail

	// Information about the product created as the result of a request. For example,
	// the output for a CloudFormation-backed product that creates an S3 bucket would
	// include the S3 bucket URL.
	RecordOutputs []types.RecordOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRecord{}, middleware.After)
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
	if err = addOpDescribeRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecord(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeRecord",
	}
}
