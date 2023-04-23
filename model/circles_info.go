package model

// CirclesInfo 圈子主表
type CirclesInfo struct {
	ID         string    `db:"id" json:"id" xml:"id"`                            // ID id 圈子id
	Logo       string    `db:"logo" json:"logo" xml:"logo"`                      // Logo logo 圈子logo id
	Name       string    `db:"name" json:"name" xml:"name"`                      // Name name 圈子名称
	Descirbe   string    `db:"descirbe" json:"descirbe" xml:"descirbe"`          // Descirbe descirbe 圈子描述
	CreateTime *JsonTime `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 圈子的创建时间
	CreateBy   string    `db:"create_by" json:"create_by" xml:"create_by"`       // CreateBy create_by 圈子的创建者id
	Owner      string    `db:"owner" json:"owner" xml:"owner"`                   // Owner owner 圈子的所有者id
	UpdateTime *JsonTime `db:"update_time" json:"update_time" xml:"update_time"` // UpdateTime update_time 圈子的修改时间
	UpdateBy   string    `db:"update_by" json:"update_by" xml:"update_by"`       // UpdateBy update_by 圈子的修改者id
}

// TableName CirclesInfo circles_info
func (CirclesInfo) TableName() string {
	return "circles_info"
}
