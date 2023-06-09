package store

import (
	"errors"

	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
)

// 圈子持久层
type circlesStore struct{}

type ICirclesStore interface {
	Create(model.CirclesInfo) (string, error)               // Create 创建圈子
	FindById(accountId string) ([]model.CirclesInfo, error) // FindById 查询该用户拥有的圈子
}

var CirclesStore ICirclesStore = (*circlesStore)(nil)

// Create 创建圈子
func (*circlesStore) Create(circlesInfo model.CirclesInfo) (string, error) {
	checkEngine := sqlt.NewSQLEngine[model.CirclesInfo](lib.DB)
	var count int
	if err := checkEngine.Where("name = ?", circlesInfo.Name).Count().Get(&count); err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("当前圈子名称已存在")
	}

	insertEngine := sqlt.NewSQLEngine[model.CirclesInfo](lib.DB)

	circlesInfo.ID = lib.UID()
	circlesInfo.CreateTime = model.JsonTimeNow()

	_, err := insertEngine.InsertNamed("db", circlesInfo).Exec()
	if err != nil {
		return "", err
	}
	return circlesInfo.ID, nil
}

// FindById 查询该用户拥有的圈子
func (*circlesStore) FindById(accountId string) ([]model.CirclesInfo, error) {
	findEngine := sqlt.NewSQLEngine[model.CirclesInfo](lib.DB)
	var result []model.CirclesInfo
	err := findEngine.Where("owner = ?", accountId).Find(&result)
	return result, err
}
