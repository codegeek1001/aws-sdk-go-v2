// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	kinesiscust "github.com/aws/aws-sdk-go-v2/service/kinesis/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets data records from a Kinesis data stream's shard. Specify a shard iterator
// using the ShardIterator parameter. The shard iterator specifies the position in
// the shard from which you want to start reading data records sequentially. If
// there are no records available in the portion of the shard that the iterator
// points to, GetRecords returns an empty list. It might take multiple calls to get
// to a portion of the shard that contains records. You can scale by provisioning
// multiple shards per stream while considering service limits (for more
// information, see Amazon Kinesis Data Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide). Your application should
// have one thread per shard, each reading continuously from its stream. To read
// from a stream continually, call GetRecords in a loop. Use GetShardIterator to
// get the shard iterator to specify in the first GetRecords call. GetRecords
// returns a new shard iterator in NextShardIterator. Specify the shard iterator
// returned in NextShardIterator in subsequent calls to GetRecords. If the shard
// has been closed, the shard iterator can't return more data and GetRecords
// returns null in NextShardIterator. You can terminate the loop when the shard is
// closed, or when the shard iterator reaches the record with the sequence number
// or other attribute that marks it as the last record to process. Each data record
// can be up to 1 MiB in size, and each shard can read up to 2 MiB per second. You
// can ensure that your calls don't exceed the maximum supported size or throughput
// by using the Limit parameter to specify the maximum number of records that
// GetRecords can return. Consider your average record size when determining this
// limit. The maximum number of records that can be returned per call is 10,000.
// The size of the data returned by GetRecords varies depending on the utilization
// of the shard. The maximum size of data that GetRecords can return is 10 MiB. If
// a call returns this amount of data, subsequent calls made within the next 5
// seconds throw ProvisionedThroughputExceededException. If there is insufficient
// provisioned throughput on the stream, subsequent calls made within the next 1
// second throw ProvisionedThroughputExceededException. GetRecords doesn't return
// any data when it throws an exception. For this reason, we recommend that you
// wait 1 second between calls to GetRecords. However, it's possible that the
// application will get exceptions for longer than 1 second. To detect whether the
// application is falling behind in processing, you can use the MillisBehindLatest
// response attribute. You can also monitor the stream using CloudWatch metrics and
// other mechanisms (see Monitoring
// (https://docs.aws.amazon.com/kinesis/latest/dev/monitoring.html) in the Amazon
// Kinesis Data Streams Developer Guide). Each Amazon Kinesis record includes a
// value, ApproximateArrivalTimestamp, that is set when a stream successfully
// receives and stores a record. This is commonly referred to as a server-side time
// stamp, whereas a client-side time stamp is set when a data producer creates or
// sends the record to a stream (a data producer is any data source putting data
// records into a stream, for example with PutRecords). The time stamp has
// millisecond precision. There are no guarantees about the time stamp accuracy, or
// that the time stamp is always increasing. For example, records in a shard or
// across a stream might have time stamps that are out of order. This operation has
// a limit of five transactions per second per shard.
func (c *Client) GetRecords(ctx context.Context, params *GetRecordsInput, optFns ...func(*Options)) (*GetRecordsOutput, error) {
	if params == nil {
		params = &GetRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecords", params, optFns, c.addOperationGetRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for GetRecords.
type GetRecordsInput struct {

	// The position in the shard from which you want to start sequentially reading data
	// records. A shard iterator specifies this position using the sequence number of a
	// data record in the shard.
	//
	// This member is required.
	ShardIterator *string

	// The maximum number of records to return. Specify a value of up to 10,000. If you
	// specify a value that is greater than 10,000, GetRecords throws
	// InvalidArgumentException. The default value is 10,000.
	Limit *int32

	noSmithyDocumentSerde
}

// Represents the output for GetRecords.
type GetRecordsOutput struct {

	// The data records retrieved from the shard.
	//
	// This member is required.
	Records []types.Record

	ChildShards []types.ChildShard

	// The number of milliseconds the GetRecords response is from the tip of the
	// stream, indicating how far behind current time the consumer is. A value of zero
	// indicates that record processing is caught up, and there are no new records to
	// process at this moment.
	MillisBehindLatest *int64

	// The next position in the shard from which to start sequentially reading data
	// records. If set to null, the shard has been closed and the requested iterator
	// does not return any more data.
	NextShardIterator *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRecords{}, middleware.After)
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
	if err = addOpGetRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecords(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = awshttp.AddResponseReadTimeoutMiddleware(stack, kinesiscust.ReadTimeoutDuration); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "GetRecords",
	}
}
