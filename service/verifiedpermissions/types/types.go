// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains information about an action for a request for which an authorization
// decision is made. This data type is used as a request parameter to the
// IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// , BatchIsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorized.html)
// , and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations. Example: { "actionId": "<action name>", "actionType": "Action" }
type ActionIdentifier struct {

	// The ID of an action.
	//
	// This member is required.
	ActionId *string

	// The type of an action.
	//
	// This member is required.
	ActionType *string

	noSmithyDocumentSerde
}

// The value of an attribute. Contains information about the runtime context for a
// request for which an authorization decision is made. This data type is used as a
// member of the ContextDefinition (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ContextDefinition.html)
// structure which is uses as a request parameter for the IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// , BatchIsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorized.html)
// , and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations.
//
// The following types satisfy this interface:
//
//	AttributeValueMemberBoolean
//	AttributeValueMemberEntityIdentifier
//	AttributeValueMemberLong
//	AttributeValueMemberRecord
//	AttributeValueMemberSet
//	AttributeValueMemberString
type AttributeValue interface {
	isAttributeValue()
}

// An attribute value of Boolean (https://docs.cedarpolicy.com/policies/syntax-datatypes.html#boolean)
// type. Example: {"boolean": true}
type AttributeValueMemberBoolean struct {
	Value bool

	noSmithyDocumentSerde
}

func (*AttributeValueMemberBoolean) isAttributeValue() {}

// An attribute value of type EntityIdentifier (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_EntityIdentifier.html)
// . Example: "entityIdentifier": { "entityId": "<id>", "entityType": "<entity
// type>"}
type AttributeValueMemberEntityIdentifier struct {
	Value EntityIdentifier

	noSmithyDocumentSerde
}

func (*AttributeValueMemberEntityIdentifier) isAttributeValue() {}

// An attribute value of Long (https://docs.cedarpolicy.com/policies/syntax-datatypes.html#long)
// type. Example: {"long": 0}
type AttributeValueMemberLong struct {
	Value int64

	noSmithyDocumentSerde
}

func (*AttributeValueMemberLong) isAttributeValue() {}

// An attribute value of Record (https://docs.cedarpolicy.com/policies/syntax-datatypes.html#record)
// type. Example: {"record": { "keyName": {} } }
type AttributeValueMemberRecord struct {
	Value map[string]AttributeValue

	noSmithyDocumentSerde
}

func (*AttributeValueMemberRecord) isAttributeValue() {}

// An attribute value of Set (https://docs.cedarpolicy.com/policies/syntax-datatypes.html#set)
// type. Example: {"set": [ {} ] }
type AttributeValueMemberSet struct {
	Value []AttributeValue

	noSmithyDocumentSerde
}

func (*AttributeValueMemberSet) isAttributeValue() {}

// An attribute value of String (https://docs.cedarpolicy.com/policies/syntax-datatypes.html#string)
// type. Example: {"string": "abc"}
type AttributeValueMemberString struct {
	Value string

	noSmithyDocumentSerde
}

func (*AttributeValueMemberString) isAttributeValue() {}

// An authorization request that you include in a BatchIsAuthorized API request.
type BatchIsAuthorizedInputItem struct {

	// Specifies the requested action to be authorized. For example, is the principal
	// authorized to perform this action on the resource?
	Action *ActionIdentifier

	// Specifies additional context that can be used to make more granular
	// authorization decisions.
	Context ContextDefinition

	// Specifies the principal for which the authorization decision is to be made.
	Principal *EntityIdentifier

	// Specifies the resource for which the authorization decision is to be made.
	Resource *EntityIdentifier

	noSmithyDocumentSerde
}

// The decision, based on policy evaluation, from an individual authorization
// request in a BatchIsAuthorized API request.
type BatchIsAuthorizedOutputItem struct {

	// An authorization decision that indicates if the authorization request should be
	// allowed or denied.
	//
	// This member is required.
	Decision Decision

	// The list of determining policies used to make the authorization decision. For
	// example, if there are two matching policies, where one is a forbid and the other
	// is a permit, then the forbid policy will be the determining policy. In the case
	// of multiple matching permit policies then there would be multiple determining
	// policies. In the case that no policies match, and hence the response is DENY,
	// there would be no determining policies.
	//
	// This member is required.
	DeterminingPolicies []DeterminingPolicyItem

	// Errors that occurred while making an authorization decision, for example, a
	// policy references an Entity or entity Attribute that does not exist in the
	// slice.
	//
	// This member is required.
	Errors []EvaluationErrorItem

	// The authorization request that initiated the decision.
	//
	// This member is required.
	Request *BatchIsAuthorizedInputItem

	noSmithyDocumentSerde
}

