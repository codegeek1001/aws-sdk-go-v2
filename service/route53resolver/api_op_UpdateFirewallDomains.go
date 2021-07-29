// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the firewall domain list from an array of domain specifications.
func (c *Client) UpdateFirewallDomains(ctx context.Context, params *UpdateFirewallDomainsInput, optFns ...func(*Options)) (*UpdateFirewallDomainsOutput, error) {
	if params == nil {
		params = &UpdateFirewallDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallDomains", params, optFns, c.addOperationUpdateFirewallDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallDomainsInput struct {

	// A list of domains to use in the update operation. Each domain specification in
	// your domain list must satisfy the following requirements:
	//
	// * It can optionally
	// start with * (asterisk).
	//
	// * With the exception of the optional starting
	// asterisk, it must only contain the following characters: A-Z, a-z, 0-9, -
	// (hyphen).
	//
	// * It must be from 1-255 characters in length.
	//
	// This member is required.
	Domains []string

	// The ID of the domain list whose domains you want to update.
	//
	// This member is required.
	FirewallDomainListId *string

	// What you want DNS Firewall to do with the domains that you are providing:
	//
	// * ADD
	// - Add the domains to the ones that are already in the domain list.
	//
	// * REMOVE -
	// Search the domain list for the domains and remove them from the list.
	//
	// * REPLACE
	// - Update the domain list to exactly match the list that you are providing.
	//
	// This member is required.
	Operation types.FirewallDomainUpdateOperation

	noSmithyDocumentSerde
}

type UpdateFirewallDomainsOutput struct {

	// The ID of the firewall domain list that DNS Firewall just updated.
	Id *string

	// The name of the domain list.
	Name *string

	//
	Status types.FirewallDomainListStatus

	// Additional information about the status of the list, if available.
	StatusMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFirewallDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFirewallDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFirewallDomains{}, middleware.After)
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
	if err = addOpUpdateFirewallDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "UpdateFirewallDomains",
	}
}
