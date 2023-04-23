package logic

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/store"
)

type circlesLogic struct{}

type ICirclesLogic interface {
	Create(*gin.Context) // Create 创建圈子
	Find(*gin.Context)   // Find 查询自己所拥有的圈子
}

var CirclesLogic ICirclesLogic = (*circlesLogic)(nil)

// createCirclesRequest 创建圈子请求参数
type createCirclesRequest struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`                                   // 文件对象
	MD5     string                `form:"md5" binding:"required"`                                    // 文件md5
	Region  string                `form:"region" binding:"required,oneof=miajiodb beijing hangzhou"` // 文件上传地区
	Name    string                `form:"name" binding:"required,gte=2,lte=32"`                      // 圈子名称
	Descibe string                `form:"describe" binding:"required,gte=2,lte=255"`                 // 圈子描述
}

// Create 创建圈子
// 用户圈子个数依据字典 USER_VIP_CIRCLES 进行等级判断,对应等级用户可生成的圈子个数依据此字典设定, 当字典值为-1时将表示无限制
func (*circlesLogic) Create(ctx *gin.Context) {
	var req createCirclesRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	file, err := SystemFileLogic.UploadLogic(mo.AccountInfo.ID, "CIRCLES_LOGO", fileDefaultUploadRequest{File: req.File, MD5: req.MD5, Region: req.Region}, ".jpg", ".jpeg", ".png", ".svg", ".webp")
	if !lib.ServerFail(ctx, err) {
		return
	}

	maxCirclesNumStr, err := store.SystemDictionaryStore.FindGroupAndValByName("USER_VIP_CIRCLES", mo.DetailedInfo.Vip)
	if err != nil && err.Error() == "sql: no rows in result set" {
		maxCirclesNumStr = "-1"
	} else if !lib.ServerFail(ctx, err) {
		return
	}
	maxCirclesNum, err := strconv.Atoi(maxCirclesNumStr)
	if !lib.ServerFail(ctx, err) {
		return
	}

	if maxCirclesNum != -1 {
		userNeedCircles, err := store.CirclesStore.FindById(mo.AccountInfo.ID)
		if !lib.ServerFail(ctx, err) {
			return
		}
		if len(userNeedCircles) >= maxCirclesNum {
			lib.ServerFailf(ctx, 500, "创建失败", errors.New("当前用户已达到最大创建圈子数"))
			return
		}
	}

	circles := model.CirclesInfo{
		Logo:     file.ID,
		Name:     req.Name,
		Descirbe: req.Descibe,
		CreateBy: mo.AccountInfo.ID,
		Owner:    mo.AccountInfo.ID,
	}
	id, err := store.CirclesStore.Create(circles)
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "创建成功", id)
}

// Find 查询自己所拥有的圈子
func (*circlesLogic) Find(ctx *gin.Context) {
	mo, ok := store.TokenGet(ctx)
	if !ok {
		return
	}
	circles, err := store.CirclesStore.FindById(mo.AccountInfo.ID)
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "查询成功", circles)
}
