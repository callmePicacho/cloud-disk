package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	// 检查文件同级目录下是否存在相同文件名称
	var cnt int64
	err = l.svcCtx.Engine.Model(&models.UserRepository{}).Where("user_identity = ? and name = ? and parent_id = (select parent_id from user_repository ur where ur.identity = ?)", userIdentity, req.Name, req.Identity).
		Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	// 修改文件名称
	err = l.svcCtx.Engine.Model(&models.UserRepository{}).
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Update("name", req.Name).Error
	return
}
