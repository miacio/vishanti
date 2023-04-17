package example_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
	"github.com/miacio/vishanti/store"
)

func TestSQL(t *testing.T) {
	Runner()
	userAccountInfo, err := store.UserStore.FindAccountByEmailAndPwd("miajio@163.com", "123456")
	if err != nil {
		t.Fatal(err)
	}
	userDetailesInfo, err := store.UserStore.FindDetailedByUserId(userAccountInfo.ID)
	if err != nil {
		t.Fatal(err)
	}
	status, err := store.SystemDictionaryStore.FindGroupAndValByName("USER_ACCESS_STATUS", userAccountInfo.Status)
	if err != nil {
		t.Fatal(err)
	}

	vip, err := store.SystemDictionaryStore.FindGroupAndValByName("USER_VIP", userDetailesInfo.Vip)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("当前用户状态为:", status)
	fmt.Println("当前用户会员等级为:", vip)
}

func TestInsert(t *testing.T) {
	Runner()
	se := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB)
	te := time.Now()
	jt := model.JsonTime(te)

	sql, param, err := se.InsertNamed("db", model.UserAccountInfo{
		ID:         lib.UID(),
		Mobile:     "18570088134",
		Email:      "29160047@qq.com",
		Account:    "29160047",
		Password:   util.MD5([]byte("123456")),
		CreateTime: &jt,
		Status:     "1",
	}).Value()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sql)
	fmt.Println(param)
	_, err = se.Exec()
	if err != nil {
		t.Fatal(err)
	}
}
