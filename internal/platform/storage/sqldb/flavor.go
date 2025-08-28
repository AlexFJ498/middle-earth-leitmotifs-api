package sqldb

import (
	"errors"

	"github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
)

// defaultFlavor is the default SQL flavor used in the application.
var defaultFlavor = sqlbuilder.PostgreSQL

// Infra-level errors
var (
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrUniqueViolation     = errors.New("unique violation")
	ErrUnknown             = errors.New("unknown database error")
)

func mapSQLError(code string) error {
	switch defaultFlavor {
	case sqlbuilder.PostgreSQL:
		switch code {
		case "23503":
			return ErrForeignKeyViolation
		case "23505":
			return ErrUniqueViolation
		}
	}
	return ErrUnknown
}

func extractSQLErrorCode(err error) string {
	switch defaultFlavor {
	case sqlbuilder.PostgreSQL:
		if err, ok := err.(*pq.Error); ok {
			return string(err.Code)
		}
	}
	return ""
}
