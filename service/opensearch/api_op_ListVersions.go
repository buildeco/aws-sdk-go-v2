// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all supported versions of OpenSearch and Elasticsearch.
func (c *Client) ListVersions(ctx context.Context, params *ListVersionsInput, optFns ...func(*Options)) (*ListVersionsOutput, error) {
	if params == nil {
		params = &ListVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVersions", params, optFns, c.addOperationListVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the ListVersions operation. Use MaxResults to
// control the maximum number of results to retrieve in a single call. Use
// NextToken in response to retrieve more results. If the received response does
// not contain a NextToken, there are no more results to retrieve.
type ListVersionsInput struct {

	// Set this value to limit the number of results returned. Value must be greater
	// than 10 or it won't be honored.
	MaxResults int32

	// Paginated APIs accept the NextToken input to return the next page of results and
	// provide a NextToken output in the response, which you can use to retrieve more
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for the parameters for response received from the ListVersions
// operation.
type ListVersionsOutput struct {

	// Paginated APIs accept the NextToken input to return the next page of results and
	// provide a NextToken output in the response, which you can use to retrieve more
	// results.
	NextToken *string

	// List of supported OpenSearch versions.
	Versions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVersions(options.Region), middleware.Before); err != nil {
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

// ListVersionsAPIClient is a client that implements the ListVersions operation.
type ListVersionsAPIClient interface {
	ListVersions(context.Context, *ListVersionsInput, ...func(*Options)) (*ListVersionsOutput, error)
}

var _ ListVersionsAPIClient = (*Client)(nil)

// ListVersionsPaginatorOptions is the paginator options for ListVersions
type ListVersionsPaginatorOptions struct {
	// Set this value to limit the number of results returned. Value must be greater
	// than 10 or it won't be honored.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVersionsPaginator is a paginator for ListVersions
type ListVersionsPaginator struct {
	options   ListVersionsPaginatorOptions
	client    ListVersionsAPIClient
	params    *ListVersionsInput
	nextToken *string
	firstPage bool
}

// NewListVersionsPaginator returns a new ListVersionsPaginator
func NewListVersionsPaginator(client ListVersionsAPIClient, params *ListVersionsInput, optFns ...func(*ListVersionsPaginatorOptions)) *ListVersionsPaginator {
	if params == nil {
		params = &ListVersionsInput{}
	}

	options := ListVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListVersions page.
func (p *ListVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListVersions",
	}
}
