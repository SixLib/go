# go
## dev 运行
```sh
go run *.go
```

## 编译 windows 二进制文件
**GOOS** 目标可执行文件运行操作系统 支持 `darwin` `freebsd` `linux` `windows`
**GOARCH** 目标可执行程序操作系统架构 包括 386 amd64 arm
``` sh
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```