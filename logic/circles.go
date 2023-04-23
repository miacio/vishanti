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
	Create(*gin.Context)      // Create 创建圈子
	Find(*gin.Context)        // Find 查询自己所拥有的圈子
	InviteJoin(*gin.Context)  // InviteJoin 邀请用户加入圈子
	RequestJoin(*gin.Context) // RequestJoin 用户申请加入圈子
	FindMyJoin(*gin.Context)  // FindMyJoin 查询当前用户加入中的圈子信息
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
		userNeedCircles, err := store.CirclesStore.FindByUserId(mo.AccountInfo.ID)
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
	circles, err := store.CirclesStore.FindByUserId(mo.AccountInfo.ID)
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "查询成功", circles)
}

// inviteJoinRequest 邀请用户加入圈子请求
type inviteJoinRequest struct {
	CirclesId   string `form:"circles_id" json:"circles_id" binding:"required"`     // 加入的圈子id
	UserAccount string `form:"user_account" json:"user_account" binding:"required"` // 受邀用户的账号
}

// InviteJoin 邀请用户加入圈子
// 提供用户的账号,系统将对应用户拉入圈子用户,状态为邀请中,只有用户同意加入该圈子时,圈子用户信息才能生效
func (*circlesLogic) InviteJoin(ctx *gin.Context) {
	var req inviteJoinRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	loginUser, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	// 校验当前用户是否在圈子中,如果不在,将无法邀请用户
	userIsInCircles := false

	circlesInfos, err := store.CirclesStore.FindByUserId(loginUser.AccountInfo.ID)
	if !lib.ServerFail(ctx, err) {
		return
	}
	// 当前用户是该圈子的拥有者,将拥有拉取用户的权力
	for _, circirclesInfo := range circlesInfos {
		if circirclesInfo.ID == req.CirclesId {
			userIsInCircles = true
			break
		}
	}
	if !userIsInCircles {
		// 当前用户加入的圈子列表
		loginUserJoinCircles, err := store.CirclesUsersStore.FindByUserId(loginUser.AccountInfo.ID)
		if !lib.ServerFail(ctx, err) {
			return
		}
		for _, userCircles := range loginUserJoinCircles {
			if userCircles.CirclesID == req.CirclesId {
				userIsInCircles = true
				break
			}
		}
	}

	if !userIsInCircles {
		lib.ServerResult(ctx, 500, "当前用户未加入该圈子,没有邀请其他用户权限", nil, nil)
		return
	}

	userAccount, err := store.UserStore.FindByAccount(req.UserAccount)
	if !lib.ServerFail(ctx, err) {
		return
	}
	if userAccount.ID == "" {
		lib.ServerFailf(ctx, 500, "受邀用户不存在", nil)
		return
	}

	_, err = store.CirclesUsersStore.Create(model.CirclesUsers{
		CirclesID: req.CirclesId,
		UserID:    userAccount.ID,
		IsSignOut: "1",
	})
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "邀请成功", nil)
}

// requestJoinRequest 用户申请加入圈子请求
type requestJoinRequest struct {
	CirclesId string `form:"circles_id" json:"circles_id" uri:"circles_id" binding:"required"` // 申请加入的圈子id
}

// RequestJoin 用户申请加入圈子
func (*circlesLogic) RequestJoin(ctx *gin.Context) {
	var req requestJoinRequest
	if !lib.ShouldBind(ctx, &req) {
		return
	}

	loginUser, ok := store.TokenGet(ctx)
	if !ok {
		return
	}

	_, err := store.CirclesUsersStore.Create(model.CirclesUsers{
		CirclesID: req.CirclesId,
		UserID:    loginUser.AccountInfo.ID,
		IsSignOut: "2",
	})
	if !lib.ServerFail(ctx, err) {
		return
	}
	lib.ServerSuccess(ctx, "申请成功", nil)
}

// FindMyJoin 查询当前用户加入中的圈子信息
func (*circlesLogic) FindMyJoin(ctx *gin.Context) {
	loginUser, ok := store.TokenGet(ctx)
	if !ok {
		return
	}
	loginUserJoinCircles, err := store.CirclesUsersStore.FindByUserId(loginUser.AccountInfo.ID)
	if !lib.ServerFail(ctx, err) {
		return
	}
	ids := make([]string, 0)
	for i := range loginUserJoinCircles {
		ids = append(ids, loginUserJoinCircles[i].CirclesID)
	}
	if len(ids) > 0 {
		circlesInfos, err := store.CirclesStore.FindByIds(ids...)
		if !lib.ServerFail(ctx, err) {
			return
		}
		lib.ServerSuccess(ctx, "查询成功", circlesInfos)
		return
	}
	lib.ServerSuccess(ctx, "当前用户暂未加入圈子", nil)
}
