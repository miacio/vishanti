package lib

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioCfgParam minio参数配置
type MinioCfgParam struct {
	EndPoint  string `toml:"endPoint"`  // minio服务器地址 服务ID-地址
	AccessKey string `toml:"accessKey"` // minio 访问密钥
	SecretKey string `toml:"secretKey"` // minio 保密密钥
	UseSSL    bool   `toml:"useSSL"`    // minio 服务协议是否使用https

	Bucket string `toml:"bucket"` // 文件上传使用的桶
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

// MinioClient minio服务
type MinioClient struct {
	cli *minio.Client
}

// 设置客户端
func (m *MinioClient) SetClient(cli *minio.Client) {
	m.cli = cli
}

// 获取客户端
func (m *MinioClient) GetClient() *minio.Client {
	return m.cli
}

// MakeBucket 创建minio桶
func (m *MinioClient) MakeBucket(bucket string, region string) error {
	ctx := context.Background()
	ok, err := m.cli.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	err = m.cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{
		Region: region,
	})
	return err
}

// PutObject 文件上传
func (m *MinioClient) PutObject(bucket, region, objectName string, file io.Reader, fileSize int64) error {
	if err := m.MakeBucket(bucket, region); err != nil {
		return err
	}
	ctx := context.Background()
	_, err := m.cli.PutObject(ctx, bucket, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	return err
}

// FPutObject 文件上传
func (m *MinioClient) FPutObject(bucket, region, objectName, filePath string) error {
	if err := m.MakeBucket(bucket, region); err != nil {
		return err
	}
	ctx := context.Background()
	_, err := m.cli.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	return err
}

// GetObject 获取文件
func (m *MinioClient) GetObject(bucket, objectName string) ([]byte, error) {
	ctx := context.Background()
	obj, err := m.cli.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()
	file, err := io.ReadAll(obj)
	return file, err
}

// FGetObject 获取文件
func (m *MinioClient) FGetObject(bucket, objectName, filePath string) error {
	ctx := context.Background()
	return m.cli.FGetObject(ctx, bucket, objectName, filePath, minio.GetObjectOptions{})
}
