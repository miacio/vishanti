package store

import (
	"errors"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
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
	err := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB).Where("email = ?", email).Get(&c, "count(1)")
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

// 校验账号是否存在 returnes 存在返回true, 错误信息
func (*userStore) AccountRepeat(account string) (bool, error) {
	var c int
	err := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB).Where("account = ?", account).Get(&c, "count(1)")
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
	accountSQLEngine := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB)
	_, err = accountSQLEngine.InsertNamed("db", model.UserAccountInfo{
		ID:         id,
		Email:      email,
		Account:    account,
		Password:   util.MD5([]byte(password)),
		Status:     "1",
		CreateTime: model.JsonTimeNow(),
	}).Exec()
	if err != nil {
		return "", err
	}

	detailedSQLEngine := sqlt.NewSQLEngine[model.UserDetailedInfo](lib.DB)
	_, err = detailedSQLEngine.InsertNamed("db", model.UserDetailedInfo{
		ID:            lib.UID(),
		UserAccountID: id,
		NickName:      nickName,
	}).Exec()
	return id, err
}

// 依据用户邮箱号获取用户信息
func (*userStore) FindAccountByEmail(email string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB).Where("email = ?", email).Get(&result)
	result.Password = ""
	return &result, err
}

// 依据用户邮箱号和密码获取用户信息
func (*userStore) FindAccountByEmailAndPwd(email, password string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB).Where("email = ? and password = MD5(?)", email, password).Get(&result)
	result.Password = ""
	return &result, err
}

// 依据id获取用户账号信息
func (*userStore) FindAccountById(id string) (*model.UserAccountInfo, error) {
	var result model.UserAccountInfo
	err := sqlt.NewSQLEngine[model.UserAccountInfo](lib.DB).Where("id = ?", id).Get(&result)
	result.Password = ""
	return &result, err
}

// 依据用户id获取用户信息
func (*userStore) FindDetailedByUserId(id string) (*model.UserDetailedInfo, error) {
	var result model.UserDetailedInfo
	err := sqlt.NewSQLEngine[model.UserDetailedInfo](lib.DB).Where("user_account_id = ?", id).Get(&result)
	return &result, err
}

// 修改用户信息
// ID 与 用户账号ID字段不能为空,其余字段数据依据是否为零值判断修改
// VIP字段无法通过此方法进行修改
func (*userStore) UpdateDetailed(userDetailedInfo model.UserDetailedInfo) error {
	se := sqlt.NewSQLEngine[model.UserDetailedInfo](lib.DB)

	oldUserDetailedInfo, err := UserStore.FindDetailedByUserId(userDetailedInfo.UserAccountID)
	if err != nil {
		return err
	}

	if userDetailedInfo.NickName != "" && oldUserDetailedInfo.NickName != userDetailedInfo.NickName {
		se.Set("nick_name = ?", userDetailedInfo.NickName)
	}
	if userDetailedInfo.Sex != "" && oldUserDetailedInfo.Sex != userDetailedInfo.Sex {
		se.Set("sex = ?", userDetailedInfo.Sex)
	}
	if userDetailedInfo.BirthdayYear != 0 && oldUserDetailedInfo.BirthdayYear != userDetailedInfo.BirthdayYear {
		se.Set("birthday_year = ?", userDetailedInfo.BirthdayYear)
	}
	if userDetailedInfo.BirthdayMonth != 0 && oldUserDetailedInfo.BirthdayMonth != userDetailedInfo.BirthdayMonth {
		se.Set("birthday_month = ?", userDetailedInfo.BirthdayMonth)
	}
	if userDetailedInfo.BirthdayDay != 0 && oldUserDetailedInfo.BirthdayDay != userDetailedInfo.BirthdayDay {
		se.Set("birthday_day = ?", userDetailedInfo.BirthdayDay)
	}
	se.Where("id = ? and user_account_id = ?", userDetailedInfo.ID, userDetailedInfo.UserAccountID)
	_, err = se.Update().Exec()
	if err == nil {
		lib.Log.Infof("%s进行修改用户信息操作\n 原数据为: %s", userDetailedInfo.UserAccountID, util.ToJSON(oldUserDetailedInfo))
	}
	return err
}
