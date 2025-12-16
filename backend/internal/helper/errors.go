package helper

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// IsUniqueViolation returns true if the unique constraint is violated.
func IsUniqueViolation(err error) bool {
	const uniqueViolationCode = "23505"
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == uniqueViolationCode
	}
	return false
}
