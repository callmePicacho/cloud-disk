package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	// 获取资源信息
	rp := new(models.RepositoryPool)
	err = l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).First(rp).Error
	if err != nil {
		return
	}
	// 用户资源保存
	ur := &models.UserRepository{
		Identity:           helper.GenerateUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: rp.Identity,
		Ext:                rp.Ext,
		Name:               rp.Name,
	}
	err = l.svcCtx.Engine.Create(ur).Error
	return
}
