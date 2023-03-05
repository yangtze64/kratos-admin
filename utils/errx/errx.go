package errx

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type ErrxReason string

const (
	UnknownError     ErrxReason = "UNKNOWN_ERROR"
	Success          ErrxReason = "SUCCESS"
	ServerError      ErrxReason = "SERVER_ERROR"
	TowPasswordDiff  ErrxReason = "TOW_PASSWORD_DIFF"
	UserNameExist    ErrxReason = "USERNAME_EXIST"
	UserMobileExist  ErrxReason = "USER_MOBILE_EXIST"
	UserEmailExist   ErrxReason = "USER_EMAIL_EXIST"
	UserNotFound     ErrxReason = "USER_NOT_FOUND"
)

var reasonMessage = map[ErrxReason]string{
	UnknownError:     "unknown error",
	Success:          "success",
	ServerError:      "server error",
	TowPasswordDiff:  "the two passwords are different",
	UserNameExist:    "the username already exists",
	UserMobileExist:  "the user mobile already exists",
	UserEmailExist:   "the user email already exists",
	UserNotFound:     "This user was not found",
}

var reasonCode = map[ErrxReason]int{
	UnknownError:     0,
	Success:          http.StatusOK,
	ServerError:      http.StatusInternalServerError,
	TowPasswordDiff:  http.StatusBadRequest,
	UserNameExist:    http.StatusConflict,
	UserMobileExist:  http.StatusConflict,
	UserEmailExist:   http.StatusConflict,
	UserNotFound:     http.StatusNotFound,
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
