package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	// 查询文件夹是否存在
	parentData := new(models.UserRepository)
	err = l.svcCtx.Engine.Debug().Select("id").
		Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).
		First(parentData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文件夹不存在")
		}
		return
	}
	// 更新记录的 parentId
	err = l.svcCtx.Engine.Debug().Where("identity = ? and user_identity = ?", req.Identity, userIdentity).Updates(models.UserRepository{
		ParentId: int64(parentData.ID),
	}).Error
	return
}
