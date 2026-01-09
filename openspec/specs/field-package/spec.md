# field-package Specification

## Purpose
TBD - created by archiving change add-field-package. Update Purpose after archive.
## Requirements
### Requirement: Field Package Structure

项目 SHALL 在根目录提供 `field` 包，包含通用的 SQL 字段表达式构建功能。

#### Scenario: Package structure exists

- **WHEN** 导入 `github.com/rockcookies/go-gorms/field`
- **THEN** 包应包含以下核心类型:
  - `Expr` - 表达式接口
  - `Field` - 通用字段类型（唯一字段类型）
  - `AssignExpr` - 赋值表达式接口
  - `OrderExpr` - 排序表达式接口

#### Scenario: Core expression interfaces exist

- **WHEN** 使用 field 包
- **THEN** 以下接口应可用:
  - `Expr` - 查询表达式接口，包含 Build, As, BuildColumn 等方法
  - `AssignExpr` - 继承 Expr，用于赋值操作
  - `OrderExpr` - 继承 Expr，用于排序操作

---

### Requirement: Field Construction

包 SHALL 提供通用的字段构造函数。

#### Scenario: Create generic field

- **GIVEN** 表名 "users" 和列名 "age"
- **WHEN** 调用 `field.NewField("users", "age")`
- **THEN** 返回可用于构建查询的 `Field` 类型

#### Scenario: Create field with raw SQL

- **GIVEN** 原始 SQL "COUNT(*)"
- **WHEN** 调用 `field.NewUnsafeFieldRaw("COUNT(*)")`
- **THEN** 返回包含原始 SQL 的 `Field` 类型

---

### Requirement: Comparison Operations

`Field` 类型 SHALL 支持标准比较操作，接受任意类型值。

#### Scenario: Equal operation

- **GIVEN** 一个 `Field` 表示用户名
- **WHEN** 调用 `field.Eq("john_doe")`
- **THEN** 返回表示 `username = 'john_doe'` 的表达式

#### Scenario: Equal with integer

- **GIVEN** 一个 `Field` 表示年龄
- **WHEN** 调用 `field.Eq(18)`
- **THEN** 返回表示 `age = 18` 的表达式

#### Scenario: Not equal operation

- **GIVEN** 一个 `Field` 表示年龄
- **WHEN** 调用 `field.Neq(18)`
- **THEN** 返回表示 `age != 18` 的表达式

#### Scenario: Greater than operation

- **GIVEN** 一个 `Field` 表示年龄
- **WHEN** 调用 `field.Gt(18)`
- **THEN** 返回表示 `age > 18` 的表达式

#### Scenario: Less than operation

- **GIVEN** 一个 `Field` 表示价格
- **WHEN** 调用 `field.Lt(100.0)`
- **THEN** 返回表示 `price < 100.0` 的表达式

#### Scenario: Between operation

- **GIVEN** 一个 `Field` 表示年龄
- **WHEN** 调用 `field.Between([]interface{}{18, 65})`
- **THEN** 返回表示 `age BETWEEN 18 AND 65` 的表达式

#### Scenario: In operation

- **GIVEN** 一个 `Field` 表示状态
- **WHEN** 调用 `field.In(1, 2, 3)`
- **THEN** 返回表示 `status IN (1, 2, 3)` 的表达式

#### Scenario: Like operation

- **GIVEN** 一个 `Field` 表示名称
- **WHEN** 调用 `field.Like("John%")`
- **THEN** 返回表示 `name LIKE 'John%'` 的表达式

---

### Requirement: Arithmetic Operations

`Field` 类型 SHALL 支持算术操作。

#### Scenario: Addition operation

- **GIVEN** 一个 `Field` 表示库存
- **WHEN** 调用 `field.Add(10)`
- **THEN** 返回表示 `inventory + 10` 的表达式

#### Scenario: Subtraction operation

- **GIVEN** 一个 `Field` 表示库存
- **WHEN** 调用 `field.Sub(5)`
- **THEN** 返回表示 `inventory - 5` 的表达式

#### Scenario: Multiplication operation

- **GIVEN** 一个 `Field` 表示价格
- **WHEN** 调用 `field.Mul(1.1)`
- **THEN** 返回表示 `price * 1.1` 的表达式

#### Scenario: Division operation

- **GIVEN** 一个 `Field` 表示价格
- **WHEN** 调用 `field.Div(2.0)`
- **THEN** 返回表示 `price / 2.0` 的表达式

#### Scenario: Modulo operation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Mod(10)`
- **THEN** 返回表示 `field % 10` 的表达式

---

### Requirement: Column Operations

包 SHALL 支持列间操作。

#### Scenario: Column equality comparison

- **GIVEN** 两个 `Field` 实例 `field1` 和 `field2`
- **WHEN** 调用 `field1.EqCol(field2)`
- **THEN** 返回表示 `field1 = field2` 的表达式

#### Scenario: Column addition

- **GIVEN** 两个 `Field` 实例 `price` 和 `tax`
- **WHEN** 调用 `price.AddCol(tax)`
- **THEN** 返回表示 `price + tax` 的表达式

#### Scenario: Column multiplication

- **GIVEN** 两个 `Field` 实例 `quantity` 和 `unit_price`
- **WHEN** 调用 `quantity.MulCol(unit_price)`
- **THEN** 返回表示 `quantity * unit_price` 的表达式

#### Scenario: Column concatenation

- **GIVEN** 多个 `Field` 实例
- **WHEN** 调用 `first.ConcatCol(second, third)`
- **THEN** 返回表示 `CONCAT(first, second, third)` 的表达式

---

### Requirement: Aggregation Functions

`Field` SHALL 支持聚合函数。

