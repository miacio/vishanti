package model

import (
	"time"
)

// UserAccountInfo 用户账号表
type UserAccountInfo struct {
	ID         string     `db:"id" json:"id" xml:"id"`                            // ID id 主键uuid
	Mobile     *string    `db:"mobile" json:"mobile" xml:"mobile"`                // Mobile mobile 手机号
	Email      *string    `db:"email" json:"email" xml:"email"`                   // Email email 邮箱
	Account    string     `db:"account" json:"account" xml:"account"`             // Account account 账号
	Password   string     `db:"password" json:"password" xml:"password"`          // Password password 密码
	CreateTime *time.Time `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 创建时间
	UpdateTime *time.Time `db:"update_time" json:"update_time" xml:"update_time"` // UpdateTime update_time 修改时间
	Status     string     `db:"status" json:"status" xml:"status"`                // Status status 账号状态: [USER_ACCESS_STATUS]
	LockTime   *time.Time `db:"lock_time" json:"lock_time" xml:"lock_time"`       // LockTime lock_time 封号时间: 封号的到期时间
}

// TableName UserAccountInfo user_account_info
func (UserAccountInfo) TableName() string {
	return "user_account_info"
}
