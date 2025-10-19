# 使用官方的Go镜像作为构建环境
FROM golang:1.22.2-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jwt-secret-fake-random-goat .

# 使用更小的alpine镜像作为运行环境
FROM alpine:latest

# 安装ca证书（如果需要HTTPS）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从builder阶段复制构建好的二进制文件
COPY --from=builder /app/jwt-secret-fake-random-goat .

# 暴露默认端口
EXPOSE 10086

# 设置启动命令
CMD ["./jwt-secret-fake-random-goat", "server", "--port", "10086"]