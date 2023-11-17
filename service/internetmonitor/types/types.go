// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Amazon CloudWatch Internet Monitor calculates measurements about the
// availability for your application's internet traffic between client locations
// and Amazon Web Services. Amazon Web Services has substantial historical data
// about internet performance and availability between Amazon Web Services services
// and different network providers and geographies. By applying statistical
// analysis to the data, Internet Monitor can detect when the performance and
// availability for your application has dropped, compared to an estimated baseline
// that's already calculated. To make it easier to see those drops, we report that
// information to you in the form of health scores: a performance score and an
// availability score. Availability in Internet Monitor represents the estimated
// percentage of traffic that is not seeing an availability drop. For example, an
// availability score of 99% for an end user and service location pair is
// equivalent to 1% of the traffic experiencing an availability drop for that pair.
// For more information, see How Internet Monitor calculates performance and
// availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
// in the Amazon CloudWatch Internet Monitor section of the Amazon CloudWatch User
// Guide.
type AvailabilityMeasurement struct {

	// Experience scores, or health scores are calculated for different geographic and
	// network provider combinations (that is, different granularities) and also summed
	// into global scores. If you view performance or availability scores without
	// filtering for any specific geography or service provider, Amazon CloudWatch
	// Internet Monitor provides global health scores. The Amazon CloudWatch Internet
	// Monitor chapter in the CloudWatch User Guide includes detailed information about
	// how Internet Monitor calculates health scores, including performance and
	// availability scores, and when it creates and resolves health events. For more
	// information, see How Amazon Web Services calculates performance and
	// availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	ExperienceScore *float64

	// The percentage of impact caused by a health event for client location traffic
	// globally. For information about how Internet Monitor calculates impact, see
	// Inside Internet Monitor (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html)
	// in the Amazon CloudWatch Internet Monitor section of the Amazon CloudWatch User
	// Guide.
	PercentOfClientLocationImpacted *float64

	// The impact on total traffic that a health event has, in increased latency or
	// reduced availability. This is the percentage of how much latency has increased
	// or availability has decreased during the event, compared to what is typical for
	// traffic from this client location to the Amazon Web Services location using this
	// client network. For information about how Internet Monitor calculates impact,
	// see How Internet Monitor works (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html)
	// in the Amazon CloudWatch Internet Monitor section of the Amazon CloudWatch User
	// Guide.
	PercentOfTotalTrafficImpacted *float64

	noSmithyDocumentSerde
}

// A filter that you use with the results of a Amazon CloudWatch Internet Monitor
// query that you created and ran. The query sets up a repository of data that is a
// subset of your application's Internet Monitor data. FilterParameter is a string
// that defines how you want to filter the repository of data to return a set of
// results, based on your criteria. The filter parameters that you can specify
// depend on the query type that you used to create the repository, since each
// query type returns a different set of Internet Monitor data. For each filter,
// you specify a field (such as city ), an operator (such as not_equals , and a
// value or array of values (such as ["Seattle", "Redmond"] ). Separate values in
// the array with commas. For more information about specifying filter parameters,
// see Using the Amazon CloudWatch Internet Monitor query interface (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-view-cw-tools-cwim-query.html)
// in the Amazon CloudWatch Internet Monitor User Guide.
type FilterParameter struct {

	// A data field that you want to filter, to further scope your application's
	// Internet Monitor data in a repository that you created by running a query. A
	// field might be city , for example. The field must be one of the fields that was
	// returned by the specific query that you used to create the repository.
	Field *string

	// The operator to use with the filter field and a value, such as not_equals .
	Operator Operator

	// One or more values to be used, together with the specified operator, to filter
	// data for a query. For example, you could specify an array of values such as
	// ["Seattle", "Redmond"] . Values in the array are separated by commas.
	Values []string

	noSmithyDocumentSerde
}

