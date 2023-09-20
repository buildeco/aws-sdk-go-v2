// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates one or more records and, optionally, creates a health check
// based on the settings in a specified service. When you submit a RegisterInstance
// request, the following occurs:
//   - For each DNS record that you define in the service that's specified by
//     ServiceId , a record is created or updated in the hosted zone that's
//     associated with the corresponding namespace.
//   - If the service includes HealthCheckConfig , a health check is created based
//     on the settings in the health check configuration.
//   - The health check, if any, is associated with each of the new or updated
//     records.
//
// One RegisterInstance request must complete before you can submit another
// request and specify the same service ID and instance ID. For more information,
// see CreateService (https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html)
// . When Cloud Map receives a DNS query for the specified DNS name, it returns the
// applicable value:
//   - If the health check is healthy: returns all the records
//   - If the health check is unhealthy: returns the applicable value for the last
//     healthy instance
//   - If you didn't specify a health check configuration: returns all the records
//
// For the current quota on the number of instances that you can register using
// the same namespace and using the same service, see Cloud Map quotas (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the Cloud Map Developer Guide.
func (c *Client) RegisterInstance(ctx context.Context, params *RegisterInstanceInput, optFns ...func(*Options)) (*RegisterInstanceOutput, error) {
	if params == nil {
		params = &RegisterInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterInstance", params, optFns, c.addOperationRegisterInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterInstanceInput struct {

	// A string map that contains the following information for the service that you
	// specify in ServiceId :
	//   - The attributes that apply to the records that are defined in the service.
	//   - For each attribute, the applicable value.
	// Do not include sensitive information in the attributes if the namespace is
	// discoverable by public DNS queries. Supported attribute keys include the
	// following: AWS_ALIAS_DNS_NAME If you want Cloud Map to create an Amazon Route 53
	// alias record that routes traffic to an Elastic Load Balancing load balancer,
	// specify the DNS name that's associated with the load balancer. For information
	// about how to get the DNS name, see "DNSName" in the topic AliasTarget (https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html)
	// in the Route 53 API Reference. Note the following:
	//   - The configuration for the service that's specified by ServiceId must include
	//   settings for an A record, an AAAA record, or both.
	//   - In the service that's specified by ServiceId , the value of RoutingPolicy
	//   must be WEIGHTED .
	//   - If the service that's specified by ServiceId includes HealthCheckConfig
	//   settings, Cloud Map will create the Route 53 health check, but it doesn't
	//   associate the health check with the alias record.
	//   - Cloud Map currently doesn't support creating alias records that route
	//   traffic to Amazon Web Services resources other than Elastic Load Balancing load
	//   balancers.
	//   - If you specify a value for AWS_ALIAS_DNS_NAME , don't specify values for any
	//   of the AWS_INSTANCE attributes.
	// AWS_EC2_INSTANCE_ID HTTP namespaces only. The Amazon EC2 instance ID for the
	// instance. If the AWS_EC2_INSTANCE_ID attribute is specified, then the only
	// other attribute that can be specified is AWS_INIT_HEALTH_STATUS . When the
	// AWS_EC2_INSTANCE_ID attribute is specified, then the AWS_INSTANCE_IPV4
	// attribute will be filled out with the primary private IPv4 address.
	// AWS_INIT_HEALTH_STATUS If the service configuration includes
	// HealthCheckCustomConfig , you can optionally use AWS_INIT_HEALTH_STATUS to
	// specify the initial status of the custom health check, HEALTHY or UNHEALTHY . If
	// you don't specify a value for AWS_INIT_HEALTH_STATUS , the initial status is
	// HEALTHY . AWS_INSTANCE_CNAME If the service configuration includes a CNAME
	// record, the domain name that you want Route 53 to return in response to DNS
	// queries (for example, example.com ). This value is required if the service
	// specified by ServiceId includes settings for an CNAME record. AWS_INSTANCE_IPV4
	// If the service configuration includes an A record, the IPv4 address that you
	// want Route 53 to return in response to DNS queries (for example, 192.0.2.44 ).
	// This value is required if the service specified by ServiceId includes settings
	// for an A record. If the service includes settings for an SRV record, you must
	// specify a value for AWS_INSTANCE_IPV4 , AWS_INSTANCE_IPV6 , or both.
	// AWS_INSTANCE_IPV6 If the service configuration includes an AAAA record, the
	// IPv6 address that you want Route 53 to return in response to DNS queries (for
	// example, 2001:0db8:85a3:0000:0000:abcd:0001:2345 ). This value is required if
	// the service specified by ServiceId includes settings for an AAAA record. If the
	// service includes settings for an SRV record, you must specify a value for
	// AWS_INSTANCE_IPV4 , AWS_INSTANCE_IPV6 , or both. AWS_INSTANCE_PORT If the
	// service includes an SRV record, the value that you want Route 53 to return for
	// the port. If the service includes HealthCheckConfig , the port on the endpoint
	// that you want Route 53 to send requests to. This value is required if you
	// specified settings for an SRV record or a Route 53 health check when you
	// created the service. Custom attributes You can add up to 30 custom attributes.
	// For each key-value pair, the maximum length of the attribute name is 255
	// characters, and the maximum length of the attribute value is 1,024 characters.
	// The total size of all provided attributes (sum of all keys and values) must not
	// exceed 5,000 characters.
	//
	// This member is required.
	Attributes map[string]string

	// An identifier that you want to associate with the instance. Note the following:
	//   - If the service that's specified by ServiceId includes settings for an SRV
	//   record, the value of InstanceId is automatically included as part of the value
	//   for the SRV record. For more information, see DnsRecord > Type (https://docs.aws.amazon.com/cloud-map/latest/api/API_DnsRecord.html#cloudmap-Type-DnsRecord-Type)
	//   .
	//   - You can use this value to update an existing instance.
	//   - To register a new instance, you must specify a value that's unique among
	//   instances that you register by using the same service.
	//   - If you specify an existing InstanceId and ServiceId , Cloud Map updates the
	//   existing DNS records, if any. If there's also an existing health check, Cloud
	//   Map deletes the old health check and creates a new one. The health check isn't
	//   deleted immediately, so it will still appear for a while if you submit a
	//   ListHealthChecks request, for example.
	// Do not include sensitive information in InstanceId if the namespace is
	// discoverable by public DNS queries and any Type member of DnsRecord for the
	// service contains SRV because the InstanceId is discoverable by public DNS
	// queries.
	//
	// This member is required.
	InstanceId *string

	// The ID of the service that you want to use for settings for the instance.
	//
	// This member is required.
	ServiceId *string

	// A unique string that identifies the request and that allows failed
	// RegisterInstance requests to be retried without the risk of executing the
	// operation twice. You must use a unique CreatorRequestId string every time you
	// submit a RegisterInstance request if you're registering additional instances
	// for the same namespace and service. CreatorRequestId can be any unique string
	// (for example, a date/time stamp).
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type RegisterInstanceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addRegisterInstanceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opRegisterInstanceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpRegisterInstance struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterInstance) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterInstanceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterInstanceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterInstance{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "RegisterInstance",
	}
}

type opRegisterInstanceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opRegisterInstanceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRegisterInstanceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "servicediscovery"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "servicediscovery"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("servicediscovery")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addRegisterInstanceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opRegisterInstanceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
