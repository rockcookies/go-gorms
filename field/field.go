package field

import (
	"database/sql/driver"

	"gorm.io/gorm/clause"
)

// ScanValuer interface for Field
type ScanValuer interface {
	Scan(src interface{}) error   // sql.Scanner
	Value() (driver.Value, error) // driver.Valuer
}

// Field a standard field struct
type Field struct{ expr }

// Eq judge equal
func (field Field) Eq(value interface{}) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq judge not equal
func (field Field) Neq(value interface{}) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// In ...
func (field Field) In(values ...interface{}) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: field.toSlice(values...)}}
}

// NotIn ...
func (field Field) NotIn(values ...interface{}) Expr {
	return expr{e: clause.Not(field.In(values...).expression())}
}

// Gt ...
func (field Field) Gt(value interface{}) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte ...
func (field Field) Gte(value interface{}) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt ...
func (field Field) Lt(value interface{}) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte ...
func (field Field) Lte(value interface{}) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Like ...
func (field Field) Like(value interface{}) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// Value ...
func (field Field) Value(value interface{}) AssignExpr {
	return field.value(value)
}

// Sum ...
func (field Field) Sum() Field {
	return Field{field.sum()}
}

// IfNull ...
func (field Field) IfNull(value interface{}) Expr {
	return field.ifNull(value)
}

// Field ...
func (field Field) Field(value []interface{}) Expr {
	return field.field(value)
}

// Between ...
func (field Field) Between(values []interface{}) Expr {
	return field.between(values)
}

// Add ...
func (field Field) Add(value interface{}) Field {
	return Field{field.add(value)}
}

// Sub ...
func (field Field) Sub(value interface{}) Field {
	return Field{field.sub(value)}
}

// Mul ...
func (field Field) Mul(value interface{}) Field {
	return Field{field.mul(value)}
}

// Div ...
func (field Field) Div(value interface{}) Field {
	return Field{field.div(value)}
}

// Mod ...
func (field Field) Mod(value interface{}) Field {
	return Field{field.mod(value)}
}

func (field Field) toSlice(values ...interface{}) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
