package model

// CirclesServerInfo 圈子三方服务信息
type CirclesServerInfo struct {
	ID         string    `db:"id" json:"id" xml:"id"`                            // ID id 服务信息id
	CirclesID  string    `db:"circles_id" json:"circles_id" xml:"circles_id"`    // CirclesID circles_id 圈子id
	ServerUtl  string    `db:"server_utl" json:"server_utl" xml:"server_utl"`    // ServerUtl server_utl 圈子服务器地址
	PublicKey  string    `db:"public_key" json:"public_key" xml:"public_key"`    // PublicKey public_key 公钥
	SecretKey  string    `db:"secret_key" json:"secret_key" xml:"secret_key"`    // SecretKey secret_key 私钥
	Down       string    `db:"down" json:"down" xml:"down"`                      // Down down 通信是否宕机 字典: [CIRCLES_SERVER_DOWN]
	CreateTime *JsonTime `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 创建时间
	CreateBy   string    `db:"create_by" json:"create_by" xml:"create_by"`       // CreateBy create_by 创建者id
	UpdateTime *JsonTime `db:"update_time" json:"update_time" xml:"update_time"` // UpdateTime update_time 修改时间
	UpdateBy   string    `db:"update_by" json:"update_by" xml:"update_by"`       // UpdateBy update_by 修改者id
	Used       string    `db:"used" json:"used" xml:"used"`                      // Used used 是由启用 字典: [CIRCLES_SERVER_USED]
}

// TableName CirclesServerInfo circles_server_info
func (CirclesServerInfo) TableName() string {
	return "circles_server_info"
}
