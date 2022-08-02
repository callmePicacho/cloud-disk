package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 参数检查
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}
		if r.PostForm.Get("uploadId") == "" {
			httpx.Error(w, errors.New("uploadId is empty"))
			return
		}
		if r.PostForm.Get("partNumber") == "" {
			httpx.Error(w, errors.New("partNumber is empty"))
			return
		}

		// 分片文件上传
		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		// 返回 MD5 值
		resp = &types.FileUploadChunkReply{
			Etag: etag,
		}
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
