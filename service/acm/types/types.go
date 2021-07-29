// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains metadata about an ACM certificate. This structure is returned in the
// response to a DescribeCertificate request.
type CertificateDetail struct {

	// The Amazon Resource Name (ARN) of the certificate. For more information about
	// ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	CertificateArn *string

	// The Amazon Resource Name (ARN) of the ACM PCA private certificate authority (CA)
	// that issued the certificate. This has the following format:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	CertificateAuthorityArn *string

	// The time at which the certificate was requested.
	CreatedAt *time.Time

	// The fully qualified domain name for the certificate, such as www.example.com or
	// example.com.
	DomainName *string

	// Contains information about the initial validation of each domain name that
	// occurs as a result of the RequestCertificate request. This field exists only
	// when the certificate type is AMAZON_ISSUED.
	DomainValidationOptions []DomainValidation

	// Contains a list of Extended Key Usage X.509 v3 extension objects. Each object
	// specifies a purpose for which the certificate public key can be used and
	// consists of a name and an object identifier (OID).
	ExtendedKeyUsages []ExtendedKeyUsage

	// The reason the certificate request failed. This value exists only when the
	// certificate status is FAILED. For more information, see Certificate Request
	// Failed
	// (https://docs.aws.amazon.com/acm/latest/userguide/troubleshooting.html#troubleshooting-failed)
	// in the Amazon Web Services Certificate Manager User Guide.
	FailureReason FailureReason

	// The date and time at which the certificate was imported. This value exists only
	// when the certificate type is IMPORTED.
	ImportedAt *time.Time

	// A list of ARNs for the Amazon Web Services resources that are using the
	// certificate. A certificate can be used by multiple Amazon Web Services
	// resources.
	InUseBy []string

	// The time at which the certificate was issued. This value exists only when the
	// certificate type is AMAZON_ISSUED.
	IssuedAt *time.Time

	// The name of the certificate authority that issued and signed the certificate.
	Issuer *string

	// The algorithm that was used to generate the public-private key pair.
	KeyAlgorithm KeyAlgorithm

	// A list of Key Usage X.509 v3 extension objects. Each object is a string value
	// that identifies the purpose of the public key contained in the certificate.
	// Possible extension values include DIGITAL_SIGNATURE, KEY_ENCHIPHERMENT,
	// NON_REPUDIATION, and more.
	KeyUsages []KeyUsage

	// The time after which the certificate is not valid.
	NotAfter *time.Time

	// The time before which the certificate is not valid.
	NotBefore *time.Time

	// Value that specifies whether to add the certificate to a transparency log.
	// Certificate transparency makes it possible to detect SSL certificates that have
	// been mistakenly or maliciously issued. A browser might respond to certificate
	// that has not been logged by showing an error message. The logs are
	// cryptographically secure.
	Options *CertificateOptions

	// Specifies whether the certificate is eligible for renewal. At this time, only
	// exported private certificates can be renewed with the RenewCertificate command.
	RenewalEligibility RenewalEligibility

	// Contains information about the status of ACM's managed renewal
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the
	// certificate. This field exists only when the certificate type is AMAZON_ISSUED.
	RenewalSummary *RenewalSummary

	// The reason the certificate was revoked. This value exists only when the
	// certificate status is REVOKED.
	RevocationReason RevocationReason

	// The time at which the certificate was revoked. This value exists only when the
	// certificate status is REVOKED.
	RevokedAt *time.Time

	// The serial number of the certificate.
	Serial *string

	// The algorithm that was used to sign the certificate.
	SignatureAlgorithm *string

	// The status of the certificate.
	Status CertificateStatus

	// The name of the entity that is associated with the public key contained in the
	// certificate.
	Subject *string

	// One or more domain names (subject alternative names) included in the
	// certificate. This list contains the domain names that are bound to the public
	// key that is contained in the certificate. The subject alternative names include
	// the canonical domain name (CN) of the certificate and additional domain names
	// that can be used to connect to the website.
	SubjectAlternativeNames []string

	// The source of the certificate. For certificates provided by ACM, this value is
	// AMAZON_ISSUED. For certificates that you imported with ImportCertificate, this
	// value is IMPORTED. ACM does not provide managed renewal
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for imported
	// certificates. For more information about the differences between certificates
	// that you import and those that ACM provides, see Importing Certificates
	// (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in
	// the Amazon Web Services Certificate Manager User Guide.
	Type CertificateType

	noSmithyDocumentSerde
}

