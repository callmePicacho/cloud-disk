package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	// 查询分享文件是否存在
	ur := new(models.UserRepository)
	err = l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).First(ur).Error
	if err != nil {
		return
	}
	// 插入分享记录
	data := &models.ShareBasic{
		Identity:               helper.GenerateUUID(),
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	err = l.svcCtx.Engine.Create(data).Error
	return &types.ShareBasicCreateReply{Identity: data.Identity}, err
}