// The configuration for an identity source that represents a connection to an
// Amazon Cognito user pool used as an identity provider for Verified Permissions.
// This data type is used as a field that is part of an Configuration (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_Configuration.html)
// structure that is used as a parameter to the Configuration (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_Configuration.html)
// . Example:
// "CognitoUserPoolConfiguration":{"UserPoolArn":"arn:aws:cognito-idp:us-east-1:123456789012:userpool/us-east-1_1a2b3c4d5","ClientIds":
// ["a1b2c3d4e5f6g7h8i9j0kalbmc"]}
type CognitoUserPoolConfiguration struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the Amazon Cognito user pool that contains the identities to be authorized.
	// Example: "UserPoolArn":
	// "arn:aws:cognito-idp:us-east-1:123456789012:userpool/us-east-1_1a2b3c4d5"
	//
	// This member is required.
	UserPoolArn *string

	// The unique application client IDs that are associated with the specified Amazon
	// Cognito user pool. Example: "ClientIds": ["&ExampleCogClientId;"]
	ClientIds []string

	noSmithyDocumentSerde
}

// Contains configuration information used when creating a new identity source. At
// this time, the only valid member of this structure is a Amazon Cognito user pool
// configuration. You must specify a userPoolArn , and optionally, a ClientId .
// This data type is used as a request parameter for the CreateIdentitySource (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html)
// operation.
//
// The following types satisfy this interface:
//
//	ConfigurationMemberCognitoUserPoolConfiguration
type Configuration interface {
	isConfiguration()
}

// Contains configuration details of a Amazon Cognito user pool that Verified
// Permissions can use as a source of authenticated identities as entities. It
// specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// of a Amazon Cognito user pool and one or more application client IDs. Example:
// "configuration":{"cognitoUserPoolConfiguration":{"userPoolArn":"arn:aws:cognito-idp:us-east-1:123456789012:userpool/us-east-1_1a2b3c4d5","clientIds":
// ["a1b2c3d4e5f6g7h8i9j0kalbmc"]}}
type ConfigurationMemberCognitoUserPoolConfiguration struct {
	Value CognitoUserPoolConfiguration

	noSmithyDocumentSerde
}

func (*ConfigurationMemberCognitoUserPoolConfiguration) isConfiguration() {}

// Contains additional details about the context of the request. Verified
// Permissions evaluates this information in an authorization request as part of
// the when and unless clauses in a policy. This data type is used as a request
// parameter for the IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// , BatchIsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorized.html)
// , and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations. Example:
// "context":{"contextMap":{"<KeyName1>":{"boolean":true},"<KeyName2>":{"long":1234}}}
//
// The following types satisfy this interface:
//
//	ContextDefinitionMemberContextMap
type ContextDefinition interface {
	isContextDefinition()
}

// An list of attributes that are needed to successfully evaluate an authorization
// request. Each attribute in this array must include a map of a data type and its
// value. Example:
// "contextMap":{"<KeyName1>":{"boolean":true},"<KeyName2>":{"long":1234}}
type ContextDefinitionMemberContextMap struct {
	Value map[string]AttributeValue

	noSmithyDocumentSerde
}

func (*ContextDefinitionMemberContextMap) isContextDefinition() {}

// Contains information about one of the policies that determined an authorization
// decision. This data type is used as an element in a response parameter for the
// IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// , BatchIsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorized.html)
// , and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations. Example:
// "determiningPolicies":[{"policyId":"SPEXAMPLEabcdefg111111"}]
type DeterminingPolicyItem struct {

	// The Id of a policy that determined to an authorization decision. Example:
	// "policyId":"SPEXAMPLEabcdefg111111"
	//
	// This member is required.
	PolicyId *string

	noSmithyDocumentSerde
}

