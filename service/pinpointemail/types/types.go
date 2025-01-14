// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object that contains information about a blacklisting event that impacts one
// of the dedicated IP addresses that is associated with your account.
type BlacklistEntry struct {

	// Additional information about the blacklisting event, as provided by the
	// blacklist maintainer.
	Description *string

	// The time when the blacklisting event occurred, shown in Unix time format.
	ListingTime *time.Time

	// The name of the blacklist that the IP address appears on.
	RblName *string

	noSmithyDocumentSerde
}

// Represents the body of the email message.
type Body struct {

	// An object that represents the version of the message that is displayed in email
	// clients that support HTML. HTML messages can include formatted text, hyperlinks,
	// images, and more.
	Html *Content

	// An object that represents the version of the message that is displayed in email
	// clients that don't support HTML, or clients where the recipient has disabled
	// HTML rendering.
	Text *Content

	noSmithyDocumentSerde
}

// An object that defines an Amazon CloudWatch destination for email events. You
// can use Amazon CloudWatch to monitor and gain insights on your email sending
// metrics.
type CloudWatchDestination struct {

	// An array of objects that define the dimensions to use when you send email events
	// to Amazon CloudWatch.
	//
	// This member is required.
	DimensionConfigurations []CloudWatchDimensionConfiguration

	noSmithyDocumentSerde
}

// An object that defines the dimension configuration to use when you send Amazon
// Pinpoint email events to Amazon CloudWatch.
type CloudWatchDimensionConfiguration struct {

	// The default value of the dimension that is published to Amazon CloudWatch if you
	// don't provide the value of the dimension when you send an email. This value has
	// to meet the following criteria:
	//
	// * It can only contain ASCII letters (a-z, A-Z),
	// numbers (0-9), underscores (_), or dashes (-).
	//
	// * It can contain no more than
	// 256 characters.
	//
	// This member is required.
	DefaultDimensionValue *string

	// The name of an Amazon CloudWatch dimension associated with an email sending
	// metric. The name has to meet the following criteria:
	//
	// * It can only contain
	// ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	//
	// * It
	// can contain no more than 256 characters.
	//
	// This member is required.
	DimensionName *string

	// The location where Amazon Pinpoint finds the value of a dimension to publish to
	// Amazon CloudWatch. If you want Amazon Pinpoint to use the message tags that you
	// specify using an X-SES-MESSAGE-TAGS header or a parameter to the
	// SendEmail/SendRawEmail API, choose messageTag. If you want Amazon Pinpoint to
	// use your own email headers, choose emailHeader. If you want Amazon Pinpoint to
	// use link tags, choose linkTags.
	//
	// This member is required.
	DimensionValueSource DimensionValueSource

	noSmithyDocumentSerde
}

// An object that represents the content of the email, and optionally a character
// set specification.
type Content struct {

	// The content of the message itself.
	//
	// This member is required.
	Data *string

	// The character set for the content. Because of the constraints of the SMTP
	// protocol, Amazon Pinpoint uses 7-bit ASCII by default. If the text includes
	// characters outside of the ASCII range, you have to specify a character set. For
	// example, you could specify UTF-8, ISO-8859-1, or Shift_JIS.
	Charset *string

	noSmithyDocumentSerde
}

// An object that contains information about the volume of email sent on each day
// of the analysis period.
type DailyVolume struct {

	// An object that contains inbox placement metrics for a specified day in the
	// analysis period, broken out by the recipient's email provider.
	DomainIspPlacements []DomainIspPlacement

	// The date that the DailyVolume metrics apply to, in Unix time.
	StartDate *time.Time

	// An object that contains inbox placement metrics for a specific day in the
	// analysis period.
	VolumeStatistics *VolumeStatistics

	noSmithyDocumentSerde
}

