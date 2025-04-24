# 构建阶段
FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /builder
COPY . .

# 设置代理和依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o randomimage main.go && chmod +x randomimage

# 最小镜像阶段
FROM alpine:latest
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update && apk add --no-cache tzdata && rm -rf /var/cache/apk/* \
    && mkdir /app &&ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /app

# 拷贝可执行文件
COPY --from=builder /builder/randomimage /app/

EXPOSE 8080
CMD ["/app/randomimage"]