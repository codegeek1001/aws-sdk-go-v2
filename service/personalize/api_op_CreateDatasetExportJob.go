// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a job that exports data from your dataset to an Amazon S3 bucket. To
// allow Amazon Personalize to export the training data, you must specify an
// service-linked AWS Identity and Access Management (IAM) role that gives Amazon
// Personalize PutObject permissions for your Amazon S3 bucket. For information,
// see Exporting a dataset
// (https://docs.aws.amazon.com/personalize/latest/dg/export-data.html) in the
// Amazon Personalize developer guide. Status A dataset export job can be in one of
// the following states:
//
// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or-
// CREATE FAILED
//
// To get the status of the export job, call
// DescribeDatasetExportJob, and specify the Amazon Resource Name (ARN) of the
// dataset export job. The dataset export is complete when the status shows as
// ACTIVE. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed.
func (c *Client) CreateDatasetExportJob(ctx context.Context, params *CreateDatasetExportJobInput, optFns ...func(*Options)) (*CreateDatasetExportJobOutput, error) {
	if params == nil {
		params = &CreateDatasetExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatasetExportJob", params, optFns, c.addOperationCreateDatasetExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetExportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset that contains the data to export.
	//
	// This member is required.
	DatasetArn *string

	// The name for the dataset export job.
	//
	// This member is required.
	JobName *string

	// The path to the Amazon S3 bucket where the job's output is stored.
	//
	// This member is required.
	JobOutput *types.DatasetExportJobOutput

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management service
	// role that has permissions to add data to your output Amazon S3 bucket.
	//
	// This member is required.
	RoleArn *string

	// The data to export, based on how you imported the data. You can choose to export
	// only BULK data that you imported using a dataset import job, only PUT data that
	// you imported incrementally (using the console, PutEvents, PutUsers and PutItems
	// operations), or ALL for both types. The default value is PUT.
	IngestionMode types.IngestionMode

	noSmithyDocumentSerde
}

type CreateDatasetExportJobOutput struct {

	// The Amazon Resource Name (ARN) of the dataset export job.
	DatasetExportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatasetExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetExportJob{}, middleware.After)
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
	if err = addOpCreateDatasetExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatasetExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateDatasetExportJob",
	}
}
