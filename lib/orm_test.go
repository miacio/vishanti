package lib_test

import (
	"fmt"
	"testing"

	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
)

func TestSelect(t *testing.T) {
	o := model.UserAccountInfo{}
	se := lib.NewSQLEngine(nil, o)
	se.And("mobile = ? and password = MD5(?)", "18616220047", "123456").Or("id = ?", "123456")
	sql, params := se.Select()
	fmt.Println(sql, params)
}

func TestUpdate(t *testing.T) {
	o := model.UserAccountInfo{}
	se := lib.NewSQLEngine(nil, o)
	se.Set("mobile = ?", "18570088134").Set("password = MD5(?)", "654321")
	sql, params, err := se.And("id = ?", "123456").Update()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql, params)
}