// Information about a health event created in a monitor in Amazon CloudWatch
// Internet Monitor.
type HealthEvent struct {

	// The Amazon Resource Name (ARN) of the event.
	//
	// This member is required.
	EventArn *string

	// The internally-generated identifier of a specific network traffic impairment
	// health event.
	//
	// This member is required.
	EventId *string

	// The type of impairment for a health event.
	//
	// This member is required.
	ImpactType HealthEventImpactType

	// The locations impacted by the health event.
	//
	// This member is required.
	ImpactedLocations []ImpactedLocation

	// When the health event was last updated.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// When a health event started.
	//
	// This member is required.
	StartedAt *time.Time

	// Health event list member.
	//
	// This member is required.
	Status HealthEventStatus

	// When the health event was created.
	CreatedAt *time.Time

	// The time when a health event ended. If the health event is still active, then
	// the end time is not set.
	EndedAt *time.Time

	// The value of the threshold percentage for performance or availability that was
	// configured when Amazon CloudWatch Internet Monitor created the health event.
	HealthScoreThreshold float64

	// The impact on total traffic that a health event has, in increased latency or
	// reduced availability. This is the percentage of how much latency has increased
	// or availability has decreased during the event, compared to what is typical for
	// traffic from this client location to the Amazon Web Services location using this
	// client network.
	PercentOfTotalTrafficImpacted *float64

	noSmithyDocumentSerde
}

// A complex type with the configuration information that determines the threshold
// and other conditions for when Internet Monitor creates a health event for an
// overall performance or availability issue, across an application's geographies.
// Defines the percentages, for overall performance scores and availability scores
// for an application, that are the thresholds for when Amazon CloudWatch Internet
// Monitor creates a health event. You can override the defaults to set a custom
// threshold for overall performance or availability scores, or both. You can also
// set thresholds for local health scores,, where Internet Monitor creates a health
// event when scores cross a threshold for one or more city-networks, in addition
// to creating an event when an overall score crosses a threshold. If you don't set
// a health event threshold, the default value is 95%. For local thresholds, you
// also set a minimum percentage of overall traffic that is impacted by an issue
// before Internet Monitor creates an event. In addition, you can disable local
// thresholds, for performance scores, availability scores, or both. For more
// information, see Change health event thresholds (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview)
// in the Internet Monitor section of the CloudWatch User Guide.
type HealthEventsConfig struct {

	// The configuration that determines the threshold and other conditions for when
	// Internet Monitor creates a health event for a local availability issue.
	AvailabilityLocalHealthEventsConfig *LocalHealthEventsConfig

	// The health event threshold percentage set for availability scores.
	AvailabilityScoreThreshold float64

	// The configuration that determines the threshold and other conditions for when
	// Internet Monitor creates a health event for a local performance issue.
	PerformanceLocalHealthEventsConfig *LocalHealthEventsConfig

	// The health event threshold percentage set for performance scores.
	PerformanceScoreThreshold float64

	noSmithyDocumentSerde
}

// Information about a location impacted by a health event in Amazon CloudWatch
// Internet Monitor. Geographic regions are hierarchically categorized into
// country, subdivision, metro and city geographic granularities. The geographic
// region is identified based on the IP address used at the client locations.
type ImpactedLocation struct {

	// The name of the network at an impacted location.
	//
	// This member is required.
	ASName *string

	// The Autonomous System Number (ASN) of the network at an impacted location.
	//
	// This member is required.
	ASNumber *int64

	// The name of the country where the health event is located.
	//
	// This member is required.
	Country *string

	// The status of the health event at an impacted location.
	//
	// This member is required.
	Status HealthEventStatus

	// The cause of the impairment. There are two types of network impairments: Amazon
	// Web Services network issues or internet issues. Internet issues are typically a
	// problem with a network provider, like an internet service provider (ISP).
	CausedBy *NetworkImpairment

	// The name of the city where the health event is located.
	City *string

	// The country code where the health event is located. The ISO 3166-2 codes for
	// the country is provided, when available.
	CountryCode *string

	// The calculated health at a specific location.
	InternetHealth *InternetHealth

	// The latitude where the health event is located.
	Latitude *float64

	// The longitude where the health event is located.
	Longitude *float64

	// The metro area where the health event is located. Metro indicates a
	// metropolitan region in the United States, such as the region around New York
	// City. In non-US countries, this is a second-level subdivision. For example, in
	// the United Kingdom, it could be a county, a London borough, a unitary authority,
	// council area, and so on.
	Metro *string

	// The service location where the health event is located.
	ServiceLocation *string

	// The subdivision location where the health event is located. The subdivision
	// usually maps to states in most countries (including the United States). For
	// United Kingdom, it maps to a country (England, Scotland, Wales) or province
	// (Northern Ireland).
	Subdivision *string

	// The subdivision code where the health event is located. The ISO 3166-2 codes
	// for country subdivisions is provided, when available.
	SubdivisionCode *string

	noSmithyDocumentSerde
}

