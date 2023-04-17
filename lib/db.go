package lib

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBCfgParam struct {
	Host      string `toml:"host"`      // 数据库地址 127.0.0.1:3306
	User      string `toml:"user"`      // 数据库用户名
	Password  string `toml:"password"`  // 数据库密码
	Database  string `toml:"database"`  // 数据库名
	Charset   string `toml:"charset"`   // 数据库字符集 utf8mb4
	ParseTime string `toml:"parseTime"` // 是否分析时间 True
	Loc       string `toml:"loc"`       // loc Local
}

func (dp *DBCfgParam) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dp.User, dp.Password, dp.Host, dp.Database)

	params := make(map[string]string, 0)
	if dp.Charset != "" {
		params["charset"] = dp.Charset
	}
	if dp.ParseTime != "" {
		params["parseTime"] = dp.ParseTime
	}
	if dp.Loc != "" {
		params["loc"] = dp.Loc
	}

	vals := make([]string, 0)
	for k, v := range params {
		vals = append(vals, k+"="+v)
	}
	ps := strings.Join(vals, "&")
	if ps != "" {
		dsn = dsn + "?" + ps
	}
	return dsn
}

func (dp *DBCfgParam) Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dp.DSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
