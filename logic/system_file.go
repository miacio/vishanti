package logic

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
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
	Upload(*gin.Context)                                                                                                            // Upload 文件上传
	UploadLogic(accountId string, logicName string, req fileDefaultUploadRequest, suffixs ...string) (*model.SystemFileInfo, error) // UploadLogic 上传逻辑 - 封装
	Load(*gin.Context)                                                                                                              // Load 文件读取
}

var SystemFileLogic ISystemFileLogic = (*systemFileLogic)(nil)

// fileDefaultUploadRequest 文件默认上传请求参数
type fileDefaultUploadRequest struct {
	File   *multipart.FileHeader `form:"file" binding:"required"`                                   // 文件对象
	MD5    string                `form:"md5" binding:"required"`                                    // 文件md5
	Region string                `form:"region" binding:"required,oneof=miajiodb beijing hangzhou"` // 文件上传地区
}

// 绑定文件默认上传请求结构体
func (f *fileDefaultUploadRequest) CheckMD5() error {
	mf, err := f.File.Open()
	if err != nil {
		return err
	}
	defer mf.Close()
	reMD5 := util.MD5File(mf)
	if reMD5 != f.MD5 {
		lib.Log.Errorf("md5不一致,系统检测的md5是:%s 接收的md5是:%s", reMD5, f.MD5)
		return errors.New("md5不一致")
	}
	return nil
}

// Upload 文件上传
func (*systemFileLogic) Upload(ctx *gin.Context) {
	var req fileDefaultUploadRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	systemFileInfoModel, err := SystemFileLogic.UploadLogic(mo.AccountInfo.ID, "USER_FILES", req)
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "上传成功", systemFileInfoModel.ID)
}

// UploadLogic 上传逻辑 - 封装
// suffixs is lower string demo: [".jpg",".text"...]
func (*systemFileLogic) UploadLogic(accountId string, logicName string, req fileDefaultUploadRequest, suffixs ...string) (*model.SystemFileInfo, error) {
	if err := req.CheckMD5(); err != nil {
		return nil, err
	}
	objTmpl := "%s/%s/%d/%s"
	suffix := filepath.Ext(filepath.Base(req.File.Filename))
	if suffixs != nil && len(suffix) > 0 {
		if !util.SliceContain(suffixs, strings.ToLower(suffix)) {
			return nil, fmt.Errorf("文件格式错误,目前仅支持[%s]文件格式", strings.Join(suffixs, ","))
		}
	}

	fileName := strings.Join([]string{lib.UID(), suffix}, "")

	objectName := fmt.Sprintf(objTmpl, accountId, strings.ToUpper(logicName), time.Now().UnixMicro(), fileName)
	fileSize := req.File.Size

	mf, err := req.File.Open()
	if err != nil {
		return nil, err
	}
	defer mf.Close()

	err = lib.Minio.PutObject(lib.MinioCfg.Bucket, req.Region, objectName, mf, -1)
	if err != nil {
		return nil, err
	}

	systemFileInfoModel := model.SystemFileInfo{
		ID:         lib.UID(),
		FileName:   req.File.Filename,
		ObjectName: objectName,
		Region:     req.Region,
		Bucket:     lib.MinioCfg.Bucket,
		FileSize:   fileSize,
		FileMd5:    req.MD5,
		CreateTime: model.JsonTimeNow(),
		CreateBy:   accountId,
		Used:       0,
	}
	err = store.SystemFileStore.Insert(systemFileInfoModel)
	if err != nil {
		return nil, err
	}
	return &systemFileInfoModel, nil
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
