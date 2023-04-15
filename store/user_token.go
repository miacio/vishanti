package store

import (
	"context"
	"encoding/json"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
)

type userTokenStore struct{}

type IUserTokenStore interface {
	LoginSave(userId string) (string, error) // 登录信息存储 returnes 存储key, 错误信息
	Get(key string) (*UserStoreModel, error) // 获取登录数据结构体
}

// 用户存储数据结构体
type UserStoreModel struct {
	AccountInfo  *model.UserAccountInfo  `json:"accountInfo"`  // 账号信息
	DetailedInfo *model.UserDetailedInfo `json:"detailedInfo"` // 用户信息
}

var UserTokenStore IUserTokenStore = (*userTokenStore)(nil)

// 登录信息存储 returnes 存储key, 错误信息
func (*userTokenStore) LoginSave(userId string) (string, error) {
	accountInfo, err := UserStore.FindAccountById(userId)
	if err != nil {
		return "", err
	}

	detailInfo, err := UserStore.FindDetailedByUserId(userId)
	if err != nil {
		return "", err
	}

	msg := UserStoreModel{
		AccountInfo:  accountInfo,
		DetailedInfo: detailInfo,
	}

	key := lib.UID()
	rc := context.Background()
	err = lib.RedisClient.Set(rc, "USER:LOGIN:"+key, util.ToJSON(msg), 0).Err()
	return key, err
}

// 获取登录数据结构体
func (*userTokenStore) Get(key string) (*UserStoreModel, error) {
	rc := context.Background()
	msg, err := lib.RedisClient.Get(rc, "USER:LOGIN:"+key).Result()
	if err != nil {
		return nil, err
	}
	var result UserStoreModel
	err = json.Unmarshal([]byte(msg), &result)
	return &result, err
}
