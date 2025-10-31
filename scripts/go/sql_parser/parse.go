package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/xwb1989/sqlparser"
)

// FieldInfo 表示解析出的字段信息
type FieldInfo struct {
	Name       string `json:"name"`                 // 字段名称或表达式
	Alias      string `json:"alias,omitempty"`      // 字段别名
	Table      string `json:"table,omitempty"`      // 表名
	IsStar     bool   `json:"is_star"`              // 是否是 * 通配符
	IsFunction bool   `json:"is_function"`          // 是否是函数
	Function   string `json:"function,omitempty"`   // 函数名
	RawExpr    string `json:"raw_expr,omitempty"`   // 原始表达式
}

// SelectQueryInfo 表示完整的 SELECT 查询信息
type SelectQueryInfo struct {
	Fields     []FieldInfo `json:"fields"`               // 查询字段列表
	Tables     []string    `json:"tables"`               // 涉及的表
	Where      string      `json:"where,omitempty"`      // WHERE 条件
	HasStar    bool        `json:"has_star"`             // 是否包含 *
	FieldCount int         `json:"field_count"`          // 字段数量
	RawSQL     string      `json:"raw_sql,omitempty"`    // 原始 SQL
}

// SQLParser SQL 解析器
type SQLParser struct {
	logger *log.Logger
}

// NewSQLParser 创建新的 SQL 解析器
func NewSQLParser(logger *log.Logger) *SQLParser {
	if logger == nil {
		logger = log.Default()
	}
	return &SQLParser{logger: logger}
}

// ParseSelect 解析 SELECT 语句
func (p *SQLParser) ParseSelect(sql string) (*SelectQueryInfo, error) {
	p.logger.Printf("开始解析 SQL: %s", sql)

	// 清理 SQL 字符串
	cleanedSQL := strings.TrimSpace(sql)
	if cleanedSQL == "" {
		return nil, fmt.Errorf("SQL 语句为空")
	}

	// 解析 SQL
	stmt, err := sqlparser.Parse(cleanedSQL)
	if err != nil {
		p.logger.Printf("SQL 解析失败: %v", err)
		return nil, fmt.Errorf("SQL 解析失败: %w", err)
	}

	// 类型断言，确保是 SELECT 语句
	selectStmt, ok := stmt.(*sqlparser.Select)
	if !ok {
		return nil, fmt.Errorf("不是 SELECT 语句")
	}

	info := &SelectQueryInfo{
		RawSQL: cleanedSQL,
	}

	// 解析查询字段
	if err := p.parseSelectFields(selectStmt, info); err != nil {
		return nil, err
	}

	// 解析表信息
	p.parseTables(selectStmt, info)

	// 解析 WHERE 条件
	p.parseWhereClause(selectStmt, info)

	info.FieldCount = len(info.Fields)

	p.logger.Printf("解析完成: 找到 %d 个字段, %d 个表", info.FieldCount, len(info.Tables))
	return info, nil
}

// parseSelectFields 解析 SELECT 字段
func (p *SQLParser) parseSelectFields(selectStmt *sqlparser.Select, info *SelectQueryInfo) error {
	for i, expr := range selectStmt.SelectExprs {
		field, err := p.parseSelectExpr(expr, i)
		if err != nil {
			return fmt.Errorf("解析第 %d 个字段失败: %w", i+1, err)
		}
		info.Fields = append(info.Fields, field)
		if field.IsStar {
			info.HasStar = true
		}
	}
	return nil
}