// Contains the list of entities to be considered during an authorization request.
// This includes all principals, resources, and actions required to successfully
// evaluate the request. This data type is used as a field in the response
// parameter for the IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations.
//
// The following types satisfy this interface:
//
//	EntitiesDefinitionMemberEntityList
type EntitiesDefinition interface {
	isEntitiesDefinition()
}

// An array of entities that are needed to successfully evaluate an authorization
// request. Each entity in this array must include an identifier for the entity,
// the attributes of the entity, and a list of any parent entities.
type EntitiesDefinitionMemberEntityList struct {
	Value []EntityItem

	noSmithyDocumentSerde
}

func (*EntitiesDefinitionMemberEntityList) isEntitiesDefinition() {}

// Contains the identifier of an entity, including its ID and type. This data type
// is used as a request parameter for IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// operation, and as a response parameter for the CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// , GetPolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_GetPolicy.html)
// , and UpdatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicy.html)
// operations. Example: {"entityId":"string","entityType":"string"}
type EntityIdentifier struct {

	// The identifier of an entity. "entityId":"identifier"
	//
	// This member is required.
	EntityId *string

	// The type of an entity. Example: "entityType":"typeName"
	//
	// This member is required.
	EntityType *string

	noSmithyDocumentSerde
}

// Contains information about an entity that can be referenced in a Cedar policy.
// This data type is used as one of the fields in the EntitiesDefinition (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_EntitiesDefinition.html)
// structure. { "identifier": { "entityType": "Photo", "entityId":
// "VacationPhoto94.jpg" }, "attributes": {}, "parents": [ { "entityType": "Album",
// "entityId": "alice_folder" } ] }
type EntityItem struct {

	// The identifier of the entity.
	//
	// This member is required.
	Identifier *EntityIdentifier

	// A list of attributes for the entity.
	Attributes map[string]AttributeValue

	// The parents in the hierarchy that contains the entity.
	Parents []EntityIdentifier

	noSmithyDocumentSerde
}

// Contains information about a principal or resource that can be referenced in a
// Cedar policy. This data type is used as part of the PolicyFilter (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_PolicyFilter.html)
// structure that is used as a request parameter for the ListPolicies (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicies.html)
// operation..
//
// The following types satisfy this interface:
//
//	EntityReferenceMemberIdentifier
//	EntityReferenceMemberUnspecified
type EntityReference interface {
	isEntityReference()
}

// The identifier of the entity. It can consist of either an EntityType and
// EntityId, a principal, or a resource.
type EntityReferenceMemberIdentifier struct {
	Value EntityIdentifier

	noSmithyDocumentSerde
}

func (*EntityReferenceMemberIdentifier) isEntityReference() {}

// Used to indicate that a principal or resource is not specified. This can be
// used to search for policies that are not associated with a specific principal or
// resource.
type EntityReferenceMemberUnspecified struct {
	Value bool

	noSmithyDocumentSerde
}

func (*EntityReferenceMemberUnspecified) isEntityReference() {}

// Contains a description of an evaluation error. This data type is a response
// parameter of the IsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorized.html)
// , BatchIsAuthorized (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_BatchIsAuthorized.html)
// , and IsAuthorizedWithToken (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_IsAuthorizedWithToken.html)
// operations.
type EvaluationErrorItem struct {

	// The error description.
	//
	// This member is required.
	ErrorDescription *string

	noSmithyDocumentSerde
}

// A structure that contains configuration of the identity source. This data type
// is used as a response parameter for the CreateIdentitySource (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html)
// operation.
type IdentitySourceDetails struct {

	// The application client IDs associated with the specified Amazon Cognito user
	// pool that are enabled for this identity source.
	ClientIds []string

	// The well-known URL that points to this user pool's OIDC discovery endpoint.
	// This is a URL string in the following format. This URL replaces the placeholders
	// for both the Amazon Web Services Region and the user pool identifier with those
	// appropriate for this user pool.
	// https://cognito-idp.<region>.amazonaws.com/<user-pool-id>/.well-known/openid-configuration
	DiscoveryUrl *string

	// A string that identifies the type of OIDC service represented by this identity
	// source. At this time, the only valid value is cognito .
	OpenIdIssuer OpenIdIssuer

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the Amazon Cognito user pool whose identities are accessible to this Verified
	// Permissions policy store.
	UserPoolArn *string

	noSmithyDocumentSerde
}

