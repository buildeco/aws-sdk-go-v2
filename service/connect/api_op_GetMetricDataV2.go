// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets metric data from the specified Amazon Connect instance. GetMetricDataV2
// offers more features than GetMetricData (https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricData.html)
// , the previous version of this API. It has new metrics, offers filtering at a
// metric level, and offers the ability to filter and group data by channels,
// queues, routing profiles, agents, and agent hierarchy levels. It can retrieve
// historical data for the last 3 months, at varying intervals. For a description
// of the historical metrics that are supported by GetMetricDataV2 and
// GetMetricData , see Historical metrics definitions (https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html)
// in the Amazon Connect Administrator's Guide.
func (c *Client) GetMetricDataV2(ctx context.Context, params *GetMetricDataV2Input, optFns ...func(*Options)) (*GetMetricDataV2Output, error) {
	if params == nil {
		params = &GetMetricDataV2Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricDataV2", params, optFns, c.addOperationGetMetricDataV2Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricDataV2Output)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricDataV2Input struct {

	// The timestamp, in UNIX Epoch time format, at which to end the reporting
	// interval for the retrieval of historical metrics data. The time must be later
	// than the start time timestamp. It cannot be later than the current timestamp.
	//
	// This member is required.
	EndTime *time.Time

	// The filters to apply to returned metrics. You can filter on the following
	// resources:
	//   - Queues
	//   - Routing profiles
	//   - Agents
	//   - Channels
	//   - User hierarchy groups
	//   - Feature
	// At least one filter must be passed from queues, routing profiles, agents, or
	// user hierarchy groups. To filter by phone number, see Create a historical
	// metrics report (https://docs.aws.amazon.com/connect/latest/adminguide/create-historical-metrics-report.html)
	// in the Amazon Connect Administrator's Guide. Note the following limits:
	//   - Filter keys: A maximum of 5 filter keys are supported in a single request.
	//   Valid filter keys: QUEUE | ROUTING_PROFILE | AGENT | CHANNEL |
	//   AGENT_HIERARCHY_LEVEL_ONE | AGENT_HIERARCHY_LEVEL_TWO |
	//   AGENT_HIERARCHY_LEVEL_THREE | AGENT_HIERARCHY_LEVEL_FOUR |
	//   AGENT_HIERARCHY_LEVEL_FIVE | FEATURE |
	//   contact/segmentAttributes/connect:Subtype
	//   - Filter values: A maximum of 100 filter values are supported in a single
	//   request. VOICE, CHAT, and TASK are valid filterValue for the CHANNEL filter
	//   key. They do not count towards limitation of 100 filter values. For example, a
	//   GetMetricDataV2 request can filter by 50 queues, 35 agents, and 15 routing
	//   profiles for a total of 100 filter values, along with 3 channel filters.
	//   contact_lens_conversational_analytics is a valid filterValue for the FEATURE
	//   filter key. It is available only to contacts analyzed by Contact Lens
	//   conversational analytics. connect:Chat , connect:SMS , connect:Telephony , and
	//   connect:WebRTC are valid filterValue examples (not exhaustive) for the
	//   contact/segmentAttributes/connect:Subtype filter key.
	//
	// This member is required.
	Filters []types.FilterV2

	// The metrics to retrieve. Specify the name, groupings, and filters for each
	// metric. The following historical metrics are available. For a description of
	// each metric, see Historical metrics definitions (https://docs.aws.amazon.com/connect/latest/adminguide/historical-metrics-definitions.html)
	// in the Amazon Connect Administrator's Guide. ABANDONMENT_RATE Unit: Percent
	// Valid groupings and filters: Queue, Channel, Routing Profile, Agent, Agent
	// Hierarchy, Feature, contact/segmentAttributes/connect:Subtype
	// AGENT_ADHERENT_TIME This metric is available only in Amazon Web Services Regions
	// where Forecasting, capacity planning, and scheduling (https://docs.aws.amazon.com/connect/latest/adminguide/regions.html#optimization_region)
	// is available. Unit: Seconds Valid groupings and filters: Queue, Channel, Routing
	// Profile, Agent, Agent Hierarchy AGENT_ANSWER_RATE Unit: Percent Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// AGENT_NON_ADHERENT_TIME Unit: Seconds Valid groupings and filters: Queue,
	// Channel, Routing Profile, Agent, Agent Hierarchy AGENT_NON_RESPONSE Unit: Count
	// Valid groupings and filters: Queue, Channel, Routing Profile, Agent, Agent
	// Hierarchy AGENT_NON_RESPONSE_WITHOUT_CUSTOMER_ABANDONS Unit: Count Valid
	// groupings and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// Data for this metric is available starting from October 1, 2023 0:00:00 GMT.
	// AGENT_OCCUPANCY Unit: Percentage Valid groupings and filters: Routing Profile,
	// Agent, Agent Hierarchy AGENT_SCHEDULE_ADHERENCE This metric is available only in
	// Amazon Web Services Regions where Forecasting, capacity planning, and scheduling (https://docs.aws.amazon.com/connect/latest/adminguide/regions.html#optimization_region)
	// is available. Unit: Percent Valid groupings and filters: Queue, Channel, Routing
	// Profile, Agent, Agent Hierarchy AGENT_SCHEDULED_TIME This metric is available
	// only in Amazon Web Services Regions where Forecasting, capacity planning, and
	// scheduling (https://docs.aws.amazon.com/connect/latest/adminguide/regions.html#optimization_region)
	// is available. Unit: Seconds Valid groupings and filters: Queue, Channel, Routing
	// Profile, Agent, Agent Hierarchy AVG_ABANDON_TIME Unit: Seconds Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy, Feature,
	// contact/segmentAttributes/connect:Subtype AVG_AFTER_CONTACT_WORK_TIME Unit:
	// Seconds Valid metric filter key: INITIATION_METHOD Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy, Feature,
	// contact/segmentAttributes/connect:Subtype Feature is a valid filter but not a
	// valid grouping. AVG_AGENT_CONNECTING_TIME Unit: Seconds Valid metric filter key:
	// INITIATION_METHOD . For now, this metric only supports the following as
	// INITIATION_METHOD : INBOUND | OUTBOUND | CALLBACK | API Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy The Negate key
	// in Metric Level Filters is not applicable for this metric. AVG_CONTACT_DURATION
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, Feature, contact/segmentAttributes/connect:Subtype
	// Feature is a valid filter but not a valid grouping. AVG_CONVERSATION_DURATION
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, Feature, contact/segmentAttributes/connect:Subtype
	// AVG_GREETING_TIME_AGENT This metric is available only for contacts analyzed by
	// Contact Lens conversational analytics. Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype AVG_HANDLE_TIME Unit: Seconds Valid
	// groupings and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// Feature, contact/segmentAttributes/connect:Subtype Feature is a valid filter but
	// not a valid grouping. AVG_HOLD_TIME Unit: Seconds Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy, Feature,
	// contact/segmentAttributes/connect:Subtype Feature is a valid filter but not a
	// valid grouping. AVG_HOLD_TIME_ALL_CONTACTS Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype AVG_HOLDS Unit: Count Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy, Feature,
	// contact/segmentAttributes/connect:Subtype Feature is a valid filter but not a
	// valid grouping. AVG_INTERACTION_AND_HOLD_TIME Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype AVG_INTERACTION_TIME Unit: Seconds
	// Valid metric filter key: INITIATION_METHOD Valid groupings and filters: Queue,
	// Channel, Routing Profile, Feature, contact/segmentAttributes/connect:Subtype
	// Feature is a valid filter but not a valid grouping. AVG_INTERRUPTIONS_AGENT This
	// metric is available only for contacts analyzed by Contact Lens conversational
	// analytics. Unit: Count Valid groupings and filters: Queue, Channel, Routing
	// Profile, Agent, Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// AVG_INTERRUPTION_TIME_AGENT This metric is available only for contacts analyzed
	// by Contact Lens conversational analytics. Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype AVG_NON_TALK_TIME This metric is
	// available only for contacts analyzed by Contact Lens conversational analytics.
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// AVG_QUEUE_ANSWER_TIME Unit: Seconds Valid groupings and filters: Queue, Channel,
	// Routing Profile, Feature, contact/segmentAttributes/connect:Subtype Feature is a
	// valid filter but not a valid grouping. AVG_RESOLUTION_TIME Unit: Seconds Valid
	// groupings and filters: Queue, Channel, Routing Profile,
	// contact/segmentAttributes/connect:Subtype AVG_TALK_TIME This metric is available
	// only for contacts analyzed by Contact Lens conversational analytics. Unit:
	// Seconds Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy, contact/segmentAttributes/connect:Subtype AVG_TALK_TIME_AGENT
	// This metric is available only for contacts analyzed by Contact Lens
	// conversational analytics. Unit: Seconds Valid groupings and filters: Queue,
	// Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype AVG_TALK_TIME_CUSTOMER This metric is
	// available only for contacts analyzed by Contact Lens conversational analytics.
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// CONTACTS_ABANDONED Unit: Count Valid groupings and filters: Queue, Channel,
	// Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype CONTACTS_CREATED Unit: Count Valid
	// metric filter key: INITIATION_METHOD Valid groupings and filters: Queue,
	// Channel, Routing Profile, Feature, contact/segmentAttributes/connect:Subtype
	// Feature is a valid filter but not a valid grouping. CONTACTS_HANDLED Unit: Count
	// Valid metric filter key: INITIATION_METHOD , DISCONNECT_REASON Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy, Feature,
	// contact/segmentAttributes/connect:Subtype Feature is a valid filter but not a
	// valid grouping. CONTACTS_HOLD_ABANDONS Unit: Count Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype CONTACTS_ON_HOLD_AGENT_DISCONNECT
	// Unit: Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy CONTACTS_ON_HOLD_CUSTOMER_DISCONNECT Unit: Count Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// CONTACTS_PUT_ON_HOLD Unit: Count Valid groupings and filters: Queue, Channel,
	// Routing Profile, Agent, Agent Hierarchy CONTACTS_TRANSFERRED_OUT_EXTERNAL Unit:
	// Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent, Agent
	// Hierarchy CONTACTS_TRANSFERRED_OUT_INTERNAL Unit: Percent Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy CONTACTS_QUEUED
	// Unit: Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// CONTACTS_RESOLVED_IN_X Unit: Count Valid groupings and filters: Queue, Channel,
	// Routing Profile, contact/segmentAttributes/connect:Subtype Threshold: For
	// ThresholdValue enter any whole number from 1 to 604800 (inclusive), in seconds.
	// For Comparison , you must enter LT (for "Less than"). CONTACTS_TRANSFERRED_OUT
	// Unit: Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy, Feature, contact/segmentAttributes/connect:Subtype Feature is a
	// valid filter but not a valid grouping. CONTACTS_TRANSFERRED_OUT_BY_AGENT Unit:
	// Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent, Agent
	// Hierarchy, contact/segmentAttributes/connect:Subtype
	// CONTACTS_TRANSFERRED_OUT_FROM_QUEUE Unit: Count Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype MAX_QUEUED_TIME Unit: Seconds Valid
	// groupings and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype PERCENT_NON_TALK_TIME This metric is
	// available only for contacts analyzed by Contact Lens conversational analytics.
	// Unit: Percentage Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// PERCENT_TALK_TIME This metric is available only for contacts analyzed by Contact
	// Lens conversational analytics. Unit: Percentage Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype PERCENT_TALK_TIME_AGENT This metric is
	// available only for contacts analyzed by Contact Lens conversational analytics.
	// Unit: Percentage Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy, contact/segmentAttributes/connect:Subtype
	// PERCENT_TALK_TIME_CUSTOMER This metric is available only for contacts analyzed
	// by Contact Lens conversational analytics. Unit: Percentage Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy,
	// contact/segmentAttributes/connect:Subtype SERVICE_LEVEL You can include up to 20
	// SERVICE_LEVEL metrics in a request. Unit: Percent Valid groupings and filters:
	// Queue, Channel, Routing Profile Threshold: For ThresholdValue , enter any whole
	// number from 1 to 604800 (inclusive), in seconds. For Comparison , you must enter
	// LT (for "Less than"). SUM_AFTER_CONTACT_WORK_TIME Unit: Seconds Valid groupings
	// and filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// SUM_CONNECTING_TIME_AGENT Unit: Seconds Valid metric filter key:
	// INITIATION_METHOD . This metric only supports the following filter keys as
	// INITIATION_METHOD : INBOUND | OUTBOUND | CALLBACK | API Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy The Negate key
	// in Metric Level Filters is not applicable for this metric. SUM_CONTACT_FLOW_TIME
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy SUM_CONTACT_TIME_AGENT Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// SUM_CONTACTS_ANSWERED_IN_X Unit: Count Valid groupings and filters: Queue,
	// Channel, Routing Profile, contact/segmentAttributes/connect:Subtype Threshold:
	// For ThresholdValue , enter any whole number from 1 to 604800 (inclusive), in
	// seconds. For Comparison , you must enter LT (for "Less than").
	// SUM_CONTACTS_ABANDONED_IN_X Unit: Count Valid groupings and filters: Queue,
	// Channel, Routing Profile, contact/segmentAttributes/connect:Subtype Threshold:
	// For ThresholdValue , enter any whole number from 1 to 604800 (inclusive), in
	// seconds. For Comparison , you must enter LT (for "Less than").
	// SUM_CONTACTS_DISCONNECTED Valid metric filter key: DISCONNECT_REASON Unit:
	// Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent, Agent
	// Hierarchy, contact/segmentAttributes/connect:Subtype SUM_ERROR_STATUS_TIME_AGENT
	// Unit: Seconds Valid groupings and filters: Queue, Channel, Routing Profile,
	// Agent, Agent Hierarchy SUM_HANDLE_TIME Unit: Seconds Valid groupings and
	// filters: Queue, Channel, Routing Profile, Agent, Agent Hierarchy SUM_HOLD_TIME
	// Unit: Count Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy SUM_IDLE_TIME_AGENT Unit: Seconds Valid groupings and filters:
	// Routing Profile, Agent, Agent Hierarchy SUM_INTERACTION_AND_HOLD_TIME Unit:
	// Seconds Valid groupings and filters: Queue, Channel, Routing Profile, Agent,
	// Agent Hierarchy SUM_INTERACTION_TIME Unit: Seconds Valid groupings and filters:
	// Queue, Channel, Routing Profile, Agent, Agent Hierarchy
	// SUM_NON_PRODUCTIVE_TIME_AGENT Unit: Seconds Valid groupings and filters: Routing
	// Profile, Agent, Agent Hierarchy SUM_ONLINE_TIME_AGENT Unit: Seconds Valid
	// groupings and filters: Routing Profile, Agent, Agent Hierarchy
	// SUM_RETRY_CALLBACK_ATTEMPTS Unit: Count Valid groupings and filters: Queue,
	// Channel, Routing Profile, contact/segmentAttributes/connect:Subtype
	//
	// This member is required.
	Metrics []types.MetricV2

	// The Amazon Resource Name (ARN) of the resource. This includes the instanceId an
	// Amazon Connect instance.
	//
	// This member is required.
	ResourceArn *string

	// The timestamp, in UNIX Epoch time format, at which to start the reporting
	// interval for the retrieval of historical metrics data. The time must be before
	// the end time timestamp. The start and end time depends on the IntervalPeriod
	// selected. By default the time range between start and end time is 35 days.
	// Historical metrics are available for 3 months.
	//
	// This member is required.
	StartTime *time.Time

	// The grouping applied to the metrics that are returned. For example, when
	// results are grouped by queue, the metrics returned are grouped by queue. The
	// values that are returned apply to the metrics for each queue. They are not
	// aggregated for all queues. If no grouping is specified, a summary of all metrics
	// is returned. Valid grouping keys: QUEUE | ROUTING_PROFILE | AGENT | CHANNEL |
	// AGENT_HIERARCHY_LEVEL_ONE | AGENT_HIERARCHY_LEVEL_TWO |
	// AGENT_HIERARCHY_LEVEL_THREE | AGENT_HIERARCHY_LEVEL_FOUR |
	// AGENT_HIERARCHY_LEVEL_FIVE , contact/segmentAttributes/connect:Subtype
	Groupings []string

	// The interval period and timezone to apply to returned metrics.
	//   - IntervalPeriod : An aggregated grouping applied to request metrics. Valid
	//   IntervalPeriod values are: FIFTEEN_MIN | THIRTY_MIN | HOUR | DAY | WEEK |
	//   TOTAL . For example, if IntervalPeriod is selected THIRTY_MIN , StartTime and
	//   EndTime differs by 1 day, then Amazon Connect returns 48 results in the
	//   response. Each result is aggregated by the THIRTY_MIN period. By default Amazon
	//   Connect aggregates results based on the TOTAL interval period. The following
	//   list describes restrictions on StartTime and EndTime based on which
	//   IntervalPeriod is requested.
	//   - FIFTEEN_MIN : The difference between StartTime and EndTime must be less than
	//   3 days.
	//   - THIRTY_MIN : The difference between StartTime and EndTime must be less than
	//   3 days.
	//   - HOUR : The difference between StartTime and EndTime must be less than 3
	//   days.
	//   - DAY : The difference between StartTime and EndTime must be less than 35
	//   days.
	//   - WEEK : The difference between StartTime and EndTime must be less than 35
	//   days.
	//   - TOTAL : The difference between StartTime and EndTime must be less than 35
	//   days.
	//   - TimeZone : The timezone applied to requested metrics.
	Interval *types.IntervalDetails

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetMetricDataV2Output struct {

	// Information about the metrics requested in the API request If no grouping is
	// specified, a summary of metric data is returned.
	MetricResults []types.MetricResultV2

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetricDataV2Middlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMetricDataV2{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMetricDataV2{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMetricDataV2"); err != nil {
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
	if err = addOpGetMetricDataV2ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricDataV2(options.Region), middleware.Before); err != nil {
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

// GetMetricDataV2APIClient is a client that implements the GetMetricDataV2
// operation.
type GetMetricDataV2APIClient interface {
	GetMetricDataV2(context.Context, *GetMetricDataV2Input, ...func(*Options)) (*GetMetricDataV2Output, error)
}

var _ GetMetricDataV2APIClient = (*Client)(nil)

// GetMetricDataV2PaginatorOptions is the paginator options for GetMetricDataV2
type GetMetricDataV2PaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMetricDataV2Paginator is a paginator for GetMetricDataV2
type GetMetricDataV2Paginator struct {
	options   GetMetricDataV2PaginatorOptions
	client    GetMetricDataV2APIClient
	params    *GetMetricDataV2Input
	nextToken *string
	firstPage bool
}

// NewGetMetricDataV2Paginator returns a new GetMetricDataV2Paginator
func NewGetMetricDataV2Paginator(client GetMetricDataV2APIClient, params *GetMetricDataV2Input, optFns ...func(*GetMetricDataV2PaginatorOptions)) *GetMetricDataV2Paginator {
	if params == nil {
		params = &GetMetricDataV2Input{}
	}

	options := GetMetricDataV2PaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMetricDataV2Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMetricDataV2Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetMetricDataV2 page.
func (p *GetMetricDataV2Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMetricDataV2Output, error) {
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

	result, err := p.client.GetMetricDataV2(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetMetricDataV2(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMetricDataV2",
	}
}
