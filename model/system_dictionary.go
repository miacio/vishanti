package model

import (
	"time"
)

// SystemDictionary 字典表
type SystemDictionary struct {
	ID          string    `db:"id" json:"id" xml:"id"`                               // ID id 主键uuid
	Name        string    `db:"name" json:"name" xml:"name"`                         // Name name 名称
	Group       string    `db:"group" json:"group" xml:"group"`                      // Group group 组名
	ParentGroup string    `db:"parent_group" json:"parent_group" xml:"parent_group"` // ParentGroup parent_group 上级组名
	Describe    string    `db:"describe" json:"describe" xml:"describe"`             // Describe describe 描述
	Val         string    `db:"val" json:"val" xml:"val"`                            // Val val 值
	CreateTime  time.Time `db:"create_time" json:"create_time" xml:"create_time"`    // CreateTime create_time 创建时间
	CreateBy    string    `db:"create_by" json:"create_by" xml:"create_by"`          // CreateBy create_by 创建人id
	UpdateTime  time.Time `db:"update_time" json:"update_time" xml:"update_time"`    // UpdateTime update_time 修改时间
	UpdateBy    time.Time `db:"update_by" json:"update_by" xml:"update_by"`          // UpdateBy update_by 修改人id
}

// TableName SystemDictionary system_dictionary
func (SystemDictionary) TableName() string {
	return "system_dictionary"
}
