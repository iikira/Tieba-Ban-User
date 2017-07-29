# Tieba-Ban-User

## 功能

贴吧: 支持吧务封禁用户

## 如何使用
---
无需安装Go语言即可下载使用 [点此查看发布页](https://github.com/iikira/Tieba-Ban-User/releases)

在 (Linux, MacOS)终端／(Windows)命令提示符 中运行。

---

## 如何手动编译安装
---
### 1. 安装Go语言运行环境

* 访问 [Go语言官网](https://golang.org) 下载安装Golang
* 设置GOPATH环境变量

Linux: 
```shell
export GOPATH=/path/to/your/gopath
```
Windows:
```shell
set GOPATH=C:\path\to\your\gopath
```

### 2. 安装

* 编译安装(需要设置GOPATH环境变量)
```shell
go get -u -v github.com/iikira/Tieba-Ban-User
```
> 编译生成的文件在GOPATH的bin目录下

* 手动编译安装(需要设置GOPATH环境变量)

1. 下载源码到源码目录
2. 安装依赖包

```shell
go get -u -v github.com/bitly/go-simplejson
go get -u -v github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil
```

3. 进入源码目录，编译

```shell
go build
```
---