// Contains information about a dedicated IP address that is associated with your
// Amazon Pinpoint account.
type DedicatedIp struct {

	// An IP address that is reserved for use by your Amazon Pinpoint account.
	//
	// This member is required.
	Ip *string

	// Indicates how complete the dedicated IP warm-up process is. When this value
	// equals 1, the address has completed the warm-up process and is ready for use.
	//
	// This member is required.
	WarmupPercentage *int32

	// The warm-up status of a dedicated IP address. The status can have one of the
	// following values:
	//
	// * IN_PROGRESS – The IP address isn't ready to use because the
	// dedicated IP warm-up process is ongoing.
	//
	// * DONE – The dedicated IP warm-up
	// process is complete, and the IP address is ready to use.
	//
	// This member is required.
	WarmupStatus WarmupStatus

	// The name of the dedicated IP pool that the IP address is associated with.
	PoolName *string

	noSmithyDocumentSerde
}

// An object that contains metadata related to a predictive inbox placement test.
type DeliverabilityTestReport struct {

	// The date and time when the predictive inbox placement test was created, in Unix
	// time format.
	CreateDate *time.Time

	// The status of the predictive inbox placement test. If the status is IN_PROGRESS,
	// then the predictive inbox placement test is currently running. Predictive inbox
	// placement tests are usually complete within 24 hours of creating the test. If
	// the status is COMPLETE, then the test is finished, and you can use the
	// GetDeliverabilityTestReport to view the results of the test.
	DeliverabilityTestStatus DeliverabilityTestStatus

	// The sender address that you specified for the predictive inbox placement test.
	FromEmailAddress *string

	// A unique string that identifies the predictive inbox placement test.
	ReportId *string

	// A name that helps you identify a predictive inbox placement test report.
	ReportName *string

	// The subject line for an email that you submitted in a predictive inbox placement
	// test.
	Subject *string

	noSmithyDocumentSerde
}

// Used to associate a configuration set with a dedicated IP pool.
type DeliveryOptions struct {

	// The name of the dedicated IP pool that you want to associate with the
	// configuration set.
	SendingPoolName *string

	// Specifies whether messages that use the configuration set are required to use
	// Transport Layer Security (TLS). If the value is Require, messages are only
	// delivered if a TLS connection can be established. If the value is Optional,
	// messages can be delivered in plain text if a TLS connection can't be
	// established.
	TlsPolicy TlsPolicy

	noSmithyDocumentSerde
}

// An object that describes the recipients for an email.
type Destination struct {

	// An array that contains the email addresses of the "BCC" (blind carbon copy)
	// recipients for the email.
	BccAddresses []string

	// An array that contains the email addresses of the "CC" (carbon copy) recipients
	// for the email.
	CcAddresses []string

	// An array that contains the email addresses of the "To" recipients for the email.
	ToAddresses []string

	noSmithyDocumentSerde
}

// An object that contains information about the DKIM configuration for an email
// identity.
type DkimAttributes struct {

	// If the value is true, then the messages that Amazon Pinpoint sends from the
	// identity are DKIM-signed. If the value is false, then the messages that Amazon
	// Pinpoint sends from the identity aren't DKIM-signed.
	SigningEnabled bool

	// Describes whether or not Amazon Pinpoint has successfully located the DKIM
	// records in the DNS records for the domain. The status can be one of the
	// following:
	//
	// * PENDING – Amazon Pinpoint hasn't yet located the DKIM records in
	// the DNS configuration for the domain, but will continue to attempt to locate
	// them.
	//
	// * SUCCESS – Amazon Pinpoint located the DKIM records in the DNS
	// configuration for the domain and determined that they're correct. Amazon
	// Pinpoint can now send DKIM-signed email from the identity.
	//
	// * FAILED – Amazon
	// Pinpoint was unable to locate the DKIM records in the DNS settings for the
	// domain, and won't continue to search for them.
	//
	// * TEMPORARY_FAILURE – A
	// temporary issue occurred, which prevented Amazon Pinpoint from determining the
	// DKIM status for the domain.
	//
	// * NOT_STARTED – Amazon Pinpoint hasn't yet started
	// searching for the DKIM records in the DKIM records for the domain.
	Status DkimStatus

	// A set of unique strings that you use to create a set of CNAME records that you
	// add to the DNS configuration for your domain. When Amazon Pinpoint detects these
	// records in the DNS configuration for your domain, the DKIM authentication
	// process is complete. Amazon Pinpoint usually detects these records within about
	// 72 hours of adding them to the DNS configuration for your domain.
	Tokens []string

	noSmithyDocumentSerde
}