// parseSelectExpr 解析单个 SELECT 表达式
func (p *SQLParser) parseSelectExpr(expr sqlparser.SelectExpr, index int) (FieldInfo, error) {
	field := FieldInfo{}

	switch expr := expr.(type) {
	case *sqlparser.AliasedExpr:
		// 处理带别名的表达式
		field.Alias = expr.As.String()
		field.RawExpr = sqlparser.String(expr.Expr)
		
		switch subExpr := expr.Expr.(type) {
		case *sqlparser.ColName:
			// 普通列名
			field.Name = subExpr.Name.String()
			if !subExpr.Qualifier.IsEmpty() {
				field.Table = subExpr.Qualifier.Name.String()
			}
			
		case *sqlparser.FuncExpr:
			// 函数表达式
			field.IsFunction = true
			field.Function = sqlparser.String(subExpr.Name)
			field.Name = sqlparser.String(subExpr)
			
		case *sqlparser.BinaryExpr:
			// 二元表达式，如 a + b
			field.Name = sqlparser.String(subExpr)
			
		case *sqlparser.ParenExpr:
			// 括号表达式
			field.Name = sqlparser.String(subExpr)
			
		case *sqlparser.SQLVal:
			// 字面值，如 SELECT 1, 'test'
			field.Name = string(subExpr.Val)
			
		case *sqlparser.NullVal:
			// NULL 值
			field.Name = "NULL"
			
		case *sqlparser.Subquery:
			// 子查询
			field.Name = fmt.Sprintf("(%s)", sqlparser.String(subExpr))
			
		default:
			// 其他未知表达式类型
			field.Name = sqlparser.String(subExpr)
			p.logger.Printf("警告: 未知表达式类型 %T 在位置 %d", subExpr, index)
		}

	case *sqlparser.StarExpr:
		// 处理 * 通配符
		field.IsStar = true
		field.Name = "*"
		if !expr.TableName.IsEmpty() {
			field.Table = expr.TableName.Name.String()
			field.Name = field.Table + ".*"
		}

	case *sqlparser.Nextval:
		// 序列下一个值，如 NEXT VALUE FOR sequence_name
		field.Name = fmt.Sprintf("NEXT VALUE FOR %s", sqlparser.String(expr.Expr))

	default:
		return field, fmt.Errorf("不支持的 SELECT 表达式类型: %T", expr)
	}

	// 如果没有明确的名称，使用原始表达式
	if field.Name == "" && field.RawExpr != "" {
		field.Name = field.RawExpr
	}

	return field, nil
}

// parseTables 解析涉及的表
func (p *SQLParser) parseTables(selectStmt *sqlparser.Select, info *SelectQueryInfo) {
	for _, tableExpr := range selectStmt.From {
		switch expr := tableExpr.(type) {
		case *sqlparser.AliasedTableExpr:
			switch table := expr.Expr.(type) {
			case sqlparser.TableName:
				if !table.Name.IsEmpty() {
					info.Tables = append(info.Tables, table.Name.String())
				}
			case *sqlparser.Subquery:
				// 子查询作为表
				info.Tables = append(info.Tables, fmt.Sprintf("subquery_%d", len(info.Tables)+1))
			}
		case *sqlparser.JoinTableExpr:
			// 处理 JOIN 语句中的表（简化处理）
			p.logger.Printf("检测到 JOIN 表达式，需要更复杂的表解析")
		}
	}
}

// parseWhereClause 解析 WHERE 条件
func (p *SQLParser) parseWhereClause(selectStmt *sqlparser.Select, info *SelectQueryInfo) {
	if selectStmt.Where != nil {
		info.Where = sqlparser.String(selectStmt.Where.Expr)
	}
}

