// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the scan findings for the specified image.
func (c *Client) DescribeImageScanFindings(ctx context.Context, params *DescribeImageScanFindingsInput, optFns ...func(*Options)) (*DescribeImageScanFindingsOutput, error) {
	if params == nil {
		params = &DescribeImageScanFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageScanFindings", params, optFns, c.addOperationDescribeImageScanFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageScanFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageScanFindingsInput struct {

	// An object with identifying information for an Amazon ECR image.
	//
	// This member is required.
	ImageId *types.ImageIdentifier

	// The repository for the image for which to describe the scan findings.
	//
	// This member is required.
	RepositoryName *string

	// The maximum number of image scan results returned by DescribeImageScanFindings
	// in paginated output. When this parameter is used, DescribeImageScanFindings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeImageScanFindings request with the returned nextToken value.
	// This value can be between 1 and 1000. If this parameter is not used, then
	// DescribeImageScanFindings returns up to 100 results and a nextToken value, if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeImageScanFindings
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string

	// The AWS account ID associated with the registry that contains the repository in
	// which to describe the image scan findings for. If you do not specify a registry,
	// the default registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type DescribeImageScanFindingsOutput struct {

	// An object with identifying information for an Amazon ECR image.
	ImageId *types.ImageIdentifier

	// The information contained in the image scan findings.
	ImageScanFindings *types.ImageScanFindings

	// The current state of the scan.
	ImageScanStatus *types.ImageScanStatus

	// The nextToken value to include in a future DescribeImageScanFindings request.
	// When the results of a DescribeImageScanFindings request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageScanFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageScanFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageScanFindings{}, middleware.After)
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
	if err = addOpDescribeImageScanFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageScanFindings(options.Region), middleware.Before); err != nil {
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

// DescribeImageScanFindingsAPIClient is a client that implements the
// DescribeImageScanFindings operation.
type DescribeImageScanFindingsAPIClient interface {
	DescribeImageScanFindings(context.Context, *DescribeImageScanFindingsInput, ...func(*Options)) (*DescribeImageScanFindingsOutput, error)
}

var _ DescribeImageScanFindingsAPIClient = (*Client)(nil)

// DescribeImageScanFindingsPaginatorOptions is the paginator options for
// DescribeImageScanFindings
type DescribeImageScanFindingsPaginatorOptions struct {
	// The maximum number of image scan results returned by DescribeImageScanFindings
	// in paginated output. When this parameter is used, DescribeImageScanFindings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeImageScanFindings request with the returned nextToken value.
	// This value can be between 1 and 1000. If this parameter is not used, then
	// DescribeImageScanFindings returns up to 100 results and a nextToken value, if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImageScanFindingsPaginator is a paginator for DescribeImageScanFindings
type DescribeImageScanFindingsPaginator struct {
	options   DescribeImageScanFindingsPaginatorOptions
	client    DescribeImageScanFindingsAPIClient
	params    *DescribeImageScanFindingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeImageScanFindingsPaginator returns a new
// DescribeImageScanFindingsPaginator
func NewDescribeImageScanFindingsPaginator(client DescribeImageScanFindingsAPIClient, params *DescribeImageScanFindingsInput, optFns ...func(*DescribeImageScanFindingsPaginatorOptions)) *DescribeImageScanFindingsPaginator {
	if params == nil {
		params = &DescribeImageScanFindingsInput{}
	}

	options := DescribeImageScanFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeImageScanFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImageScanFindingsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeImageScanFindings page.
func (p *DescribeImageScanFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImageScanFindingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeImageScanFindings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ImageScanCompleteWaiterOptions are waiter options for ImageScanCompleteWaiter
type ImageScanCompleteWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ImageScanCompleteWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ImageScanCompleteWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeImageScanFindingsInput, *DescribeImageScanFindingsOutput, error) (bool, error)
}

// ImageScanCompleteWaiter defines the waiters for ImageScanComplete
type ImageScanCompleteWaiter struct {
	client DescribeImageScanFindingsAPIClient

	options ImageScanCompleteWaiterOptions
}

// NewImageScanCompleteWaiter constructs a ImageScanCompleteWaiter.
func NewImageScanCompleteWaiter(client DescribeImageScanFindingsAPIClient, optFns ...func(*ImageScanCompleteWaiterOptions)) *ImageScanCompleteWaiter {
	options := ImageScanCompleteWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = imageScanCompleteStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ImageScanCompleteWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ImageScanComplete waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *ImageScanCompleteWaiter) Wait(ctx context.Context, params *DescribeImageScanFindingsInput, maxWaitDur time.Duration, optFns ...func(*ImageScanCompleteWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeImageScanFindings(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ImageScanComplete waiter")
}

func imageScanCompleteStateRetryable(ctx context.Context, input *DescribeImageScanFindingsInput, output *DescribeImageScanFindingsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("imageScanStatus.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETE"
		value, ok := pathValue.(types.ScanStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ScanStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("imageScanStatus.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.ScanStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ScanStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeImageScanFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DescribeImageScanFindings",
	}
}
