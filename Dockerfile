# 使用官方的 Golang 镜像作为构建环境
FROM golang:latest AS builder

# 设置工作目录为容器的工作空间
WORKDIR /usr/src/app

# 将本地的包文件复制到容器的工作空间
COPY . .

# 编译 Go 应用
RUN go build -tags netgo -o main ./main.go

EXPOSE 8080

CMD ["/usr/src/app/main"]
