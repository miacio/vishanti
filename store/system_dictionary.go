package store

import (
	"errors"

	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/sqlt"
)

// 系统字典
type systemDictionaryStore struct{}

type ISystemDictionaryStore interface {
	FindGroupAndValByName(group, val string) (string, error)                                         // 依据组名和值获取字典名称
	FindGroupAndNameByVal(group, name string) (string, error)                                        // 依据组名和名称获取字典值
	FindById(id string) (model.SystemDictionary, error)                                              // 依据字典id值获取字典信息
	FindByGroup(group string) ([]model.SystemDictionary, error)                                      // 依据字典组获取该组所有字典
	InsertSystemDictionary(systemDictionary model.SystemDictionary) (*model.SystemDictionary, error) // 添加字典方法
	Inserts(systemDictionaryStores []model.SystemDictionary) error                                   // 批量写入字典
}

var SystemDictionaryStore ISystemDictionaryStore = (*systemDictionaryStore)(nil)

// 依据组名和值获取字典名称
func (*systemDictionaryStore) FindGroupAndValByName(group, val string) (string, error) {
	var name string
	err := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB).Where("`group` = ? and val = ?", group, val).Get(&name, "name")
	return name, err
}

// 依据组名和名称获取字典值
func (*systemDictionaryStore) FindGroupAndNameByVal(group, name string) (string, error) {
	var val string
	err := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB).Where("`group` = ? and name = ?", group, name).Get(&name, "val")
	return val, err
}

// 依据字典id值获取字典信息
func (*systemDictionaryStore) FindById(id string) (model.SystemDictionary, error) {
	var systemDictionary model.SystemDictionary
	err := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB).Where("id = ?", id).Get(&systemDictionary)
	return systemDictionary, err
}

// 依据字典组获取该组所有字典
func (*systemDictionaryStore) FindByGroup(group string) ([]model.SystemDictionary, error) {
	var systemDictionarys []model.SystemDictionary
	err := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB).Where("`group` = ?", group).Find(&systemDictionarys)
	return systemDictionarys, err
}

// 添加字典方法
func (*systemDictionaryStore) InsertSystemDictionary(systemDictionary model.SystemDictionary) (*model.SystemDictionary, error) {
	se := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB)

	if systemDictionary.Name == "" {
		return nil, errors.New("名称不能为空")
	}

	if systemDictionary.Val == "" {
		return nil, errors.New("值不能为空")
	}

	if systemDictionary.Group == "" {
		return nil, errors.New("组名不能为空")
	}

	if systemDictionary.ParentGroup != "" {
		var c int
		err := se.Where("`group` = ?", systemDictionary.ParentGroup).Count().Get(&c)
		if err != nil {
			return nil, err
		}
		if c == 0 {
			return nil, errors.New("上级组名不存在")
		}
	}

	systemDictionary.ID = lib.UID()
	_, err := se.InsertNamed("db", systemDictionary).Exec()

	//(:id, :name, :group, :parent_group, :describe, :val, NOW(), :create_by)
	if err != nil {
		return nil, err
	}

	return &systemDictionary, nil
}

// 批量写入字典
func (*systemDictionaryStore) Inserts(systemDictionaryStores []model.SystemDictionary) error {
	_, err := sqlt.NewSQLEngine[model.SystemDictionary](lib.DB).InsertNamed("db", systemDictionaryStores...).Exec()
	return err
}
