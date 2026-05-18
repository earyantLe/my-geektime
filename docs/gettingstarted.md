# 快速开始 ⚡️

## 安装
* 目前支持三种安装方式, 首推 docker compose 方式
* 如果，你的本地已经有redis，mysql等服务了，也可以考虑docker方式和二进制发行包方式

#### docker compose 方式，推荐指数 🌟🌟🌟🌟🌟

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
    src="../images/sponsor.jpg"
    width="356px"
  />
</picture>


#### docker 方式，推荐指数 🌟🌟🌟
使用宿主机目录替换下面的 ${directory}
```shell
docker run -d --restart always \
-p 8090:8090 \
--name mygeektime \
-v ${directory}:/repo  \
-v config.yml:/config.yml \
zkep/mygeektime:latest  server --config=config.yml
```
浏览器访问:  http://127.0.0.1:8090


#### [二进制包安装](https://github.com/zkep/my-geektime/releases) ，推荐指数 🌟🌟

下载对应操作系统的二进制包，下面以MacOS为例
```shell
# 下载
wget https://github.com/zkep/my-geektime/releases/download/v0.0.6/mygeektime_Darwin_arm64.tar.gz
# 解压
tar -zxvf my-geektime_Darwin_arm64.tar.gz

# 切入解压目录
cd mygeektime_Darwin_arm64

# 默认配置启动服务
./mygeektime server 

# 执行生成自定义配置模版命令，会生成 config_templete.yml文件，自行修改配置内容 
./mygeektime cli config

# 自定义配置启动服务
./mygeektime server --config=config_templete.yml

```
二进制方式，缺少一些依赖项，比如最重要的是ffmpeg，用于音视频合成的，需要自行安装后，加入环境变量中
#### [FFmpeg 处理视频](https://ffmpeg.org/download.html)
MacOS
```shell
brew install ffmpeg        
```
Linux
```shell
git clone https://github.com/FFmpeg/FFmpeg.git ffmpeg

cd ffmpeg

./configure --enable-gpl --enable-libx264

make && make install
```

### 安装python环境依赖，mkdocs 用于生成 markdown文档

```shell
pip install mkdocs-material
```


[配置项](./config.md)  👉