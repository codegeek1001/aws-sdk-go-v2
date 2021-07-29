// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of Challenge-Handshake Authentication Protocol (CHAP)
// credentials information for a specified iSCSI target, one for each
// target-initiator pair. This operation is supported in the volume and tape
// gateway types.
func (c *Client) DescribeChapCredentials(ctx context.Context, params *DescribeChapCredentialsInput, optFns ...func(*Options)) (*DescribeChapCredentialsOutput, error) {
	if params == nil {
		params = &DescribeChapCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChapCredentials", params, optFns, c.addOperationDescribeChapCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the iSCSI volume
// target.
type DescribeChapCredentialsInput struct {

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	//
	// This member is required.
	TargetARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type DescribeChapCredentialsOutput struct {

	// An array of ChapInfo objects that represent CHAP credentials. Each object in the
	// array contains CHAP credential information for one target-initiator pair. If no
	// CHAP credentials are set, an empty array is returned. CHAP credential
	// information is provided in a JSON object with the following fields:
	//
	// *
	// InitiatorName: The iSCSI initiator that connects to the target.
	//
	// *
	// SecretToAuthenticateInitiator: The secret key that the initiator (for example,
	// the Windows client) must provide to participate in mutual CHAP with the
	// target.
	//
	// * SecretToAuthenticateTarget: The secret key that the target must
	// provide to participate in mutual CHAP with the initiator (e.g. Windows
	// client).
	//
	// * TargetARN: The Amazon Resource Name (ARN) of the storage volume.
	ChapCredentials []types.ChapInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChapCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeChapCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeChapCredentials{}, middleware.After)
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
	if err = addOpDescribeChapCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChapCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChapCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeChapCredentials",
	}
}
