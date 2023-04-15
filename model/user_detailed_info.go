package model

// UserDetailedInfo 用户信息表
type UserDetailedInfo struct {
	ID            string `db:"id" json:"id" xml:"id"`                                        // ID id 主键uuid
	UserAccountID string `db:"user_account_id" json:"user_account_id" xml:"user_account_id"` // UserAccountID user_account_id 用户账号id
	Vip           string `db:"vip" json:"vip" xml:"vip"`                                     // Vip vip vip类别: [USER_VIP]
	HeadPicID     string `db:"head_pic_id" json:"head_pic_id" xml:"head_pic_id"`             // HeadPicID head_pic_id 用户头像 - 文件id
	NickName      string `db:"nick_name" json:"nick_name" xml:"nick_name"`                   // NickName nick_name 用户昵称
	Sex           string `db:"sex" json:"sex" xml:"sex"`                                     // Sex sex 用户性别: 0 未知 1 男 2 女
	BirthdayYear  int    `db:"birthday_year" json:"birthday_year" xml:"birthday_year"`       // BirthdayYear birthday_year 用户生日-年
	BirthdayMonth int    `db:"birthday_month" json:"birthday_month" xml:"birthday_month"`    // BirthdayMonth birthday_month 用户生日-月
	BirthdayDay   int    `db:"birthday_day" json:"birthday_day" xml:"birthday_day"`          // BirthdayDay birthday_day 用户生日-日
	Profile       string `db:"profile" json:"profile" xml:"profile"`                         // Profile profile 个人简介
}

// TableName UserDetailedInfo user_detailed_info
func (UserDetailedInfo) TableName() string {
	return "user_detailed_info"
}
