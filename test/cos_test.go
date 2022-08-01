package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func NewClient() *cos.Client {
	u, _ := url.Parse(define.TencentCosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	return client
}

func TestFileUploadByFilepath(t *testing.T) {
	client := NewClient()

	key := "test.png"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/test.png", nil,
	)
	if err != nil {
		t.Log(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	client := NewClient()

	key := "test.png"

	f, err := os.ReadFile("./img/test.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Log(err)
	}
}

// 分片上传初始化 获取对应的 uploadId
func TestInitChunkUpload(t *testing.T) {
	client := NewClient()

	name := "chunk.png"
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID)
}

// 分片上传 使用初始化获得的 uploadId 来上传分片
func TestChunkUpload(t *testing.T) {
	client := NewClient()

	// 分片初始化获得 UploadID
	UploadID := "1659337315a0dcd2aa9fbc08efd344f3f83336f2f2fb803f3c7d550f76bd3133a87c37635b"
	key := "chunk.png"
	f, err := os.ReadFile("./0.chunk")
	if err != nil {
		t.Fatal(err)
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag") // ETag 分块的 Md5 值
	fmt.Println(PartETag)
}

// 分片上传完成
func TestChunkUploadComplete(t *testing.T) {
	client := NewClient()

	UploadID := "1659337315a0dcd2aa9fbc08efd344f3f83336f2f2fb803f3c7d550f76bd3133a87c37635b"
	key := "chunk.png"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts,
		// 使用上传 partNum 分片时得到的 ETag 填充
		cos.Object{PartNumber: 1, ETag: "41ac7a05cdc4998dcd3a2e69123db964"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
