package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 1. 从数据库查询当前用户
	user := new(models.UserBasic)
	err = models.Engine.Where("name = ?", req.Name).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名不存在")
		}
		return nil, err
	}
	if helper.Md5(req.Password) != user.Password {
		return nil, errors.New("密码错误")
	}
	// 2. 生成 token
	token, err := helper.GenerateToken(user.ID, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	return &types.LoginReply{Token: token}, nil
}
