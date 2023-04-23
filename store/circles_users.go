package store

import (
	"errors"

	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
)

type circlesUsersStore struct{}

type ICirclesUsersStore interface {
	FindByUserId(userId string) ([]model.CirclesUsers, error) // FindByUserId 查看当前用户所在圈子列表
	Create(model.CirclesUsers) (string, error)                // Create 创建圈子用户
}

var CirclesUsersStore ICirclesUsersStore = (*circlesUsersStore)(nil)

// FindByUserId 查看当前用户所在圈子列表
func (*circlesUsersStore) FindByUserId(userId string) ([]model.CirclesUsers, error) {
	var result []model.CirclesUsers
	circlesUsersEngine := sqlt.NewSQLEngine[model.CirclesUsers](lib.DB)
	err := circlesUsersEngine.Where("user_id = ? and out_time is null and is_sign_out = 3", userId).Find(&result)
	return result, err
}

// Create 创建圈子用户
func (*circlesUsersStore) Create(circlesUsers model.CirclesUsers) (string, error) {
	circlesUsersEngine := sqlt.NewSQLEngine[model.CirclesUsers](lib.DB)
	var oldCirclesUsers model.CirclesUsers
	err := circlesUsersEngine.Where("circles_id = ? and user_id = ?", circlesUsers.CirclesID, circlesUsers.UserID).Select().Get(&oldCirclesUsers)
	if err != nil {
		return "", err
	}

	if oldCirclesUsers.ID != "" {
		name, err := SystemDictionaryStore.FindGroupAndValByName("CIRCLES_SIGN_OUT", oldCirclesUsers.IsSignOut)
		if err != nil {
			return "", err
		}
		switch name {
		case "邀请中":
			return "", errors.New("当前用户正在受邀加入该圈子中")
		case "申请中":
			return "", errors.New("当前用户正在申请加入该圈子中")
		case "登记":
			return "", errors.New("当前用户已经在该圈子中")
		case "退出":
			updateCirclesUsersEngine := sqlt.NewSQLEngine[model.CirclesUsers](lib.DB)
			if _, err = updateCirclesUsersEngine.Set("is_sign_out = ?, create_time = ?, out_time = null", circlesUsers.IsSignOut, model.JsonTimeNow()).Where("id = ?", oldCirclesUsers.ID).Update().Exec(); err != nil {
				return "", err
			}
			return oldCirclesUsers.ID, err
		default:
			return "", errors.New("当前用户已经在该圈子中")
		}
	}

	circlesUsers.ID = lib.UID()
	circlesUsers.CreateTime = model.JsonTimeNow()

	insertCirclesUsersEngine := sqlt.NewSQLEngine[model.CirclesUsers](lib.DB)
	_, err = insertCirclesUsersEngine.InsertNamed("db", circlesUsers).Exec()
	return circlesUsers.ID, err

}
