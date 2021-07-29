// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains details about a package version asset.
type AssetSummary struct {

	// The name of the asset.
	//
	// This member is required.
	Name *string

	// The hashes of the asset.
	Hashes map[string]string

	// The size of the asset.
	Size *int64

	noSmithyDocumentSerde
}

// Information about a domain. A domain is a container for repositories. When you
// create a domain, it is empty until you add one or more repositories.
type DomainDescription struct {

	// The Amazon Resource Name (ARN) of the domain.
	Arn *string

	// The total size of all assets in the domain.
	AssetSizeBytes int64

	// A timestamp that represents the date and time the domain was created.
	CreatedTime *time.Time

	// The ARN of an AWS Key Management Service (AWS KMS) key associated with a domain.
	EncryptionKey *string

	// The name of the domain.
	Name *string

	// The AWS account ID that owns the domain.
	Owner *string

	// The number of repositories in the domain.
	RepositoryCount int32

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket that is used to store
	// package assets in the domain.
	S3BucketArn *string

	// The current status of a domain. The valid values are
	//
	// * Active
	//
	// * Deleted
	Status DomainStatus

	noSmithyDocumentSerde
}

// Information about a domain, including its name, Amazon Resource Name (ARN), and
// status. The ListDomains
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListDomains.html)
// operation returns a list of DomainSummary objects.
type DomainSummary struct {

	// The ARN of the domain.
	Arn *string

	// A timestamp that contains the date and time the domain was created.
	CreatedTime *time.Time

	// The key used to encrypt the domain.
	EncryptionKey *string

	// The name of the domain.
	Name *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	Owner *string

	// A string that contains the status of the domain. The valid values are:
	//
	// *
	// Active
	//
	// * Deleted
	Status DomainStatus

	noSmithyDocumentSerde
}

// Details of the license data.
type LicenseInfo struct {

	// Name of the license.
	Name *string

	// The URL for license data.
	Url *string

	noSmithyDocumentSerde
}

// Details about a package dependency.
type PackageDependency struct {

	// The type of a package dependency. The possible values depend on the package
	// type. Example types are compile, runtime, and test for Maven packages, and dev,
	// prod, and optional for npm packages.
	DependencyType *string

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The name of the package that this package depends on.
	Package *string

	// The required version, or version range, of the package that this package depends
	// on. The version format is specific to the package type. For example, the
	// following are possible valid required versions: 1.2.3, ^2.3.4, or 4.x.
	VersionRequirement *string

	noSmithyDocumentSerde
}

// Details about a package, including its format, namespace, and name. The
// ListPackages
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackages.html)
// operation returns a list of PackageSummary objects.
type PackageSummary struct {

	// The format of the package. Valid values are:
	//
	// * npm
	//
	// * pypi
	//
	// * maven
	Format PackageFormat

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The name of the package.
	Package *string

	noSmithyDocumentSerde
}

// Details about a package version.
type PackageVersionDescription struct {

	// The name of the package that is displayed. The displayName varies depending on
	// the package version's format. For example, if an npm package is named ui, is in
	// the namespace vue, and has the format npm, then the displayName is @vue/ui.
	DisplayName *string

	// The format of the package version. The valid package formats are:
	//
	// * npm: A Node
	// Package Manager (npm) package.
	//
	// * pypi: A Python Package Index (PyPI)
	// package.
	//
	// * maven: A Maven package that contains compiled code in a
	// distributable format, such as a JAR file.
	Format PackageFormat

	// The homepage associated with the package.
	HomePage *string

	// Information about licenses associated with the package version.
	Licenses []LicenseInfo

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The name of the requested package.
	PackageName *string

	// A timestamp that contains the date and time the package version was published.
	PublishedTime *time.Time

	// The revision of the package version.
	Revision *string

	// The repository for the source code in the package version, or the source code
	// used to build it.
	SourceCodeRepository *string

	// A string that contains the status of the package version. It can be one of the
	// following:
	//
	// * Published
	//
	// * Unfinished
	//
	// * Unlisted
	//
	// * Archived
	//
	// * Disposed
	Status PackageVersionStatus

	// A summary of the package version. The summary is extracted from the package. The
	// information in and detail level of the summary depends on the package version's
	// format.
	Summary *string

	// The version of the package.
	Version *string

	noSmithyDocumentSerde
}

