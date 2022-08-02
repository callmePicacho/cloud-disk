package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"
	"gorm.io/gorm"
	"path"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.RepositoryPool)
	err = l.svcCtx.Engine.Where("hash = ?", req.Md5).First(rp).Error
	// 未找到，上传到 cos 后返回 uploadId 和 key
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 获取该文件的 uploadId 和 key，准备进行文件分片上传
		key, uploadId, err := helper.CosInitPart(path.Ext(req.Name))
		return &types.FileUploadPrepareReply{Key: key, UploadId: uploadId}, err
	}

	// 已经存在该文件，不必重复上传 cos，直接返回 identity
	return &types.FileUploadPrepareReply{Identity: rp.Identity}, err
}