// Structure that contains options for your certificate. Currently, you can use
// this only to specify whether to opt in to or out of certificate transparency
// logging. Some browsers require that public certificates issued for your domain
// be recorded in a log. Certificates that are not logged typically generate a
// browser error. Transparency makes it possible for you to detect SSL/TLS
// certificates that have been mistakenly or maliciously issued for your domain.
// For general information, see Certificate Transparency Logging
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency).
type CertificateOptions struct {

	// You can opt out of certificate transparency logging by specifying the DISABLED
	// option. Opt in by specifying ENABLED.
	CertificateTransparencyLoggingPreference CertificateTransparencyLoggingPreference

	noSmithyDocumentSerde
}

// This structure is returned in the response object of ListCertificates action.
type CertificateSummary struct {

	// Amazon Resource Name (ARN) of the certificate. This is of the form:
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	// For more information about ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string

	// Fully qualified domain name (FQDN), such as www.example.com or example.com, for
	// the certificate.
	DomainName *string

	noSmithyDocumentSerde
}

// Contains information about the validation of each domain name in the
// certificate.
type DomainValidation struct {

	// A fully qualified domain name (FQDN) in the certificate. For example,
	// www.example.com or example.com.
	//
	// This member is required.
	DomainName *string

	// Contains the CNAME record that you add to your DNS database for domain
	// validation. For more information, see Use DNS to Validate Domain Ownership
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html).
	// Note: The CNAME information that you need does not include the name of your
	// domain. If you include  your domain name in the DNS database CNAME record,
	// validation fails.  For example, if the name is
	// "_a79865eb4cd1a6ab990a45779b4e0b96.yourdomain.com", only
	// "_a79865eb4cd1a6ab990a45779b4e0b96" must be used.
	ResourceRecord *ResourceRecord

	// The domain name that ACM used to send domain validation emails.
	ValidationDomain *string

	// A list of email addresses that ACM used to send domain validation emails.
	ValidationEmails []string

	// Specifies the domain validation method.
	ValidationMethod ValidationMethod

	// The validation status of the domain name. This can be one of the following
	// values:
	//
	// * PENDING_VALIDATION
	//
	// * SUCCESS
	//
	// * FAILED
	ValidationStatus DomainStatus

	noSmithyDocumentSerde
}

// Contains information about the domain names that you want ACM to use to send you
// emails that enable you to validate domain ownership.
type DomainValidationOption struct {

	// A fully qualified domain name (FQDN) in the certificate request.
	//
	// This member is required.
	DomainName *string

	// The domain name that you want ACM to use to send you validation emails. This
	// domain name is the suffix of the email addresses that you want ACM to use. This
	// must be the same as the DomainName value or a superdomain of the DomainName
	// value. For example, if you request a certificate for testing.example.com, you
	// can specify example.com for this value. In that case, ACM sends domain
	// validation emails to the following five addresses:
	//
	// * admin@example.com
	//
	// *
	// administrator@example.com
	//
	// * hostmaster@example.com
	//
	// * postmaster@example.com
	//
	// *
	// webmaster@example.com
	//
	// This member is required.
	ValidationDomain *string

	noSmithyDocumentSerde
}

// Object containing expiration events options associated with an Amazon Web
// Services account.
type ExpiryEventsConfiguration struct {

	// Specifies the number of days prior to certificate expiration when ACM starts
	// generating EventBridge events. ACM sends one event per day per certificate until
	// the certificate expires. By default, accounts receive events starting 45 days
	// before certificate expiration.
	DaysBeforeExpiry *int32

	noSmithyDocumentSerde
}

