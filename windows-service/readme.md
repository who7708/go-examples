# 安装rsrc工具
go install github.com/akavel/rsrc@latest

# 在项目目录下，生成 .syso 资源文件
rsrc -manifest app.manifest -o app.syso



-ldflags="-linkmode internal"  -ldflags="-H windowsgui"