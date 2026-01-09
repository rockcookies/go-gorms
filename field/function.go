package field

import (
	"strings"

	"gorm.io/gorm/clause"
)

// Func sql functions
var Func = new(function)

type function struct{}

// UnixTimestamp same as UNIX_TIMESTAMP([date])
func (f *function) UnixTimestamp(date ...string) Field {
	if len(date) > 0 {
		return Field{expr: expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []interface{}{date[0]}}}}
	}
	return Field{expr: expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP()"}}}
}

// FromUnixTime FROM_UNIXTIME(unix_timestamp[,format])
func (f *function) FromUnixTime(date uint64, format string) Field {
	if strings.TrimSpace(format) != "" {
		return Field{expr: expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?, ?)", Vars: []interface{}{date, format}}}}
	}
	return Field{expr: expr{e: clause.Expr{SQL: "FROM_UNIXTIME(?)", Vars: []interface{}{date}}}}
}

func (f *function) Rand() Field {
	return Field{expr: expr{e: clause.Expr{SQL: "RAND()"}}}
}

func (f *function) Random() Field {
	return Field{expr: expr{e: clause.Expr{SQL: "RANDOM()"}}}
}
