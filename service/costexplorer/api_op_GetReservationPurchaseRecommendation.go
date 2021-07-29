// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets recommendations for which reservations to purchase. These recommendations
// could help you reduce your costs. Reservations provide a discounted hourly rate
// (up to 75%) compared to On-Demand pricing. AWS generates your recommendations by
// identifying your On-Demand usage during a specific time period and collecting
// your usage into categories that are eligible for a reservation. After AWS has
// these categories, it simulates every combination of reservations in each
// category of usage to identify the best number of each type of RI to purchase to
// maximize your estimated savings. For example, AWS automatically aggregates your
// Amazon EC2 Linux, shared tenancy, and c4 family usage in the US West (Oregon)
// Region and recommends that you buy size-flexible regional reservations to apply
// to the c4 family usage. AWS recommends the smallest size instance in an instance
// family. This makes it easier to purchase a size-flexible RI. AWS also shows the
// equal number of normalized units so that you can purchase any instance size that
// you want. For this example, your RI recommendation would be for c4.large because
// that is the smallest size instance in the c4 instance family.
func (c *Client) GetReservationPurchaseRecommendation(ctx context.Context, params *GetReservationPurchaseRecommendationInput, optFns ...func(*Options)) (*GetReservationPurchaseRecommendationOutput, error) {
	if params == nil {
		params = &GetReservationPurchaseRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservationPurchaseRecommendation", params, optFns, c.addOperationGetReservationPurchaseRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservationPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservationPurchaseRecommendationInput struct {

	// The specific service that you want recommendations for.
	//
	// This member is required.
	Service *string

	// The account ID that is associated with the recommendation.
	AccountId *string

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the management account and member accounts
	// if the value is set to PAYER. If the value is LINKED, recommendations are
	// calculated for individual member accounts only.
	AccountScope types.AccountScope

	// Use Expression to filter by cost or by usage. There are two patterns:
	//
	// * Simple
	// dimension values - You can set the dimension name and values for the filters
	// that you plan to use. For example, you can filter for REGION==us-east-1 OR
	// REGION==us-west-1. For GetRightsizingRecommendation, the Region is a full name
	// (for example, REGION==US East (N. Virginia). The Expression example looks like:
	// { "Dimensions": { "Key": "REGION", "Values": [ "us-east-1", “us-west-1” ] } }
	// The list of dimension values are OR'd together to retrieve cost or usage data.
	// You can create Expression and DimensionValues objects using either with* methods
	// or set* methods in multiple lines.
	//
	// * Compound dimension values with logical
	// operations - You can use multiple Expression types and the logical operators
	// AND/OR/NOT to create a list of one or more Expression objects. This allows you
	// to filter on more advanced options. For example, you can filter on ((REGION ==
	// us-east-1 OR REGION == us-west-1) OR (TAG.Type == Type1)) AND (USAGE_TYPE !=
	// DataTransfer). The Expression for that looks like this: { "And": [ {"Or": [
	// {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ] }},
	// {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not": {"Dimensions":
	// { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }  Because each
	// Expression can have only one operator, the service returns an error if more than
	// one is specified. The following example shows an Expression object that creates
	// an error.  { "And": [ ... ], "DimensionValues": { "Dimension": "USAGE_TYPE",
	// "Values": [ "DataTransfer" ] } }
	//
	// For the GetRightsizingRecommendation action, a
	// combination of OR and NOT is not supported. OR is not supported between
	// different dimensions, or dimensions and tags. NOT operators aren't supported.
	// Dimensions are also limited to LINKED_ACCOUNT, REGION, or RIGHTSIZING_TYPE. For
	// the GetReservationPurchaseRecommendation action, only NOT is supported. AND and
	// OR are not supported. Dimensions are limited to LINKED_ACCOUNT.
	Filter *types.Expression

	// The number of previous days that you want AWS to consider when it calculates
	// your recommendations.
	LookbackPeriodInDays types.LookbackPeriodInDays

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	// The reservation purchase option that you want recommendations for.
	PaymentOption types.PaymentOption

	// The hardware specifications for the service instances that you want
	// recommendations for, such as standard or convertible Amazon EC2 instances.
	ServiceSpecification *types.ServiceSpecification

	// The reservation term that you want recommendations for.
	TermInYears types.TermInYears

	noSmithyDocumentSerde
}

type GetReservationPurchaseRecommendationOutput struct {

	// Information about this specific recommendation call, such as the time stamp for
	// when Cost Explorer generated this recommendation.
	Metadata *types.ReservationPurchaseRecommendationMetadata

	// The pagination token for the next set of retrievable results.
	NextPageToken *string

	// Recommendations for reservations to purchase.
	Recommendations []types.ReservationPurchaseRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservationPurchaseRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReservationPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReservationPurchaseRecommendation{}, middleware.After)
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
	if err = addOpGetReservationPurchaseRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReservationPurchaseRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetReservationPurchaseRecommendation",
	}
}
