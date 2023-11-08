// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object that represents the custom domain name association.
type Association struct {

	// The custom domain name’s certificate Amazon resource name (ARN).
	CustomDomainCertificateArn *string

	// The expiration time for the certificate.
	CustomDomainCertificateExpiryTime *time.Time

	// The custom domain name associated with the workgroup.
	CustomDomainName *string

	// The name of the workgroup associated with the database.
	WorkgroupName *string

	noSmithyDocumentSerde
}

// An array of key-value pairs to set for advanced control over Amazon Redshift
// Serverless.
type ConfigParameter struct {

	// The key of the parameter. The options are auto_mv , datestyle ,
	// enable_case_sensitivity_identifier , enable_user_activity_logging , query_group
	// , search_path , and query monitoring metrics that let you define performance
	// boundaries. For more information about query monitoring rules and available
	// metrics, see Query monitoring metrics for Amazon Redshift Serverless (https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless)
	// .
	ParameterKey *string

	// The value of the parameter to set.
	ParameterValue *string

	noSmithyDocumentSerde
}

// The VPC endpoint object.
type Endpoint struct {

	// The DNS address of the VPC endpoint.
	Address *string

	// The port that Amazon Redshift Serverless listens on.
	Port *int32

	// An array of VpcEndpoint objects.
	VpcEndpoints []VpcEndpoint

	noSmithyDocumentSerde
}

// Information about an Amazon Redshift Serverless VPC endpoint.
type EndpointAccess struct {

	// The DNS address of the endpoint.
	Address *string

	// The Amazon Resource Name (ARN) of the VPC endpoint.
	EndpointArn *string

	// The time that the endpoint was created.
	EndpointCreateTime *time.Time

	// The name of the VPC endpoint.
	EndpointName *string

	// The status of the VPC endpoint.
	EndpointStatus *string

	// The port number on which Amazon Redshift Serverless accepts incoming
	// connections.
	Port *int32

	// The unique identifier of subnets where Amazon Redshift Serverless choose to
	// deploy the VPC endpoint.
	SubnetIds []string

	// The connection endpoint for connecting to Amazon Redshift Serverless.
	VpcEndpoint *VpcEndpoint

	// The security groups associated with the endpoint.
	VpcSecurityGroups []VpcSecurityGroupMembership

	// The name of the workgroup associated with the endpoint.
	WorkgroupName *string

	noSmithyDocumentSerde
}

// A collection of database objects and users.
type Namespace struct {

	// The Amazon Resource Name (ARN) for the namespace's admin user credentials
	// secret.
	AdminPasswordSecretArn *string

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// namespace's admin credentials secret.
	AdminPasswordSecretKmsKeyId *string

	// The username of the administrator for the first database created in the
	// namespace.
	AdminUsername *string

	// The date of when the namespace was created.
	CreationDate *time.Time

	// The name of the first database created in the namespace.
	DbName *string

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the
	// namespace.
	DefaultIamRoleArn *string

	// A list of IAM roles to associate with the namespace.
	IamRoles []string

	// The ID of the Amazon Web Services Key Management Service key used to encrypt
	// your data.
	KmsKeyId *string

	// The types of logs the namespace can export. Available export types are User
	// log, Connection log, and User activity log.
	LogExports []LogExport

	// The Amazon Resource Name (ARN) associated with a namespace.
	NamespaceArn *string

	// The unique identifier of a namespace.
	NamespaceId *string

	// The name of the namespace. Must be between 3-64 alphanumeric characters in
	// lowercase, and it cannot be a reserved word. A list of reserved words can be
	// found in Reserved Words (https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html)
	// in the Amazon Redshift Database Developer Guide.
	NamespaceName *string

	// The status of the namespace.
	Status NamespaceStatus

	noSmithyDocumentSerde
}

// Contains information about a network interface in an Amazon Redshift Serverless
// managed VPC endpoint.
type NetworkInterface struct {

	// The availability Zone.
	AvailabilityZone *string

	// The unique identifier of the network interface.
	NetworkInterfaceId *string

	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress *string

	// The unique identifier of the subnet.
	SubnetId *string

	noSmithyDocumentSerde
}

