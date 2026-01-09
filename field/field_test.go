package field

import (
	"reflect"
	"strings"
	"sync"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

var db, _ = gorm.Open(tests.DummyDialector{}, nil)

func GetStatement() *gorm.Statement {
	user, _ := schema.Parse(&User{}, &sync.Map{}, db.NamingStrategy)
	return &gorm.Statement{DB: db, Table: user.Table, Schema: user, Clauses: map[string]clause.Clause{}}
}

func CheckBuildExpr(t *testing.T, e Expr, result string, vars []interface{}) {
	stmt := GetStatement()

	e.expression().Build(stmt)

	sql := strings.TrimSpace(stmt.SQL.String())
	if sql != result {
		t.Errorf("SQL expects %v got %v", result, sql)
	}

	if !reflect.DeepEqual(stmt.Vars, vars) {
		t.Errorf("Vars expects %+v got %v", stmt.Vars, vars)
	}
}

func BuildToString(e Expr) (string, []interface{}) {
	stmt := GetStatement()

	e.expression().Build(stmt)

	return stmt.SQL.String(), stmt.Vars
}

type User struct {
	gorm.Model
	Name string
	Age  uint
}

// TestNewField tests basic field creation
func TestNewField(t *testing.T) {
	age := NewField("users", "age")
	expr := age.Eq(18)
	if expr == nil {
		t.Fatal("NewField expression returned nil")
	}
}

// TestEq tests equal comparison
func TestEq(t *testing.T) {
	age := NewField("users", "age")
	expr := age.Eq(18)
	CheckBuildExpr(t, expr, "`users`.`age` = ?", []interface{}{18})
}

// TestNeq tests not equal comparison
func TestNeq(t *testing.T) {
	age := NewField("users", "age")
	expr := age.Neq(18)
	CheckBuildExpr(t, expr, "`users`.`age` <> ?", []interface{}{18})
}

// TestGt tests greater than comparison
func TestGt(t *testing.T) {
	age := NewField("users", "age")
	expr := age.Gt(18)
	CheckBuildExpr(t, expr, "`users`.`age` > ?", []interface{}{18})
}

// TestLt tests less than comparison
func TestLt(t *testing.T) {
	price := NewField("products", "price")
	expr := price.Lt(100.0)
	CheckBuildExpr(t, expr, "`products`.`price` < ?", []interface{}{100.0})
}

// TestIn tests IN clause
func TestIn(t *testing.T) {
	status := NewField("orders", "status")
	expr := status.In(1, 2, 3)
	CheckBuildExpr(t, expr, "`orders`.`status` IN (?,?,?)", []interface{}{1, 2, 3})
}

// TestLike tests LIKE clause
func TestLike(t *testing.T) {
	name := NewField("users", "name")
	expr := name.Like("John%")
	CheckBuildExpr(t, expr, "`users`.`name` LIKE ?", []interface{}{"John%"})
}

// TestBetween tests BETWEEN clause
func TestBetween(t *testing.T) {
	age := NewField("users", "age")
	expr := age.Between([]interface{}{18, 65})
	CheckBuildExpr(t, expr, "`users`.`age` BETWEEN ? AND ?", []interface{}{18, 65})
}

// TestAddCol tests column addition
func TestAddCol(t *testing.T) {
	age := NewField("users", "age")
	bonus := NewField("users", "bonus")
	expr := age.AddCol(bonus)
	CheckBuildExpr(t, expr, "`users`.`age` + `users`.`bonus`", nil)
}

// TestSum tests SUM aggregation
func TestSum(t *testing.T) {
	price := NewField("products", "price")
	expr := price.Sum()
	CheckBuildExpr(t, expr, "SUM(`products`.`price`)", nil)
}

// TestCount tests COUNT aggregation
func TestCount(t *testing.T) {
	id := NewField("users", "id")
	expr := id.Count()
	CheckBuildExpr(t, expr, "COUNT(`users`.`id`)", nil)
}

// TestMax tests MAX aggregation
func TestMax(t *testing.T) {
	price := NewField("products", "price")
	expr := price.Max()
	CheckBuildExpr(t, expr, "MAX(`products`.`price`)", nil)
}

// TestMin tests MIN aggregation
func TestMin(t *testing.T) {
	price := NewField("products", "price")
	expr := price.Min()
	CheckBuildExpr(t, expr, "MIN(`products`.`price`)", nil)
}

// TestAvg tests AVG aggregation
func TestAvg(t *testing.T) {
	price := NewField("products", "price")
	expr := price.Avg()
	CheckBuildExpr(t, expr, "AVG(`products`.`price`)", nil)
}

// TestDistinct tests DISTINCT clause
func TestDistinct(t *testing.T) {
	category := NewField("products", "category")
	expr := category.Distinct()
	CheckBuildExpr(t, expr, "DISTINCT `products`.`category`", nil)
}

// TestIsNull tests IS NULL clause
func TestIsNull(t *testing.T) {
	name := NewField("users", "name")
	expr := name.IsNull()
	CheckBuildExpr(t, expr, "`users`.`name` IS NULL", nil)
}

// TestIsNotNull tests IS NOT NULL clause
func TestIsNotNull(t *testing.T) {
	name := NewField("users", "name")
	expr := name.IsNotNull()
	CheckBuildExpr(t, expr, "`users`.`name` IS NOT NULL", nil)
}

// TestIfNull tests IFNULL function
func TestIfNull(t *testing.T) {
	name := NewField("users", "name")
	expr := name.IfNull("Anonymous")
	CheckBuildExpr(t, expr, "IFNULL(`users`.`name`,?)", []interface{}{"Anonymous"})
}

// TestAsc tests ASC ordering
func TestAsc(t *testing.T) {
	name := NewField("users", "name")
	expr := name.Asc()
	CheckBuildExpr(t, expr, "`users`.`name` ASC", nil)
}

// TestDesc tests DESC ordering
func TestDesc(t *testing.T) {
	name := NewField("users", "name")
	expr := name.Desc()
	CheckBuildExpr(t, expr, "`users`.`name` DESC", nil)
}

// TestAs tests alias
func TestAs(t *testing.T) {
	name := NewField("users", "name")
	expr := name.As("username")
	CheckBuildExpr(t, expr, "`users`.`name` AS `username`", nil)
}

// TestValue tests assignment value
func TestValue(t *testing.T) {
	name := NewField("users", "name")
	expr := name.Value("John Doe")
	CheckBuildExpr(t, expr, "`name` = ?", []interface{}{"John Doe"})
}

// TestEqCol tests column equality comparison
func TestEqCol(t *testing.T) {
	field1 := NewField("table1", "id")
	field2 := NewField("table2", "ref_id")
	expr := field1.EqCol(field2)
	CheckBuildExpr(t, expr, "`table1`.`id` = `table2`.`ref_id`", nil)
}

// TestMulCol tests column multiplication
func TestMulCol(t *testing.T) {
	quantity := NewField("orders", "quantity")
	unitPrice := NewField("orders", "unit_price")
	expr := quantity.MulCol(unitPrice)
	CheckBuildExpr(t, expr, "(`orders`.`quantity`) * (`orders`.`unit_price`)", nil)
}

// TestConcatCol tests column concatenation
func TestConcatCol(t *testing.T) {
	first := NewField("users", "first_name")
	last := NewField("users", "last_name")
	expr := first.ConcatCol(last)
	CheckBuildExpr(t, expr, "CONCAT(`users`.`first_name`,`users`.`last_name`)", nil)
}

// TestOr tests OR condition
func TestOr(t *testing.T) {
	age := NewField("users", "age")
	name := NewField("users", "name")
	expr := Or(age.Eq(18), name.Like("John%"))
	CheckBuildExpr(t, expr, "(`users`.`age` = ? OR `users`.`name` LIKE ?)", []interface{}{18, "John%"})
}

// TestAnd tests AND condition
func TestAnd(t *testing.T) {
	age := NewField("users", "age")
	name := NewField("users", "name")
	expr := And(age.Gte(18), name.Like("John%"))
	CheckBuildExpr(t, expr, "(`users`.`age` >= ? AND `users`.`name` LIKE ?)", []interface{}{18, "John%"})
}

// TestNot tests NOT condition
func TestNot(t *testing.T) {
	age := NewField("users", "age")
	expr := Not(age.Eq(18))
	CheckBuildExpr(t, expr, "`users`.`age` <> ?", []interface{}{18})
}

// TestNewUnsafeFieldRaw tests raw SQL field creation
func TestNewUnsafeFieldRaw(t *testing.T) {
	expr := NewUnsafeFieldRaw("COUNT(*)")
	CheckBuildExpr(t, expr, "COUNT(*)", nil)
}
