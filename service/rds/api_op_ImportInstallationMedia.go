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

// Imports the installation media for a DB engine that requires an on-premises
// customer provided license, such as SQL Server.
func (c *Client) ImportInstallationMedia(ctx context.Context, params *ImportInstallationMediaInput, optFns ...func(*Options)) (*ImportInstallationMediaOutput, error) {
	if params == nil {
		params = &ImportInstallationMediaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportInstallationMedia", params, optFns, c.addOperationImportInstallationMediaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportInstallationMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportInstallationMediaInput struct {

	// The identifier of the custom Availability Zone (AZ) to import the installation
	// media to.
	//
	// This member is required.
	CustomAvailabilityZoneId *string

	// The name of the database engine to be used for this instance. The list only
	// includes supported DB engines that require an on-premises customer provided
	// license. Valid Values:
	//
	// * sqlserver-ee
	//
	// * sqlserver-se
	//
	// * sqlserver-ex
	//
	// *
	// sqlserver-web
	//
	// This member is required.
	Engine *string

	// The path to the installation medium for the specified DB engine. Example:
	// SQLServerISO/en_sql_server_2016_enterprise_x64_dvd_8701793.iso
	//
	// This member is required.
	EngineInstallationMediaPath *string

	// The version number of the database engine to use. For a list of valid engine
	// versions, call DescribeDBEngineVersions. The following are the database engines
	// and links to information about the major and minor versions. The list only
	// includes DB engines that require an on-premises customer provided license.
	// Microsoft SQL Server See  Microsoft SQL Server Versions on Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport)
	// in the Amazon RDS User Guide.
	//
	// This member is required.
	EngineVersion *string

	// The path to the installation medium for the operating system associated with the
	// specified DB engine. Example:
	// WindowsISO/en_windows_server_2016_x64_dvd_9327751.iso
	//
	// This member is required.
	OSInstallationMediaPath *string

	noSmithyDocumentSerde
}

// Contains the installation media for a DB engine that requires an on-premises
// customer provided license, such as Microsoft SQL Server.
type ImportInstallationMediaOutput struct {

	// The custom Availability Zone (AZ) that contains the installation media.
	CustomAvailabilityZoneId *string

	// The DB engine.
	Engine *string

	// The path to the installation medium for the DB engine.
	EngineInstallationMediaPath *string

	// The engine version of the DB engine.
	EngineVersion *string

	// If an installation media failure occurred, the cause of the failure.
	FailureCause *types.InstallationMediaFailureCause

	// The installation medium ID.
	InstallationMediaId *string

	// The path to the installation medium for the operating system associated with the
	// DB engine.
	OSInstallationMediaPath *string

	// The status of the installation medium.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportInstallationMediaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpImportInstallationMedia{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpImportInstallationMedia{}, middleware.After)
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
	if err = addOpImportInstallationMediaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportInstallationMedia(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportInstallationMedia(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ImportInstallationMedia",
	}
}
