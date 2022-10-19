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
)

var (
	//TODO come up with status codes
	ErrNilRequest          = status.Error(codes.Code(statusUnprocessableEntity), "nil request")
	ErrNotFound            = status.Error(codes.Code(statusLogicError), "token not found")
	ErrInvalidUUID         = status.Error(codes.Code(statusBadRequest), "invalid user uuid")
	ErrInvalidAccessToken  = status.Error(codes.Code(statusBadRequest), "invalid access token")
	ErrInvalidRefreshToken = status.Error(codes.Code(statusBadRequest), "invalid refresh token")
)

func ErrorWithCode400(msg string) error {
	return status.Error(codes.Code(statusBadRequest), msg)
}
func ErrorWithCode500(msg string) error {
	return status.Error(codes.Code(statusInternalError), msg)
}
