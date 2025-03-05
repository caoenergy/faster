## faster

Faster 是一个极简命令行工具; 用于获取 GitHub | Docker 加速访问地址、另外可直接拉取 Docker 镜像

```
NAME:
   faster - A new cli application

USAGE:
   faster [global options] [command [command options]]

VERSION:
   1.0.0

DESCRIPTION:
   Faster 是一个极简命令行工具; 用于获取 GitHub | Docker 加速访问地址、另外可直接拉取 Docker 镜像

COMMANDS:
   github   获取 GitHub 加速访问地址
   docker   获取 Docker 加速访问地址
   pull     拉取 Docker 镜像
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```
faster github  				   ## 获取 github 加速地址
faster github download xxx 	## 下载 github 上的文件
faster github clone xxx    	## 克隆 github 上的仓库

faster docker				      ## 获取 docker 加速地址
faster pull xxx				   ## 拉取 docker 镜像
```
