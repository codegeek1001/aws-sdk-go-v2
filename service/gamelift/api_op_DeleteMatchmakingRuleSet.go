// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing matchmaking rule set. To delete the rule set, provide the
// rule set name. Rule sets cannot be deleted if they are currently being used by a
// matchmaking configuration. Learn more
//
// * Build a rule set
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
//
// Related
// actions CreateMatchmakingConfiguration | DescribeMatchmakingConfigurations |
// UpdateMatchmakingConfiguration | DeleteMatchmakingConfiguration |
// CreateMatchmakingRuleSet | DescribeMatchmakingRuleSets |
// ValidateMatchmakingRuleSet | DeleteMatchmakingRuleSet | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteMatchmakingRuleSet(ctx context.Context, params *DeleteMatchmakingRuleSetInput, optFns ...func(*Options)) (*DeleteMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &DeleteMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMatchmakingRuleSet", params, optFns, c.addOperationDeleteMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteMatchmakingRuleSetInput struct {

	// A unique identifier for the matchmaking rule set to be deleted. (Note: The rule
	// set name is different from the optional "name" field in the rule set body.) You
	// can use either the rule set name or ARN value.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type DeleteMatchmakingRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMatchmakingRuleSet{}, middleware.After)
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
	if err = addOpDeleteMatchmakingRuleSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMatchmakingRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMatchmakingRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteMatchmakingRuleSet",
	}
}