// A structure that defines characteristics of an identity source that you can use
// to filter. This data type is used as a request parameter for the
// ListIdentityStores (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListIdentityStores.html)
// operation.
type IdentitySourceFilter struct {

	// The Cedar entity type of the principals returned by the identity provider (IdP)
	// associated with this identity source.
	PrincipalEntityType *string

	noSmithyDocumentSerde
}

// A structure that defines an identity source. This data type is used as a
// request parameter for the ListIdentityStores (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListIdentityStores.html)
// operation.
type IdentitySourceItem struct {

	// The date and time the identity source was originally created.
	//
	// This member is required.
	CreatedDate *time.Time

	// A structure that contains the details of the associated identity provider (IdP).
	//
	// This member is required.
	Details *IdentitySourceItemDetails

	// The unique identifier of the identity source.
	//
	// This member is required.
	IdentitySourceId *string

	// The date and time the identity source was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The identifier of the policy store that contains the identity source.
	//
	// This member is required.
	PolicyStoreId *string

	// The Cedar entity type of the principals returned from the IdP associated with
	// this identity source.
	//
	// This member is required.
	PrincipalEntityType *string

	noSmithyDocumentSerde
}

// A structure that contains configuration of the identity source. This data type
// is used as a response parameter for the CreateIdentitySource (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html)
// operation.
type IdentitySourceItemDetails struct {

	// The application client IDs associated with the specified Amazon Cognito user
	// pool that are enabled for this identity source.
	ClientIds []string

	// The well-known URL that points to this user pool's OIDC discovery endpoint.
	// This is a URL string in the following format. This URL replaces the placeholders
	// for both the Amazon Web Services Region and the user pool identifier with those
	// appropriate for this user pool.
	// https://cognito-idp.<region>.amazonaws.com/<user-pool-id>/.well-known/openid-configuration
	DiscoveryUrl *string

	// A string that identifies the type of OIDC service represented by this identity
	// source. At this time, the only valid value is cognito .
	OpenIdIssuer OpenIdIssuer

	// The Amazon Cognito user pool whose identities are accessible to this Verified
	// Permissions policy store.
	UserPoolArn *string

	noSmithyDocumentSerde
}

// A structure that contains the details for a Cedar policy definition. It
// includes the policy type, a description, and a policy body. This is a top level
// data type used to create a policy. This data type is used as a request parameter
// for the CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// operation. This structure must always have either an static or a templateLinked
// element.
//
// The following types satisfy this interface:
//
//	PolicyDefinitionMemberStatic
//	PolicyDefinitionMemberTemplateLinked
type PolicyDefinition interface {
	isPolicyDefinition()
}

// A structure that describes a static policy. An static policy doesn't use a
// template or allow placeholders for entities.
type PolicyDefinitionMemberStatic struct {
	Value StaticPolicyDefinition

	noSmithyDocumentSerde
}

func (*PolicyDefinitionMemberStatic) isPolicyDefinition() {}

// A structure that describes a policy that was instantiated from a template. The
// template can specify placeholders for principal and resource . When you use
// CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// to create a policy from a template, you specify the exact principal and resource
// to use for the instantiated policy.
type PolicyDefinitionMemberTemplateLinked struct {
	Value TemplateLinkedPolicyDefinition

	noSmithyDocumentSerde
}

func (*PolicyDefinitionMemberTemplateLinked) isPolicyDefinition() {}

// A structure that describes a policy definition. It must always have either an
// static or a templateLinked element. This data type is used as a response
// parameter for the GetPolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_GetPolicy.html)
// operation.
//
// The following types satisfy this interface:
//
//	PolicyDefinitionDetailMemberStatic
//	PolicyDefinitionDetailMemberTemplateLinked
type PolicyDefinitionDetail interface {
	isPolicyDefinitionDetail()
}

// Information about a static policy that wasn't created with a policy template.
type PolicyDefinitionDetailMemberStatic struct {
	Value StaticPolicyDefinitionDetail

	noSmithyDocumentSerde
}

func (*PolicyDefinitionDetailMemberStatic) isPolicyDefinitionDetail() {}

