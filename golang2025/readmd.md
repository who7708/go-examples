# 初始化
go mod init golang2025/hello

# 创建文件
go mod tidy

# 运行
go run main.go



go help


go get rsc.io/quote


go get golang.org/x/tools/cmd/stringer
go list -f '{{.Dir}}'
go list -f '{{.Dir}}' rsc.io/quote
go list -f '{{.Target}}' rsc.io/quote

# 找到 Go 安装路径，go命令将在其中安装当前包
go list -f '{{.Target}}'


# go env 命令设置 GOBIN 变量来更改安装目标：
go env -w GOBIN=/tmp/bin
