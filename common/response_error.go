package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gertd/go-pluralize"
)

var plrl = pluralize.NewClient()

var ErrRecordNotFound = errors.New("record not found")

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomErrorResponse(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong", err.Error(), "ErrDB")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternalServer(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong", err.Error(), "ErrInternalServer")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot list %s", plrl.Plural(strings.ToLower(entity))),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("%s has been deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("%s has been existed", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExisted", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomErrorResponse(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrPermissionDenied(err error) *AppError {
	return NewCustomErrorResponse(err, "permission denied", "ErrPermissionDenied")
}

func ErrNoPermission(err error) *AppError {
	return NewCustomErrorResponse(err, "no permission", "ErrNoPermission")
}
