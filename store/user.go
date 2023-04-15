package store

import (
	"errors"
	"strings"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
)

// 用户持久层
type userStore struct{}

type IUserStore interface {
	EmailRepeat(email string) (bool, error)                                          // 校验邮箱是否存在 returnes 存在返回true, 错误信息
	AccountRepeat(account string) (bool, error)                                      // 校验账号是否存在 returnes 存在返回true, 错误信息
	EmailRegister(email, nickName, account, password string) (string, error)         // 邮箱注册用户 returnes 用户ID, 错误信息
	FindAccountByEmail(email string) (*model.UserAccountInfo, error)                 // 依据用户邮箱号获取用户信息
	FindAccountByEmailAndPwd(email, password string) (*model.UserAccountInfo, error) // 依据用户邮箱号和密码获取用户信息
	FindAccountById(id string) (*model.UserAccountInfo, error)                       // 依据用户账号id获取用户信息
	FindDetailedByUserId(id string) (*model.UserDetailedInfo, error)                 // 依据用户id获取用户信息
	UpdateDetailed(userDetailedInfo model.UserDetailedInfo) error                    // 修改用户信息
}

var UserStore IUserStore = (*userStore)(nil)

// 校验邮箱是否存在 returnes 存在返回true, 错误信息
func (*userStore) EmailRepeat(email string) (bool, error) {
	var c int
	err := lib.DB.Get(&c, "select count(1) from user_account_info where email = ?", email)
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

// 校验账号是否存在 returnes 存在返回true, 错误信息
func (*userStore) AccountRepeat(account string) (bool, error) {
	var c int
	err := lib.DB.Get(&c, "select count(1) from user_account_info where account = ?", account)
	if err != nil {
		return false, err
	}
	return c > 0, err
}

// 邮箱注册用户
func (u *userStore) EmailRegister(email, nickName, account, password string) (string, error) {
	ok, err := u.EmailRepeat(email)
	if err != nil {
		return "", err
	}
	if ok {
		return "", errors.New("当前邮箱已经被注册,请勿重复注册")
	}

	ok, err = u.AccountRepeat(account)
	if err != nil {
		return "", err
	}
	if ok {
		return "", errors.New("当前账号已经被注册,请勿重复注册")
	}

	id := lib.UID()
	if err := u.emailRegisterTx(id, email, account, password, nickName); err != nil {
		return "", err
	}
	return id, nil
}

// 邮箱注册-事务方式
func (*userStore) emailRegisterTx(id, email, account, password, nickName string) error {
	tx, err := lib.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	_, err = tx.Exec("insert into user_account_info (id, email, account, password, create_time, status) values (?, ?, ?, MD5(?), NOW(), 1)", id, email, account, password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into user_detailed_info (id, user_account_id, nick_name) values (UPPER(REPLACE(UUID(),'-', '')), ?, ?)", id, nickName)
	if err != nil {
		return err
	}
	return nil
}

// 依据用户邮箱号获取用户信息
func (*userStore) FindAccountByEmail(email string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := lib.DB.Get(&result, "select * from user_account_info where email = ?", email)
	result.Password = ""
	return &result, err
}

// 依据用户邮箱号和密码获取用户信息
func (*userStore) FindAccountByEmailAndPwd(email, password string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := lib.DB.Get(&result, "select * from user_account_info where email = ? and password = MD5(?)", email, password)
	result.Password = ""
	return &result, err
}

// 依据id获取用户账号信息
func (*userStore) FindAccountById(id string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := lib.DB.Get(&result, "select * from user_account_info where id = ?", id)
	result.Password = ""
	return &result, err
}

// 依据用户id获取用户信息
func (*userStore) FindDetailedByUserId(id string) (*model.UserDetailedInfo, error) {
	var result model.UserDetailedInfo
	err := lib.DB.Get(&result, "select * from user_detailed_info where user_account_id = ?", id)
	return &result, err
}

// 修改用户信息
// ID 与 用户账号ID字段不能为空,其余字段数据依据是否为零值判断修改
// VIP字段无法通过此方法进行修改
func (*userStore) UpdateDetailed(userDetailedInfo model.UserDetailedInfo) error {
	set_sql := make([]string, 0)
	set_params := make([]interface{}, 0)

	oldUserDetailedInfo, err := UserStore.FindDetailedByUserId(userDetailedInfo.UserAccountID)
	if err != nil {
		return err
	}

	if userDetailedInfo.HeadPicID != "" && oldUserDetailedInfo.HeadPicID != userDetailedInfo.HeadPicID {
		set_sql = append(set_sql, "head_pic_id = ?")
		set_params = append(set_params, userDetailedInfo.HeadPicID)
	}
	if userDetailedInfo.NickName != "" && oldUserDetailedInfo.NickName != userDetailedInfo.NickName {
		set_sql = append(set_sql, "nick_name = ?")
		set_params = append(set_params, userDetailedInfo.NickName)
	}
	if userDetailedInfo.Sex != "" && oldUserDetailedInfo.Sex != userDetailedInfo.Sex {
		set_sql = append(set_sql, "sex = ?")
		set_params = append(set_params, userDetailedInfo.Sex)
	}
	if userDetailedInfo.BirthdayYear != 0 && oldUserDetailedInfo.BirthdayYear != userDetailedInfo.BirthdayYear {
		set_sql = append(set_sql, "birthday_year = ?")
		set_params = append(set_params, userDetailedInfo.BirthdayYear)
	}
	if userDetailedInfo.BirthdayMonth != 0 && oldUserDetailedInfo.BirthdayMonth != userDetailedInfo.BirthdayMonth {
		set_sql = append(set_sql, "birthday_month = ?")
		set_params = append(set_params, userDetailedInfo.BirthdayMonth)
	}
	if userDetailedInfo.BirthdayDay != 0 && oldUserDetailedInfo.BirthdayDay != userDetailedInfo.BirthdayDay {
		set_sql = append(set_sql, "birthday_day = ?")
		set_params = append(set_params, userDetailedInfo.BirthdayDay)
	}

	if len(set_sql) > 0 {
		update_sql := "update user_detailed_info set " + strings.Join(set_sql, ",") + " where id = ? and user_account_id = ?"
		set_params = append(set_params, userDetailedInfo.ID, userDetailedInfo.UserAccountID)
		_, err := lib.DB.Exec(update_sql, set_params...)
		if err != nil {
			return err
		}
		lib.Log.Infof("%s进行修改用户信息操作\n 原数据为: %s", userDetailedInfo.UserAccountID, util.ToJSON(oldUserDetailedInfo))
	}
	return nil
}
