package errors

import (
	"errors"
	"net/http"
)

var (
	ErrDuplicate             = errors.New("duplicate")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrUserNotFound          = errors.New("user not found")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrIdentityAlreadyExists = errors.New("identity already exists")
	ErrInternalServerError   = errors.New("internal server error")
	ErrCurrencyNotFound      = errors.New("currency not found")
	ErrInvalidAmount         = errors.New("invalid amount")
	ErrWalletNotFound        = errors.New("wallet not found")
	ErrInsufficientBalance   = errors.New("insufficient balance")
)

func ErrorResPonse(err error) (int, error) {
	switch err {
	case ErrInvalidPassword,
		ErrInvalidAmount,
		ErrInsufficientBalance:
		return http.StatusBadRequest, err
	case ErrUserNotFound,
		ErrCurrencyNotFound,
		ErrWalletNotFound:
		return http.StatusNotFound, err
	case ErrIdentityAlreadyExists:
		return http.StatusConflict, err
	default:
		return http.StatusInternalServerError, ErrInternalServerError
	}
}
