# Implementation Tasks

## 1. 准备工作
- [x] 1.1 验证 go.mod 中是否有 gorm.io/gorm 依赖
- [x] 1.2 如需要，添加 gorm.io/gorm 依赖

## 2. 代码移植
- [x] 2.1 创建 `field/` 目录
- [x] 2.2 复制核心文件:
  - [x] 2.2.1 `doc.go` - 包文档
  - [x] 2.2.2 `expr.go` - 表达式接口和实现
  - [x] 2.2.3 `field.go` - 基础字段类型
- [x] 2.3 创建简化的 `export.go`:
  - [x] 2.3.1 只保留 `NewField` 函数
  - [x] 2.3.2 只保留 `NewUnsafeFieldRaw` 函数
  - [x] 2.3.3 只保留 `toColumn` 辅助函数
  - [x] 2.3.4 移除所有类型化构造函数（NewInt, NewString 等）
  - [x] 2.3.5 移除 Star/ALL 常量（除非需要）
- [x] 2.4 复制辅助文件:
  - [x] 2.4.1 `tag.go` - 字段标签选项
  - [x] 2.4.2 `function.go` - 函数表达式

## 3. 测试文件
- [x] 3.1 创建简化的 `field_test.go`
- [x] 3.2 添加通用 Field 类型的测试用例

## 4. 验证
- [x] 4.1 运行 `go build ./field/` 确保编译通过
- [x] 4.2 运行 `go test ./field/` 确保测试通过
- [x] 4.3 运行 `go vet ./field/` 检查代码质量

## 5. 清理
- [x] 5.1 考虑是否需要删除 `_obj/gen` 目录（取决于项目需求）

## 简化版文件清单

**需要复制的文件（约 770 行）**:
- `doc.go` (3 行)
- `expr.go` (428 行)
- `field.go` (90 行)
- `tag.go` (110 行)
- `function.go` (30 行)

**需要创建的文件**:
- `export.go` (约 60 行，简化版)
- `field_test.go` (约 50 行)

**不需要的文件（约 3200+ 行）**:
- `int.go` - 跳过
- `float.go` - 跳过
- `string.go` - 跳过
- `bool.go` - 跳过
- `time.go` - 跳过
- `association.go` - 跳过
- `serializer.go` - 跳过
- `asterisk.go` - 跳过
- `assign_attr.go` - 跳过
- `export_test.go` - 跳过
- `example_test.go` - 跳过

