// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a virtual interface with a specified link aggregation group (LAG) or
// connection. Connectivity to Amazon Web Services is temporarily interrupted as
// the virtual interface is being migrated. If the target connection or LAG has an
// associated virtual interface with a conflicting VLAN number or a conflicting IP
// address, the operation fails. Virtual interfaces associated with a hosted
// connection cannot be associated with a LAG; hosted connections must be migrated
// along with their virtual interfaces using AssociateHostedConnection. To
// reassociate a virtual interface to a new connection or LAG, the requester must
// own either the virtual interface itself or the connection to which the virtual
// interface is currently associated. Additionally, the requester must own the
// connection or LAG for the association.
func (c *Client) AssociateVirtualInterface(ctx context.Context, params *AssociateVirtualInterfaceInput, optFns ...func(*Options)) (*AssociateVirtualInterfaceOutput, error) {
	if params == nil {
		params = &AssociateVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateVirtualInterface", params, optFns, c.addOperationAssociateVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateVirtualInterfaceInput struct {

	// The ID of the LAG or connection.
	//
	// This member is required.
	ConnectionId *string

	// The ID of the virtual interface.
	//
	// This member is required.
	VirtualInterfaceId *string

	noSmithyDocumentSerde
}

// Information about a virtual interface.
type AssociateVirtualInterfaceOutput struct {

	// The address family for the BGP peer.
	AddressFamily types.AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration. The valid values are 1-2147483647.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates a physical connection's BGP
	// sessions.
	AwsLogicalDeviceId *string

	// The BGP peers configured on this virtual interface.
	BgpPeers []types.BGPPeer

	// The ID of the connection.
	ConnectionId *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The customer router configuration.
	CustomerRouterConfig *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The location of the connection.
	Location *string

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500 and
	// 9001. The default value is 1500.
	Mtu *int32

	// The ID of the account that owns the virtual interface.
	OwnerAccount *string

	// The Region where the virtual interface is located.
	Region *string

	// The routes to be advertised to the Amazon Web Services network in this Region.
	// Applies to public virtual interfaces.
	RouteFilterPrefixes []types.RouteFilterPrefix

	// The tags associated with the virtual interface.
	Tags []types.Tag

	// The ID of the virtual private gateway. Applies only to private virtual
	// interfaces.
	VirtualGatewayId *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	// The name of the virtual interface assigned by the customer network. The name has
	// a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a
	// hyphen (-).
	VirtualInterfaceName *string

	// The state of the virtual interface. The following are the possible values:
	//
	// *
	// confirming: The creation of the virtual interface is pending confirmation from
	// the virtual interface owner. If the owner of the virtual interface is different
	// from the owner of the connection on which it is provisioned, then the virtual
	// interface will remain in this state until it is confirmed by the virtual
	// interface owner.
	//
	// * verifying: This state only applies to public virtual
	// interfaces. Each public virtual interface needs validation before the virtual
	// interface can be created.
	//
	// * pending: A virtual interface is in this state from
	// the time that it is created until the virtual interface is ready to forward
	// traffic.
	//
	// * available: A virtual interface that is able to forward traffic.
	//
	// *
	// down: A virtual interface that is BGP down.
	//
	// * deleting: A virtual interface is
	// in this state immediately after calling DeleteVirtualInterface until it can no
	// longer forward traffic.
	//
	// * deleted: A virtual interface that cannot forward
	// traffic.
	//
	// * rejected: The virtual interface owner has declined creation of the
	// virtual interface. If a virtual interface in the Confirming state is deleted by
	// the virtual interface owner, the virtual interface enters the Rejected state.
	//
	// *
	// unknown: The state of the virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateVirtualInterface{}, middleware.After)
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
	if err = addOpAssociateVirtualInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateVirtualInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateVirtualInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AssociateVirtualInterface",
	}
}
