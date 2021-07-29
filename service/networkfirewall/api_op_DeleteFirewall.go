// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified Firewall and its FirewallStatus. This operation requires
// the firewall's DeleteProtection flag to be FALSE. You can't revert this
// operation. You can check whether a firewall is in use by reviewing the route
// tables for the Availability Zones where you have firewall subnet mappings.
// Retrieve the subnet mappings by calling DescribeFirewall. You define and update
// the route tables through Amazon VPC. As needed, update the route tables for the
// zones to remove the firewall endpoints. When the route tables no longer use the
// firewall endpoints, you can remove the firewall safely. To delete a firewall,
// remove the delete protection if you need to using
// UpdateFirewallDeleteProtection, then delete the firewall by calling
// DeleteFirewall.
func (c *Client) DeleteFirewall(ctx context.Context, params *DeleteFirewallInput, optFns ...func(*Options)) (*DeleteFirewallOutput, error) {
	if params == nil {
		params = &DeleteFirewallInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFirewall", params, optFns, c.addOperationDeleteFirewallMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFirewallOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFirewallInput struct {

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	noSmithyDocumentSerde
}

type DeleteFirewallOutput struct {

	// The firewall defines the configuration settings for an AWS Network Firewall
	// firewall. These settings include the firewall policy, the subnets in your VPC to
	// use for the firewall endpoints, and any tags that are attached to the firewall
	// AWS resource. The status of the firewall, for example whether it's ready to
	// filter network traffic, is provided in the corresponding FirewallStatus. You can
	// retrieve both objects by calling DescribeFirewall.
	Firewall *types.Firewall

	// Detailed information about the current status of a Firewall. You can retrieve
	// this for a firewall by calling DescribeFirewall and providing the firewall name
	// and ARN.
	FirewallStatus *types.FirewallStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFirewallMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteFirewall{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteFirewall{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFirewall(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFirewall(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "DeleteFirewall",
	}
}