// Information about a template-linked policy that was created by instantiating a
// policy template.
type PolicyDefinitionDetailMemberTemplateLinked struct {
	Value TemplateLinkedPolicyDefinitionDetail

	noSmithyDocumentSerde
}

func (*PolicyDefinitionDetailMemberTemplateLinked) isPolicyDefinitionDetail() {}

// A structure that describes a PolicyDefinintion (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_PolicyDefinintion.html)
// . It will always have either an StaticPolicy or a TemplateLinkedPolicy element.
// This data type is used as a response parameter for the CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// and ListPolicies (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicies.html)
// operations.
//
// The following types satisfy this interface:
//
//	PolicyDefinitionItemMemberStatic
//	PolicyDefinitionItemMemberTemplateLinked
type PolicyDefinitionItem interface {
	isPolicyDefinitionItem()
}

// Information about a static policy that wasn't created with a policy template.
type PolicyDefinitionItemMemberStatic struct {
	Value StaticPolicyDefinitionItem

	noSmithyDocumentSerde
}

func (*PolicyDefinitionItemMemberStatic) isPolicyDefinitionItem() {}

// Information about a template-linked policy that was created by instantiating a
// policy template.
type PolicyDefinitionItemMemberTemplateLinked struct {
	Value TemplateLinkedPolicyDefinitionItem

	noSmithyDocumentSerde
}

func (*PolicyDefinitionItemMemberTemplateLinked) isPolicyDefinitionItem() {}

// Contains information about a filter to refine policies returned in a query.
// This data type is used as a response parameter for the ListPolicies (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicies.html)
// operation.
type PolicyFilter struct {

	// Filters the output to only template-linked policies that were instantiated from
	// the specified policy template.
	PolicyTemplateId *string

	// Filters the output to only policies of the specified type.
	PolicyType PolicyType

	// Filters the output to only policies that reference the specified principal.
	Principal EntityReference

	// Filters the output to only policies that reference the specified resource.
	Resource EntityReference

	noSmithyDocumentSerde
}

// Contains information about a policy. This data type is used as a response
// parameter for the ListPolicies (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicies.html)
// operation.
type PolicyItem struct {

	// The date and time the policy was created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The policy definition of an item in the list of policies returned.
	//
	// This member is required.
	Definition PolicyDefinitionItem

	// The date and time the policy was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The identifier of the policy you want information about.
	//
	// This member is required.
	PolicyId *string

	// The identifier of the PolicyStore where the policy you want information about
	// is stored.
	//
	// This member is required.
	PolicyStoreId *string

	// The type of the policy. This is one of the following values:
	//   - static
	//   - templateLinked
	//
	// This member is required.
	PolicyType PolicyType

	// The principal associated with the policy.
	Principal *EntityIdentifier

	// The resource associated with the policy.
	Resource *EntityIdentifier

	noSmithyDocumentSerde
}

// Contains information about a policy store. This data type is used as a response
// parameter for the ListPolicyStores (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicyStores.html)
// operation.
type PolicyStoreItem struct {

	// The Amazon Resource Name (ARN) of the policy store.
	//
	// This member is required.
	Arn *string

	// The date and time the policy was created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The unique identifier of the policy store.
	//
	// This member is required.
	PolicyStoreId *string

	noSmithyDocumentSerde
}

// Contains details about a policy template This data type is used as a response
// parameter for the ListPolicyTemplates (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_ListPolicyTemplates.html)
// operation.
type PolicyTemplateItem struct {

	// The date and time that the policy template was created.
	//
	// This member is required.
	CreatedDate *time.Time

	// The date and time that the policy template was most recently updated.
	//
	// This member is required.
	LastUpdatedDate *time.Time

	// The unique identifier of the policy store that contains the template.
	//
	// This member is required.
	PolicyStoreId *string

	// The unique identifier of the policy template.
	//
	// This member is required.
	PolicyTemplateId *string

	// The description attached to the policy template.
	Description *string

	noSmithyDocumentSerde
}

