# Cloud-disk

> 轻量级云盘系统，技术栈包括：go-zero、gorm、redis、COS

GORM 文档：https://gorm.io/zh_CN/docs/index.html  
go-zero 文档：https://go-zero.dev/cn/ 
腾讯云COS后台地址：https://console.cloud.tencent.com/cos/bucket

启动服务
```text
go run .\core.go -f .\etc\core-api.yaml
```

使用 api 文件生成接口代码
```text
goctl api go -api core.api -dir . -style go_zero
```