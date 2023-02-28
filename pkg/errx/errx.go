package errx

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type ErrxReason string

const (
	UnknownError    ErrxReason = "UNKNOWN_ERROR"
	Success         ErrxReason = "SUCCESS"
	ServerError     ErrxReason = "SERVER_ERROR"
	TowPasswordDiff ErrxReason = "TOW_PASSWORD_DIFF"
)

var reasonMessage = map[ErrxReason]string{
	UnknownError:    "unknown error",
	Success:         "success",
	ServerError:     "server error",
	TowPasswordDiff: "the two passwords are different",
}

var reasonCode = map[ErrxReason]int{
	UnknownError:    0,
	Success:         http.StatusOK,
	ServerError:     http.StatusInternalServerError,
	TowPasswordDiff: http.StatusBadRequest,
}

func New(reason ErrxReason) *errors.Error {
	var (
		msg  string
		code int
		ok   bool
	)
	if msg, ok = reasonMessage[reason]; !ok {
		reason = UnknownError
		msg = reasonMessage[reason]
	}
	if code, ok = reasonCode[reason]; !ok {
		code = reasonCode[UnknownError]
	}
	return errors.New(code, string(reason), msg)
}
