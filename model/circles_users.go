package model

// CirclesUsers 圈子用户表
type CirclesUsers struct {
	ID         string    `db:"id" json:"id" xml:"id"`                            // ID id 圈子用户id
	CirclesID  string    `db:"circles_id" json:"circles_id" xml:"circles_id"`    // CirclesID circles_id 圈子id
	UserID     string    `db:"user_id" json:"user_id" xml:"user_id"`             // UserID user_id 用户id
	HeadPic    string    `db:"head_pic" json:"head_pic" xml:"head_pic"`          // HeadPic head_pic 圈子用户头像id
	Experience string    `db:"experience" json:"experience" xml:"experience"`    // Experience experience 用户在该圈子的经验值
	CreateTime *JsonTime `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 用户加入圈子的时间
	OutTime    *JsonTime `db:"out_time" json:"out_time" xml:"out_time"`          // OutTime out_time 用户退出圈子的时间
	IsSignOut  string    `db:"is_sign_out" json:"is_sign_out" xml:"is_sign_out"` // IsSignOut is_sign_out 是否退出圈子 字典[CIRCLES_SIGN_OUT]
}

// TableName CirclesUsers circles_users
func (CirclesUsers) TableName() string {
	return "circles_users"
}