// An object that contains the deliverability data for a specific campaign. This
// data is available for a campaign only if the campaign sent email by using a
// domain that the Deliverability dashboard is enabled for
// (PutDeliverabilityDashboardOption operation).
type DomainDeliverabilityCampaign struct {

	// The unique identifier for the campaign. Amazon Pinpoint automatically generates
	// and assigns this identifier to a campaign. This value is not the same as the
	// campaign identifier that Amazon Pinpoint assigns to campaigns that you create
	// and manage by using the Amazon Pinpoint API or the Amazon Pinpoint console.
	CampaignId *string

	// The percentage of email messages that were deleted by recipients, without being
	// opened first. Due to technical limitations, this value only includes recipients
	// who opened the message by using an email client that supports images.
	DeleteRate *float64

	// The major email providers who handled the email message.
	Esps []string

	// The first time, in Unix time format, when the email message was delivered to any
	// recipient's inbox. This value can help you determine how long it took for a
	// campaign to deliver an email message.
	FirstSeenDateTime *time.Time

	// The verified email address that the email message was sent from.
	FromAddress *string

	// The URL of an image that contains a snapshot of the email message that was sent.
	ImageUrl *string

	// The number of email messages that were delivered to recipients’ inboxes.
	InboxCount *int64

	// The last time, in Unix time format, when the email message was delivered to any
	// recipient's inbox. This value can help you determine how long it took for a
	// campaign to deliver an email message.
	LastSeenDateTime *time.Time

	// The projected number of recipients that the email message was sent to.
	ProjectedVolume *int64

	// The percentage of email messages that were opened and then deleted by
	// recipients. Due to technical limitations, this value only includes recipients
	// who opened the message by using an email client that supports images.
	ReadDeleteRate *float64

	// The percentage of email messages that were opened by recipients. Due to
	// technical limitations, this value only includes recipients who opened the
	// message by using an email client that supports images.
	ReadRate *float64

	// The IP addresses that were used to send the email message.
	SendingIps []string

	// The number of email messages that were delivered to recipients' spam or junk
	// mail folders.
	SpamCount *int64

	// The subject line, or title, of the email message.
	Subject *string

	noSmithyDocumentSerde
}

// An object that contains information about the Deliverability dashboard
// subscription for a verified domain that you use to send email and currently has
// an active Deliverability dashboard subscription. If a Deliverability dashboard
// subscription is active for a domain, you gain access to reputation, inbox
// placement, and other metrics for the domain.
type DomainDeliverabilityTrackingOption struct {

	// A verified domain that’s associated with your AWS account and currently has an
	// active Deliverability dashboard subscription.
	Domain *string

	// An object that contains information about the inbox placement data settings for
	// the domain.
	InboxPlacementTrackingOption *InboxPlacementTrackingOption

	// The date, in Unix time format, when you enabled the Deliverability dashboard for
	// the domain.
	SubscriptionStartDate *time.Time

	noSmithyDocumentSerde
}

// An object that contains inbox placement data for email sent from one of your
// email domains to a specific email provider.
type DomainIspPlacement struct {

	// The percentage of messages that were sent from the selected domain to the
	// specified email provider that arrived in recipients' inboxes.
	InboxPercentage *float64

	// The total number of messages that were sent from the selected domain to the
	// specified email provider that arrived in recipients' inboxes.
	InboxRawCount *int64

	// The name of the email provider that the inbox placement data applies to.
	IspName *string

	// The percentage of messages that were sent from the selected domain to the
	// specified email provider that arrived in recipients' spam or junk mail folders.
	SpamPercentage *float64

	// The total number of messages that were sent from the selected domain to the
	// specified email provider that arrived in recipients' spam or junk mail folders.
	SpamRawCount *int64

	noSmithyDocumentSerde
}

