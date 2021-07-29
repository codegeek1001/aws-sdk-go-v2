// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You are not authorized to use this operation with the given parameters.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've exceeded the notification or subscriber limit.
type CreationLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CreationLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreationLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreationLimitExceededException) ErrorCode() string             { return "CreationLimitExceededException" }
func (e *CreationLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The budget name already exists. Budget names must be unique within an account.
type DuplicateRecordException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DuplicateRecordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateRecordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateRecordException) ErrorCode() string             { return "DuplicateRecordException" }
func (e *DuplicateRecordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pagination token expired.
type ExpiredNextTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ExpiredNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredNextTokenException) ErrorCode() string             { return "ExpiredNextTokenException" }
func (e *ExpiredNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error on the server occurred during the processing of your request. Try again
// later.
type InternalErrorException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string             { return "InternalErrorException" }
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The pagination token is invalid.
type InvalidNextTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error on the client occurred. Typically, the cause is an invalid input value.
type InvalidParameterException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can’t locate the resource that you specified.
type NotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was received and recognized by the server, but the server rejected
// that particular method for the requested resource.
type ResourceLockedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceLockedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLockedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLockedException) ErrorCode() string             { return "ResourceLockedException" }
func (e *ResourceLockedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
