package logic

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/miacio/vishanti/lib"
	"github.com/miacio/vishanti/model"
	"github.com/miacio/vishanti/store"
)

type systemDictionaryLogic struct{}

type ISystemDictionaryLogic interface {
	FindByGroup(ctx *gin.Context) // 依据组获取字典列表
	Inserts(ctx *gin.Context)     // 批量写入字典
}

var SystemDictionaryLogic ISystemDictionaryLogic = (*systemDictionaryLogic)(nil)

type (
	// systemDictionaryInsertRequest 字典写入请求结构体
	systemDictionaryInsertRequest struct {
		Name        string `json:"name" form:"name" uri:"name" binding:"required"`      // Name name 名称
		Group       string `json:"group" form:"group" uri:"group" binding:"required"`   // Group group 组名
		ParentGroup string `json:"parent_group" form:"parent_group" uri:"parent_group"` // ParentGroup parent_group 上级组名
		Describe    string `json:"describe" form:"describe" uri:"describe"`             // Describe describe 描述
		Val         string `json:"val" form:"val" uri:"val" binding:"required"`         // Val val 值
		CreateBy    string `json:"create_by" form:"create_by" uri:"create_by"`          // CreateBy create_by 创建人id
	}

	systemDictionaryBatchRequest []systemDictionaryInsertRequest // 字典批量写入
)

// ToModels 转换成字典结构体列表
func (ts systemDictionaryBatchRequest) ToModels() ([]model.SystemDictionary, error) {
	result := make([]model.SystemDictionary, 0)
	for i := range ts {
		t := ts[i]

		if t.Name == "" {
			return nil, errors.New("名称不能为空")
		}

		if t.Val == "" {
			return nil, errors.New("值不能为空")
		}

		if t.Group == "" {
			return nil, errors.New("组名不能为空")
		}

		result = append(result, model.SystemDictionary{
			ID:          lib.UID(),
			Name:        t.Name,
			Group:       t.Group,
			ParentGroup: t.ParentGroup,
			Describe:    t.Describe,
			Val:         t.Val,
			CreateBy:    t.CreateBy,
		})
	}
	return result, nil
}

// 依据组获取字典列表
func (*systemDictionaryLogic) FindByGroup(ctx *gin.Context) {
	group := ctx.Query("group")
	systemDictionaryGroup, err := store.SystemDictionaryStore.FindByGroup(group)
	if !lib.ServerFailf(ctx, 500, "获取失败", err) {
		return
	}
	lib.ServerSuccess(ctx, "获取成功", systemDictionaryGroup)
}

// 批量写入字典
func (*systemDictionaryLogic) Inserts(ctx *gin.Context) {
	var req systemDictionaryBatchRequest
	if !lib.ShouldBindJSON(ctx, &req) {
		return
	}

	models, err := req.ToModels()
	if !lib.ServerFailf(ctx, 400, "参数错误", err) {
		return
	}

	err = store.SystemDictionaryStore.Inserts(models)
	if !lib.ServerFailf(ctx, 500, "写入失败", err) {
		return
	}
	lib.ServerSuccess(ctx, "写入成功", nil)
}