// An object that defines the entire content of the email, including the message
// headers and the body content. You can create a simple email message, in which
// you specify the subject and the text and HTML versions of the message body. You
// can also create raw messages, in which you specify a complete MIME-formatted
// message. Raw messages can include attachments and custom headers.
type EmailContent struct {

	// The raw email message. The message has to meet the following criteria:
	//
	// * The
	// message has to contain a header and a body, separated by one blank line.
	//
	// * All
	// of the required header fields must be present in the message.
	//
	// * Each part of a
	// multipart MIME message must be formatted properly.
	//
	// * If you include
	// attachments, they must be in a file format that Amazon Pinpoint supports.
	//
	// * The
	// entire message must be Base64 encoded.
	//
	// * If any of the MIME parts in your
	// message contain content that is outside of the 7-bit ASCII character range, you
	// should encode that content to ensure that recipients' email clients render the
	// message properly.
	//
	// * The length of any single line of text in the message can't
	// exceed 1,000 characters. This restriction is defined in RFC 5321
	// (https://tools.ietf.org/html/rfc5321).
	Raw *RawMessage

	// The simple email message. The message consists of a subject and a message body.
	Simple *Message

	// The template to use for the email message.
	Template *Template

	noSmithyDocumentSerde
}

// In Amazon Pinpoint, events include message sends, deliveries, opens, clicks,
// bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints, or
// you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
type EventDestination struct {

	// The types of events that Amazon Pinpoint sends to the specified event
	// destinations.
	//
	// This member is required.
	MatchingEventTypes []EventType

	// A name that identifies the event destination.
	//
	// This member is required.
	Name *string

	// An object that defines an Amazon CloudWatch destination for email events. You
	// can use Amazon CloudWatch to monitor and gain insights on your email sending
	// metrics.
	CloudWatchDestination *CloudWatchDestination

	// If true, the event destination is enabled. When the event destination is
	// enabled, the specified event types are sent to the destinations in this
	// EventDestinationDefinition. If false, the event destination is disabled. When
	// the event destination is disabled, events aren't sent to the specified
	// destinations.
	Enabled bool

	// An object that defines an Amazon Kinesis Data Firehose destination for email
	// events. You can use Amazon Kinesis Data Firehose to stream data to other
	// services, such as Amazon S3 and Amazon Redshift.
	KinesisFirehoseDestination *KinesisFirehoseDestination

	// An object that defines a Amazon Pinpoint destination for email events. You can
	// use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You
	// can use these attributes to create segments for your campaigns.
	PinpointDestination *PinpointDestination

	// An object that defines an Amazon SNS destination for email events. You can use
	// Amazon SNS to send notification when certain email events occur.
	SnsDestination *SnsDestination

	noSmithyDocumentSerde
}

// An object that defines the event destination. Specifically, it defines which
// services receive events from emails sent using the configuration set that the
// event destination is associated with. Also defines the types of events that are
// sent to the event destination.
type EventDestinationDefinition struct {

	// An object that defines an Amazon CloudWatch destination for email events. You
	// can use Amazon CloudWatch to monitor and gain insights on your email sending
	// metrics.
	CloudWatchDestination *CloudWatchDestination

	// If true, the event destination is enabled. When the event destination is
	// enabled, the specified event types are sent to the destinations in this
	// EventDestinationDefinition. If false, the event destination is disabled. When
	// the event destination is disabled, events aren't sent to the specified
	// destinations.
	Enabled bool

	// An object that defines an Amazon Kinesis Data Firehose destination for email
	// events. You can use Amazon Kinesis Data Firehose to stream data to other
	// services, such as Amazon S3 and Amazon Redshift.
	KinesisFirehoseDestination *KinesisFirehoseDestination

	// An array that specifies which events Amazon Pinpoint should send to the
	// destinations in this EventDestinationDefinition.
	MatchingEventTypes []EventType

	// An object that defines a Amazon Pinpoint destination for email events. You can
	// use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You
	// can use these attributes to create segments for your campaigns.
	PinpointDestination *PinpointDestination

	// An object that defines an Amazon SNS destination for email events. You can use
	// Amazon SNS to send notification when certain email events occur.
	SnsDestination *SnsDestination

	noSmithyDocumentSerde
}

