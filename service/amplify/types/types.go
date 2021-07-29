// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the different branches of a repository for building, deploying, and
// hosting an Amplify app.
type App struct {

	// The Amazon Resource Name (ARN) of the Amplify app.
	//
	// This member is required.
	AppArn *string

	// The unique ID of the Amplify app.
	//
	// This member is required.
	AppId *string

	// Creates a date and time for the Amplify app.
	//
	// This member is required.
	CreateTime *time.Time

	// The default domain for the Amplify app.
	//
	// This member is required.
	DefaultDomain *string

	// The description for the Amplify app.
	//
	// This member is required.
	Description *string

	// Enables basic authorization for the Amplify app's branches.
	//
	// This member is required.
	EnableBasicAuth *bool

	// Enables the auto-building of branches for the Amplify app.
	//
	// This member is required.
	EnableBranchAutoBuild *bool

	// The environment variables for the Amplify app.
	//
	// This member is required.
	EnvironmentVariables map[string]string

	// The name for the Amplify app.
	//
	// This member is required.
	Name *string

	// The platform for the Amplify app.
	//
	// This member is required.
	Platform Platform

	// The repository for the Amplify app.
	//
	// This member is required.
	Repository *string

	// Updates the date and time for the Amplify app.
	//
	// This member is required.
	UpdateTime *time.Time

	// Describes the automated branch creation configuration for the Amplify app.
	AutoBranchCreationConfig *AutoBranchCreationConfig

	// Describes the automated branch creation glob patterns for the Amplify app.
	AutoBranchCreationPatterns []string

	// The basic authorization credentials for branches for the Amplify app.
	BasicAuthCredentials *string

	// Describes the content of the build specification (build spec) for the Amplify
	// app.
	BuildSpec *string

	// Describes the custom HTTP headers for the Amplify app.
	CustomHeaders *string

	// Describes the custom redirect and rewrite rules for the Amplify app.
	CustomRules []CustomRule

	// Enables automated branch creation for the Amplify app.
	EnableAutoBranchCreation *bool

	// Automatically disconnect a branch in the Amplify Console when you delete a
	// branch from your Git repository.
	EnableBranchAutoDeletion *bool

	// The AWS Identity and Access Management (IAM) service role for the Amazon
	// Resource Name (ARN) of the Amplify app.
	IamServiceRoleArn *string

	// Describes the information about a production branch of the Amplify app.
	ProductionBranch *ProductionBranch

	// The tag for the Amplify app.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes an artifact.
type Artifact struct {

	// The file name for the artifact.
	//
	// This member is required.
	ArtifactFileName *string

	// The unique ID for the artifact.
	//
	// This member is required.
	ArtifactId *string

	noSmithyDocumentSerde
}

// Describes the automated branch creation configuration.
type AutoBranchCreationConfig struct {

	// The basic authorization credentials for the autocreated branch.
	BasicAuthCredentials *string

	// The build specification (build spec) for the autocreated branch.
	BuildSpec *string

	// Enables auto building for the autocreated branch.
	EnableAutoBuild *bool

	// Enables basic authorization for the autocreated branch.
	EnableBasicAuth *bool

	// Enables performance mode for the branch. Performance mode optimizes for faster
	// hosting performance by keeping content cached at the edge for a longer interval.
	// When performance mode is enabled, hosting configuration or code changes can take
	// up to 10 minutes to roll out.
	EnablePerformanceMode *bool

	// Enables pull request previews for the autocreated branch.
	EnablePullRequestPreview *bool

	// The environment variables for the autocreated branch.
	EnvironmentVariables map[string]string

	// The framework for the autocreated branch.
	Framework *string

	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string

	// Describes the current stage for the autocreated branch.
	Stage Stage

	noSmithyDocumentSerde
}

// Describes the backend environment for an Amplify app.
type BackendEnvironment struct {

	// The Amazon Resource Name (ARN) for a backend environment that is part of an
	// Amplify app.
	//
	// This member is required.
	BackendEnvironmentArn *string

	// The creation date and time for a backend environment that is part of an Amplify
	// app.
	//
	// This member is required.
	CreateTime *time.Time

	// The name for a backend environment that is part of an Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The last updated date and time for a backend environment that is part of an
	// Amplify app.
	//
	// This member is required.
	UpdateTime *time.Time

	// The name of deployment artifacts.
	DeploymentArtifacts *string

	// The AWS CloudFormation stack name of a backend environment.
	StackName *string

	noSmithyDocumentSerde
}

// The branch for an Amplify app, which maps to a third-party repository branch.
type Branch struct {

	// The ID of the active job for a branch of an Amplify app.
	//
	// This member is required.
	ActiveJobId *string

	// The Amazon Resource Name (ARN) for a branch that is part of an Amplify app.
	//
	// This member is required.
	BranchArn *string

	// The name for the branch that is part of an Amplify app.
	//
	// This member is required.
	BranchName *string

	// The creation date and time for a branch that is part of an Amplify app.
	//
	// This member is required.
	CreateTime *time.Time

	// The custom domains for a branch of an Amplify app.
	//
	// This member is required.
	CustomDomains []string

	// The description for the branch that is part of an Amplify app.
	//
	// This member is required.
	Description *string

	// The display name for the branch. This is used as the default domain prefix.
	//
	// This member is required.
	DisplayName *string

	// Enables auto-building on push for a branch of an Amplify app.
	//
	// This member is required.
	EnableAutoBuild *bool

	// Enables basic authorization for a branch of an Amplify app.
	//
	// This member is required.
	EnableBasicAuth *bool

	// Enables notifications for a branch that is part of an Amplify app.
	//
	// This member is required.
	EnableNotification *bool

	// Enables pull request previews for the branch.
	//
	// This member is required.
	EnablePullRequestPreview *bool

	// The environment variables specific to a branch of an Amplify app.
	//
	// This member is required.
	EnvironmentVariables map[string]string

	// The framework for a branch of an Amplify app.
	//
	// This member is required.
	Framework *string

	// The current stage for the branch that is part of an Amplify app.
	//
	// This member is required.
	Stage Stage

	// The total number of jobs that are part of an Amplify app.
	//
	// This member is required.
	TotalNumberOfJobs *string

	// The content Time to Live (TTL) for the website in seconds.
	//
	// This member is required.
	Ttl *string

	// The last updated date and time for a branch that is part of an Amplify app.
	//
	// This member is required.
	UpdateTime *time.Time

	// A list of custom resources that are linked to this branch.
	AssociatedResources []string

	// The Amazon Resource Name (ARN) for a backend environment that is part of an
	// Amplify app.
	BackendEnvironmentArn *string

	// The basic authorization credentials for a branch of an Amplify app.
	BasicAuthCredentials *string

	// The build specification (build spec) content for the branch of an Amplify app.
	BuildSpec *string

	// The destination branch if the branch is a pull request branch.
	DestinationBranch *string

	// Enables performance mode for the branch. Performance mode optimizes for faster
	// hosting performance by keeping content cached at the edge for a longer interval.
	// When performance mode is enabled, hosting configuration or code changes can take
	// up to 10 minutes to roll out.
	EnablePerformanceMode *bool

	// The Amplify environment name for the pull request.
	PullRequestEnvironmentName *string

	// The source branch if the branch is a pull request branch.
	SourceBranch *string

	// The tag for the branch of an Amplify app.
	Tags map[string]string

	// The thumbnail URL for the branch of an Amplify app.
	ThumbnailUrl *string

	noSmithyDocumentSerde
}

// Describes a custom rewrite or redirect rule.
type CustomRule struct {

	// The source pattern for a URL rewrite or redirect rule.
	//
	// This member is required.
	Source *string

	// The target pattern for a URL rewrite or redirect rule.
	//
	// This member is required.
	Target *string

	// The condition for a URL rewrite or redirect rule, such as a country code.
	Condition *string

	// The status code for a URL rewrite or redirect rule. 200 Represents a 200 rewrite
	// rule. 301 Represents a 301 (moved pemanently) redirect rule. This and all future
	// requests should be directed to the target URL. 302 Represents a 302 temporary
	// redirect rule. 404 Represents a 404 redirect rule. 404-200 Represents a 404
	// rewrite rule.
	Status *string

	noSmithyDocumentSerde
}

// Describes a domain association that associates a custom domain with an Amplify
// app.
type DomainAssociation struct {

	// The Amazon Resource Name (ARN) for the domain association.
	//
	// This member is required.
	DomainAssociationArn *string

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The current status of the domain association.
	//
	// This member is required.
	DomainStatus DomainStatus

	// Enables the automated creation of subdomains for branches.
	//
	// This member is required.
	EnableAutoSubDomain *bool

	// The reason for the current status of the domain association.
	//
	// This member is required.
	StatusReason *string

	// The subdomains for the domain association.
	//
	// This member is required.
	SubDomains []SubDomain

	// Sets branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns []string

	// The required AWS Identity and Access Management (IAM) service role for the
	// Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIAMRole *string

	// The DNS record for certificate verification.
	CertificateVerificationDNSRecord *string

	noSmithyDocumentSerde
}

// Describes an execution job for an Amplify app.
type Job struct {

	// The execution steps for an execution job, for an Amplify app.
	//
	// This member is required.
	Steps []Step

	// Describes the summary for an execution job for an Amplify app.
	//
	// This member is required.
	Summary *JobSummary

	noSmithyDocumentSerde
}

// Describes the summary for an execution job for an Amplify app.
type JobSummary struct {

	// The commit ID from a third-party repository provider for the job.
	//
	// This member is required.
	CommitId *string

	// The commit message from a third-party repository provider for the job.
	//
	// This member is required.
	CommitMessage *string

	// The commit date and time for the job.
	//
	// This member is required.
	CommitTime *time.Time

	// The Amazon Resource Name (ARN) for the job.
	//
	// This member is required.
	JobArn *string

	// The unique ID for the job.
	//
	// This member is required.
	JobId *string

	// The type for the job. If the value is RELEASE, the job was manually released
	// from its source by using the StartJob API. If the value is RETRY, the job was
	// manually retried using the StartJob API. If the value is WEB_HOOK, the job was
	// automatically triggered by webhooks.
	//
	// This member is required.
	JobType JobType

	// The start date and time for the job.
	//
	// This member is required.
	StartTime *time.Time

	// The current status for the job.
	//
	// This member is required.
	Status JobStatus

	// The end date and time for the job.
	EndTime *time.Time

	noSmithyDocumentSerde
}

// Describes the information about a production branch for an Amplify app.
type ProductionBranch struct {

	// The branch name for the production branch.
	BranchName *string

	// The last deploy time of the production branch.
	LastDeployTime *time.Time

	// The status of the production branch.
	Status *string

	// The thumbnail URL for the production branch.
	ThumbnailUrl *string

	noSmithyDocumentSerde
}

// Describes an execution step, for an execution job, for an Amplify app.
type Step struct {

	// The end date and time of the execution step.
	//
	// This member is required.
	EndTime *time.Time

	// The start date and time of the execution step.
	//
	// This member is required.
	StartTime *time.Time

	// The status of the execution step.
	//
	// This member is required.
	Status JobStatus

	// The name of the execution step.
	//
	// This member is required.
	StepName *string

	// The URL to the artifact for the execution step.
	ArtifactsUrl *string

	// The context for the current step. Includes a build image if the step is build.
	Context *string

	// The URL to the logs for the execution step.
	LogUrl *string

	// The list of screenshot URLs for the execution step, if relevant.
	Screenshots map[string]string

	// The reason for the current step status.
	StatusReason *string

	// The URL to the test artifact for the execution step.
	TestArtifactsUrl *string

	// The URL to the test configuration for the execution step.
	TestConfigUrl *string

	noSmithyDocumentSerde
}

// The subdomain for the domain association.
type SubDomain struct {

	// The DNS record for the subdomain.
	//
	// This member is required.
	DnsRecord *string

	// Describes the settings for the subdomain.
	//
	// This member is required.
	SubDomainSetting *SubDomainSetting

	// The verified status of the subdomain
	//
	// This member is required.
	Verified *bool

	noSmithyDocumentSerde
}

// Describes the settings for the subdomain.
type SubDomainSetting struct {

	// The branch name setting for the subdomain.
	//
	// This member is required.
	BranchName *string

	// The prefix setting for the subdomain.
	//
	// This member is required.
	Prefix *string

	noSmithyDocumentSerde
}

// Describes a webhook that connects repository events to an Amplify app.
type Webhook struct {

	// The name for a branch that is part of an Amplify app.
	//
	// This member is required.
	BranchName *string

	// The create date and time for a webhook.
	//
	// This member is required.
	CreateTime *time.Time

	// The description for a webhook.
	//
	// This member is required.
	Description *string

	// Updates the date and time for a webhook.
	//
	// This member is required.
	UpdateTime *time.Time

	// The Amazon Resource Name (ARN) for the webhook.
	//
	// This member is required.
	WebhookArn *string

	// The ID of the webhook.
	//
	// This member is required.
	WebhookId *string

	// The URL of the webhook.
	//
	// This member is required.
	WebhookUrl *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
