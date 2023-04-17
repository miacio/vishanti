package sqlt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
)

func TestSelect(t *testing.T) {
	se := sqlt.NewSQLEngine[model.UserAccountInfo](nil)
	se.And("mobile = ? and password = MD5(?)", "18616220047", "123456").Or("id = ?", "123456")
	sql, params := se.Select("count(1)")
	fmt.Println(sql, params)
}

func TestUpdate(t *testing.T) {
	se := sqlt.NewSQLEngine[model.UserAccountInfo](nil)
	se.Set("mobile = ?", "18570088134").Set("password = MD5(?)", "654321")
	sql, params, err := se.And("id = ?", "123456").Update()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql, params)
}

func TestDelete(t *testing.T) {
	se := sqlt.NewSQLEngine[model.UserAccountInfo](nil)
	sql, params, err := se.Where("id = ?", "123456").Or("id = ?", "654321").Delete()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql, params)
}

func TestInsert(t *testing.T) {
	se := sqlt.NewSQLEngine[model.UserAccountInfo](nil)
	sql, err := se.InsertNamed("db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql)

	te := time.Now()
	jt := model.JsonTime(te)
	val, err := sqlt.ObjectToTagMap(model.UserAccountInfo{
		ID:         lib.UID(),
		Mobile:     "18570088134",
		Email:      "29160047@qq.com",
		Account:    "29160047",
		Password:   util.MD5([]byte("123456")),
		CreateTime: &jt,
		Status:     "1",
	}, "db")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)
}
