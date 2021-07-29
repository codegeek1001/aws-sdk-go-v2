// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact. If the output contains a
// TagOption key with an empty list of values, there is a TagOption conflict for
// that key. The end user cannot take action to fix the conflict, and launch is not
// blocked. In subsequent calls to ProvisionProduct, do not include conflicted
// TagOption keys as tags, or this causes the error "Parameter validation failed:
// Missing required parameter in Tags[N]:Value". Tag the provisioned product with
// the value sc-tagoption-conflict-portfolioId-productId.
func (c *Client) DescribeProvisioningParameters(ctx context.Context, params *DescribeProvisioningParametersInput, optFns ...func(*Options)) (*DescribeProvisioningParametersOutput, error) {
	if params == nil {
		params = &DescribeProvisioningParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProvisioningParameters", params, optFns, c.addOperationDescribeProvisioningParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProvisioningParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProvisioningParametersInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths. You must provide the name or ID, but
	// not both.
	PathId *string

	// The name of the path. You must provide the name or ID, but not both.
	PathName *string

	// The product identifier. You must provide the product name or ID, but not both.
	ProductId *string

	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string

	// The identifier of the provisioning artifact. You must provide the name or ID,
	// but not both.
	ProvisioningArtifactId *string

	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string

	noSmithyDocumentSerde
}

type DescribeProvisioningParametersOutput struct {

	// Information about the constraints used to provision the product.
	ConstraintSummaries []types.ConstraintSummary

	// The output of the provisioning artifact.
	ProvisioningArtifactOutputs []types.ProvisioningArtifactOutput

	// Information about the parameters used to provision the product.
	ProvisioningArtifactParameters []types.ProvisioningArtifactParameter

	// An object that contains information about preferences, such as regions and
	// accounts, for the provisioning artifact.
	ProvisioningArtifactPreferences *types.ProvisioningArtifactPreferences

	// Information about the TagOptions associated with the resource.
	TagOptions []types.TagOptionSummary

	// Any additional metadata specifically related to the provisioning of the product.
	// For example, see the Version field of the CloudFormation template.
	UsageInstructions []types.UsageInstruction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProvisioningParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProvisioningParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProvisioningParameters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisioningParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeProvisioningParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeProvisioningParameters",
	}
}