// Contains information about a resource conflict.
type ResourceConflict struct {

	// The unique identifier of the resource involved in a conflict.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource involved in a conflict.
	//
	// This member is required.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Contains a list of principal types, resource types, and actions that can be
// specified in policies stored in the same policy store. If the validation mode
// for the policy store is set to STRICT , then policies that can't be validated by
// this schema are rejected by Verified Permissions and can't be stored in the
// policy store.
//
// The following types satisfy this interface:
//
//	SchemaDefinitionMemberCedarJson
type SchemaDefinition interface {
	isSchemaDefinition()
}

// A JSON string representation of the schema supported by applications that use
// this policy store. For more information, see Policy store schema (https://docs.aws.amazon.com/verifiedpermissions/latest/userguide/schema.html)
// in the Amazon Verified Permissions User Guide.
type SchemaDefinitionMemberCedarJson struct {
	Value string

	noSmithyDocumentSerde
}

func (*SchemaDefinitionMemberCedarJson) isSchemaDefinition() {}

// Contains information about a static policy. This data type is used as a field
// that is part of the PolicyDefinitionDetail (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_PolicyDefinitionDetail.html)
// type.
type StaticPolicyDefinition struct {

	// The policy content of the static policy, written in the Cedar policy language.
	//
	// This member is required.
	Statement *string

	// The description of the static policy.
	Description *string

	noSmithyDocumentSerde
}

// A structure that contains details about a static policy. It includes the
// description and policy body. This data type is used within a PolicyDefinition (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_PolicyDefinition.html)
// structure as part of a request parameter for the CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// operation.
type StaticPolicyDefinitionDetail struct {

	// The content of the static policy written in the Cedar policy language.
	//
	// This member is required.
	Statement *string

	// A description of the static policy.
	Description *string

	noSmithyDocumentSerde
}

// A structure that contains details about a static policy. It includes the
// description and policy statement. This data type is used within a
// PolicyDefinition (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_PolicyDefinition.html)
// structure as part of a request parameter for the CreatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicy.html)
// operation.
type StaticPolicyDefinitionItem struct {

	// A description of the static policy.
	Description *string

	noSmithyDocumentSerde
}

// Contains information about a policy created by instantiating a policy template.
type TemplateLinkedPolicyDefinition struct {

	// The unique identifier of the policy template used to create this policy.
	//
	// This member is required.
	PolicyTemplateId *string

	// The principal associated with this template-linked policy. Verified Permissions
	// substitutes this principal for the ?principal placeholder in the policy
	// template when it evaluates an authorization request.
	Principal *EntityIdentifier

	// The resource associated with this template-linked policy. Verified Permissions
	// substitutes this resource for the ?resource placeholder in the policy template
	// when it evaluates an authorization request.
	Resource *EntityIdentifier

	noSmithyDocumentSerde
}

// Contains information about a policy that was created by instantiating a policy
// template. This
type TemplateLinkedPolicyDefinitionDetail struct {

	// The unique identifier of the policy template used to create this policy.
	//
	// This member is required.
	PolicyTemplateId *string

	// The principal associated with this template-linked policy. Verified Permissions
	// substitutes this principal for the ?principal placeholder in the policy
	// template when it evaluates an authorization request.
	Principal *EntityIdentifier

	// The resource associated with this template-linked policy. Verified Permissions
	// substitutes this resource for the ?resource placeholder in the policy template
	// when it evaluates an authorization request.
	Resource *EntityIdentifier

	noSmithyDocumentSerde
}

// Contains information about a policy created by instantiating a policy template.
// This
type TemplateLinkedPolicyDefinitionItem struct {

	// The unique identifier of the policy template used to create this policy.
	//
	// This member is required.
	PolicyTemplateId *string

	// The principal associated with this template-linked policy. Verified Permissions
	// substitutes this principal for the ?principal placeholder in the policy
	// template when it evaluates an authorization request.
	Principal *EntityIdentifier

	// The resource associated with this template-linked policy. Verified Permissions
	// substitutes this resource for the ?resource placeholder in the policy template
	// when it evaluates an authorization request.
	Resource *EntityIdentifier

	noSmithyDocumentSerde
}

// Contains configuration details of a Amazon Cognito user pool for use with an
// identity source.
type UpdateCognitoUserPoolConfiguration struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the Amazon Cognito user pool associated with this identity source.
	//
	// This member is required.
	UserPoolArn *string

	// The client ID of an app client that is configured for the specified Amazon
	// Cognito user pool.
	ClientIds []string

	noSmithyDocumentSerde
}

