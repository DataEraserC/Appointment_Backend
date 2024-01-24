# 构建说明

## 如何使用golang国内镜像

```shell
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
```

## 如何进入nix的go基础开发环境

> 实际上如果同时安装了direnv和nix,进入文件夹时会自动执行nix develop

```shell
nix develop
# 等于 nix develop .#default
```

## 如何编译/运行

- 编译

```shell
# 本机平台
go build -ldflags "-s -w" -o Appointment -trimpath
# x64 Linux 平台 如各种云服务器
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o Appointment -trimpath
# x64 Windows 平台 如大多数家用电脑
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o Appointment.exe -trimpath
# armv6 Linux 平台 如树莓派 zero W
GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -ldflags "-s -w" -o Appointment -trimpath
# mips Linux 平台 如 路由器 wndr4300
GOOS=linux GOARCH=mips GOMIPS=softfloat CGO_ENABLED=0 go build -ldflags "-s -w" -o Appointment -trimpath
```

- 编译并运行(不显式产生二进制文件)

```shell
go run .
```

- 使用nix 编译并运行

```shell
nix run
# 等于 nix run .#default
```

- 进入一个仅编译并加入PATH的shell环境(退出后无法访问编译结果)

```shell
nix shell
# 等于 nix shell .#default
```

## 如何更新依赖

1. go mod 的更新

```shell
go get -u
# 等于 go get -u .
# 更新当前目录

# 或者 循环更新所有目录
# go get -u ./...

go mod tidy
# 用于确保 go.mod 与模块中的源代码匹配
```

2. nix 项目管理也需要更新

```shell
nix develop
# 进入nix管理的go devshell

gomod2nix
# 转换go.mod至gomod2nix.toml
```