// An error associated with package.
type PackageVersionError struct {

	// The error code associated with the error. Valid error codes are:
	//
	// *
	// ALREADY_EXISTS
	//
	// * MISMATCHED_REVISION
	//
	// * MISMATCHED_STATUS
	//
	// * NOT_ALLOWED
	//
	// *
	// NOT_FOUND
	//
	// * SKIPPED
	ErrorCode PackageVersionErrorCode

	// The error message associated with the error.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// Details about a package version, including its status, version, and revision.
// The ListPackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)
// operation returns a list of PackageVersionSummary objects.
type PackageVersionSummary struct {

	// A string that contains the status of the package version. It can be one of the
	// following:
	//
	// * Published
	//
	// * Unfinished
	//
	// * Unlisted
	//
	// * Archived
	//
	// * Disposed
	//
	// This member is required.
	Status PackageVersionStatus

	// Information about a package version.
	//
	// This member is required.
	Version *string

	// The revision associated with a package version.
	Revision *string

	noSmithyDocumentSerde
}

// The details of a repository stored in AWS CodeArtifact. A CodeArtifact
// repository contains a set of package versions, each of which maps to a set of
// assets. Repositories are polyglot—a single repository can contain packages of
// any supported type. Each repository exposes endpoints for fetching and
// publishing packages using tools like the npm CLI, the Maven CLI (mvn), and pip.
// You can create up to 100 repositories per AWS account.
type RepositoryDescription struct {

	// The 12-digit account number of the AWS account that manages the repository.
	AdministratorAccount *string

	// The Amazon Resource Name (ARN) of the repository.
	Arn *string

	// A text description of the repository.
	Description *string

	// The name of the domain that contains the repository.
	DomainName *string

	// The 12-digit account number of the AWS account that owns the domain that
	// contains the repository. It does not include dashes or spaces.
	DomainOwner *string

	// An array of external connections associated with the repository.
	ExternalConnections []RepositoryExternalConnectionInfo

	// The name of the repository.
	Name *string

	// A list of upstream repositories to associate with the repository. The order of
	// the upstream repositories in the list determines their priority order when AWS
	// CodeArtifact looks for a requested package version. For more information, see
	// Working with upstream repositories
	// (https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html).
	Upstreams []UpstreamRepositoryInfo

	noSmithyDocumentSerde
}

// Contains information about the external connection of a repository.
type RepositoryExternalConnectionInfo struct {

	// The name of the external connection associated with a repository.
	ExternalConnectionName *string

	// The package format associated with a repository's external connection. The valid
	// package formats are:
	//
	// * npm: A Node Package Manager (npm) package.
	//
	// * pypi: A
	// Python Package Index (PyPI) package.
	//
	// * maven: A Maven package that contains
	// compiled code in a distributable format, such as a JAR file.
	PackageFormat PackageFormat

	// The status of the external connection of a repository. There is one valid value,
	// Available.
	Status ExternalConnectionStatus

	noSmithyDocumentSerde
}

// Details about a repository, including its Amazon Resource Name (ARN),
// description, and domain information. The ListRepositories
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListRepositories.html)
// operation returns a list of RepositorySummary objects.
type RepositorySummary struct {

	// The AWS account ID that manages the repository.
	AdministratorAccount *string

	// The ARN of the repository.
	Arn *string

	// The description of the repository.
	Description *string

	// The name of the domain that contains the repository.
	DomainName *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The name of the repository.
	Name *string

	noSmithyDocumentSerde
}

// An AWS CodeArtifact resource policy that contains a resource ARN, document
// details, and a revision.
type ResourcePolicy struct {

	// The resource policy formatted in JSON.
	Document *string

	// The ARN of the resource associated with the resource policy
	ResourceArn *string

	// The current revision of the resource policy.
	Revision *string

	noSmithyDocumentSerde
}

// Contains the revision and status of a package version.
type SuccessfulPackageVersionInfo struct {

	// The revision of a package version.
	Revision *string

	// The status of a package version. Valid statuses are:
	//
	// * Published
	//
	// *
	// Unfinished
	//
	// * Unlisted
	//
	// * Archived
	//
	// * Disposed
	Status PackageVersionStatus

	noSmithyDocumentSerde
}

// A tag is a key-value pair that can be used to manage, search for, or filter
// resources in AWS CodeArtifact.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Information about an upstream repository. A list of UpstreamRepository objects
// is an input parameter to CreateRepository
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_CreateRepository.html)
// and UpdateRepository
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdateRepository.html).
type UpstreamRepository struct {

	// The name of an upstream repository.
	//
	// This member is required.
	RepositoryName *string

	noSmithyDocumentSerde
}

// Information about an upstream repository.
type UpstreamRepositoryInfo struct {

	// The name of an upstream repository.
	RepositoryName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
