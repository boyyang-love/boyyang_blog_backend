## 个人博客网站（go-zero单体服务）

### go-zero
1. [go-zero文档地址](https://go-zero.dev/cn/)

### etc
1. etc 下的blog-api-c.yaml文件替换成blog-api.yaml
2. 关于数据库以及对象存储需要替换成自己的相关信息
```yaml
Name: blog-api
Host: 0.0.0.0
Port: 8888
MaxBytes: 10240000
Auth:
  AccessSecret: $AccessSecret
  AccessExpire: $AccessExpire
Mysql:
  Host: $host
  Port: $port
  Database: 数据库名称
  Username: 用户名
  Password: 密码
  Charset: utf8
  Timeout: 10s
CloudBase:
  ClientUrl: 对象存储连接
  ClientSecretId: 对象存储SecretId
  ClientSecretKey: 对象存储key
AppidAndSecret:
  AppId: $appid // 微信小程序appid
  Secret: $secret // 密钥
```
### gorm 
1. orm框架使用的是gorm
2. [gorm文档地址](https://learnku.com/docs/gorm/v2)

### 实现功能
1. 登录注册 √
2. 修改,获取,用户信息 √
3. cos文件上传 √
4. 图片墙上传,修改,收藏 √
5. 博客发布,修改,点赞,评论,删除 √

### 后续新增功能
1. 留言板
2. 聊天室
3. 时间轴
4. 热度博客
5. 用户关注
### 接口在线文档
[接口文档地址](https://console-docs.apipost.cn/preview/03ccd55c68247833/3eb25d53d437d3ea)
### 前后端代码
1. [后端gitHub仓库地址](https://github.com/boyyang-love/boyyang_blog_backend)
2. [前端gitHub仓库地址](https://github.com/boyyang-love/boyyang_blog)
3. [前端页面在线预览地址](https://prod-2g5hif5wbec83baa-1301921121.tcloudbaseapp.com)
