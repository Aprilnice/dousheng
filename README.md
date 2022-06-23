# 抖声 v1.0
默认分支beta为最终分支，dev分支已弃用
## 一、项目功能
1. 具有登录和游客两种模式，下拉获取视频流，返回按投稿时间倒序视频列表
2. 新用户可通过用户名、密码注册以及登录
3. 可获取登录用户的详细信息（用户名、关注数、粉丝数、投稿列表、关注列表、粉丝列表）
4. 用户登录后可上传视频
5. 登录用户可对视频点赞或取消点赞以及评论
6. 登录用户可获取所有已点赞视频列表
7. 用户可获取视频的所有评论，按发布时间倒序
8. 登录用户可关注或取消关注其他用户
## 二、项目架构
1. 使用RestfulAPI与前端对接
2. 具体服务通过go-micro微服务框架实现gRPC协议
3. 使用proto3序列化协定与服务进行通信
4. 使用etcd实现服务发现
5. 服务拆分
   - 用户服务
   - 视频服务
   - 评论服务
   - 点赞服务
   - 关注服务
6. 中间件：
   - JWT鉴权
   - 密码md5加密
   - token生成
   - 雪花算法生成唯一id
7. 数据存储 
   - 视频文件：
     - 本地保存
   - 结构化数据：
     - MySQL存储
     - Redis存储
## 三、项目结构

| package | package     | introduction |
|---------|-------------|--------------|
| cmd     |             | 核心代码         |
|         | api         | 获取请求，实例化服务   |
|         | comment     | 评论服务具体实现     |
|         | favourite   | 点赞服务具体实现     |
|         | relation    | 关注服务具体实现     |
|         | video       | 视频服务具体实现     |
|         | user        | 用户服务具体实现     |
| config  |             | 配置文件         |
|         | config.yaml | 配置文件         |
|         | config.go   | 配置文件读取       |
| file    |             | 文件存储         |
|         | video       | 视频文件存储       |
|         | image       | 视频封面文件存储     |
| pkg     |             | 公共包          |
|         | errdeal     | 错误处理         |
|         | constant    | 一些常量         |
|         | ffmpeg      | 视频关键帧截取作为封面  |
|         | log         | 日志           |
|         | middleware  | 中间件          |
|         | snowflaker  | ID生成器        |
|         | doushengjwt | 用户鉴权相关       |

## 四 运行
    Windows 用户需要下载ffmpeg可执行文件，放在项目目录下。
- `git clone https://github.com/Aprilnice/dousheng.git`
- `在config.yaml中配置好自己的参数`
- `分别执行cmd目录下api comment favorite video user relation服务`

## 五 项目展示
[项目展示文档](https://bytedancecampus1.feishu.cn/docx/doxcnHUUV6ZiJyWny6ohAPyhnob)

## 鸣谢

[字节后端青训营](https://youthcamp.bytedance.com/)


