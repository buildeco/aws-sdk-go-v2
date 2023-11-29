// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty dataset and adds it to the specified dataset group. Use
// CreateDatasetImportJob (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html)
// to import your training data to a dataset. There are 5 types of datasets:
//   - Item interactions
//   - Items
//   - Users
//   - Action interactions
//   - Actions
//
// Each dataset type has an associated schema with required field types. Only the
// Item interactions dataset is required in order to train a model (also referred
// to as creating a solution). A dataset can be in one of the following states:
//   - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//   - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the dataset, call DescribeDataset (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html)
// . Related APIs
//   - CreateDatasetGroup (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetGroup.html)
//   - ListDatasets (https://docs.aws.amazon.com/personalize/latest/dg/API_ListDatasets.html)
//   - DescribeDataset (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeDataset.html)
//   - DeleteDataset (https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteDataset.html)
func (c *Client) CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) {
	if params == nil {
		params = &CreateDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataset", params, optFns, c.addOperationCreateDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetInput struct {

	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to.
	//
	// This member is required.
	DatasetGroupArn *string

	// The type of dataset. One of the following (case insensitive) values:
	//   - Interactions
	//   - Items
	//   - Users
	//   - Actions
	//   - Action_Interactions
	//
	// This member is required.
	DatasetType *string

	// The name for the dataset.
	//
	// This member is required.
	Name *string

	// The ARN of the schema to associate with the dataset. The schema defines the
	// dataset fields.
	//
	// This member is required.
	SchemaArn *string

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html)
	// to apply to the dataset.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDatasetOutput struct {

	// The ARN of the dataset.
	DatasetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataset"); err != nil {
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
	if err = addOpCreateDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataset",
	}
}
