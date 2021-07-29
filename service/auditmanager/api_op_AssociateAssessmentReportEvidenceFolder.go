// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an evidence folder to the specified assessment report in AWS Audit
// Manager.
func (c *Client) AssociateAssessmentReportEvidenceFolder(ctx context.Context, params *AssociateAssessmentReportEvidenceFolderInput, optFns ...func(*Options)) (*AssociateAssessmentReportEvidenceFolderOutput, error) {
	if params == nil {
		params = &AssociateAssessmentReportEvidenceFolderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAssessmentReportEvidenceFolder", params, optFns, c.addOperationAssociateAssessmentReportEvidenceFolderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAssessmentReportEvidenceFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAssessmentReportEvidenceFolderInput struct {

	// The identifier for the specified assessment.
	//
	// This member is required.
	AssessmentId *string

	// The identifier for the folder in which evidence is stored.
	//
	// This member is required.
	EvidenceFolderId *string

	noSmithyDocumentSerde
}

type AssociateAssessmentReportEvidenceFolderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateAssessmentReportEvidenceFolderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateAssessmentReportEvidenceFolder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateAssessmentReportEvidenceFolder{}, middleware.After)
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
	if err = addOpAssociateAssessmentReportEvidenceFolderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAssessmentReportEvidenceFolder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateAssessmentReportEvidenceFolder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "AssociateAssessmentReportEvidenceFolder",
	}
}
