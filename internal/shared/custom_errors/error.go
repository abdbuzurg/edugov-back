package custom_errors

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

type AppError struct {
	StatusCode int
	ErrType    string
	Err        error
}

func (ae *AppError) Error() string {
	return ae.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (ae *AppError) Log(requestID string) string {
	return fmt.Sprintf("[%s] %s: %v", requestID, ae.ErrType, ae.Err.Error())
}

func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func IsUniqueConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" 
	}

	return false
}

func InternalServerError(err error) error {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		ErrType:    "INTERNAL_SERVER_ERROR",
		Err:        err,
	}
}

func BadRequest(err error) error {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		ErrType:    "CLIENT_ERROR",
		Err:        err,
	}
}

func MethodNotAllowed(err error) error {
	return &AppError{
		StatusCode: http.StatusMethodNotAllowed,
		ErrType:    "CLIENT_ERROR",
		Err:        err,
	}
}

func NotFound(err error) error {
	return &AppError{
		StatusCode: http.StatusNotFound,
		ErrType:    "CLIENT_ERROR",
		Err:        err,
	}
}

func Unauthorized(err error) error {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		ErrType:    "CLIENT_ERROR",
		Err:        err,
	}
}
