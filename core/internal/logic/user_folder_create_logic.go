package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	// 检查文件同级目录下是否存在相同文件夹名称
	var cnt int64
	err = l.svcCtx.Engine.Model(&models.UserRepository{}).Where("user_identity = ? and name = ? and parent_id = ?", userIdentity, req.Name, req.ParentId).
		Count(&cnt).Error
	if err != nil {
		return
	}
	if cnt > 0 {
		return nil, errors.New("该名称已存在")
	}
	// 创建文件夹
	data := &models.UserRepository{
		Identity:     helper.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	err = l.svcCtx.Engine.Create(data).Error
	return
}
