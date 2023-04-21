package model

type SystemFileInfo struct {
	ID         string    `db:"id" json:"id" xml:"id"`                            // ID id 文件id
	FileName   string    `db:"file_name" json:"file_name" xml:"file_name"`       // FileName file_name 文件名
	ObjectName string    `db:"object_name" json:"object_name" xml:"object_name"` // ObjectName object_name minio object name
	Region     string    `db:"region" json:"region" xml:"region"`                // Region region minio region
	Bucket     string    `db:"bucket" json:"bucket" xml:"bucket"`                // Bucket bucket minio bucket
	FileSize   int64     `db:"file_size" json:"file_size" xml:"file_size"`       // FileSize file_size 文件大小
	FileMd5    string    `db:"file_md5" json:"file_md5" xml:"file_md5"`          // FileMd5 file_md5 文件md5
	CreateTime *JsonTime `db:"create_time" json:"create_time" xml:"create_time"` // CreateTime create_time 创建时间
	CreateBy   string    `db:"create_by" json:"create_by" xml:"create_by"`       // CreateBy create_by 上传者id
	Used       int       `db:"used" json:"used" xml:"used"`                      // Used used 是否使用
}

// TableName SystemFileInfo system_file_info
func (SystemFileInfo) TableName() string {
	return "system_file_info"
}
