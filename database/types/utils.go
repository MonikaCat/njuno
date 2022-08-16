package types

import (
	"database/sql"
	"strings"
)

func ToString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

// _________________________________________________________

func ToNullString(value string) sql.NullString {
	value = strings.TrimSpace(value)
	return sql.NullString{
		Valid:  value != "",
		String: value,
	}
}
