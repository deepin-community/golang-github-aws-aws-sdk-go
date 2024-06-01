// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costoptimizationhub

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You are not authorized to use this operation with the given parameters.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An error on the server occurred during the processing of your request. Try
	// again later.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified Amazon Resource Name (ARN) in the request doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an Amazon Web Services
	// service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
	"ValidationException":       newErrorValidationException,
}