// Internet health includes measurements calculated by Amazon CloudWatch Internet
// Monitor about the performance and availability for your application on the
// internet. Amazon Web Services has substantial historical data about internet
// performance and availability between Amazon Web Services services and different
// network providers and geographies. By applying statistical analysis to the data,
// Internet Monitor can detect when the performance and availability for your
// application has dropped, compared to an estimated baseline that's already
// calculated. To make it easier to see those drops, Internet Monitor reports the
// information to you in the form of health scores: a performance score and an
// availability score.
type InternetHealth struct {

	// Availability in Internet Monitor represents the estimated percentage of traffic
	// that is not seeing an availability drop. For example, an availability score of
	// 99% for an end user and service location pair is equivalent to 1% of the traffic
	// experiencing an availability drop for that pair. For more information, see How
	// Internet Monitor calculates performance and availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	Availability *AvailabilityMeasurement

	// Performance in Internet Monitor represents the estimated percentage of traffic
	// that is not seeing a performance drop. For example, a performance score of 99%
	// for an end user and service location pair is equivalent to 1% of the traffic
	// experiencing a performance drop for that pair. For more information, see How
	// Internet Monitor calculates performance and availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	Performance *PerformanceMeasurement

	noSmithyDocumentSerde
}

// Publish internet measurements to an Amazon S3 bucket in addition to CloudWatch
// Logs.
type InternetMeasurementsLogDelivery struct {

	// The configuration information for publishing Internet Monitor internet
	// measurements to Amazon S3. The configuration includes the bucket name and
	// (optionally) prefix for the S3 bucket to store the measurements, and the
	// delivery status. The delivery status is ENABLED or DISABLED , depending on
	// whether you choose to deliver internet measurements to S3 logs.
	S3Config *S3Config

	noSmithyDocumentSerde
}

// A complex type with the configuration information that determines the threshold
// and other conditions for when Internet Monitor creates a health event for a
// local performance or availability issue, when scores cross a threshold for one
// or more city-networks. Defines the percentages, for performance scores or
// availability scores, that are the local thresholds for when Amazon CloudWatch
// Internet Monitor creates a health event. Also defines whether a local threshold
// is enabled or disabled, and the minimum percentage of overall traffic that must
// be impacted by an issue before Internet Monitor creates an event when a
// threshold is crossed for a local health score. If you don't set a local health
// event threshold, the default value is 60%. For more information, see Change
// health event thresholds (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-overview.html#IMUpdateThresholdFromOverview)
// in the Internet Monitor section of the CloudWatch User Guide.
type LocalHealthEventsConfig struct {

	// The health event threshold percentage set for a local health score.
	HealthScoreThreshold float64

	// The minimum percentage of overall traffic for an application that must be
	// impacted by an issue before Internet Monitor creates an event when a threshold
	// is crossed for a local health score. If you don't set a minimum traffic impact
	// threshold, the default value is 0.01%.
	MinTrafficImpact float64

	// The status of whether Internet Monitor creates a health event based on a
	// threshold percentage set for a local health score. The status can be ENABLED or
	// DISABLED .
	Status LocalHealthEventsConfigStatus

	noSmithyDocumentSerde
}

// The description of and information about a monitor in Amazon CloudWatch
// Internet Monitor.
type Monitor struct {

	// The Amazon Resource Name (ARN) of the monitor.
	//
	// This member is required.
	MonitorArn *string

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// The status of a monitor.
	//
	// This member is required.
	Status MonitorConfigState

	// The health of data processing for the monitor.
	ProcessingStatus MonitorProcessingStatusCode

	noSmithyDocumentSerde
}

// An internet service provider (ISP) or network in Amazon CloudWatch Internet
// Monitor.
type Network struct {

	// The internet provider name or network name.
	//
	// This member is required.
	ASName *string

	// The Autonomous System Number (ASN) of the internet provider or network.
	//
	// This member is required.
	ASNumber *int64

	noSmithyDocumentSerde
}

// Information about the network impairment for a specific network measured by
// Amazon CloudWatch Internet Monitor.
type NetworkImpairment struct {

	// The combination of the Autonomous System Number (ASN) of the network and the
	// name of the network.
	//
	// This member is required.
	AsPath []Network

	// Type of network impairment.
	//
	// This member is required.
	NetworkEventType TriangulationEventType

	// The networks that could be impacted by a network impairment event.
	//
	// This member is required.
	Networks []Network

	noSmithyDocumentSerde
}

