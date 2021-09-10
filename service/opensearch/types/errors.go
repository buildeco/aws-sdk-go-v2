// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An error occurred because user does not have permissions to access the resource.
// Returns HTTP status code 403.
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

// An error occurred while processing the request.
type BaseException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BaseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BaseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BaseException) ErrorCode() string             { return "BaseException" }
func (e *BaseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred because the client attempts to remove a resource that is
// currently in use. Returns HTTP status code 409.
type ConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occured because the client wanted to access a not supported operation.
// Gives http status code of 409.
type DisabledOperationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DisabledOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DisabledOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DisabledOperationException) ErrorCode() string             { return "DisabledOperationException" }
func (e *DisabledOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request processing has failed because of an unknown error, exception or
// failure (the failure is internal to the service) . Gives http status code of
// 500.
type InternalException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string             { return "InternalException" }
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request processing has failed because of invalid pagination token provided
// by customer. Returns an HTTP status code of 400.
type InvalidPaginationTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidPaginationTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPaginationTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPaginationTokenException) ErrorCode() string {
	return "InvalidPaginationTokenException"
}
func (e *InvalidPaginationTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception for trying to create or access sub-resource that is either invalid
// or not supported. Gives http status code of 409.
type InvalidTypeException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTypeException) ErrorCode() string             { return "InvalidTypeException" }
func (e *InvalidTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception for trying to create more than allowed resources or sub-resources.
// Gives http status code of 409.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception for creating a resource that already exists. Gives http status code
// of 400.
type ResourceAlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception for accessing or deleting a resource that does not exist. Gives
// http status code of 400.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception for missing / invalid input fields. Gives http status code of 400.
type ValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