// Information about an email identity.
type IdentityInfo struct {

	// The address or domain of the identity.
	IdentityName *string

	// The email identity type. The identity type can be one of the following:
	//
	// *
	// EMAIL_ADDRESS – The identity is an email address.
	//
	// * DOMAIN – The identity is a
	// domain.
	//
	// * MANAGED_DOMAIN – The identity is a domain that is managed by AWS.
	IdentityType IdentityType

	// Indicates whether or not you can send email from the identity. In Amazon
	// Pinpoint, an identity is an email address or domain that you send email from.
	// Before you can send email from an identity, you have to demostrate that you own
	// the identity, and that you authorize Amazon Pinpoint to send email from that
	// identity.
	SendingEnabled bool

	noSmithyDocumentSerde
}

// An object that contains information about the inbox placement data settings for
// a verified domain that’s associated with your AWS account. This data is
// available only if you enabled the Deliverability dashboard for the domain
// (PutDeliverabilityDashboardOption operation).
type InboxPlacementTrackingOption struct {

	// Specifies whether inbox placement data is being tracked for the domain.
	Global bool

	// An array of strings, one for each major email provider that the inbox placement
	// data applies to.
	TrackedIsps []string

	noSmithyDocumentSerde
}

// An object that describes how email sent during the predictive inbox placement
// test was handled by a certain email provider.
type IspPlacement struct {

	// The name of the email provider that the inbox placement data applies to.
	IspName *string

	// An object that contains inbox placement metrics for a specific email provider.
	PlacementStatistics *PlacementStatistics

	noSmithyDocumentSerde
}

// An object that defines an Amazon Kinesis Data Firehose destination for email
// events. You can use Amazon Kinesis Data Firehose to stream data to other
// services, such as Amazon S3 and Amazon Redshift.
type KinesisFirehoseDestination struct {

	// The Amazon Resource Name (ARN) of the Amazon Kinesis Data Firehose stream that
	// Amazon Pinpoint sends email events to.
	//
	// This member is required.
	DeliveryStreamArn *string

	// The Amazon Resource Name (ARN) of the IAM role that Amazon Pinpoint uses when
	// sending email events to the Amazon Kinesis Data Firehose stream.
	//
	// This member is required.
	IamRoleArn *string

	noSmithyDocumentSerde
}

// A list of attributes that are associated with a MAIL FROM domain.
type MailFromAttributes struct {

	// The action that Amazon Pinpoint to takes if it can't read the required MX record
	// for a custom MAIL FROM domain. When you set this value to UseDefaultValue,
	// Amazon Pinpoint uses amazonses.com as the MAIL FROM domain. When you set this
	// value to RejectMessage, Amazon Pinpoint returns a MailFromDomainNotVerified
	// error, and doesn't attempt to deliver the email. These behaviors are taken when
	// the custom MAIL FROM domain configuration is in the Pending, Failed, and
	// TemporaryFailure states.
	//
	// This member is required.
	BehaviorOnMxFailure BehaviorOnMxFailure

	// The name of a domain that an email identity uses as a custom MAIL FROM domain.
	//
	// This member is required.
	MailFromDomain *string

	// The status of the MAIL FROM domain. This status can have the following
	// values:
	//
	// * PENDING – Amazon Pinpoint hasn't started searching for the MX record
	// yet.
	//
	// * SUCCESS – Amazon Pinpoint detected the required MX record for the MAIL
	// FROM domain.
	//
	// * FAILED – Amazon Pinpoint can't find the required MX record, or
	// the record no longer exists.
	//
	// * TEMPORARY_FAILURE – A temporary issue occurred,
	// which prevented Amazon Pinpoint from determining the status of the MAIL FROM
	// domain.
	//
	// This member is required.
	MailFromDomainStatus MailFromDomainStatus

	noSmithyDocumentSerde
}

// Represents the email message that you're sending. The Message object consists of
// a subject line and a message body.
type Message struct {

	// The body of the message. You can specify an HTML version of the message, a
	// text-only version of the message, or both.
	//
	// This member is required.
	Body *Body

	// The subject line of the email. The subject line can only contain 7-bit ASCII
	// characters. However, you can specify non-ASCII characters in the subject line by
	// using encoded-word syntax, as described in RFC 2047
	// (https://tools.ietf.org/html/rfc2047).
	//
	// This member is required.
	Subject *Content

	noSmithyDocumentSerde
}

