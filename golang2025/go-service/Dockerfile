# 使用官方 Golang 镜像构建 Go 管理程序
FROM golang:1.21 AS builder

WORKDIR /app

# 拷贝 go.mod 和源码
COPY go.mod ./
COPY . ./

# 构建 Go 可执行文件
RUN go build -o go-spring-boot-manager main.go

# 生产环境镜像
FROM debian:bookworm-slim

WORKDIR /app

# 拷贝编译好的可执行文件和 Spring Boot jar 包
COPY --from=builder /app/go-spring-boot-manager /app/go-spring-boot-manager
COPY ./springboot.jar /app/springboot.jar

# （可选）拷贝静态资源，如果有
COPY ./internal/web/static /app/internal/web/static

# 暴露 Web 端口（Gin 的管理界面）
EXPOSE 8888

# 启动 Go 管理器
CMD ["./go-spring-boot-manager", "./springboot.jar"]