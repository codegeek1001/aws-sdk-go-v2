// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains information about the directories that belong to this account. You can
// retrieve information about specific directories by passing the directory
// identifiers in the DirectoryIds parameter. Otherwise, all directories that
// belong to the current account are returned. This operation supports pagination
// with the use of the NextToken request and response parameters. If more results
// are available, the DescribeDirectoriesResult.NextToken member contains a token
// that you pass in the next call to DescribeDirectories to retrieve the next set
// of items. You can also specify a maximum number of return results with the Limit
// parameter.
func (c *Client) DescribeDirectories(ctx context.Context, params *DescribeDirectoriesInput, optFns ...func(*Options)) (*DescribeDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectories", params, optFns, c.addOperationDescribeDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the DescribeDirectories operation.
type DescribeDirectoriesInput struct {

	// A list of identifiers of the directories for which to obtain the information. If
	// this member is null, all directories that belong to the current account are
	// returned. An empty list results in an InvalidParameterException being thrown.
	DirectoryIds []string

	// The maximum number of items to return. If this value is zero, the maximum number
	// of items is specified by the limitations of the operation.
	Limit *int32

	// The DescribeDirectoriesResult.NextToken value from a previous call to
	// DescribeDirectories. Pass null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains the results of the DescribeDirectories operation.
type DescribeDirectoriesOutput struct {

	// The list of DirectoryDescription objects that were retrieved. It is possible
	// that this list contains less than the number of items specified in the Limit
	// member of the request. This occurs if there are less than the requested number
	// of items left to retrieve, or if the limitations of the operation have been
	// exceeded.
	DirectoryDescriptions []types.DirectoryDescription

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to DescribeDirectories to retrieve the next set
	// of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectories{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeDirectories",
	}
}
