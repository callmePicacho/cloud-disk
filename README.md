# Cloud-disk

> 轻量级云盘系统，技术栈包括：go-zero、gorm、redis、COS

涉及的相关文档包括：
> go-zero 文档：https://go-zero.dev/cn/  
> GORM 文档：https://gorm.io/zh_CN/docs/index.html  
> Redis 文档：https://redis.io/docs/  
> 腾讯云 COS 后台地址：https://console.cloud.tencent.com/cos/bucket  
> 腾讯云 COS 帮助文档：https://cloud.tencent.com/document/product/436/31215  

用到的库包括：
> Go 邮箱库：https://github.com/jordan-wright/email  
> Go Redis库：https://github.com/go-redis/redis  
> Go UUID 库：https://github.com/satori/go.uuid

使用到的命令：
```shell
# 创建 API 服务
goctl api new core
# 使用 .api 文件生成代码
goctl api go -api core.api -dir . -style go_zero
# 启动服务
go run .\core.go -f .\etc\core-api.yaml
```

实现功能包括：
1. 用户模块
   1. 用户账号密码登陆
   2. Authorization 刷新
   3. 邮箱验证码注册
   4. 用户详情
2. 存储池模块
   1. 中心存储池资源管理
      1. 文件上传
      2. 文件秒传
      3. 文件分片上传
   2. 个人存储池资源管理
      1. 文件关联存储
      2. 文件列表
      3. 文件名称修改
      4. 文件夹创建
      5. 文件删除
      6. 文件移动
3. 文件分享模块
   1. 创建分享记录
   2. 获取资源详情
   3. 分享资源保存


食用建议：
1. 参考每个 commits
2. api 文件定义接口和参数
3. 使用 `goctl api` 命令生成代码
4. 验证引用库/函数先开 test 跑通，再考虑如何集成到业务代码
5. 在 handler 或 logic 目录下对应文件写接口逻辑层代码
6. 完成编码后，postman 自测通过