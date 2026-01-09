# Change: 移植通用 Field 表达式构建器

## Why

项目需要一个轻量级的 SQL 字段表达式构建器，用于构建类型安全的查询条件。不需要完整的类型化字段（Int, Float, String 等），只需要一个通用的 `Field` 类型来构建 SQL 表达式。

当前 `field` 包位于 `_obj/gen/field/` 目录下（从 `github.com/go-gorm/gen` 克隆的仓库），我们将只移植其核心表达式构建能力。

## What Changes

- 从 `_obj/gen/field/` 提取核心代码到项目根目录的 `field/` 目录
- 只保留通用的 `Field` 类型，移除所有类型化字段（Int, Float, String, Bool, Time 等）
- 保持核心表达式构建 API
- 添加必要的测试文件
- 更新 go.mod 添加 gorm.io/gorm 依赖（如果尚未添加）

## Impact

- 受影响的规范: 新增 `field-package` 能力规范
- 受影响的代码:
  - 新增 `field/` 目录及其子文件
  - 更新 `go.mod` 添加依赖

## Files to be Added

从 `_obj/gen/field/` 移植以下文件（简化版本）：

| 文件 | 描述 | 行数 (约) |
|------|------|----------|
| `doc.go` | 包文档 | 3 |
| `expr.go` | 核心表达式接口和实现 | 428 |
| `field.go` | 基础字段类型 | 90 |
| `export.go` | 简化的构造函数（仅通用部分） | ~60 |
| `tag.go` | 字段标签选项 | 110 |
| `function.go` | 函数表达式 | 30 |
| `field_test.go` | 简化的单元测试 | ~50 |

**总计约 770 行代码**

### 不包含的文件

以下类型化字段文件**不包含**在本提案中：
- `int.go` - 整数类型字段（~1000+ 行）
- `float.go` - 浮点类型字段（~180+ 行）
- `string.go` - 字符串类型字段（~270+ 行）
- `bool.go` - 布尔类型字段（~55 行）
- `time.go` - 时间类型字段（~200+ 行）
- `association.go` - 关联字段（~300+ 行）
- `serializer.go` - 序列化字段（~110+ 行）
- `asterisk.go` - 通配符字段（~40 行）
- `assign_attr.go` - 属性赋值（~80 行）

## Key Features

通用 Field 包提供以下核心功能：

1. **通用字段表达式**: 单一 `Field` 类型，接受任意 `interface{}` 值
2. **查询条件构建**: Eq, Neq, Gt, Lt, In, Like, Between 等比较操作
3. **字段操作**: Add, Sub, Mul, Div, Mod 等算术操作
4. **聚合函数**: Count, Sum, Max, Min, Avg, Distinct 等
5. **列间比较**: EqCol, GtCol, AddCol 等列操作
6. **空值处理**: IsNull, IsNotNull, IfNull, Null
7. **排序**: Asc, Desc
8. **别名**: As

## 设计决策

**为什么移除类型化字段？**

- 简化代码维护：减少 3000+ 行类型化代码
- 灵活性：通用 `Field` 类型接受任意值，使用 `interface{}`
- 减少依赖：不需要维护多种类型的重复逻辑

**使用示例**：

```go
// 创建字段
age := field.NewField("users", "age")
name := field.NewField("users", "name")

// 构建表达式（值可以是任意类型）
age.Eq(18)              // age = 18
name.Like("John%")      // name LIKE 'John%'
age.Between(18, 65)     // age BETWEEN 18 AND 65
age.AddCol(field.NewField("users", "bonus"))  // age + bonus
```
