package store

import (
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
)

// 文件持久层
type systemFileStore struct{}

type ISystemFileStore interface {
	Insert(systemFileInfo model.SystemFileInfo) error // Insert 写入文件
	FindById(string) (model.SystemFileInfo, error)    // FindById 依据id获取文件信息
}

var SystemFileStore ISystemFileStore = (*systemFileStore)(nil)

// Insert 写入文件
func (*systemFileStore) Insert(systemFileInfo model.SystemFileInfo) error {
	eg := sqlt.NewSQLEngine[model.SystemFileInfo](lib.DB)
	_, err := eg.InsertNamed("db", systemFileInfo).Exec()
	return err
}

// FindById 依据id获取文件信息
func (*systemFileStore) FindById(id string) (model.SystemFileInfo, error) {
	var result model.SystemFileInfo
	err := sqlt.NewSQLEngine[model.SystemFileInfo](lib.DB).Where("id = ?", id).Get(&result)
	return result, err
}
