package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RefreshAuthorization 刷新 token
// 1. 此时传入的 token 应该是 login 接口返回的 refresh token
// 2. 解析 refresh token 中的用户信息重新生成 token 和 refresh token
func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationReply, err error) {
	// 解析 token
	uc, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return
	}
	// 生成 token
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	// 生成 refresh token
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.RefreshTokenExpire)
	return &types.RefreshAuthorizationReply{
		Token:        token,
		RefreshToken: refreshToken,
	}, err
}