// The Extended Key Usage X.509 v3 extension defines one or more purposes for which
// the public key can be used. This is in addition to or in place of the basic
// purposes specified by the Key Usage extension.
type ExtendedKeyUsage struct {

	// The name of an Extended Key Usage value.
	Name ExtendedKeyUsageName

	// An object identifier (OID) for the extension value. OIDs are strings of numbers
	// separated by periods. The following OIDs are defined in RFC 3280 and RFC
	// 5280.
	//
	// * 1.3.6.1.5.5.7.3.1 (TLS_WEB_SERVER_AUTHENTICATION)
	//
	// * 1.3.6.1.5.5.7.3.2
	// (TLS_WEB_CLIENT_AUTHENTICATION)
	//
	// * 1.3.6.1.5.5.7.3.3 (CODE_SIGNING)
	//
	// *
	// 1.3.6.1.5.5.7.3.4 (EMAIL_PROTECTION)
	//
	// * 1.3.6.1.5.5.7.3.8 (TIME_STAMPING)
	//
	// *
	// 1.3.6.1.5.5.7.3.9 (OCSP_SIGNING)
	//
	// * 1.3.6.1.5.5.7.3.5 (IPSEC_END_SYSTEM)
	//
	// *
	// 1.3.6.1.5.5.7.3.6 (IPSEC_TUNNEL)
	//
	// * 1.3.6.1.5.5.7.3.7 (IPSEC_USER)
	OID *string

	noSmithyDocumentSerde
}

// This structure can be used in the ListCertificates action to filter the output
// of the certificate list.
type Filters struct {

	// Specify one or more ExtendedKeyUsage extension values.
	ExtendedKeyUsage []ExtendedKeyUsageName

	// Specify one or more algorithms that can be used to generate key pairs. Default
	// filtering returns only RSA_1024 and RSA_2048 certificates that have at least one
	// domain. To return other certificate types, provide the desired type signatures
	// in a comma-separated list. For example, "keyTypes": ["RSA_2048,RSA_4096"]
	// returns both RSA_2048 and RSA_4096 certificates.
	KeyTypes []KeyAlgorithm

	// Specify one or more KeyUsage extension values.
	KeyUsage []KeyUsageName

	noSmithyDocumentSerde
}

// The Key Usage X.509 v3 extension defines the purpose of the public key contained
// in the certificate.
type KeyUsage struct {

	// A string value that contains a Key Usage extension name.
	Name KeyUsageName

	noSmithyDocumentSerde
}

// Contains information about the status of ACM's managed renewal
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) for the
// certificate. This structure exists only when the certificate type is
// AMAZON_ISSUED.
type RenewalSummary struct {

	// Contains information about the validation of each domain name in the
	// certificate, as it pertains to ACM's managed renewal
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html). This is
	// different from the initial validation that occurs as a result of the
	// RequestCertificate request. This field exists only when the certificate type is
	// AMAZON_ISSUED.
	//
	// This member is required.
	DomainValidationOptions []DomainValidation

	// The status of ACM's managed renewal
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) of the
	// certificate.
	//
	// This member is required.
	RenewalStatus RenewalStatus

	// The time at which the renewal summary was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The reason that a renewal request was unsuccessful.
	RenewalStatusReason FailureReason

	noSmithyDocumentSerde
}

// Contains a DNS record value that you can use to validate ownership or control of
// a domain. This is used by the DescribeCertificate action.
type ResourceRecord struct {

	// The name of the DNS record to create in your domain. This is supplied by ACM.
	//
	// This member is required.
	Name *string

	// The type of DNS record. Currently this can be CNAME.
	//
	// This member is required.
	Type RecordType

	// The value of the CNAME record to add to your DNS database. This is supplied by
	// ACM.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// A key-value pair that identifies or specifies metadata about an ACM resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
