// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This request can be sent after CreateRestoreTestingPlan request returns
// successfully. This is the second part of creating a resource testing plan, and
// it must be completed sequentially. This consists of RestoreTestingSelectionName
// , ProtectedResourceType , and one of the following:
//   - ProtectedResourceArns
//   - ProtectedResourceConditions
//
// Each protected resource type can have one single value. A restore testing
// selection can include a wildcard value ("*") for ProtectedResourceArns along
// with ProtectedResourceConditions . Alternatively, you can include up to 30
// specific protected resource ARNs in ProtectedResourceArns . Cannot select by
// both protected resource types AND specific ARNs. Request will fail if both are
// included.
func (c *Client) CreateRestoreTestingSelection(ctx context.Context, params *CreateRestoreTestingSelectionInput, optFns ...func(*Options)) (*CreateRestoreTestingSelectionOutput, error) {
	if params == nil {
		params = &CreateRestoreTestingSelectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRestoreTestingSelection", params, optFns, c.addOperationCreateRestoreTestingSelectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRestoreTestingSelectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRestoreTestingSelectionInput struct {

	// Input the restore testing plan name that was returned from the related
	// CreateRestoreTestingPlan request.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// This consists of RestoreTestingSelectionName , ProtectedResourceType , and one
	// of the following:
	//   - ProtectedResourceArns
	//   - ProtectedResourceConditions
	// Each protected resource type can have one single value. A restore testing
	// selection can include a wildcard value ("*") for ProtectedResourceArns along
	// with ProtectedResourceConditions . Alternatively, you can include up to 30
	// specific protected resource ARNs in ProtectedResourceArns .
	//
	// This member is required.
	RestoreTestingSelection *types.RestoreTestingSelectionForCreate

	// This is an optional unique string that identifies the request and allows failed
	// requests to be retried without the risk of running the operation twice. If used,
	// this parameter must contain 1 to 50 alphanumeric or '-_.' characters.
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type CreateRestoreTestingSelectionOutput struct {

	// This is the time the resource testing selection was created successfully.
	//
	// This member is required.
	CreationTime *time.Time

	// This is the ARN of the restore testing plan with which the restore testing
	// selection is associated.
	//
	// This member is required.
	RestoreTestingPlanArn *string

	// Unique string that is the name of the restore testing plan. The name cannot be
	// changed after creation. The name consists of only alphanumeric characters and
	// underscores. Maximum length is 50.
	//
	// This member is required.
	RestoreTestingPlanName *string

	// This is the unique name of the restore testing selection that belongs to the
	// related restore testing plan.
	//
	// This member is required.
	RestoreTestingSelectionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRestoreTestingSelectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRestoreTestingSelection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRestoreTestingSelection"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateRestoreTestingSelectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRestoreTestingSelection(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateRestoreTestingSelection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRestoreTestingSelection",
	}
}
