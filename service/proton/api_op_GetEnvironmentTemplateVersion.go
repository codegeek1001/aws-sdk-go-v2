// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// View detail data for a major or minor version of an environment template.
func (c *Client) GetEnvironmentTemplateVersion(ctx context.Context, params *GetEnvironmentTemplateVersionInput, optFns ...func(*Options)) (*GetEnvironmentTemplateVersionOutput, error) {
	if params == nil {
		params = &GetEnvironmentTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironmentTemplateVersion", params, optFns, c.addOperationGetEnvironmentTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentTemplateVersionInput struct {

	// To view environment template major version detail data, include majorVersion.
	//
	// This member is required.
	MajorVersion *string

	// To view environment template minor version detail data, include minorVersion.
	//
	// This member is required.
	MinorVersion *string

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

type GetEnvironmentTemplateVersionOutput struct {

	// The environment template version detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentTemplateVersion *types.EnvironmentTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnvironmentTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnvironmentTemplateVersion{}, middleware.After)
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
	if err = addOpGetEnvironmentTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnvironmentTemplateVersion(options.Region), middleware.Before); err != nil {
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

// GetEnvironmentTemplateVersionAPIClient is a client that implements the
// GetEnvironmentTemplateVersion operation.
type GetEnvironmentTemplateVersionAPIClient interface {
	GetEnvironmentTemplateVersion(context.Context, *GetEnvironmentTemplateVersionInput, ...func(*Options)) (*GetEnvironmentTemplateVersionOutput, error)
}

var _ GetEnvironmentTemplateVersionAPIClient = (*Client)(nil)

// EnvironmentTemplateVersionRegisteredWaiterOptions are waiter options for
// EnvironmentTemplateVersionRegisteredWaiter
type EnvironmentTemplateVersionRegisteredWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// EnvironmentTemplateVersionRegisteredWaiter will use default minimum delay of 2
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, EnvironmentTemplateVersionRegisteredWaiter will use default max delay
	// of 300 seconds. Note that MaxDelay must resolve to value greater than or equal
	// to the MinDelay.
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
	Retryable func(context.Context, *GetEnvironmentTemplateVersionInput, *GetEnvironmentTemplateVersionOutput, error) (bool, error)
}

// EnvironmentTemplateVersionRegisteredWaiter defines the waiters for
// EnvironmentTemplateVersionRegistered
type EnvironmentTemplateVersionRegisteredWaiter struct {
	client GetEnvironmentTemplateVersionAPIClient

	options EnvironmentTemplateVersionRegisteredWaiterOptions
}

// NewEnvironmentTemplateVersionRegisteredWaiter constructs a
// EnvironmentTemplateVersionRegisteredWaiter.
func NewEnvironmentTemplateVersionRegisteredWaiter(client GetEnvironmentTemplateVersionAPIClient, optFns ...func(*EnvironmentTemplateVersionRegisteredWaiterOptions)) *EnvironmentTemplateVersionRegisteredWaiter {
	options := EnvironmentTemplateVersionRegisteredWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 300 * time.Second
	options.Retryable = environmentTemplateVersionRegisteredStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EnvironmentTemplateVersionRegisteredWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EnvironmentTemplateVersionRegistered waiter.
// The maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur
// is required and must be greater than zero.
func (w *EnvironmentTemplateVersionRegisteredWaiter) Wait(ctx context.Context, params *GetEnvironmentTemplateVersionInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentTemplateVersionRegisteredWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 300 * time.Second
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

		out, err := w.client.GetEnvironmentTemplateVersion(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for EnvironmentTemplateVersionRegistered waiter")
}

func environmentTemplateVersionRegisteredStateRetryable(ctx context.Context, input *GetEnvironmentTemplateVersionInput, output *GetEnvironmentTemplateVersionOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("environmentTemplateVersion.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DRAFT"
		value, ok := pathValue.(types.TemplateVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TemplateVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("environmentTemplateVersion.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PUBLISHED"
		value, ok := pathValue.(types.TemplateVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TemplateVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("environmentTemplateVersion.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "REGISTRATION_FAILED"
		value, ok := pathValue.(types.TemplateVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TemplateVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetEnvironmentTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetEnvironmentTemplateVersion",
	}
}