#### Scenario: Count aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Count()`
- **THEN** 返回表示 `COUNT(field)` 的 `Field` 表达式

#### Scenario: Sum aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Sum()`
- **THEN** 返回表示 `SUM(field)` 的 `Field` 表达式

#### Scenario: Max aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Max()`
- **THEN** 返回表示 `MAX(field)` 的 `Field` 表达式

#### Scenario: Min aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Min()`
- **THEN** 返回表示 `MIN(field)` 的 `Field` 表达式

#### Scenario: Average aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Avg()`
- **THEN** 返回表示 `AVG(field)` 的 `Field` 表达式

#### Scenario: Distinct aggregation

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Distinct()`
- **THEN** 返回表示 `DISTINCT field` 的 `Field` 表达式

---

### Requirement: String Functions

`Field` SHALL 支持字符串特定的函数。

#### Scenario: Find in set

- **GIVEN** 一个 `Field` 和逗号分隔的值列表 "a,b,c"
- **WHEN** 调用 `field.FindInSet("a,b,c")`
- **THEN** 返回表示 `FIND_IN_SET(field, 'a,b,c')` 的表达式

#### Scenario: Regular expression match

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Regexp("^[a-z]+$")`
- **THEN** 返回表示 `field REGEXP '^[a-z]+$'` 的表达式

---

### Requirement: Assignment Operations

包 SHALL 支持字段赋值操作。

#### Scenario: Set value for update

- **GIVEN** 一个 `Field` 表示用户名
- **WHEN** 调用 `field.Value("new_name")`
- **THEN** 返回可用于 UPDATE 语句的 `AssignExpr`

#### Scenario: Set NULL value

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Null()`
- **THEN** 返回表示将字段设为 NULL 的 `AssignExpr`

---

### Requirement: Sorting Operations

包 SHALL 支持排序表达式。

#### Scenario: Ascending order

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Asc()`
- **THEN** 返回表示 `field ASC` 的表达式

#### Scenario: Descending order

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Desc()`
- **THEN** 返回表示 `field DESC` 的表达式

---

### Requirement: Null Handling

包 SHALL 支持 NULL 值检查和处理。

#### Scenario: Is NULL check

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.IsNull()`
- **THEN** 返回表示 `field IS NULL` 的表达式

#### Scenario: Is NOT NULL check

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.IsNotNull()`
- **THEN** 返回表示 `field IS NOT NULL` 的表达式

#### Scenario: IFNULL function

- **GIVEN** 一个 `Field` 和默认值
- **WHEN** 调用 `field.IfNull(defaultValue)`
- **THEN** 返回表示 `IFNULL(field, defaultValue)` 的表达式

---

### Requirement: Alias Support

包 SHALL 支持字段别名。

#### Scenario: Set field alias

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.As("total")`
- **THEN** 返回表示 `field AS total` 的表达式

---

### Requirement: Expression Building

包 SHALL 支持 GORM clause.Expression 接口。

#### Scenario: Build expression with GORM

- **GIVEN** 一个 `Field` 表达式
- **WHEN** 传递给 GORM 的查询方法
- **THEN** 表达式应正确构建为 SQL 片段

#### Scenario: Build column with table name

- **GIVEN** 一个带表名的 `Field`
- **WHEN** 调用 `BuildColumn(stmt, field.WithTable)`
- **THEN** 返回带表名的列引用（如 `table.column`）

#### Scenario: Build column with alias

- **GIVEN** 一个带别名的 `Field`
- **WHEN** 调用 `BuildColumn(stmt, field.WithAll)`
- **THEN** 返回带表名和别名的列引用（如 `table.column AS alias`）

---

### Requirement: Field Options

包 SHALL 支持字段选项。

#### Scenario: Create field with options

- **GIVEN** 表名和列名
- **WHEN** 调用 `field.NewField(table, column, opts...)`
- **THEN** 返回应用了选项的 `Field`

#### Scenario: WithTable option exists

- **WHEN** 检查 BuildOpt 常量
- **THEN** 应存在 `field.WithTable` 选项

#### Scenario: WithAll option exists

- **WHEN** 检查 BuildOpt 常量
- **THEN** 应存在 `field.WithAll` 选项

#### Scenario: WithoutQuote option exists

- **WHEN** 检查 BuildOpt 常量
- **THEN** 应存在 `field.WithoutQuote` 选项

---

### Requirement: Go Module Dependencies

field 包的 go.mod SHALL 包含必要的依赖。

#### Scenario: gorm.io/gorm dependency exists

- **WHEN** 检查项目的 go.mod
- **THEN** 应包含 `gorm.io/gorm` 依赖

---

### Requirement: Test Coverage

field 包 SHALL 包含测试文件。

#### Scenario: Unit tests exist

- **WHEN** 列出 field/ 目录中的测试文件
- **THEN** 应包含 `field_test.go`

#### Scenario: Tests pass

- **WHEN** 运行 `go test ./field/`
- **THEN** 所有测试应通过

---

### Requirement: Type Flexibility

`Field` 类型 SHALL 接受任意类型的值。

#### Scenario: Field accepts integer values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq(42)`
- **THEN** 应正确处理整数类型值

#### Scenario: Field accepts string values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq("hello")`
- **THEN** 应正确处理字符串类型值

#### Scenario: Field accepts float values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq(3.14)`
- **THEN** 应正确处理浮点类型值

#### Scenario: Field accepts boolean values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq(true)`
- **THEN** 应正确处理布尔类型值

#### Scenario: Field accepts time.Time values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq(time.Now())`
- **THEN** 应正确处理时间类型值

#### Scenario: Field accepts nil values

- **GIVEN** 一个 `Field`
- **WHEN** 调用 `field.Eq(nil)`
- **THEN** 应正确处理 nil 值

