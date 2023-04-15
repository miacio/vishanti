package store

import (
	"errors"
	"time"

	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
)

// 系统字典
type systemDictionaryStore struct{}

type ISystemDictionaryStore interface {
	FindGroupAndValByName(group, val string) (string, error)                                         // 依据组名和值获取字典名称
	FindById(id string) (model.SystemDictionary, error)                                              // 依据字典id值获取字典信息
	FindByGroup(group string) ([]model.SystemDictionary, error)                                      // 依据字典组获取该组所有字典
	InsertSystemDictionary(systemDictionary model.SystemDictionary) (*model.SystemDictionary, error) // 添加字典方法
}

var SystemDictionaryStore ISystemDictionaryStore = (*systemDictionaryStore)(nil)

const (
	sql_system_dictionary_insert = "insert into system_dictionary (id, `name`, `group`, parent_group, `describe`, val, create_time, create_by) values (:id, :name, :group, :parent_group, :describe, :val, :create_time, :create_by)"
)

// 依据组名和值获取字典名称
func (*systemDictionaryStore) FindGroupAndValByName(group, val string) (string, error) {
	var name string
	err := lib.DB.Get(&name, "select name from system_dictionary where group = ? and val = ?", group, val)
	return name, err
}

// 依据字典id值获取字典信息
func (*systemDictionaryStore) FindById(id string) (model.SystemDictionary, error) {
	var systemDictionary model.SystemDictionary
	err := lib.DB.Get(&systemDictionary, "select * from system_dictionary where id = ?", id)
	return systemDictionary, err
}

// 依据字典组获取该组所有字典
func (*systemDictionaryStore) FindByGroup(group string) ([]model.SystemDictionary, error) {
	var systemDictionarys []model.SystemDictionary
	err := lib.DB.Select(&systemDictionarys, "select * from system_dictionary where group = ?", group)
	return systemDictionarys, err
}

// 添加字典方法
func (*systemDictionaryStore) InsertSystemDictionary(systemDictionary model.SystemDictionary) (*model.SystemDictionary, error) {
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
		err := lib.DB.Get(&c, "select count(1) from system_dictionary where group = ?", systemDictionary.ParentGroup)
		if err != nil {
			return nil, err
		}
		if c == 0 {
			return nil, errors.New("上级组名不存在")
		}
	}

	systemDictionary.ID = lib.UID()
	systemDictionary.CreateTime = time.Now()

	systemDictionaryParam, err := util.Object2Tag(systemDictionary, "db")
	if err != nil {
		return nil, err
	}

	_, err = lib.DB.NamedExec(sql_system_dictionary_insert, systemDictionaryParam)
	if err != nil {
		return nil, err
	}

	return &systemDictionary, nil
}