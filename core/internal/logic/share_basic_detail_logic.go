package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	// 查询用户分享的文件是否还在
	var cnt int64
	err = l.svcCtx.Engine.Model(&models.UserRepository{}).
		Where("identity = (select user_repository_identity from share_basic sc where sc.identity = ?)", req.Identity).Count(&cnt).Error
	if err != nil {
		return
	}
	// 对分享记录的点击次数 + 1
	err = l.svcCtx.Engine.Exec("update share_basic set click_num = click_num+1 where identity = ?", req.Identity).Error
	if err != nil {
		return
	}
	// 获取资源详情
	resp = new(types.ShareBasicDetailReply)
	err = l.svcCtx.Engine.Table("share_basic sb").
		Select("ur.repository_identity, ur.name, ur.ext, rp.size, rp.path").
		Joins("LEFT JOIN user_repository ur on sb.user_repository_identity = ur.identity").
		Joins("LEFT JOIN repository_pool rp on ur.repository_identity = rp.identity").
		Where("sb.identity = ?", req.Identity).
		Take(resp).Error
	return
}