// FormatAsJSON 将解析结果格式化为 JSON
func (info *SelectQueryInfo) FormatAsJSON() string {
	jsonData, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "JSON 格式化失败: %v"}`, err)
	}
	return string(jsonData)
}

// GetFieldNames 获取所有字段名称（优先使用别名）
func (info *SelectQueryInfo) GetFieldNames() []string {
	names := make([]string, 0, len(info.Fields))
	for _, field := range info.Fields {
		if field.Alias != "" {
			names = append(names, field.Alias)
		} else {
			names = append(names, field.Name)
		}
	}
	return names
}

// HasTableQualifier 检查是否包含表限定符
func (info *SelectQueryInfo) HasTableQualifier() bool {
	for _, field := range info.Fields {
		if field.Table != "" {
			return true
		}
	}
	return false
}

// 示例使用
func main() {
	// 创建解析器
	parser := NewSQLParser(nil)

	// 测试各种 SQL 语句
	testSQLs := []string{
		// 基本查询
		// "SELECT *, id, name, email FROM users WHERE age > 18",
		
		// // 带别名
		// "SELECT u.\"id\" AS user_id, u.`name` AS username, u.email FROM users u",
		
		// // 函数调用
		// "SELECT COUNT(*) as total, MAX(age) as max_age, AVG(salary) FROM employees",
		
		// // 表达式
		// "SELECT price * quantity as total_price, (salary * 1.1) as new_salary FROM orders",
		
		// // 复杂情况
		// "SELECT u.*, p.product_name, COUNT(o.id) as order_count FROM users u LEFT JOIN orders o ON u.id = o.user_id LEFT JOIN products p ON o.product_id = p.id GROUP BY u.id, p.product_name",
		
		// // 子查询
		// "SELECT name, (SELECT COUNT(*) FROM orders WHERE users.id = orders.user_id) as order_count FROM users",
		
		// // 字面值
		// "SELECT 1 as constant, 'test' as text, NULL as empty FROM dual",

		// "SELECT lft.customer_id, lft.customer_name, lft.order_id, rgt.product_name FROM {{.node4}} AS lft INNER JOIN {{.node3}} AS rgt ON lft.product_id = rgt.product_id",
	}

	for i, sql := range testSQLs {
		fmt.Printf("=== 测试用例 %d ===\n", i+1)
		fmt.Printf("SQL: %s\n", sql)
		
		info, err := parser.ParseSelect(sql)
		if err != nil {
			fmt.Printf("解析错误: %v\n\n", err)
			continue
		}

		// 输出解析结果
		fmt.Printf("解析结果:\n")
		fmt.Printf("- 字段数量: %d\n", info.FieldCount)
		fmt.Printf("- 包含通配符: %t\n", info.HasStar)
		fmt.Printf("- 涉及表: %v\n", info.Tables)
		
		if info.Where != "" {
			fmt.Printf("- WHERE 条件: %s\n", info.Where)
		}
		
		fmt.Printf("- 字段详情:\n")
		for j, field := range info.Fields {
			output := fmt.Sprintf("  %d. %s", j+1, field.Name)
			if field.Alias != "" && field.Alias != field.Name {
				output += fmt.Sprintf(" AS %s", field.Alias)
			}
			if field.Table != "" {
				output += fmt.Sprintf(" [表: %s]", field.Table)
			}
			if field.IsFunction {
				output += " [函数]"
			}
			if field.IsStar {
				output += " [通配符]"
			}
			fmt.Println(output)
		}
		
		// 输出 JSON 格式
		fmt.Printf("\nJSON 格式:\n%s\n\n", info.FormatAsJSON())
	}
}

// 单元测试示例
func TestSQLParser() {
	parser := NewSQLParser(nil)
	
	// 测试用例
	tests := []struct {
		name     string
		sql      string
		expected int // 期望的字段数量
	}{
		{"简单查询", "SELECT a, b, c FROM table1", 3},
		{"带别名", "SELECT id AS user_id, name FROM users", 2},
		{"函数调用", "SELECT COUNT(*), MAX(age) FROM employees", 2},
		{"通配符", "SELECT * FROM products", 1},
	}

	for _, test := range tests {
		info, err := parser.ParseSelect(test.sql)
		if err != nil {
			log.Printf("测试 %s 失败: %v", test.name, err)
			continue
		}
		
		if info.FieldCount != test.expected {
			log.Printf("测试 %s 失败: 期望 %d 个字段, 实际 %d 个", 
				test.name, test.expected, info.FieldCount)
		} else {
			log.Printf("测试 %s 通过", test.name)
		}
	}
}