// The automatically created recovery point of a namespace. Recovery points are
// created every 30 minutes and kept for 24 hours.
type RecoveryPoint struct {

	// The Amazon Resource Name (ARN) of the namespace the recovery point is
	// associated with.
	NamespaceArn *string

	// The name of the namespace the recovery point is associated with.
	NamespaceName *string

	// The time the recovery point is created.
	RecoveryPointCreateTime *time.Time

	// The unique identifier of the recovery point.
	RecoveryPointId *string

	// The total size of the data in the recovery point in megabytes.
	TotalSizeInMegaBytes *float64

	// The name of the workgroup the recovery point is associated with.
	WorkgroupName *string

	noSmithyDocumentSerde
}

// The resource policy object. Currently, you can use policies to share snapshots
// across Amazon Web Services accounts.
type ResourcePolicy struct {

	// The resource policy.
	Policy *string

	// The Amazon Resource Name (ARN) of the policy.
	ResourceArn *string

	noSmithyDocumentSerde
}

// A snapshot object that contains databases.
type Snapshot struct {

	// All of the Amazon Web Services accounts that have access to restore a snapshot
	// to a provisioned cluster.
	AccountsWithProvisionedRestoreAccess []string

	// All of the Amazon Web Services accounts that have access to restore a snapshot
	// to a namespace.
	AccountsWithRestoreAccess []string

	// The size of the incremental backup in megabytes.
	ActualIncrementalBackupSizeInMegaBytes *float64

	// The Amazon Resource Name (ARN) for the namespace's admin user credentials
	// secret.
	AdminPasswordSecretArn *string

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// namespace's admin credentials secret.
	AdminPasswordSecretKmsKeyId *string

	// The username of the database within a snapshot.
	AdminUsername *string

	// The size in megabytes of the data that has been backed up to a snapshot.
	BackupProgressInMegaBytes *float64

	// The rate at which data is backed up into a snapshot in megabytes per second.
	CurrentBackupRateInMegaBytesPerSecond *float64

	// The amount of time it took to back up data into a snapshot.
	ElapsedTimeInSeconds *int64

	// The estimated amount of seconds until the snapshot completes backup.
	EstimatedSecondsToCompletion *int64

	// The unique identifier of the KMS key used to encrypt the snapshot.
	KmsKeyId *string

	// The Amazon Resource Name (ARN) of the namespace the snapshot was created from.
	NamespaceArn *string

	// The name of the namepsace.
	NamespaceName *string

	// The owner Amazon Web Services; account of the snapshot.
	OwnerAccount *string

	// The Amazon Resource Name (ARN) of the snapshot.
	SnapshotArn *string

	// The timestamp of when the snapshot was created.
	SnapshotCreateTime *time.Time

	// The name of the snapshot.
	SnapshotName *string

	// The amount of days until the snapshot is deleted.
	SnapshotRemainingDays *int32

	// The period of time, in days, of how long the snapshot is retained.
	SnapshotRetentionPeriod *int32

	// The timestamp of when data within the snapshot started getting retained.
	SnapshotRetentionStartTime *time.Time

	// The status of the snapshot.
	Status SnapshotStatus

	// The total size, in megabytes, of how big the snapshot is.
	TotalBackupSizeInMegaBytes *float64

	noSmithyDocumentSerde
}

// Contains information about a table restore request.
type TableRestoreStatus struct {

	// A description of the status of the table restore request. Status values include
	// SUCCEEDED , FAILED , CANCELED , PENDING , IN_PROGRESS .
	Message *string

	// The namespace of the table being restored from.
	NamespaceName *string

	// The name of the table to create from the restore operation.
	NewTableName *string

	// The amount of data restored to the new table so far, in megabytes (MB).
	ProgressInMegaBytes *int64

	// The time that the table restore request was made, in Universal Coordinated Time
	// (UTC).
	RequestTime *time.Time

	// The name of the snapshot being restored from.
	SnapshotName *string

	// The name of the source database being restored from.
	SourceDatabaseName *string

	// The name of the source schema being restored from.
	SourceSchemaName *string

	// The name of the source table being restored from.
	SourceTableName *string

	// A value that describes the current state of the table restore request. Possible
	// values include SUCCEEDED , FAILED , CANCELED , PENDING , IN_PROGRESS .
	Status *string

	// The ID of the RestoreTableFromSnapshot request.
	TableRestoreRequestId *string

	// The name of the database to restore to.
	TargetDatabaseName *string

	// The name of the schema to restore to.
	TargetSchemaName *string

	// The total amount of data to restore to the new table, in megabytes (MB).
	TotalDataInMegaBytes *int64

	// The name of the workgroup being restored from.
	WorkgroupName *string

	noSmithyDocumentSerde
}

