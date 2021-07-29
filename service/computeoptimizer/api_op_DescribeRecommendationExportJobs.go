// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes recommendation export jobs created in the last seven days. Use the
// ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations
// actions to request an export of your recommendations. Then use the
// DescribeRecommendationExportJobs action to view your export jobs.
func (c *Client) DescribeRecommendationExportJobs(ctx context.Context, params *DescribeRecommendationExportJobsInput, optFns ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error) {
	if params == nil {
		params = &DescribeRecommendationExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommendationExportJobs", params, optFns, c.addOperationDescribeRecommendationExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommendationExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationExportJobsInput struct {

	// An array of objects that describe a filter to return a more specific list of
	// export jobs.
	Filters []types.JobFilter

	// The identification numbers of the export jobs to return. An export job ID is
	// returned when you create an export using the
	// ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations
	// actions. All export jobs created in the last seven days are returned if this
	// parameter is omitted.
	JobIds []string

	// The maximum number of export jobs to return with a single request. To retrieve
	// the remaining results, make another request with the returned NextToken value.
	MaxResults *int32

	// The token to advance to the next page of export jobs.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRecommendationExportJobsOutput struct {

	// The token to use to advance to the next page of export jobs. This value is null
	// when there are no more pages of export jobs to return.
	NextToken *string

	// An array of objects that describe recommendation export jobs.
	RecommendationExportJobs []types.RecommendationExportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecommendationExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRecommendationExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRecommendationExportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "DescribeRecommendationExportJobs",
	}
}
