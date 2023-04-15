package example_test

import (
	"fmt"
	"testing"

	"github.com/miacio/vishanti/store"
)

func init() {
	start()
}

func TestSQL(t *testing.T) {
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
