package lib

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TableObject interface {
	TableName() string // 返回表名
}

// 子句 - 单个块的子句为 and 当存在多个子句时用 or 拼接
type Clause struct {
	Condition []string      // 语句(条件语句或设置语句) xxx = ? | xxx in (?) ...
	Params    []interface{} // 值
	End       bool          // 是否结束当前子句
}

// 创建一个新的子句
func NewClause() Clause {
	return Clause{
		Condition: make([]string, 0),
		Params:    make([]interface{}, 0),
	}
}

// SQL引擎 - 基础单表操作sql引擎,快速进行单表操作
type SQLEngine struct {
	to          TableObject
	db          *sqlx.DB // sqlx db 用于直接操作数据库
	whereClause []Clause // where子句
	setClause   []Clause // set子句
}

func NewSQLEngine(db *sqlx.DB, to TableObject) *SQLEngine {
	return &SQLEngine{
		db:          db,
		to:          to,
		whereClause: make([]Clause, 0),
	}
}

// whereAppend
// end 是否结束当前子句
func (se *SQLEngine) whereAppend(end bool, condition string, vals ...interface{}) {
	var clause Clause
	if se.whereClause == nil || len(se.whereClause) == 0 {
		clause = NewClause()
		clause.Condition = append(clause.Condition, condition)
		clause.Params = append(clause.Params, vals...)
		clause.End = end
		se.whereClause = append(se.whereClause, clause)
		return
	}
	clause = se.whereClause[len(se.whereClause)-1]
	if !clause.End {
		clause.Condition = append(clause.Condition, condition)
		clause.Params = append(clause.Params, vals...)
		se.whereClause[len(se.whereClause)-1] = clause
	} else {
		newClause := NewClause()
		newClause.Condition = append(newClause.Condition, condition)
		newClause.Params = append(newClause.Params, vals...)
		se.whereClause = append(se.whereClause, newClause)
	}
}

// 关闭当前子句 - 用于进入到下一条or子句 或结束当前or子句操作
func (se *SQLEngine) CloseClause() *SQLEngine {
	if len(se.whereClause) > 0 {
		clause := se.whereClause[len(se.whereClause)-1]
		clause.End = true
		se.whereClause[len(se.whereClause)-1] = clause
	}
	return se
}

// And
func (se *SQLEngine) And(condition string, vals ...interface{}) *SQLEngine {
	se.whereAppend(false, condition, vals...)
	return se
}

// Or
func (se *SQLEngine) Or(condition string, vals ...interface{}) *SQLEngine {
	se.CloseClause()
	se.whereAppend(false, condition, vals...)
	return se
}

// where where子句生成器
func (se *SQLEngine) where() (string, []interface{}) {
	sqlChain := make([]string, 0)
	params := make([]interface{}, 0)
	for _, c := range se.whereClause {
		sqlChain = append(sqlChain, strings.Join(c.Condition, " and "))
		params = append(params, c.Params...)
	}
	if len(sqlChain) > 1 {
		for i := range sqlChain {
			sqlChain[i] = "(" + sqlChain[i] + ")"
		}
	}
	return strings.Join(sqlChain, " or "), params
}

// Set
func (se *SQLEngine) Set(condition string, vals ...interface{}) *SQLEngine {
	var clause Clause
	if se.setClause == nil || len(se.setClause) == 0 {
		clause = NewClause()
	}
	clause.Condition = append(clause.Condition, condition)
	clause.Params = append(clause.Params, vals...)
	se.setClause = append(se.setClause, clause)
	return se
}

// set set子句生成器
func (se *SQLEngine) set() (string, []interface{}) {
	sqlChain := make([]string, 0)
	params := make([]interface{}, 0)
	for _, c := range se.setClause {
		sqlChain = append(sqlChain, strings.Join(c.Condition, ","))
		params = append(params, c.Params...)
	}
	return strings.Join(sqlChain, ","), params
}

// Select select语句生成器
func (se *SQLEngine) Select(columns ...string) (string, []interface{}) {
	sqlTemp := "SELECT %s FROM %s"
	var sql string
	if columns != nil {
		sql = fmt.Sprintf(sqlTemp, strings.Join(columns, ","), se.to.TableName())
	} else {
		sql = fmt.Sprintf(sqlTemp, "*", se.to.TableName())
	}

	where, params := se.where()
	if where != "" {
		sql = sql + " WHERE " + where
	}
	return sql, params
}

// Update update语句生成器
func (se *SQLEngine) Update() (string, []interface{}, error) {
	sqlTemp := "UPDATE %s SET %s WHERE %s"
	params := make([]interface{}, 0)
	setSql, setParams := se.set()
	if setSql == "" {
		return "", nil, errors.New("no update column")
	}
	whereSql, whereParams := se.where()
	if whereSql == "" {
		return "", nil, errors.New("where clause is empty")
	}
	params = append(params, setParams...)
	params = append(params, whereParams...)
	sql := fmt.Sprintf(sqlTemp, se.to.TableName(), setSql, whereSql)
	return sql, params, nil
}
