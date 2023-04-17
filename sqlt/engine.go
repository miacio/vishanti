package sqlt

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

// 操作类型
type optionType int

const (
	_      optionType = iota
	SELECT            // 查询
	UPDATE            // 修改
	DELETE            // 删除
	INSERT            // 新增
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
type SQLEngine[T TableObject] struct {
	to          T
	db          *sqlx.DB // sqlx db 用于直接操作数据库
	whereClause []Clause // where子句
	setClause   []Clause // set子句

	optionType  optionType // 操作类型
	optionError error      // 执行操作类型时产生的错误

	sql    string        // 生成的sql语句
	params []interface{} // 参数集
}

func NewSQLEngine[T TableObject](db *sqlx.DB) *SQLEngine[T] {
	return &SQLEngine[T]{
		db:          db,
		whereClause: make([]Clause, 0),
		setClause:   make([]Clause, 0),
		optionType:  0,
		optionError: nil,
		sql:         "",
		params:      make([]interface{}, 0),
	}
}

func (se *SQLEngine[T]) Clear() *SQLEngine[T] {
	se.whereClause = make([]Clause, 0)
	se.setClause = make([]Clause, 0)
	se.optionType = 0
	se.optionError = nil
	se.sql = ""
	se.params = make([]interface{}, 0)
	return se
}

// whereAppend
// end 是否结束当前子句
func (se *SQLEngine[T]) whereAppend(end bool, condition string, vals ...interface{}) {
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
func (se *SQLEngine[T]) CloseClause() *SQLEngine[T] {
	if len(se.whereClause) > 0 {
		clause := se.whereClause[len(se.whereClause)-1]
		clause.End = true
		se.whereClause[len(se.whereClause)-1] = clause
	}
	return se
}

// Where
func (se *SQLEngine[T]) Where(condition string, vals ...interface{}) *SQLEngine[T] {
	se.CloseClause()
	se.whereAppend(false, condition, vals...)
	return se
}

// And
func (se *SQLEngine[T]) And(condition string, vals ...interface{}) *SQLEngine[T] {
	se.whereAppend(false, condition, vals...)
	return se
}

// Or
func (se *SQLEngine[T]) Or(condition string, vals ...interface{}) *SQLEngine[T] {
	se.CloseClause()
	se.whereAppend(false, condition, vals...)
	return se
}

// where where子句生成器
func (se *SQLEngine[T]) where() (string, []interface{}) {
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
func (se *SQLEngine[T]) Set(condition string, vals ...interface{}) *SQLEngine[T] {
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
func (se *SQLEngine[T]) set() (string, []interface{}) {
	sqlChain := make([]string, 0)
	params := make([]interface{}, 0)
	for _, c := range se.setClause {
		sqlChain = append(sqlChain, strings.Join(c.Condition, ","))
		params = append(params, c.Params...)
	}
	return strings.Join(sqlChain, ","), params
}

// Select select语句生成器
func (se *SQLEngine[T]) Select(columns ...string) *SQLEngine[T] {
	if se.optionType == 0 && se.optionError == nil {
		sqlTemp := "SELECT %s FROM %s"
		var sql string
		if columns == nil {
			columns = se.extractColumn("db", true)
		}
		if columns != nil {
			sql = fmt.Sprintf(sqlTemp, strings.Join(columns, ","), se.to.TableName())
		} else {
			sql = fmt.Sprintf(sqlTemp, "*", se.to.TableName())
		}

		where, params := se.where()
		if where != "" {
			sql = sql + " WHERE " + where
		}
		se.sql = sql
		se.params = params
		se.optionType = SELECT
	} else {
		se.optionError = errors.New("sql engine only option error")
	}
	return se
}

// Update update语句生成器
// 此方法执行时如果update语句没有where条件将会抛出错误
func (se *SQLEngine[T]) Update() *SQLEngine[T] {
	if se.optionType == 0 && se.optionError == nil {
		sqlTemp := "UPDATE %s SET %s WHERE %s"
		params := make([]interface{}, 0)
		setSql, setParams := se.set()
		if setSql == "" {
			se.optionError = errors.New("no update column")
			return se
		}
		whereSql, whereParams := se.where()
		if whereSql == "" {
			se.optionError = errors.New("where clause is empty")
			return se
		}
		params = append(params, setParams...)
		params = append(params, whereParams...)
		sql := fmt.Sprintf(sqlTemp, se.to.TableName(), setSql, whereSql)
		se.sql = sql
		se.params = params
		se.optionType = UPDATE
	} else {
		se.optionError = errors.New("sql engine only option error")
	}
	return se
}

// Delete delete语句生成器
// 此方法执行时如果delete语句没有where条件将会抛出错误
func (se *SQLEngine[T]) Delete() *SQLEngine[T] {
	if se.optionType == 0 && se.optionError == nil {
		sqlTemp := "DELETE FROM %s WHERE %s"
		params := make([]interface{}, 0)
		whereSql, whereParams := se.where()
		if whereSql == "" {
			se.optionError = errors.New("where clause is empty")
			return se
		}
		params = append(params, whereParams...)
		se.sql = fmt.Sprintf(sqlTemp, se.to.TableName(), whereSql)
		se.params = params
		se.optionType = DELETE
	} else {
		se.optionError = errors.New("sql engine only option error")
	}
	return se
}

// Insert insert named语句生成器(允许生成批量插入)
// 此方法依据tag获取字段名称,并将依据此tag的值设定为列名进行插入语句生成
func (se *SQLEngine[T]) InsertNamed(tag string, objs ...T) *SQLEngine[T] {
	if se.optionType == 0 && se.optionError == nil {
		sqlTemp := "INSERT INTO %s (%s) VALUES (%s)"

		columns := se.extractColumn(tag, false)
		if columns == nil {
			se.optionError = errors.New("columns is not found")
			return se
		}
		valColumns := make([]string, 0)
		for i := range columns {
			valColumns = append(valColumns, ":"+columns[i])
			columns[i] = keywordTo(columns[i])
		}
		se.sql = fmt.Sprintf(sqlTemp, se.to.TableName(), strings.Join(columns, ","), strings.Join(valColumns, ","))
		se.params = make([]interface{}, 0)
		for i := range objs {
			param, err := ObjectToTagMap(objs[i], tag)
			if err != nil {
				se.optionError = err
				return se
			}
			se.params = append(se.params, param)
		}
		se.optionType = INSERT
	} else {
		se.optionError = errors.New("sql engine only option error")
	}
	return se
}

func (se *SQLEngine[T]) extractColumn(tag string, keyword bool) []string {
	result := make([]string, 0)

	valueOf := reflect.ValueOf(se.to)
	typeOf := reflect.TypeOf(se.to)

	if reflect.TypeOf(se.to).Kind() == reflect.Ptr {
		valueOf = reflect.ValueOf(se.to).Elem()
		typeOf = reflect.TypeOf(se.to).Elem()
	}
	numField := valueOf.NumField()
	for i := 0; i < numField; i++ {
		tag := typeOf.Field(i).Tag.Get(tag)
		if len(tag) > 0 && tag != "-" {
			if keyword {
				result = append(result, keywordTo(tag))
			} else {
				result = append(result, tag)
			}
		}
	}
	return result
}

func (se *SQLEngine[T]) Get(obj any, columns ...string) error {
	if se.optionError != nil {
		return se.optionError
	}
	if se.optionType == 0 {
		se.Select(columns...)
	}
	if se.optionType == SELECT {
		err := se.db.Get(obj, se.sql, se.params...)
		se.Clear()
		return err
	}
	return errors.New("unknown option type")
}

func (se *SQLEngine[T]) Find(obj any, columns ...string) error {
	if se.optionError != nil {
		return se.optionError
	}
	if se.optionType == 0 {
		se.Select(columns...)
	}
	if se.optionType == SELECT {
		err := se.db.Select(obj, se.sql, se.params...)
		se.Clear()
		return err
	}
	return errors.New("unknown option type")
}

func (se *SQLEngine[T]) Exec() (sql.Result, error) {
	if se.optionError != nil {
		return nil, se.optionError
	}
	var result sql.Result
	var err error
	switch se.optionType {
	case UPDATE, DELETE:
		result, err = se.db.Exec(se.sql, se.params...)
		se.Clear()
		return result, err
	case INSERT:
		result, err = se.db.NamedExec(se.sql, se.params)
		se.Clear()
		return result, err
	}
	return nil, errors.New("unknow option type")
}

func (se *SQLEngine[T]) Value() (string, []interface{}, error) {
	return se.sql, se.params, se.optionError
}
