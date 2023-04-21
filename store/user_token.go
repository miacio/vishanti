package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
)

type userTokenStore struct{}

type IUserTokenStore interface {
	LoginSave(userId string) (string, error)       // 登录信息存储 returnes 存储key, 错误信息
	LoginFlush(key, userId string) (string, error) // 登录信息存储刷新 returnes 存储key, 错误信息
	Get(key string) (*UserStoreModel, error)       // 获取登录数据结构体
}

// 用户存储数据结构体
type UserStoreModel struct {
	AccountInfo  *model.UserAccountInfo  `json:"accountInfo"`  // 账号信息
	DetailedInfo *model.UserDetailedInfo `json:"detailedInfo"` // 用户信息
	HeadPic      string                  `json:"headPic"`      // 用户头像地址
}

var UserTokenStore IUserTokenStore = (*userTokenStore)(nil)

// 登录信息存储 returnes 存储key, 错误信息
func (*userTokenStore) LoginSave(userId string) (string, error) {
	key := lib.UID()
	return UserTokenStore.LoginFlush(key, userId)
}

// 登录信息存储刷新 returnes 存储key, 错误信息
func (*userTokenStore) LoginFlush(key, userId string) (string, error) {
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

func TokenGet(ctx *gin.Context) (*UserStoreModel, bool) {
	obj, ok := ctx.Get("token")
	if !ok {
		lib.ServerResult(ctx, 500, "获取登录信息失败", nil, nil)
		return nil, false
	}
	return obj.(*UserStoreModel), ok
}

func TokenFlush(ctx *gin.Context) error {
	tk := ctx.GetHeader("token")
	if tk == "" {
		return errors.New("token获取失败")
	}
	obj, err := UserTokenStore.Get(tk)
	if err != nil {
		return err
	}
	_, err = UserTokenStore.LoginFlush(tk, obj.AccountInfo.ID)
	return err
}
