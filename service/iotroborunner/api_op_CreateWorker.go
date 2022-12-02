// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to create a worker
func (c *Client) CreateWorker(ctx context.Context, params *CreateWorkerInput, optFns ...func(*Options)) (*CreateWorkerOutput, error) {
	if params == nil {
		params = &CreateWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorker", params, optFns, c.addOperationCreateWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkerInput struct {

	// Full ARN of the worker fleet.
	//
	// This member is required.
	Fleet *string

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// JSON blob containing unstructured worker properties that are fixed and won't
	// change during regular operation.
	AdditionalFixedProperties *string

	// JSON blob containing unstructured worker properties that are transient and may
	// change during regular operation.
	AdditionalTransientProperties *string

	// Token used for detecting replayed requests. Replayed requests will not be
	// performed multiple times.
	ClientToken *string

	// Worker orientation measured in units clockwise from north.
	Orientation types.Orientation

	// Supported coordinates for worker position.
	Position types.PositionCoordinates

	// Properties of the worker that are provided by the vendor FMS.
	VendorProperties *types.VendorProperties

	noSmithyDocumentSerde
}

type CreateWorkerOutput struct {

	// Full ARN of the worker.
	//
	// This member is required.
	Arn *string

	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Filters access by the workers identifier
	//
	// This member is required.
	Id *string

	// Site ARN.
	//
	// This member is required.
	Site *string

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorker{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateWorkerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorker(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorker struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorker) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorkerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkerInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWorkerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorker{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "CreateWorker",
	}
}