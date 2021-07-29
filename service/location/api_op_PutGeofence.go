// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stores a geofence geometry in a given geofence collection, or updates the
// geometry of an existing geofence if a geofence ID is included in the request.
func (c *Client) PutGeofence(ctx context.Context, params *PutGeofenceInput, optFns ...func(*Options)) (*PutGeofenceOutput, error) {
	if params == nil {
		params = &PutGeofenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGeofence", params, optFns, c.addOperationPutGeofenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGeofenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutGeofenceInput struct {

	// The geofence collection to store the geofence in.
	//
	// This member is required.
	CollectionName *string

	// An identifier for the geofence. For example, ExampleGeofence-1.
	//
	// This member is required.
	GeofenceId *string

	// Contains the polygon details to specify the position of the geofence. Each
	// geofence polygon
	// (https://docs.aws.amazon.com/location-geofences/latest/APIReference/API_GeofenceGeometry.html)
	// can have a maximum of 1,000 vertices.
	//
	// This member is required.
	Geometry *types.GeofenceGeometry

	noSmithyDocumentSerde
}

type PutGeofenceOutput struct {

	// The timestamp for when the geofence was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	CreateTime *time.Time

	// The geofence identifier entered in the request.
	//
	// This member is required.
	GeofenceId *string

	// The timestamp for when the geofence was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutGeofenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutGeofence{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutGeofence{}, middleware.After)
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
	if err = addOpPutGeofenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGeofence(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutGeofence(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "PutGeofence",
	}
}
