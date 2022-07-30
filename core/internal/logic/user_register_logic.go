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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// 验证code是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱的验证码为空")
	}
	if code != req.Code {
		return nil, errors.New("验证码错误")
	}
	// 检查用户名是否存在
	user := new(models.UserBasic)
	result := l.svcCtx.Engine.Where("name = ?", req.Name).First(&user)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return
	}
	if result.RowsAffected > 0 {
		return nil, errors.New("用户已存在")
	}
	// 数据入库
	user = &models.UserBasic{
		Identity: helper.GenerateUUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	err = l.svcCtx.Engine.Create(user).Error
	return
}
