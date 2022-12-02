// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List existing ephemerides.
func (c *Client) ListEphemerides(ctx context.Context, params *ListEphemeridesInput, optFns ...func(*Options)) (*ListEphemeridesOutput, error) {
	if params == nil {
		params = &ListEphemeridesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEphemerides", params, optFns, c.addOperationListEphemeridesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEphemeridesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEphemeridesInput struct {

	// The end time to list in UTC. The operation will return an ephemeris if its
	// expiration time is within the time range defined by the startTime and endTime.
	//
	// This member is required.
	EndTime *time.Time

	// The AWS Ground Station satellite ID to list ephemeris for.
	//
	// This member is required.
	SatelliteId *string

	// The start time to list in UTC. The operation will return an ephemeris if its
	// expiration time is within the time range defined by the startTime and endTime.
	//
	// This member is required.
	StartTime *time.Time

	// Maximum number of ephemerides to return.
	MaxResults *int32

	// Pagination token.
	NextToken *string

	// The list of ephemeris status to return.
	StatusList []types.EphemerisStatus

	noSmithyDocumentSerde
}

type ListEphemeridesOutput struct {

	// List of ephemerides.
	Ephemerides []types.EphemerisItem

	// Pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEphemeridesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEphemerides{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEphemerides{}, middleware.After)
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
	if err = addOpListEphemeridesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEphemerides(options.Region), middleware.Before); err != nil {
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

// ListEphemeridesAPIClient is a client that implements the ListEphemerides
// operation.
type ListEphemeridesAPIClient interface {
	ListEphemerides(context.Context, *ListEphemeridesInput, ...func(*Options)) (*ListEphemeridesOutput, error)
}

var _ ListEphemeridesAPIClient = (*Client)(nil)

// ListEphemeridesPaginatorOptions is the paginator options for ListEphemerides
type ListEphemeridesPaginatorOptions struct {
	// Maximum number of ephemerides to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEphemeridesPaginator is a paginator for ListEphemerides
type ListEphemeridesPaginator struct {
	options   ListEphemeridesPaginatorOptions
	client    ListEphemeridesAPIClient
	params    *ListEphemeridesInput
	nextToken *string
	firstPage bool
}

// NewListEphemeridesPaginator returns a new ListEphemeridesPaginator
func NewListEphemeridesPaginator(client ListEphemeridesAPIClient, params *ListEphemeridesInput, optFns ...func(*ListEphemeridesPaginatorOptions)) *ListEphemeridesPaginator {
	if params == nil {
		params = &ListEphemeridesInput{}
	}

	options := ListEphemeridesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEphemeridesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEphemeridesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEphemerides page.
func (p *ListEphemeridesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEphemeridesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListEphemerides(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListEphemerides(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "ListEphemerides",
	}
}