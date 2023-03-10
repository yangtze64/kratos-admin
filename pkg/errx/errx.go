package errx

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type ErrxReason string

const (
	UnknownError                ErrxReason = "UNKNOWN_ERROR"
	Success                     ErrxReason = "SUCCESS"
	ServerError                 ErrxReason = "SERVER_ERROR"
	TowPasswordDiff             ErrxReason = "TOW_PASSWORD_DIFF"
	UserNameExist               ErrxReason = "USERNAME_EXIST"
	UserMobileExist             ErrxReason = "USER_MOBILE_EXIST"
	UserEmailExist              ErrxReason = "USER_EMAIL_EXIST"
	UserNotFound                ErrxReason = "USER_NOT_FOUND"
	UsernameOrPasswordIncorrect ErrxReason = "USERNAME_OR_PASSWORD_INCORRECT"
	UnauthorizedInfoMissing     ErrxReason = "UNAUTHORIZED_INFO_MISSING"
	UnauthorizedTokenInvalid    ErrxReason = "UNAUTHORIZED_TOKEN_INVALID"
)

var reasonMessage = map[ErrxReason]string{
	UnknownError:                "Unknown error",
	Success:                     "Success",
	ServerError:                 "Server error",
	TowPasswordDiff:             "The two passwords are different",
	UserNameExist:               "The username already exists",
	UserMobileExist:             "The user mobile already exists",
	UserEmailExist:              "The user email already exists",
	UserNotFound:                "The user was not found",
	UsernameOrPasswordIncorrect: "The user name or password is incorrect",
	UnauthorizedInfoMissing:     "The authorization information not found",
	UnauthorizedTokenInvalid:    "The authorization token is invalid",
}

var reasonCode = map[ErrxReason]int{
	UnknownError:                0,
	Success:                     http.StatusOK,
	ServerError:                 http.StatusInternalServerError,
	TowPasswordDiff:             http.StatusBadRequest,
	UserNameExist:               http.StatusConflict,
	UserMobileExist:             http.StatusConflict,
	UserEmailExist:              http.StatusConflict,
	UserNotFound:                http.StatusNotFound,
	UsernameOrPasswordIncorrect: http.StatusUnauthorized,
	UnauthorizedInfoMissing:     http.StatusUnauthorized,
	UnauthorizedTokenInvalid:    http.StatusUnauthorized,
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
