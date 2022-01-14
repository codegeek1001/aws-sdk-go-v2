// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
)

func ExampleAssistantAssociationInputData_outputUsage() {
	var union types.AssistantAssociationInputData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AssistantAssociationInputDataMemberKnowledgeBaseId:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string

func ExampleAssistantAssociationOutputData_outputUsage() {
	var union types.AssistantAssociationOutputData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AssistantAssociationOutputDataMemberKnowledgeBaseAssociation:
		_ = v.Value // Value is types.KnowledgeBaseAssociationData

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.KnowledgeBaseAssociationData

func ExampleSourceConfiguration_outputUsage() {
	var union types.SourceConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SourceConfigurationMemberAppIntegrations:
		_ = v.Value // Value is types.AppIntegrationsConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AppIntegrationsConfiguration