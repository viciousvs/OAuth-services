package customErrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	statusLogicError          = 409
	statusBadRequest          = 400
	statusUnprocessableEntity = 422
	statusInternalError       = 500
	statusUnauthorized        = 401
)

var (
	//TODO come up with status codes
	ErrNilRequest        = status.Error(codes.Code(statusUnprocessableEntity), "nil request")
	ErrInvalidUUID       = status.Error(codes.Code(statusBadRequest), "invalid user uuid")
	ErrInvalidData       = status.Error(codes.Code(statusBadRequest), "invalid user data")
	ErrNilHashResponse   = status.Error(codes.Code(statusLogicError), "nil hash response")
	ErrNilUUIDResponse   = status.Error(codes.Code(statusLogicError), "nil uuid response")
	ErrNilTokensResponse = status.Error(codes.Code(statusLogicError), "nil tokens response")
	ErrNilUserResponse   = status.Error(codes.Code(statusLogicError), "nil user response")
)

func ErrorWithCode400(msg string) error {
	return status.Error(codes.Code(statusBadRequest), msg)
}
func ErrorWithCode500(msg string) error {
	return status.Error(codes.Code(statusInternalError), msg)
}
