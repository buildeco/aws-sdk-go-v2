// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A resource that is used to connect third-party source providers with services
// like CodePipeline. Note: A connection created through CloudFormation, the CLI,
// or the SDK is in `PENDING` status by default. You can make its status
// `AVAILABLE` by updating the connection in the console.
type Connection struct {

	// The Amazon Resource Name (ARN) of the connection. The ARN is used as the
	// connection reference when the connection is shared between Amazon Web Services.
	// The ARN is never reused if the connection is deleted.
	ConnectionArn *string

	// The name of the connection. Connection names must be unique in an Amazon Web
	// Services account.
	ConnectionName *string

	// The current status of the connection.
	ConnectionStatus ConnectionStatus

	// The Amazon Resource Name (ARN) of the host associated with the connection.
	HostArn *string

	// The identifier of the external provider where your third-party code repository
	// is configured. For Bitbucket, this is the account ID of the owner of the
	// Bitbucket repository.
	OwnerAccountId *string

	// The name of the external provider where your third-party code repository is
	// configured.
	ProviderType ProviderType

	noSmithyDocumentSerde
}

// A resource that represents the infrastructure where a third-party provider is
// installed. The host is used when you create connections to an installed
// third-party provider type, such as GitHub Enterprise Server. You create one host
// for all connections to that provider. A host created through the CLI or the SDK
// is in `PENDING` status by default. You can make its status `AVAILABLE` by
// setting up the host in the console.
type Host struct {

	// The Amazon Resource Name (ARN) of the host.
	HostArn *string

	// The name of the host.
	Name *string

	// The endpoint of the infrastructure where your provider type is installed.
	ProviderEndpoint *string

	// The name of the installed provider to be associated with your connection. The
	// host resource represents the infrastructure where your provider type is
	// installed. The valid provider type is GitHub Enterprise Server.
	ProviderType ProviderType

	// The status of the host, such as PENDING, AVAILABLE, VPC_CONFIG_DELETING,
	// VPC_CONFIG_INITIALIZING, and VPC_CONFIG_FAILED_INITIALIZATION.
	Status *string

	// The status description for the host.
	StatusMessage *string

	// The VPC configuration provisioned for the host.
	VpcConfiguration *VpcConfiguration

	noSmithyDocumentSerde
}

// Information about the repository link resource, such as the repository link
// ARN, the associated connection ARN, encryption key ARN, and owner ID.
type RepositoryLinkInfo struct {

	// The Amazon Resource Name (ARN) of the connection associated with the repository
	// link.
	//
	// This member is required.
	ConnectionArn *string

	// The owner ID for the repository associated with the repository link, such as
	// the owner ID in GitHub.
	//
	// This member is required.
	OwnerId *string

	// The provider type for the connection, such as GitHub, associated with the
	// repository link.
	//
	// This member is required.
	ProviderType ProviderType

	// The Amazon Resource Name (ARN) of the repository link.
	//
	// This member is required.
	RepositoryLinkArn *string

	// The ID of the repository link.
	//
	// This member is required.
	RepositoryLinkId *string

	// The name of the repository associated with the repository link.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Resource Name (ARN) of the encryption key for the repository
	// associated with the repository link.
	EncryptionKeyArn *string

	noSmithyDocumentSerde
}

// Information about a repository sync attempt for a repository with a sync
// configuration.
type RepositorySyncAttempt struct {

	// The events associated with a specific sync attempt.
	//
	// This member is required.
	Events []RepositorySyncEvent

	// The start time of a specific sync attempt.
	//
	// This member is required.
	StartedAt *time.Time

	// The status of a specific sync attempt. The following are valid statuses:
	//   - INITIATED - A repository sync attempt has been created and will begin soon.
	//   - IN_PROGRESS - A repository sync attempt has started and work is being done
	//   to reconcile the branch.
	//   - SUCCEEDED - The repository sync attempt has completed successfully.
	//   - FAILED - The repository sync attempt has failed.
	//   - QUEUED - The repository sync attempt didn't execute and was queued.
	//
	// This member is required.
	Status RepositorySyncStatus

	noSmithyDocumentSerde
}

// The definition for a repository with a sync configuration.
type RepositorySyncDefinition struct {

	// The branch specified for a repository sync definition.
	//
	// This member is required.
	Branch *string

	// The configuration file for a repository sync definition. This value comes from
	// creating or updating the config-file field of a sync-configuration .
	//
	// This member is required.
	Directory *string

	// The parent resource specified for a repository sync definition.
	//
	// This member is required.
	Parent *string

	// The target resource specified for a repository sync definition. In some cases,
	// such as CFN_STACK_SYNC, the parent and target resource are the same.
	//
	// This member is required.
	Target *string

	noSmithyDocumentSerde
}

// Information about a repository sync event.
type RepositorySyncEvent struct {

	// A description of a repository sync event.
	//
	// This member is required.
	Event *string

	// The time that a repository sync event occurred.
	//
	// This member is required.
	Time *time.Time

	// The event type for a repository sync event.
	//
	// This member is required.
	Type *string

	// The ID for a repository sync event.
	ExternalId *string

	noSmithyDocumentSerde
}

