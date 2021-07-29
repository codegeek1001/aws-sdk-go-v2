// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the current status of an asynchronous request to create an account.
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator for an AWS service.
func (c *Client) DescribeCreateAccountStatus(ctx context.Context, params *DescribeCreateAccountStatusInput, optFns ...func(*Options)) (*DescribeCreateAccountStatusOutput, error) {
	if params == nil {
		params = &DescribeCreateAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCreateAccountStatus", params, optFns, c.addOperationDescribeCreateAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCreateAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCreateAccountStatusInput struct {

	// Specifies the Id value that uniquely identifies the CreateAccount request. You
	// can get the value from the CreateAccountStatus.Id response in an earlier
	// CreateAccount request, or from the ListCreateAccountStatus operation. The regex
	// pattern (http://wikipedia.org/wiki/regex) for a create account request ID string
	// requires "car-" followed by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	CreateAccountRequestId *string

	noSmithyDocumentSerde
}

type DescribeCreateAccountStatusOutput struct {

	// A structure that contains the current status of an account creation request.
	CreateAccountStatus *types.CreateAccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCreateAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCreateAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCreateAccountStatus{}, middleware.After)
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
	if err = addOpDescribeCreateAccountStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCreateAccountStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCreateAccountStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeCreateAccountStatus",
	}
}