// Contains the name and value of a tag that you apply to an email. You can use
// message tags when you publish email sending events.
type MessageTag struct {

	// The name of the message tag. The message tag name has to meet the following
	// criteria:
	//
	// * It can only contain ASCII letters (a–z, A–Z), numbers (0–9),
	// underscores (_), or dashes (-).
	//
	// * It can contain no more than 256 characters.
	//
	// This member is required.
	Name *string

	// The value of the message tag. The message tag value has to meet the following
	// criteria:
	//
	// * It can only contain ASCII letters (a–z, A–Z), numbers (0–9),
	// underscores (_), or dashes (-).
	//
	// * It can contain no more than 256 characters.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// An object that contains information about email that was sent from the selected
// domain.
type OverallVolume struct {

	// An object that contains inbox and junk mail placement metrics for individual
	// email providers.
	DomainIspPlacements []DomainIspPlacement

	// The percentage of emails that were sent from the domain that were read by their
	// recipients.
	ReadRatePercent *float64

	// An object that contains information about the numbers of messages that arrived
	// in recipients' inboxes and junk mail folders.
	VolumeStatistics *VolumeStatistics

	noSmithyDocumentSerde
}

// An object that defines a Amazon Pinpoint destination for email events. You can
// use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You
// can use these attributes to create segments for your campaigns.
type PinpointDestination struct {

	// The Amazon Resource Name (ARN) of the Amazon Pinpoint project that you want to
	// send email events to.
	ApplicationArn *string

	noSmithyDocumentSerde
}

// An object that contains inbox placement data for an email provider.
type PlacementStatistics struct {

	// The percentage of emails that were authenticated by using DomainKeys Identified
	// Mail (DKIM) during the predictive inbox placement test.
	DkimPercentage *float64

	// The percentage of emails that arrived in recipients' inboxes during the
	// predictive inbox placement test.
	InboxPercentage *float64

	// The percentage of emails that didn't arrive in recipients' inboxes at all during
	// the predictive inbox placement test.
	MissingPercentage *float64

	// The percentage of emails that arrived in recipients' spam or junk mail folders
	// during the predictive inbox placement test.
	SpamPercentage *float64

	// The percentage of emails that were authenticated by using Sender Policy
	// Framework (SPF) during the predictive inbox placement test.
	SpfPercentage *float64

	noSmithyDocumentSerde
}

// The raw email message.
type RawMessage struct {

	// The raw email message. The message has to meet the following criteria:
	//
	// * The
	// message has to contain a header and a body, separated by one blank line.
	//
	// * All
	// of the required header fields must be present in the message.
	//
	// * Each part of a
	// multipart MIME message must be formatted properly.
	//
	// * Attachments must be in a
	// file format that Amazon Pinpoint supports.
	//
	// * The entire message must be Base64
	// encoded.
	//
	// * If any of the MIME parts in your message contain content that is
	// outside of the 7-bit ASCII character range, you should encode that content to
	// ensure that recipients' email clients render the message properly.
	//
	// * The length
	// of any single line of text in the message can't exceed 1,000 characters. This
	// restriction is defined in RFC 5321 (https://tools.ietf.org/html/rfc5321).
	//
	// This member is required.
	Data []byte

	noSmithyDocumentSerde
}

// Enable or disable collection of reputation metrics for emails that you send
// using this configuration set in the current AWS Region.
type ReputationOptions struct {

	// The date and time (in Unix time) when the reputation metrics were last given a
	// fresh start. When your account is given a fresh start, your reputation metrics
	// are calculated starting from the date of the fresh start.
	LastFreshStart *time.Time

	// If true, tracking of reputation metrics is enabled for the configuration set. If
	// false, tracking of reputation metrics is disabled for the configuration set.
	ReputationMetricsEnabled bool

	noSmithyDocumentSerde
}

// Used to enable or disable email sending for messages that use this configuration
// set in the current AWS Region.
type SendingOptions struct {

	// If true, email sending is enabled for the configuration set. If false, email
	// sending is disabled for the configuration set.
	SendingEnabled bool

	noSmithyDocumentSerde
}

// An object that contains information about the per-day and per-second sending
// limits for your Amazon Pinpoint account in the current AWS Region.
type SendQuota struct {

	// The maximum number of emails that you can send in the current AWS Region over a
	// 24-hour period. This value is also called your sending quota.
	Max24HourSend float64

	// The maximum number of emails that you can send per second in the current AWS
	// Region. This value is also called your maximum sending rate or your maximum TPS
	// (transactions per second) rate.
	MaxSendRate float64

	// The number of emails sent from your Amazon Pinpoint account in the current AWS
	// Region over the past 24 hours.
	SentLast24Hours float64

	noSmithyDocumentSerde
}

// An object that defines an Amazon SNS destination for email events. You can use
// Amazon SNS to send notification when certain email events occur.
type SnsDestination struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish
	// email events to. For more information about Amazon SNS topics, see the Amazon
	// SNS Developer Guide
	// (https://docs.aws.amazon.com/sns/latest/dg/CreateTopic.html).
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

// An object that defines the tags that are associated with a resource. A tag is a
// label that you optionally define and associate with a resource in Amazon
// Pinpoint. Tags can help you categorize and manage resources in different ways,
// such as by purpose, owner, environment, or other criteria. A resource can have
// as many as 50 tags. Each tag consists of a required tag key and an associated
// tag value, both of which you define. A tag key is a general label that acts as a
// category for a more specific tag value. A tag value acts as a descriptor within
// a tag key. A tag key can contain as many as 128 characters. A tag value can
// contain as many as 256 characters. The characters can be Unicode letters,
// digits, white space, or one of the following symbols: _ . : / = + -. The
// following additional restrictions apply to tags:
//
// * Tag keys and values are case
// sensitive.
//
// * For each associated resource, each tag key must be unique and it
// can have only one value.
//
// * The aws: prefix is reserved for use by AWS; you
// can’t use it in any tag keys or values that you define. In addition, you can't
// edit or remove tag keys or values that use this prefix. Tags that use this
// prefix don’t count against the limit of 50 tags per resource.
//
// * You can
// associate tags with public or shared resources, but the tags are available only
// for your AWS account, not any other accounts that share the resource. In
// addition, the tags are available only for resources that are located in the
// specified AWS Region for your AWS account.
type Tag struct {

	// One part of a key-value pair that defines a tag. The maximum length of a tag key
	// is 128 characters. The minimum length is 1 character.
	//
	// This member is required.
	Key *string

	// The optional part of a key-value pair that defines a tag. The maximum length of
	// a tag value is 256 characters. The minimum length is 0 characters. If you don’t
	// want a resource to have a specific tag value, don’t specify a value for this
	// parameter. Amazon Pinpoint will set the value to an empty string.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type Template struct {

	// The Amazon Resource Name (ARN) of the template.
	TemplateArn *string

	// An object that defines the values to use for message variables in the template.
	// This object is a set of key-value pairs. Each key defines a message variable in
	// the template. The corresponding value defines the value to use for that
	// variable.
	TemplateData *string

	noSmithyDocumentSerde
}

// An object that defines the tracking options for a configuration set. When you
// use Amazon Pinpoint to send an email, it contains an invisible image that's used
// to track when recipients open your email. If your email contains links, those
// links are changed slightly in order to track when recipients click them. These
// images and links include references to a domain operated by AWS. You can
// optionally configure Amazon Pinpoint to use a domain that you operate for these
// images and links.
type TrackingOptions struct {

	// The domain that you want to use for tracking open and click events.
	//
	// This member is required.
	CustomRedirectDomain *string

	noSmithyDocumentSerde
}

// An object that contains information about the amount of email that was delivered
// to recipients.
type VolumeStatistics struct {

	// The total number of emails that arrived in recipients' inboxes.
	InboxRawCount *int64

	// An estimate of the percentage of emails sent from the current domain that will
	// arrive in recipients' inboxes.
	ProjectedInbox *int64

	// An estimate of the percentage of emails sent from the current domain that will
	// arrive in recipients' spam or junk mail folders.
	ProjectedSpam *int64

	// The total number of emails that arrived in recipients' spam or junk mail
	// folders.
	SpamRawCount *int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