// Information about a resource sync attempt.
type ResourceSyncAttempt struct {

	// The events related to a resource sync attempt.
	//
	// This member is required.
	Events []ResourceSyncEvent

	// The current state of the resource as defined in the resource's config-file in
	// the linked repository.
	//
	// This member is required.
	InitialRevision *Revision

	// The start time for a resource sync attempt.
	//
	// This member is required.
	StartedAt *time.Time

	// The status for a resource sync attempt. The follow are valid statuses:
	//   - SYNC-INITIATED - A resource sync attempt has been created and will begin
	//   soon.
	//   - SYNCING - Syncing has started and work is being done to reconcile state.
	//   - SYNCED - Syncing has completed successfully.
	//   - SYNC_FAILED - A resource sync attempt has failed.
	//
	// This member is required.
	Status ResourceSyncStatus

	// The name of the Amazon Web Services resource that is attempted to be
	// synchronized.
	//
	// This member is required.
	Target *string

	// The desired state of the resource as defined in the resource's config-file in
	// the linked repository. Git sync attempts to update the resource to this state.
	//
	// This member is required.
	TargetRevision *Revision

	noSmithyDocumentSerde
}

// Information about a resource sync event for the resource associated with a sync
// configuration.
type ResourceSyncEvent struct {

	// The event for a resource sync event.
	//
	// This member is required.
	Event *string

	// The time that a resource sync event occurred.
	//
	// This member is required.
	Time *time.Time

	// The type of resource sync event.
	//
	// This member is required.
	Type *string

	// The ID for a resource sync event.
	ExternalId *string

	noSmithyDocumentSerde
}

// Information about the revision for a specific sync event, such as the branch,
// owner ID, and name of the repository.
type Revision struct {

	// The branch name for a specific revision.
	//
	// This member is required.
	Branch *string

	// The directory, if any, for a specific revision.
	//
	// This member is required.
	Directory *string

	// The owner ID for a specific revision, such as the GitHub owner ID for a GitHub
	// repository.
	//
	// This member is required.
	OwnerId *string

	// The provider type for a revision, such as GitHub.
	//
	// This member is required.
	ProviderType ProviderType

	// The repository name for a specific revision.
	//
	// This member is required.
	RepositoryName *string

	// The SHA, such as the commit ID, for a specific revision.
	//
	// This member is required.
	Sha *string

	noSmithyDocumentSerde
}

// Information about a blocker for a sync event.
type SyncBlocker struct {

	// The creation time for a specific sync blocker.
	//
	// This member is required.
	CreatedAt *time.Time

	// The provided reason for a specific sync blocker.
	//
	// This member is required.
	CreatedReason *string

	// The ID for a specific sync blocker.
	//
	// This member is required.
	Id *string

	// The status for a specific sync blocker.
	//
	// This member is required.
	Status BlockerStatus

	// The sync blocker type.
	//
	// This member is required.
	Type BlockerType

	// The contexts for a specific sync blocker.
	Contexts []SyncBlockerContext

	// The time that a specific sync blocker was resolved.
	ResolvedAt *time.Time

	// The resolved reason for a specific sync blocker.
	ResolvedReason *string

	noSmithyDocumentSerde
}

// The context for a specific sync blocker.
type SyncBlockerContext struct {

	// The key provided for a context key-value pair for a specific sync blocker.
	//
	// This member is required.
	Key *string

	// The value provided for a context key-value pair for a specific sync blocker.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// A summary for sync blockers.
type SyncBlockerSummary struct {

	// The resource name for sync blocker summary.
	//
	// This member is required.
	ResourceName *string

	// The latest events for a sync blocker summary.
	LatestBlockers []SyncBlocker

	// The parent resource name for a sync blocker summary.
	ParentResourceName *string

	noSmithyDocumentSerde
}

// Information, such as repository, branch, provider, and resource names for a
// specific sync configuration.
type SyncConfiguration struct {

	// The branch associated with a specific sync configuration.
	//
	// This member is required.
	Branch *string

	// The owner ID for the repository associated with a specific sync configuration,
	// such as the owner ID in GitHub.
	//
	// This member is required.
	OwnerId *string

	// The connection provider type associated with a specific sync configuration,
	// such as GitHub.
	//
	// This member is required.
	ProviderType ProviderType

	// The ID of the repository link associated with a specific sync configuration.
	//
	// This member is required.
	RepositoryLinkId *string

	// The name of the repository associated with a specific sync configuration.
	//
	// This member is required.
	RepositoryName *string

	// The name of the connection resource associated with a specific sync
	// configuration.
	//
	// This member is required.
	ResourceName *string

	// The Amazon Resource Name (ARN) of the IAM role associated with a specific sync
	// configuration.
	//
	// This member is required.
	RoleArn *string

	// The type of sync for a specific sync configuration.
	//
	// This member is required.
	SyncType SyncConfigurationType

	// The file path to the configuration file associated with a specific sync
	// configuration. The path should point to an actual file in the sync
	// configurations linked repository.
	ConfigFile *string

	noSmithyDocumentSerde
}

// A tag is a key-value pair that is used to manage the resource. This tag is
// available for use by Amazon Web Services services that support tags.
type Tag struct {

	// The tag's key.
	//
	// This member is required.
	Key *string

	// The tag's value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The VPC configuration provisioned for the host.
type VpcConfiguration struct {

	// The ID of the security group or security groups associated with the Amazon VPC
	// connected to the infrastructure where your provider type is installed.
	//
	// This member is required.
	SecurityGroupIds []string

	// The ID of the subnet or subnets associated with the Amazon VPC connected to the
	// infrastructure where your provider type is installed.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the Amazon VPC connected to the infrastructure where your provider
	// type is installed.
	//
	// This member is required.
	VpcId *string

	// The value of the Transport Layer Security (TLS) certificate associated with the
	// infrastructure where your provider type is installed.
	TlsCertificate *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
