package model

// CirclesManagerUsers 圈子管理员表
type CirclesManagerUsers struct {
	ID         string    `db:"id" json:"id" xml:"id"`                            // ID id 圈子管理员id
	CirclesID  string    `db:"circles_id" json:"circles_id" xml:"circles_id"`    // CirclesID circles_id 圈子id
	UserID     string    `db:"user_id" json:"user_id" xml:"user_id"`             // UserID user_id 用户id
	Level      string    `db:"level" json:"level" xml:"level"`                   // Level level 管理员权限等级-圈子的管理权限将依据对应圈子开发者设定的规则而定义,该位置仅用于存储
	CreateTime *JsonTime `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 成为管理员的时间
}

// TableName CirclesManagerUsers circles_manager_users
func (CirclesManagerUsers) TableName() string {
	return "circles_manager_users"
}
