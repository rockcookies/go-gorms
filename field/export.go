package field

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Option field option
type Option func(clause.Column) clause.Column

var (
	banColumnRaw Option = func(col clause.Column) clause.Column {
		col.Raw = false
		return col
	}
)

// ======================== generic field =======================

// NewField create new field
func NewField(table, column string, opts ...Option) Field {
	return Field{expr: expr{col: toColumn(table, column, opts...)}}
}

// NewUnsafeFieldRaw create new field by native sql
//
// Warning: Using NewUnsafeFieldRaw with raw SQL exposes your application to SQL injection vulnerabilities.
// Always validate/sanitize inputs and prefer parameterized queries or NewField methods for field construction.
// Use this low-level function only when absolutely necessary, and ensure any embedded values are properly escaped.
func NewUnsafeFieldRaw(rawSQL string, vars ...interface{}) Field {
	return Field{expr: expr{e: clause.Expr{SQL: rawSQL, Vars: vars}}}
}

func toColumn(table, column string, opts ...Option) clause.Column {
	col := clause.Column{Table: table, Name: column}
	for _, opt := range opts {
		col = opt(col)
	}
	return banColumnRaw(col)
}

// ======================== boolean operate ========================

// Or return or condition
func Or(exprs ...Expr) Expr {
	return &expr{e: clause.Or(toExpression(exprs...)...)}
}

// And return and condition
func And(exprs ...Expr) Expr {
	return &expr{e: clause.And(toExpression(exprs...)...)}
}

// Not return not condition
func Not(exprs ...Expr) Expr {
	return &expr{e: clause.Not(toExpression(exprs...)...)}
}

func toExpression(conds ...Expr) []clause.Expression {
	exprs := make([]clause.Expression, len(conds))
	for i, cond := range conds {
		exprs[i] = cond.expression()
	}
	return exprs
}

// ======================== subquery method ========================

// ContainsSubQuery return contains subquery
// when len(columns) == 1, equal to columns[0] IN (subquery)
// when len(columns) > 1, equal to (columns[0], columns[1], ...) IN (subquery)
func ContainsSubQuery(columns []Expr, subQuery *gorm.DB) Expr {
	switch len(columns) {
	case 0:
		return expr{e: clause.Expr{}}
	case 1:
		return expr{e: clause.Expr{
			SQL:  "? IN (?)",
			Vars: []interface{}{columns[0].RawExpr(), subQuery},
		}}
	default: // len(columns) > 0
		placeholders := make([]string, len(columns))
		cols := make([]interface{}, len(columns))
		for i, c := range columns {
			placeholders[i], cols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) IN (?)", strings.Join(placeholders, ",")),
			Vars: append(cols, subQuery),
		}}
	}
}

// AssignSubQuery assign with subquery
func AssignSubQuery(columns []Expr, subQuery *gorm.DB) AssignExpr {
	cols := make([]string, len(columns))
	for i, c := range columns {
		cols[i] = string(c.BuildColumn(subQuery.Statement))
	}

	name := cols[0]
	if len(cols) > 1 {
		name = "(" + strings.Join(cols, ",") + ")"
	}

	return expr{e: clause.Set{{
		Column: clause.Column{Name: name, Raw: true},
		Value:  gorm.Expr("(?)", subQuery),
	}}}
}

// CompareOperator compare operator
type CompareOperator string

const (
	// EqOp =
	EqOp CompareOperator = " = "
	// NeqOp <>
	NeqOp CompareOperator = " <> "
	// GtOp >
	GtOp CompareOperator = " > "
	// GteOp >=
	GteOp CompareOperator = " >= "
	// LtOp <
	LtOp CompareOperator = " < "
	// LteOp <=
	LteOp CompareOperator = " <= "
	// ExistsOp EXISTS
	ExistsOp CompareOperator = "EXISTS "
)

// CompareSubQuery compare with sub query
func CompareSubQuery(op CompareOperator, column Expr, subQuery *gorm.DB) Expr {
	if op == ExistsOp {
		return expr{e: clause.Expr{
			SQL:  fmt.Sprint(op, "(?)"),
			Vars: []interface{}{subQuery},
		}}
	}
	return expr{e: clause.Expr{
		SQL:  fmt.Sprint("?", op, "(?)"),
		Vars: []interface{}{column.RawExpr(), subQuery},
	}}
}

// Value ...
type Value interface {
	expr() clause.Expr

	// implement Condition
	BeCond() interface{}
	CondError() error
}

type val clause.Expr

func (v val) expr() clause.Expr   { return clause.Expr(v) }
func (v val) BeCond() interface{} { return v }
func (val) CondError() error      { return nil }

// Values convert value to expression which implement Value
func Values(value interface{}) Value {
	return val(clause.Expr{
		SQL:                "?",
		Vars:               []interface{}{value},
		WithoutParentheses: true,
	})
}

// ContainsValue return expression which compare with value
func ContainsValue(columns []Expr, value Value) Expr {
	switch len(columns) {
	case 0:
		return expr{e: clause.Expr{}}
	case 1:
		return expr{e: clause.Expr{
			SQL:  "? IN (?)",
			Vars: []interface{}{columns[0].RawExpr(), value.expr()},
		}}
	default: // len(columns) > 0
		vars := make([]string, len(columns))
		queryCols := make([]interface{}, len(columns))
		for i, c := range columns {
			vars[i], queryCols[i] = "?", c.RawExpr()
		}
		return expr{e: clause.Expr{
			SQL:  fmt.Sprintf("(%s) IN (?)", strings.Join(vars, ", ")),
			Vars: append(queryCols, value.expr()),
		}}
	}
}

// EmptyExpr return a empty expression
func EmptyExpr() Expr { return expr{e: clause.Expr{}} }
