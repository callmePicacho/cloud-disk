package logic

import (
	"cloud-disk/core/define"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Size
	if page == 0 {
		page = 1
	}
	// 分页偏移量
	offset := (page - 1) * size

	// 查询列表
	userList := make([]*types.UserFile, 0)
	err = l.svcCtx.Engine.Table("user_repository u").
		Joins("LEFT JOIN repository_pool p on u.repository_identity = p.identity").
		Select("u.id, u.identity, u.repository_identity, u.name, u.ext, p.path, p.size").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Limit(size).
		Offset(offset).
		Find(&userList).Error
	if err != nil {
		return
	}
	var cnt int64
	err = l.svcCtx.Engine.Table("user_repository u").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Count(&cnt).Error
	return &types.UserFileListReply{List: userList, Count: cnt}, err
}
