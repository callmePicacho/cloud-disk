package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	// 查询该邮箱是否被注册
	user := new(models.UserBasic)
	result := models.Engine.Where("email = ?", req.Email).First(&user)
	if result.Error != err && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if result.RowsAffected != 0 {
		return nil, errors.New("该邮箱已被注册")
	}
	// 获取验证码
	code := helper.GenerateVerifyCode()
	// 存储验证码
	models.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	// 发送验证码
	err = helper.MailSendCode(req.Email, code)
	return
}
