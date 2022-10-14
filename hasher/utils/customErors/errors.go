package customErors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	statusBadRequest = 400
)

var (
	//TODO come up with status codes
	ErrNilRequest  = status.Error(codes.Code(0), "nil request")
	ErrInvalidData = status.Error(codes.Code(statusBadRequest), "invalid input data")
)