// A map of key-value pairs.
type Tag struct {

	// The key to use in the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The usage limit object.
type UsageLimit struct {

	// The limit amount. If time-based, this amount is in RPUs consumed per hour. If
	// data-based, this amount is in terabytes (TB). The value must be a positive
	// number.
	Amount *int64

	// The action that Amazon Redshift Serverless takes when the limit is reached.
	BreachAction UsageLimitBreachAction

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly.
	Period UsageLimitPeriod

	// The Amazon Resource Name (ARN) that identifies the Amazon Redshift Serverless
	// resource.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the resource associated with the usage limit.
	UsageLimitArn *string

	// The identifier of the usage limit.
	UsageLimitId *string

	// The Amazon Redshift Serverless feature to limit.
	UsageType UsageLimitUsageType

	noSmithyDocumentSerde
}

// The connection endpoint for connecting to Amazon Redshift Serverless through
// the proxy.
type VpcEndpoint struct {

	// One or more network interfaces of the endpoint. Also known as an interface
	// endpoint.
	NetworkInterfaces []NetworkInterface

	// The connection endpoint ID for connecting to Amazon Redshift Serverless.
	VpcEndpointId *string

	// The VPC identifier that the endpoint is associated with.
	VpcId *string

	noSmithyDocumentSerde
}

// Describes the members of a VPC security group.
type VpcSecurityGroupMembership struct {

	// The status of the VPC security group.
	Status *string

	// The unique identifier of the VPC security group.
	VpcSecurityGroupId *string

	noSmithyDocumentSerde
}

// The collection of computing resources from which an endpoint is created.
type Workgroup struct {

	// The base data warehouse capacity of the workgroup in Redshift Processing Units
	// (RPUs).
	BaseCapacity *int32

	// An array of parameters to set for advanced control over a database. The options
	// are auto_mv , datestyle , enable_case_sensitivity_identifier ,
	// enable_user_activity_logging , query_group , , search_path , and query
	// monitoring metrics that let you define performance boundaries. For more
	// information about query monitoring rules and available metrics, see Query
	// monitoring metrics for Amazon Redshift Serverless (https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless)
	// .
	ConfigParameters []ConfigParameter

	// The creation date of the workgroup.
	CreationDate *time.Time

	// The custom domain name’s certificate Amazon resource name (ARN).
	CustomDomainCertificateArn *string

	// The expiration time for the certificate.
	CustomDomainCertificateExpiryTime *time.Time

	// The custom domain name associated with the workgroup.
	CustomDomainName *string

	// The endpoint that is created from the workgroup.
	Endpoint *Endpoint

	// The value that specifies whether to enable enhanced virtual private cloud (VPC)
	// routing, which forces Amazon Redshift Serverless to route traffic through your
	// VPC.
	EnhancedVpcRouting *bool

	// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve
	// queries. The max capacity is specified in RPUs.
	MaxCapacity *int32

	// The namespace the workgroup is associated with.
	NamespaceName *string

	// The patch version of your Amazon Redshift Serverless workgroup. For more
	// information about patch versions, see Cluster versions for Amazon Redshift (https://docs.aws.amazon.com/redshift/latest/mgmt/cluster-versions.html)
	// .
	PatchVersion *string

	// The custom port to use when connecting to a workgroup. Valid port ranges are
	// 5431-5455 and 8191-8215. The default is 5439.
	Port *int32

	// A value that specifies whether the workgroup can be accessible from a public
	// network
	PubliclyAccessible *bool

	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds []string

	// The status of the workgroup.
	Status WorkgroupStatus

	// An array of subnet IDs the workgroup is associated with.
	SubnetIds []string

	// The Amazon Resource Name (ARN) that links to the workgroup.
	WorkgroupArn *string

	// The unique identifier of the workgroup.
	WorkgroupId *string

	// The name of the workgroup.
	WorkgroupName *string

	// The Amazon Redshift Serverless version of your workgroup. For more information
	// about Amazon Redshift Serverless versions, see Cluster versions for Amazon
	// Redshift (https://docs.aws.amazon.com/redshift/latest/mgmt/cluster-versions.html)
	// .
	WorkgroupVersion *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
