package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// FileUploadHandler
// 1. 解析读取文件
// 2. 计算 Hash 判断文件是否已存在，如果有相同 hash，直接返回已有文件 Identity
// 3. 文件上传到 cos
// 4. 落库文件信息
func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 解析读取文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 判断是否已在cos中存在
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		err = svcCtx.Engine.Select("identity, name, ext").Where("hash = ?", hash).First(&rp).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			httpx.Error(w, err)
			return
		}
		if rp.Identity != "" { // 已存在直接返回Identity
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		// 上传到cos
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 赋值文件信息并落库
		req = types.FileUploadRequest{
			Hash: hash,
			Name: fileHeader.Filename,
			Ext:  path.Ext(fileHeader.Filename),
			Size: fileHeader.Size,
			Path: cosPath,
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
