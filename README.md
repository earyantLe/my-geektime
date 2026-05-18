 [English](./README_US.md) | 中文
 

## 我的极客时间

不仅是下载器，更是在线文档

 * 支持极客时间VIP账号一次缓存数据，永久观看
 * 支持一键发布整个课程为在线文档
 * 支持一键下载整个课程音视频资源到本地目录



### [项目文档](https://zkep.github.io/my-geektime/) | [在线体验](http://8.141.6.243:8090)


### 安装

1. 下载项目

```shell
git clone https://github.com/zkep/my-geektime.git
```
2. 获取镜像 （任选其一）
- 拉取仓库镜像
```shell
# 拉取linux/amd64架构的镜像
docker pull --platform=linux/amd64  zkep/mygeektime:latest

# 拉取linux/arm64架构的镜像
# docker pull --platform=linux/arm64  zkep/mygeektime:latest
```
- 本地构建镜像
```shell
docker build --platform linux/amd64  -t zkep/mygeektime:latest .
```

3. 启动服务
```shell
cd my-geektime/docker

docker-compose -f docker-compose.yml up -d
```

浏览器访问:  http://127.0.0.1:8090


#### 微信赞赏

如若有用，不吝赞赏👏

赞赏并留言 <b>邮箱账号</b>，回赠<b> [数据库](https://zkep.github.io/my-geektime/guide/data_default/) </b>

<picture>
  <img
    alt="sponsor"
    src="docs/images/sponsor.jpg"
    width="356px"
  />
</picture>


