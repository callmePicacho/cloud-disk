package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// 先找数据库中是否已存在 UserIdentity - parentId - name 相等的记录，如果已有，更新 RepositoryIdentity，否则插入记录
	err = l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ? AND name = ? AND ext = ?", req.ParentId, UserIdentity, req.Name, req.Ext).
		Assign(&models.UserRepository{
			Identity:           helper.GenerateUUID(),
			UserIdentity:       UserIdentity,
			ParentId:           req.ParentId,
			RepositoryIdentity: req.RepositoryIdentity,
			Ext:                req.Ext,
			Name:               req.Name,
		}).
		FirstOrCreate(&models.UserRepository{}).Error
	return
}