// Contains an updated configuration to replace the configuration in an existing
// identity source. At this time, the only valid member of this structure is a
// Amazon Cognito user pool configuration. You must specify a userPoolArn , and
// optionally, a ClientId .
//
// The following types satisfy this interface:
//
//	UpdateConfigurationMemberCognitoUserPoolConfiguration
type UpdateConfiguration interface {
	isUpdateConfiguration()
}

// Contains configuration details of a Amazon Cognito user pool.
type UpdateConfigurationMemberCognitoUserPoolConfiguration struct {
	Value UpdateCognitoUserPoolConfiguration

	noSmithyDocumentSerde
}

func (*UpdateConfigurationMemberCognitoUserPoolConfiguration) isUpdateConfiguration() {}

// Contains information about updates to be applied to a policy. This data type is
// used as a request parameter in the UpdatePolicy (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicy.html)
// operation.
//
// The following types satisfy this interface:
//
//	UpdatePolicyDefinitionMemberStatic
type UpdatePolicyDefinition interface {
	isUpdatePolicyDefinition()
}

// Contains details about the updates to be applied to a static policy.
type UpdatePolicyDefinitionMemberStatic struct {
	Value UpdateStaticPolicyDefinition

	noSmithyDocumentSerde
}

func (*UpdatePolicyDefinitionMemberStatic) isUpdatePolicyDefinition() {}

// Contains information about an update to a static policy.
type UpdateStaticPolicyDefinition struct {

	// Specifies the Cedar policy language text to be added to or replaced on the
	// static policy. You can change only the following elements from the original
	// content:
	//   - The action referenced by the policy.
	//   - Any conditional clauses, such as when or unless clauses.
	// You can't change the following elements:
	//   - Changing from StaticPolicy to TemplateLinkedPolicy .
	//   - The effect ( permit or forbid ) of the policy.
	//   - The principal referenced by the policy.
	//   - The resource referenced by the policy.
	//
	// This member is required.
	Statement *string

	// Specifies the description to be added to or replaced on the static policy.
	Description *string

	noSmithyDocumentSerde
}

// Details about a field that failed policy validation.
type ValidationExceptionField struct {

	// Describes the policy validation error.
	//
	// This member is required.
	Message *string

	// The path to the specific element that Verified Permissions found to be not
	// valid.
	//
	// This member is required.
	Path *string

	noSmithyDocumentSerde
}

// A structure that contains Cedar policy validation settings for the policy
// store. The validation mode determines which validation failures that Cedar
// considers serious enough to block acceptance of a new or edited static policy or
// policy template. This data type is used as a request parameter in the
// CreatePolicyStore (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicyStore.html)
// and UpdatePolicyStore (https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_UpdatePolicyStore.html)
// operations.
type ValidationSettings struct {

	// The validation mode currently configured for this policy store. The valid
	// values are:
	//   - OFF – Neither Verified Permissions nor Cedar perform any validation on
	//   policies. No validation errors are reported by either service.
	//   - STRICT – Requires a schema to be present in the policy store. Cedar
	//   performs validation on all submitted new or updated static policies and policy
	//   templates. Any that fail validation are rejected and Cedar doesn't store them in
	//   the policy store.
	// If Mode=STRICT and the policy store doesn't contain a schema, Verified
	// Permissions rejects all static policies and policy templates because there is no
	// schema to validate against. To submit a static policy or policy template without
	// a schema, you must turn off validation.
	//
	// This member is required.
	Mode ValidationMode

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAttributeValue()         {}
func (*UnknownUnionMember) isConfiguration()          {}
func (*UnknownUnionMember) isContextDefinition()      {}
func (*UnknownUnionMember) isEntitiesDefinition()     {}
func (*UnknownUnionMember) isEntityReference()        {}
func (*UnknownUnionMember) isPolicyDefinition()       {}
func (*UnknownUnionMember) isPolicyDefinitionDetail() {}
func (*UnknownUnionMember) isPolicyDefinitionItem()   {}
func (*UnknownUnionMember) isSchemaDefinition()       {}
func (*UnknownUnionMember) isUpdateConfiguration()    {}
func (*UnknownUnionMember) isUpdatePolicyDefinition() {}
