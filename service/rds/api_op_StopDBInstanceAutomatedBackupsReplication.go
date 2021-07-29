// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops automated backup replication for a DB instance. For more information, see
// Replicating Automated Backups to Another Amazon Web Services Region
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReplicateBackups.html)
// in the Amazon RDS User Guide.
func (c *Client) StopDBInstanceAutomatedBackupsReplication(ctx context.Context, params *StopDBInstanceAutomatedBackupsReplicationInput, optFns ...func(*Options)) (*StopDBInstanceAutomatedBackupsReplicationOutput, error) {
	if params == nil {
		params = &StopDBInstanceAutomatedBackupsReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDBInstanceAutomatedBackupsReplication", params, optFns, c.addOperationStopDBInstanceAutomatedBackupsReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDBInstanceAutomatedBackupsReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDBInstanceAutomatedBackupsReplicationInput struct {

	// The Amazon Resource Name (ARN) of the source DB instance for which to stop
	// replicating automated backups, for example,
	// arn:aws:rds:us-west-2:123456789012:db:mydatabase.
	//
	// This member is required.
	SourceDBInstanceArn *string

	noSmithyDocumentSerde
}

type StopDBInstanceAutomatedBackupsReplicationOutput struct {

	// An automated backup of a DB instance. It consists of system backups, transaction
	// logs, and the database instance properties that existed at the time you deleted
	// the source instance.
	DBInstanceAutomatedBackup *types.DBInstanceAutomatedBackup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopDBInstanceAutomatedBackupsReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStopDBInstanceAutomatedBackupsReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStopDBInstanceAutomatedBackupsReplication{}, middleware.After)
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
	if err = addOpStopDBInstanceAutomatedBackupsReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopDBInstanceAutomatedBackupsReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopDBInstanceAutomatedBackupsReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopDBInstanceAutomatedBackupsReplication",
	}
}
