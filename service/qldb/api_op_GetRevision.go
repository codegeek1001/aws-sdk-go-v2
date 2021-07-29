// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a revision data object for a specified document ID and block address.
// Also returns a proof of the specified revision for verification if
// DigestTipAddress is provided.
func (c *Client) GetRevision(ctx context.Context, params *GetRevisionInput, optFns ...func(*Options)) (*GetRevisionOutput, error) {
	if params == nil {
		params = &GetRevisionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRevision", params, optFns, c.addOperationGetRevisionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRevisionInput struct {

	// The block location of the document revision to be verified. An address is an
	// Amazon Ion structure that has two fields: strandId and sequenceNo. For example:
	// {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:14}.
	//
	// This member is required.
	BlockAddress *types.ValueHolder

	// The UUID (represented in Base62-encoded text) of the document to be verified.
	//
	// This member is required.
	DocumentId *string

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The latest block location covered by the digest for which to request a proof. An
	// address is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:49}.
	DigestTipAddress *types.ValueHolder

	noSmithyDocumentSerde
}

type GetRevisionOutput struct {

	// The document revision data object in Amazon Ion format.
	//
	// This member is required.
	Revision *types.ValueHolder

	// The proof object in Amazon Ion format returned by a GetRevision request. A proof
	// contains the list of hash values that are required to recalculate the specified
	// digest using a Merkle tree, starting with the specified document revision.
	Proof *types.ValueHolder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRevisionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRevision{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRevision{}, middleware.After)
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
	if err = addOpGetRevisionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRevision(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRevision(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "GetRevision",
	}
}
