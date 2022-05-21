# 抖声
1. 使用proto3
2. gRPC框架
   - 用户服务：
     - 鉴权：jwt
     - 密码加密：md5 // v2.0反爬
     - token登录（7天过期）
     - 限制同一时间只能一台设备登录（已登录信息存储在redis）
     - 密钥（douyin）
     - gorm
   - 视频服务：
     - 实现播放器
     - redis维护一个视频列表（定期更新，设置过期时间）以及视频详细信息
     - 封面（图床）
     - 投稿（文件上传）
   - 交互服务：
     - gorm
     - redis存储评论信息，点赞列表，关注列表
3. 高可用性
   - 消息队列削峰/负载均衡
   - **数据操作一定要加锁**



用户服务 1人

视频服务+redis定期向MySQL刷盘  2人

交互服务+消息队列/负载均衡 1人





项目结构

| package     | package        | function         |
| ----------- | -------------- | ---------------- |
| config      |                | 配置文件         |
| interaction |                | 交互服务         |
|             | cmd            | 交互服务具体实现 |
|             | pb             | proto文件        |
|             | interaction.go | 交互服务入口     |
| user        | cmd            |                  |
|             | pb             |                  |
|             | user.go        | 用户服务入口     |
| video       | cmd            |                  |
|             | pb             |                  |
|             | video.go       | 视频服务入口     |
| pkg         |                | 公共包           |
|             | errordeal      | 错误处理         |
|             | etcd           |                  |
|             | log            | 日志             |
|             | middleware     | 中间件           |
|             | mysql          |                  |
|             | redis          |                  |
