// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation is not supported by directory buckets. Returns the status of the
// resource policy associated with an Object Lambda Access Point.
func (c *Client) GetAccessPointPolicyStatusForObjectLambda(ctx context.Context, params *GetAccessPointPolicyStatusForObjectLambdaInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusForObjectLambdaOutput, error) {
	if params == nil {
		params = &GetAccessPointPolicyStatusForObjectLambdaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPointPolicyStatusForObjectLambda", params, optFns, c.addOperationGetAccessPointPolicyStatusForObjectLambdaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointPolicyStatusForObjectLambdaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyStatusForObjectLambdaInput struct {

	// The account ID for the account that owns the specified Object Lambda Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// The name of the Object Lambda Access Point.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

func (in *GetAccessPointPolicyStatusForObjectLambdaInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type GetAccessPointPolicyStatusForObjectLambdaOutput struct {

	// Indicates whether this access point policy is public. For more information
	// about how Amazon S3 evaluates policies to determine whether they are public, see
	// The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon S3 User Guide.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessPointPolicyStatusForObjectLambdaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicyStatusForObjectLambda{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicyStatusForObjectLambda{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessPointPolicyStatusForObjectLambda"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAccessPointPolicyStatusForObjectLambdaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicyStatusForObjectLambda(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetAccessPointPolicyStatusForObjectLambdaUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware struct {
}

func (*endpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*GetAccessPointPolicyStatusForObjectLambdaInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAccessPointPolicyStatusForObjectLambdaMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetAccessPointPolicyStatusForObjectLambda(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessPointPolicyStatusForObjectLambda",
	}
}

func copyGetAccessPointPolicyStatusForObjectLambdaInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetAccessPointPolicyStatusForObjectLambdaInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetAccessPointPolicyStatusForObjectLambdaInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *GetAccessPointPolicyStatusForObjectLambdaInput) copy() interface{} {
	v := *in
	return &v
}
func backFillGetAccessPointPolicyStatusForObjectLambdaAccountID(input interface{}, v string) error {
	in := input.(*GetAccessPointPolicyStatusForObjectLambdaInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetAccessPointPolicyStatusForObjectLambdaUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetAccessPointPolicyStatusForObjectLambdaInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
