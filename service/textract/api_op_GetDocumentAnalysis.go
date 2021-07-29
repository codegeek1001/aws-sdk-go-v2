// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a document. You start asynchronous text analysis by calling
// StartDocumentAnalysis, which returns a job identifier (JobId). When the text
// analysis operation finishes, Amazon Textract publishes a completion status to
// the Amazon Simple Notification Service (Amazon SNS) topic that's registered in
// the initial call to StartDocumentAnalysis. To get the results of the
// text-detection operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If so, call GetDocumentAnalysis, and pass the job
// identifier (JobId) from the initial call to StartDocumentAnalysis.
// GetDocumentAnalysis returns an array of Block objects. The following types of
// information are returned:
//
// * Form data (key-value pairs). The related
// information is returned in two Block objects, each of type KEY_VALUE_SET: a KEY
// Block object and a VALUE Block object. For example, Name: Ana Silva Carolina
// contains a key and value. Name: is the key. Ana Silva Carolina is the value.
//
// *
// Table and table cell data. A TABLE Block object contains information about a
// detected table. A CELL Block object is returned for each cell in a table.
//
// *
// Lines and words of text. A LINE Block object contains one or more WORD Block
// objects. All lines and words that are detected in the document are returned
// (including text that doesn't have a relationship with the value of the
// StartDocumentAnalysisFeatureTypes input parameter).
//
// Selection elements such as
// check boxes and option buttons (radio buttons) can be detected in form data and
// in tables. A SELECTION_ELEMENT Block object contains information about a
// selection element, including the selection status. Use the MaxResults parameter
// to limit the number of blocks that are returned. If there are more results than
// specified in MaxResults, the value of NextToken in the operation response
// contains a pagination token for getting the next set of results. To get the next
// page of results, call GetDocumentAnalysis, and populate the NextToken request
// parameter with the token value that's returned from the previous call to
// GetDocumentAnalysis. For more information, see Document Text Analysis
// (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html).
func (c *Client) GetDocumentAnalysis(ctx context.Context, params *GetDocumentAnalysisInput, optFns ...func(*Options)) (*GetDocumentAnalysisOutput, error) {
	if params == nil {
		params = &GetDocumentAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDocumentAnalysis", params, optFns, c.addOperationGetDocumentAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDocumentAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDocumentAnalysisInput struct {

	// A unique identifier for the text-detection job. The JobId is returned from
	// StartDocumentAnalysis. A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	// The maximum number of results to return per paginated call. The largest value
	// that you can specify is 1,000. If you specify a value greater than 1,000, a
	// maximum of 1,000 results is returned. The default value is 1,000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more blocks to
	// retrieve), Amazon Textract returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of blocks.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDocumentAnalysisOutput struct {

	//
	AnalyzeDocumentModelVersion *string

	// The results of the text-analysis operation.
	Blocks []types.Block

	// Information about a document that Amazon Textract processed. DocumentMetadata is
	// returned in every page of paginated responses from an Amazon Textract video
	// operation.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the text detection job.
	JobStatus types.JobStatus

	// If the response is truncated, Amazon Textract returns this token. You can use
	// this token in the subsequent request to retrieve the next set of text detection
	// results.
	NextToken *string

	// Returns if the detection job could not be completed. Contains explanation for
	// what error occured.
	StatusMessage *string

	// A list of warnings that occurred during the document-analysis operation.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDocumentAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDocumentAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDocumentAnalysis{}, middleware.After)
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
	if err = addOpGetDocumentAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDocumentAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDocumentAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetDocumentAnalysis",
	}
}
