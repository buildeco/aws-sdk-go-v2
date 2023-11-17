// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the inputs for the change set and a list of changes that CloudFormation
// will make if you execute the change set. For more information, see Updating
// Stacks Using Change Sets (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html)
// in the CloudFormation User Guide.
func (c *Client) DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) {
	if params == nil {
		params = &DescribeChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChangeSet", params, optFns, c.addOperationDescribeChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeChangeSet action.
type DescribeChangeSetInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// describe.
	//
	// This member is required.
	ChangeSetName *string

	// A string (provided by the DescribeChangeSet response output) that identifies
	// the next page of information that you want to retrieve.
	NextToken *string

	// If you specified the name of a change set, specify the stack name or ID (ARN)
	// of the change set you want to describe.
	StackName *string

	noSmithyDocumentSerde
}

// The output for the DescribeChangeSet action.
type DescribeChangeSetOutput struct {

	// If you execute the change set, the list of capabilities that were explicitly
	// acknowledged when the change set was created.
	Capabilities []types.Capability

	// The Amazon Resource Name (ARN) of the change set.
	ChangeSetId *string

	// The name of the change set.
	ChangeSetName *string

	// A list of Change structures that describes the resources CloudFormation changes
	// if you execute the change set.
	Changes []types.Change

	// The start time when the change set was created, in UTC.
	CreationTime *time.Time

	// Information about the change set.
	Description *string

	// If the change set execution status is AVAILABLE , you can execute the change
	// set. If you can't execute the change set, the status indicates why. For example,
	// a change set might be in an UNAVAILABLE state because CloudFormation is still
	// creating it or in an OBSOLETE state because the stack was already updated.
	ExecutionStatus types.ExecutionStatus

	// Indicates if the stack set imports resources that already exist. This parameter
	// can only import resources that have custom names (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html)
	// in templates. To import resources that do not accept custom names, such as EC2
	// instances, use the resource import (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import.html)
	// feature instead.
	ImportExistingResources *bool

	// Verifies if IncludeNestedStacks is set to True .
	IncludeNestedStacks *bool

	// If the output exceeds 1 MB, a string that identifies the next page of changes.
	// If there is no additional page, this value is null.
	NextToken *string

	// The ARNs of the Amazon Simple Notification Service (Amazon SNS) topics that
	// will be associated with the stack if you execute the change set.
	NotificationARNs []string

	// Determines what action will be taken if stack creation fails. When this
	// parameter is specified, the DisableRollback parameter to the ExecuteChangeSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteChangeSet.html)
	// API operation must not be specified. This must be one of these values:
	//   - DELETE - Deletes the change set if the stack creation fails. This is only
	//   valid when the ChangeSetType parameter is set to CREATE . If the deletion of
	//   the stack fails, the status of the stack is DELETE_FAILED .
	//   - DO_NOTHING - if the stack creation fails, do nothing. This is equivalent to
	//   specifying true for the DisableRollback parameter to the ExecuteChangeSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteChangeSet.html)
	//   API operation.
	//   - ROLLBACK - if the stack creation fails, roll back the stack. This is
	//   equivalent to specifying false for the DisableRollback parameter to the
	//   ExecuteChangeSet (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteChangeSet.html)
	//   API operation.
	OnStackFailure types.OnStackFailure

	// A list of Parameter structures that describes the input parameters and their
	// values used to create the change set. For more information, see the Parameter (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []types.Parameter

	// Specifies the change set ID of the parent change set in the current nested
	// change set hierarchy.
	ParentChangeSetId *string

	// The rollback triggers for CloudFormation to monitor during stack creation and
	// updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Specifies the change set ID of the root change set in the current nested change
	// set hierarchy.
	RootChangeSetId *string

	// The Amazon Resource Name (ARN) of the stack that's associated with the change
	// set.
	StackId *string

	// The name of the stack that's associated with the change set.
	StackName *string

	// The current status of the change set, such as CREATE_IN_PROGRESS ,
	// CREATE_COMPLETE , or FAILED .
	Status types.ChangeSetStatus

	// A description of the change set's status. For example, if your attempt to
	// create a change set failed, CloudFormation shows the error message.
	StatusReason *string

	// If you execute the change set, the tags that will be associated with the stack.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeChangeSet"); err != nil {
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
	if err = addOpDescribeChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChangeSet(options.Region), middleware.Before); err != nil {
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

// DescribeChangeSetAPIClient is a client that implements the DescribeChangeSet
// operation.
type DescribeChangeSetAPIClient interface {
	DescribeChangeSet(context.Context, *DescribeChangeSetInput, ...func(*Options)) (*DescribeChangeSetOutput, error)
}

var _ DescribeChangeSetAPIClient = (*Client)(nil)

// ChangeSetCreateCompleteWaiterOptions are waiter options for
// ChangeSetCreateCompleteWaiter
type ChangeSetCreateCompleteWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ChangeSetCreateCompleteWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ChangeSetCreateCompleteWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeChangeSetInput, *DescribeChangeSetOutput, error) (bool, error)
}

// ChangeSetCreateCompleteWaiter defines the waiters for ChangeSetCreateComplete
type ChangeSetCreateCompleteWaiter struct {
	client DescribeChangeSetAPIClient

	options ChangeSetCreateCompleteWaiterOptions
}

// NewChangeSetCreateCompleteWaiter constructs a ChangeSetCreateCompleteWaiter.
func NewChangeSetCreateCompleteWaiter(client DescribeChangeSetAPIClient, optFns ...func(*ChangeSetCreateCompleteWaiterOptions)) *ChangeSetCreateCompleteWaiter {
	options := ChangeSetCreateCompleteWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = changeSetCreateCompleteStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ChangeSetCreateCompleteWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ChangeSetCreateComplete waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ChangeSetCreateCompleteWaiter) Wait(ctx context.Context, params *DescribeChangeSetInput, maxWaitDur time.Duration, optFns ...func(*ChangeSetCreateCompleteWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ChangeSetCreateComplete waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ChangeSetCreateCompleteWaiter) WaitForOutput(ctx context.Context, params *DescribeChangeSetInput, maxWaitDur time.Duration, optFns ...func(*ChangeSetCreateCompleteWaiterOptions)) (*DescribeChangeSetOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeChangeSet(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ChangeSetCreateComplete waiter")
}

func changeSetCreateCompleteStateRetryable(ctx context.Context, input *DescribeChangeSetInput, output *DescribeChangeSetOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_COMPLETE"
		value, ok := pathValue.(types.ChangeSetStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ChangeSetStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.ChangeSetStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ChangeSetStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationError" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeChangeSet",
	}
}
