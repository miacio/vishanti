package lib

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioCfgParam struct {
	EndPoint  string `toml:"endPoint"`  // minio服务器地址 服务ID-地址
	AccessKey string `toml:"accessKey"` // minio 访问密钥
	SecretKey string `toml:"secretKey"` // minio 保密密钥
	UseSSL    bool   `toml:"useSSL"`    // minio 服务协议是否使用https
}

func (mcp *MinioCfgParam) Connect() (*minio.Client, error) {
	minioClient, err := minio.New(mcp.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(mcp.AccessKey, mcp.SecretKey, ""),
		Secure: mcp.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClient, err
}
