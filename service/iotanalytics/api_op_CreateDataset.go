// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a dataset. A dataset stores data retrieved from a data store by applying
// a queryAction (a SQL query) or a containerAction (executing a containerized
// application). This operation creates the skeleton of a dataset. The dataset can
// be populated manually by calling CreateDatasetContent or automatically according
// to a trigger you specify.
func (c *Client) CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) {
	if params == nil {
		params = &CreateDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataset", params, optFns, c.addOperationCreateDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetInput struct {

	// A list of actions that create the data set contents.
	//
	// This member is required.
	Actions []types.DatasetAction

	// The name of the data set.
	//
	// This member is required.
	DatasetName *string

	// When dataset contents are created, they are delivered to destinations specified
	// here.
	ContentDeliveryRules []types.DatasetContentDeliveryRule

	// A list of data rules that send notifications to Amazon CloudWatch, when data
	// arrives late. To specify lateDataRules, the dataset must use a DeltaTimer
	// (https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_DeltaTime.html)
	// filter.
	LateDataRules []types.LateDataRule

	// Optional. How long, in days, versions of dataset contents are kept for the
	// dataset. If not specified or set to null, versions of dataset contents are
	// retained for at most 90 days. The number of versions of dataset contents
	// retained is determined by the versioningConfiguration parameter. For more
	// information, see Keeping Multiple Versions of AWS IoT Analytics Data Sets
	// (https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	// in the AWS IoT Analytics User Guide.
	RetentionPeriod *types.RetentionPeriod

	// Metadata which can be used to manage the data set.
	Tags []types.Tag

	// A list of triggers. A trigger causes data set contents to be populated at a
	// specified time interval or when another data set's contents are created. The
	// list of triggers can be empty or contain up to five DataSetTrigger objects.
	Triggers []types.DatasetTrigger

	// Optional. How many versions of dataset contents are kept. If not specified or
	// set to null, only the latest version plus the latest succeeded version (if they
	// are different) are kept for the time period specified by the retentionPeriod
	// parameter. For more information, see Keeping Multiple Versions of AWS IoT
	// Analytics Data Sets
	// (https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	// in the AWS IoT Analytics User Guide.
	VersioningConfiguration *types.VersioningConfiguration

	noSmithyDocumentSerde
}

type CreateDatasetOutput struct {

	// The ARN of the dataset.
	DatasetArn *string

	// The name of the dataset.
	DatasetName *string

	// How long, in days, dataset contents are kept for the dataset.
	RetentionPeriod *types.RetentionPeriod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataset{}, middleware.After)
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
	if err = addOpCreateDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CreateDataset",
	}
}