// Amazon CloudWatch Internet Monitor calculates measurements about the
// performance for your application's internet traffic between client locations and
// Amazon Web Services. Amazon Web Services has substantial historical data about
// internet performance and availability between Amazon Web Services services and
// different network providers and geographies. By applying statistical analysis to
// the data, Internet Monitor can detect when the performance and availability for
// your application has dropped, compared to an estimated baseline that's already
// calculated. To make it easier to see those drops, we report that information to
// you in the form of health scores: a performance score and an availability score.
// Performance in Internet Monitor represents the estimated percentage of traffic
// that is not seeing a performance drop. For example, a performance score of 99%
// for an end user and service location pair is equivalent to 1% of the traffic
// experiencing a performance drop for that pair. For more information, see How
// Internet Monitor calculates performance and availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
type PerformanceMeasurement struct {

	// Experience scores, or health scores, are calculated for different geographic
	// and network provider combinations (that is, different granularities) and also
	// totaled into global scores. If you view performance or availability scores
	// without filtering for any specific geography or service provider, Amazon
	// CloudWatch Internet Monitor provides global health scores. The Amazon CloudWatch
	// Internet Monitor chapter in the CloudWatch User Guide includes detailed
	// information about how Internet Monitor calculates health scores, including
	// performance and availability scores, and when it creates and resolves health
	// events. For more information, see How Amazon Web Services calculates
	// performance and availability scores (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMExperienceScores)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	ExperienceScore *float64

	// How much performance impact was caused by a health event at a client location.
	// For performance, this is the percentage of how much latency increased during the
	// event compared to typical performance for traffic, from this client location to
	// an Amazon Web Services location, using a specific client network. For more
	// information, see When Amazon Web Services creates and resolves health events (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMHealthEventStartStop)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	PercentOfClientLocationImpacted *float64

	// The impact on total traffic that a health event has, in increased latency or
	// reduced availability. This is the percentage of how much latency has increased
	// or availability has decreased during the event, compared to what is typical for
	// traffic from this client location to the Amazon Web Services location using this
	// client network. For more information, see When Amazon Web Services creates and
	// resolves health events (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMHealthEventStartStop)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	PercentOfTotalTrafficImpacted *float64

	// This is the percentage of how much round-trip time increased during the event
	// compared to typical round-trip time for your application for traffic. For more
	// information, see When Amazon Web Services creates and resolves health events (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-inside-internet-monitor.html#IMHealthEventStartStop)
	// in the Amazon CloudWatch Internet Monitor section of the CloudWatch User Guide.
	RoundTripTime *RoundTripTime

	noSmithyDocumentSerde
}

// Defines a field to query for your application's Amazon CloudWatch Internet
// Monitor data. You create a data repository by running a query of a specific
// type. Each QueryType includes a specific set of fields and datatypes to
// retrieve data for.
type QueryField struct {

	// The name of a field to query your application's Amazon CloudWatch Internet
	// Monitor data for, such as availability_score .
	Name *string

	// The data type for a query field, which must correspond to the field you're
	// defining for QueryField . For example, if the query field name is
	// availability_score , the data type is float .
	Type *string

	noSmithyDocumentSerde
}

// Round-trip time (RTT) is how long it takes for a request from the user to
// return a response to the user. Amazon CloudWatch Internet Monitor calculates RTT
// at different percentiles: p50, p90, and p95.
type RoundTripTime struct {

	// RTT at the 50th percentile (p50).
	P50 *float64

	// RTT at the 90th percentile (p90).
	P90 *float64

	// RTT at the 95th percentile (p95).
	P95 *float64

	noSmithyDocumentSerde
}

// The configuration for publishing Amazon CloudWatch Internet Monitor internet
// measurements to Amazon S3. The configuration includes the bucket name and
// (optionally) prefix for the S3 bucket to store the measurements, and the
// delivery status. The delivery status is ENABLED or DISABLED , depending on
// whether you choose to deliver internet measurements to S3 logs.
type S3Config struct {

	// The Amazon S3 bucket name.
	BucketName *string

	// The Amazon S3 bucket prefix.
	BucketPrefix *string

	// The status of publishing Internet Monitor internet measurements to an Amazon S3
	// bucket.
	LogDeliveryStatus LogDeliveryStatus

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
