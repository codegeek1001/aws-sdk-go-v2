// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all managed policies that are attached to the specified IAM group. An IAM
// group can also have inline policies embedded with it. To list the inline
// policies for a group, use ListGroupPolicies. For information about policies, see
// Managed policies and inline policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. You can paginate the results using the MaxItems and
// Marker parameters. You can use the PathPrefix parameter to limit the list of
// policies to only those matching the specified path prefix. If there are no
// policies attached to the specified group (or none that match the specified path
// prefix), the operation returns an empty list.
func (c *Client) ListAttachedGroupPolicies(ctx context.Context, params *ListAttachedGroupPoliciesInput, optFns ...func(*Options)) (*ListAttachedGroupPoliciesOutput, error) {
	if params == nil {
		params = &ListAttachedGroupPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttachedGroupPolicies", params, optFns, c.addOperationListAttachedGroupPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttachedGroupPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedGroupPoliciesInput struct {

	// The name (friendly name, not ARN) of the group to list attached policies for.
	// This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	GroupName *string

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The path prefix for filtering the results. This parameter is optional. If it is
	// not included, it defaults to a slash (/), listing all policies. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of either a forward slash (/) by itself or a string that
	// must begin and end with forward slashes. In addition, it can contain any ASCII
	// character from the ! (\u0021) through the DEL character (\u007F), including most
	// punctuation characters, digits, and upper and lowercased letters.
	PathPrefix *string

	noSmithyDocumentSerde
}

// Contains the response to a successful ListAttachedGroupPolicies request.
type ListAttachedGroupPoliciesOutput struct {

	// A list of the attached policies.
	AttachedPolicies []types.AttachedPolicy

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttachedGroupPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListAttachedGroupPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListAttachedGroupPolicies{}, middleware.After)
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
	if err = addOpListAttachedGroupPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedGroupPolicies(options.Region), middleware.Before); err != nil {
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

// ListAttachedGroupPoliciesAPIClient is a client that implements the
// ListAttachedGroupPolicies operation.
type ListAttachedGroupPoliciesAPIClient interface {
	ListAttachedGroupPolicies(context.Context, *ListAttachedGroupPoliciesInput, ...func(*Options)) (*ListAttachedGroupPoliciesOutput, error)
}

var _ ListAttachedGroupPoliciesAPIClient = (*Client)(nil)

// ListAttachedGroupPoliciesPaginatorOptions is the paginator options for
// ListAttachedGroupPolicies
type ListAttachedGroupPoliciesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttachedGroupPoliciesPaginator is a paginator for ListAttachedGroupPolicies
type ListAttachedGroupPoliciesPaginator struct {
	options   ListAttachedGroupPoliciesPaginatorOptions
	client    ListAttachedGroupPoliciesAPIClient
	params    *ListAttachedGroupPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListAttachedGroupPoliciesPaginator returns a new
// ListAttachedGroupPoliciesPaginator
func NewListAttachedGroupPoliciesPaginator(client ListAttachedGroupPoliciesAPIClient, params *ListAttachedGroupPoliciesInput, optFns ...func(*ListAttachedGroupPoliciesPaginatorOptions)) *ListAttachedGroupPoliciesPaginator {
	if params == nil {
		params = &ListAttachedGroupPoliciesInput{}
	}

	options := ListAttachedGroupPoliciesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttachedGroupPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttachedGroupPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAttachedGroupPolicies page.
func (p *ListAttachedGroupPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttachedGroupPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListAttachedGroupPolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAttachedGroupPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListAttachedGroupPolicies",
	}
}
