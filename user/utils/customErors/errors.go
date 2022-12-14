package customErors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PgCodeRecordAlreadyExists      = "23505"
	PgCodeViolatesCheckConstraints = "23514"

	statusLogicError = 409
	statusBadRequest = 400
)

var (
	//TODO come up with status codes
	ErrNilRequest               = status.Error(codes.Code(0), "nil request")
	ErrNotFound                 = status.Error(codes.Code(statusLogicError), "not found")
	ErrRecordAlreadyExists      = status.Error(codes.Code(statusLogicError), "record already exists")
	ErrViolatesCheckConstraints = status.Error(codes.Code(statusLogicError), "violates check constraints")
	ErrRowsAffected             = status.Error(codes.Code(0), "count of affected rows not equal one")
	ErrInvalidData              = status.Error(codes.Code(statusBadRequest), "invalid input data")
	ErrIncorrectPasswordOrLogin = status.Error(codes.Code(statusLogicError), "incorrect password or login")
)
