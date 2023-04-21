package logic

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miacio/varietas/util"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/store"
)

// 文件逻辑
type systemFileLogic struct{}

type ISystemFileLogic interface {
	Upload(ctx *gin.Context) // Upload 文件上传
	Load(ctx *gin.Context)   // Load 文件读取
}

var SystemFileLogic ISystemFileLogic = (*systemFileLogic)(nil)

// fileDefaultUploadRequest 文件默认上传请求参数
type fileDefaultUploadRequest struct {
	File   *multipart.FileHeader `json:"file"`   // 文件对象
	MD5    string                `json:"md5"`    // 文件md5
	Region string                `json:"region"` // 文件上传地区
}

// 绑定文件默认上传请求结构体
func formFileDefaultUploadRequest(ctx *gin.Context, obj *fileDefaultUploadRequest) error {
	reqMd5 := ctx.PostForm("md5")
	region := ctx.PostForm("region")
	if region == "" {
		region = "miajiodb"
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	obj.File = file
	obj.MD5 = reqMd5
	mf, err := file.Open()
	if err != nil {
		return err
	}
	defer mf.Close()
	rmd5 := util.MD5File(mf)
	if rmd5 != reqMd5 {
		return errors.New("md5不一致")
	}

	obj.Region = region
	return nil
}

// Upload 文件上传
func (*systemFileLogic) Upload(ctx *gin.Context) {
	var req fileDefaultUploadRequest
	err := formFileDefaultUploadRequest(ctx, &req)
	if !lib.ServerFail(ctx, err) {
		return
	}

	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	objTmpl := "files/%s/%d/%s"
	fileName := req.File.Filename
	objectName := fmt.Sprintf(objTmpl, mo.AccountInfo.ID, time.Now().UnixMicro(), fileName)
	fileSize := req.File.Size

	mf, err := req.File.Open()
	if !lib.ServerFail(ctx, err) {
		return
	}
	defer mf.Close()

	err = lib.Minio.PutObject(lib.MinioCfg.Bucket, req.Region, objectName, mf, -1)
	if !lib.ServerFail(ctx, err) {
		return
	}

	systemFileInfoModel := model.SystemFileInfo{
		ID:         lib.UID(),
		FileName:   fileName,
		ObjectName: objectName,
		Region:     req.Region,
		Bucket:     lib.MinioCfg.Bucket,
		FileSize:   fileSize,
		FileMd5:    req.MD5,
		CreateTime: model.JsonTimeNow(),
		CreateBy:   mo.AccountInfo.ID,
		Used:       0,
	}

	err = store.SystemFileStore.Insert(systemFileInfoModel)
	if !lib.ServerFail(ctx, err) {
		return
	}

	lib.ServerSuccess(ctx, "上传成功", systemFileInfoModel.ID)
}

// Load 文件读取
func (*systemFileLogic) Load(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		lib.ServerFailf(ctx, 400, "文件id不能为空", nil)
		return
	}

	systemFileInfo, err := store.SystemFileStore.FindById(id)
	if !lib.ServerFail(ctx, err) {
		return
	}
	fileByte, err := lib.Minio.GetObject(systemFileInfo.Bucket, systemFileInfo.ObjectName)
	if !lib.ServerFail(ctx, err) {
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+systemFileInfo.FileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Data(http.StatusOK, "application/octet-stream", fileByte)
}
