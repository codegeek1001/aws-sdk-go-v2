// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes one or more import tasks, each identified by their import ID. Each
// import task has a number of records that can identify servers or applications.
// AWS Application Discovery Service has built-in matching logic that will identify
// when discovered servers match existing entries that you've previously
// discovered, the information for the already-existing discovered server is
// updated. When you delete an import task that contains records that were used to
// match, the information in those matched records that comes from the deleted
// records will also be deleted.
func (c *Client) BatchDeleteImportData(ctx context.Context, params *BatchDeleteImportDataInput, optFns ...func(*Options)) (*BatchDeleteImportDataOutput, error) {
	if params == nil {
		params = &BatchDeleteImportDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteImportData", params, optFns, c.addOperationBatchDeleteImportDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteImportDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteImportDataInput struct {

	// The IDs for the import tasks that you want to delete.
	//
	// This member is required.
	ImportTaskIds []string

	noSmithyDocumentSerde
}

type BatchDeleteImportDataOutput struct {

	// Error messages returned for each import task that you deleted as a response for
	// this command.
	Errors []types.BatchDeleteImportDataError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteImportDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteImportData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteImportData{}, middleware.After)
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
	if err = addOpBatchDeleteImportDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteImportData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteImportData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "BatchDeleteImportData",
	}
}
