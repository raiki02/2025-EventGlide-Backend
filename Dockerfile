# 第一阶段：构建 Go 应用
FROM golang:alpine AS builder

# 设置 Go 代理为七牛云的代理
ENV GOPROXY=https://goproxy.cn,direct

# 安装需要的依赖
RUN apk update && apk add --no-cache git

# 切换到 app 目录，构建二进制文件
WORKDIR /app

# 复制 be-ccnu 服务代码
COPY . /app

RUN go mod tidy && go build -o app

# 第二阶段：复制编译结果到最终镜像
FROM alpine

# 安装 tzdata 来设置时区
RUN apk add --no-cache tzdata

# 设置时区为 Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录为
WORKDIR /app

# 从 builder 复制编译好的二进制文件
COPY --from=builder /app/app .

# 开放端口（根据需要设置）
EXPOSE 8080

# 启动用户服务
CMD ["./app"]
